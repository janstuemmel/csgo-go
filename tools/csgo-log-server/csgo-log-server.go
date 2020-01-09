package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	csgologreceiver "github.com/janstuemmel/csgo-go/log-receiver"
)

func main() {

	receiver, err := csgologreceiver.New("0.0.0.0", 1234)

	if err != nil {
		panic(err)
	}

	for {

		response, _ := receiver.Read()

		fmt.Printf(ToJSON(response))
	}

}

// helper
func ToJSON(m csgologreceiver.Response) string {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	enc.Encode(m)
	return buf.String()
}
