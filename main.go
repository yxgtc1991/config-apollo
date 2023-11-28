package main

import (
	"fmt"

	"github.com/shima-park/agollo"
)

func main() {
	client, err := agollo.New("127.0.0.1:8080", "SampleApp", agollo.PreloadNamespaces("application"))
	if err != nil {
		panic(err)
	}
	errCh := client.Start()
	watchCh := client.Watch()
	for {
		select {
		case err := <-errCh:
			fmt.Println("Error: ", err)
		case resp := <-watchCh:
			fmt.Println(
				"Namespace: ", resp.Namespace,
				"\t Old Value:", resp.OldValue,
				"\t New Value: ", resp.NewValue,
				"\t Error: ", resp.Error,
			)
		}
	}
}
