# RPSL Go

RPSL Go is a Go library for parsing and handling [RFC 2622: Routing Policy Specification Language (RPSL)](https://datatracker.ietf.org/doc/rfc2622/) attributes and objects.

> [!IMPORTANT]  
> The goal of this library is to parse responses provided by the RIPE database. It is not a full implementation of the RPSL specification. See the [Restrictions](#restrictions) section for more information.

## Installation

To install the library, use `go get`:

```sh
go get github.com/frederic-arr/rpsl-go
```

## Usage

```go
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
	fmt.Printf("%s\n\n", address.Value)

	maintainers := obj.GetAll("mnt-by")
	fmt.Printf("--- Maintainers ---\n")
	for _, m := range maintainers {
		fmt.Printf("%s\n", m.Value)
	}
}

```

See the output by running `go run examples/object.go` in this repository.


## Restrictions

- It is case-sensitive.
- Multi-line attributes are not supported.
- Attribute names shall only contains characters from the set `[a-z0-9-]`.
- No validation regarding the object is performed.
- No validation regarding the attribute values is performed.


## License

Copyright (c) The RPSL Go Authors. [Apache 2.0 License](./LICENSE).
