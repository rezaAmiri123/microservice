package tag

import (
	"fmt"
)

// Error returns tag for Error
func Error(err error) Tag {
	return newErrorTag("error", err)
}

// QueryID returns tag for QueryID
func QueryID(queryID string) Tag {
	return newStringTag("query-id", queryID)
}

// OperationName returns tag for OperationName
func OperationName(operationName string) Tag {
	return newStringTag("operation-name", operationName)
}

// storeOperation returns tag for storeOperation
func storeOperation(storeOperation string) Tag {
	return newPredefinedStringTag("store-operation", storeOperation)
}

// Addresses returns tag for Addresses
func Addresses(ads []string) Tag {
	return newObjectTag("addresses", ads)
}

// Env return tag for runtime environment
func Env(env string) Tag {
	return newStringTag("env", env)
}

// ValueType returns tag for ValueType
func ValueType(v interface{}) Tag {
	return newStringTag("value-type", fmt.Sprintf("%T", v))
}

// DefaultValue returns tag for DefaultValue
func DefaultValue(v interface{}) Tag {
	return newObjectTag("default-value", v)
}

// Port returns tag for Port
func Port(p int) Tag {
	return newInt("port", p)
}

// NextNumber returns tag for NextNumber
func NextNumber(n int64) Tag {
	return newInt64("next-number", n)
}

// Bool returns tag for Bool
func Bool(b bool) Tag {
	return newBoolTag("bool", b)
}

// LoggingCallAtKey is reserved tag
const LoggingCallAtKey = "logging-call-at"

// Kafka related

// KafkaTopicName returns tag for TopicName
func KafkaTopicName(topicName string) Tag {
	return newStringTag("kafka-topic-name", topicName)
}

// KafkaConsumerName returns tag for ConsumerName
func KafkaConsumerName(consumerName string) Tag {
	return newStringTag("kafka-consumer-name", consumerName)
}

// KafkaPartition returns tag for Partition
func KafkaPartition(partition int32) Tag {
	return newInt32("kafka-partition", partition)
}

// KafkaPartitionKey returns tag for PartitionKey
func KafkaPartitionKey(partitionKey interface{}) Tag {
	return newObjectTag("kafka-partition-key", partitionKey)
}

// KafkaOffset returns tag for Offset
func KafkaOffset(offset int64) Tag {
	return newInt64("kafka-offset", offset)
}

// TokenLastEventID returns tag for TokenLastEventID
func TokenLastEventID(id int64) Tag {
	return newInt64("token-last-event-id", id)
}

// WorkflowAction returns tag for WorkflowAction
func workflowAction(action string) Tag {
	return newPredefinedStringTag("wf-action", action)
}
