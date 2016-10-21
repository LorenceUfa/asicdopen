//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package pluginCommon

import (
	"net"
)

/*
#include "pluginCommon.h"
*/
import "C"

var SwitchMacAddr net.HardwareAddr

//Consolidated list of constants used by underlying plugins
const (
	//Generic consts
	DEFAULT_SWITCH_MAC_ADDR = "00:11:22:33:44:55"
	MAC_ADDR_LEN            = C.MAC_ADDR_LEN
	INVALID_OBJECT_ID       = 0xFFFFFFFFFFFFFFFF
	INVALID_IFINDEX         = -1

	//System related consts
	BOOT_MODE_COLDBOOT = C.BOOT_MODE_COLDBOOT
	BOOT_MODE_WARMBOOT = C.BOOT_MODE_WARMBOOT
	INTF_STATE_UP      = C.INTF_STATE_UP
	INTF_STATE_DOWN    = C.INTF_STATE_DOWN
	MIN_SYS_PORTS      = 0
	MAX_SYS_PORTS      = 256
	ASICD_CONFIG_FILE  = "asicdConf.json"
	MAX_VENDOR_ID_LEN  = C.MAX_VENDOR_ID_LEN
	MAX_PART_NUM_LEN   = C.MAX_PART_NUM_LEN
	MAX_REV_ID_LEN     = C.MAX_REV_ID_LEN

	//FDB relate consts
	MAC_ENTRY_LEARNED = C.MAC_ENTRY_LEARNED
	MAC_ENTRY_AGED    = C.MAC_ENTRY_AGED

	//Vlan related consts
	SVI_PREFIX        = "vlan"
	SYS_RSVD_VLAN     = -1
	MAX_VLAN_ID       = C.MAX_VLAN_ID
	DEFAULT_VLAN_ID   = C.DEFAULT_VLAN_ID
	SYS_RSVD_VLAN_MIN = 3835
	SYS_RSVD_VLAN_MAX = 4090

	//Port related consts
	MAX_PORT_STAT_TYPES                   = C.portStatTypesMax
	MAX_BUFFER_STAT_TYPES                 = C.bufferPortStatTypesMax
	MAX_BUFFER_GLOBAL_STAT_TYPES          = C.bufferGlobalStatTypesMax
	PORT_BREAKOUT_MODE_UNSUPPORTED_STRING = "unsupported"
	PORT_BREAKOUT_MODE_UNSUPPORTED        = C.PORT_BREAKOUT_MODE_UNSUPPORTED
	PORT_BREAKOUT_MODE_1x40               = C.PORT_BREAKOUT_MODE_1x40
	PORT_BREAKOUT_MODE_4x10               = C.PORT_BREAKOUT_MODE_4x10
	PORT_BREAKOUT_MODE_1x100              = C.PORT_BREAKOUT_MODE_1x100
	PORT_ATTR_PHY_INTF_TYPE               = C.PORT_ATTR_PHY_INTF_TYPE
	PORT_ATTR_ADMIN_STATE                 = C.PORT_ATTR_ADMIN_STATE
	PORT_ATTR_MAC_ADDR                    = C.PORT_ATTR_MAC_ADDR
	PORT_ATTR_SPEED                       = C.PORT_ATTR_SPEED
	PORT_ATTR_DUPLEX                      = C.PORT_ATTR_DUPLEX
	PORT_ATTR_AUTONEG                     = C.PORT_ATTR_AUTONEG
	PORT_ATTR_MEDIA_TYPE                  = C.PORT_ATTR_MEDIA_TYPE
	PORT_ATTR_MTU                         = C.PORT_ATTR_MTU
	PORT_ATTR_BREAKOUT_MODE               = C.PORT_ATTR_BREAKOUT_MODE
	PORT_ATTR_LOOPBACK_MODE               = C.PORT_ATTR_LOOPBACK_MODE
	PORT_ATTR_ENABLE_FEC                  = C.PORT_ATTR_ENABLE_FEC
	PORT_ATTR_TX_PRBS_EN                  = C.PORT_ATTR_TX_PRBS_EN
	PORT_ATTR_RX_PRBS_EN                  = C.PORT_ATTR_RX_PRBS_EN
	PORT_ATTR_PRBS_POLY                   = C.PORT_ATTR_PRBS_POLY
	PORT_MODE_UNCONFIGURED                = "Unconfigured"
	PORT_MODE_L2                          = "L2"
	PORT_MODE_L3                          = "L3"
	PORT_MODE_INTERNAL                    = "Internal"
	PORT_LOOPBACK_MODE_NONE               = C.PORT_LOOPBACK_MODE_NONE
	PORT_LOOPBACK_MODE_MAC                = C.PORT_LOOPBACK_MODE_MAC
	PORT_LOOPBACK_MODE_PHY                = C.PORT_LOOPBACK_MODE_PHY

	//Intf related constants
	IP_TYPE_IPV4 = C.IP_TYPE_IPV4
	IP_TYPE_IPV6 = C.IP_TYPE_IPV6

	//STP related consts
	STP_PORT_STATE_BLOCKING   = C.StpPortStateBlocking
	STP_PORT_STATE_LEARNING   = C.StpPortStateLearning
	STP_PORT_STATE_FORWARDING = C.StpPortStateForwarding

	//Lag related consts
	LAG_PREFIX             = "LAG"
	HASHTYPE_SRCMAC_DSTMAC = C.HASHTYPE_SRCMAC_DSTMAC
	HASHTYPE_SRCIP_DSTIP   = C.HASHTYPE_SRCIP_DSTIP

	//Next hop related consts
	NEIGHBOR_TYPE_COPY_TO_CPU       = C.NEIGHBOR_TYPE_COPY_TO_CPU
	NEIGHBOR_TYPE_BLACKHOLE         = C.NEIGHBOR_TYPE_BLACKHOLE
	NEIGHBOR_TYPE_FULL_SPEC_NEXTHOP = C.NEIGHBOR_TYPE_FULL_SPEC_NEXTHOP
	NEIGHBOR_L2_ACCESS_TYPE_PORT    = C.NEIGHBOR_L2_ACCESS_TYPE_PORT
	NEIGHBOR_L2_ACCESS_TYPE_LAG     = C.NEIGHBOR_L2_ACCESS_TYPE_LAG

	//Next hop related consts
	NEXTHOP_TYPE_COPY_TO_CPU       = C.NEXTHOP_TYPE_COPY_TO_CPU
	NEXTHOP_TYPE_BLACKHOLE         = C.NEXTHOP_TYPE_BLACKHOLE
	NEXTHOP_TYPE_FULL_SPEC_NEXTHOP = C.NEXTHOP_TYPE_FULL_SPEC_NEXTHOP
	NEXTHOP_L2_ACCESS_TYPE_PORT    = C.NEXTHOP_L2_ACCESS_TYPE_PORT
	NEXTHOP_L2_ACCESS_TYPE_LAG     = C.NEXTHOP_L2_ACCESS_TYPE_LAG
	INVALID_NEXTHOP_ID             = 0xFFFFFFFFFFFFFFFF

	// vxlan related consts
	VXLAN_VTEP_PREFIX  = "Vtep"
	VXLAN_VXLAN_PREFIX = "Vxlan"

	//Route related consts
	MAX_NEXTHOPS_PER_GROUP      = C.MAX_NEXTHOPS_PER_GROUP
	ROUTE_OPERATION_TYPE_UPDATE = C.ROUTE_OPERATION_TYPE_UPDATE
	ROUTE_TYPE_CONNECTED        = C.ROUTE_TYPE_CONNECTED
	ROUTE_TYPE_MULTIPATH        = C.ROUTE_TYPE_MULTIPATH
	ROUTE_TYPE_SINGLEPATH       = C.ROUTE_TYPE_SINGLEPATH
	ROUTE_TYPE_V6               = C.ROUTE_TYPE_V6
	ROUTE_TYPE_NULL             = C.ROUTE_TYPE_NULL

	//ACL and COPP related consts
	MAX_COPP_CLASS_COUNT = C.CoppClassCount
)

// Interface id/type mgmt
const (
	INTF_TYPE_MASK  = C.INTF_TYPE_MASK
	INTF_TYPE_SHIFT = C.INTF_TYPE_SHIFT
	INTF_ID_MASK    = C.INTF_ID_MASK
	INTF_ID_SHIFT   = C.INTF_ID_SHIFT
)

const (
	PORT_PROTOCOL_ARP       = 0x1
	PORT_PROTOCOL_DHCP      = 0x2
	PORT_PROTOCOL_DHCPRELAY = 0x4
	PORT_PROTOCOL_BGP       = 0x8
	PORT_PROTOCOL_OSPF      = 0x10
	PORT_PROTOCOL_VXLAN     = 0x20
	PORT_PROTOCOL_MPLS      = 0x40
	PORT_PROTOCOL_BFD       = 0x40
)

const (
	ACL_DIRECTION_IN  = "IN"
	ACL_DIRECTION_OUT = "OUT"
	ACL_TYPE_IP       = "IP"
	ACL_TYPE_MAC      = "MAC"
	ACL_TYPE_MLAG     = "MLAG"
)

// Func types for intf id mgmt
type GetId func(int32) int
type GetType func(int32) int
type GetIfIndex func(int, int) int32

func GetTypeFromIfIndex(ifIndex int32) int {
	return int((ifIndex & INTF_TYPE_MASK) >> INTF_TYPE_SHIFT)
}
func GetIdFromIfIndex(ifIndex int32) int {
	return int((ifIndex & INTF_ID_MASK) >> INTF_ID_SHIFT)
}
func GetIfIndexFromIdType(ifId, ifType int) int32 {
	return int32((ifId & INTF_ID_MASK) | ((ifType << INTF_TYPE_SHIFT) & INTF_TYPE_MASK))
}
func IsZeroIP(ip []uint8, ipType int) bool {
	/*    if ipType == syscall.AF_INET {
	    if bytes.Equal(ip,net.IPv4zero) {
		    return true
	    }
	} else {
	    if bytes.Equal(ip,net.IPv6zero) {
		    return true
	    }
	}
	return false*/
	for i := 0; i < len(ip); i++ {
		if ip[i] != 0 {
			return false
		}
	}
	return true
}

const (
	LB_MODE_NONE = "NONE"
	LB_MODE_MAC  = "MAC"
	LB_MODE_PHY  = "PHY"
	LB_MODE_RMT  = "RMT"
)

var LBModeEnumToStr map[uint8]string = map[uint8]string{
	uint8(C.PORT_LOOPBACK_MODE_NONE): LB_MODE_NONE,
	uint8(C.PORT_LOOPBACK_MODE_MAC):  LB_MODE_MAC,
	uint8(C.PORT_LOOPBACK_MODE_PHY):  LB_MODE_PHY,
	uint8(C.PORT_LOOPBACK_MODE_RMT):  LB_MODE_RMT,
}

var LBModeStrToEnum map[string]uint8 = map[string]uint8{
	LB_MODE_NONE: uint8(C.PORT_LOOPBACK_MODE_NONE),
	LB_MODE_MAC:  uint8(C.PORT_LOOPBACK_MODE_MAC),
	LB_MODE_PHY:  uint8(C.PORT_LOOPBACK_MODE_PHY),
	LB_MODE_RMT:  uint8(C.PORT_LOOPBACK_MODE_RMT),
}

const (
	PRBS_POLY_2POW7  = "2^7"
	PRBS_POLY_2POW23 = "2^23"
	PRBS_POLY_2POW31 = "2^31"
)

var PrbsStrToEnum map[string]uint8 = map[string]uint8{
	PRBS_POLY_2POW7:  uint8(C.PRBS_POLY_2POW7),
	PRBS_POLY_2POW23: uint8(C.PRBS_POLY_2POW23),
	PRBS_POLY_2POW31: uint8(C.PRBS_POLY_2POW31),
}

var PrbsEnumToStr map[uint8]string = map[uint8]string{
	uint8(C.PRBS_POLY_2POW7):  PRBS_POLY_2POW7,
	uint8(C.PRBS_POLY_2POW23): PRBS_POLY_2POW23,
	uint8(C.PRBS_POLY_2POW31): PRBS_POLY_2POW31,
}

var OnOffState map[int]string = map[int]string{0: "OFF", 1: "ON"}

const (
	STATE_UP   = "UP"
	STATE_DOWN = "DOWN"
)

var UpDownState map[int]string = map[int]string{0: STATE_DOWN, 1: STATE_UP}
var UpDownStateStr map[string]int = map[string]int{STATE_DOWN: 0, STATE_UP: 1}
var IfType map[int]string = map[int]string{
	int(C.PortIfTypeMII):    "MII",
	int(C.PortIfTypeGMII):   "GMII",
	int(C.PortIfTypeSGMII):  "SGMII",
	int(C.PortIfTypeQSGMII): "QSGMII",
	int(C.PortIfTypeSFI):    "SFI",
	int(C.PortIfTypeXFI):    "XFI",
	int(C.PortIfTypeXAUI):   "XAUI",
	int(C.PortIfTypeXLAUI):  "XLAUI",
	int(C.PortIfTypeRXAUI):  "RXAUI",
	int(C.PortIfTypeCR):     "CR",
	int(C.PortIfTypeCR2):    "CR2",
	int(C.PortIfTypeCR4):    "CR4",
	int(C.PortIfTypeKR):     "KR",
	int(C.PortIfTypeKR2):    "KR2",
	int(C.PortIfTypeKR4):    "KR4",
	int(C.PortIfTypeSR):     "SR",
	int(C.PortIfTypeSR2):    "SR2",
	int(C.PortIfTypeSR4):    "SR4",
	int(C.PortIfTypeSR10):   "SR10",
	int(C.PortIfTypeLR):     "LR",
	int(C.PortIfTypeLR4):    "LR4",
}

const (
	FULL_DUPLEX = "Full_Duplex"
	HALF_DUPLEX = "Half_Duplex"
)
const (
	PluginOp_NA = iota
	PluginOp_Add
	PluginOp_Del
)

var DuplexType map[int]string = map[int]string{
	int(C.HalfDuplex): HALF_DUPLEX,
	int(C.FullDuplex): FULL_DUPLEX,
}

const (
	//Asicd notification msgs
	NOTIFY_L2INTF_STATE_CHANGE = iota
	NOTIFY_IPV4_L3INTF_STATE_CHANGE
	NOTIFY_IPV6_L3INTF_STATE_CHANGE
	NOTIFY_VLAN_CREATE
	NOTIFY_VLAN_DELETE
	NOTIFY_VLAN_UPDATE
	NOTIFY_LOGICAL_INTF_CREATE
	NOTIFY_LOGICAL_INTF_DELETE
	NOTIFY_LOGICAL_INTF_UPDATE
	NOTIFY_IPV4INTF_CREATE
	NOTIFY_IPV4INTF_DELETE
	NOTIFY_IPV6INTF_CREATE
	NOTIFY_IPV6INTF_DELETE
	NOTIFY_LAG_CREATE
	NOTIFY_LAG_DELETE
	NOTIFY_LAG_UPDATE
	NOTIFY_IPV4NBR_MAC_MOVE
	NOTIFY_IPV6NBR_MAC_MOVE
	NOTIFY_IPV4_ROUTE_CREATE_FAILURE
	NOTIFY_IPV4_ROUTE_DELETE_FAILURE
	NOTIFY_IPV6_ROUTE_CREATE_FAILURE
	NOTIFY_IPV6_ROUTE_DELETE_FAILURE
	NOTIFY_VTEP_CREATE
	NOTIFY_VTEP_DELETE
	NOTIFY_MPLSINTF_STATE_CHANGE
	NOTIFY_MPLSINTF_CREATE
	NOTIFY_MPLSINTF_DELETE
	NOTIFY_PORT_CONFIG_MODE_CHANGE
	NOTIFY_PORT_CONFIG_MTU_CHANGE
)

// Format of asicd's published messages
type AsicdNotification struct {
	MsgType uint8
	Msg     []byte
}

// Following sections contains formats for individual message types
type L2IntfStateNotifyMsg struct {
	IfIndex int32
	IfState uint8
}
type MplsIntfStateNotifyMsg struct {
	IfIndex int32
	IfState uint8
}
type IPv4L3IntfStateNotifyMsg struct {
	IpAddr  string
	IfIndex int32
	IfState uint8
}
type IPv6L3IntfStateNotifyMsg struct {
	IpAddr  string
	IfIndex int32
	IfState uint8
}
type VlanNotifyMsg struct {
	VlanId      uint16
	VlanIfIndex int32
	VlanName    string
	TagPorts    []int32
	UntagPorts  []int32
}
type LogicalIntfNotifyMsg struct {
	IfIndex         int32
	LogicalIntfName string
}
type LagNotifyMsg struct {
	LagName     string
	IfIndex     int32
	IfIndexList []int32
}
type MplsIntfNotifyMsg struct {
	IfIndex int32
}
type IPv4IntfNotifyMsg struct {
	IpAddr  string
	IfIndex int32
	IntfRef string
}
type IPv4NbrMacMoveNotifyMsg struct {
	IpAddr  string
	IfIndex int32
	VlanId  int32
}
type IPv6NbrMacMoveNotifyMsg struct {
	IpAddr  string
	IfIndex int32
	VlanId  int32
}
type IPv6IntfNotifyMsg struct {
	IpAddr  string
	IfIndex int32
	IntfRef string
}
type VtepNotifyMsg struct {
	VtepName   string
	IfIndex    int32
	Vni        int32
	SrcIfIndex int32
	SrcIfName  string
}
type PortConfigModeChgNotifyMsg struct {
	IfIndex int32
	OldMode string
	NewMode string
}
type PortConfigMtuChangeNotifyMsg struct {
	IfIndex int32
	Mtu     int32
}

// Struct containing info required for mapping asic ports to linux interfaces
type IfMapInfo struct {
	IfName string
	Port   int
}

// Struct used for configuring sub interface on all plugins
type SubIntfPluginObj struct {
	IfIndex        int32
	IpAddr         uint32
	MaskLen        int
	VlanId         int
	StateUp        bool
	SubIntfIfIndex int32
	MacAddr        net.HardwareAddr
}

// Struct used for configuring ports
type PortConfig struct {
	PortNum           int32
	PortName          string
	Description       string
	PhyIntfType       string
	AdminState        string
	MacAddr           string
	Speed             int32
	Duplex            string
	Autoneg           string
	MediaType         string
	Mtu               int32
	LogicalPortInfo   bool
	MappedToHw        bool
	BreakOutMode      int32
	BreakOutSupported bool
	LoopbackMode      uint8
	EnableFEC         bool
	PRBSTxEnable      bool
	PRBSRxEnable      bool
	PRBSPolynomial    uint8
}

// Struct used for retrieving port state
type PortState struct {
	PortNum                     int32
	IfIndex                     int32
	Name                        string
	OperState                   string
	NumUpEvents                 int32
	LastUpEventTime             string
	NumDownEvents               int32
	LastDownEventTime           string
	Pvid                        int32
	IfInOctets                  int64
	IfInUcastPkts               int64
	IfInDiscards                int64
	IfInErrors                  int64
	IfInUnknownProtos           int64
	IfOutOctets                 int64
	IfOutUcastPkts              int64
	IfOutDiscards               int64
	IfOutErrors                 int64
	IfEtherUnderSizePktCnt      int64
	IfEtherOverSizePktCnt       int64
	IfEtherFragments            int64
	IfEtherCRCAlignError        int64
	IfEtherJabber               int64
	IfEtherPkts                 int64
	IfEtherMCPkts               int64
	IfEtherBcastPkts            int64
	IfEtherPkts64OrLessOctets   int64
	IfEtherPkts65To127Octets    int64
	IfEtherPkts128To255Octets   int64
	IfEtherPkts256To511Octets   int64
	IfEtherPkts512To1023Octets  int64
	IfEtherPkts1024To1518Octets int64
	ErrDisableReason            string
	PRBSRxErrCnt                int64
}

// Buffer State for retrieving buffer state
type BufferPortState struct {
	IntfRef            string
	IfIndex            int32
	EgressPort         uint64
	PortBufferPortStat uint64
}

// Buffer state for global buffer stats
type BufferGlobalState struct {
	DeviceId          int32
	BufferStat        uint64
	EgressBufferStat  uint64
	IngressBufferStat uint64
}

// Struct used for ECMP/WCMP group creation
type NextHopGroupMemberInfo struct {
	IpAddr    interface{} //uint32
	NextHopId uint64
	Weight    int
	RifId     int32
}

var CoppType map[int]string = map[int]string{
	int(C.CoppArpUC):    "ArpUC",
	int(C.CoppArpMC):    "ArpMC",
	int(C.CoppBgp):      "BGP",
	int(C.CoppIcmpV4UC): "ICMPv4UC",
	int(C.CoppIcmpV4BC): "ICMPv4BC",
	int(C.CoppStp):      "STP",
	int(C.CoppLacp):     "LACP",
	int(C.CoppBfd):      "BFD",
	int(C.CoppIcmpv6):   "ICMPv6",
	int(C.CoppLldp):     "LLDP",
	//	int(C.CoppSsh):      "SSH",
	//	int(C.CoppHttp):     "HTTP",
}

type CoppStatState struct {
	CoppClass    string
	PeakRate     int32
	BurstRate    int32
	GreenPackets int64
	RedPackets   int64
}

//struct used for ACL creation
//Acl proto type to IP number map
var AclProtoType map[string]int = map[string]int{
	"TCP":    C.COPP_IP_PROTOCOL_IP_NUMBER_TCP,
	"UDP":    C.COPP_IP_PROTOCOL_IP_NUMBER_UDP,
	"ICMPV4": C.COPP_IP_PROTOCOL_IPV4_NUMBER_ICMP,
	"ICMPV6": C.COPP_IP_PROTOCOL_IPV6_NUMBER_ICMP,
	"OSPFV2": C.COPP_IP_PROTOCOL_IP_NUMBER_OSPFV2,
}

type Acl struct {
	AclName      string
	AclType      string
	RuleNameList []string
	Direction    string
}

type AclRule struct {
	RuleName    string
	RuleIndex   int
	SourceMac   net.HardwareAddr
	DestMac     net.HardwareAddr
	SourceIp    []uint8
	DestIp      []uint8
	SourceMask  []uint8
	DestMask    []uint8
	Action      string
	Proto       int
	SrcPort     int32
	DstPort     int32
	L4SrcPort   int32
	L4DstPort   int32
	L4PortMatch string
	L4MinPort   int32
	L4MaxPort   int32
}

type AclRuleState struct {
	SliceIndex int
	Stat       uint64
}

//Format of callback functions for updating DBs in individual resource managers
type ProcessLinkStateChangeCB func(int32, int32, string, string)
type InitPortConfigDBCB func(*PortConfig)
type InitPortStateDBCB func(int32, string)
type UpdatePortStateDBCB func(*PortState)
type UpdateLagDBCB func(int, int, []int32)
type UpdateIPNeighborDBCB func(*PluginIPNeighborInfo)
type UpdateIPv4NeighborDBCB func(uint32, net.HardwareAddr, int32, uint64, int)
type UpdateIPv4RouteDBCB func(uint32, uint32, uint32)
type UpdateIPv4NextHopDBCB func()
type UpdateIPv4NextHopGroupDBCB func()
type MacEntryTableCB func(int, int32, int32, net.HardwareAddr)
type InitBufferPortStateDBCB func(int32, string)
type UpdateBufferPortStateDBCB func(*BufferPortState)
type InitBufferGlobalStateDBCB func(int) int
type UpdateBufferGlobalStateDBCB func(*BufferGlobalState)
type UpdateCoppStatStateDBCB func(*CoppStatState)
type UpdateAclStateDBCB func(*AclRuleState)

func ComputeSetDifference(a, b []int32) (aDiffb []int32) {
	var bMap map[int32]bool = make(map[int32]bool, 0)
	if len(a) == 0 {
		return a
	}
	if len(b) == 0 {
		return a
	}
	for _, elem := range b {
		bMap[elem] = true
	}
	for _, elem := range a {
		if _, ok := bMap[elem]; !ok {
			aDiffb = append(aDiffb, elem)
		}
	}
	return aDiffb
}

type PluginIPInfo struct {
	IfIndex      int32
	IpAddr       []uint32
	MaskLen      int
	IpType       int
	VlanId       int
	IfName       string
	Address      string // if a plugin uses string then it will be copied in here... for e.g. 192.168.1.1/32
	RefCount     int
	InstallRoute int
	IPv6Type     int
	LinkLocalIp  string
	AdminState   int
}

type PluginIPRouteInfo struct {
	PrefixIp      []uint8
	NwAddr        string
	IpType        int //destination ip type (v4/v6)
	Mask          []uint8
	NextHopIp     []uint8
	NextHopIpType int //next hop ip type (v4/v6)
	NextHopIfType int
	Weight        int
	//info to use to communicate with plugins
	Op            uint8
	RouteFlags    uint32
	NextHopId     uint64
	RouterIfIndex int32
	IfName        string
	RifId         int32
	RouteDBKeys   interface{}
	DoneChannel   chan int
	RouteManager  interface{}
	EcmpRoute     bool
	RouteEntity   interface{}
	RetHdlrFunc   func(ipInfo *PluginIPRouteInfo, rMgr interface{}, plugin interface{}, ret int)
}

type PluginIPNeighborInfo struct {
	IfIndex       int32
	L2IfIndex     int32
	IpAddr        []uint32
	NeighborFlags uint32
	NextHopFlags  uint32
	NextHopId     uint64
	VlanId        int
	MacAddr       net.HardwareAddr
	IpType        int
	Address       string
	OperationType int
}

type PluginLagInfo struct {
	IfName     string
	HwId       *int
	HashType   int
	MemberList []int32
}

type PluginUpdateLagInfo struct {
	IfName        string
	HwId          *int
	HashType      int
	OldMemberList []int32
	MemberList    []int32
}

type Inventory struct {
	VendorId   string
	PartNumber string
	RevisionId string
}
