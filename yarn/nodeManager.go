package yarn

type NodeManager struct {
	Beans []struct {
		Name string `json:"name"`
		ModelerType string `json:"modelerType"`
		TagPort string `json:"tag.port,omitempty"`
		TagContext string `json:"tag.Context,omitempty"`
		TagHostname string `json:"tag.Hostname,omitempty"`
		ReceivedBytes int64 `json:"ReceivedBytes,omitempty"`
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
		StartContainersNumOps int `json:"StartContainersNumOps,omitempty"`
		StartContainersAvgTime float64 `json:"StartContainersAvgTime,omitempty"`
		StopContainersNumOps int `json:"StopContainersNumOps,omitempty"`
		StopContainersAvgTime float64 `json:"StopContainersAvgTime,omitempty"`
		HeartbeatNumOps int `json:"HeartbeatNumOps,omitempty"`
		HeartbeatAvgTime float64 `json:"HeartbeatAvgTime,omitempty"`
		TagProcessName string `json:"tag.ProcessName,omitempty"`
		TagSessionID interface{} `json:"tag.SessionId,omitempty"`
		MemNonHeapUsedM float64 `json:"MemNonHeapUsedM,omitempty"`
		MemNonHeapCommittedM float64 `json:"MemNonHeapCommittedM,omitempty"`
		MemNonHeapMaxM float64 `json:"MemNonHeapMaxM,omitempty"`
		MemHeapUsedM float64 `json:"MemHeapUsedM,omitempty"`
		MemHeapCommittedM float64 `json:"MemHeapCommittedM,omitempty"`
		MemHeapMaxM float64 `json:"MemHeapMaxM,omitempty"`
		MemMaxM float64 `json:"MemMaxM,omitempty"`
		GcCount int `json:"GcCount,omitempty"`
		GcTimeMillis int `json:"GcTimeMillis,omitempty"`
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
		LoginSuccessNumOps int `json:"LoginSuccessNumOps,omitempty"`
		LoginSuccessAvgTime float64 `json:"LoginSuccessAvgTime,omitempty"`
		LoginFailureNumOps int `json:"LoginFailureNumOps,omitempty"`
		LoginFailureAvgTime float64 `json:"LoginFailureAvgTime,omitempty"`
		GetGroupsNumOps int `json:"GetGroupsNumOps,omitempty"`
		GetGroupsAvgTime float64 `json:"GetGroupsAvgTime,omitempty"`
		ShuffleOutputBytes int64 `json:"ShuffleOutputBytes,omitempty"`
		ShuffleOutputsFailed int `json:"ShuffleOutputsFailed,omitempty"`
		ShuffleOutputsOK int `json:"ShuffleOutputsOK,omitempty"`
		ShuffleConnections int `json:"ShuffleConnections,omitempty"`
		NumActiveSources int `json:"NumActiveSources,omitempty"`
		NumAllSources int `json:"NumAllSources,omitempty"`
		NumActiveSinks int `json:"NumActiveSinks,omitempty"`
		NumAllSinks int `json:"NumAllSinks,omitempty"`
		SnapshotNumOps int `json:"SnapshotNumOps,omitempty"`
		SnapshotAvgTime float64 `json:"SnapshotAvgTime,omitempty"`
		PublishNumOps int `json:"PublishNumOps,omitempty"`
		PublishAvgTime float64 `json:"PublishAvgTime,omitempty"`
		DroppedPubAll int `json:"DroppedPubAll,omitempty"`
		AllocatedVCores int `json:"AllocatedVCores,omitempty"`
		AvailableVCores int `json:"AvailableVCores,omitempty"`
	} `json:"beans"`
}