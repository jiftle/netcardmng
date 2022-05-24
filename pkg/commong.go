package pkg

type NetCardInfo struct {
	Type               string
	BootProto          string
	Defroute           bool
	PeerDns            bool
	PeerRoutes         bool
	IPV4_FAILURE_FATAL bool
	IPV6Init           bool
	IPV6_AUTOCONF      bool
	IPV6_DefRoute      bool
	IPV6_PeerDNS       bool
	IPV6_PeerRoutes    bool
	IPV6_FAILURE_FATAL bool
	IPV6_ADDR_GEN_MODE string
	Name               string
	UUID               string
	Device             string
	OnBoot             bool
	MAC                string
	IpAddr             string
	Prefix             int32
	PrefixStr          string
	Gateway            string
	DNS1               string
	DNS2               string

	// interal
	File string
}

// IPADDR=192.168.116.46
// PREFIX=24
// GATEWAY=192.168.116.1
// DNS1=211.138.24.66

//TYPE="Ethernet"
//BOOTPROTO="dhcp"
//DEFROUTE="yes"
//PEERDNS="yes"
//PEERROUTES="yes"
//IPV4_FAILURE_FATAL="no"
//IPV6INIT="yes"
//IPV6_AUTOCONF="yes"
//IPV6_DEFROUTE="yes"
//IPV6_PEERDNS="yes"
//IPV6_PEERROUTES="yes"
//IPV6_FAILURE_FATAL="no"
//IPV6_ADDR_GEN_MODE="stable-privacy"
//NAME="enp0s3"
//UUID="eb3b92bd-bdca-43c2-8ff6-fd5d7806bea3"
//DEVICE="enp0s3"
//ONBOOT="yes"
