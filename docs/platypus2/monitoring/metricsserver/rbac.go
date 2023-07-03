// Copyright (c) 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package metricsserver

import (
	ku "github.com/volvo-cars/lingon/pkg/kubeutil"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AuthReaderRB = &rbacv1.RoleBinding{
	TypeMeta: ku.TypeRoleBindingV1,
	ObjectMeta: metav1.ObjectMeta{
		Labels:    M.Labels(),
		Name:      M.Name + "-auth-reader",
		Namespace: ku.NSKubeSystem,
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: ku.TypeRoleV1.GroupVersionKind().Group, // "rbac.authorization.k8s.io",
		Kind:     ku.TypeRoleV1.Kind,
		Name:     "extension-apiserver-authentication-reader", // predefined ?
	},
	Subjects: []rbacv1.Subject{
		{
			Kind:      ku.TypeServiceAccountV1.Kind,
			Name:      SA.Name,
			Namespace: SA.Namespace,
		},
	},
}

var SystemAggregatedReaderCR = &rbacv1.ClusterRole{
	TypeMeta: ku.TypeClusterRoleV1,
	ObjectMeta: metav1.ObjectMeta{
		Labels: ku.MergeLabels(
			M.Labels(), map[string]string{
				ku.LabelRbacAggregateToAdmin: "true",
				ku.LabelRbacAggregateToEdit:  "true",
				ku.LabelRbacAggregateToView:  "true",
			},
		),
		Name: "system:metrics-server-aggregated-reader",
	},
	Rules: []rbacv1.PolicyRule{
		{
			APIGroups: []string{"metrics.k8s.io"},
			Resources: []string{"pods", "nodes"},
			Verbs:     []string{"get", "list", "watch"},
		},
	},
}

var SystemCR = &rbacv1.ClusterRole{
	TypeMeta: ku.TypeClusterRoleV1,
	ObjectMeta: metav1.ObjectMeta{
		Labels: M.Labels(),
		Name:   "system:" + M.Name,
	},
	Rules: []rbacv1.PolicyRule{
		{
			APIGroups: []string{""},
			Resources: []string{"nodes/metrics"},
			Verbs:     []string{"get"},
		}, {
			APIGroups: []string{""},
			Resources: []string{"pods", "nodes", "namespaces", "configmaps"},
			Verbs:     []string{"get", "list", "watch"},
		},
	},
}

var SystemAuthDelegatorCRB = &rbacv1.ClusterRoleBinding{
	TypeMeta: ku.TypeClusterRoleBindingV1,
	ObjectMeta: metav1.ObjectMeta{
		Labels: M.Labels(),
		Name:   M.Name + ":system:auth-delegator",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: ku.TypeClusterRoleV1.GroupVersionKind().Group, // "rbac.authorization.k8s.io",
		Kind:     ku.TypeClusterRoleV1.Kind,
		Name:     "system:auth-delegator",
	},
	Subjects: []rbacv1.Subject{
		{
			Kind:      ku.TypeServiceAccountV1.Kind,
			Name:      SA.Name,
			Namespace: SA.Namespace,
		},
	},
}

var SystemCRB = &rbacv1.ClusterRoleBinding{
	TypeMeta: ku.TypeClusterRoleBindingV1,
	ObjectMeta: metav1.ObjectMeta{
		Labels: M.Labels(),
		Name:   "system:" + M.Name,
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: ku.TypeClusterRoleV1.GroupVersionKind().Group, // "rbac.authorization.k8s.io",
		Kind:     ku.TypeClusterRoleV1.Kind,
		Name:     SystemCR.Name,
	},
	Subjects: []rbacv1.Subject{
		{
			Kind:      ku.TypeServiceAccountV1.Kind,
			Name:      SA.Name,
			Namespace: SA.Namespace,
		},
	},
}