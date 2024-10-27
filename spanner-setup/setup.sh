#!/bin/bash
# Script to setup fleetspeak database schema and resources on Spanner
# This is a temporary and intermediate step before integrating this into the fleetspeak code base
export ACTIONS=19

echo "1/$ACTIONS : Creating fleetspeak database on Spanner..."
gcloud spanner databases create fleetspeak --instance fleetspeak-instance

echo "2/$ACTIONS : Adding common proto..."
gcloud spanner databases ddl update fleetspeak --instance=fleetspeak-instance --ddl='CREATE PROTO BUNDLE (`fleetspeak.Address`,`fleetspeak.Annotations`,`fleetspeak.Annotations.Entry`,`fleetspeak.Label`,`fleetspeak.MessageResult`,`fleetspeak.ValidationInfo`);' --proto-descriptors-file=./common.pb

echo "3/$ACTIONS : Adding broadcast proto..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='ALTER PROTO BUNDLE INSERT (`fleetspeak.server.Broadcast`);' --proto-descriptors-file=./broadcast.pb

echo "4/$ACTIONS : Adding resource proto..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='ALTER PROTO BUNDLE INSERT (`fleetspeak.server.ClientResourceUsageRecord`);' --proto-descriptors-file=./resource.pb

echo "5/$ACTIONS : Adding any proto..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='ALTER PROTO BUNDLE INSERT (`google.protobuf.Any`);' --proto-descriptors-file=./any.pb

echo "6/$ACTIONS : Adding timestamp proto..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='ALTER PROTO BUNDLE INSERT (`google.protobuf.Timestamp`);' --proto-descriptors-file=./timestamp.pb

echo "7/$ACTIONS : Creating Clients table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE Clients (
  ClientID BYTES(8) NOT NULL,
  ClientKey BYTES(8192) NOT NULL,
  LastContactTime TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
  LastContactAddress STRING(64),
  LastContactStreamingTo STRING(128),
  Blacklisted BOOL NOT NULL,
  LastClock `google.protobuf.Timestamp`,
) PRIMARY KEY (ClientID);'

echo "8/$ACTIONS : Creating ClientLabels table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE ClientLabels (
  ClientID BYTES(8) NOT NULL,
  Name STRING(MAX) NOT NULL AS (Label.label) STORED, 
  ServiceName STRING(MAX) NOT NULL AS (Label.service_name) STORED,
  Label `fleetspeak.Label`
) PRIMARY KEY(ClientID, ServiceName, Name),
  INTERLEAVE IN PARENT Clients ON DELETE CASCADE;'

echo "9/$ACTIONS : Creating ClientContacts table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE ClientContacts (
  ClientID BYTES(8) NOT NULL,
  Time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
  SentNonce BYTES(8) NOT NULL,
  ReceivedNonce BYTES(8),
  Address STRING(64)
) PRIMARY KEY(ClientID, Time),
  INTERLEAVE IN PARENT Clients ON DELETE CASCADE,
  ROW DELETION POLICY (OLDER_THAN(Time, INTERVAL 14 DAY));'

echo "10/$ACTIONS : Creating ClientContactMessages table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE ClientContactMessages (
  ClientID BYTES(8) NOT NULL,
  Time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
  MessageID BYTES(32) NOT NULL
) PRIMARY KEY(ClientID, Time, MessageID),
  INTERLEAVE IN PARENT ClientContacts ON DELETE CASCADE;'

echo "11/$ACTIONS : Creating Messages table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE Messages (
  MessageID BYTES(32) NOT NULL,
  Source `fleetspeak.Address` NOT NULL,
  SourceMessageID BYTES(16),
  Destination `fleetspeak.Address` NOT NULL,
  MessageType STRING(64) NOT NULL,
  CreationTime TIMESTAMP,
  EncryptedData `google.protobuf.Any`,
  Result `fleetspeak.MessageResult`,
  ValidationInformation `fleetspeak.ValidationInfo`,
  Annotations `fleetspeak.Annotations`
) PRIMARY KEY(MessageID),
ROW DELETION POLICY (OLDER_THAN(CreationTime, INTERVAL 15 DAY));'

echo "12/$ACTIONS : ClientPendingMessages table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE ClientPendingMessages (
  ClientID BYTES(8) NOT NULL,
  MessageID BYTES(32) NOT NULL,
  RetryCount INT64 NOT NULL,
  ScheduledTime TIMESTAMP NOT NULL,
  DestinationServiceName STRING(128),
  CreationTime TIMESTAMP
) PRIMARY KEY(ClientID, MessageID),
  INTERLEAVE IN PARENT Clients ON DELETE CASCADE,
  ROW DELETION POLICY (OLDER_THAN(CreationTime, INTERVAL 14 DAY));'

echo "13/$ACTIONS : Creating ClientResourceUsageRecords table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE ClientResourceUsageRecords (
  ClientID BYTES(8) NOT NULL,
  ServerTimestamp TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
  Record `fleetspeak.server.ClientResourceUsageRecord`
) PRIMARY KEY(ClientID, ServerTimestamp),
  INTERLEAVE IN PARENT Clients ON DELETE CASCADE,
  ROW DELETION POLICY (OLDER_THAN(ServerTimestamp, INTERVAL 14 DAY));'

echo "14/$ACTIONS : Creating Broadcasts table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE Broadcasts(
  BroadcastID BYTES(MAX) NOT NULL AS (Broadcast.broadcast_id) STORED,
  BroadcastService STRING(MAX) NOT NULL AS (Broadcast.source.service_name) STORED,
  BroadcastExpirySeconds INT64 NOT NULL AS (Broadcast.expiration_time.seconds) STORED,
  BroadcastExpiryNanos INT64 NOT NULL AS (Broadcast.expiration_time.nanos) STORED,
  Broadcast `fleetspeak.server.Broadcast` NOT NULL,
  Sent INT64 NOT NULL,
  Allocated INT64 NOT NULL,
  MessageLimit INT64 NOT NULL
) PRIMARY KEY(BroadcastID);'

echo "15/$ACTIONS : Creating BroadcastsByService index..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE INDEX BroadcastsByService ON Broadcasts(BroadcastService);'

echo "16/$ACTIONS : Creating BroadcastsByExpiry index..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE INDEX BroadcastsByExpiry ON Broadcasts(BroadcastExpirySeconds, BroadcastExpiryNanos);'

echo "17/$ACTIONS : Creating BroadcastAllocations table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE BroadcastAllocations (
  BroadcastID BYTES(MAX) NOT NULL,
  AllocationID BYTES(8) NOT NULL,
  Sent INT64 NOT NULL,
  MessageLimit INT64 NOT NULL,
  ExpiresTime `google.protobuf.Timestamp` NOT NULL
) PRIMARY KEY (BroadcastID, AllocationID),
  INTERLEAVE IN PARENT Broadcasts ON DELETE CASCADE;'

echo "18/$ACTIONS : Creating BroadcastSent table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE BroadcastSent (
  BroadcastID BYTES(MAX) NOT NULL,
  ClientID BYTES(8) NOT NULL
) PRIMARY KEY (ClientID, BroadcastID),
  INTERLEAVE IN PARENT Clients;'

echo "19/$ACTIONS : Creating Files table..."
gcloud spanner databases ddl update fleetspeak --instance fleetspeak-instance --ddl='CREATE TABLE Files (
  Service STRING(128) NOT NULL,
  Name STRING(256) NOT NULL,
  ModifiedTime `google.protobuf.Timestamp`,
  Data BYTES(MAX)
  ) PRIMARY KEY (Service, Name);'