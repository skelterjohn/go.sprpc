package main

import (
	"github.com/skelterjohn/go.sprpc"
	"os"
)

type Things struct {
	X int
}

func (d *Things) Increment(in int, out *int) (err error) {
	d.X++
	*out = d.X
	return
}

func (d *Things) Fetch(in int, out *int) (err error) {
	*out = d.X
	return
}

func main() {
	os.Stderr.WriteString("Guest running.")
	d := &Things{}

	hp := sprpc.NewHostPeer(os.Stdin, os.Stdout)
	hp.RegisterName("Things", d)
	hp.Serve()
}
