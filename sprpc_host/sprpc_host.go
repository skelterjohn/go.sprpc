package main

import (
	"fmt"
	"github.com/skelterjohn/go.sprpc"
	"log"
)

func main() {
	gp, err := sprpc.NewGuestPeer("sprpc_guest")
	if err != nil {
		log.Fatal(err)
	}

	var x int

	err = gp.Call("Things.Increment", 5, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = gp.Call("Things.Increment", 5, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = gp.Call("Things.Increment", 5, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = gp.Call("Things.Fetch", 5, &x)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(x)
}
