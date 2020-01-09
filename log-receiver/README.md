[![Godoc](https://godoc.org/github.com/janstuemmel/csgo-go/log-receiver?status.svg)](https://godoc.org/github.com/janstuemmel/csgo-go/log-receiver)

# csgo logreceiver

A udp handler to receive logs from a remote srcds csgo server.
It parses the header sent via udp and returns data in a response struct. 

## Usage

```go
package main

import (
	"fmt"

	csgologreceiver "github.com/janstuemmel/csgo-go/log-receiver"
)

func main() {

	receiver, _ := csgologreceiver.New("0.0.0.0", 1234)

	for {
		response, _ := receiver.Read()
		fmt.Print(response)
	}
}
```