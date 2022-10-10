// Code generated by protoc-gen-goext. DO NOT EDIT.

package triggers

import (
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Trigger) SetId(v string) {
	m.Id = v
}

func (m *Trigger) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Trigger) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Trigger) SetName(v string) {
	m.Name = v
}

func (m *Trigger) SetDescription(v string) {
	m.Description = v
}

func (m *Trigger) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Trigger) SetRule(v *Trigger_Rule) {
	m.Rule = v
}

func (m *Trigger) SetStatus(v Trigger_Status) {
	m.Status = v
}

type Trigger_Rule_Rule = isTrigger_Rule_Rule

func (m *Trigger_Rule) SetRule(v Trigger_Rule_Rule) {
	m.Rule = v
}

func (m *Trigger_Rule) SetTimer(v *Trigger_Timer) {
	m.Rule = &Trigger_Rule_Timer{
		Timer: v,
	}
}

func (m *Trigger_Rule) SetMessageQueue(v *Trigger_MessageQueue) {
	m.Rule = &Trigger_Rule_MessageQueue{
		MessageQueue: v,
	}
}

func (m *Trigger_Rule) SetIotMessage(v *Trigger_IoTMessage) {
	m.Rule = &Trigger_Rule_IotMessage{
		IotMessage: v,
	}
}

func (m *Trigger_Rule) SetObjectStorage(v *Trigger_ObjectStorage) {
	m.Rule = &Trigger_Rule_ObjectStorage{
		ObjectStorage: v,
	}
}

func (m *Trigger_Rule) SetContainerRegistry(v *Trigger_ContainerRegistry) {
	m.Rule = &Trigger_Rule_ContainerRegistry{
		ContainerRegistry: v,
	}
}

func (m *Trigger_Rule) SetCloudLogs(v *Trigger_CloudLogs) {
	m.Rule = &Trigger_Rule_CloudLogs{
		CloudLogs: v,
	}
}

func (m *Trigger_Rule) SetLogging(v *Trigger_Logging) {
	m.Rule = &Trigger_Rule_Logging{
		Logging: v,
	}
}

func (m *Trigger_Rule) SetBillingBudget(v *BillingBudget) {
	m.Rule = &Trigger_Rule_BillingBudget{
		BillingBudget: v,
	}
}

func (m *Trigger_Rule) SetDataStream(v *DataStream) {
	m.Rule = &Trigger_Rule_DataStream{
		DataStream: v,
	}
}

func (m *Trigger_Rule) SetMail(v *Mail) {
	m.Rule = &Trigger_Rule_Mail{
		Mail: v,
	}
}

type Trigger_Timer_Action = isTrigger_Timer_Action

func (m *Trigger_Timer) SetAction(v Trigger_Timer_Action) {
	m.Action = v
}

func (m *Trigger_Timer) SetCronExpression(v string) {
	m.CronExpression = v
}

func (m *Trigger_Timer) SetInvokeFunction(v *InvokeFunctionOnce) {
	m.Action = &Trigger_Timer_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *Trigger_Timer) SetInvokeFunctionWithRetry(v *InvokeFunctionWithRetry) {
	m.Action = &Trigger_Timer_InvokeFunctionWithRetry{
		InvokeFunctionWithRetry: v,
	}
}

func (m *Trigger_Timer) SetInvokeContainerWithRetry(v *InvokeContainerWithRetry) {
	m.Action = &Trigger_Timer_InvokeContainerWithRetry{
		InvokeContainerWithRetry: v,
	}
}

type Trigger_MessageQueue_Action = isTrigger_MessageQueue_Action

func (m *Trigger_MessageQueue) SetAction(v Trigger_MessageQueue_Action) {
	m.Action = v
}

func (m *Trigger_MessageQueue) SetQueueId(v string) {
	m.QueueId = v
}

func (m *Trigger_MessageQueue) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *Trigger_MessageQueue) SetBatchSettings(v *BatchSettings) {
	m.BatchSettings = v
}

func (m *Trigger_MessageQueue) SetVisibilityTimeout(v *durationpb.Duration) {
	m.VisibilityTimeout = v
}

func (m *Trigger_MessageQueue) SetInvokeFunction(v *InvokeFunctionOnce) {
	m.Action = &Trigger_MessageQueue_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *Trigger_MessageQueue) SetInvokeContainer(v *InvokeContainerOnce) {
	m.Action = &Trigger_MessageQueue_InvokeContainer{
		InvokeContainer: v,
	}
}

type Trigger_IoTMessage_Action = isTrigger_IoTMessage_Action

func (m *Trigger_IoTMessage) SetAction(v Trigger_IoTMessage_Action) {
	m.Action = v
}

func (m *Trigger_IoTMessage) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *Trigger_IoTMessage) SetDeviceId(v string) {
	m.DeviceId = v
}

func (m *Trigger_IoTMessage) SetMqttTopic(v string) {
	m.MqttTopic = v
}

func (m *Trigger_IoTMessage) SetInvokeFunction(v *InvokeFunctionWithRetry) {
	m.Action = &Trigger_IoTMessage_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *Trigger_IoTMessage) SetInvokeContainer(v *InvokeContainerWithRetry) {
	m.Action = &Trigger_IoTMessage_InvokeContainer{
		InvokeContainer: v,
	}
}

type Trigger_ObjectStorage_Action = isTrigger_ObjectStorage_Action

func (m *Trigger_ObjectStorage) SetAction(v Trigger_ObjectStorage_Action) {
	m.Action = v
}

func (m *Trigger_ObjectStorage) SetEventType(v []Trigger_ObjectStorageEventType) {
	m.EventType = v
}

func (m *Trigger_ObjectStorage) SetBucketId(v string) {
	m.BucketId = v
}

func (m *Trigger_ObjectStorage) SetPrefix(v string) {
	m.Prefix = v
}

func (m *Trigger_ObjectStorage) SetSuffix(v string) {
	m.Suffix = v
}

func (m *Trigger_ObjectStorage) SetInvokeFunction(v *InvokeFunctionWithRetry) {
	m.Action = &Trigger_ObjectStorage_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *Trigger_ObjectStorage) SetInvokeContainer(v *InvokeContainerWithRetry) {
	m.Action = &Trigger_ObjectStorage_InvokeContainer{
		InvokeContainer: v,
	}
}

type Trigger_ContainerRegistry_Action = isTrigger_ContainerRegistry_Action

func (m *Trigger_ContainerRegistry) SetAction(v Trigger_ContainerRegistry_Action) {
	m.Action = v
}

func (m *Trigger_ContainerRegistry) SetEventType(v []Trigger_ContainerRegistryEventType) {
	m.EventType = v
}

func (m *Trigger_ContainerRegistry) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *Trigger_ContainerRegistry) SetImageName(v string) {
	m.ImageName = v
}

func (m *Trigger_ContainerRegistry) SetTag(v string) {
	m.Tag = v
}

func (m *Trigger_ContainerRegistry) SetInvokeFunction(v *InvokeFunctionWithRetry) {
	m.Action = &Trigger_ContainerRegistry_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *Trigger_ContainerRegistry) SetInvokeContainer(v *InvokeContainerWithRetry) {
	m.Action = &Trigger_ContainerRegistry_InvokeContainer{
		InvokeContainer: v,
	}
}

type Trigger_CloudLogs_Action = isTrigger_CloudLogs_Action

func (m *Trigger_CloudLogs) SetAction(v Trigger_CloudLogs_Action) {
	m.Action = v
}

func (m *Trigger_CloudLogs) SetLogGroupId(v []string) {
	m.LogGroupId = v
}

func (m *Trigger_CloudLogs) SetBatchSettings(v *CloudLogsBatchSettings) {
	m.BatchSettings = v
}

func (m *Trigger_CloudLogs) SetInvokeFunction(v *InvokeFunctionWithRetry) {
	m.Action = &Trigger_CloudLogs_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *Trigger_CloudLogs) SetInvokeContainer(v *InvokeContainerWithRetry) {
	m.Action = &Trigger_CloudLogs_InvokeContainer{
		InvokeContainer: v,
	}
}

type Trigger_Logging_Action = isTrigger_Logging_Action

func (m *Trigger_Logging) SetAction(v Trigger_Logging_Action) {
	m.Action = v
}

func (m *Trigger_Logging) SetLogGroupId(v string) {
	m.LogGroupId = v
}

func (m *Trigger_Logging) SetResourceType(v []string) {
	m.ResourceType = v
}

func (m *Trigger_Logging) SetResourceId(v []string) {
	m.ResourceId = v
}

func (m *Trigger_Logging) SetLevels(v []v1.LogLevel_Level) {
	m.Levels = v
}

func (m *Trigger_Logging) SetBatchSettings(v *LoggingBatchSettings) {
	m.BatchSettings = v
}

func (m *Trigger_Logging) SetInvokeFunction(v *InvokeFunctionWithRetry) {
	m.Action = &Trigger_Logging_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *Trigger_Logging) SetInvokeContainer(v *InvokeContainerWithRetry) {
	m.Action = &Trigger_Logging_InvokeContainer{
		InvokeContainer: v,
	}
}

func (m *InvokeFunctionOnce) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *InvokeFunctionOnce) SetFunctionTag(v string) {
	m.FunctionTag = v
}

func (m *InvokeFunctionOnce) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *InvokeFunctionWithRetry) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *InvokeFunctionWithRetry) SetFunctionTag(v string) {
	m.FunctionTag = v
}

func (m *InvokeFunctionWithRetry) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *InvokeFunctionWithRetry) SetRetrySettings(v *RetrySettings) {
	m.RetrySettings = v
}

func (m *InvokeFunctionWithRetry) SetDeadLetterQueue(v *PutQueueMessage) {
	m.DeadLetterQueue = v
}

func (m *InvokeContainerOnce) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *InvokeContainerOnce) SetPath(v string) {
	m.Path = v
}

func (m *InvokeContainerOnce) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *InvokeContainerWithRetry) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *InvokeContainerWithRetry) SetPath(v string) {
	m.Path = v
}

func (m *InvokeContainerWithRetry) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *InvokeContainerWithRetry) SetRetrySettings(v *RetrySettings) {
	m.RetrySettings = v
}

func (m *InvokeContainerWithRetry) SetDeadLetterQueue(v *PutQueueMessage) {
	m.DeadLetterQueue = v
}

func (m *PutQueueMessage) SetQueueId(v string) {
	m.QueueId = v
}

func (m *PutQueueMessage) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *BatchSettings) SetSize(v int64) {
	m.Size = v
}

func (m *BatchSettings) SetCutoff(v *durationpb.Duration) {
	m.Cutoff = v
}

func (m *CloudLogsBatchSettings) SetSize(v int64) {
	m.Size = v
}

func (m *CloudLogsBatchSettings) SetCutoff(v *durationpb.Duration) {
	m.Cutoff = v
}

func (m *LoggingBatchSettings) SetSize(v int64) {
	m.Size = v
}

func (m *LoggingBatchSettings) SetCutoff(v *durationpb.Duration) {
	m.Cutoff = v
}

func (m *RetrySettings) SetRetryAttempts(v int64) {
	m.RetryAttempts = v
}

func (m *RetrySettings) SetInterval(v *durationpb.Duration) {
	m.Interval = v
}

type BillingBudget_Action = isBillingBudget_Action

func (m *BillingBudget) SetAction(v BillingBudget_Action) {
	m.Action = v
}

func (m *BillingBudget) SetBillingAccountId(v string) {
	m.BillingAccountId = v
}

func (m *BillingBudget) SetBudgetId(v string) {
	m.BudgetId = v
}

func (m *BillingBudget) SetInvokeFunction(v *InvokeFunctionWithRetry) {
	m.Action = &BillingBudget_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *BillingBudget) SetInvokeContainer(v *InvokeContainerWithRetry) {
	m.Action = &BillingBudget_InvokeContainer{
		InvokeContainer: v,
	}
}

func (m *DataStreamBatchSettings) SetSize(v int64) {
	m.Size = v
}

func (m *DataStreamBatchSettings) SetCutoff(v *durationpb.Duration) {
	m.Cutoff = v
}

type DataStream_Action = isDataStream_Action

func (m *DataStream) SetAction(v DataStream_Action) {
	m.Action = v
}

func (m *DataStream) SetEndpoint(v string) {
	m.Endpoint = v
}

func (m *DataStream) SetDatabase(v string) {
	m.Database = v
}

func (m *DataStream) SetStream(v string) {
	m.Stream = v
}

func (m *DataStream) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *DataStream) SetBatchSettings(v *DataStreamBatchSettings) {
	m.BatchSettings = v
}

func (m *DataStream) SetInvokeFunction(v *InvokeFunctionWithRetry) {
	m.Action = &DataStream_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *DataStream) SetInvokeContainer(v *InvokeContainerWithRetry) {
	m.Action = &DataStream_InvokeContainer{
		InvokeContainer: v,
	}
}

type Mail_Action = isMail_Action

func (m *Mail) SetAction(v Mail_Action) {
	m.Action = v
}

func (m *Mail) SetEmail(v string) {
	m.Email = v
}

func (m *Mail) SetInvokeFunction(v *InvokeFunctionWithRetry) {
	m.Action = &Mail_InvokeFunction{
		InvokeFunction: v,
	}
}

func (m *Mail) SetInvokeContainer(v *InvokeContainerWithRetry) {
	m.Action = &Mail_InvokeContainer{
		InvokeContainer: v,
	}
}
