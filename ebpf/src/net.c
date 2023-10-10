
#include "net.h"

/*****************************************************************************
 * CGroupV2
 *****************************************************************************/

SEC("cgroup/connect4")
int sock_connect4(struct bpf_sock_addr *ctx)
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
