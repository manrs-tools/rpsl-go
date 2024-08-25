# RPSL-Go

RPSL-Go is a Go library for parsing and handling Routing Policy Specification Language (RPSL) attributes and objects.

## Installation

To install the library, use `go get`:

```sh
go get github.com/frederic-arr/rpsl-go
```

## Usage

You can parse RPSL objects using the `ParseObject` function. Here are some examples:

```go
package main

import (
    "fmt"
    "log"

    rpslgo "github.com/frederic-arr/rpsl-go"
)

func main() {
    lines := []string{
        "person:        John Doe",
        "address:       1234 Elm Street",
        "phone:         +1 555 123456",
        "nic-hdl:       JD1234-RIPE",
        "mnt-by:        EXAMPLE-MNT",
        "source:        RIPE",
    }
    i := 0
    obj, err := rpslgo.ParseObject(&i, lines)
    if err != nil {
        log.Fatalf("parseObject => %v", err)
    }
    fmt.Printf("Object Type: %s\n", obj.Type)
    for _, attr := range obj.Attributes {
        fmt.Printf("Attribute Name: %s, Attribute Value: %v\n", attr.Name, attr.Value)
    }
}
```
