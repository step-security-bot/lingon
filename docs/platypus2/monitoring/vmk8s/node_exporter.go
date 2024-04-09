// Copyright (c) 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package vmk8s

import (
	"github.com/VictoriaMetrics/operator/api/victoriametrics/v1beta1"
	"github.com/golingon/lingon/pkg/kube"
	ku "github.com/golingon/lingon/pkg/kubeutil"
	"github.com/golingon/lingoneks/meta"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	NodeExportPort     = 9100
	NodeExportPortName = "metrics"
	NodeExportVersion  = "1.4.0"
)

var NE = &meta.Metadata{
	Name:      "node-exporter", // linked to the name of the JobLabel in VMServiceScrape
	Namespace: namespace,
	Instance:  "node-exporter-" + namespace,
	Component: "metrics-export",
	PartOf:    appName,
	Version:   NodeExportVersion,
	ManagedBy: "lingon",
	Img: meta.ContainerImg{
		Registry: "quay.io",
		Image:    "prometheus/node-exporter",
		Tag:      "v" + NodeExportVersion,
	},
}

type NodeExporter struct {
	kube.App

	DS         *appsv1.DaemonSet
	SA         *corev1.ServiceAccount
	SVC        *corev1.Service
	Rules      *v1beta1.VMRule
	AlertRules *v1beta1.VMRule
	Scrape     *v1beta1.VMServiceScrape
}

func NewNodeExporter() *NodeExporter {
	return &NodeExporter{
		DS:         NodeExporterDS,
		SA:         NodeExporterSA,
		SVC:        NodeExporterSVC,
		Rules:      NodeExporterRules,
		AlertRules: NodeExporterAlertRules,
		Scrape:     NodeExporterScrape,
	}
}

var NodeExporterSVC = &corev1.Service{
	TypeMeta:   ku.TypeServiceV1,
	ObjectMeta: NE.ObjectMeta(),
	Spec: corev1.ServiceSpec{
		Ports: []corev1.ServicePort{
			{
				Name:       NodeExportPortName,
				Port:       int32(NodeExportPort),
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(NodeExportPort),
			},
		},
		Selector: NE.MatchLabels(),
		Type:     corev1.ServiceTypeClusterIP,
	},
}

var NodeExporterSA = NE.ServiceAccount()

var NodeExporterDS = &appsv1.DaemonSet{
	TypeMeta:   ku.TypeDaemonSetV1,
	ObjectMeta: NE.ObjectMeta(),
	Spec: appsv1.DaemonSetSpec{
		Selector: &metav1.LabelSelector{MatchLabels: NE.MatchLabels()},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Annotations: map[string]string{"cluster-autoscaler.kubernetes.io/safe-to-evict": "true"},
				Labels:      NE.Labels(),
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Args: []string{
							"--path.procfs=/host/proc",
							"--path.sysfs=/host/sys",
							"--path.rootfs=/host/root",
							"--web.listen-address=[$(HOST_IP)]:" + d(NodeExportPort),
							"--collector.filesystem.ignored-mount-points=^/(dev|proc|sys|var/lib/docker/.+|var/lib/kubelet/.+)($|/)",
							"--collector.filesystem.ignored-fs-types=^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs)$",
						},
						Env: []corev1.EnvVar{
							{Name: "HOST_IP", Value: "0.0.0.0"},
						},
						Image:           NE.Img.URL(),
						ImagePullPolicy: corev1.PullIfNotPresent,
						LivenessProbe: &corev1.Probe{
							ProbeHandler: ku.ProbeHTTP("/", NodeExportPort),
						},
						Name: NE.Name,
						Ports: []corev1.ContainerPort{
							{
								ContainerPort: int32(NodeExportPort),
								Name:          NodeExportPortName,
								Protocol:      corev1.ProtocolTCP,
							},
						},
						ReadinessProbe: &corev1.Probe{
							ProbeHandler: ku.ProbeHTTP("/", NodeExportPort),
						},
						Resources: ku.Resources(
							"250m",
							"300Mi",
							"250m",
							"300Mi",
						),
						VolumeMounts: []corev1.VolumeMount{
							{
								MountPath: "/host/proc",
								Name:      "proc",
								ReadOnly:  true,
							}, {
								MountPath: "/host/sys",
								Name:      "sys",
								ReadOnly:  true,
							}, {
								MountPath:        "/host/root",
								MountPropagation: P(corev1.MountPropagationHostToContainer),
								Name:             "root",
								ReadOnly:         true,
							},
						},
					},
				},
				HostNetwork: true,
				HostPID:     true,
				SecurityContext: &corev1.PodSecurityContext{
					FSGroup:      P(int64(65534)),
					RunAsGroup:   P(int64(65534)),
					RunAsNonRoot: P(true),
					RunAsUser:    P(int64(65534)),
				},
				ServiceAccountName: NodeExporterSA.Name,
				// This does not work with Fargate
				// Tolerations: []corev1.Toleration{
				// 	{
				// 		Effect:   corev1.TaintEffectNoSchedule,
				// 		Operator: corev1.TolerationOpExists,
				// 	},
				// },
				Volumes: []corev1.Volume{
					{
						Name:         "proc",
						VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/proc"}},
					}, {
						Name:         "sys",
						VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/sys"}},
					}, {
						Name:         "root",
						VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/"}},
					},
				},
			},
		},
		UpdateStrategy: appsv1.DaemonSetUpdateStrategy{
			RollingUpdate: &appsv1.RollingUpdateDaemonSet{
				MaxUnavailable: P(intstr.FromInt(1)),
			},
			Type: appsv1.RollingUpdateDaemonSetStrategyType,
		},
	},
}

var NodeExporterScrape = &v1beta1.VMServiceScrape{
	TypeMeta:   TypeVMServiceScrapeV1Beta1,
	ObjectMeta: NE.ObjectMeta(),
	Spec: v1beta1.VMServiceScrapeSpec{
		Endpoints: []v1beta1.Endpoint{
			{
				MetricRelabelConfigs: []*v1beta1.RelabelConfig{
					{
						Action:                 "drop",
						Regex:                  "/var/lib/kubelet/pods.+",
						SourceLabels:           []string{"mountpoint"},
						UnderScoreSourceLabels: []string{"mountpoint"},
					},
				},
				Port: NodeExportPortName,
			},
		},
		// JobLabel is linked to [ku.AppLabelName] which is defined by [NE.Name]
		JobLabel: ku.AppLabelName,
		Selector: metav1.LabelSelector{MatchLabels: NE.MatchLabels()},
	},
}

var NodeExporterRules = &v1beta1.VMRule{
	TypeMeta:   TypeVMRuleV1Beta1,
	ObjectMeta: NE.ObjectMetaNameSuffix("rules"),
	Spec: v1beta1.VMRuleSpec{
		Groups: []v1beta1.RuleGroup{
			{
				Name: "node-exporter.rules",
				Rules: []v1beta1.Rule{
					{
						Expr: `
count without (cpu, mode) (
  node_cpu_seconds_total{job="node-exporter",mode="idle"}
)
`,
						Record: "instance:node_num_cpu:sum",
					}, {
						Expr: `
1 - avg without (cpu) (
  sum without (mode) (rate(node_cpu_seconds_total{job="node-exporter", mode=~"idle|iowait|steal"}[5m]))
)
`,
						Record: "instance:node_cpu_utilisation:rate5m",
					}, {
						Expr: `
(
  node_load1{job="node-exporter"}
/
  instance:node_num_cpu:sum{job="node-exporter"}
)
`,
						Record: "instance:node_load1_per_cpu:ratio",
					}, {
						Expr: `
1 - (
  (
    node_memory_MemAvailable_bytes{job="node-exporter"}
    or
    (
      node_memory_Buffers_bytes{job="node-exporter"}
      +
      node_memory_Cached_bytes{job="node-exporter"}
      +
      node_memory_MemFree_bytes{job="node-exporter"}
      +
      node_memory_Slab_bytes{job="node-exporter"}
    )
  )
/
  node_memory_MemTotal_bytes{job="node-exporter"}
)
`,
						Record: "instance:node_memory_utilisation:ratio",
					}, {
						Expr:   `rate(node_vmstat_pgmajfault{job="node-exporter"}[5m])`,
						Record: "instance:node_vmstat_pgmajfault:rate5m",
					}, {
						Expr:   `rate(node_disk_io_time_seconds_total{job="node-exporter", device=~"(/dev/)?(mmcblk.p.+|nvme.+|rbd.+|sd.+|vd.+|xvd.+|dm-.+|md.+|dasd.+)"}[5m])`,
						Record: "instance_device:node_disk_io_time_seconds:rate5m",
					}, {
						Expr:   `rate(node_disk_io_time_weighted_seconds_total{job="node-exporter", device=~"(/dev/)?(mmcblk.p.+|nvme.+|rbd.+|sd.+|vd.+|xvd.+|dm-.+|md.+|dasd.+)"}[5m])`,
						Record: "instance_device:node_disk_io_time_weighted_seconds:rate5m",
					}, {
						Expr: `
sum without (device) (
  rate(node_network_receive_bytes_total{job="node-exporter", device!="lo"}[5m])
)
`,
						Record: "instance:node_network_receive_bytes_excluding_lo:rate5m",
					}, {
						Expr: `
sum without (device) (
  rate(node_network_transmit_bytes_total{job="node-exporter", device!="lo"}[5m])
)
`,
						Record: "instance:node_network_transmit_bytes_excluding_lo:rate5m",
					}, {
						Expr: `
sum without (device) (
  rate(node_network_receive_drop_total{job="node-exporter", device!="lo"}[5m])
)
`,
						Record: "instance:node_network_receive_drop_excluding_lo:rate5m",
					}, {
						Expr: `
sum without (device) (
  rate(node_network_transmit_drop_total{job="node-exporter", device!="lo"}[5m])
)
`,
						Record: "instance:node_network_transmit_drop_excluding_lo:rate5m",
					},
				},
			},
		},
	},
}

var NodeExporterAlertRules = &v1beta1.VMRule{
	TypeMeta:   TypeVMRuleV1Beta1,
	ObjectMeta: NE.ObjectMeta(),
	Spec: v1beta1.VMRuleSpec{
		Groups: []v1beta1.RuleGroup{
			{
				Name: "node-exporter",
				Rules: []v1beta1.Rule{
					{
						Alert: "NodeFilesystemSpaceFillingUp",
						Annotations: map[string]string{
							"description": `Filesystem on {{ $labels.device }} at {{ $labels.instance }} has only {{ printf "%.2f" $value }}% available space left and is filling up.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefilesystemspacefillingup",
							"summary":     "Filesystem is predicted to run out of space within the next 24 hours.",
						},
						Expr: `
(
  node_filesystem_avail_bytes{job="node-exporter",fstype!="",mountpoint!=""} / node_filesystem_size_bytes{job="node-exporter",fstype!="",mountpoint!=""} * 100 < 15
and
  predict_linear(node_filesystem_avail_bytes{job="node-exporter",fstype!="",mountpoint!=""}[6h], 24*60*60) < 0
and
  node_filesystem_readonly{job="node-exporter",fstype!="",mountpoint!=""} == 0
)
`,
						For:    "1h",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeFilesystemSpaceFillingUp",
						Annotations: map[string]string{
							"description": `Filesystem on {{ $labels.device }} at {{ $labels.instance }} has only {{ printf "%.2f" $value }}% available space left and is filling up fast.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefilesystemspacefillingup",
							"summary":     "Filesystem is predicted to run out of space within the next 4 hours.",
						},
						Expr: `
(
  node_filesystem_avail_bytes{job="node-exporter",fstype!="",mountpoint!=""} / node_filesystem_size_bytes{job="node-exporter",fstype!="",mountpoint!=""} * 100 < 10
and
  predict_linear(node_filesystem_avail_bytes{job="node-exporter",fstype!="",mountpoint!=""}[6h], 4*60*60) < 0
and
  node_filesystem_readonly{job="node-exporter",fstype!="",mountpoint!=""} == 0
)
`,
						For:    "1h",
						Labels: map[string]string{"severity": "critical"},
					}, {
						Alert: "NodeFilesystemAlmostOutOfSpace",
						Annotations: map[string]string{
							"description": `Filesystem on {{ $labels.device }} at {{ $labels.instance }} has only {{ printf "%.2f" $value }}% available space left.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefilesystemalmostoutofspace",
							"summary":     "Filesystem has less than 5% space left.",
						},
						Expr: `
(
  node_filesystem_avail_bytes{job="node-exporter",fstype!="",mountpoint!=""} / node_filesystem_size_bytes{job="node-exporter",fstype!="",mountpoint!=""} * 100 < 5
and
  node_filesystem_readonly{job="node-exporter",fstype!="",mountpoint!=""} == 0
)
`,
						For:    "30m",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeFilesystemAlmostOutOfSpace",
						Annotations: map[string]string{
							"description": `Filesystem on {{ $labels.device }} at {{ $labels.instance }} has only {{ printf "%.2f" $value }}% available space left.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefilesystemalmostoutofspace",
							"summary":     "Filesystem has less than 3% space left.",
						},
						Expr: `
(
  node_filesystem_avail_bytes{job="node-exporter",fstype!="",mountpoint!=""} / node_filesystem_size_bytes{job="node-exporter",fstype!="",mountpoint!=""} * 100 < 3
and
  node_filesystem_readonly{job="node-exporter",fstype!="",mountpoint!=""} == 0
)
`,
						For:    "30m",
						Labels: map[string]string{"severity": "critical"},
					}, {
						Alert: "NodeFilesystemFilesFillingUp",
						Annotations: map[string]string{
							"description": `Filesystem on {{ $labels.device }} at {{ $labels.instance }} has only {{ printf "%.2f" $value }}% available inodes left and is filling up.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefilesystemfilesfillingup",
							"summary":     "Filesystem is predicted to run out of inodes within the next 24 hours.",
						},
						Expr: `
(
  node_filesystem_files_free{job="node-exporter",fstype!="",mountpoint!=""} / node_filesystem_files{job="node-exporter",fstype!="",mountpoint!=""} * 100 < 40
and
  predict_linear(node_filesystem_files_free{job="node-exporter",fstype!="",mountpoint!=""}[6h], 24*60*60) < 0
and
  node_filesystem_readonly{job="node-exporter",fstype!="",mountpoint!=""} == 0
)
`,
						For:    "1h",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeFilesystemFilesFillingUp",
						Annotations: map[string]string{
							"description": `Filesystem on {{ $labels.device }} at {{ $labels.instance }} has only {{ printf "%.2f" $value }}% available inodes left and is filling up fast.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefilesystemfilesfillingup",
							"summary":     "Filesystem is predicted to run out of inodes within the next 4 hours.",
						},
						Expr: `
(
  node_filesystem_files_free{job="node-exporter",fstype!="",mountpoint!=""} / node_filesystem_files{job="node-exporter",fstype!="",mountpoint!=""} * 100 < 20
and
  predict_linear(node_filesystem_files_free{job="node-exporter",fstype!="",mountpoint!=""}[6h], 4*60*60) < 0
and
  node_filesystem_readonly{job="node-exporter",fstype!="",mountpoint!=""} == 0
)
`,
						For:    "1h",
						Labels: map[string]string{"severity": "critical"},
					}, {
						Alert: "NodeFilesystemAlmostOutOfFiles",
						Annotations: map[string]string{
							"description": `Filesystem on {{ $labels.device }} at {{ $labels.instance }} has only {{ printf "%.2f" $value }}% available inodes left.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefilesystemalmostoutoffiles",
							"summary":     "Filesystem has less than 5% inodes left.",
						},
						Expr: `
(
  node_filesystem_files_free{job="node-exporter",fstype!="",mountpoint!=""} / node_filesystem_files{job="node-exporter",fstype!="",mountpoint!=""} * 100 < 5
and
  node_filesystem_readonly{job="node-exporter",fstype!="",mountpoint!=""} == 0
)
`,
						For:    "1h",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeFilesystemAlmostOutOfFiles",
						Annotations: map[string]string{
							"description": `Filesystem on {{ $labels.device }} at {{ $labels.instance }} has only {{ printf "%.2f" $value }}% available inodes left.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefilesystemalmostoutoffiles",
							"summary":     "Filesystem has less than 3% inodes left.",
						},
						Expr: `
(
  node_filesystem_files_free{job="node-exporter",fstype!="",mountpoint!=""} / node_filesystem_files{job="node-exporter",fstype!="",mountpoint!=""} * 100 < 3
and
  node_filesystem_readonly{job="node-exporter",fstype!="",mountpoint!=""} == 0
)
`,
						For:    "1h",
						Labels: map[string]string{"severity": "critical"},
					}, {
						Alert: "NodeNetworkReceiveErrs",
						Annotations: map[string]string{
							"description": `{{ $labels.instance }} interface {{ $labels.device }} has encountered {{ printf "%.0f" $value }} receive errors in the last two minutes.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodenetworkreceiveerrs",
							"summary":     "Network interface is reporting many receive errors.",
						},
						Expr:   "rate(node_network_receive_errs_total[2m]) / rate(node_network_receive_packets_total[2m]) > 0.01",
						For:    "1h",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeNetworkTransmitErrs",
						Annotations: map[string]string{
							"description": `{{ $labels.instance }} interface {{ $labels.device }} has encountered {{ printf "%.0f" $value }} transmit errors in the last two minutes.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodenetworktransmiterrs",
							"summary":     "Network interface is reporting many transmit errors.",
						},
						Expr:   "rate(node_network_transmit_errs_total[2m]) / rate(node_network_transmit_packets_total[2m]) > 0.01",
						For:    "1h",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeHighNumberConntrackEntriesUsed",
						Annotations: map[string]string{
							"description": "{{ $value | humanizePercentage }} of conntrack entries are used.",
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodehighnumberconntrackentriesused",
							"summary":     "Number of conntrack are getting close to the limit.",
						},
						Expr:   "(node_nf_conntrack_entries / node_nf_conntrack_entries_limit) > 0.75",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeTextFileCollectorScrapeError",
						Annotations: map[string]string{
							"description": "Node Exporter text file collector failed to scrape.",
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodetextfilecollectorscrapeerror",
							"summary":     "Node Exporter text file collector failed to scrape.",
						},
						Expr:   `node_textfile_scrape_error{job="node-exporter"} == 1`,
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeClockSkewDetected",
						Annotations: map[string]string{
							"description": "Clock on {{ $labels.instance }} is out of sync by more than 300s. Ensure NTP is configured correctly on this host.",
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodeclockskewdetected",
							"summary":     "Clock skew detected.",
						},
						Expr: `
(
  node_timex_offset_seconds{job="node-exporter"} > 0.05
and
  deriv(node_timex_offset_seconds{job="node-exporter"}[5m]) >= 0
)
or
(
  node_timex_offset_seconds{job="node-exporter"} < -0.05
and
  deriv(node_timex_offset_seconds{job="node-exporter"}[5m]) <= 0
)
`,
						For:    "10m",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeClockNotSynchronising",
						Annotations: map[string]string{
							"description": "Clock on {{ $labels.instance }} is not synchronising. Ensure NTP is configured on this host.",
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodeclocknotsynchronising",
							"summary":     "Clock not synchronising.",
						},
						Expr: `
min_over_time(node_timex_sync_status{job="node-exporter"}[5m]) == 0
and
node_timex_maxerror_seconds{job="node-exporter"} >= 16
`,
						For:    "10m",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeRAIDDegraded",
						Annotations: map[string]string{
							"description": "RAID array '{{ $labels.device }}' on {{ $labels.instance }} is in degraded state due to one or more disks failures. Number of spare drives is insufficient to fix issue automatically.",
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/noderaiddegraded",
							"summary":     "RAID Array is degraded",
						},
						Expr:   `node_md_disks_required{job="node-exporter",device=~"(/dev/)?(mmcblk.p.+|nvme.+|rbd.+|sd.+|vd.+|xvd.+|dm-.+|md.+|dasd.+)"} - ignoring (state) (node_md_disks{state="active",job="node-exporter",device=~"(/dev/)?(mmcblk.p.+|nvme.+|rbd.+|sd.+|vd.+|xvd.+|dm-.+|md.+|dasd.+)"}) > 0`,
						For:    "15m",
						Labels: map[string]string{"severity": "critical"},
					}, {
						Alert: "NodeRAIDDiskFailure",
						Annotations: map[string]string{
							"description": "At least one device in RAID array on {{ $labels.instance }} failed. Array '{{ $labels.device }}' needs attention and possibly a disk swap.",
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/noderaiddiskfailure",
							"summary":     "Failed device in RAID array",
						},
						Expr:   `node_md_disks{state="failed",job="node-exporter",device=~"(/dev/)?(mmcblk.p.+|nvme.+|rbd.+|sd.+|vd.+|xvd.+|dm-.+|md.+|dasd.+)"} > 0`,
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeFileDescriptorLimit",
						Annotations: map[string]string{
							"description": `File descriptors limit at {{ $labels.instance }} is currently at {{ printf "%.2f" $value }}%.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefiledescriptorlimit",
							"summary":     "Kernel is predicted to exhaust file descriptors limit soon.",
						},
						Expr: `
(
  node_filefd_allocated{job="node-exporter"} * 100 / node_filefd_maximum{job="node-exporter"} > 70
)
`,
						For:    "15m",
						Labels: map[string]string{"severity": "warning"},
					}, {
						Alert: "NodeFileDescriptorLimit",
						Annotations: map[string]string{
							"description": `File descriptors limit at {{ $labels.instance }} is currently at {{ printf "%.2f" $value }}%.`,
							"runbook_url": "https://runbooks.prometheus-operator.dev/runbooks/node/nodefiledescriptorlimit",
							"summary":     "Kernel is predicted to exhaust file descriptors limit soon.",
						},
						Expr: `
(
  node_filefd_allocated{job="node-exporter"} * 100 / node_filefd_maximum{job="node-exporter"} > 90
)
`,
						For:    "15m",
						Labels: map[string]string{"severity": "critical"},
					},
				},
			},
		},
	},
}
