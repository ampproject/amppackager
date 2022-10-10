// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *PostgresqlConfig9_6) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *PostgresqlConfig9_6) SetSharedBuffers(v *wrapperspb.Int64Value) {
	m.SharedBuffers = v
}

func (m *PostgresqlConfig9_6) SetTempBuffers(v *wrapperspb.Int64Value) {
	m.TempBuffers = v
}

func (m *PostgresqlConfig9_6) SetMaxPreparedTransactions(v *wrapperspb.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *PostgresqlConfig9_6) SetWorkMem(v *wrapperspb.Int64Value) {
	m.WorkMem = v
}

func (m *PostgresqlConfig9_6) SetMaintenanceWorkMem(v *wrapperspb.Int64Value) {
	m.MaintenanceWorkMem = v
}

func (m *PostgresqlConfig9_6) SetReplacementSortTuples(v *wrapperspb.Int64Value) {
	m.ReplacementSortTuples = v
}

func (m *PostgresqlConfig9_6) SetAutovacuumWorkMem(v *wrapperspb.Int64Value) {
	m.AutovacuumWorkMem = v
}

func (m *PostgresqlConfig9_6) SetTempFileLimit(v *wrapperspb.Int64Value) {
	m.TempFileLimit = v
}

func (m *PostgresqlConfig9_6) SetVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.VacuumCostDelay = v
}

func (m *PostgresqlConfig9_6) SetVacuumCostPageHit(v *wrapperspb.Int64Value) {
	m.VacuumCostPageHit = v
}

func (m *PostgresqlConfig9_6) SetVacuumCostPageMiss(v *wrapperspb.Int64Value) {
	m.VacuumCostPageMiss = v
}

func (m *PostgresqlConfig9_6) SetVacuumCostPageDirty(v *wrapperspb.Int64Value) {
	m.VacuumCostPageDirty = v
}

func (m *PostgresqlConfig9_6) SetVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.VacuumCostLimit = v
}

func (m *PostgresqlConfig9_6) SetBgwriterDelay(v *wrapperspb.Int64Value) {
	m.BgwriterDelay = v
}

func (m *PostgresqlConfig9_6) SetBgwriterLruMaxpages(v *wrapperspb.Int64Value) {
	m.BgwriterLruMaxpages = v
}

func (m *PostgresqlConfig9_6) SetBgwriterLruMultiplier(v *wrapperspb.DoubleValue) {
	m.BgwriterLruMultiplier = v
}

func (m *PostgresqlConfig9_6) SetBgwriterFlushAfter(v *wrapperspb.Int64Value) {
	m.BgwriterFlushAfter = v
}

func (m *PostgresqlConfig9_6) SetBackendFlushAfter(v *wrapperspb.Int64Value) {
	m.BackendFlushAfter = v
}

func (m *PostgresqlConfig9_6) SetOldSnapshotThreshold(v *wrapperspb.Int64Value) {
	m.OldSnapshotThreshold = v
}

func (m *PostgresqlConfig9_6) SetWalLevel(v PostgresqlConfig9_6_WalLevel) {
	m.WalLevel = v
}

func (m *PostgresqlConfig9_6) SetSynchronousCommit(v PostgresqlConfig9_6_SynchronousCommit) {
	m.SynchronousCommit = v
}

func (m *PostgresqlConfig9_6) SetCheckpointTimeout(v *wrapperspb.Int64Value) {
	m.CheckpointTimeout = v
}

func (m *PostgresqlConfig9_6) SetCheckpointCompletionTarget(v *wrapperspb.DoubleValue) {
	m.CheckpointCompletionTarget = v
}

func (m *PostgresqlConfig9_6) SetCheckpointFlushAfter(v *wrapperspb.Int64Value) {
	m.CheckpointFlushAfter = v
}

func (m *PostgresqlConfig9_6) SetMaxWalSize(v *wrapperspb.Int64Value) {
	m.MaxWalSize = v
}

func (m *PostgresqlConfig9_6) SetMinWalSize(v *wrapperspb.Int64Value) {
	m.MinWalSize = v
}

func (m *PostgresqlConfig9_6) SetMaxStandbyStreamingDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyStreamingDelay = v
}

func (m *PostgresqlConfig9_6) SetDefaultStatisticsTarget(v *wrapperspb.Int64Value) {
	m.DefaultStatisticsTarget = v
}

func (m *PostgresqlConfig9_6) SetConstraintExclusion(v PostgresqlConfig9_6_ConstraintExclusion) {
	m.ConstraintExclusion = v
}

func (m *PostgresqlConfig9_6) SetCursorTupleFraction(v *wrapperspb.DoubleValue) {
	m.CursorTupleFraction = v
}

func (m *PostgresqlConfig9_6) SetFromCollapseLimit(v *wrapperspb.Int64Value) {
	m.FromCollapseLimit = v
}

func (m *PostgresqlConfig9_6) SetJoinCollapseLimit(v *wrapperspb.Int64Value) {
	m.JoinCollapseLimit = v
}

func (m *PostgresqlConfig9_6) SetForceParallelMode(v PostgresqlConfig9_6_ForceParallelMode) {
	m.ForceParallelMode = v
}

func (m *PostgresqlConfig9_6) SetClientMinMessages(v PostgresqlConfig9_6_LogLevel) {
	m.ClientMinMessages = v
}

func (m *PostgresqlConfig9_6) SetLogMinMessages(v PostgresqlConfig9_6_LogLevel) {
	m.LogMinMessages = v
}

func (m *PostgresqlConfig9_6) SetLogMinErrorStatement(v PostgresqlConfig9_6_LogLevel) {
	m.LogMinErrorStatement = v
}

func (m *PostgresqlConfig9_6) SetLogMinDurationStatement(v *wrapperspb.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *PostgresqlConfig9_6) SetLogCheckpoints(v *wrapperspb.BoolValue) {
	m.LogCheckpoints = v
}

func (m *PostgresqlConfig9_6) SetLogConnections(v *wrapperspb.BoolValue) {
	m.LogConnections = v
}

func (m *PostgresqlConfig9_6) SetLogDisconnections(v *wrapperspb.BoolValue) {
	m.LogDisconnections = v
}

func (m *PostgresqlConfig9_6) SetLogDuration(v *wrapperspb.BoolValue) {
	m.LogDuration = v
}

func (m *PostgresqlConfig9_6) SetLogErrorVerbosity(v PostgresqlConfig9_6_LogErrorVerbosity) {
	m.LogErrorVerbosity = v
}

func (m *PostgresqlConfig9_6) SetLogLockWaits(v *wrapperspb.BoolValue) {
	m.LogLockWaits = v
}

func (m *PostgresqlConfig9_6) SetLogStatement(v PostgresqlConfig9_6_LogStatement) {
	m.LogStatement = v
}

func (m *PostgresqlConfig9_6) SetLogTempFiles(v *wrapperspb.Int64Value) {
	m.LogTempFiles = v
}

func (m *PostgresqlConfig9_6) SetSearchPath(v string) {
	m.SearchPath = v
}

func (m *PostgresqlConfig9_6) SetRowSecurity(v *wrapperspb.BoolValue) {
	m.RowSecurity = v
}

func (m *PostgresqlConfig9_6) SetDefaultTransactionIsolation(v PostgresqlConfig9_6_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *PostgresqlConfig9_6) SetStatementTimeout(v *wrapperspb.Int64Value) {
	m.StatementTimeout = v
}

func (m *PostgresqlConfig9_6) SetLockTimeout(v *wrapperspb.Int64Value) {
	m.LockTimeout = v
}

func (m *PostgresqlConfig9_6) SetIdleInTransactionSessionTimeout(v *wrapperspb.Int64Value) {
	m.IdleInTransactionSessionTimeout = v
}

func (m *PostgresqlConfig9_6) SetByteaOutput(v PostgresqlConfig9_6_ByteaOutput) {
	m.ByteaOutput = v
}

func (m *PostgresqlConfig9_6) SetXmlbinary(v PostgresqlConfig9_6_XmlBinary) {
	m.Xmlbinary = v
}

func (m *PostgresqlConfig9_6) SetXmloption(v PostgresqlConfig9_6_XmlOption) {
	m.Xmloption = v
}

func (m *PostgresqlConfig9_6) SetGinPendingListLimit(v *wrapperspb.Int64Value) {
	m.GinPendingListLimit = v
}

func (m *PostgresqlConfig9_6) SetDeadlockTimeout(v *wrapperspb.Int64Value) {
	m.DeadlockTimeout = v
}

func (m *PostgresqlConfig9_6) SetMaxLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxLocksPerTransaction = v
}

func (m *PostgresqlConfig9_6) SetMaxPredLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxPredLocksPerTransaction = v
}

func (m *PostgresqlConfig9_6) SetArrayNulls(v *wrapperspb.BoolValue) {
	m.ArrayNulls = v
}

func (m *PostgresqlConfig9_6) SetBackslashQuote(v PostgresqlConfig9_6_BackslashQuote) {
	m.BackslashQuote = v
}

func (m *PostgresqlConfig9_6) SetDefaultWithOids(v *wrapperspb.BoolValue) {
	m.DefaultWithOids = v
}

func (m *PostgresqlConfig9_6) SetEscapeStringWarning(v *wrapperspb.BoolValue) {
	m.EscapeStringWarning = v
}

func (m *PostgresqlConfig9_6) SetLoCompatPrivileges(v *wrapperspb.BoolValue) {
	m.LoCompatPrivileges = v
}

func (m *PostgresqlConfig9_6) SetOperatorPrecedenceWarning(v *wrapperspb.BoolValue) {
	m.OperatorPrecedenceWarning = v
}

func (m *PostgresqlConfig9_6) SetQuoteAllIdentifiers(v *wrapperspb.BoolValue) {
	m.QuoteAllIdentifiers = v
}

func (m *PostgresqlConfig9_6) SetStandardConformingStrings(v *wrapperspb.BoolValue) {
	m.StandardConformingStrings = v
}

func (m *PostgresqlConfig9_6) SetSynchronizeSeqscans(v *wrapperspb.BoolValue) {
	m.SynchronizeSeqscans = v
}

func (m *PostgresqlConfig9_6) SetTransformNullEquals(v *wrapperspb.BoolValue) {
	m.TransformNullEquals = v
}

func (m *PostgresqlConfig9_6) SetExitOnError(v *wrapperspb.BoolValue) {
	m.ExitOnError = v
}

func (m *PostgresqlConfig9_6) SetSeqPageCost(v *wrapperspb.DoubleValue) {
	m.SeqPageCost = v
}

func (m *PostgresqlConfig9_6) SetRandomPageCost(v *wrapperspb.DoubleValue) {
	m.RandomPageCost = v
}

func (m *PostgresqlConfig9_6) SetSqlInheritance(v *wrapperspb.BoolValue) {
	m.SqlInheritance = v
}

func (m *PostgresqlConfig9_6) SetAutovacuumMaxWorkers(v *wrapperspb.Int64Value) {
	m.AutovacuumMaxWorkers = v
}

func (m *PostgresqlConfig9_6) SetAutovacuumVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostDelay = v
}

func (m *PostgresqlConfig9_6) SetAutovacuumVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostLimit = v
}

func (m *PostgresqlConfig9_6) SetAutovacuumNaptime(v *wrapperspb.Int64Value) {
	m.AutovacuumNaptime = v
}

func (m *PostgresqlConfig9_6) SetArchiveTimeout(v *wrapperspb.Int64Value) {
	m.ArchiveTimeout = v
}

func (m *PostgresqlConfig9_6) SetTrackActivityQuerySize(v *wrapperspb.Int64Value) {
	m.TrackActivityQuerySize = v
}

func (m *PostgresqlConfig9_6) SetEffectiveIoConcurrency(v *wrapperspb.Int64Value) {
	m.EffectiveIoConcurrency = v
}

func (m *PostgresqlConfig9_6) SetEffectiveCacheSize(v *wrapperspb.Int64Value) {
	m.EffectiveCacheSize = v
}

func (m *PostgresqlConfigSet9_6) SetEffectiveConfig(v *PostgresqlConfig9_6) {
	m.EffectiveConfig = v
}

func (m *PostgresqlConfigSet9_6) SetUserConfig(v *PostgresqlConfig9_6) {
	m.UserConfig = v
}

func (m *PostgresqlConfigSet9_6) SetDefaultConfig(v *PostgresqlConfig9_6) {
	m.DefaultConfig = v
}
