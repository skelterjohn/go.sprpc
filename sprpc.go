package sprpc

import (
	"io"
	"net/rpc"
	"os/exec"
)

type IOCloser struct {
	io.Reader
	io.WriteCloser
}

type GuestPeer struct {
	*rpc.Client
}

func NewGuestPeer(args ...string) (gp GuestPeer, err error) {
	// First run the process, collecting in/out.
	cmd := exec.Command(args[0], args[1:]...)
	in, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	out, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	if err != nil {
		return
	}

	gp = GuestPeer{rpc.NewClient(IOCloser{Reader: in, WriteCloser: out})}
	return
}

type HostPeer struct {
	*rpc.Server
	ioc IOCloser
}

func NewHostPeer(in io.Reader, out io.WriteCloser) (hp *HostPeer) {
	hp = &HostPeer{
		Server: rpc.NewServer(),
		ioc:    IOCloser{Reader: in, WriteCloser: out},
	}
	return
}

func (hp *HostPeer) Serve() {
	hp.ServeConn(hp.ioc)
}
