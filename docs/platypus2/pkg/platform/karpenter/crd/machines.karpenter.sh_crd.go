// Copyright (c) 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package karpentercrd

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MachinesKarpenterShCRD = &apiextensionsv1.CustomResourceDefinition{
	ObjectMeta: metav1.ObjectMeta{
		Annotations: map[string]string{"controller-gen.kubebuilder.io/version": "v0.11.3"},
		Name:        "machines.karpenter.sh",
	},
	Spec: apiextensionsv1.CustomResourceDefinitionSpec{
		Group: "karpenter.sh",
		Names: apiextensionsv1.CustomResourceDefinitionNames{
			Categories: []string{"karpenter"},
			Kind:       "Machine",
			ListKind:   "MachineList",
			Plural:     "machines",
			Singular:   "machine",
		},
		Scope: apiextensionsv1.ResourceScope("Cluster"),
		Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
			{
				AdditionalPrinterColumns: []apiextensionsv1.CustomResourceColumnDefinition{
					{
						JSONPath: ".metadata.labels.node\\.kubernetes\\.io/instance-type",
						Name:     "Type",
						Type:     "string",
					}, {
						JSONPath: ".metadata.labels.topology\\.kubernetes\\.io/zone",
						Name:     "Zone",
						Type:     "string",
					}, {
						JSONPath: ".status.nodeName",
						Name:     "Node",
						Type:     "string",
					}, {
						JSONPath: ".status.conditions[?(@.type==\"Ready\")].status",
						Name:     "Ready",
						Type:     "string",
					}, {
						JSONPath: ".metadata.creationTimestamp",
						Name:     "Age",
						Type:     "date",
					}, {
						JSONPath: ".metadata.labels.karpenter\\.sh/capacity-type",
						Name:     "Capacity",
						Priority: int32(1),
						Type:     "string",
					}, {
						JSONPath: ".metadata.labels.karpenter\\.sh/provisioner-name",
						Name:     "Provisioner",
						Priority: int32(1),
						Type:     "string",
					}, {
						JSONPath: ".spec.machineTemplateRef.name",
						Name:     "Template",
						Priority: int32(1),
						Type:     "string",
					},
				},
				Name: "v1alpha5",
				Schema: &apiextensionsv1.CustomResourceValidation{
					OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
						Description: "Machine is the Schema for the Machines API",
						Properties: map[string]apiextensionsv1.JSONSchemaProps{
							"apiVersion": {
								Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								Type:        "string",
							},
							"kind": {
								Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Type:        "string",
							},
							"metadata": {Type: "object"},
							"spec": {
								Description: "MachineSpec describes the desired state of the Machine",
								Properties: map[string]apiextensionsv1.JSONSchemaProps{
									"kubelet": {
										Description: "Kubelet are options passed to the kubelet when provisioning nodes",
										Properties: map[string]apiextensionsv1.JSONSchemaProps{
											"clusterDNS": {
												Description: "clusterDNS is a list of IP addresses for the cluster DNS server. Note that not all providers may use all addresses.",
												Items:       &apiextensionsv1.JSONSchemaPropsOrArray{Schema: &apiextensionsv1.JSONSchemaProps{Type: "string"}},
												Type:        "array",
											},
											"containerRuntime": {
												Description: "ContainerRuntime is the container runtime to be used with your worker nodes.",
												Type:        "string",
											},
											"cpuCFSQuota": {
												Description: "CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits.",
												Type:        "boolean",
											},
											"evictionHard": {
												AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
													Allows: true,
													Schema: &apiextensionsv1.JSONSchemaProps{Type: "string"},
												},
												Description: "EvictionHard is the map of signal names to quantities that define hard eviction thresholds",
												Type:        "object",
											},
											"evictionMaxPodGracePeriod": {
												Description: "EvictionMaxPodGracePeriod is the maximum allowed grace period (in seconds) to use when terminating pods in response to soft eviction thresholds being met.",
												Format:      "int32",
												Type:        "integer",
											},
											"evictionSoft": {
												AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
													Allows: true,
													Schema: &apiextensionsv1.JSONSchemaProps{Type: "string"},
												},
												Description: "EvictionSoft is the map of signal names to quantities that define soft eviction thresholds",
												Type:        "object",
											},
											"evictionSoftGracePeriod": {
												AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
													Allows: true,
													Schema: &apiextensionsv1.JSONSchemaProps{Type: "string"},
												},
												Description: "EvictionSoftGracePeriod is the map of signal names to quantities that define grace periods for each eviction signal",
												Type:        "object",
											},
											"imageGCHighThresholdPercent": {
												Description: "ImageGCHighThresholdPercent is the percent of disk usage after which image garbage collection is always run. The percent is calculated by dividing this field value by 100, so this field must be between 0 and 100, inclusive. When specified, the value must be greater than ImageGCLowThresholdPercent.",
												Format:      "int32",
												Maximum:     P(100.0),
												Type:        "integer",
											},
											"imageGCLowThresholdPercent": {
												Description: "ImageGCLowThresholdPercent is the percent of disk usage before which image garbage collection is never run. Lowest disk usage to garbage collect to. The percent is calculated by dividing this field value by 100, so the field value must be between 0 and 100, inclusive. When specified, the value must be less than imageGCHighThresholdPercent",
												Format:      "int32",
												Maximum:     P(100.0),
												Type:        "integer",
											},
											"kubeReserved": {
												AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
													Allows: true,
													Schema: &apiextensionsv1.JSONSchemaProps{
														AnyOf: []apiextensionsv1.JSONSchemaProps{
															{Type: "integer"},
															{Type: "string"},
														},
														Pattern:      "^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$",
														XIntOrString: true,
													},
												},
												Description: "KubeReserved contains resources reserved for Kubernetes system components.",
												Type:        "object",
											},
											"maxPods": {
												Description: "MaxPods is an override for the maximum number of pods that can run on a worker node instance.",
												Format:      "int32",
												Type:        "integer",
											},
											"podsPerCore": {
												Description: "PodsPerCore is an override for the number of pods that can run on a worker node instance based on the number of cpu cores. This value cannot exceed MaxPods, so, if MaxPods is a lower value, that value will be used.",
												Format:      "int32",
												Type:        "integer",
											},
											"systemReserved": {
												AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
													Allows: true,
													Schema: &apiextensionsv1.JSONSchemaProps{
														AnyOf: []apiextensionsv1.JSONSchemaProps{
															{Type: "integer"},
															{Type: "string"},
														},
														Pattern:      "^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$",
														XIntOrString: true,
													},
												},
												Description: "SystemReserved contains resources reserved for OS system daemons and kernel memory.",
												Type:        "object",
											},
										},
										Type: "object",
									},
									"machineTemplateRef": {
										Description: "MachineTemplateRef is a reference to an object that defines provider specific configuration",
										Properties: map[string]apiextensionsv1.JSONSchemaProps{
											"apiVersion": {
												Description: "API version of the referent",
												Type:        "string",
											},
											"kind": {
												Description: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds\"",
												Type:        "string",
											},
											"name": {
												Description: "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
												Type:        "string",
											},
										},
										Required: []string{"name"},
										Type:     "object",
									},
									"requirements": {
										Description: "Requirements are layered with Labels and applied to every node.",
										Items: &apiextensionsv1.JSONSchemaPropsOrArray{
											Schema: &apiextensionsv1.JSONSchemaProps{
												Description: "A node selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
												Properties: map[string]apiextensionsv1.JSONSchemaProps{
													"key": {
														Description: "The label key that the selector applies to.",
														Type:        "string",
													},
													"operator": {
														Description: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
														Type:        "string",
													},
													"values": {
														Description: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
														Items:       &apiextensionsv1.JSONSchemaPropsOrArray{Schema: &apiextensionsv1.JSONSchemaProps{Type: "string"}},
														Type:        "array",
													},
												},
												Required: []string{
													"key",
													"operator",
												},
												Type: "object",
											},
										},
										Type: "array",
									},
									"resources": {
										Description: "Resources models the resource requirements for the Machine to launch",
										Properties: map[string]apiextensionsv1.JSONSchemaProps{
											"requests": {
												AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
													Allows: true,
													Schema: &apiextensionsv1.JSONSchemaProps{
														AnyOf: []apiextensionsv1.JSONSchemaProps{
															{Type: "integer"},
															{Type: "string"},
														},
														Pattern:      "^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$",
														XIntOrString: true,
													},
												},
												Description: "Requests describes the minimum required resources for the Machine to launch",
												Type:        "object",
											},
										},
										Type: "object",
									},
									"startupTaints": {
										Description: "StartupTaints are taints that are applied to nodes upon startup which are expected to be removed automatically within a short period of time, typically by a DaemonSet that tolerates the taint. These are commonly used by daemonsets to allow initialization and enforce startup ordering.  StartupTaints are ignored for provisioning purposes in that pods are not required to tolerate a StartupTaint in order to have nodes provisioned for them.",
										Items: &apiextensionsv1.JSONSchemaPropsOrArray{
											Schema: &apiextensionsv1.JSONSchemaProps{
												Description: "The node this Taint is attached to has the \"effect\" on any pod that does not tolerate the Taint.",
												Properties: map[string]apiextensionsv1.JSONSchemaProps{
													"effect": {
														Description: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
														Type:        "string",
													},
													"key": {
														Description: "Required. The taint key to be applied to a node.",
														Type:        "string",
													},
													"timeAdded": {
														Description: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
														Format:      "date-time",
														Type:        "string",
													},
													"value": {
														Description: "The taint value corresponding to the taint key.",
														Type:        "string",
													},
												},
												Required: []string{
													"effect",
													"key",
												},
												Type: "object",
											},
										},
										Type: "array",
									},
									"taints": {
										Description: "Taints will be applied to the machine's node.",
										Items: &apiextensionsv1.JSONSchemaPropsOrArray{
											Schema: &apiextensionsv1.JSONSchemaProps{
												Description: "The node this Taint is attached to has the \"effect\" on any pod that does not tolerate the Taint.",
												Properties: map[string]apiextensionsv1.JSONSchemaProps{
													"effect": {
														Description: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
														Type:        "string",
													},
													"key": {
														Description: "Required. The taint key to be applied to a node.",
														Type:        "string",
													},
													"timeAdded": {
														Description: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
														Format:      "date-time",
														Type:        "string",
													},
													"value": {
														Description: "The taint value corresponding to the taint key.",
														Type:        "string",
													},
												},
												Required: []string{
													"effect",
													"key",
												},
												Type: "object",
											},
										},
										Type: "array",
									},
								},
								Type: "object",
							},
							"status": {
								Description: "MachineStatus defines the observed state of Machine",
								Properties: map[string]apiextensionsv1.JSONSchemaProps{
									"allocatable": {
										AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
											Allows: true,
											Schema: &apiextensionsv1.JSONSchemaProps{
												AnyOf: []apiextensionsv1.JSONSchemaProps{
													{Type: "integer"},
													{Type: "string"},
												},
												Pattern:      "^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$",
												XIntOrString: true,
											},
										},
										Description: "Allocatable is the estimated allocatable capacity of the machine",
										Type:        "object",
									},
									"capacity": {
										AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
											Allows: true,
											Schema: &apiextensionsv1.JSONSchemaProps{
												AnyOf: []apiextensionsv1.JSONSchemaProps{
													{Type: "integer"},
													{Type: "string"},
												},
												Pattern:      "^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$",
												XIntOrString: true,
											},
										},
										Description: "Capacity is the estimated full capacity of the machine",
										Type:        "object",
									},
									"conditions": {
										Description: "Conditions contains signals for health and readiness",
										Items: &apiextensionsv1.JSONSchemaPropsOrArray{
											Schema: &apiextensionsv1.JSONSchemaProps{
												Description: "Condition defines a readiness condition for a Knative resource. See: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties",
												Properties: map[string]apiextensionsv1.JSONSchemaProps{
													"lastTransitionTime": {
														Description: "LastTransitionTime is the last time the condition transitioned from one status to another. We use VolatileTime in place of metav1.Time to exclude this from creating equality.Semantic differences (all other things held constant).",
														Type:        "string",
													},
													"message": {
														Description: "A human readable message indicating details about the transition.",
														Type:        "string",
													},
													"reason": {
														Description: "The reason for the condition's last transition.",
														Type:        "string",
													},
													"severity": {
														Description: "Severity with which to treat failures of this type of condition. When this is not specified, it defaults to Error.",
														Type:        "string",
													},
													"status": {
														Description: "Status of the condition, one of True, False, Unknown.",
														Type:        "string",
													},
													"type": {
														Description: "Type of condition.",
														Type:        "string",
													},
												},
												Required: []string{
													"status",
													"type",
												},
												Type: "object",
											},
										},
										Type: "array",
									},
									"nodeName": {
										Description: "NodeName is the name of the corresponding node object",
										Type:        "string",
									},
									"providerID": {
										Description: "ProviderID of the corresponding node object",
										Type:        "string",
									},
								},
								Type: "object",
							},
						},
						Type: "object",
					},
				},
				Served:       true,
				Storage:      true,
				Subresources: &apiextensionsv1.CustomResourceSubresources{},
			},
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "apiextensions.k8s.io/v1",
		Kind:       "CustomResourceDefinition",
	},
}
