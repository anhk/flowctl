
#include "net.h"

/*****************************************************************************
 * CGroupV2
 *****************************************************************************/

SEC("cgroup/connect4")
int sock_connect4(struct bpf_sock_addr *ctx)
{
    return SK_PASS;
}

SEC("cgroup/connect6")
int sock_connect6(struct bpf_sock_addr *ctx)
{
    return SK_PASS;
}

SEC("cgroup/sendmsg4")
int sock_sendmsg4(struct bpf_sock_addr *ctx)
{
    return SK_PASS;
}

SEC("cgroup/sendmsg6")
int sock_sendmsg6(struct bpf_sock_addr *ctx)
{
    return SK_PASS;
}

SEC("cgroup/recvmsg4")
int sock_recvmsg4(struct bpf_sock_addr *ctx)
{
    return SK_PASS;
}

SEC("cgroup/recvmsg6")
int sock_recvmsg6(struct bpf_sock_addr *ctx)
{
    return SK_PASS;
}
SEC("cgroup/getpeername4")
int sock_getpeername4(struct bpf_sock_addr *ctx)
{
    return SK_PASS;
}

SEC("cgroup/getpeername6")
int sock_getpeername6(struct bpf_sock_addr *ctx)
{
    return SK_PASS;
}
/*****************************************************************************
 * TC
 *****************************************************************************/

SEC("classifier/ingress")
int tc_ingress(struct __sk_buff *skb)
{
    return TC_ACT_OK;
}

SEC("classifier/egress")
int tc_egress(struct __sk_buff *skb)
{
    return TC_ACT_OK;
}
