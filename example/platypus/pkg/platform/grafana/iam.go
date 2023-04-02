// Copyright (c) 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package grafana

import rbacv1 "k8s.io/api/rbac/v1"

var RoleRules = []rbacv1.PolicyRule{
	{
		APIGroups:     []string{"extensions"},
		ResourceNames: []string{"grafana"},
		Resources:     []string{"podsecuritypolicies"},
		Verbs:         []string{"use"},
	},
}