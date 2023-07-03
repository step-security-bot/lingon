// Copyright (c) 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package promstack

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var KubePromtheusStackGrafanaSA = &corev1.ServiceAccount{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "grafana",
			"app.kubernetes.io/version":    "9.5.3",
			"helm.sh/chart":                "grafana-6.57.3",
		},
		Name:      "kube-promtheus-stack-grafana",
		Namespace: "monitoring",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ServiceAccount",
	},
}

var KubePromtheusStackGrafanaTestSA = &corev1.ServiceAccount{
	ObjectMeta: metav1.ObjectMeta{
		Annotations: map[string]string{
			"helm.sh/hook":               "test-success",
			"helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded",
		},
		Labels: map[string]string{
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "grafana",
			"app.kubernetes.io/version":    "9.5.3",
			"helm.sh/chart":                "grafana-6.57.3",
		},
		Name:      "kube-promtheus-stack-grafana-test",
		Namespace: "monitoring",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ServiceAccount",
	},
}

var KubePromtheusStackKubeAdmissionSA = &corev1.ServiceAccount{
	ObjectMeta: metav1.ObjectMeta{
		Annotations: map[string]string{
			"helm.sh/hook":               "pre-install,pre-upgrade,post-install,post-upgrade",
			"helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded",
		},
		Labels: map[string]string{
			"app":                          "kube-prometheus-stack-admission",
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/part-of":    "kube-prometheus-stack",
			"app.kubernetes.io/version":    "47.0.0",
			"chart":                        "kube-prometheus-stack-47.0.0",
			"heritage":                     "Helm",
			"release":                      "kube-promtheus-stack",
		},
		Name:      "kube-promtheus-stack-kube-admission",
		Namespace: "monitoring",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ServiceAccount",
	},
}

var KubePromtheusStackKubeAlertmanagerSA = &corev1.ServiceAccount{
	AutomountServiceAccountToken: P(true),
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app":                          "kube-prometheus-stack-alertmanager",
			"app.kubernetes.io/component":  "alertmanager",
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "kube-prometheus-stack-alertmanager",
			"app.kubernetes.io/part-of":    "kube-prometheus-stack",
			"app.kubernetes.io/version":    "47.0.0",
			"chart":                        "kube-prometheus-stack-47.0.0",
			"heritage":                     "Helm",
			"release":                      "kube-promtheus-stack",
		},
		Name:      "kube-promtheus-stack-kube-alertmanager",
		Namespace: "monitoring",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ServiceAccount",
	},
}

var KubePromtheusStackKubeOperatorSA = &corev1.ServiceAccount{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app":                          "kube-prometheus-stack-operator",
			"app.kubernetes.io/component":  "prometheus-operator",
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "kube-prometheus-stack-prometheus-operator",
			"app.kubernetes.io/part-of":    "kube-prometheus-stack",
			"app.kubernetes.io/version":    "47.0.0",
			"chart":                        "kube-prometheus-stack-47.0.0",
			"heritage":                     "Helm",
			"release":                      "kube-promtheus-stack",
		},
		Name:      "kube-promtheus-stack-kube-operator",
		Namespace: "monitoring",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ServiceAccount",
	},
}

var KubePromtheusStackKubePrometheusSA = &corev1.ServiceAccount{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app":                          "kube-prometheus-stack-prometheus",
			"app.kubernetes.io/component":  "prometheus",
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "kube-prometheus-stack-prometheus",
			"app.kubernetes.io/part-of":    "kube-prometheus-stack",
			"app.kubernetes.io/version":    "47.0.0",
			"chart":                        "kube-prometheus-stack-47.0.0",
			"heritage":                     "Helm",
			"release":                      "kube-promtheus-stack",
		},
		Name:      "kube-promtheus-stack-kube-prometheus",
		Namespace: "monitoring",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ServiceAccount",
	},
}

var KubePromtheusStackKubeStateMetricsSA = &corev1.ServiceAccount{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component":  "metrics",
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "kube-state-metrics",
			"app.kubernetes.io/part-of":    "kube-state-metrics",
			"app.kubernetes.io/version":    "2.9.2",
			"helm.sh/chart":                "kube-state-metrics-5.7.0",
			"release":                      "kube-promtheus-stack",
		},
		Name:      "kube-promtheus-stack-kube-state-metrics",
		Namespace: "monitoring",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ServiceAccount",
	},
}

var KubePromtheusStackPrometheusNodeExporterSA = &corev1.ServiceAccount{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component":  "metrics",
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "prometheus-node-exporter",
			"app.kubernetes.io/part-of":    "prometheus-node-exporter",
			"app.kubernetes.io/version":    "1.5.0",
			"helm.sh/chart":                "prometheus-node-exporter-4.17.5",
			"jobLabel":                     "node-exporter",
			"release":                      "kube-promtheus-stack",
		},
		Name:      "kube-promtheus-stack-prometheus-node-exporter",
		Namespace: "monitoring",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ServiceAccount",
	},
}