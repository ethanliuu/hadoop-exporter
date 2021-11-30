package yarn

const (
	namespace    = "hadoop"
	subNamespace = "yarn_resourcemanager"
)

type RsmNodem struct {
	Beans []struct {
		Name                                       string      `json:"name"`
		ModelerType                                string      `json:"modelerType"`
		LiveNodeManagers                           string      `json:"LiveNodeManagers,omitempty"`
		TagPort                                    string      `json:"tag.port,omitempty"`
		TagContext                                 string      `json:"tag.Context,omitempty"`
		TagHostname                                string      `json:"tag.Hostname,omitempty"`
		ReceivedBytes                              int         `json:"ReceivedBytes,omitempty"`
		SentBytes                                  int         `json:"SentBytes,omitempty"`
		RpcQueueTimeNumOps                         int         `json:"RpcQueueTimeNumOps,omitempty"`
		RpcQueueTimeAvgTime                        float64     `json:"RpcQueueTimeAvgTime,omitempty"`
		RpcProcessingTimeNumOps                    int         `json:"RpcProcessingTimeNumOps,omitempty"`
		RpcProcessingTimeAvgTime                   float64     `json:"RpcProcessingTimeAvgTime,omitempty"`
		RpcAuthenticationFailures                  int         `json:"RpcAuthenticationFailures,omitempty"`
		RpcAuthenticationSuccesses                 int         `json:"RpcAuthenticationSuccesses,omitempty"`
		RpcAuthorizationFailures                   int         `json:"RpcAuthorizationFailures,omitempty"`
		RpcAuthorizationSuccesses                  int         `json:"RpcAuthorizationSuccesses,omitempty"`
		NumOpenConnections                         int         `json:"NumOpenConnections,omitempty"`
		CallQueueLength                            int         `json:"CallQueueLength,omitempty"`
		GetClusterMetricsNumOps                    int         `json:"GetClusterMetricsNumOps,omitempty"`
		GetClusterMetricsAvgTime                   float64     `json:"GetClusterMetricsAvgTime,omitempty"`
		GetNewApplicationNumOps                    int         `json:"GetNewApplicationNumOps,omitempty"`
		GetNewApplicationAvgTime                   float64     `json:"GetNewApplicationAvgTime,omitempty"`
		SubmitApplicationNumOps                    int         `json:"SubmitApplicationNumOps,omitempty"`
		SubmitApplicationAvgTime                   float64     `json:"SubmitApplicationAvgTime,omitempty"`
		GetApplicationReportNumOps                 int         `json:"GetApplicationReportNumOps,omitempty"`
		GetApplicationReportAvgTime                float64     `json:"GetApplicationReportAvgTime,omitempty"`
		GetApplicationsNumOps                      int         `json:"GetApplicationsNumOps,omitempty"`
		GetApplicationsAvgTime                     float64     `json:"GetApplicationsAvgTime,omitempty"`
		ApplicationNotFoundExceptionNumOps         int         `json:"ApplicationNotFoundExceptionNumOps,omitempty"`
		ApplicationNotFoundExceptionAvgTime        float64     `json:"ApplicationNotFoundExceptionAvgTime,omitempty"`
		LoginSuccessNumOps                         int         `json:"LoginSuccessNumOps,omitempty"`
		LoginSuccessAvgTime                        float64     `json:"LoginSuccessAvgTime,omitempty"`
		LoginFailureNumOps                         int         `json:"LoginFailureNumOps,omitempty"`
		LoginFailureAvgTime                        float64     `json:"LoginFailureAvgTime,omitempty"`
		GetGroupsNumOps                            int         `json:"GetGroupsNumOps,omitempty"`
		GetGroupsAvgTime                           float64     `json:"GetGroupsAvgTime,omitempty"`
		NumActiveSources                           int         `json:"NumActiveSources,omitempty"`
		NumAllSources                              int         `json:"NumAllSources,omitempty"`
		NumActiveSinks                             int         `json:"NumActiveSinks,omitempty"`
		NumAllSinks                                int         `json:"NumAllSinks,omitempty"`
		SnapshotNumOps                             int         `json:"SnapshotNumOps,omitempty"`
		SnapshotAvgTime                            float64     `json:"SnapshotAvgTime,omitempty"`
		PublishNumOps                              int         `json:"PublishNumOps,omitempty"`
		PublishAvgTime                             float64     `json:"PublishAvgTime,omitempty"`
		DroppedPubAll                              int         `json:"DroppedPubAll,omitempty"`
		TagQueue                                   string      `json:"tag.Queue,omitempty"`
		Running0                                   int         `json:"running_0,omitempty"`
		Running60                                  int         `json:"running_60,omitempty"`
		Running300                                 int         `json:"running_300,omitempty"`
		Running1440                                int         `json:"running_1440,omitempty"`
		FairShareMB                                int         `json:"FairShareMB,omitempty"`
		FairShareVCores                            int         `json:"FairShareVCores,omitempty"`
		SteadyFairShareMB                          int         `json:"SteadyFairShareMB,omitempty"`
		SteadyFairShareVCores                      int         `json:"SteadyFairShareVCores,omitempty"`
		MinShareMB                                 int         `json:"MinShareMB,omitempty"`
		MinShareVCores                             int         `json:"MinShareVCores,omitempty"`
		MaxShareMB                                 int         `json:"MaxShareMB,omitempty"`
		MaxShareVCores                             int         `json:"MaxShareVCores,omitempty"`
		AppsSubmitted                              int         `json:"AppsSubmitted,omitempty"`
		AppsRunning                                int         `json:"AppsRunning,omitempty"`
		AppsPending                                int         `json:"AppsPending,omitempty"`
		AppsCompleted                              int         `json:"AppsCompleted,omitempty"`
		AppsKilled                                 int         `json:"AppsKilled,omitempty"`
		AppsFailed                                 int         `json:"AppsFailed,omitempty"`
		AllocatedMB                                int         `json:"AllocatedMB,omitempty"`
		AllocatedVCores                            int         `json:"AllocatedVCores,omitempty"`
		AllocatedContainers                        int         `json:"AllocatedContainers,omitempty"`
		AggregateContainersAllocated               int         `json:"AggregateContainersAllocated,omitempty"`
		AggregateContainersReleased                int         `json:"AggregateContainersReleased,omitempty"`
		AvailableMB                                int         `json:"AvailableMB,omitempty"`
		AvailableVCores                            int         `json:"AvailableVCores,omitempty"`
		PendingMB                                  int         `json:"PendingMB,omitempty"`
		PendingVCores                              int         `json:"PendingVCores,omitempty"`
		PendingContainers                          int         `json:"PendingContainers,omitempty"`
		ReservedMB                                 int         `json:"ReservedMB,omitempty"`
		ReservedVCores                             int         `json:"ReservedVCores,omitempty"`
		ReservedContainers                         int         `json:"ReservedContainers,omitempty"`
		ActiveUsers                                int         `json:"ActiveUsers,omitempty"`
		ActiveApplications                         int         `json:"ActiveApplications,omitempty"`
		TagUser                                    string      `json:"tag.User,omitempty"`
		TagProcessName                             string      `json:"tag.ProcessName,omitempty"`
		TagSessionID                               interface{} `json:"tag.SessionId,omitempty"`
		MemNonHeapUsedM                            float64     `json:"MemNonHeapUsedM,omitempty"`
		MemNonHeapCommittedM                       float64     `json:"MemNonHeapCommittedM,omitempty"`
		MemNonHeapMaxM                             float64     `json:"MemNonHeapMaxM,omitempty"`
		MemHeapUsedM                               float64     `json:"MemHeapUsedM,omitempty"`
		MemHeapCommittedM                          float64     `json:"MemHeapCommittedM,omitempty"`
		MemHeapMaxM                                float64     `json:"MemHeapMaxM,omitempty"`
		MemMaxM                                    float64     `json:"MemMaxM,omitempty"`
		GcCount                                    int         `json:"GcCount,omitempty"`
		GcTimeMillis                               int         `json:"GcTimeMillis,omitempty"`
		ThreadsNew                                 int         `json:"ThreadsNew,omitempty"`
		ThreadsRunnable                            int         `json:"ThreadsRunnable,omitempty"`
		ThreadsBlocked                             int         `json:"ThreadsBlocked,omitempty"`
		ThreadsWaiting                             int         `json:"ThreadsWaiting,omitempty"`
		ThreadsTimedWaiting                        int         `json:"ThreadsTimedWaiting,omitempty"`
		ThreadsTerminated                          int         `json:"ThreadsTerminated,omitempty"`
		LogFatal                                   int         `json:"LogFatal,omitempty"`
		LogError                                   int         `json:"LogError,omitempty"`
		LogWarn                                    int         `json:"LogWarn,omitempty"`
		LogInfo                                    int         `json:"LogInfo,omitempty"`
		RegisterApplicationMasterNumOps            int         `json:"RegisterApplicationMasterNumOps,omitempty"`
		RegisterApplicationMasterAvgTime           float64     `json:"RegisterApplicationMasterAvgTime,omitempty"`
		AllocateNumOps                             int         `json:"AllocateNumOps,omitempty"`
		AllocateAvgTime                            float64     `json:"AllocateAvgTime,omitempty"`
		FinishApplicationMasterNumOps              int         `json:"FinishApplicationMasterNumOps,omitempty"`
		FinishApplicationMasterAvgTime             float64     `json:"FinishApplicationMasterAvgTime,omitempty"`
		ApplicationAttemptNotFoundExceptionNumOps  int         `json:"ApplicationAttemptNotFoundExceptionNumOps,omitempty"`
		ApplicationAttemptNotFoundExceptionAvgTime float64     `json:"ApplicationAttemptNotFoundExceptionAvgTime,omitempty"`
		TagFSOpDurations                           string      `json:"tag.FSOpDurations,omitempty"`
		ContinuousSchedulingRunNumOps              int         `json:"ContinuousSchedulingRunNumOps,omitempty"`
		ContinuousSchedulingRunAvgTime             float64     `json:"ContinuousSchedulingRunAvgTime,omitempty"`
		ContinuousSchedulingRunStdevTime           float64     `json:"ContinuousSchedulingRunStdevTime,omitempty"`
		ContinuousSchedulingRunIMinTime            float64     `json:"ContinuousSchedulingRunIMinTime,omitempty"`
		ContinuousSchedulingRunIMaxTime            float64     `json:"ContinuousSchedulingRunIMaxTime,omitempty"`
		ContinuousSchedulingRunMinTime             float64     `json:"ContinuousSchedulingRunMinTime,omitempty"`
		ContinuousSchedulingRunMaxTime             float64     `json:"ContinuousSchedulingRunMaxTime,omitempty"`
		NodeUpdateCallNumOps                       int         `json:"NodeUpdateCallNumOps,omitempty"`
		NodeUpdateCallAvgTime                      float64     `json:"NodeUpdateCallAvgTime,omitempty"`
		NodeUpdateCallStdevTime                    float64     `json:"NodeUpdateCallStdevTime,omitempty"`
		NodeUpdateCallIMinTime                     float64     `json:"NodeUpdateCallIMinTime,omitempty"`
		NodeUpdateCallIMaxTime                     float64     `json:"NodeUpdateCallIMaxTime,omitempty"`
		NodeUpdateCallMinTime                      float64     `json:"NodeUpdateCallMinTime,omitempty"`
		NodeUpdateCallMaxTime                      float64     `json:"NodeUpdateCallMaxTime,omitempty"`
		UpdateThreadRunNumOps                      int         `json:"UpdateThreadRunNumOps,omitempty"`
		UpdateThreadRunAvgTime                     float64     `json:"UpdateThreadRunAvgTime,omitempty"`
		UpdateThreadRunStdevTime                   float64     `json:"UpdateThreadRunStdevTime,omitempty"`
		UpdateThreadRunIMinTime                    float64     `json:"UpdateThreadRunIMinTime,omitempty"`
		UpdateThreadRunIMaxTime                    float64     `json:"UpdateThreadRunIMaxTime,omitempty"`
		UpdateThreadRunMinTime                     float64     `json:"UpdateThreadRunMinTime,omitempty"`
		UpdateThreadRunMaxTime                     float64     `json:"UpdateThreadRunMaxTime,omitempty"`
		UpdateCallNumOps                           int         `json:"UpdateCallNumOps,omitempty"`
		UpdateCallAvgTime                          float64     `json:"UpdateCallAvgTime,omitempty"`
		UpdateCallStdevTime                        float64     `json:"UpdateCallStdevTime,omitempty"`
		UpdateCallIMinTime                         float64     `json:"UpdateCallIMinTime,omitempty"`
		UpdateCallIMaxTime                         float64     `json:"UpdateCallIMaxTime,omitempty"`
		UpdateCallMinTime                          float64     `json:"UpdateCallMinTime,omitempty"`
		UpdateCallMaxTime                          float64     `json:"UpdateCallMaxTime,omitempty"`
		PreemptCallNumOps                          int         `json:"PreemptCallNumOps,omitempty"`
		PreemptCallAvgTime                         float64     `json:"PreemptCallAvgTime,omitempty"`
		PreemptCallStdevTime                       float64     `json:"PreemptCallStdevTime,omitempty"`
		PreemptCallIMinTime                        float64     `json:"PreemptCallIMinTime,omitempty"`
		PreemptCallIMaxTime                        float64     `json:"PreemptCallIMaxTime,omitempty"`
		PreemptCallMinTime                         float64     `json:"PreemptCallMinTime,omitempty"`
		PreemptCallMaxTime                         float64     `json:"PreemptCallMaxTime,omitempty"`
		NodeHeartbeatNumOps                        int         `json:"NodeHeartbeatNumOps,omitempty"`
		NodeHeartbeatAvgTime                       float64     `json:"NodeHeartbeatAvgTime,omitempty"`
		RegisterNodeManagerNumOps                  int         `json:"RegisterNodeManagerNumOps,omitempty"`
		RegisterNodeManagerAvgTime                 float64     `json:"RegisterNodeManagerAvgTime,omitempty"`
		TagClusterMetrics                          string      `json:"tag.ClusterMetrics,omitempty"`
		NumActiveNMs                               int         `json:"NumActiveNMs,omitempty"`
		NumDecommissionedNMs                       int         `json:"NumDecommissionedNMs,omitempty"`
		NumLostNMs                                 int         `json:"NumLostNMs,omitempty"`
		NumUnhealthyNMs                            int         `json:"NumUnhealthyNMs,omitempty"`
		NumRebootedNMs                             int         `json:"NumRebootedNMs,omitempty"`
		// nodemanager
		ContainersLaunched int `json:"ContainersLaunched,omitempty"`
		ContainersCompleted int `json:"ContainersCompleted,omitempty"`
		ContainersFailed int `json:"ContainersFailed,omitempty"`
		ContainersKilled int `json:"ContainersKilled,omitempty"`
		ContainersIniting int `json:"ContainersIniting,omitempty"`
		ContainersRunning int `json:"ContainersRunning,omitempty"`
		AllocatedGB int `json:"AllocatedGB,omitempty"`
		AvailableGB int `json:"AvailableGB,omitempty"`
	} `json:"beans"`
}

type LiveNodeManagers struct {
	HostName           string `json:"HostName"`
	Rack               string `json:"Rack"`
	State              string `json:"State"`
	NodeID             string `json:"NodeId"`
	NodeHTTPAddress    string `json:"NodeHTTPAddress"`
	LastHealthUpdate   int64  `json:"LastHealthUpdate"`
	HealthReport       string `json:"HealthReport"`
	NodeManagerVersion string `json:"NodeManagerVersion"`
	NumContainers      int    `json:"NumContainers"`
	UsedMemoryMB       int    `json:"UsedMemoryMB"`
	AvailableMemoryMB  int    `json:"AvailableMemoryMB"`
}

//
//
//type Expoter struct {
//	url                                            string
//	name                                           prometheus.Gauge
//	modelerType                                    prometheus.Gauge
//	tagPort 	prometheus.Gauge
//	tagQueue                                       prometheus.Gauge
//	tagContext                                     prometheus.Gauge
//	tagHostName                                    prometheus.Gauge
//	amResourceLimitMB                              prometheus.Gauge
//	amResourceLimitVCores                          prometheus.Gauge
//	usedAMResourceMB                               prometheus.Gauge
//	usedAMResourceVCores                           prometheus.Gauge
//	usedCapacity                                   prometheus.Gauge
//	absoluteUsedCapacity                           prometheus.Gauge
//	guaranteedMB                                   prometheus.Gauge
//	guaranteedVCores                               prometheus.Gauge
//	maxCapacityMB                                  prometheus.Gauge
//	maxCapacityVCores                              prometheus.Gauge
//	guaranteedCapacity                             prometheus.Gauge
//	guaranteedAbsoluteCapacity                     prometheus.Gauge
//	maxCapacity                                    prometheus.Gauge
//	maxAbsoluteCapacity                            prometheus.Gauge
//	appsSubmitted                                  prometheus.Gauge
//	appsRunning                                    prometheus.Gauge
//	appsPending                                    prometheus.Gauge
//	appsCompleted                                  prometheus.Gauge
//	appsKilled                                     prometheus.Gauge
//	appsFailed                                     prometheus.Gauge
//	aggregateNodeLocalContainersAllocated          prometheus.Gauge
//	aggregateRackLocalContainersAllocated          prometheus.Gauge
//	aggregateOffSwitchContainersAllocated          prometheus.Gauge
//	aggregateContainersPreempted                   prometheus.Gauge
//	aggregateMemoryMBSecondsPreempted              prometheus.Gauge
//	aggregateVcoreSecondsPreempted                 prometheus.Gauge
//	activeUsers                                    prometheus.Gauge
//	activeApplications                             prometheus.Gauge
//	appAttemptFirstContainerAllocationDelayNumOps  prometheus.Gauge
//	appAttemptFirstContainerAllocationDelayAvgTime prometheus.Gauge
//	aggregateMemoryMBPreempted                     prometheus.Gauge
//	aggregateVcoresPreempted                       prometheus.Gauge
//	allocatedMB                                    prometheus.Gauge
//	allocatedVCores                                prometheus.Gauge
//	allocatedContainers                            prometheus.Gauge
//	aggregateContainersAllocated                  prometheus.Gauge
//	aggregateContainersReleased                    prometheus.Gauge
//	availableMB                                    prometheus.Gauge
//	availableVCores                                prometheus.Gauge
//	pendingMB                                      prometheus.Gauge
//	pendingVCores                                  prometheus.Gauge
//	pendingContainers                              prometheus.Gauge
//	reservedMB                                     prometheus.Gauge
//	reservedVCores                                 prometheus.Gauge
//	reservedContainers                             prometheus.Gauge
//}
//func newServerMetric(metricName string, docString string, labels []string) *prometheus.Desc {
//	return prometheus.NewDesc(
//		prometheus.BuildFQName(namespace, "server", metricName),
//		docString, labels, nil,
//	)
//}
//
//func NewExporter(url string) *Expoter {
//	return &Expoter{
//		url: url,
//		name: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "name",
//			Help:      "name",
//		}),
//		modelerType: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "modelerType",
//			Help:      "modelerType",
//		}),
//		tagQueue: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "tagQueue",
//			Help:      "tagQueue",
//		}),
//		tagContext: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "tagContext",
//			Help:      "tagContext",
//		}),
//		tagHostName: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "tagHostName",
//			Help:      "tagHostName",
//		}),
//		amResourceLimitMB: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "amResourceLimitMB",
//			Help:      "amResourceLimitMB",
//		}),
//		amResourceLimitVCores: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "amResourceLimitVCores",
//			Help:      "amResourceLimitVCores",
//		}),
//		usedAMResourceMB: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "usedAMResourceMB",
//			Help:      "usedAMResourceMB",
//		}),
//		usedAMResourceVCores: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "usedAMResourceVCores",
//			Help:      "usedAMResourceVCores",
//		}),
//		usedCapacity: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "usedCapacity",
//			Help:      "Current used capacity across all DataNodes in bytesn",
//		}),
//		absoluteUsedCapacity: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "absoluteUsedCapacity",
//			Help:      "absoluteUsedCapacity",
//		}),
//		guaranteedMB: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "guaranteedMB",
//			Help:      "guaranteedMB",
//		}),
//		guaranteedVCores: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "guaranteedVCores",
//			Help:      "guaranteedVCores",
//		}),
//		maxCapacityMB: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "maxCapacityMB",
//			Help:      "maxCapacityMB",
//		}),
//		maxCapacityVCores: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "maxCapacityVCores",
//			Help:      "maxCapacityVCores",
//		}),
//		guaranteedCapacity: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "guaranteedCapacity",
//			Help:      "guaranteedCapacity",
//		}),
//		guaranteedAbsoluteCapacity: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "guaranteedAbsoluteCapacity",
//			Help:      "guaranteedAbsoluteCapacity",
//		}),
//		maxCapacity: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "maxCapacity",
//			Help:      "maxCapacity",
//		}),
//		maxAbsoluteCapacity: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "maxAbsoluteCapacity",
//			Help:      "maxAbsoluteCapacity",
//		}),
//		appsSubmitted: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "appsSubmitted",
//			Help:      "Current number of running applications",
//		}),
//		appsRunning: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "appsRunning",
//			Help:      "appsRunning",
//		}),
//		appsPending: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "appsPending",
//			Help:      "Current number of applications that have not yet been assigned by any containers",
//		}),
//		appsCompleted: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "appsCompleted",
//			Help:      "Total number of completed applications",
//		}),
//		appsKilled: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "appsKilled",
//			Help:      "Total number of killed applications",
//		}),
//		appsFailed: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "appsFailed",
//			Help:      "Total number of failed applications",
//		}),
//		aggregateNodeLocalContainersAllocated: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aggregateNodeLocalContainersAllocated",
//			Help:      "Total number of node local containers allocated",
//		}),
//		aggregateRackLocalContainersAllocated: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aggregateRackLocalContainersAllocated",
//			Help:      "Total number of rack local containers allocated",
//		}),
//		aggregateOffSwitchContainersAllocated: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aggregateOffSwitchContainersAllocated",
//			Help:      "Total number of off switch containers allocated",
//		}),
//		aggregateContainersPreempted: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aggregateContainersPreempted",
//			Help:      "aggregateContainersPreempted",
//		}),
//		aggregateMemoryMBSecondsPreempted: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aggregateMemoryMBSecondsPreempted",
//			Help:      "aggregateMemoryMBSecondsPreempted",
//		}),
//		aggregateVcoreSecondsPreempted: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aggregateVcoreSecondsPreempted",
//			Help:      "aggregateVcoreSecondsPreempted",
//		}),
//		activeUsers: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "activeUsers",
//			Help:      "Current number of active users",
//		}),
//		activeApplications: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "activeApplications",
//			Help:      "Current number of active applications",
//		}),
//		appAttemptFirstContainerAllocationDelayNumOps: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "appAttemptFirstContainerAllocationDelayNumOps",
//			Help:      "Total number of first container allocated for all attemptsn",
//		}),
//		appAttemptFirstContainerAllocationDelayAvgTime: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "appAttemptFirstContainerAllocationDelayAvgTime",
//			Help:      "Average time RM spends to allocate the first container for all attempts. For managed AM, the first container is AM container. So, this indicates the time duration to allocate AM container. For unmanaged AM, this is the time duration to allocate the first container asked by unmanaged AM.",
//		}),
//		aggregateMemoryMBPreempted: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aggregateMemoryMBPreempted",
//			Help:      "aggregateMemoryMBPreempted",
//		}),
//		aggregateVcoresPreempted: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aggregateVcoresPreempted",
//			Help:      "aggregateVcoresPreempted",
//		}),
//		allocatedMB: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "allocatedMB",
//			Help:      "Current allocated memory in MB",
//		}),
//		allocatedVCores: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "allocatedVCores",
//			Help:      "Current allocated CPU in virtual cores",
//		}),
//		allocatedContainers: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "allocatedContainers",
//			Help:      "Current number of allocated containers",
//		}),
//		aggregateContainersAllocated: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aAggregateContainersAllocated",
//			Help:      "Total number of allocated containers",
//		}),
//		aggregateContainersReleased: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "aggregateContainersReleased",
//			Help:      "Total number of released containers",
//		}),
//		availableMB: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "availableMB",
//			Help:      "Current available memory in MB",
//		}),
//		availableVCores: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "availableVCores",
//			Help:      "Current available CPU in virtual cores",
//		}),
//		pendingMB: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "pendingMB",
//			Help:      "Current memory requests in MB that are pending to be fulfilled by the scheduler",
//		}),
//		pendingVCores: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "pendingVCores",
//			Help:      "Current CPU requests in virtual cores that are pending to be fulfilled by the scheduler",
//		}),
//		pendingContainers: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "pendingContainers",
//			Help:      "Current number of containers that are pending to be fulfilled by the scheduler",
//		}),
//		reservedMB: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "reservedMB",
//			Help:      "Current reserved memory in MB",
//		}),
//		reservedVCores: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "reservedVCores",
//			Help:      "Current reserved CPU in virtual cores",
//		}),
//		reservedContainers: prometheus.NewGauge(prometheus.GaugeOpts{
//			Namespace: namespace,
//			Subsystem: subNamespace,
//			Name:      "reservedContainers",
//			Help:      "Current number of reserved containers",
//		}),
//	}
//}
//
//func (e *Expoter)Describe(ch chan <- *prometheus.Desc)  {
//	e.name.Describe(ch)
//	e.modelerType.Describe(ch)
//	e.tagQueue.Describe(ch)
//	e.tagContext.Describe(ch)
//	e.tagHostName.Describe(ch)
//	e.amResourceLimitMB.Describe(ch)
//	e.amResourceLimitVCores.Describe(ch)
//	e.usedAMResourceMB.Describe(ch)
//	e.usedAMResourceVCores.Describe(ch)
//	e.usedCapacity.Describe(ch)
//	e.absoluteUsedCapacity.Describe(ch)
//	e.guaranteedMB.Describe(ch)
//	e.guaranteedVCores.Describe(ch)
//	e.maxCapacityMB.Describe(ch)
//	e.maxCapacityVCores.Describe(ch)
//	e.guaranteedCapacity.Describe(ch)
//	e.guaranteedAbsoluteCapacity.Describe(ch)
//	e.maxCapacity.Describe(ch)
//	e.maxAbsoluteCapacity.Describe(ch)
//	e.appsSubmitted.Describe(ch)
//	e.appsRunning.Describe(ch)
//	e.appsPending.Describe(ch)
//	e.appsCompleted.Describe(ch)
//	e.appsKilled.Describe(ch)
//	e.appsFailed.Describe(ch)
//	e.aggregateNodeLocalContainersAllocated.Describe(ch)
//	e.aggregateRackLocalContainersAllocated.Describe(ch)
//	e.aggregateOffSwitchContainersAllocated.Describe(ch)
//	e.aggregateContainersPreempted.Describe(ch)
//	e.aggregateMemoryMBSecondsPreempted.Describe(ch)
//	e.aggregateVcoreSecondsPreempted.Describe(ch)
//	e.activeUsers.Describe(ch)
//	e.activeApplications.Describe(ch)
//	e.appAttemptFirstContainerAllocationDelayNumOps.Describe(ch)
//	e.appAttemptFirstContainerAllocationDelayAvgTime.Describe(ch)
//	e.aggregateMemoryMBPreempted.Describe(ch)
//	e.aggregateVcoresPreempted.Describe(ch)
//	e.allocatedMB.Describe(ch)
//	e.allocatedVCores.Describe(ch)
//	e.allocatedContainers.Describe(ch)
//	e.aggregateContainersAllocated.Describe(ch)
//	e.aggregateContainersPreempted.Describe(ch)
//	e.availableMB.Describe(ch)
//	e.availableVCores.Describe(ch)
//	e.pendingMB.Describe(ch)
//	e.pendingVCores.Describe(ch)
//	e.pendingContainers.Describe(ch)
//	e.reservedMB.Describe(ch)
//	e.reservedVCores.Describe(ch)
//	e.reservedContainers.Describe(ch)
//
//}
//
//func (e *Expoter)Collect(ch chan<- prometheus.Metric)  {
//	resp, err := http.Get(e.url + "/jmx?qry=Hadoop:*")
//	if err != nil {
//		glog.Errorln(err)
//	}
//	defer resp.Body.Close()
//	data, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		glog.Errorln(err)
//	}
//
//	var rm ResourceManager
//	err = json.Unmarshal(data, &rm)
//	if err != nil {
//		glog.Errorln(err)
//	}
//
//	for _, l := range rm.Beans{
//		if l.ModelerType == "QueueMetrics"{
//			ch <- prometheus.MustNewConstMetric(
//				queue, prometheus.GaugeValue, float64(lag), group.GroupId, topic, strconv.FormatInt(int64(partition), 10),
//			)
//		}
//	}
//	e.name.Set(cm["name"].(float64))
//	e.name.Collect(ch)
//}
//
//
//func toFlag(name string, help string) *kingpin.FlagClause {
//	flag.CommandLine.String(name, "", help) // hack around flag.Parse and glog.init flags
//	return kingpin.Flag(name, help)
//}
//
//var (
//	listenAddress      = flag.String("web.listen-address", ":9088", "Address on which to expose metrics and web interface.")
//	metricsPath        = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
//	resourceManagerUrl = flag.String("resourcemanager.url", "http://localhost:8088", "Hadoop ResourceManager URL.")
//)
//
//func main() {
//	flag.Parse()
//	exporter := NewExporter(*resourceManagerUrl)
//	prometheus.MustRegister(exporter)
//
//	http.Handle(*metricsPath, promhttp.Handler())
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte(`<html>
//	        <head><title>Kafka Exporter</title></head>
//	        <body>
//	        <h1>Kafka Exporter</h1>
//	        <p><a href='` + *metricsPath + `'>Metrics</a></p>
//	        </body>
//	        </html>`))
//	})
//	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
//		// need more specific sarama check
//		w.Write([]byte("ok"))
//	})
//
//	glog.Infoln("Listening on", *listenAddress)
//	glog.Fatal(http.ListenAndServe(*listenAddress, nil))
//}
