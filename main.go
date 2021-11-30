package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethanliuu/hadoop-exporter/hdfs"
	"github.com/ethanliuu/hadoop-exporter/yarn"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const(
	rsm = "resourcemanager"
	nodem = "nodemanager"
	datan = "datanode"
	namen = "namenode"
)
type Exporter struct {
	URI                    map[string]string
	yarnMetrics , hdfsMetrics map[string]*prometheus.Desc
}


func newYarnMetric(subNamespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName("hadoop_yarn", subNamespace, metricName),
		docString, labels, nil,
	)
}

func newHdfsMetric(subNamespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName("hadoop_hdfs", subNamespace, metricName),
		docString, labels, nil,
	)
}


func NewExporter(uri map[string]string) *Exporter {
	return &Exporter{
		URI: uri,
		yarnMetrics: map[string]*prometheus.Desc{
			//RpcActivityForPort
			"ReceivedBytes":              newYarnMetric("rpc", "received_bytes", "Total number of received bytes", []string{"name", "type", "port", "hostname", "context"}),
			"SentBytes":                  newYarnMetric("rpc", "sent_bytes", "Total number of sent bytes", []string{"name", "type", "port", "hostname", "context"}),
			"RpcQueueTimeNumOps":         newYarnMetric("", "rpc_queue_time_num_ops", "Total number of Rpc calls", []string{"name", "type", "port", "hostname", "context"}),
			"RpcQueueTimeAvgTime":        newYarnMetric("", "rpc_queue_time_avg_time", "Average time waiting for lock acquisition in milliseconds", []string{"name", "type", "port", "hostname", "context"}),
			"RpcProcessingTimeNumOps":    newYarnMetric("", "rpc_processing_time_num_ops", "Total number of Rpc calls (same to RpcQueueTimeNumOps)", []string{"name", "type", "port", "hostname", "context"}),
			"RpcProcessingTimeAvgTime":   newYarnMetric("", "rpc_processing_time_avg_time", "Average time waiting for Rpc calls (same to RpcQueueTimeAvgTime)", []string{"name", "type", "port", "hostname", "context"}),
			"RpcAuthenticationFailures":  newYarnMetric("", "rpc_authentication_failures", "Total number of authentication failures", []string{"name", "type", "port", "hostname", "context"}),
			"RpcAuthenticationSuccesses": newYarnMetric("", "rpc_authentication_successes", "Total number of authentication successes", []string{"name", "type", "port", "hostname", "context"}),
			"RpcAuthorizationFailures":   newYarnMetric("", "rpc_authorization_failures", "Total number of authorization failures", []string{"name", "type", "port", "hostname", "context"}),
			"RpcAuthorizationSuccesses":  newYarnMetric("", "rpc_authorization_successes", "Total number of authorization successes", []string{"name", "type", "port", "hostname", "context"}),
			"NumOpenConnections":         newYarnMetric("rpc", "num_open_connections", "Current number of open connections", []string{"name", "type", "port", "hostname", "context"}),
			"CallQueueLength":            newYarnMetric("rpc", "call_queue_length", "Current length of the call queue", []string{"name", "type", "port", "hostname", "context"}),
			//jvm
			"MemNonHeapUsedM":      newYarnMetric("jvm", "mem_non_heap_used", "Current non-heap memory used in MB", []string{"name", "type", "context", "hostname"}),
			"MemNonHeapCommittedM": newYarnMetric("jvm", "mem_non_heap_committed", "Current non-heap memory committed in MB", []string{"name", "type", "context", "hostname"}),
			"MemNonHeapMaxM":       newYarnMetric("jvm", "mem_non_heap_max", "Max non-heap memory size in MB", []string{"name", "type", "context", "hostname"}),
			"MemHeapUsedM":         newYarnMetric("jvm", "mem_heap_used", "Current heap memory used in MB", []string{"name", "type", "context", "hostname"}),
			"MemHeapCommittedM":    newYarnMetric("jvm", "mem_heap_committed", "Current heap memory committed in MB", []string{"name", "type", "context", "hostname"}),
			"MemHeapMaxM":          newYarnMetric("jvm", "mem_heap_max", "Max heap memory size in MB", []string{"name", "type", "context", "hostname"}),
			"MemMaxM":              newYarnMetric("jvm", "mem_max", "Max memory size in MB", []string{"name", "type", "context", "hostname"}),
			"GcCount":              newYarnMetric("jvm", "gc_count", "Total GC count", []string{"name", "type", "context", "hostname"}),
			"GcTimeMillis":         newYarnMetric("jvm", "gc_time_millis", "Total GC time in msec", []string{"name", "type", "context", "hostname"}),
			"ThreadsNew":           newYarnMetric("jvm", "threads_new", "Current number of NEW threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsRunnable":      newYarnMetric("jvm", "threads_runnable", "Current number of RUNNABLE threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsBlocked":       newYarnMetric("jvm", "threads_blocked", "Current number of BLOCKED threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsWaiting":       newYarnMetric("jvm", "threads_waiting", "Current number of WAITING threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsTimedWaiting":  newYarnMetric("jvm", "threads_timed_waiting", "Current number of TIMED_WAITING threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsTerminated":    newYarnMetric("jvm", "threads_terminated", "Current number of TERMINATED threads", []string{"name", "type", "context", "hostname"}),
			//queueMetrics
			"Running0":    newYarnMetric("queue", "running_0", "Current number of running applications whose elapsed time are less than 60 minutes", []string{"name", "type", "queue", "context", "hostname"}),
			"Running60":   newYarnMetric("queue", "running_60", "Current number of running applications whose elapsed time are between 60 and 300 minutes", []string{"name", "type", "queue", "context", "hostname"}),
			"Running300":  newYarnMetric("queue", "running_300", "Current number of running applications whose elapsed time are between 300 and 1440 minutes", []string{"name", "type", "queue", "context", "hostname"}),
			"Running1440": newYarnMetric("queue", "running_1440", "Current number of running applications elapsed time are more than 1440 minutes", []string{"name", "type", "queue", "context", "hostname"}),
			//			"FairShareMB":                  newYarnMetric("fair_share_mb", "(FairScheduler only) Current fair share of memory in MB", []string{"name", "type", "queue", "context", "hostname"}),
			//			"FairShareVCores":              newYarnMetric("fair_share_virtual_cores", "(FairScheduler only) Current fair share of CPU in virtual cores", []string{"name", "type", "queue", "context", "hostname"}),
			//			"SteadyFairShareMB":            newYarnMetric("steady_fair_share_mb", "SteadyFairShareMB", []string{"name", "type", "queue", "context", "hostname"}),
			//			"SteadyFairShareVCores":        newYarnMetric("steady_fair_share_virtual_cores", "SteadyFairShareVCores", []string{"name", "type", "queue", "context", "hostname"}),
			//			"MinShareMB":                   newYarnMetric("min_share_mb", "(FairScheduler only) Minimum share of memory in MB", []string{"name", "type", "queue", "context", "hostname"}),
			//			"MinShareVCores":               newYarnMetric("min_share_virtual_cores", "(FairScheduler only) Minimum share of CPU in virtual cores", []string{"name", "type", "queue", "context", "hostname"}),
			//			"MaxShareMB":                   newYarnMetric("max_share_virtual_cores", "(FairScheduler only) Maximum share of memory in MB", []string{"name", "type", "queue", "context", "hostname"}),
			//			"MaxShareVCores":               newYarnMetric("max_share_virtual_cores", "(FairScheduler only) Maximum share of CPU in virtual cores", []string{"name", "type", "queue", "context", "hostname"}),
			"AppsSubmitted":                newYarnMetric("queue", "apps_submitted", "Total number of submitted applications", []string{"name", "type", "queue", "context", "hostname"}),
			"AppsRunning":                  newYarnMetric("queue", "apps_running", "Current number of running applications", []string{"name", "type", "queue", "context", "hostname"}),
			"AppsPending":                  newYarnMetric("queue", "apps_pending", "Current number of applications that have not yet been assigned by any containers", []string{"name", "type", "queue", "context", "hostname"}),
			"AppsCompleted":                newYarnMetric("queue", "apps_completed", "Total number of completed applications", []string{"name", "type", "queue", "context", "hostname"}),
			"AppsKilled":                   newYarnMetric("queue", "apps_killed", "Total number of killed applications", []string{"name", "type", "queue", "context", "hostname"}),
			"AppsFailed":                   newYarnMetric("queue", "apps_failed", "Total number of failed applications", []string{"name", "type", "queue", "context", "hostname"}),
			"AllocatedMB":                  newYarnMetric("queue", "allocated_mb", "Current allocated memory in MB", []string{"name", "type", "queue", "context", "hostname"}),
			"AllocatedVCores":              newYarnMetric("queue", "allocated_virtual_cores", "Current allocated CPU in virtual cores", []string{"name", "type", "queue", "context", "hostname"}),
			"AllocatedContainers":          newYarnMetric("queue", "allocated_containers", "Current number of allocated containers", []string{"name", "type", "queue", "context", "hostname"}),
			"AggregateContainersAllocated": newYarnMetric("queue", "aggregate_containers_allocated", "Total number of allocated containers", []string{"name", "type", "queue", "context", "hostname"}),
			"AggregateContainersReleased":  newYarnMetric("queue", "aggregate_containers_released", "Total number of released containers", []string{"name", "type", "queue", "context", "hostname"}),
			"AvailableMB":                  newYarnMetric("queue", "available_mb", "Current available memory in MB", []string{"name", "type", "queue", "context", "hostname"}),
			"AvailableVCores":              newYarnMetric("queue", "available_virtual_cores", "Current available CPU in virtual cores", []string{"name", "type", "queue", "context", "hostname"}),
			"PendingMB":                    newYarnMetric("queue", "pending_mb", "Current memory requests in MB that are pending to be fulfilled by the scheduler", []string{"name", "type", "queue", "context", "hostname"}),
			"PendingVCores":                newYarnMetric("queue", "pending_virtual_cores", "Current CPU requests in virtual cores that are pending to be fulfilled by the scheduler", []string{"name", "type", "queue", "context", "hostname"}),
			"PendingContainers":            newYarnMetric("queue", "pending_containers", "Current number of containers that are pending to be fulfilled by the scheduler", []string{"name", "type", "queue", "context", "hostname"}),
			"ReservedMB":                   newYarnMetric("queue", "reserved_mb", "Current reserved memory in MB", []string{"name", "type", "queue", "context", "hostname"}),
			"ReservedVCores":               newYarnMetric("queue", "reserved_virtual_cores", "Current reserved CPU in virtual cores", []string{"name", "type", "queue", "context", "hostname"}),
			"ReservedContainers":           newYarnMetric("queue", "reserved_containers", "Current number of reserved containers", []string{"name", "type", "queue", "context", "hostname"}),
			"ActiveUsers":                  newYarnMetric("queue", "active_users", "Current number of active users", []string{"name", "type", "queue", "context", "hostname"}),
			"ActiveApplications":           newYarnMetric("queue", "active_applications", "Current number of active applications", []string{"name", "type", "queue", "context", "hostname"}),
			//ClusterMetrics
			"NumActiveNMs":         newYarnMetric("cluster", "num_active_node_managers", "Current number of active NodeManagers", []string{"name", "type", "context", "hostname"}),
			"NumDecommissionedNMs": newYarnMetric("cluster", "num_decommissioned_node_managers", "Current number of NodeManagers being decommissioned", []string{"name", "type", "context", "hostname"}),
			"NumLostNMs":           newYarnMetric("cluster", "num_lost_node_managers", "Current number of lost NodeManagers for not sending heartbeats.", []string{"name", "type", "context", "hostname"}),
			"NumUnhealthyNMs":      newYarnMetric("cluster", "num_unhealthy_node_managers", "Current number of unhealthy NodeManagers", []string{"name", "type", "context", "hostname"}),
			"NumRebootedNMs":       newYarnMetric("cluster", "num_rebooted_node_managers", "Current number of rebooted NodeManagers", []string{"name", "type", "context", "hostname"}),
			//system
			"NumActiveSources":       newYarnMetric("system", "num_active_sources", "Current number of active metrics sources", []string{"name", "type", "context", "hostname"}),
			"NumAllSources":       newYarnMetric("system", "num_all_sources", "Total number of metrics sources", []string{"name", "type", "context", "hostname"}),
			"NumActiveSinks":       newYarnMetric("system", "num_active_sinks", "Current number of active sinks", []string{"name", "type", "context", "hostname"}),
			"NumAllSinks":       newYarnMetric("system", "num_all_sinks", "Total number of sinks", []string{"name", "type", "context", "hostname"}),
			"SnapshotNumOps":       newYarnMetric("system", "snapshot_num_ops", "Total number of operations to snapshot statistics from a metrics source", []string{"name", "type", "context", "hostname"}),
			"SnapshotAvgTime":       newYarnMetric("system", "snapshot_avg_time", "Average time in milliseconds to snapshot statistics from a metrics source", []string{"name", "type", "context", "hostname"}),
			"PublishNumOps":       newYarnMetric("system", "publish_num_ops", "Total number of operations to publish statistics to a sink", []string{"name", "type", "context", "hostname"}),
			"PublishAvgTime":       newYarnMetric("system", "publish_avg_time", "Average time in milliseconds to publish statistics to a sink", []string{"name", "type", "context", "hostname"}),
			"DroppedPubAll":       newYarnMetric("system", "dropped_pub_all", "Total number of dropped publishes", []string{"name", "type", "context", "hostname"}),
		},
		hdfsMetrics: map[string]*prometheus.Desc{
			//dfs
			"BytesWritten":              newHdfsMetric("dfs", "bytes_written", "Total number of bytes written to DataNode", []string{"name", "type", "hostname", "context"}),
			"BytesRead":              newHdfsMetric("dfs", "bytes_read", "Total number of bytes read from DataNode", []string{"name", "type", "hostname", "context"}),
			"BlocksWritten":              newHdfsMetric("dfs", "blocks_written", "Total number of blocks written to DataNode", []string{"name", "type", "hostname", "context"}),
			"BlocksRead":              newHdfsMetric("dfs", "blocks_read", "Total number of blocks read from DataNode", []string{"name", "type", "hostname", "context"}),
			"BlocksReplicated":              newHdfsMetric("dfs", "blocks_replicated", "", []string{"name", "type", "hostname", "context"}),
			"BlocksRemoved":              newHdfsMetric("dfs", "blocks_removed", "", []string{"name", "type", "hostname", "context"}),
			"BlocksVerified":              newHdfsMetric("dfs", "blocks_verified", "", []string{"name", "type", "hostname", "context"}),
			"BlockVerificationFailures":              newHdfsMetric("dfs", "block_verification_failures", "", []string{"name", "type", "hostname", "context"}),
			"BlocksCached":              newHdfsMetric("dfs", "blocks_cached", "", []string{"name", "type", "hostname", "context"}),
			"BlocksUncached":              newHdfsMetric("dfs", "blocks_uncached", "", []string{"name", "type", "hostname", "context"}),
			"ReadsFromLocalClient":              newHdfsMetric("dfs", "reads_from_local_client", "", []string{"name", "type", "hostname", "context"}),
			"ReadsFromRemoteClient":              newHdfsMetric("dfs", "reads_from_remote_client", "", []string{"name", "type", "hostname", "context"}),
			"WritesFromLocalClient":              newHdfsMetric("dfs", "writes_from_local_client", "", []string{"name", "type", "hostname", "context"}),
			"WritesFromRemoteClient":              newHdfsMetric("dfs", "writes_from_remote_client", "", []string{"name", "type", "hostname", "context"}),
			"BlocksGetLocalPathInfo":              newHdfsMetric("dfs", "bocks_get_local_path_info", "", []string{"name", "type", "hostname", "context"}),
			"FsyncCount":              newHdfsMetric("dfs", "fsync_count", "", []string{"name", "type", "hostname", "context"}),
			"VolumeFailures":              newHdfsMetric("dfs", "volume_failures", "", []string{"name", "type", "hostname", "context"}),
			"ReadBlockOpNumOps":              newHdfsMetric("dfs", "read_block_op_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"ReadBlockOpAvgTime":              newHdfsMetric("dfs", "read_block_op_Avg_time", "", []string{"name", "type", "hostname", "context"}),
			"WriteBlockOpNumOps":              newHdfsMetric("dfs", "write_block_op_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"WriteBlockOpAvgTime":              newHdfsMetric("dfs", "write_block_op_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"BlockChecksumOpNumOps":              newHdfsMetric("dfs", "block_checksum_op_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"BlockChecksumOpAvgTime":              newHdfsMetric("dfs", "block_checksum_op_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"CopyBlockOpNumOps":              newHdfsMetric("dfs", "copy_block_op_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"CopyBlockOpAvgTime":              newHdfsMetric("dfs", "copy_block_op_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"ReplaceBlockOpNumOps":              newHdfsMetric("dfs", "replace_block_op_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"ReplaceBlockOpAvgTime":              newHdfsMetric("dfs", "replace_block_op_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"HeartbeatsNumOps":              newHdfsMetric("dfs", "heartbeats_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"HeartbeatsAvgTime":              newHdfsMetric("dfs", "heartbeats_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"BlockReportsNumOps":              newHdfsMetric("dfs", "block_reports_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"BlockReportsAvgTime":              newHdfsMetric("dfs", "block_reports_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"CacheReportsNumOps":              newHdfsMetric("dfs", "cache_reports_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"CacheReportsAvgTime":              newHdfsMetric("dfs", "cache_reports_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"PacketAckRoundTripTimeNanosNumOps":              newHdfsMetric("dfs", "packet_ack_round_trip_time_nanos_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"PacketAckRoundTripTimeNanosAvgTime":              newHdfsMetric("dfs", "packet_ack_round_trip_time_nanos_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"FlushNanosNumOps":              newHdfsMetric("dfs", "flush_nanos_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"FlushNanosAvgTime":              newHdfsMetric("dfs", "flush_nanos_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"FsyncNanosNumOps":              newHdfsMetric("dfs", "fsync_nanos_numOps", "", []string{"name", "type", "hostname", "context"}),
			"FsyncNanosAvgTime":              newHdfsMetric("dfs", "fsync_nanos_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"SendDataPacketBlockedOnNetworkNanosNumOps":              newHdfsMetric("dfs", "send_data_packet_blocked_on_network_nanos_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"SendDataPacketBlockedOnNetworkNanosAvgTime":              newHdfsMetric("dfs", "send_data_packet_blocked_on_network_nanos_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"SendDataPacketTransferNanosNumOps":              newHdfsMetric("dfs", "send_data_packet_transfer_nanos_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"SendDataPacketTransferNanosAvgTime":              newHdfsMetric("dfs", "send_data_packet_transfer_nanos_avg_time", "", []string{"name", "type", "hostname", "context"}),
			// nameNode dfs
			"CreateFileOps":              newHdfsMetric("dfs", "create_file_ops", "", []string{"name", "type", "hostname", "context"}),
			"FilesCreated":              newHdfsMetric("dfs", "files_created", "", []string{"name", "type", "hostname", "context"}),
			"FilesAppended":              newHdfsMetric("dfs", "files_appended", "", []string{"name", "type", "hostname", "context"}),
			"GetBlockLocations":              newHdfsMetric("dfs", "get_block_locations", "", []string{"name", "type", "hostname", "context"}),
			"GetListingOps":              newHdfsMetric("dfs", "get_listing_ops", "", []string{"name", "type", "hostname", "context"}),
			"DeleteFileOps":              newHdfsMetric("dfs", "delete_file_pps", "", []string{"name", "type", "hostname", "context"}),
			"FilesDeleted":              newHdfsMetric("dfs", "files_deleted", "", []string{"name", "type", "hostname", "context"}),
			"FileInfoOps":              newHdfsMetric("dfs", "file_info_ops", "", []string{"name", "type", "hostname", "context"}),
			"AddBlockOps":              newHdfsMetric("dfs", "add_block_ops", "", []string{"name", "type", "hostname", "context"}),
			"GetAdditionalDatanodeOps":              newHdfsMetric("dfs", "get_additional_datanode_ops", "", []string{"name", "type", "hostname", "context"}),
			"CreateSymlinkOps":              newHdfsMetric("dfs", "create_symlink_ops", "", []string{"name", "type", "hostname", "context"}),
			"GetLinkTargetOps":              newHdfsMetric("dfs", "get_link_target_ops", "", []string{"name", "type", "hostname", "context"}),
			"FilesInGetListingOps":              newHdfsMetric("dfs", "files_in_get_listing_ops", "", []string{"name", "type", "hostname", "context"}),
			"AllowSnapshotOps":              newHdfsMetric("dfs", "allow_snapshot_ops", "", []string{"name", "type", "hostname", "context"}),
			"DisallowSnapshotOps":              newHdfsMetric("dfs", "disallow_snapshot_ops", "", []string{"name", "type", "hostname", "context"}),
			"CreateSnapshotOps":              newHdfsMetric("dfs", "create_snapshot_ops", "", []string{"name", "type", "hostname", "context"}),
			"DeleteSnapshotOps":              newHdfsMetric("dfs", "delete_snapshot_ops", "", []string{"name", "type", "hostname", "context"}),
			"RenameSnapshotOps":              newHdfsMetric("dfs", "rename_snapshot_ops", "", []string{"name", "type", "hostname", "context"}),
			"ListSnapshottableDirOps":              newHdfsMetric("dfs", "list_snapshottable_dir_ops", "", []string{"name", "type", "hostname", "context"}),
			"SnapshotDiffReportOps":              newHdfsMetric("dfs", "snapshot_diff_report_ops", "", []string{"name", "type", "hostname", "context"}),
			"BlockReceivedAndDeletedOps":              newHdfsMetric("dfs", "block_received_and_deleted_ops", "", []string{"name", "type", "hostname", "context"}),
			"StorageBlockReportOps":              newHdfsMetric("dfs", "storage_block_report_ops", "", []string{"name", "type", "hostname", "context"}),
			"TransactionsNumOps":              newHdfsMetric("dfs", "transactions_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"TransactionsAvgTime":              newHdfsMetric("dfs", "transactions_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"SyncsNumOps":              newHdfsMetric("dfs", "syncs_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"SyncsAvgTime":              newHdfsMetric("dfs", "syncs_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"TransactionsBatchedInSync":              newHdfsMetric("dfs", "transactions_batched_in_sync", "", []string{"name", "type", "hostname", "context"}),
			"BlockReportNumOps":              newHdfsMetric("dfs", "block_report_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"BlockReportAvgTime":              newHdfsMetric("dfs", "block_report_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"CacheReportNumOps":              newHdfsMetric("dfs", "cache_report_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"CacheReportAvgTime":              newHdfsMetric("dfs", "cache_report_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"SafeModeTime":              newHdfsMetric("dfs", "safe_mode_time", "", []string{"name", "type", "hostname", "context"}),
			"FsImageLoadTime":              newHdfsMetric("dfs", "fs_image_load_time", "", []string{"name", "type", "hostname", "context"}),
			"GetEditNumOps":              newHdfsMetric("dfs", "get_edit_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"GetEditAvgTime":              newHdfsMetric("dfs", "get_edit_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"GetImageNumOps":              newHdfsMetric("dfs", "get_image_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"GetImageAvgTime":              newHdfsMetric("dfs", "get_image_avg_time", "", []string{"name", "type", "hostname", "context"}),
			"PutImageNumOps":              newHdfsMetric("dfs", "put_image_num_ops", "", []string{"name", "type", "hostname", "context"}),
			"PutImageAvgTime":              newHdfsMetric("dfs", "put_image_avg_time", "", []string{"name", "type", "hostname", "context"}),


			//jvm
			"MemNonHeapUsedM":      newHdfsMetric("jvm", "mem_non_heap_used", "Current non-heap memory used in MB", []string{"name", "type", "context", "hostname"}),
			"MemNonHeapCommittedM": newHdfsMetric("jvm", "mem_non_heap_committed", "Current non-heap memory committed in MB", []string{"name", "type", "context", "hostname"}),
			"MemNonHeapMaxM":       newHdfsMetric("jvm", "mem_non_heap_max", "Max non-heap memory size in MB", []string{"name", "type", "context", "hostname"}),
			"MemHeapUsedM":         newHdfsMetric("jvm", "mem_heap_used", "Current heap memory used in MB", []string{"name", "type", "context", "hostname"}),
			"MemHeapCommittedM":    newHdfsMetric("jvm", "mem_heap_committed", "Current heap memory committed in MB", []string{"name", "type", "context", "hostname"}),
			"MemHeapMaxM":          newHdfsMetric("jvm", "mem_heap_max", "Max heap memory size in MB", []string{"name", "type", "context", "hostname"}),
			"MemMaxM":              newHdfsMetric("jvm", "mem_max", "Max memory size in MB", []string{"name", "type", "context", "hostname"}),
			"GcCount":              newHdfsMetric("jvm", "gc_count", "Total GC count", []string{"name", "type", "context", "hostname"}),
			"GcTimeMillis":         newHdfsMetric("jvm", "gc_time_millis", "Total GC time in msec", []string{"name", "type", "context", "hostname"}),
			"ThreadsNew":           newHdfsMetric("jvm", "threads_new", "Current number of NEW threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsRunnable":      newHdfsMetric("jvm", "threads_runnable", "Current number of RUNNABLE threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsBlocked":       newHdfsMetric("jvm", "threads_blocked", "Current number of BLOCKED threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsWaiting":       newHdfsMetric("jvm", "threads_waiting", "Current number of WAITING threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsTimedWaiting":  newHdfsMetric("jvm", "threads_timed_waiting", "Current number of TIMED_WAITING threads", []string{"name", "type", "context", "hostname"}),
			"ThreadsTerminated":    newHdfsMetric("jvm", "threads_terminated", "Current number of TERMINATED threads", []string{"name", "type", "context", "hostname"}),
			//RpcActivityForPort
			"ReceivedBytes":              newHdfsMetric("rpc", "received_bytes", "Total number of received bytes", []string{"name", "type", "port", "hostname", "context"}),
			"SentBytes":                  newHdfsMetric("rpc", "sent_bytes", "Total number of sent bytes", []string{"name", "type", "port", "hostname", "context"}),
			"RpcQueueTimeNumOps":         newHdfsMetric("", "rpc_queue_time_num_ops", "Total number of Rpc calls", []string{"name", "type", "port", "hostname", "context"}),
			"RpcQueueTimeAvgTime":        newHdfsMetric("", "rpc_queue_time_avg_time", "Average time waiting for lock acquisition in milliseconds", []string{"name", "type", "port", "hostname", "context"}),
			"RpcProcessingTimeNumOps":    newHdfsMetric("", "rpc_processing_time_num_ops", "Total number of Rpc calls (same to RpcQueueTimeNumOps)", []string{"name", "type", "port", "hostname", "context"}),
			"RpcProcessingTimeAvgTime":   newHdfsMetric("", "rpc_processing_time_avg_time", "Average time waiting for Rpc calls (same to RpcQueueTimeAvgTime)", []string{"name", "type", "port", "hostname", "context"}),
			"RpcAuthenticationFailures":  newHdfsMetric("", "rpc_authentication_failures", "Total number of authentication failures", []string{"name", "type", "port", "hostname", "context"}),
			"RpcAuthenticationSuccesses": newHdfsMetric("", "rpc_authentication_successes", "Total number of authentication successes", []string{"name", "type", "port", "hostname", "context"}),
			"RpcAuthorizationFailures":   newHdfsMetric("", "rpc_authorization_failures", "Total number of authorization failures", []string{"name", "type", "port", "hostname", "context"}),
			"RpcAuthorizationSuccesses":  newHdfsMetric("", "rpc_authorization_successes", "Total number of authorization successes", []string{"name", "type", "port", "hostname", "context"}),
			"NumOpenConnections":         newHdfsMetric("rpc", "num_open_connections", "Current number of open connections", []string{"name", "type", "port", "hostname", "context"}),
			"CallQueueLength":            newHdfsMetric("rpc", "call_queue_length", "Current length of the call queue", []string{"name", "type", "port", "hostname", "context"}),
			//system
			"NumActiveSources":       newHdfsMetric("system", "num_active_sources", "Current number of active metrics sources", []string{"name", "type", "context", "hostname"}),
			"NumAllSources":       newHdfsMetric("system", "num_all_sources", "Total number of metrics sources", []string{"name", "type", "context", "hostname"}),
			"NumActiveSinks":       newHdfsMetric("system", "num_active_sinks", "Current number of active sinks", []string{"name", "type", "context", "hostname"}),
			"NumAllSinks":       newHdfsMetric("system", "num_all_sinks", "Total number of sinks", []string{"name", "type", "context", "hostname"}),
			"SnapshotNumOps":       newHdfsMetric("system", "snapshot_num_ops", "Total number of operations to snapshot statistics from a metrics source", []string{"name", "type", "context", "hostname"}),
			"SnapshotAvgTime":       newHdfsMetric("system", "snapshot_avg_time", "Average time in milliseconds to snapshot statistics from a metrics source", []string{"name", "type", "context", "hostname"}),
			"PublishNumOps":       newHdfsMetric("system", "publish_num_ops", "Total number of operations to publish statistics to a sink", []string{"name", "type", "context", "hostname"}),
			"PublishAvgTime":       newHdfsMetric("system", "publish_avg_time", "Average time in milliseconds to publish statistics to a sink", []string{"name", "type", "context", "hostname"}),
			"DroppedPubAll":       newHdfsMetric("system", "dropped_pub_all", "Total number of dropped publishes", []string{"name", "type", "context", "hostname"}),
			// FSNamesystem
			"MissingBlocks":       newHdfsMetric("dfs", "missing_blocks", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"ExpiredHeartbeats":       newHdfsMetric("dfs", "expired_heartbeats", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"TransactionsSinceLastCheckpoint":       newHdfsMetric("dfs", "transactions_since_last_checkpoint", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"TransactionsSinceLastLogRoll":       newHdfsMetric("dfs", "transactions_since_lastLog_roll", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"LastWrittenTransactionId":       newHdfsMetric("dfs", "last_written_transaction_id", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"LastCheckpointTime":       newHdfsMetric("dfs", "last_checkpoint_time", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"CapacityTotal":       newHdfsMetric("dfs", "capacity_total", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"CapacityTotalGB":       newHdfsMetric("dfs", "capacity_total_gb", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"CapacityUsed":       newHdfsMetric("dfs", "capacity_used", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"CapacityUsedGB":       newHdfsMetric("dfs", "capacity_used_gb", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"CapacityRemaining":       newHdfsMetric("dfs", "capacity_remaining", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"CapacityRemainingGB":       newHdfsMetric("dfs", "capacity_remaining_gb", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"CapacityUsedNonDFS":       newHdfsMetric("dfs", "capacity_used_non_dfs", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"TotalLoad":       newHdfsMetric("dfs", "total_load", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"BlocksTotal":       newHdfsMetric("dfs", "blocks_total", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"FilesTotal":       newHdfsMetric("dfs", "files_total", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"PendingReplicationBlocks":       newHdfsMetric("dfs", "pending_replication_blocks", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"UnderReplicatedBlocks":       newHdfsMetric("dfs", "under_replicated_blocks", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"CorruptBlocks":       newHdfsMetric("dfs", "corrupt_blocks", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"ScheduledReplicationBlocks":       newHdfsMetric("dfs", "scheduled_replication_blocks", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"PendingDeletionBlocks":       newHdfsMetric("dfs", "pending_deletion_blocks", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"ExcessBlocks":       newHdfsMetric("dfs", "excess_blocks", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"PostponedMisreplicatedBlocks":       newHdfsMetric("dfs", "postponed_misreplicated_blocks", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"PendingDataNodeMessageCount":       newHdfsMetric("dfs", "pending_data_node_message_count", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"MillisSinceLastLoadedEdits":       newHdfsMetric("dfs", "millis_since_last_loaded_edits", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"BlockCapacity":       newHdfsMetric("dfs", "block_capacity", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"StaleDataNodes":       newHdfsMetric("dfs", "stale_data_nodes", "", []string{"name", "type", "context","ha_state", "hostname"}),
			"TotalFiles":       newHdfsMetric("dfs", "total_files", "", []string{"name", "type", "context","ha_state", "hostname"}),
		},
	}
}

var (
	listenAddress       = flag.String("web.listen_address", ":9088", "Address on which to expose metrics and web interface.")
	metricsPath         = flag.String("web.telemetry_path", "/metrics", "Path under which to expose metrics.")
	resourceManagerURI  = flag.String("hadoop.resource_manager_uri", "", "Hadoop ResourceManager URI.")
	nameNodeURI         = flag.String("hadoop.name_node_uri", "", "Hadoop NameNode URI.")
	dataNodeURI         = flag.String("hadoop.data_node_uri", "", "Hadoop DataNode URI.")
	nodeManagerURI      = flag.String("hadoop.node_manager_uri", "", "Hadoop NodeManager URI.")
	hadoopScrapeTimeout = flag.Int("hadoop.scrape_timeout", 2, "The number of seconds to wait for an HTTP response from the hadoop")
	insecure            = flag.Bool("insecure", false, "Ignore server certificate if using https")
)

func fetchHTTP(uri string, timeout time.Duration) func() (io.ReadCloser, error) {
	http.DefaultClient.Timeout = timeout
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: *insecure}

	return func() (io.ReadCloser, error) {
		resp, err := http.DefaultClient.Get(uri)
		if err != nil {
			return nil, err
		}
		if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
			resp.Body.Close()
			return nil, fmt.Errorf("HTTP status %d", resp.StatusCode)
		}
		return resp.Body, nil
	}
}
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range e.yarnMetrics {
		ch <- m
	}
}


func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	var resourceManager yarn.RsmNodem
	var nodeManager yarn.RsmNodem

	var nameNode hdfs.NamenDatan
	var dataNode hdfs.NamenDatan
	yarnMetrics := func(r yarn.RsmNodem,name string) {
		for _, r := range r.Beans {
			if i := strings.Index(r.ModelerType, "RpcActivityForPort"); i != -1 {
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ReceivedBytes"], prometheus.GaugeValue, float64(r.ReceivedBytes), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["SentBytes"], prometheus.GaugeValue, float64(r.SentBytes), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["RpcQueueTimeNumOps"], prometheus.GaugeValue, float64(r.RpcQueueTimeNumOps), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["RpcQueueTimeAvgTime"], prometheus.GaugeValue, r.RpcQueueTimeAvgTime, name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["RpcProcessingTimeNumOps"], prometheus.GaugeValue, float64(r.RpcProcessingTimeNumOps), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["RpcProcessingTimeAvgTime"], prometheus.GaugeValue, r.RpcProcessingTimeAvgTime, name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["RpcAuthenticationFailures"], prometheus.GaugeValue, float64(r.RpcAuthenticationFailures), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["RpcAuthenticationSuccesses"], prometheus.GaugeValue, float64(r.RpcAuthenticationSuccesses), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["RpcAuthorizationFailures"], prometheus.GaugeValue, float64(r.RpcAuthorizationFailures), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["RpcAuthorizationSuccesses"], prometheus.GaugeValue, float64(r.RpcAuthorizationSuccesses), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumOpenConnections"], prometheus.GaugeValue, float64(r.NumOpenConnections), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["CallQueueLength"], prometheus.GaugeValue, float64(r.CallQueueLength), name,r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
			}
			if i := strings.Index(r.ModelerType, "JvmMetrics"); i != -1 {
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["MemNonHeapUsedM"], prometheus.GaugeValue, r.MemNonHeapUsedM, name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["MemNonHeapCommittedM"], prometheus.GaugeValue, r.MemNonHeapCommittedM, name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["MemNonHeapMaxM"], prometheus.GaugeValue, r.MemNonHeapMaxM, name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["MemHeapUsedM"], prometheus.GaugeValue, r.MemHeapUsedM, name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["MemHeapCommittedM"], prometheus.GaugeValue, r.MemHeapCommittedM, name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["MemHeapMaxM"], prometheus.GaugeValue, r.MemHeapMaxM, name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["MemMaxM"], prometheus.GaugeValue, r.MemMaxM, name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["GcCount"], prometheus.GaugeValue, float64(r.GcCount), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["GcTimeMillis"], prometheus.GaugeValue, float64(r.GcTimeMillis), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ThreadsNew"], prometheus.GaugeValue, float64(r.ThreadsNew), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ThreadsRunnable"], prometheus.GaugeValue, float64(r.ThreadsRunnable), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ThreadsBlocked"], prometheus.GaugeValue, float64(r.ThreadsBlocked), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ThreadsWaiting"], prometheus.GaugeValue, float64(r.ThreadsWaiting), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ThreadsTimedWaiting"], prometheus.GaugeValue, float64(r.ThreadsTimedWaiting), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ThreadsTerminated"], prometheus.GaugeValue, float64(r.ThreadsTerminated), name,r.ModelerType, r.TagContext, r.TagHostname)
			}
			if i := strings.Index(r.ModelerType, "QueueMetrics"); i != -1 {
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["Running0"], prometheus.GaugeValue, float64(r.Running0), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["Running60"], prometheus.GaugeValue, float64(r.Running60), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["Running300"], prometheus.GaugeValue, float64(r.Running300), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["Running1440"], prometheus.GaugeValue, float64(r.Running1440), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AppsSubmitted"], prometheus.GaugeValue, float64(r.AppsSubmitted), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AppsRunning"], prometheus.GaugeValue, float64(r.AppsRunning), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AppsPending"], prometheus.GaugeValue, float64(r.AppsPending), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AppsCompleted"], prometheus.GaugeValue, float64(r.AppsCompleted), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AppsKilled"], prometheus.GaugeValue, float64(r.AppsKilled), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AppsFailed"], prometheus.GaugeValue, float64(r.AppsFailed), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AllocatedMB"], prometheus.GaugeValue, float64(r.AllocatedMB), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AllocatedVCores"], prometheus.GaugeValue, float64(r.AllocatedVCores), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AllocatedContainers"], prometheus.GaugeValue, float64(r.AllocatedContainers), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AggregateContainersAllocated"], prometheus.GaugeValue, float64(r.AggregateContainersAllocated), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AggregateContainersReleased"], prometheus.GaugeValue, float64(r.AggregateContainersReleased), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AvailableMB"], prometheus.GaugeValue, float64(r.AvailableMB), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["AvailableVCores"], prometheus.GaugeValue, float64(r.AvailableVCores), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["PendingMB"], prometheus.GaugeValue, float64(r.PendingMB), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["PendingVCores"], prometheus.GaugeValue, float64(r.PendingVCores), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["PendingContainers"], prometheus.GaugeValue, float64(r.PendingContainers), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ReservedMB"], prometheus.GaugeValue, float64(r.ReservedMB), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ReservedVCores"], prometheus.GaugeValue, float64(r.ReservedVCores), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ReservedContainers"], prometheus.GaugeValue, float64(r.ReservedContainers), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ActiveUsers"], prometheus.GaugeValue, float64(r.ActiveUsers), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["ActiveApplications"], prometheus.GaugeValue, float64(r.ActiveApplications), name,r.ModelerType, r.TagQueue, r.TagContext, r.TagHostname)
			}
			if i:= strings.Index(r.ModelerType,"MetricsSystem,sub=Stats"); i != -1{
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumActiveSources"], prometheus.GaugeValue, float64(r.NumActiveSources), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumAllSources"], prometheus.GaugeValue, float64(r.NumAllSources), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumActiveSinks"], prometheus.GaugeValue, float64(r.NumActiveSinks), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumAllSinks"], prometheus.GaugeValue, float64(r.NumAllSinks), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["SnapshotNumOps"], prometheus.GaugeValue, float64(r.SnapshotNumOps), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["SnapshotAvgTime"], prometheus.GaugeValue, r.SnapshotAvgTime, name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["PublishNumOps"], prometheus.GaugeValue, float64(r.PublishNumOps), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["PublishAvgTime"], prometheus.GaugeValue, r.PublishAvgTime, name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["DroppedPubAll"], prometheus.GaugeValue, float64(r.DroppedPubAll), name,r.ModelerType, r.TagContext, r.TagHostname)

			}

			if i := strings.Index(r.ModelerType, "ClusterMetrics"); i != -1 {
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumActiveNMs"], prometheus.GaugeValue, float64(r.NumActiveNMs), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumDecommissionedNMs"], prometheus.GaugeValue, float64(r.NumDecommissionedNMs), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumLostNMs"], prometheus.GaugeValue, float64(r.NumLostNMs), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumUnhealthyNMs"], prometheus.GaugeValue, float64(r.NumUnhealthyNMs), name,r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.yarnMetrics["NumRebootedNMs"], prometheus.GaugeValue, float64(r.NumRebootedNMs), name,r.ModelerType, r.TagContext, r.TagHostname)

			}
		}
	}

	hdfsMetrics := func(r hdfs.NamenDatan,name string) {
		for _,r := range r.Beans {
			if i := strings.Index(r.ModelerType, "JvmMetrics"); i != -1 {
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["MemNonHeapUsedM"], prometheus.GaugeValue, r.MemNonHeapUsedM, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["MemNonHeapCommittedM"], prometheus.GaugeValue, r.MemNonHeapCommittedM, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["MemNonHeapMaxM"], prometheus.GaugeValue, r.MemNonHeapMaxM, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["MemHeapUsedM"], prometheus.GaugeValue, r.MemHeapUsedM, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["MemHeapCommittedM"], prometheus.GaugeValue, r.MemHeapCommittedM, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["MemHeapMaxM"], prometheus.GaugeValue, r.MemHeapMaxM, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["MemMaxM"], prometheus.GaugeValue, r.MemMaxM, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GcCount"], prometheus.GaugeValue, float64(r.GcCount), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GcTimeMillis"], prometheus.GaugeValue, float64(r.GcTimeMillis), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ThreadsNew"], prometheus.GaugeValue, float64(r.ThreadsNew), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ThreadsRunnable"], prometheus.GaugeValue, float64(r.ThreadsRunnable), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ThreadsBlocked"], prometheus.GaugeValue, float64(r.ThreadsBlocked), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ThreadsWaiting"], prometheus.GaugeValue, float64(r.ThreadsWaiting), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ThreadsTimedWaiting"], prometheus.GaugeValue, float64(r.ThreadsTimedWaiting), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ThreadsTerminated"], prometheus.GaugeValue, float64(r.ThreadsTerminated), name, r.ModelerType, r.TagContext, r.TagHostname)
			}

			if i := strings.Index(r.ModelerType, "DataNodeActivity"); i != -1 {
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BytesWritten"], prometheus.GaugeValue, float64(r.BytesWritten), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BytesRead"], prometheus.GaugeValue, float64(r.BytesRead), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlocksWritten"], prometheus.GaugeValue, float64(r.BlocksWritten), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlocksRead"], prometheus.GaugeValue, float64(r.BlocksRead), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlocksReplicated"], prometheus.GaugeValue, float64(r.BlocksReplicated), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlocksRemoved"], prometheus.GaugeValue, float64(r.BlocksRemoved), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlocksVerified"], prometheus.GaugeValue, float64(r.BlocksVerified), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlockVerificationFailures"], prometheus.GaugeValue, float64(r.BlockVerificationFailures), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlocksCached"], prometheus.GaugeValue, float64(r.BlocksCached), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlocksUncached"], prometheus.GaugeValue, float64(r.BlocksUncached), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ReadsFromLocalClient"], prometheus.GaugeValue, float64(r.ReadsFromLocalClient), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ReadsFromRemoteClient"], prometheus.GaugeValue, float64(r.ReadsFromRemoteClient), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["WritesFromLocalClient"], prometheus.GaugeValue, float64(r.WritesFromLocalClient), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["WritesFromRemoteClient"], prometheus.GaugeValue, float64(r.WritesFromRemoteClient), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlocksGetLocalPathInfo"], prometheus.GaugeValue, float64(r.BlocksGetLocalPathInfo), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FsyncCount"], prometheus.GaugeValue, float64(r.FsyncCount), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ReadBlockOpNumOps"], prometheus.GaugeValue, float64(r.ReadBlockOpNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ReadBlockOpAvgTime"], prometheus.GaugeValue, r.ReadBlockOpAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["WriteBlockOpNumOps"], prometheus.GaugeValue, float64(r.WriteBlockOpNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["WriteBlockOpAvgTime"], prometheus.GaugeValue, r.WriteBlockOpAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlockChecksumOpNumOps"], prometheus.GaugeValue, float64(r.BlockChecksumOpNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlockChecksumOpAvgTime"], prometheus.GaugeValue, r.BlockChecksumOpAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CopyBlockOpNumOps"], prometheus.GaugeValue, float64(r.CopyBlockOpNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CopyBlockOpAvgTime"], prometheus.GaugeValue, r.CopyBlockOpAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ReplaceBlockOpNumOps"], prometheus.GaugeValue, float64(r.ReplaceBlockOpNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ReplaceBlockOpAvgTime"], prometheus.GaugeValue, r.ReplaceBlockOpAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["HeartbeatsNumOps"], prometheus.GaugeValue, float64(r.HeartbeatsNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["HeartbeatsAvgTime"], prometheus.GaugeValue, r.HeartbeatsAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlockReportsNumOps"], prometheus.GaugeValue, float64(r.BlockReportsNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlockReportsAvgTime"], prometheus.GaugeValue, r.BlockReportsAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CacheReportsNumOps"], prometheus.GaugeValue, float64(r.CacheReportsNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CacheReportsAvgTime"], prometheus.GaugeValue, r.CacheReportsAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PacketAckRoundTripTimeNanosNumOps"], prometheus.GaugeValue, float64(r.PacketAckRoundTripTimeNanosNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PacketAckRoundTripTimeNanosAvgTime"], prometheus.GaugeValue, r.PacketAckRoundTripTimeNanosAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FlushNanosNumOps"], prometheus.GaugeValue, float64(r.FlushNanosNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FlushNanosAvgTime"], prometheus.GaugeValue, r.FlushNanosAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FsyncNanosNumOps"], prometheus.GaugeValue, float64(r.FsyncNanosNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FsyncNanosAvgTime"], prometheus.GaugeValue, r.FsyncNanosAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SendDataPacketBlockedOnNetworkNanosNumOps"], prometheus.GaugeValue, float64(r.SendDataPacketBlockedOnNetworkNanosNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SendDataPacketBlockedOnNetworkNanosAvgTime"], prometheus.GaugeValue, r.SendDataPacketBlockedOnNetworkNanosAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SendDataPacketTransferNanosNumOps"], prometheus.GaugeValue, float64(r.SendDataPacketTransferNanosNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SendDataPacketTransferNanosAvgTime"], prometheus.GaugeValue, r.SendDataPacketTransferNanosAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
			}
			if i := strings.Index(r.ModelerType, "NameNodeActivity"); i != -1 {
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CreateFileOps"], prometheus.GaugeValue, float64(r.CreateFileOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FilesCreated"], prometheus.GaugeValue, float64(r.FilesCreated), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FilesAppended"], prometheus.GaugeValue, float64(r.FilesAppended), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GetBlockLocations"], prometheus.GaugeValue, float64(r.GetBlockLocations), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GetListingOps"], prometheus.GaugeValue, float64(r.GetListingOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["DeleteFileOps"], prometheus.GaugeValue, float64(r.DeleteFileOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FilesDeleted"], prometheus.GaugeValue, float64(r.FilesDeleted), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FileInfoOps"], prometheus.GaugeValue, float64(r.FileInfoOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["AddBlockOps"], prometheus.GaugeValue, float64(r.AddBlockOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GetAdditionalDatanodeOps"], prometheus.GaugeValue, float64(r.GetAdditionalDatanodeOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CreateSymlinkOps"], prometheus.GaugeValue, float64(r.CreateSymlinkOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GetLinkTargetOps"], prometheus.GaugeValue, float64(r.GetLinkTargetOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FilesInGetListingOps"], prometheus.GaugeValue, float64(r.FilesInGetListingOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["AllowSnapshotOps"], prometheus.GaugeValue, float64(r.AllowSnapshotOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["DisallowSnapshotOps"], prometheus.GaugeValue, float64(r.DisallowSnapshotOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CreateSnapshotOps"], prometheus.GaugeValue, float64(r.CreateSnapshotOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["DeleteSnapshotOps"], prometheus.GaugeValue, float64(r.DeleteSnapshotOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["RenameSnapshotOps"], prometheus.GaugeValue, float64(r.RenameSnapshotOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ListSnapshottableDirOps"], prometheus.GaugeValue, float64(r.ListSnapshottableDirOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SnapshotDiffReportOps"], prometheus.GaugeValue, float64(r.SnapshotDiffReportOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlockReceivedAndDeletedOps"], prometheus.GaugeValue, float64(r.BlockReceivedAndDeletedOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["StorageBlockReportOps"], prometheus.GaugeValue, float64(r.StorageBlockReportOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["TransactionsNumOps"], prometheus.GaugeValue, float64(r.TransactionsNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["TransactionsAvgTime"], prometheus.GaugeValue, r.TransactionsAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SyncsNumOps"], prometheus.GaugeValue, float64(r.SyncsNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SyncsAvgTime"], prometheus.GaugeValue, r.SyncsAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["TransactionsBatchedInSync"], prometheus.GaugeValue, float64(r.TransactionsBatchedInSync), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlockReportNumOps"], prometheus.GaugeValue, float64(r.BlockReportNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlockReportAvgTime"], prometheus.GaugeValue, r.BlockReportAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CacheReportNumOps"], prometheus.GaugeValue, float64(r.CacheReportNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CacheReportAvgTime"], prometheus.GaugeValue, r.CacheReportAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SafeModeTime"], prometheus.GaugeValue, float64(r.SafeModeTime), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FsImageLoadTime"], prometheus.GaugeValue, float64(r.FsImageLoadTime), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GetEditNumOps"], prometheus.GaugeValue, float64(r.GetEditNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GetEditAvgTime"], prometheus.GaugeValue, r.GetEditAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GetImageNumOps"], prometheus.GaugeValue, float64(r.GetImageNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["GetImageAvgTime"], prometheus.GaugeValue, r.GetImageAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PutImageNumOps"], prometheus.GaugeValue, float64(r.PutImageNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PutImageAvgTime"], prometheus.GaugeValue, r.PutImageAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)

			}
			if i := strings.Index(r.ModelerType, "RpcActivityForPort"); i != -1 {
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ReceivedBytes"], prometheus.GaugeValue, float64(r.ReceivedBytes), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SentBytes"], prometheus.GaugeValue, float64(r.SentBytes), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["RpcQueueTimeNumOps"], prometheus.GaugeValue, float64(r.RpcProcessingTimeNumOps), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["RpcQueueTimeAvgTime"], prometheus.GaugeValue, r.RpcQueueTimeAvgTime, name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["RpcProcessingTimeNumOps"], prometheus.GaugeValue, float64(r.RpcProcessingTimeNumOps), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["RpcProcessingTimeAvgTime"], prometheus.GaugeValue, r.RpcProcessingTimeAvgTime, name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["RpcAuthenticationFailures"], prometheus.GaugeValue, float64(r.RpcAuthenticationFailures), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["RpcAuthenticationSuccesses"], prometheus.GaugeValue, float64(r.RpcAuthenticationSuccesses), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["RpcAuthorizationFailures"], prometheus.GaugeValue, float64(r.RpcAuthorizationFailures), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["RpcAuthorizationSuccesses"], prometheus.GaugeValue, float64(r.RpcAuthorizationSuccesses), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["NumOpenConnections"], prometheus.GaugeValue, float64(r.NumOpenConnections), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CallQueueLength"], prometheus.GaugeValue, float64(r.CallQueueLength), name, r.ModelerType, r.TagPort, r.TagHostname, r.TagContext)
			}
			if i := strings.Index(r.ModelerType, "MetricsSystem,sub=Stats"); i != -1 {
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["NumActiveSources"], prometheus.GaugeValue, float64(r.NumActiveSources), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["NumAllSources"], prometheus.GaugeValue, float64(r.NumAllSources), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["NumActiveSinks"], prometheus.GaugeValue, float64(r.NumActiveSinks), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["NumAllSinks"], prometheus.GaugeValue, float64(r.NumAllSinks), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SnapshotNumOps"], prometheus.GaugeValue, float64(r.SnapshotNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["SnapshotAvgTime"], prometheus.GaugeValue, r.SnapshotAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PublishNumOps"], prometheus.GaugeValue, float64(r.PublishNumOps), name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PublishAvgTime"], prometheus.GaugeValue, r.PublishAvgTime, name, r.ModelerType, r.TagContext, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["DroppedPubAll"], prometheus.GaugeValue, float64(r.DroppedPubAll), name, r.ModelerType, r.TagContext, r.TagHostname)
			}
			if i := strings.Index(r.ModelerType, "FSNamesystem"); i != -1 && i == 0{
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["MissingBlocks"], prometheus.GaugeValue, float64(r.MissingBlocks), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ExpiredHeartbeats"], prometheus.GaugeValue, float64(r.ExpiredHeartbeats), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["TransactionsSinceLastCheckpoint"], prometheus.GaugeValue, float64(r.TransactionsSinceLastCheckpoint), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["TransactionsSinceLastLogRoll"], prometheus.GaugeValue, float64(r.TransactionsSinceLastLogRoll), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["LastWrittenTransactionId"], prometheus.GaugeValue, float64(r.LastWrittenTransactionID), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["LastCheckpointTime"], prometheus.GaugeValue, float64(r.LastCheckpointTime), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CapacityTotal"], prometheus.GaugeValue, float64(r.CapacityTotal), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CapacityTotalGB"], prometheus.GaugeValue, r.CapacityTotalGB, name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CapacityUsed"], prometheus.GaugeValue, float64(r.CapacityUsed), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CapacityUsedGB"], prometheus.GaugeValue, r.CapacityUsedGB, name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CapacityRemaining"], prometheus.GaugeValue, float64(r.CapacityRemaining), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CapacityRemainingGB"], prometheus.GaugeValue, r.CapacityRemainingGB, name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CapacityUsedNonDFS"], prometheus.GaugeValue, float64(r.CapacityUsedNonDFS), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["TotalLoad"], prometheus.GaugeValue, float64(r.TotalLoad), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlocksTotal"], prometheus.GaugeValue, float64(r.BlocksTotal), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["FilesTotal"], prometheus.GaugeValue, float64(r.FilesTotal), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PendingReplicationBlocks"], prometheus.GaugeValue, float64(r.PendingReplicationBlocks), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["UnderReplicatedBlocks"], prometheus.GaugeValue, float64(r.UnderReplicatedBlocks), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["CorruptBlocks"], prometheus.GaugeValue, float64(r.CorruptBlocks), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ScheduledReplicationBlocks"], prometheus.GaugeValue, float64(r.ScheduledReplicationBlocks), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PendingDeletionBlocks"], prometheus.GaugeValue, float64(r.PendingDeletionBlocks), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["ExcessBlocks"], prometheus.GaugeValue, float64(r.ExcessBlocks), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PostponedMisreplicatedBlocks"], prometheus.GaugeValue, float64(r.PostponedMisreplicatedBlocks), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["PendingDataNodeMessageCount"], prometheus.GaugeValue, float64(r.PendingDataNodeMessageCount), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["MillisSinceLastLoadedEdits"], prometheus.GaugeValue, float64(r.MillisSinceLastLoadedEdits), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["BlockCapacity"], prometheus.GaugeValue, float64(r.BlockCapacity), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["StaleDataNodes"], prometheus.GaugeValue, float64(r.StaleDataNodes), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
				ch <- prometheus.MustNewConstMetric(e.hdfsMetrics["TotalFiles"], prometheus.GaugeValue, float64(r.TotalFiles), name, r.ModelerType, r.TagContext,r.TagHAState, r.TagHostname)
			}
		}
	}

	log.Println("range e.URI")
	for name, uri := range e.URI {
		body, err := fetchHTTP(uri+"/jmx?qry=Hadoop:*", time.Duration(*hadoopScrapeTimeout)*time.Second)()
		if err != nil {
			log.Println("fetchHTTP failed", err)
			return
		}

		data, err := ioutil.ReadAll(body)
		if err != nil {
			log.Println("ioutil.ReadAll failed", err)
			return
		}
		body.Close()

		if name == "resourceManager" {
			log.Println("reso")
			err = json.Unmarshal(data, &resourceManager)
			if err != nil {
				log.Println("json.Unmarshal failed", err)
				return
			}
			yarnMetrics(resourceManager,rsm)
		}else if name =="nodeManager"{
			err = json.Unmarshal(data, &nodeManager)
			if err != nil {
				log.Println("json.Unmarshal failed", err)
				return
			}
			yarnMetrics(nodeManager,nodem)
		}else if name =="dataNode" {
			log.Println("dataNode")
			err = json.Unmarshal(data, &dataNode)
			if err != nil {
				log.Println("json.Unmarshal failed", err)
				return
			}
			hdfsMetrics(dataNode,datan)

		}else if name == "nameNode" {
			log.Println("nameNode")
			err = json.Unmarshal(data, &nameNode)
			if err != nil {
				log.Println("json.Unmarshal failed", err)
				return
			}
			hdfsMetrics(nameNode,namen)
		}

	}
	log.Println("range resourceManager.Beans")


	log.Println("finish")

}
func init() {
	prometheus.MustRegister(version.NewCollector("hadoop_exporter"))
}
func main() {
	flag.Parse()
	log.Println("start111")

	log.Printf("Starting nginx_vts_exporter %s", version.Info())
	log.Printf("Build context %s", version.BuildContext())

	uri := make(map[string]string)
	if *nameNodeURI != "" {
		uri["nameNode"] = *nameNodeURI
	}
	if *dataNodeURI != "" {
		uri["dataNode"] = *dataNodeURI
	}
	if *nodeManagerURI != "" {
		uri["nodeManager"] = *nodeManagerURI
	}
	if *resourceManagerURI != "" {
		uri["resourceManager"] = *resourceManagerURI
	}

	exporter := NewExporter(uri)
	prometheus.MustRegister(exporter)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Hadoop Exporter</title></head>
			<body>
			<h1>Hadoop Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}
