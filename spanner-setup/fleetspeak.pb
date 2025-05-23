
�
google/protobuf/any.protogoogle.protobuf"6
Any
type_url (	RtypeUrl
value (RvalueBv
com.google.protobufBAnyProtoPZ,google.golang.org/protobuf/types/known/anypb�GPB�Google.Protobuf.WellKnownTypesbproto3
�
google/protobuf/timestamp.protogoogle.protobuf";
	Timestamp
seconds (Rseconds
nanos (RnanosB�
com.google.protobufBTimestampProtoPZ2google.golang.org/protobuf/types/known/timestamppb��GPB�Google.Protobuf.WellKnownTypesbproto3
�
3fleetspeak/src/common/proto/fleetspeak/common.proto
fleetspeakgoogle/protobuf/any.protogoogle/protobuf/timestamp.proto"I
Address
	client_id (RclientId!
service_name (	RserviceName"�
ValidationInfo8
tags (2$.fleetspeak.ValidationInfo.TagsEntryRtags7
	TagsEntry
key (	Rkey
value (	Rvalue:8"�
Message

message_id (R	messageId+
source (2.fleetspeak.AddressRsource2
is_blocklisted_source (RisBlocklistedSource*
source_message_id (RsourceMessageId5
destination (2.fleetspeak.AddressRdestination!
message_type (	RmessageType?
creation_time (2.google.protobuf.TimestampRcreationTime(
data (2.google.protobuf.AnyRdataC
validation_info (2.fleetspeak.ValidationInfoRvalidationInfo1
result	 (2.fleetspeak.MessageResultRresult8
priority
 (2.fleetspeak.Message.PriorityRpriority

background (R
background9
annotations (2.fleetspeak.AnnotationsRannotations")
Priority

MEDIUM 
LOW
HIGH"�
MessageResultA
processed_time (2.google.protobuf.TimestampRprocessedTime
failed (Rfailed#
failed_reason (	RfailedReason"w
Annotations7
entries (2.fleetspeak.Annotations.EntryRentries/
Entry
key (	Rkey
value (	Rvalue"@
Label!
service_name (	RserviceName
label (	Rlabel"i
	Signature 
certificate (Rcertificate
	algorithm (R	algorithm
	signature (R	signature"�
WrappedContactData!
contact_data (RcontactData5

signatures (2.fleetspeak.SignatureR
signatures#
client_labels (	RclientLabels"�
ContactData)
sequencing_nonce (RsequencingNonce/
messages (2.fleetspeak.MessageRmessages=
client_clock (2.google.protobuf.TimestampRclientClock
	ack_index (RackIndex!
done_sending (RdoneSendingV
AllowedMessages (2,.fleetspeak.ContactData.AllowedMessagesEntryRAllowedMessagesB
AllowedMessagesEntry
key (	Rkey
value (Rvalue:8"
EmptyMessage*E
CompressionAlgorithm
COMPRESSION_NONE 
COMPRESSION_DEFLATEBEZCgithub.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeakbproto3
�
>fleetspeak/src/server/proto/fleetspeak_server/broadcasts.protofleetspeak.server3fleetspeak/src/common/proto/fleetspeak/common.protogoogle/protobuf/any.protogoogle/protobuf/timestamp.proto"�
	Broadcast!
broadcast_id (RbroadcastId+
source (2.fleetspeak.AddressRsource!
message_type (	RmessageType:
required_labels (2.fleetspeak.LabelRrequiredLabelsC
expiration_time (2.google.protobuf.TimestampRexpirationTime(
data (2.google.protobuf.AnyRdataBLZJgithub.com/google/fleetspeak/fleetspeak/src/server/proto/fleetspeak_serverbproto3
�
<fleetspeak/src/server/proto/fleetspeak_server/resource.protofleetspeak.servergoogle/protobuf/timestamp.proto"�
ClientResourceUsageRecord
scope (	Rscope
pid (RpidH
process_start_time (2.google.protobuf.TimestampRprocessStartTimeE
client_timestamp (2.google.protobuf.TimestampRclientTimestampE
server_timestamp (2.google.protobuf.TimestampRserverTimestamp-
process_terminated (RprocessTerminated+
mean_user_cpu_rate (RmeanUserCpuRate)
max_user_cpu_rate (RmaxUserCpuRate/
mean_system_cpu_rate (RmeanSystemCpuRate-
max_system_cpu_rate	 (RmaxSystemCpuRate7
mean_resident_memory_mib
 (RmeanResidentMemoryMib5
max_resident_memory_mib (RmaxResidentMemoryMib 
mean_num_fds (R
meanNumFds
max_num_fds (R	maxNumFdsBLZJgithub.com/google/fleetspeak/fleetspeak/src/server/proto/fleetspeak_serverbproto3