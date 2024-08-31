package main

import (
	"fmt"
	"log"

	"github.com/frederic-arr/rpsl-go"
)

func main() {
	raw := "person:	John Doe\n" +
		"address:	1234 Elm Street\n" +
		"			Iceland\n" +
		"phone:		+1 555 123456\n" +
		"nic-hdl:	JD1234-RIPE\n" +
		"mnt-by:	EXAMPLE-MNT\n" +
		"source:	RIPE"

	obj, err := rpsl.Parse(raw)
	if err != nil {
		log.Fatalf("parseObject => %v", err)
	}

	for _, attr := range obj.Attributes {
		fmt.Printf("Attribute Name: %s, Attribute Value: %v\n", attr.Name, attr.Value)
	}
}
