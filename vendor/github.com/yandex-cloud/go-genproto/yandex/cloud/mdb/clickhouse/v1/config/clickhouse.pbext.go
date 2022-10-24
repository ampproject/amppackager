// Code generated by protoc-gen-goext. DO NOT EDIT.

package clickhouse

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *ClickhouseConfig) SetLogLevel(v ClickhouseConfig_LogLevel) {
	m.LogLevel = v
}

func (m *ClickhouseConfig) SetMergeTree(v *ClickhouseConfig_MergeTree) {
	m.MergeTree = v
}

func (m *ClickhouseConfig) SetCompression(v []*ClickhouseConfig_Compression) {
	m.Compression = v
}

func (m *ClickhouseConfig) SetDictionaries(v []*ClickhouseConfig_ExternalDictionary) {
	m.Dictionaries = v
}

func (m *ClickhouseConfig) SetGraphiteRollup(v []*ClickhouseConfig_GraphiteRollup) {
	m.GraphiteRollup = v
}

func (m *ClickhouseConfig) SetKafka(v *ClickhouseConfig_Kafka) {
	m.Kafka = v
}

func (m *ClickhouseConfig) SetKafkaTopics(v []*ClickhouseConfig_KafkaTopic) {
	m.KafkaTopics = v
}

func (m *ClickhouseConfig) SetRabbitmq(v *ClickhouseConfig_Rabbitmq) {
	m.Rabbitmq = v
}

func (m *ClickhouseConfig) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *ClickhouseConfig) SetMaxConcurrentQueries(v *wrapperspb.Int64Value) {
	m.MaxConcurrentQueries = v
}

func (m *ClickhouseConfig) SetKeepAliveTimeout(v *wrapperspb.Int64Value) {
	m.KeepAliveTimeout = v
}

func (m *ClickhouseConfig) SetUncompressedCacheSize(v *wrapperspb.Int64Value) {
	m.UncompressedCacheSize = v
}

func (m *ClickhouseConfig) SetMarkCacheSize(v *wrapperspb.Int64Value) {
	m.MarkCacheSize = v
}

func (m *ClickhouseConfig) SetMaxTableSizeToDrop(v *wrapperspb.Int64Value) {
	m.MaxTableSizeToDrop = v
}

func (m *ClickhouseConfig) SetMaxPartitionSizeToDrop(v *wrapperspb.Int64Value) {
	m.MaxPartitionSizeToDrop = v
}

func (m *ClickhouseConfig) SetBuiltinDictionariesReloadInterval(v *wrapperspb.Int64Value) {
	m.BuiltinDictionariesReloadInterval = v
}

func (m *ClickhouseConfig) SetTimezone(v string) {
	m.Timezone = v
}

func (m *ClickhouseConfig) SetGeobaseUri(v string) {
	m.GeobaseUri = v
}

func (m *ClickhouseConfig) SetQueryLogRetentionSize(v *wrapperspb.Int64Value) {
	m.QueryLogRetentionSize = v
}

func (m *ClickhouseConfig) SetQueryLogRetentionTime(v *wrapperspb.Int64Value) {
	m.QueryLogRetentionTime = v
}

func (m *ClickhouseConfig) SetQueryThreadLogEnabled(v *wrapperspb.BoolValue) {
	m.QueryThreadLogEnabled = v
}

func (m *ClickhouseConfig) SetQueryThreadLogRetentionSize(v *wrapperspb.Int64Value) {
	m.QueryThreadLogRetentionSize = v
}

func (m *ClickhouseConfig) SetQueryThreadLogRetentionTime(v *wrapperspb.Int64Value) {
	m.QueryThreadLogRetentionTime = v
}

func (m *ClickhouseConfig) SetPartLogRetentionSize(v *wrapperspb.Int64Value) {
	m.PartLogRetentionSize = v
}

func (m *ClickhouseConfig) SetPartLogRetentionTime(v *wrapperspb.Int64Value) {
	m.PartLogRetentionTime = v
}

func (m *ClickhouseConfig) SetMetricLogEnabled(v *wrapperspb.BoolValue) {
	m.MetricLogEnabled = v
}

func (m *ClickhouseConfig) SetMetricLogRetentionSize(v *wrapperspb.Int64Value) {
	m.MetricLogRetentionSize = v
}

func (m *ClickhouseConfig) SetMetricLogRetentionTime(v *wrapperspb.Int64Value) {
	m.MetricLogRetentionTime = v
}

func (m *ClickhouseConfig) SetTraceLogEnabled(v *wrapperspb.BoolValue) {
	m.TraceLogEnabled = v
}

func (m *ClickhouseConfig) SetTraceLogRetentionSize(v *wrapperspb.Int64Value) {
	m.TraceLogRetentionSize = v
}

func (m *ClickhouseConfig) SetTraceLogRetentionTime(v *wrapperspb.Int64Value) {
	m.TraceLogRetentionTime = v
}

func (m *ClickhouseConfig) SetTextLogEnabled(v *wrapperspb.BoolValue) {
	m.TextLogEnabled = v
}

func (m *ClickhouseConfig) SetTextLogRetentionSize(v *wrapperspb.Int64Value) {
	m.TextLogRetentionSize = v
}

func (m *ClickhouseConfig) SetTextLogRetentionTime(v *wrapperspb.Int64Value) {
	m.TextLogRetentionTime = v
}

func (m *ClickhouseConfig) SetTextLogLevel(v ClickhouseConfig_LogLevel) {
	m.TextLogLevel = v
}

func (m *ClickhouseConfig) SetBackgroundPoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundPoolSize = v
}

func (m *ClickhouseConfig) SetBackgroundSchedulePoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundSchedulePoolSize = v
}

func (m *ClickhouseConfig_MergeTree) SetReplicatedDeduplicationWindow(v *wrapperspb.Int64Value) {
	m.ReplicatedDeduplicationWindow = v
}

func (m *ClickhouseConfig_MergeTree) SetReplicatedDeduplicationWindowSeconds(v *wrapperspb.Int64Value) {
	m.ReplicatedDeduplicationWindowSeconds = v
}

func (m *ClickhouseConfig_MergeTree) SetPartsToDelayInsert(v *wrapperspb.Int64Value) {
	m.PartsToDelayInsert = v
}

func (m *ClickhouseConfig_MergeTree) SetPartsToThrowInsert(v *wrapperspb.Int64Value) {
	m.PartsToThrowInsert = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxReplicatedMergesInQueue(v *wrapperspb.Int64Value) {
	m.MaxReplicatedMergesInQueue = v
}

func (m *ClickhouseConfig_MergeTree) SetNumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge(v *wrapperspb.Int64Value) {
	m.NumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxBytesToMergeAtMinSpaceInPool(v *wrapperspb.Int64Value) {
	m.MaxBytesToMergeAtMinSpaceInPool = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxBytesToMergeAtMaxSpaceInPool(v *wrapperspb.Int64Value) {
	m.MaxBytesToMergeAtMaxSpaceInPool = v
}

func (m *ClickhouseConfig_Kafka) SetSecurityProtocol(v ClickhouseConfig_Kafka_SecurityProtocol) {
	m.SecurityProtocol = v
}

func (m *ClickhouseConfig_Kafka) SetSaslMechanism(v ClickhouseConfig_Kafka_SaslMechanism) {
	m.SaslMechanism = v
}

func (m *ClickhouseConfig_Kafka) SetSaslUsername(v string) {
	m.SaslUsername = v
}

func (m *ClickhouseConfig_Kafka) SetSaslPassword(v string) {
	m.SaslPassword = v
}

func (m *ClickhouseConfig_KafkaTopic) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_KafkaTopic) SetSettings(v *ClickhouseConfig_Kafka) {
	m.Settings = v
}

func (m *ClickhouseConfig_Rabbitmq) SetUsername(v string) {
	m.Username = v
}

func (m *ClickhouseConfig_Rabbitmq) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_Compression) SetMethod(v ClickhouseConfig_Compression_Method) {
	m.Method = v
}

func (m *ClickhouseConfig_Compression) SetMinPartSize(v int64) {
	m.MinPartSize = v
}

func (m *ClickhouseConfig_Compression) SetMinPartSizeRatio(v float64) {
	m.MinPartSizeRatio = v
}

type ClickhouseConfig_ExternalDictionary_Lifetime = isClickhouseConfig_ExternalDictionary_Lifetime

func (m *ClickhouseConfig_ExternalDictionary) SetLifetime(v ClickhouseConfig_ExternalDictionary_Lifetime) {
	m.Lifetime = v
}

type ClickhouseConfig_ExternalDictionary_Source = isClickhouseConfig_ExternalDictionary_Source

func (m *ClickhouseConfig_ExternalDictionary) SetSource(v ClickhouseConfig_ExternalDictionary_Source) {
	m.Source = v
}

func (m *ClickhouseConfig_ExternalDictionary) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_ExternalDictionary) SetStructure(v *ClickhouseConfig_ExternalDictionary_Structure) {
	m.Structure = v
}

func (m *ClickhouseConfig_ExternalDictionary) SetLayout(v *ClickhouseConfig_ExternalDictionary_Layout) {
	m.Layout = v
}

func (m *ClickhouseConfig_ExternalDictionary) SetFixedLifetime(v int64) {
	m.Lifetime = &ClickhouseConfig_ExternalDictionary_FixedLifetime{
		FixedLifetime: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetLifetimeRange(v *ClickhouseConfig_ExternalDictionary_Range) {
	m.Lifetime = &ClickhouseConfig_ExternalDictionary_LifetimeRange{
		LifetimeRange: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetHttpSource(v *ClickhouseConfig_ExternalDictionary_HttpSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_HttpSource_{
		HttpSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetMysqlSource(v *ClickhouseConfig_ExternalDictionary_MysqlSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_MysqlSource_{
		MysqlSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetClickhouseSource(v *ClickhouseConfig_ExternalDictionary_ClickhouseSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_ClickhouseSource_{
		ClickhouseSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetMongodbSource(v *ClickhouseConfig_ExternalDictionary_MongodbSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_MongodbSource_{
		MongodbSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetPostgresqlSource(v *ClickhouseConfig_ExternalDictionary_PostgresqlSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_PostgresqlSource_{
		PostgresqlSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary_HttpSource) SetUrl(v string) {
	m.Url = v
}

func (m *ClickhouseConfig_ExternalDictionary_HttpSource) SetFormat(v string) {
	m.Format = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetDb(v string) {
	m.Db = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetTable(v string) {
	m.Table = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetReplicas(v []*ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) {
	m.Replicas = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetWhere(v string) {
	m.Where = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetInvalidateQuery(v string) {
	m.InvalidateQuery = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetHost(v string) {
	m.Host = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetPriority(v int64) {
	m.Priority = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetDb(v string) {
	m.Db = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetTable(v string) {
	m.Table = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetHost(v string) {
	m.Host = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetWhere(v string) {
	m.Where = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetDb(v string) {
	m.Db = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetCollection(v string) {
	m.Collection = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetHost(v string) {
	m.Host = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetDb(v string) {
	m.Db = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetTable(v string) {
	m.Table = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetHosts(v []string) {
	m.Hosts = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetInvalidateQuery(v string) {
	m.InvalidateQuery = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetSslMode(v ClickhouseConfig_ExternalDictionary_PostgresqlSource_SslMode) {
	m.SslMode = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetId(v *ClickhouseConfig_ExternalDictionary_Structure_Id) {
	m.Id = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetKey(v *ClickhouseConfig_ExternalDictionary_Structure_Key) {
	m.Key = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetRangeMin(v *ClickhouseConfig_ExternalDictionary_Structure_Attribute) {
	m.RangeMin = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetRangeMax(v *ClickhouseConfig_ExternalDictionary_Structure_Attribute) {
	m.RangeMax = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetAttributes(v []*ClickhouseConfig_ExternalDictionary_Structure_Attribute) {
	m.Attributes = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetType(v string) {
	m.Type = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetNullValue(v string) {
	m.NullValue = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetExpression(v string) {
	m.Expression = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetHierarchical(v bool) {
	m.Hierarchical = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetInjective(v bool) {
	m.Injective = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Id) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Key) SetAttributes(v []*ClickhouseConfig_ExternalDictionary_Structure_Attribute) {
	m.Attributes = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetType(v ClickhouseConfig_ExternalDictionary_Layout_Type) {
	m.Type = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetSizeInCells(v int64) {
	m.SizeInCells = v
}

func (m *ClickhouseConfig_ExternalDictionary_Range) SetMin(v int64) {
	m.Min = v
}

func (m *ClickhouseConfig_ExternalDictionary_Range) SetMax(v int64) {
	m.Max = v
}

func (m *ClickhouseConfig_GraphiteRollup) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_GraphiteRollup) SetPatterns(v []*ClickhouseConfig_GraphiteRollup_Pattern) {
	m.Patterns = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern) SetRegexp(v string) {
	m.Regexp = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern) SetFunction(v string) {
	m.Function = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern) SetRetention(v []*ClickhouseConfig_GraphiteRollup_Pattern_Retention) {
	m.Retention = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern_Retention) SetAge(v int64) {
	m.Age = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern_Retention) SetPrecision(v int64) {
	m.Precision = v
}

func (m *ClickhouseConfigSet) SetEffectiveConfig(v *ClickhouseConfig) {
	m.EffectiveConfig = v
}

func (m *ClickhouseConfigSet) SetUserConfig(v *ClickhouseConfig) {
	m.UserConfig = v
}

func (m *ClickhouseConfigSet) SetDefaultConfig(v *ClickhouseConfig) {
	m.DefaultConfig = v
}