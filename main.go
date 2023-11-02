package main

import (
	"fmt"

	"github.com/shima-park/agollo"
)

func main() {
	client, err := agollo.New("localhost:8080", "SampleApp", agollo.PreloadNamespaces("application"))
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
				"Old Value:", resp.OldValue,
				"New Value: ", resp.NewValue,
				"Error: ", resp.Error,
			)
		}
	}
}
