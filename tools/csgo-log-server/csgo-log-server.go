package main

import (
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

		jsn, _ := json.MarshalIndent(response, "", "  ")

		fmt.Println(string(jsn))
	}

}
