package app

import (
	"os"

	"github.com/anhk/flowctl/ebpf/objs"
	"github.com/anhk/flowctl/pkg/attach"
	"github.com/anhk/flowctl/pkg/exception"
)

const (
	pinPath   = "/sys/fs/bpf/flowctl"
	cgrouPath = "/sys/fs/cgroup"
)

type Loadbalance struct {
}

func NewLoadbalance() *Loadbalance {
	return &Loadbalance{}
}

func (lb *Loadbalance) Install() error {
	return exception.TryWithError(func() {
		exception.Must(os.Mkdir(pinPath, 0x755))
		objs.NewObject(pinPath).AttachCgroup(cgrouPath).AttachTc(2)
	})
}

func (lb *Loadbalance) UnInstall() error {
	return exception.TryWithError(func() {
		attach.ClearTc(2)
		exception.Must(os.RemoveAll(pinPath))
	})
}

func (lb *Loadbalance) Add() {

}

func (lb *Loadbalance) Del() {

}

func (lb *Loadbalance) Clear() {

}
