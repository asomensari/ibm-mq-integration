package main

import (
	"github.com/newrelic/infra-integrations-sdk/data/metric"
)

var defaultMetrics = map[string][]interface{}{
	"mq.mqstat":  {"mqstat", metric.GAUGE},
	"mq.mqdisc":  {"mqdisc", metric.GAUGE},
	"mq.mqopen":  {"mqopen", metric.GAUGE},
	"mq.mqclose": {"mqclose", metric.GAUGE},
	"mq.mqinq":   {"mqinq", metric.GAUGE},
	"mq.mqsubrq": {"mqsubrq", metric.GAUGE},
	"mq.mqctl":   {"mqctl", metric.GAUGE},
	"mq.mqset":   {"mqset", metric.GAUGE},
	"mq.mqcb":    {"mqcb", metric.GAUGE},

	"queueManager":   {"queueManager", metric.ATTRIBUTE},
	"monitoredQueue": {"monitoredQueue", metric.ATTRIBUTE},

	"fail.failedMqopen":                        {"failedMqopen", metric.GAUGE},
	"fail.failedMqclose":                       {"failedMqclose", metric.GAUGE},
	"fail.failedMqcb":                          {"failedMqcb", metric.GAUGE},
	"fail.failedMqinq":                         {"failedMqinq", metric.GAUGE},
	"fail.failedMqconnMqconnx":                 {"failedMqconnMqconnx", metric.GAUGE},
	"fail.failedMqset":                         {"failedMqset", metric.GAUGE},
	"fail.failedMqsubrq":                       {"failedMqsubrq", metric.GAUGE},
	"fail.failedMqput1":                        {"failedMqput1", metric.GAUGE},
	"fail.failedBrowse":                        {"failedBrowse", metric.GAUGE},
	"fail.failedCreateAlterResumeSubscription": {"failedCreateAlterResumeSubscription", metric.GAUGE},
	"fail.failedMqget":                         {"failedMqget", metric.GAUGE},
	"fail.failedMqput":                         {"failedMqput", metric.GAUGE},

	"log.logPhysicalWrittenBytes":            {"logPhysicalWrittenBytes", metric.GAUGE},
	"log.logWorkloadPrimarySpaceUtilization": {"logWorkloadPrimarySpaceUtilization", metric.GAUGE},
	"log.logMaxBytes":                        {"logMaxBytes", metric.GAUGE},
	"log.logLogicalWrittenBytes":             {"logLogicalWrittenBytes", metric.GAUGE},
	"log.logFileSystemMaxBytes":              {"logFileSystemMaxBytes", metric.GAUGE},
	"log.logFileSystemInUseBytes":            {"logFileSystemInUseBytes", metric.GAUGE},
	"log.logWriteLatencySeconds":             {"logWriteLatencySeconds", metric.GAUGE},
	"log.logInUseBytes":                      {"logInUseBytes", metric.GAUGE},
	"log.logWriteSize":                       {"logWriteSize", metric.GAUGE},
	"log.logCurrentPrimarySpaceInUse":        {"logCurrentPrimarySpaceInUse", metric.GAUGE},

	"ram.ramFreePercentage":                    {"ramFreePercentage", metric.GAUGE},
	"ram.ramTotalBytesEstimateForQueueManager": {"ramTotalBytesEstimateForQueueManager", metric.GAUGE},
	"ram.ramTotalBytes":                        {"ramTotalBytes", metric.GAUGE},

	"commit":          {"commit", metric.GAUGE},
	"rollback":        {"rollback", metric.GAUGE},
	"expiredMessages": {"expiredMessages", metric.GAUGE},
}

/*

createDurableSubscription

mqErrorsFileSystemInUseBytes


nonDurableSubscriberHighWaterMark
putNonPersistentMessagesBytes

nonPersistentMessageBrowseBytes
publishedToSubscribersBytes

persistentTopicMqputMqput1
persistentMessageDestructiveGet
persistentMessageMqput
persistentMessageMqput1
persistentMessageBrowse
persistentMessageBrowseBytes

topicMqputMqput1IntervalTotal


intervalTotalTopicBytesPut
resumeDurableSubscription
mqFdcFiles

intervalTotalDestructiveGetBytes
purgedQueue
nonPersistentMessageBrowse

gotNonPersistentMessagesBytes
queueManagerFileSystemInUseBytes
durableSubscriberLowWaterMark

mqTraceFileSystemInUseBytes
mqTraceFileSystemFreeSpacePercentage

deleteNonDurableSubscription
intervalTotalMqputMqput1
systemCpuTimeEstimateForQueueManagerPercentage
gotPersistentMessagesBytes

publishedToSubscribersMessages

putPersistentMessagesBytes
intervalTotalMqputMqput1Bytes


intervalTotalDestructiveGet
subscriptionDeleteFailure

alterDurableSubscription
createNonDurableSubscription


deleteDurableSubscription

mqErrorsFileSystemFreeSpacePercentage

queueManagerFileSystemFreeSpacePercentage


nonPersistentMessageMqput
nonPersistentMessageMqput1
nonPersistentMessageDestructiveGet
nonDurableSubscriberLowWaterMark


cpuLoadFiveMinuteAverage
cpuLoadFifteenMinuteAverage
cpuLoadOneMinuteAverage


userCpuTimePercentage
userCpuTimeEstimateForQueueManagerPercentage

durableSubscriberHighWaterMark
mqconnMqconnx
concurrentConnectionsHighWaterMark
systemCpuTimePercentage

nonPersistentTopicMqputMqput1



*/
