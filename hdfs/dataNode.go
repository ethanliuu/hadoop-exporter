package hdfs

type DataNode struct {
	Beans []struct {
		Name string `json:"name"`
		ModelerType string `json:"modelerType"`
		TagSessionID interface{} `json:"tag.SessionId,omitempty"`
		TagContext string `json:"tag.Context,omitempty"`
		TagHostname string `json:"tag.Hostname,omitempty"`
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
		TagProcessName string `json:"tag.ProcessName,omitempty"`
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
		Remaining int64 `json:"Remaining,omitempty"`
		StorageInfo string `json:"StorageInfo,omitempty"`
		Capacity int64 `json:"Capacity,omitempty"`
		DfsUsed int64 `json:"DfsUsed,omitempty"`
		CacheCapacity int `json:"CacheCapacity,omitempty"`
		CacheUsed int `json:"CacheUsed,omitempty"`
		NumFailedVolumes int `json:"NumFailedVolumes,omitempty"`
		NumBlocksCached int `json:"NumBlocksCached,omitempty"`
		NumBlocksFailedToCache int `json:"NumBlocksFailedToCache,omitempty"`
		NumBlocksFailedToUncache int `json:"NumBlocksFailedToUncache,omitempty"`
		LoginSuccessNumOps int `json:"LoginSuccessNumOps,omitempty"`
		LoginSuccessAvgTime float64 `json:"LoginSuccessAvgTime,omitempty"`
		LoginFailureNumOps int `json:"LoginFailureNumOps,omitempty"`
		LoginFailureAvgTime float64 `json:"LoginFailureAvgTime,omitempty"`
		GetGroupsNumOps int `json:"GetGroupsNumOps,omitempty"`
		GetGroupsAvgTime float64 `json:"GetGroupsAvgTime,omitempty"`
		XceiverCount int `json:"XceiverCount,omitempty"`
		Version string `json:"Version,omitempty"`
		RpcPort string `json:"RpcPort,omitempty"`
		HTTPPort interface{} `json:"HttpPort,omitempty"`
		NamenodeAddresses string `json:"NamenodeAddresses,omitempty"`
		VolumeInfo string `json:"VolumeInfo,omitempty"`
		ClusterID string `json:"ClusterId,omitempty"`
		TagPort string `json:"tag.port,omitempty"`
		IOExceptionNumOps int `json:"IOExceptionNumOps,omitempty"`
		IOExceptionAvgTime float64 `json:"IOExceptionAvgTime,omitempty"`
		InitReplicaRecoveryNumOps int `json:"InitReplicaRecoveryNumOps,omitempty"`
		InitReplicaRecoveryAvgTime float64 `json:"InitReplicaRecoveryAvgTime,omitempty"`
		UpdateReplicaUnderRecoveryNumOps int `json:"UpdateReplicaUnderRecoveryNumOps,omitempty"`
		UpdateReplicaUnderRecoveryAvgTime float64 `json:"UpdateReplicaUnderRecoveryAvgTime,omitempty"`
		GetReplicaVisibleLengthNumOps int `json:"GetReplicaVisibleLengthNumOps,omitempty"`
		GetReplicaVisibleLengthAvgTime float64 `json:"GetReplicaVisibleLengthAvgTime,omitempty"`
		ReceivedBytes int `json:"ReceivedBytes,omitempty"`
		SentBytes int `json:"SentBytes,omitempty"`
		RpcQueueTimeNumOps int `json:"RpcQueueTimeNumOps,omitempty"`
		RpcQueueTimeAvgTime float64 `json:"RpcQueueTimeAvgTime,omitempty"`
		RpcProcessingTimeNumOps int `json:"RpcProcessingTimeNumOps,omitempty"`
		RpcProcessingTimeAvgTime float64 `json:"RpcProcessingTimeAvgTime,omitempty"`
		RpcAuthenticationFailures int `json:"RpcAuthenticationFailures,omitempty"`
		RpcAuthenticationSuccesses int `json:"RpcAuthenticationSuccesses,omitempty"`
		RpcAuthorizationFailures int `json:"RpcAuthorizationFailures,omitempty"`
		RpcAuthorizationSuccesses int `json:"RpcAuthorizationSuccesses,omitempty"`
		NumOpenConnections int `json:"NumOpenConnections,omitempty"`
		CallQueueLength int `json:"CallQueueLength,omitempty"`
		NumActiveSources int `json:"NumActiveSources,omitempty"`
		NumAllSources int `json:"NumAllSources,omitempty"`
		NumActiveSinks int `json:"NumActiveSinks,omitempty"`
		NumAllSinks int `json:"NumAllSinks,omitempty"`
		SnapshotNumOps int `json:"SnapshotNumOps,omitempty"`
		SnapshotAvgTime float64 `json:"SnapshotAvgTime,omitempty"`
		PublishNumOps int `json:"PublishNumOps,omitempty"`
		PublishAvgTime float64 `json:"PublishAvgTime,omitempty"`
		DroppedPubAll int `json:"DroppedPubAll,omitempty"`
	} `json:"beans"`
}