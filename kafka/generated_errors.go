package kafka
// Copyright 2016-2022 Confluent Inc.
// AUTOMATICALLY GENERATED ON 2022-09-05 11:58:41.026802 +0530 IST m=+0.000584006 USING librdkafka 1.9.2

/*
#include "select_rdkafka.h"
*/
import "C"

// ErrorCode is the integer representation of local and broker error codes
type ErrorCode int

// String returns a human readable representation of an error code
func (c ErrorCode) String() string {
      return C.GoString(C.rd_kafka_err2str(C.rd_kafka_resp_err_t(c)))
}

const (
    // ErrBadMsg Local: Bad message format
    ErrBadMsg ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__BAD_MSG)
    // ErrBadCompression Local: Invalid compressed data
    ErrBadCompression ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__BAD_COMPRESSION)
    // ErrDestroy Local: Broker handle destroyed
    ErrDestroy ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__DESTROY)
    // ErrFail Local: Communication failure with broker
    ErrFail ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__FAIL)
    // ErrTransport Local: Broker transport failure
    ErrTransport ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__TRANSPORT)
    // ErrCritSysResource Local: Critical system resource failure
    ErrCritSysResource ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__CRIT_SYS_RESOURCE)
    // ErrResolve Local: Host resolution failure
    ErrResolve ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__RESOLVE)
    // ErrMsgTimedOut Local: Message timed out
    ErrMsgTimedOut ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__MSG_TIMED_OUT)
    // ErrPartitionEOF Broker: No more messages
    ErrPartitionEOF ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__PARTITION_EOF)
    // ErrUnknownPartition Local: Unknown partition
    ErrUnknownPartition ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNKNOWN_PARTITION)
    // ErrFs Local: File or filesystem error
    ErrFs ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__FS)
    // ErrUnknownTopic Local: Unknown topic
    ErrUnknownTopic ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNKNOWN_TOPIC)
    // ErrAllBrokersDown Local: All broker connections are down
    ErrAllBrokersDown ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__ALL_BROKERS_DOWN)
    // ErrInvalidArg Local: Invalid argument or configuration
    ErrInvalidArg ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__INVALID_ARG)
    // ErrTimedOut Local: Timed out
    ErrTimedOut ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__TIMED_OUT)
    // ErrQueueFull Local: Queue full
    ErrQueueFull ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__QUEUE_FULL)
    // ErrIsrInsuff Local: ISR count insufficient
    ErrIsrInsuff ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__ISR_INSUFF)
    // ErrNodeUpdate Local: Broker node update
    ErrNodeUpdate ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NODE_UPDATE)
    // ErrSsl Local: SSL error
    ErrSsl ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__SSL)
    // ErrWaitCoord Local: Waiting for coordinator
    ErrWaitCoord ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__WAIT_COORD)
    // ErrUnknownGroup Local: Unknown group
    ErrUnknownGroup ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNKNOWN_GROUP)
    // ErrInProgress Local: Operation in progress
    ErrInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__IN_PROGRESS)
    // ErrPrevInProgress Local: Previous operation in progress
    ErrPrevInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__PREV_IN_PROGRESS)
    // ErrExistingSubscription Local: Existing subscription
    ErrExistingSubscription ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__EXISTING_SUBSCRIPTION)
    // ErrAssignPartitions Local: Assign partitions
    ErrAssignPartitions ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__ASSIGN_PARTITIONS)
    // ErrRevokePartitions Local: Revoke partitions
    ErrRevokePartitions ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__REVOKE_PARTITIONS)
    // ErrConflict Local: Conflicting use
    ErrConflict ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__CONFLICT)
    // ErrState Local: Erroneous state
    ErrState ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__STATE)
    // ErrUnknownProtocol Local: Unknown protocol
    ErrUnknownProtocol ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNKNOWN_PROTOCOL)
    // ErrNotImplemented Local: Not implemented
    ErrNotImplemented ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NOT_IMPLEMENTED)
    // ErrAuthentication Local: Authentication failure
    ErrAuthentication ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__AUTHENTICATION)
    // ErrNoOffset Local: No offset stored
    ErrNoOffset ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NO_OFFSET)
    // ErrOutdated Local: Outdated
    ErrOutdated ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__OUTDATED)
    // ErrTimedOutQueue Local: Timed out in queue
    ErrTimedOutQueue ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__TIMED_OUT_QUEUE)
    // ErrUnsupportedFeature Local: Required feature not supported by broker
    ErrUnsupportedFeature ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNSUPPORTED_FEATURE)
    // ErrWaitCache Local: Awaiting cache update
    ErrWaitCache ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__WAIT_CACHE)
    // ErrIntr Local: Operation interrupted
    ErrIntr ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__INTR)
    // ErrKeySerialization Local: Key serialization error
    ErrKeySerialization ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__KEY_SERIALIZATION)
    // ErrValueSerialization Local: Value serialization error
    ErrValueSerialization ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__VALUE_SERIALIZATION)
    // ErrKeyDeserialization Local: Key deserialization error
    ErrKeyDeserialization ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__KEY_DESERIALIZATION)
    // ErrValueDeserialization Local: Value deserialization error
    ErrValueDeserialization ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__VALUE_DESERIALIZATION)
    // ErrPartial Local: Partial response
    ErrPartial ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__PARTIAL)
    // ErrReadOnly Local: Read-only object
    ErrReadOnly ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__READ_ONLY)
    // ErrNoent Local: No such entry
    ErrNoent ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NOENT)
    // ErrUnderflow Local: Read underflow
    ErrUnderflow ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNDERFLOW)
    // ErrInvalidType Local: Invalid type
    ErrInvalidType ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__INVALID_TYPE)
    // ErrRetry Local: Retry operation
    ErrRetry ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__RETRY)
    // ErrPurgeQueue Local: Purged in queue
    ErrPurgeQueue ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__PURGE_QUEUE)
    // ErrPurgeInflight Local: Purged in flight
    ErrPurgeInflight ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__PURGE_INFLIGHT)
    // ErrFatal Local: Fatal error
    ErrFatal ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__FATAL)
    // ErrInconsistent Local: Inconsistent state
    ErrInconsistent ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__INCONSISTENT)
    // ErrGaplessGuarantee Local: Gap-less ordering would not be guaranteed if proceeding
    ErrGaplessGuarantee ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__GAPLESS_GUARANTEE)
    // ErrMaxPollExceeded Local: Maximum application poll interval (max.poll.interval.ms) exceeded
    ErrMaxPollExceeded ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__MAX_POLL_EXCEEDED)
    // ErrUnknownBroker Local: Unknown broker
    ErrUnknownBroker ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNKNOWN_BROKER)
    // ErrNotConfigured Local: Functionality not configured
    ErrNotConfigured ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NOT_CONFIGURED)
    // ErrFenced Local: This instance has been fenced by a newer instance
    ErrFenced ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__FENCED)
    // ErrApplication Local: Application generated error
    ErrApplication ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__APPLICATION)
    // ErrAssignmentLost Local: Group partition assignment lost
    ErrAssignmentLost ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__ASSIGNMENT_LOST)
    // ErrNoop Local: No operation performed
    ErrNoop ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NOOP)
    // ErrAutoOffsetReset Local: No offset to automatically reset to
    ErrAutoOffsetReset ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__AUTO_OFFSET_RESET)
    // ErrUnknown Unknown broker error
    ErrUnknown ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNKNOWN)
    // ErrNoError Success
    ErrNoError ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NO_ERROR)
    // ErrOffsetOutOfRange Broker: Offset out of range
    ErrOffsetOutOfRange ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_OFFSET_OUT_OF_RANGE)
    // ErrInvalidMsg Broker: Invalid message
    ErrInvalidMsg ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_MSG)
    // ErrUnknownTopicOrPart Broker: Unknown topic or partition
    ErrUnknownTopicOrPart ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNKNOWN_TOPIC_OR_PART)
    // ErrInvalidMsgSize Broker: Invalid message size
    ErrInvalidMsgSize ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_MSG_SIZE)
    // ErrLeaderNotAvailable Broker: Leader not available
    ErrLeaderNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_LEADER_NOT_AVAILABLE)
    // ErrNotLeaderForPartition Broker: Not leader for partition
    ErrNotLeaderForPartition ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_LEADER_FOR_PARTITION)
    // ErrRequestTimedOut Broker: Request timed out
    ErrRequestTimedOut ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_REQUEST_TIMED_OUT)
    // ErrBrokerNotAvailable Broker: Broker not available
    ErrBrokerNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_BROKER_NOT_AVAILABLE)
    // ErrReplicaNotAvailable Broker: Replica not available
    ErrReplicaNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_REPLICA_NOT_AVAILABLE)
    // ErrMsgSizeTooLarge Broker: Message size too large
    ErrMsgSizeTooLarge ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_MSG_SIZE_TOO_LARGE)
    // ErrStaleCtrlEpoch Broker: StaleControllerEpochCode
    ErrStaleCtrlEpoch ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_STALE_CTRL_EPOCH)
    // ErrOffsetMetadataTooLarge Broker: Offset metadata string too large
    ErrOffsetMetadataTooLarge ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_OFFSET_METADATA_TOO_LARGE)
    // ErrNetworkException Broker: Broker disconnected before response received
    ErrNetworkException ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NETWORK_EXCEPTION)
    // ErrCoordinatorLoadInProgress Broker: Coordinator load in progress
    ErrCoordinatorLoadInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_COORDINATOR_LOAD_IN_PROGRESS)
    // ErrCoordinatorNotAvailable Broker: Coordinator not available
    ErrCoordinatorNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_COORDINATOR_NOT_AVAILABLE)
    // ErrNotCoordinator Broker: Not coordinator
    ErrNotCoordinator ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_COORDINATOR)
    // ErrTopicException Broker: Invalid topic
    ErrTopicException ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TOPIC_EXCEPTION)
    // ErrRecordListTooLarge Broker: Message batch larger than configured server segment size
    ErrRecordListTooLarge ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_RECORD_LIST_TOO_LARGE)
    // ErrNotEnoughReplicas Broker: Not enough in-sync replicas
    ErrNotEnoughReplicas ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_ENOUGH_REPLICAS)
    // ErrNotEnoughReplicasAfterAppend Broker: Message(s) written to insufficient number of in-sync replicas
    ErrNotEnoughReplicasAfterAppend ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_ENOUGH_REPLICAS_AFTER_APPEND)
    // ErrInvalidRequiredAcks Broker: Invalid required acks value
    ErrInvalidRequiredAcks ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_REQUIRED_ACKS)
    // ErrIllegalGeneration Broker: Specified group generation id is not valid
    ErrIllegalGeneration ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_ILLEGAL_GENERATION)
    // ErrInconsistentGroupProtocol Broker: Inconsistent group protocol
    ErrInconsistentGroupProtocol ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INCONSISTENT_GROUP_PROTOCOL)
    // ErrInvalidGroupID Broker: Invalid group.id
    ErrInvalidGroupID ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_GROUP_ID)
    // ErrUnknownMemberID Broker: Unknown member
    ErrUnknownMemberID ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNKNOWN_MEMBER_ID)
    // ErrInvalidSessionTimeout Broker: Invalid session timeout
    ErrInvalidSessionTimeout ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_SESSION_TIMEOUT)
    // ErrRebalanceInProgress Broker: Group rebalance in progress
    ErrRebalanceInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_REBALANCE_IN_PROGRESS)
    // ErrInvalidCommitOffsetSize Broker: Commit offset data size is not valid
    ErrInvalidCommitOffsetSize ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_COMMIT_OFFSET_SIZE)
    // ErrTopicAuthorizationFailed Broker: Topic authorization failed
    ErrTopicAuthorizationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TOPIC_AUTHORIZATION_FAILED)
    // ErrGroupAuthorizationFailed Broker: Group authorization failed
    ErrGroupAuthorizationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_GROUP_AUTHORIZATION_FAILED)
    // ErrClusterAuthorizationFailed Broker: Cluster authorization failed
    ErrClusterAuthorizationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_CLUSTER_AUTHORIZATION_FAILED)
    // ErrInvalidTimestamp Broker: Invalid timestamp
    ErrInvalidTimestamp ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_TIMESTAMP)
    // ErrUnsupportedSaslMechanism Broker: Unsupported SASL mechanism
    ErrUnsupportedSaslMechanism ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNSUPPORTED_SASL_MECHANISM)
    // ErrIllegalSaslState Broker: Request not valid in current SASL state
    ErrIllegalSaslState ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_ILLEGAL_SASL_STATE)
    // ErrUnsupportedVersion Broker: API version not supported
    ErrUnsupportedVersion ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNSUPPORTED_VERSION)
    // ErrTopicAlreadyExists Broker: Topic already exists
    ErrTopicAlreadyExists ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TOPIC_ALREADY_EXISTS)
    // ErrInvalidPartitions Broker: Invalid number of partitions
    ErrInvalidPartitions ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_PARTITIONS)
    // ErrInvalidReplicationFactor Broker: Invalid replication factor
    ErrInvalidReplicationFactor ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_REPLICATION_FACTOR)
    // ErrInvalidReplicaAssignment Broker: Invalid replica assignment
    ErrInvalidReplicaAssignment ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_REPLICA_ASSIGNMENT)
    // ErrInvalidConfig Broker: Configuration is invalid
    ErrInvalidConfig ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_CONFIG)
    // ErrNotController Broker: Not controller for cluster
    ErrNotController ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_CONTROLLER)
    // ErrInvalidRequest Broker: Invalid request
    ErrInvalidRequest ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_REQUEST)
    // ErrUnsupportedForMessageFormat Broker: Message format on broker does not support request
    ErrUnsupportedForMessageFormat ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNSUPPORTED_FOR_MESSAGE_FORMAT)
    // ErrPolicyViolation Broker: Policy violation
    ErrPolicyViolation ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_POLICY_VIOLATION)
    // ErrOutOfOrderSequenceNumber Broker: Broker received an out of order sequence number
    ErrOutOfOrderSequenceNumber ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_OUT_OF_ORDER_SEQUENCE_NUMBER)
    // ErrDuplicateSequenceNumber Broker: Broker received a duplicate sequence number
    ErrDuplicateSequenceNumber ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_DUPLICATE_SEQUENCE_NUMBER)
    // ErrInvalidProducerEpoch Broker: Producer attempted an operation with an old epoch
    ErrInvalidProducerEpoch ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_PRODUCER_EPOCH)
    // ErrInvalidTxnState Broker: Producer attempted a transactional operation in an invalid state
    ErrInvalidTxnState ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_TXN_STATE)
    // ErrInvalidProducerIDMapping Broker: Producer attempted to use a producer id which is not currently assigned to its transactional id
    ErrInvalidProducerIDMapping ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_PRODUCER_ID_MAPPING)
    // ErrInvalidTransactionTimeout Broker: Transaction timeout is larger than the maximum value allowed by the broker's max.transaction.timeout.ms
    ErrInvalidTransactionTimeout ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_TRANSACTION_TIMEOUT)
    // ErrConcurrentTransactions Broker: Producer attempted to update a transaction while another concurrent operation on the same transaction was ongoing
    ErrConcurrentTransactions ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_CONCURRENT_TRANSACTIONS)
    // ErrTransactionCoordinatorFenced Broker: Indicates that the transaction coordinator sending a WriteTxnMarker is no longer the current coordinator for a given producer
    ErrTransactionCoordinatorFenced ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TRANSACTION_COORDINATOR_FENCED)
    // ErrTransactionalIDAuthorizationFailed Broker: Transactional Id authorization failed
    ErrTransactionalIDAuthorizationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TRANSACTIONAL_ID_AUTHORIZATION_FAILED)
    // ErrSecurityDisabled Broker: Security features are disabled
    ErrSecurityDisabled ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_SECURITY_DISABLED)
    // ErrOperationNotAttempted Broker: Operation not attempted
    ErrOperationNotAttempted ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_OPERATION_NOT_ATTEMPTED)
    // ErrKafkaStorageError Broker: Disk error when trying to access log file on disk
    ErrKafkaStorageError ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_KAFKA_STORAGE_ERROR)
    // ErrLogDirNotFound Broker: The user-specified log directory is not found in the broker config
    ErrLogDirNotFound ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_LOG_DIR_NOT_FOUND)
    // ErrSaslAuthenticationFailed Broker: SASL Authentication failed
    ErrSaslAuthenticationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_SASL_AUTHENTICATION_FAILED)
    // ErrUnknownProducerID Broker: Unknown Producer Id
    ErrUnknownProducerID ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNKNOWN_PRODUCER_ID)
    // ErrReassignmentInProgress Broker: Partition reassignment is in progress
    ErrReassignmentInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_REASSIGNMENT_IN_PROGRESS)
    // ErrDelegationTokenAuthDisabled Broker: Delegation Token feature is not enabled
    ErrDelegationTokenAuthDisabled ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_DELEGATION_TOKEN_AUTH_DISABLED)
    // ErrDelegationTokenNotFound Broker: Delegation Token is not found on server
    ErrDelegationTokenNotFound ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_DELEGATION_TOKEN_NOT_FOUND)
    // ErrDelegationTokenOwnerMismatch Broker: Specified Principal is not valid Owner/Renewer
    ErrDelegationTokenOwnerMismatch ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_DELEGATION_TOKEN_OWNER_MISMATCH)
    // ErrDelegationTokenRequestNotAllowed Broker: Delegation Token requests are not allowed on this connection
    ErrDelegationTokenRequestNotAllowed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_DELEGATION_TOKEN_REQUEST_NOT_ALLOWED)
    // ErrDelegationTokenAuthorizationFailed Broker: Delegation Token authorization failed
    ErrDelegationTokenAuthorizationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_DELEGATION_TOKEN_AUTHORIZATION_FAILED)
    // ErrDelegationTokenExpired Broker: Delegation Token is expired
    ErrDelegationTokenExpired ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_DELEGATION_TOKEN_EXPIRED)
    // ErrInvalidPrincipalType Broker: Supplied principalType is not supported
    ErrInvalidPrincipalType ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_PRINCIPAL_TYPE)
    // ErrNonEmptyGroup Broker: The group is not empty
    ErrNonEmptyGroup ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NON_EMPTY_GROUP)
    // ErrGroupIDNotFound Broker: The group id does not exist
    ErrGroupIDNotFound ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_GROUP_ID_NOT_FOUND)
    // ErrFetchSessionIDNotFound Broker: The fetch session ID was not found
    ErrFetchSessionIDNotFound ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_FETCH_SESSION_ID_NOT_FOUND)
    // ErrInvalidFetchSessionEpoch Broker: The fetch session epoch is invalid
    ErrInvalidFetchSessionEpoch ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_FETCH_SESSION_EPOCH)
    // ErrListenerNotFound Broker: No matching listener
    ErrListenerNotFound ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_LISTENER_NOT_FOUND)
    // ErrTopicDeletionDisabled Broker: Topic deletion is disabled
    ErrTopicDeletionDisabled ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TOPIC_DELETION_DISABLED)
    // ErrFencedLeaderEpoch Broker: Leader epoch is older than broker epoch
    ErrFencedLeaderEpoch ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_FENCED_LEADER_EPOCH)
    // ErrUnknownLeaderEpoch Broker: Leader epoch is newer than broker epoch
    ErrUnknownLeaderEpoch ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNKNOWN_LEADER_EPOCH)
    // ErrUnsupportedCompressionType Broker: Unsupported compression type
    ErrUnsupportedCompressionType ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNSUPPORTED_COMPRESSION_TYPE)
    // ErrStaleBrokerEpoch Broker: Broker epoch has changed
    ErrStaleBrokerEpoch ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_STALE_BROKER_EPOCH)
    // ErrOffsetNotAvailable Broker: Leader high watermark is not caught up
    ErrOffsetNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_OFFSET_NOT_AVAILABLE)
    // ErrMemberIDRequired Broker: Group member needs a valid member ID
    ErrMemberIDRequired ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_MEMBER_ID_REQUIRED)
    // ErrPreferredLeaderNotAvailable Broker: Preferred leader was not available
    ErrPreferredLeaderNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_PREFERRED_LEADER_NOT_AVAILABLE)
    // ErrGroupMaxSizeReached Broker: Consumer group has reached maximum size
    ErrGroupMaxSizeReached ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_GROUP_MAX_SIZE_REACHED)
    // ErrFencedInstanceID Broker: Static consumer fenced by other consumer with same group.instance.id
    ErrFencedInstanceID ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_FENCED_INSTANCE_ID)
    // ErrEligibleLeadersNotAvailable Broker: Eligible partition leaders are not available
    ErrEligibleLeadersNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_ELIGIBLE_LEADERS_NOT_AVAILABLE)
    // ErrElectionNotNeeded Broker: Leader election not needed for topic partition
    ErrElectionNotNeeded ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_ELECTION_NOT_NEEDED)
    // ErrNoReassignmentInProgress Broker: No partition reassignment is in progress
    ErrNoReassignmentInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NO_REASSIGNMENT_IN_PROGRESS)
    // ErrGroupSubscribedToTopic Broker: Deleting offsets of a topic while the consumer group is subscribed to it
    ErrGroupSubscribedToTopic ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_GROUP_SUBSCRIBED_TO_TOPIC)
    // ErrInvalidRecord Broker: Broker failed to validate record
    ErrInvalidRecord ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_RECORD)
    // ErrUnstableOffsetCommit Broker: There are unstable offsets that need to be cleared
    ErrUnstableOffsetCommit ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNSTABLE_OFFSET_COMMIT)
    // ErrThrottlingQuotaExceeded Broker: Throttling quota has been exceeded
    ErrThrottlingQuotaExceeded ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_THROTTLING_QUOTA_EXCEEDED)
    // ErrProducerFenced Broker: There is a newer producer with the same transactionalId which fences the current one
    ErrProducerFenced ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_PRODUCER_FENCED)
    // ErrResourceNotFound Broker: Request illegally referred to resource that does not exist
    ErrResourceNotFound ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_RESOURCE_NOT_FOUND)
    // ErrDuplicateResource Broker: Request illegally referred to the same resource twice
    ErrDuplicateResource ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_DUPLICATE_RESOURCE)
    // ErrUnacceptableCredential Broker: Requested credential would not meet criteria for acceptability
    ErrUnacceptableCredential ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNACCEPTABLE_CREDENTIAL)
    // ErrInconsistentVoterSet Broker: Indicates that the either the sender or recipient of a voter-only request is not one of the expected voters
    ErrInconsistentVoterSet ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INCONSISTENT_VOTER_SET)
    // ErrInvalidUpdateVersion Broker: Invalid update version
    ErrInvalidUpdateVersion ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_UPDATE_VERSION)
    // ErrFeatureUpdateFailed Broker: Unable to update finalized features due to server error
    ErrFeatureUpdateFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_FEATURE_UPDATE_FAILED)
    // ErrPrincipalDeserializationFailure Broker: Request principal deserialization failed during forwarding
    ErrPrincipalDeserializationFailure ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_PRINCIPAL_DESERIALIZATION_FAILURE)
)
