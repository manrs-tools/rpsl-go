// Copyright (c) The RPSL Go Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"

	"github.com/frederic-arr/rpsl-go"
)

func main() {
	raw := "person:	John Doe\n" +
		"address:	1234 Elm Street Iceland\n" +
		"phone:		+1 555 123456\n" +
		"nic-hdl:	JD1234-RIPE\n" +
		"mnt-by:	FOO-MNT\n" +
		"mnt-by:	BAR-MNT\n" +
		"source:	RIPE"

	obj, err := rpsl.Parse(raw)
	if err != nil {
		log.Fatalf("parseObject => %v", err)
	}

	address := obj.GetFirst("address")
	fmt.Printf("--- Address ---\n")
	fmt.Printf("%s\n\n", address)

	maintainers := obj.GetAll("mnt-by")
	fmt.Printf("--- Maintainers ---\n")
	for _, mntner := range maintainers {
		fmt.Printf("%s\n", mntner)
	}
}
