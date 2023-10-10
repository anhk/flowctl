package objs

import (
	"fmt"

	"github.com/anhk/flowctl/pkg/attach"
	"github.com/anhk/flowctl/pkg/exception"
	"github.com/cilium/ebpf"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target bpfel ebpf ../src/net.c -- -I ../inc

type Object struct {
	objs    ebpfObjects
	pinPath string
}

func NewObject(pinPath string) *Object {
	o := &Object{pinPath: pinPath}
	exception.Must(loadEbpfObjects(&o.objs, &ebpf.CollectionOptions{
		Maps: ebpf.MapOptions{PinPath: pinPath},
	}))
	return o
}

func (o *Object) AttachCgroup(cgroupPath string) {
	var attachArray = []struct {
		name       string
		attachType ebpf.AttachType
		prog       *ebpf.Program
	}{
		{"connect4", ebpf.AttachCGroupInet4Connect, o.objs.SockConnect4},
	}

	for _, att := range attachArray {
		attach.AttachCgroup(fmt.Sprintf("%v/%v", o.pinPath, att.name),
			cgroupPath, att.attachType, att.prog)
	}
}

func (o *Object) AttachTc(ifIndex uint32) {
	attach.AttachTc(ifIndex, o.objs.TcIngress, o.objs.TcEgress)
}
