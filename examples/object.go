package main

import (
	"fmt"
	"log"

	"github.com/frederic-arr/rpsl-go"
)

func main() {
	lines := []string{
		"person:        John Doe",
		"address:       1234 Elm Street",
		"               Iceland",
		"phone:         +1 555 123456",
		"nic-hdl:       JD1234-RIPE",
		"mnt-by:        EXAMPLE-MNT",
		"source:        RIPE",
	}
	i := 0
	obj, err := rpsl.ParseObject(&i, lines)
	if err != nil {
		log.Fatalf("parseObject => %v", err)
	}

	for _, attr := range obj.Attributes {
		fmt.Printf("Attribute Name: %s, Attribute Value: %v\n", attr.Name, attr.Value)
	}
}
