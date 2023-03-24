// Copyright (c) Volvo Car AB
// SPDX-License-Identifier: Apache-2.0

package terra

import "fmt"

func ExampleBool() {
	b := Bool(true)
	fmt.Println(string(b.InternalTokens().Bytes()))
	// 	Output: true
}
