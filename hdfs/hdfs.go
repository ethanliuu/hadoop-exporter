package hdfs

type NamenDatan struct {
	Beans []struct {
		Name string `json:"name"`
		ModelerType string `json:"modelerType"`
		TagContext string `json:"tag.Context,omitempty"`
		TagProcessName string `json:"tag.ProcessName,omitempty"`
		TagSessionID interface{} `json:"tag.SessionId,omitempty"`
		TagHostname string `json:"tag.Hostname,omitempty"`
		MemNonHeapUsedM float64 `json:"MemNonHeapUsedM,omitempty"`
		MemNonHeapCommittedM float64 `json:"MemNonHeapCommittedM,omitempty"`
		MemNonHeapMaxM float64 `json:"MemNonHeapMaxM,omitempty"`
		MemHeapUsedM float64 `json:"MemHeapUsedM,omitempty"`
		MemHeapCommittedM float64 `json:"MemHeapCommittedM,omitempty"`
		MemHeapMaxM float64 `json:"MemHeapMaxM,omitempty"`
		MemMaxM float64 `json:"MemMaxM,omitempty"`
		GcCount int `json:"GcCount,omitempty"`
		GcTimeMillis int `json:"GcTimeMillis,omitempty"`
		GcNumWarnThresholdExceeded int `json:"GcNumWarnThresholdExceeded,omitempty"`
		GcNumInfoThresholdExceeded int `json:"GcNumInfoThresholdExceeded,omitempty"`
		GcTotalExtraSleepTime int `json:"GcTotalExtraSleepTime,omitempty"`
		ThreadsNew int `json:"ThreadsNew,omitempty"`
		ThreadsRunnable int `json:"ThreadsRunnable,omitempty"`
		ThreadsBlocked int `json:"ThreadsBlocked,omitempty"`
		ThreadsWaiting int `json:"ThreadsWaiting,omitempty"`
		ThreadsTimedWaiting int `json:"ThreadsTimedWaiting,omitempty"`
		ThreadsTerminated int `json:"ThreadsTerminated,omitempty"`
		LogFatal int `json:"LogFatal,omitempty"`
		LogError int `json:"LogError,omitempty"`
		LogWarn int `json:"LogWarn,omitempty"`
		LogInfo int `json:"LogInfo,omitempty"`
		CreateFileOps int `json:"CreateFileOps,omitempty"`
		FilesCreated int `json:"FilesCreated,omitempty"`
		FilesAppended int `json:"FilesAppended,omitempty"`
		GetBlockLocations int64 `json:"GetBlockLocations,omitempty"`
		FilesRenamed int `json:"FilesRenamed,omitempty"`
		GetListingOps int `json:"GetListingOps,omitempty"`
		DeleteFileOps int `json:"DeleteFileOps,omitempty"`
		FilesDeleted int `json:"FilesDeleted,omitempty"`
		FileInfoOps int `json:"FileInfoOps,omitempty"`
		AddBlockOps int `json:"AddBlockOps,omitempty"`
		GetAdditionalDatanodeOps int `json:"GetAdditionalDatanodeOps,omitempty"`
		CreateSymlinkOps int `json:"CreateSymlinkOps,omitempty"`
		GetLinkTargetOps int `json:"GetLinkTargetOps,omitempty"`
		FilesInGetListingOps int `json:"FilesInGetListingOps,omitempty"`
		AllowSnapshotOps int `json:"AllowSnapshotOps,omitempty"`
		DisallowSnapshotOps int `json:"DisallowSnapshotOps,omitempty"`
		CreateSnapshotOps int `json:"CreateSnapshotOps,omitempty"`
		DeleteSnapshotOps int `json:"DeleteSnapshotOps,omitempty"`
		RenameSnapshotOps int `json:"RenameSnapshotOps,omitempty"`
		ListSnapshottableDirOps int `json:"ListSnapshottableDirOps,omitempty"`
		SnapshotDiffReportOps int `json:"SnapshotDiffReportOps,omitempty"`
		BlockReceivedAndDeletedOps int `json:"BlockReceivedAndDeletedOps,omitempty"`
		StorageBlockReportOps int `json:"StorageBlockReportOps,omitempty"`
		TransactionsNumOps int `json:"TransactionsNumOps,omitempty"`
		TransactionsAvgTime float64 `json:"TransactionsAvgTime,omitempty"`
		SyncsNumOps int `json:"SyncsNumOps,omitempty"`
		SyncsAvgTime float64 `json:"SyncsAvgTime,omitempty"`
		TransactionsBatchedInSync int `json:"TransactionsBatchedInSync,omitempty"`
		BlockReportNumOps int `json:"BlockReportNumOps,omitempty"`
		BlockReportAvgTime float64 `json:"BlockReportAvgTime,omitempty"`
		CacheReportNumOps int `json:"CacheReportNumOps,omitempty"`
		CacheReportAvgTime float64 `json:"CacheReportAvgTime,omitempty"`
		SafeModeTime int `json:"SafeModeTime,omitempty"`
		FsImageLoadTime int `json:"FsImageLoadTime,omitempty"`
		GetEditNumOps int `json:"GetEditNumOps,omitempty"`
		GetEditAvgTime float64 `json:"GetEditAvgTime,omitempty"`
		GetImageNumOps int `json:"GetImageNumOps,omitempty"`
		GetImageAvgTime float64 `json:"GetImageAvgTime,omitempty"`
		PutImageNumOps int `json:"PutImageNumOps,omitempty"`
		PutImageAvgTime float64 `json:"PutImageAvgTime,omitempty"`
		NNRole string `json:"NNRole,omitempty"`
		HostAndPort string `json:"HostAndPort,omitempty"`
		SecurityEnabled bool `json:"SecurityEnabled,omitempty"`
		State string `json:"State,omitempty"`
		Total int64 `json:"Total,omitempty"`
		UpgradeFinalized bool `json:"UpgradeFinalized,omitempty"`
		ClusterID string `json:"ClusterId,omitempty"`
		Version string `json:"Version,omitempty"`
		Used int64 `json:"Used,omitempty"`
		Free int64 `json:"Free,omitempty"`
		Safemode string `json:"Safemode,omitempty"`
		NonDfsUsedSpace int64 `json:"NonDfsUsedSpace,omitempty"`
		PercentUsed float64 `json:"PercentUsed,omitempty"`
		BlockPoolUsedSpace int64 `json:"BlockPoolUsedSpace,omitempty"`
		PercentBlockPoolUsed float64 `json:"PercentBlockPoolUsed,omitempty"`
		PercentRemaining float64 `json:"PercentRemaining,omitempty"`
		CacheCapacity int `json:"CacheCapacity,omitempty"`
		CacheUsed int `json:"CacheUsed,omitempty"`
		TotalBlocks int `json:"TotalBlocks,omitempty"`
		TotalFiles int `json:"TotalFiles,omitempty"`
		NumberOfMissingBlocks int `json:"NumberOfMissingBlocks,omitempty"`
		LiveNodes string `json:"LiveNodes,omitempty"`
		DeadNodes string `json:"DeadNodes,omitempty"`
		DecomNodes string `json:"DecomNodes,omitempty"`
		BlockPoolID string `json:"BlockPoolId,omitempty"`
		NameDirStatuses string `json:"NameDirStatuses,omitempty"`
		NodeUsage string `json:"NodeUsage,omitempty"`
		NameJournalStatus string `json:"NameJournalStatus,omitempty"`
		JournalTransactionInfo string `json:"JournalTransactionInfo,omitempty"`
		NNStarted string `json:"NNStarted,omitempty"`
		CompileInfo string `json:"CompileInfo,omitempty"`
		CorruptFiles string `json:"CorruptFiles,omitempty"`
		DistinctVersionCount int `json:"DistinctVersionCount,omitempty"`
		DistinctVersions []struct {
			Key string `json:"key"`
			Value int `json:"value"`
		} `json:"DistinctVersions,omitempty"`
		SoftwareVersion string `json:"SoftwareVersion,omitempty"`
		RollingUpgradeStatus interface{} `json:"RollingUpgradeStatus,omitempty"`
		Threads int `json:"Threads,omitempty"`
		ElapsedTime int `json:"ElapsedTime,omitempty"`
		PercentComplete float64 `json:"PercentComplete,omitempty"`
		LoadingFsImageCount int `json:"LoadingFsImageCount,omitempty"`
		LoadingFsImageElapsedTime int `json:"LoadingFsImageElapsedTime,omitempty"`
		LoadingFsImageTotal int `json:"LoadingFsImageTotal,omitempty"`
		LoadingFsImagePercentComplete float64 `json:"LoadingFsImagePercentComplete,omitempty"`
		LoadingEditsCount int `json:"LoadingEditsCount,omitempty"`
		LoadingEditsElapsedTime int `json:"LoadingEditsElapsedTime,omitempty"`
		LoadingEditsTotal int `json:"LoadingEditsTotal,omitempty"`
		LoadingEditsPercentComplete float64 `json:"LoadingEditsPercentComplete,omitempty"`
		SavingCheckpointCount int `json:"SavingCheckpointCount,omitempty"`
		SavingCheckpointElapsedTime int `json:"SavingCheckpointElapsedTime,omitempty"`
		SavingCheckpointTotal int `json:"SavingCheckpointTotal,omitempty"`
		SavingCheckpointPercentComplete float64 `json:"SavingCheckpointPercentComplete,omitempty"`
		SafeModeCount int `json:"SafeModeCount,omitempty"`
		SafeModeElapsedTime int `json:"SafeModeElapsedTime,omitempty"`
		SafeModeTotal int `json:"SafeModeTotal,omitempty"`
		SafeModePercentComplete float64 `json:"SafeModePercentComplete,omitempty"`
		TagIsOutOfSync string `json:"tag.IsOutOfSync,omitempty"`
		QueuedEditsSize int `json:"QueuedEditsSize,omitempty"`
		LagTimeMillis int `json:"LagTimeMillis,omitempty"`
		CurrentLagTxns int `json:"CurrentLagTxns,omitempty"`
		TagHAState string `json:"tag.HAState,omitempty"`
		MissingBlocks int `json:"MissingBlocks,omitempty"`
		ExpiredHeartbeats int `json:"ExpiredHeartbeats,omitempty"`
		TransactionsSinceLastCheckpoint int `json:"TransactionsSinceLastCheckpoint,omitempty"`
		TransactionsSinceLastLogRoll int `json:"TransactionsSinceLastLogRoll,omitempty"`
		LastWrittenTransactionID int `json:"LastWrittenTransactionId,omitempty"`
		LastCheckpointTime int64 `json:"LastCheckpointTime,omitempty"`
		CapacityTotal int64 `json:"CapacityTotal,omitempty"`
		CapacityTotalGB float64 `json:"CapacityTotalGB,omitempty"`
		CapacityUsed int64 `json:"CapacityUsed,omitempty"`
		CapacityUsedGB float64 `json:"CapacityUsedGB,omitempty"`
		CapacityRemaining int64 `json:"CapacityRemaining,omitempty"`
		CapacityRemainingGB float64 `json:"CapacityRemainingGB,omitempty"`
		CapacityUsedNonDFS int64 `json:"CapacityUsedNonDFS,omitempty"`
		TotalLoad int `json:"TotalLoad,omitempty"`
		//Snapshots int `json:"Snapshots,omitempty"`
		BlocksTotal int `json:"BlocksTotal,omitempty"`
		FilesTotal int `json:"FilesTotal,omitempty"`
		PendingReplicationBlocks int `json:"PendingReplicationBlocks,omitempty"`
		UnderReplicatedBlocks int `json:"UnderReplicatedBlocks,omitempty"`
		CorruptBlocks int `json:"CorruptBlocks,omitempty"`
		ScheduledReplicationBlocks int `json:"ScheduledReplicationBlocks,omitempty"`
		PendingDeletionBlocks int `json:"PendingDeletionBlocks,omitempty"`
		ExcessBlocks int `json:"ExcessBlocks,omitempty"`
		PostponedMisreplicatedBlocks int `json:"PostponedMisreplSnapshottableDirectoriesicatedBlocks,omitempty"`
		PendingDataNodeMessageCount int `json:"PendingDataNodeMessageCount,omitempty"`
		MillisSinceLastLoadedEdits int `json:"MillisSinceLastLoadedEdits,omitempty"`
		BlockCapacity int `json:"BlockCapacity,omitempty"`
		StaleDataNodes int `json:"StaleDataNodes,omitempty"`
		SnapshotStats string `json:"SnapshotStats,omitempty"`
		MaxObjects int `json:"MaxObjects,omitempty"`
		BlockDeletionStartTime int64 `json:"BlockDeletionStartTime,omitempty"`
		FSState string `json:"FSState,omitempty"`
		NumLiveDataNodes int `json:"NumLiveDataNodes,omitempty"`
		NumDeadDataNodes int `json:"NumDeadDataNodes,omitempty"`
		NumDecomLiveDataNodes int `json:"NumDecomLiveDataNodes,omitempty"`
		NumDecomDeadDataNodes int `json:"NumDecomDeadDataNodes,omitempty"`
		NumDecommissioningDataNodes int `json:"NumDecommissioningDataNodes,omitempty"`
		NumStaleDataNodes int `json:"NumStaleDataNodes,omitempty"`
		NumStaleStorages int `json:"NumStaleStorages,omitempty"`
		CacheHit int `json:"CacheHit,omitempty"`
		CacheCleared int `json:"CacheCleared,omitempty"`
		CacheUpdated int `json:"CacheUpdated,omitempty"`
		TagPort string `json:"tag.port,omitempty"`
		ReceivedBytes int64 `json:"ReceivedBytes,omitempty"`
		SentBytes int64 `json:"SentBytes,omitempty"`
		RpcQueueTimeNumOps int64 `json:"RpcQueueTimeNumOps,omitempty"`
		RpcQueueTimeAvgTime float64 `json:"RpcQueueTimeAvgTime,omitempty"`
		RpcProcessingTimeNumOps int64 `json:"RpcProcessingTimeNumOps,omitempty"`
		RpcProcessingTimeAvgTime float64 `json:"RpcProcessingTimeAvgTime,omitempty"`
		RpcAuthenticationFailures int `json:"RpcAuthenticationFailures,omitempty"`
		RpcAuthenticationSuccesses int `json:"RpcAuthenticationSuccesses,omitempty"`
		RpcAuthorizationFailures int `json:"RpcAuthorizationFailures,omitempty"`
		RpcAuthorizationSuccesses int `json:"RpcAuthorizationSuccesses,omitempty"`
		NumOpenConnections int `json:"NumOpenConnections,omitempty"`
		CallQueueLength int `json:"CallQueueLength,omitempty"`
		LoginSuccessNumOps int `json:"LoginSuccessNumOps,omitempty"`
		LoginSuccessAvgTime float64 `json:"LoginSuccessAvgTime,omitempty"`
		LoginFailureNumOps int `json:"LoginFailureNumOps,omitempty"`
		LoginFailureAvgTime float64 `json:"LoginFailureAvgTime,omitempty"`
		GetGroupsNumOps int `json:"GetGroupsNumOps,omitempty"`
		GetGroupsAvgTime float64 `json:"GetGroupsAvgTime,omitempty"`
		StandbyExceptionNumOps int `json:"StandbyExceptionNumOps,omitempty"`
		StandbyExceptionAvgTime float64 `json:"StandbyExceptionAvgTime,omitempty"`
		GetServiceStatusNumOps int `json:"GetServiceStatusNumOps,omitempty"`
		GetServiceStatusAvgTime float64 `json:"GetServiceStatusAvgTime,omitempty"`
		MonitorHealthNumOps int `json:"MonitorHealthNumOps,omitempty"`
		MonitorHealthAvgTime float64 `json:"MonitorHealthAvgTime,omitempty"`
		SendHeartbeatNumOps int `json:"SendHeartbeatNumOps,omitempty"`
		SendHeartbeatAvgTime float64 `json:"SendHeartbeatAvgTime,omitempty"`
		TransitionToActiveNumOps int `json:"TransitionToActiveNumOps,omitempty"`
		TransitionToActiveAvgTime float64 `json:"TransitionToActiveAvgTime,omitempty"`
		RetriableExceptionNumOps int `json:"RetriableExceptionNumOps,omitempty"`
		RetriableExceptionAvgTime float64 `json:"RetriableExceptionAvgTime,omitempty"`
		VersionRequestNumOps int `json:"VersionRequestNumOps,omitempty"`
		VersionRequestAvgTime float64 `json:"VersionRequestAvgTime,omitempty"`
		RegisterDatanodeNumOps int `json:"RegisterDatanodeNumOps,omitempty"`
		RegisterDatanodeAvgTime float64 `json:"RegisterDatanodeAvgTime,omitempty"`
		GetFileInfoNumOps int `json:"GetFileInfoNumOps,omitempty"`
		GetFileInfoAvgTime float64 `json:"GetFileInfoAvgTime,omitempty"`
		BlockReceivedAndDeletedNumOps int `json:"BlockReceivedAndDeletedNumOps,omitempty"`
		BlockReceivedAndDeletedAvgTime float64 `json:"BlockReceivedAndDeletedAvgTime,omitempty"`
		GetListingNumOps int `json:"GetListingNumOps,omitempty"`
		GetListingAvgTime float64 `json:"GetListingAvgTime,omitempty"`
		GetBlockLocationsNumOps int64 `json:"GetBlockLocationsNumOps,omitempty"`
		GetBlockLocationsAvgTime float64 `json:"GetBlockLocationsAvgTime,omitempty"`
		GetServerDefaultsNumOps int `json:"GetServerDefaultsNumOps,omitempty"`
		GetServerDefaultsAvgTime float64 `json:"GetServerDefaultsAvgTime,omitempty"`
		SetSafeModeNumOps int `json:"SetSafeModeNumOps,omitempty"`
		SetSafeModeAvgTime float64 `json:"SetSafeModeAvgTime,omitempty"`
		RenewLeaseNumOps int `json:"RenewLeaseNumOps,omitempty"`
		RenewLeaseAvgTime float64 `json:"RenewLeaseAvgTime,omitempty"`
		RenameNumOps int `json:"RenameNumOps,omitempty"`
		RenameAvgTime float64 `json:"RenameAvgTime,omitempty"`
		CreateNumOps int `json:"CreateNumOps,omitempty"`
		CreateAvgTime float64 `json:"CreateAvgTime,omitempty"`
		AddBlockNumOps int `json:"AddBlockNumOps,omitempty"`
		AddBlockAvgTime float64 `json:"AddBlockAvgTime,omitempty"`
		FsyncNumOps int `json:"FsyncNumOps,omitempty"`
		FsyncAvgTime float64 `json:"FsyncAvgTime,omitempty"`
		CompleteNumOps int `json:"CompleteNumOps,omitempty"`
		CompleteAvgTime float64 `json:"CompleteAvgTime,omitempty"`
		MkdirsNumOps int `json:"MkdirsNumOps,omitempty"`
		MkdirsAvgTime float64 `json:"MkdirsAvgTime,omitempty"`
		SetPermissionNumOps int `json:"SetPermissionNumOps,omitempty"`
		SetPermissionAvgTime float64 `json:"SetPermissionAvgTime,omitempty"`
		DeleteNumOps int `json:"DeleteNumOps,omitempty"`
		DeleteAvgTime float64 `json:"DeleteAvgTime,omitempty"`
		RollEditLogNumOps int `json:"RollEditLogNumOps,omitempty"`
		RollEditLogAvgTime float64 `json:"RollEditLogAvgTime,omitempty"`
		GetContentSummaryNumOps int `json:"GetContentSummaryNumOps,omitempty"`
		GetContentSummaryAvgTime float64 `json:"GetContentSummaryAvgTime,omitempty"`
		GetAdditionalDatanodeNumOps int `json:"GetAdditionalDatanodeNumOps,omitempty"`
		GetAdditionalDatanodeAvgTime float64 `json:"GetAdditionalDatanodeAvgTime,omitempty"`
		UpdateBlockForPipelineNumOps int `json:"UpdateBlockForPipelineNumOps,omitempty"`
		UpdateBlockForPipelineAvgTime float64 `json:"UpdateBlockForPipelineAvgTime,omitempty"`
		UpdatePipelineNumOps int `json:"UpdatePipelineNumOps,omitempty"`
		UpdatePipelineAvgTime float64 `json:"UpdatePipelineAvgTime,omitempty"`
		SetReplicationNumOps int `json:"SetReplicationNumOps,omitempty"`
		SetReplicationAvgTime float64 `json:"SetReplicationAvgTime,omitempty"`
		CommitBlockSynchronizationNumOps int `json:"CommitBlockSynchronizationNumOps,omitempty"`
		CommitBlockSynchronizationAvgTime float64 `json:"CommitBlockSynchronizationAvgTime,omitempty"`
		GetFsStatsNumOps int `json:"GetFsStatsNumOps,omitempty"`
		GetFsStatsAvgTime float64 `json:"GetFsStatsAvgTime,omitempty"`
		GetDatanodeReportNumOps int `json:"GetDatanodeReportNumOps,omitempty"`
		GetDatanodeReportAvgTime float64 `json:"GetDatanodeReportAvgTime,omitempty"`
		ReportBadBlocksNumOps int `json:"ReportBadBlocksNumOps,omitempty"`
		ReportBadBlocksAvgTime float64 `json:"ReportBadBlocksAvgTime,omitempty"`
		Rename2NumOps int `json:"Rename2NumOps,omitempty"`
		Rename2AvgTime float64 `json:"Rename2AvgTime,omitempty"`
		GetEZForPathNumOps int `json:"GetEZForPathNumOps,omitempty"`
		GetEZForPathAvgTime float64 `json:"GetEZForPathAvgTime,omitempty"`
		FileNotFoundExceptionNumOps int `json:"FileNotFoundExceptionNumOps,omitempty"`
		FileNotFoundExceptionAvgTime float64 `json:"FileNotFoundExceptionAvgTime,omitempty"`
		AbandonBlockNumOps int `json:"AbandonBlockNumOps,omitempty"`
		AbandonBlockAvgTime float64 `json:"AbandonBlockAvgTime,omitempty"`
		TransitionToStandbyNumOps int `json:"TransitionToStandbyNumOps,omitempty"`
		TransitionToStandbyAvgTime float64 `json:"TransitionToStandbyAvgTime,omitempty"`
		AccessControlExceptionNumOps int `json:"AccessControlExceptionNumOps,omitempty"`
		AccessControlExceptionAvgTime float64 `json:"AccessControlExceptionAvgTime,omitempty"`
		LeaseExpiredExceptionNumOps int `json:"LeaseExpiredExceptionNumOps,omitempty"`
		LeaseExpiredExceptionAvgTime float64 `json:"LeaseExpiredExceptionAvgTime,omitempty"`
		ErrorReportNumOps int `json:"ErrorReportNumOps,omitempty"`
		ErrorReportAvgTime float64 `json:"ErrorReportAvgTime,omitempty"`
		IOExceptionNumOps int `json:"IOExceptionNumOps,omitempty"`
		IOExceptionAvgTime float64 `json:"IOExceptionAvgTime,omitempty"`
		NumActiveSources int `json:"NumActiveSources,omitempty"`
		NumAllSources int `json:"NumAllSources,omitempty"`
		NumActiveSinks int `json:"NumActiveSinks,omitempty"`
		NumAllSinks int `json:"NumAllSinks,omitempty"`
		SnapshotNumOps int `json:"SnapshotNumOps,omitempty"`
		SnapshotAvgTime float64 `json:"SnapshotAvgTime,omitempty"`
		PublishNumOps int `json:"PublishNumOps,omitempty"`
		PublishAvgTime float64 `json:"PublishAvgTime,omitempty"`
		DroppedPubAll int `json:"DroppedPubAll,omitempty"`
		BytesWritten int64 `json:"BytesWritten,omitempty"`
		BytesRead int64 `json:"BytesRead,omitempty"`
		BlocksWritten int `json:"BlocksWritten,omitempty"`
		BlocksRead int `json:"BlocksRead,omitempty"`
		BlocksReplicated int `json:"BlocksReplicated,omitempty"`
		BlocksRemoved int `json:"BlocksRemoved,omitempty"`
		BlocksVerified int `json:"BlocksVerified,omitempty"`
		BlockVerificationFailures int `json:"BlockVerificationFailures,omitempty"`
		BlocksCached int `json:"BlocksCached,omitempty"`
		BlocksUncached int `json:"BlocksUncached,omitempty"`
		ReadsFromLocalClient int `json:"ReadsFromLocalClient,omitempty"`
		ReadsFromRemoteClient int `json:"ReadsFromRemoteClient,omitempty"`
		WritesFromLocalClient int `json:"WritesFromLocalClient,omitempty"`
		WritesFromRemoteClient int `json:"WritesFromRemoteClient,omitempty"`
		BlocksGetLocalPathInfo int `json:"BlocksGetLocalPathInfo,omitempty"`
		RAMDiskBlocksWrite int `json:"RamDiskBlocksWrite,omitempty"`
		RAMDiskBlocksWriteFallback int `json:"RamDiskBlocksWriteFallback,omitempty"`
		RAMDiskBytesWrite int `json:"RamDiskBytesWrite,omitempty"`
		RAMDiskBlocksReadHits int `json:"RamDiskBlocksReadHits,omitempty"`
		RAMDiskBlocksEvicted int `json:"RamDiskBlocksEvicted,omitempty"`
		RAMDiskBlocksEvictedWithoutRead int `json:"RamDiskBlocksEvictedWithoutRead,omitempty"`
		RAMDiskBlocksEvictionWindowMsNumOps int `json:"RamDiskBlocksEvictionWindowMsNumOps,omitempty"`
		RAMDiskBlocksEvictionWindowMsAvgTime float64 `json:"RamDiskBlocksEvictionWindowMsAvgTime,omitempty"`
		RAMDiskBlocksLazyPersisted int `json:"RamDiskBlocksLazyPersisted,omitempty"`
		RAMDiskBlocksDeletedBeforeLazyPersisted int `json:"RamDiskBlocksDeletedBeforeLazyPersisted,omitempty"`
		RAMDiskBytesLazyPersisted int `json:"RamDiskBytesLazyPersisted,omitempty"`
		RAMDiskBlocksLazyPersistWindowMsNumOps int `json:"RamDiskBlocksLazyPersistWindowMsNumOps,omitempty"`
		RAMDiskBlocksLazyPersistWindowMsAvgTime float64 `json:"RamDiskBlocksLazyPersistWindowMsAvgTime,omitempty"`
		FsyncCount int `json:"FsyncCount,omitempty"`
		VolumeFailures int `json:"VolumeFailures,omitempty"`
		ReadBlockOpNumOps int `json:"ReadBlockOpNumOps,omitempty"`
		ReadBlockOpAvgTime float64 `json:"ReadBlockOpAvgTime,omitempty"`
		WriteBlockOpNumOps int `json:"WriteBlockOpNumOps,omitempty"`
		WriteBlockOpAvgTime float64 `json:"WriteBlockOpAvgTime,omitempty"`
		BlockChecksumOpNumOps int `json:"BlockChecksumOpNumOps,omitempty"`
		BlockChecksumOpAvgTime float64 `json:"BlockChecksumOpAvgTime,omitempty"`
		CopyBlockOpNumOps int `json:"CopyBlockOpNumOps,omitempty"`
		CopyBlockOpAvgTime float64 `json:"CopyBlockOpAvgTime,omitempty"`
		ReplaceBlockOpNumOps int `json:"ReplaceBlockOpNumOps,omitempty"`
		ReplaceBlockOpAvgTime float64 `json:"ReplaceBlockOpAvgTime,omitempty"`
		HeartbeatsNumOps int `json:"HeartbeatsNumOps,omitempty"`
		HeartbeatsAvgTime float64 `json:"HeartbeatsAvgTime,omitempty"`
		BlockReportsNumOps int `json:"BlockReportsNumOps,omitempty"`
		BlockReportsAvgTime float64 `json:"BlockReportsAvgTime,omitempty"`
		CacheReportsNumOps int `json:"CacheReportsNumOps,omitempty"`
		CacheReportsAvgTime float64 `json:"CacheReportsAvgTime,omitempty"`
		PacketAckRoundTripTimeNanosNumOps int64 `json:"PacketAckRoundTripTimeNanosNumOps,omitempty"`
		PacketAckRoundTripTimeNanosAvgTime float64 `json:"PacketAckRoundTripTimeNanosAvgTime,omitempty"`
		FlushNanosNumOps int64 `json:"FlushNanosNumOps,omitempty"`
		FlushNanosAvgTime float64 `json:"FlushNanosAvgTime,omitempty"`
		FsyncNanosNumOps int `json:"FsyncNanosNumOps,omitempty"`
		FsyncNanosAvgTime float64 `json:"FsyncNanosAvgTime,omitempty"`
		SendDataPacketBlockedOnNetworkNanosNumOps int64 `json:"SendDataPacketBlockedOnNetworkNanosNumOps,omitempty"`
		SendDataPacketBlockedOnNetworkNanosAvgTime float64 `json:"SendDataPacketBlockedOnNetworkNanosAvgTime,omitempty"`
		SendDataPacketTransferNanosNumOps int64 `json:"SendDataPacketTransferNanosNumOps,omitempty"`
		SendDataPacketTransferNanosAvgTime float64 `json:"SendDataPacketTransferNanosAvgTime,omitempty"`
		Remaining int64 `json:"Remaining,omitempty"`
		StorageInfo string `json:"StorageInfo,omitempty"`
		Capacity int64 `json:"Capacity,omitempty"`
		DfsUsed int64 `json:"DfsUsed,omitempty"`
		NumFailedVolumes int `json:"NumFailedVolumes,omitempty"`
		NumBlocksCached int `json:"NumBlocksCached,omitempty"`
		NumBlocksFailedToCache int `json:"NumBlocksFailedToCache,omitempty"`
		NumBlocksFailedToUncache int `json:"NumBlocksFailedToUncache,omitempty"`
		XceiverCount int `json:"XceiverCount,omitempty"`
		RpcPort string `json:"RpcPort,omitempty"`
		HTTPPort interface{} `json:"HttpPort,omitempty"`
		NamenodeAddresses string `json:"NamenodeAddresses,omitempty"`
		VolumeInfo string `json:"VolumeInfo,omitempty"`
		InitReplicaRecoveryNumOps int `json:"InitReplicaRecoveryNumOps,omitempty"`
		InitReplicaRecoveryAvgTime float64 `json:"InitReplicaRecoveryAvgTime,omitempty"`
		UpdateReplicaUnderRecoveryNumOps int `json:"UpdateReplicaUnderRecoveryNumOps,omitempty"`
		UpdateReplicaUnderRecoveryAvgTime float64 `json:"UpdateReplicaUnderRecoveryAvgTime,omitempty"`
		GetReplicaVisibleLengthNumOps int `json:"GetReplicaVisibleLengthNumOps,omitempty"`
		GetReplicaVisibleLengthAvgTime float64 `json:"GetReplicaVisibleLengthAvgTime,omitempty"`
	} `json:"beans"`
}
