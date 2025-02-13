import os

from google.cloud import spanner
from google.api_core.exceptions import AlreadyExists


def update_database_ddl(
    project_id: str,
    instance_id: str,
    database_id: str,
    ddl: str,
    descriptors_path: str | None = None,
):
    from google.cloud.spanner_admin_database_v1.types import spanner_database_admin

    client = spanner.Client(project=project_id)
    database_admin_api = client.database_admin_api
    request = spanner_database_admin.UpdateDatabaseDdlRequest(
        database=database_admin_api.database_path(
            client.project, instance_id, database_id
        ),
        statements=[ddl],
    )
    if descriptors_path is not None:
        with open(descriptors_path, "rb") as f:
            request.proto_descriptors = f.read()
    database_admin_api.update_database_ddl(request).result()


def main():
    project_id = os.environ["GCP_PROJECT"]
    instance_name = os.environ["INSTANCE_NAME"]
    database_name = os.environ["DATABASE_NAME"]
    print(f"initializing spanner instance {instance_name} database {database_name}")
    with open("schema.ddl") as f:
        db_init_sql = f.read()
    client = spanner.Client(project=project_id)
    instance = client.instance(instance_name)
    try:
        instance.create().result()
    except AlreadyExists:
        pass
    database = instance.database(database_name)
    try:
        database.create().result()
    except AlreadyExists:
        print("database already exists. Dropping...")
        database.drop()
        print("recreating database")
        database.create().result()

    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE PROTO BUNDLE (`fleetspeak.Address`,`fleetspeak.Annotations`,`fleetspeak.Annotations.Entry`,`fleetspeak.Label`,`fleetspeak.MessageResult`,`fleetspeak.ValidationInfo`,`fleetspeak.server.Broadcast`,`fleetspeak.server.ClientResourceUsageRecord`,`google.protobuf.Any`,`google.protobuf.Timestamp`)""",
        "./fleetspeak.pb",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE Clients (
      ClientID BYTES(8) NOT NULL,
      ClientKey BYTES(8192) NOT NULL,
      LastContactTime TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
      LastContactAddress STRING(64),
      LastContactStreamingTo STRING(128),
      Blacklisted BOOL NOT NULL,
      LastClock `google.protobuf.Timestamp`,
    ) PRIMARY KEY (ClientID)""",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE ClientLabels (
      ClientID BYTES(8) NOT NULL,
      Name STRING(MAX) NOT NULL AS (Label.label) STORED, 
      ServiceName STRING(MAX) NOT NULL AS (Label.service_name) STORED,
      Label `fleetspeak.Label`
    ) PRIMARY KEY(ClientID, ServiceName, Name),
      INTERLEAVE IN PARENT Clients ON DELETE CASCADE""",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE ClientContacts (
      ClientID BYTES(8) NOT NULL,
      Time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
      SentNonce BYTES(8) NOT NULL,
      ReceivedNonce BYTES(8),
      Address STRING(64)
    ) PRIMARY KEY(ClientID, Time),
      INTERLEAVE IN PARENT Clients ON DELETE CASCADE,
      ROW DELETION POLICY (OLDER_THAN(Time, INTERVAL 14 DAY))""",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE ClientContactMessages (
      ClientID BYTES(8) NOT NULL,
      Time TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
      MessageID BYTES(32) NOT NULL
    ) PRIMARY KEY(ClientID, Time, MessageID),
      INTERLEAVE IN PARENT ClientContacts ON DELETE CASCADE""",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE Messages (
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
    ROW DELETION POLICY (OLDER_THAN(CreationTime, INTERVAL 15 DAY))""",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE ClientPendingMessages (
      ClientID BYTES(8) NOT NULL,
      MessageID BYTES(32) NOT NULL,
      RetryCount INT64 NOT NULL,
      ScheduledTime TIMESTAMP NOT NULL,
      DestinationServiceName STRING(128),
      CreationTime TIMESTAMP
    ) PRIMARY KEY(ClientID, MessageID),
      INTERLEAVE IN PARENT Clients ON DELETE CASCADE,
      ROW DELETION POLICY (OLDER_THAN(CreationTime, INTERVAL 14 DAY))""",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE ClientResourceUsageRecords (
      ClientID BYTES(8) NOT NULL,
      ServerTimestamp TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
      Record `fleetspeak.server.ClientResourceUsageRecord`
    ) PRIMARY KEY(ClientID, ServerTimestamp),
      INTERLEAVE IN PARENT Clients ON DELETE CASCADE,
      ROW DELETION POLICY (OLDER_THAN(ServerTimestamp, INTERVAL 14 DAY))""",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE Broadcasts(
      BroadcastID BYTES(MAX) NOT NULL AS (Broadcast.broadcast_id) STORED,
      Broadcast `fleetspeak.server.Broadcast` NOT NULL,
      Sent INT64 NOT NULL,
      Allocated INT64 NOT NULL,
      MessageLimit INT64 NOT NULL
    ) PRIMARY KEY(BroadcastID)""",
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE BroadcastAllocations (
      BroadcastID BYTES(MAX) NOT NULL,
      AllocationID BYTES(8) NOT NULL,
      Sent INT64 NOT NULL,
      MessageLimit INT64 NOT NULL,
      ExpiresTime `google.protobuf.Timestamp` NOT NULL
    ) PRIMARY KEY (BroadcastID, AllocationID),
      INTERLEAVE IN PARENT Broadcasts ON DELETE CASCADE""",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE BroadcastSent (
      BroadcastID BYTES(MAX) NOT NULL,
      ClientID BYTES(8) NOT NULL
    ) PRIMARY KEY (ClientID, BroadcastID),
      INTERLEAVE IN PARENT Clients""",
    )
    update_database_ddl(
        project_id,
        instance_name,
        database_name,
        """CREATE TABLE Files (
      Service STRING(128) NOT NULL,
      Name STRING(256) NOT NULL,
      ModifiedTime `google.protobuf.Timestamp`,
      Data BYTES(MAX)
    ) PRIMARY KEY (Service, Name)"""

if __name__ == "__main__":
    main()
