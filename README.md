# RPSL-Go

RPSL-Go is a Go library for parsing and handling [RFC 2622: Routing Policy Specification Language (RPSL)](https://datatracker.ietf.org/doc/rfc2622/) attributes and objects.

## Installation

To install the library, use `go get`:

```sh
go get github.com/frederic-arr/rpsl-go
```

## Usage

### Parsing a single RPSL object

```go
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
		"mnt-by:	FOO-MNT\n" +
		"mnt-by:	BAR-MNT\n" +
		"source:	RIPE"

	obj, err := rpsl.Parse(raw)
	if err != nil {
		log.Fatalf("parseObject => %v", err)
	}

	address := obj.GetFirst("address")
	fmt.Printf("Address:\n")
	fmt.Printf("%s\n\n", address)

	maintainers := obj.GetAll("mnt-by")
	fmt.Printf("Maintainers:\n")
	for _, m := range maintainers {
		fmt.Printf("%s\n", m)
	}
}

```

## License
The source code of this project is licensed under the MIT License. For more information, see [LICENSE](./LICENSE).
