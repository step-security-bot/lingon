// Copyright (c) 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

package vector

import (
	"os"
	"testing"

	"github.com/golingon/lingon/pkg/kube"
	tu "github.com/golingon/lingon/pkg/testutil"
)

func TestExport(t *testing.T) {
	tu.AssertNoError(t, os.RemoveAll("out"))
	tu.AssertNoError(
		t,
		kube.Export(New(), kube.WithExportOutputDirectory("out")),
		"export",
	)
}
