package attach

import (
	"errors"

	"github.com/anhk/flowctl/pkg/exception"
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/florianl/go-tc"
	"github.com/florianl/go-tc/core"
	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)

func AttachCgroup(pinPath, cgroupPath string, attachType ebpf.AttachType, prog *ebpf.Program) {
	if l, err := link.LoadPinnedLink(pinPath, &ebpf.LoadPinOptions{}); err == nil {
		exception.Must(l.Update(prog))
		exception.Must(l.Close())
		return
	}

	l, err := link.AttachCgroup(link.CgroupOptions{
		Path:    cgroupPath,
		Attach:  attachType,
		Program: prog,
	})
	exception.Must(err)
	defer func() { exception.Must(l.Close()) }()

	err = l.Pin(pinPath)
	if errors.Is(err, link.ErrNotSupported) { // linux 4.x not support link.Pin()
		err = nil
	}
	exception.Must(err)
}

func AttachTc(ifIndex uint32, ingress, egress *ebpf.Program) {
	tcnl, err := tc.Open(&tc.Config{})
	exception.Must(err)
	defer func() { tcnl.Close() }()

	exception.Must(tcnl.SetOption(netlink.ExtendedAcknowledge, true))

	qdiscs, err := tcnl.Qdisc().Get()
	exception.Must(err)

	// clear qdisc
	for _, qdisc := range qdiscs {
		if qdisc.Msg.Ifindex == ifIndex && qdisc.Attribute.Kind == "clsact" {
			exception.Must(tcnl.Qdisc().Delete(&qdisc))
		}
	}

	// add qdisc
	exception.Must(tcnl.Qdisc().Add(&tc.Object{Msg: tc.Msg{
		Family:  unix.AF_UNSPEC,
		Ifindex: ifIndex,
		Handle:  core.BuildHandle(0xFFFF, 0x0000),
		Parent:  tc.HandleIngress,
	}, Attribute: tc.Attribute{Kind: "clsact"}}))

	// attach ingress
	infoIngress, err := ingress.Info()
	exception.Must(err)

	tcnl.Filter().Add(&tc.Object{Msg: tc.Msg{
		Family:  unix.AF_UNSPEC,
		Ifindex: ifIndex,
		Handle:  0,
		Parent:  core.BuildHandle(0xFFFF, tc.HandleMinIngress),
		Info:    0x10300,
	}, Attribute: tc.Attribute{Kind: "bpf", BPF: &tc.Bpf{
		FD:    Pointer(uint32(ingress.FD())),
		Name:  Pointer(infoIngress.Name),
		Flags: Pointer(uint32(0x1)),
	}}})

	// attach egress
	infoEgress, err := egress.Info()
	exception.Must(err)

	tcnl.Filter().Add(&tc.Object{Msg: tc.Msg{
		Family:  unix.AF_UNSPEC,
		Ifindex: ifIndex,
		Handle:  0,
		Parent:  core.BuildHandle(0xFFFF, tc.HandleMinEgress),
		Info:    0x10300,
	}, Attribute: tc.Attribute{Kind: "bpf", BPF: &tc.Bpf{
		FD:    Pointer(uint32(egress.FD())),
		Name:  Pointer(infoEgress.Name),
		Flags: Pointer(uint32(0x1)),
	}}})
}

func Pointer[T any](v T) *T {
	return &v
}
