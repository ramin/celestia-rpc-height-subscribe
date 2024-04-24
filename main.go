package main

import (
	"context"
	"fmt"

	"github.com/celestiaorg/celestia-node/api/rpc/client"
)

func main() {
	rpc, err := client.NewClient(
		context.Background(),
		"ws://0.0.0.0:26658", // note ws for subscription

		// generate your own light node JWT here (celestia light auth read)
		"",
	)

	ctx := context.Background()

	if err != nil {
		fmt.Println(err)
		return
	}

	listen, err := rpc.Header.Subscribe(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		select {
		case header := <-listen:
			fmt.Println("New block")
			fmt.Println(header.Height())
		case <-ctx.Done():
			fmt.Println("Done")
			return
		}

	}
}
