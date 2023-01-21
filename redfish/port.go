//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// FiberConnectionType is
type FiberConnectionType string

const (
	// SingleModeFiberConnectionType The connection is using single mode operation.
	SingleModeFiberConnectionType FiberConnectionType = "SingleMode"
	// MultiModeFiberConnectionType The connection is using multi mode operation.
	MultiModeFiberConnectionType FiberConnectionType = "MultiMode"
)

// FlowControl is
type FlowControl string

const (
	// NoneFlowControl No IEEE 802.3x flow control is enabled on this port.
	NoneFlowControl FlowControl = "None"
	// TXFlowControl IEEE 802.3x flow control may be initiated by this station.
	TXFlowControl FlowControl = "TX"
	// RXFlowControl IEEE 802.3x flow control may be initiated by the link partner.
	RXFlowControl FlowControl = "RX"
	// TXRXFlowControl IEEE 802.3x flow control may be initiated by this station or the link partner.
	TXRXFlowControl FlowControl = "TX_RX"
)

// IEEE802IdSubtype is
type IEEE802IdSubtype string

const (
	// ChassisCompIEEE802IdSubtype Chassis component, based in the value of entPhysicalAlias in RFC4133.
	ChassisCompIEEE802IdSubtype IEEE802IdSubtype = "ChassisComp"
	// IfAliasIEEE802IdSubtype Interface alias, based on the ifAlias MIB object.
	IfAliasIEEE802IdSubtype IEEE802IdSubtype = "IfAlias"
	// PortCompIEEE802IdSubtype Port component, based in the value of entPhysicalAlias in RFC4133.
	PortCompIEEE802IdSubtype IEEE802IdSubtype = "PortComp"
	// MacAddrIEEE802IdSubtype MAC address, based on an agent detected unicast source address as defined in IEEE Std.
	// 802.
	MacAddrIEEE802IdSubtype IEEE802IdSubtype = "MacAddr"
	// NetworkAddrIEEE802IdSubtype Network address, based on an agent detected network address.
	NetworkAddrIEEE802IdSubtype IEEE802IdSubtype = "NetworkAddr"
	// IfNameIEEE802IdSubtype Interface name, based on the ifName MIB object.
	IfNameIEEE802IdSubtype IEEE802IdSubtype = "IfName"
	// AgentIdIEEE802IdSubtype Agent circuit ID, based on the agent-local identifier of the circuit as defined in
	// RFC3046.
	AgentIdIEEE802IdSubtype IEEE802IdSubtype = "AgentId"
	// LocalAssignIEEE802IdSubtype Locally assigned, based on a alpha-numeric value locally assigned.
	LocalAssignIEEE802IdSubtype IEEE802IdSubtype = "LocalAssign"
	// NotTransmittedIEEE802IdSubtype No data to be sent to/received from remote partner.
	NotTransmittedIEEE802IdSubtype IEEE802IdSubtype = "NotTransmitted"
)

// LinkNetworkTechnology is
type LinkNetworkTechnology string

const (
	// EthernetLinkNetworkTechnology The port is capable of connecting to an Ethernet network.
	EthernetLinkNetworkTechnology LinkNetworkTechnology = "Ethernet"
	// InfiniBandLinkNetworkTechnology The port is capable of connecting to an InfiniBand network.
	InfiniBandLinkNetworkTechnology LinkNetworkTechnology = "InfiniBand"
	// FibreChannelLinkNetworkTechnology The port is capable of connecting to a Fibre Channel network.
	FibreChannelLinkNetworkTechnology LinkNetworkTechnology = "FibreChannel"
	// GenZLinkNetworkTechnology The port is capable of connecting to a Gen-Z fabric.
	GenZLinkNetworkTechnology LinkNetworkTechnology = "GenZ"
)

// LinkState is
type LinkState string

const (
	// EnabledLinkState This link is enabled.
	EnabledLinkState LinkState = "Enabled"
	// DisabledLinkState This link is disabled.
	DisabledLinkState LinkState = "Disabled"
)

// LinkStatus is
type LinkStatus string

const (
	// LinkUpLinkStatus This link on this interface is up.
	LinkUpLinkStatus LinkStatus = "LinkUp"
	// StartingLinkStatus This link on this interface is starting. A physical link has been established, but the port
	// is not able to transfer data.
	StartingLinkStatus LinkStatus = "Starting"
	// TrainingLinkStatus This physical link on this interface is training.
	TrainingLinkStatus LinkStatus = "Training"
	// LinkDownLinkStatus The link on this interface is down.
	LinkDownLinkStatus LinkStatus = "LinkDown"
	// NoLinkLinkStatus No physical link detected on this interface.
	NoLinkLinkStatus LinkStatus = "NoLink"
)

// MediumType is
type MediumType string

const (
	// CopperMediumType The medium connected is copper.
	CopperMediumType MediumType = "Copper"
	// FiberOpticMediumType The medium connected is fiber optic.
	FiberOpticMediumType MediumType = "FiberOptic"
)

// PortConnectionType is
type PortConnectionType string

const (
	// NotConnectedPortConnectionType This port is not connected.
	NotConnectedPortConnectionType PortConnectionType = "NotConnected"
	// NPortPortConnectionType This port connects through an N-Port to a switch.
	NPortPortConnectionType PortConnectionType = "NPort"
	// PointToPointPortConnectionType This port connects in a Point-to-point configuration.
	PointToPointPortConnectionType PortConnectionType = "PointToPoint"
	// PrivateLoopPortConnectionType This port connects in a private loop configuration.
	PrivateLoopPortConnectionType PortConnectionType = "PrivateLoop"
	// PublicLoopPortConnectionType This port connects in a public configuration.
	PublicLoopPortConnectionType PortConnectionType = "PublicLoop"
	// GenericPortConnectionType This port connection type is a generic fabric port.
	GenericPortConnectionType PortConnectionType = "Generic"
	// ExtenderFabricPortConnectionType This port connection type is an extender fabric port.
	ExtenderFabricPortConnectionType PortConnectionType = "ExtenderFabric"
	// FPortPortConnectionType This port connection type is a fabric port.
	FPortPortConnectionType PortConnectionType = "FPort"
	// EPortPortConnectionType This port connection type is an extender fabric port.
	EPortPortConnectionType PortConnectionType = "EPort"
	// TEPortPortConnectionType This port connection type is an trunking extender fabric port.
	TEPortPortConnectionType PortConnectionType = "TEPort"
	// NPPortPortConnectionType This port connection type is a proxy N port for N-Port virtualization.
	NPPortPortConnectionType PortConnectionType = "NPPort"
	// GPortPortConnectionType This port connection type is a generic fabric port.
	GPortPortConnectionType PortConnectionType = "GPort"
	// NLPortPortConnectionType This port connects in a node loop configuration.
	NLPortPortConnectionType PortConnectionType = "NLPort"
	// FLPortPortConnectionType This port connects in a fabric loop configuration.
	FLPortPortConnectionType PortConnectionType = "FLPort"
	// EXPortPortConnectionType This port connection type is an external fabric port.
	EXPortPortConnectionType PortConnectionType = "EXPort"
	// UPortPortConnectionType This port connection type is unassigned.
	UPortPortConnectionType PortConnectionType = "UPort"
	// DPortPortConnectionType This port connection type is a diagnostic port.
	DPortPortConnectionType PortConnectionType = "DPort"
)

// PortMedium is
type PortMedium string

const (
	// ElectricalPortMedium This port has an electrical cable connection.
	ElectricalPortMedium PortMedium = "Electrical"
	// OpticalPortMedium This port has an optical cable connection.
	OpticalPortMedium PortMedium = "Optical"
)

// PortType is
type PortType string

const (
	// UpstreamPortPortType This port connects to a host device.
	UpstreamPortPortType PortType = "UpstreamPort"
	// DownstreamPortPortType This port connects to a target device.
	DownstreamPortPortType PortType = "DownstreamPort"
	// InterswitchPortPortType This port connects to another switch.
	InterswitchPortPortType PortType = "InterswitchPort"
	// ManagementPortPortType This port connects to a switch manager.
	ManagementPortPortType PortType = "ManagementPort"
	// BidirectionalPortPortType This port connects to any type of device.
	BidirectionalPortPortType PortType = "BidirectionalPort"
	// UnconfiguredPortPortType This port has not yet been configured.
	UnconfiguredPortPortType PortType = "UnconfiguredPort"
)

// SFPType is
type SFPType string

const (
	// SFPSFPType The SFP conforms to the SFF Specification for SFP.
	SFPSFPType SFPType = "SFP"
	// SFPPlusSFPType The SFP conforms to the SFF Specification for SFP+.
	SFPPlusSFPType SFPType = "SFPPlus"
	// SFP28SFPType The SFP conforms to the SFF Specification for SFP+ and IEEE 802.3by Specification.
	SFP28SFPType SFPType = "SFP28"
	// cSFPSFPType The SFP conforms to the CSFP MSA Specification.
	cSFPSFPType SFPType = "cSFP"
	// SFPDDSFPType The SFP conforms to the SFP-DD MSA Specification.
	SFPDDSFPType SFPType = "SFPDD"
	// QSFPSFPType The SFP conforms to the SFF Specification for QSFP.
	QSFPSFPType SFPType = "QSFP"
	// QSFPPlusSFPType The SFP conforms to the SFF Specification for QSFP+.
	QSFPPlusSFPType SFPType = "QSFPPlus"
	// QSFP14SFPType The SFP conforms to the SFF Specification for QSFP14.
	QSFP14SFPType SFPType = "QSFP14"
	// QSFP28SFPType The SFP conforms to the SFF Specification for QSFP28.
	QSFP28SFPType SFPType = "QSFP28"
	// QSFP56SFPType The SFP conforms to the SFF Specification for QSFP56.
	QSFP56SFPType SFPType = "QSFP56"
	// MiniSASHDSFPType The SFP conforms to the SFF Specification SFF-8644.
	MiniSASHDSFPType SFPType = "MiniSASHD"
)

// SupportedEthernetCapabilities is
type SupportedEthernetCapabilities string

const (
	// WakeOnLANSupportedEthernetCapabilities Wake on LAN (WoL) is supported on this port.
	WakeOnLANSupportedEthernetCapabilities SupportedEthernetCapabilities = "WakeOnLAN"
	// EEESupportedEthernetCapabilities IEEE 802.3az Energy-Efficient Ethernet (EEE) is supported on this port.
	EEESupportedEthernetCapabilities SupportedEthernetCapabilities = "EEE"
)

// ConfiguredNetworkLink shall contain a set of link settings that a port is configured to use for autonegotiation.
type ConfiguredNetworkLink struct {
	// ConfiguredLinkSpeedGbps shall contain the network link speed per lane this port is configured to allow for
	// autonegotiation purposes. This value includes overhead associated with the protocol.
	ConfiguredLinkSpeedGbps float64
	// ConfiguredWidth shall contain the network link width this port is configured to use for autonegotiation
	// purposes.
	ConfiguredWidth int
}

// UnmarshalJSON unmarshals a ConfiguredNetworkLink object from the raw JSON.
func (configurednetworklink *ConfiguredNetworkLink) UnmarshalJSON(b []byte) error {
	type temp ConfiguredNetworkLink
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*configurednetworklink = ConfiguredNetworkLink(t.temp)

	// Extract the links to other entities for later

	return nil
}

// EthernetProperties shall contain Ethernet-specific properties for a port.
type EthernetProperties struct {
	// AssociatedMACAddresses shall contain an array of configured MAC addresses that are associated with this network
	// port, including the programmed address of the lowest numbered network device function, the configured but not
	// active address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedMACAddresses []string
	// EEEEnabled shall indicate whether IEEE 802.3az Energy-Efficient Ethernet (EEE) is enabled on this port.
	EEEEnabled bool
	// FlowControlConfiguration shall contain the locally configured 802.3x flow control setting for this port.
	FlowControlConfiguration FlowControl
	// FlowControlStatus shall contain the 802.3x flow control behavior negotiated with the link partner for this port.
	FlowControlStatus FlowControl
	// LLDPEnabled shall contain the state indicating whether to enable LLDP for a port. If LLDP is disabled at the
	// adapter level, this property shall be ignored.
	LLDPEnabled string
	// LLDPReceive shall contain the LLDP data being received on this link.
	LLDPReceive LLDPReceive
	// LLDPTransmit shall contain the LLDP data being transmitted on this link.
	LLDPTransmit LLDPTransmit
	// WakeOnLANEnabled shall indicate whether Wake on LAN (WoL) is enabled on this port.
	WakeOnLANEnabled bool
}

// UnmarshalJSON unmarshals a EthernetProperties object from the raw JSON.
func (ethernetproperties *EthernetProperties) UnmarshalJSON(b []byte) error {
	type temp EthernetProperties
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ethernetproperties = EthernetProperties(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FibreChannelProperties shall contain Fibre Channel-specific properties for a port.
type FibreChannelProperties struct {
	// AssociatedWorldWideNames shall contain an array of configured World Wide Names (WWN) that are associated with
	// this network port, including the programmed address of the lowest numbered network device function, the
	// configured but not active address if applicable, the address for hardware port teaming, or other network
	// addresses.
	AssociatedWorldWideNames []string
	// FabricName shall indicate the Fibre Channel Fabric Name provided by the switch.
	FabricName string
	// NumberDiscoveredRemotePorts shall contain the number of ports not on this associated device that this port has
	// discovered.
	NumberDiscoveredRemotePorts int
	// PortConnectionType shall contain the connection type for this port.
	PortConnectionType PortConnectionType
}

// UnmarshalJSON unmarshals a FibreChannelProperties object from the raw JSON.
func (fibrechannelproperties *FibreChannelProperties) UnmarshalJSON(b []byte) error {
	type temp FibreChannelProperties
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fibrechannelproperties = FibreChannelProperties(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FunctionMaxBandwidth shall describe a maximum bandwidth percentage allocation for a network device function
// associated with a port.
type FunctionMaxBandwidth struct {
	// AllocationPercent shall contain the maximum bandwidth percentage allocation for the associated network device
	// function.
	AllocationPercent int
	// NetworkDeviceFunction shall contain a link to a resource of type NetworkDeviceFunction that represents the
	// network device function associated with this bandwidth setting of this network port.
	NetworkDeviceFunction string
}

// UnmarshalJSON unmarshals a FunctionMaxBandwidth object from the raw JSON.
func (functionmaxbandwidth *FunctionMaxBandwidth) UnmarshalJSON(b []byte) error {
	type temp FunctionMaxBandwidth
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*functionmaxbandwidth = FunctionMaxBandwidth(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FunctionMinBandwidth shall describe a minimum bandwidth percentage allocation for a network device function
// associated with a port.
type FunctionMinBandwidth struct {
	// AllocationPercent shall contain the minimum bandwidth percentage allocation for the associated network device
	// function. The sum total of all minimum percentages shall not exceed 100.
	AllocationPercent int
	// NetworkDeviceFunction shall contain a link to a resource of type NetworkDeviceFunction that represents the
	// network device function associated with this bandwidth setting of this network port.
	NetworkDeviceFunction string
}

// UnmarshalJSON unmarshals a FunctionMinBandwidth object from the raw JSON.
func (functionminbandwidth *FunctionMinBandwidth) UnmarshalJSON(b []byte) error {
	type temp FunctionMinBandwidth
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*functionminbandwidth = FunctionMinBandwidth(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GenZ shall contain Gen-Z specific port properties.
type GenZ struct {
	// LPRT shall contain a link to a resource collection of type RouteEntryCollection, and shall represent the Gen-Z
	// Core Specification-defined Linear Packet Relay Table for this port.
	LPRT string
	// MPRT shall contain a link to a resource collection of type RouteEntryCollection, and shall represent the Gen-Z
	// Core Specification-defined Multi-subnet Packet Relay Table for this port.
	MPRT string
	// VCAT shall contain a link to a resource collection of type VCATEntryCollection.
	VCAT string
}

// UnmarshalJSON unmarshals a GenZ object from the raw JSON.
func (genz *GenZ) UnmarshalJSON(b []byte) error {
	type temp GenZ
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*genz = GenZ(t.temp)

	// Extract the links to other entities for later

	return nil
}

// InfiniBandProperties shall contain InfiniBand-specific properties for a port.
type InfiniBandProperties struct {
	// AssociatedNodeGUIDs shall contain an array of configured node GUIDs that are associated with this network port,
	// including the programmed address of the lowest numbered network device function, the configured but not active
	// address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedNodeGUIDs []string
	// AssociatedPortGUIDs shall contain an array of configured port GUIDs that are associated with this network port,
	// including the programmed address of the lowest numbered network device function, the configured but not active
	// address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedPortGUIDs []string
	// AssociatedSystemGUIDs shall contain an array of configured system GUIDs that are associated with this network
	// port, including the programmed address of the lowest numbered network device function, the configured but not
	// active address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedSystemGUIDs []string
}

// UnmarshalJSON unmarshals a InfiniBandProperties object from the raw JSON.
func (infinibandproperties *InfiniBandProperties) UnmarshalJSON(b []byte) error {
	type temp InfiniBandProperties
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*infinibandproperties = InfiniBandProperties(t.temp)

	// Extract the links to other entities for later

	return nil
}

// LLDPReceive shall contain the LLDP data from the remote partner across this link.
type LLDPReceive struct {
	// ChassisId shall contain the chassis ID received from the remote partner across this link. If no such chassis ID
	// has been received, this property should not be present.
	ChassisId string
	// ChassisIdSubtype shall contain the IEEE 802.1AB-2009 chassis ID subtype received from the remote partner across
	// this link. If no such chassis ID subtype has been received, this property should not be present.
	ChassisIdSubtype IEEE802IdSubtype
	// ManagementAddressIPv4 shall contain the IPv4 management address received from the remote partner across this
	// link. If no such management address has been received, this property should not be present.
	ManagementAddressIPv4 string
	// ManagementAddressIPv6 shall contain the IPv6 management address received from the remote partner across this
	// link. If no such management address has been received, this property should not be present.
	ManagementAddressIPv6 string
	// ManagementAddressMAC shall contain the management MAC address received from the remote partner across this link.
	// If no such management address has been received, this property should not be present.
	ManagementAddressMAC string
	// ManagementVlanId shall contain the management VLAN ID received from the remote partner across this link. If no
	// such management VLAN ID has been received, this property should not be present.
	ManagementVlanId int
	// PortId shall contain a colon delimited string of hexadecimal octets identifying the port received from the
	// remote partner across this link. If no such port ID has been received, this property should not be present.
	PortId string
	// PortIdSubtype shall contain the port ID subtype from IEEE 802.1AB-2009 Table 8-3 received from the remote
	// partner across this link. If no such port ID subtype has been received, this property should not be present.
	PortIdSubtype IEEE802IdSubtype
}

// UnmarshalJSON unmarshals a LLDPReceive object from the raw JSON.
func (lldpreceive *LLDPReceive) UnmarshalJSON(b []byte) error {
	type temp LLDPReceive
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*lldpreceive = LLDPReceive(t.temp)

	// Extract the links to other entities for later

	return nil
}

// LLDPTransmit shall contain the LLDP data to be transmitted from this endpoint.
type LLDPTransmit struct {
	// ChassisId shall contain the chassis ID to be transmitted from this endpoint. If no such chassis ID is to be
	// transmitted, this value shall be an empty string.
	ChassisId string
	// ChassisIdSubtype shall contain the IEEE 802.1AB-2009 chassis ID subtype to be transmitted from this endpoint. If
	// no such chassis ID subtype is to be transmitted, this value shall be 'NotTransmitted'.
	ChassisIdSubtype IEEE802IdSubtype
	// ManagementAddressIPv4 shall contain the IPv4 management address to be transmitted from this endpoint. If no such
	// management address is to be transmitted, this value shall be an empty string.
	ManagementAddressIPv4 string
	// ManagementAddressIPv6 shall contain the IPv6 management address to be transmitted from this endpoint. If no such
	// management address is to be transmitted, this value shall be an empty string.
	ManagementAddressIPv6 string
	// ManagementAddressMAC shall contain the management MAC address to be transmitted from this endpoint. If no such
	// management address is to be transmitted, this value shall be an empty string.
	ManagementAddressMAC string
	// ManagementVlanId shall contain the management VLAN ID to be transmitted from this endpoint. If no such port ID
	// is to be transmitted, this value shall be '4095'.
	ManagementVlanId int
	// PortId shall contain a colon delimited string of hexadecimal octets identifying the port for an LLDP endpoint.
	// If no such port ID is to be transmitted, this value shall be an empty string.
	PortId string
	// PortIdSubtype shall contain the port ID subtype from IEEE 802.1AB-2009 Table 8-3 to be transmitted from this
	// endpoint. If no such port ID subtype is to be transmitted, this value shall be 'NotTransmitted'.
	PortIdSubtype IEEE802IdSubtype
}

// UnmarshalJSON unmarshals a LLDPTransmit object from the raw JSON.
func (lldptransmit *LLDPTransmit) UnmarshalJSON(b []byte) error {
	type temp LLDPTransmit
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*lldptransmit = LLDPTransmit(t.temp)

	// Extract the links to other entities for later

	return nil
}

// LinkConfiguration shall contain properties of the link for which port has been configured.
type LinkConfiguration struct {
	// AutoSpeedNegotiationCapable shall indicate whether the port is capable of autonegotiating speed.
	AutoSpeedNegotiationCapable bool
	// AutoSpeedNegotiationEnabled shall indicate whether the port is configured to autonegotiate speed.
	AutoSpeedNegotiationEnabled bool
	// CapableLinkSpeedGbps shall contain all of the possible network link speed capabilities of this port. This value
	// includes overhead associated with the protocol.
	CapableLinkSpeedGbps []string
	// ConfiguredNetworkLinks shall contain the set of link speed and width pairs to which this port is restricted for
	// autonegotiation purposes.
	ConfiguredNetworkLinks []ConfiguredNetworkLink
}

// UnmarshalJSON unmarshals a LinkConfiguration object from the raw JSON.
func (linkconfiguration *LinkConfiguration) UnmarshalJSON(b []byte) error {
	type temp LinkConfiguration
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*linkconfiguration = LinkConfiguration(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// AssociatedEndpoints shall contain an array of links to resources of type Endpoint that represent the endpoints
	// to which this port is connected.
	AssociatedEndpoints []Endpoint
	// AssociatedEndpoints@odata.count
	AssociatedEndpointsCount int `json:"AssociatedEndpoints@odata.count"`
	// Cables shall contain an array of links to resources of type Cable that represent the cables connected to this
	// port.
	Cables []Cable
	// Cables@odata.count
	CablesCount int `json:"Cables@odata.count"`
	// ConnectedPorts shall contain an array of links to resources of type Port that represent the remote device ports
	// to which this port is connected.
	ConnectedPorts []Port
	// ConnectedPorts@odata.count
	ConnectedPortsCount int `json:"ConnectedPorts@odata.count"`
	// ConnectedSwitchPorts shall contain an array of links to resources of type Port that represent the switch ports
	// to which this port is connected.
	ConnectedSwitchPorts []Port
	// ConnectedSwitchPorts@odata.count
	ConnectedSwitchPortsCount int `json:"ConnectedSwitchPorts@odata.count"`
	// ConnectedSwitches shall contain an array of links to resources of type Switch that represent the switches to
	// which this port is connected.
	ConnectedSwitches []Switch
	// ConnectedSwitches@odata.count
	ConnectedSwitchesCount int `json:"ConnectedSwitches@odata.count"`
	// EthernetInterfaces shall contain an array of links to resources of type EthernetInterface that represent the
	// Ethernet interfaces this port provides. This property shall not include Ethernet interfaces that are not
	// directly associated to a physical port.
	EthernetInterfaces []EthernetInterface
	// EthernetInterfaces@odata.count
	EthernetInterfacesCount int `json:"EthernetInterfaces@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a Links object from the raw JSON.
func (links *Links) UnmarshalJSON(b []byte) error {
	type temp Links
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*links = Links(t.temp)

	// Extract the links to other entities for later

	return nil
}

// OemActions shall contain the available OEM-specific actions for this resource.
type OemActions struct {
}

// UnmarshalJSON unmarshals a OemActions object from the raw JSON.
func (oemactions *OemActions) UnmarshalJSON(b []byte) error {
	type temp OemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*oemactions = OemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Port This resource contains a simple port for a Redfish implementation.
type Port struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ActiveWidth shall contain the number of active lanes for this interface.
	ActiveWidth string
	// CapableProtocolVersions shall contain the protocol versions capable of being sent over this port. This property
	// should only be used for protocols where the version and not the speed is of primary interest such as USB,
	// DisplayPort, or HDMI.
	CapableProtocolVersions []string
	// CurrentProtocolVersion shall contain the protocol version being sent over this port. This property should only
	// be used for protocols where the version and not the speed is of primary interest such as USB, DisplayPort, or
	// HDMI.
	CurrentProtocolVersion string
	// CurrentSpeedGbps shall contain the speed of this port currently negotiated and running. This value includes
	// overhead associated with the protocol.
	CurrentSpeedGbps float64
	// Description provides a description of this resource.
	Description string
	// Enabled shall indicate if this port is enabled. Disabling a port will disconnect any devices only connected to
	// the system through this port.
	Enabled string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that represents the environment
	// metrics for this port or any attached small form-factor pluggable (SFP) device.
	EnvironmentMetrics string
	// Ethernet shall contain Ethernet-specific properties of the port.
	Ethernet EthernetProperties
	// FibreChannel shall contain Fibre Channel-specific properties of the port.
	FibreChannel FibreChannelProperties
	// FunctionMaxBandwidth shall contain an array of maximum bandwidth allocation percentages for the functions
	// associated with this port.
	FunctionMaxBandwidth []FunctionMaxBandwidth
	// FunctionMinBandwidth shall contain an array of minimum bandwidth percentage allocations for each of the
	// functions associated with this port.
	FunctionMinBandwidth []FunctionMinBandwidth
	// GenZ shall contain Gen-Z specific properties for this interface.
	GenZ string
	// InfiniBand shall contain InfiniBand-specific properties of the port.
	InfiniBand InfiniBandProperties
	// InterfaceEnabled shall indicate whether the interface is enabled.
	InterfaceEnabled bool
	// LinkConfiguration shall contain the static capabilities and configuration settings of the port.
	LinkConfiguration []LinkConfiguration
	// LinkNetworkTechnology shall contain a network technology capability of this port.
	LinkNetworkTechnology LinkNetworkTechnology
	// LinkState shall contain the desired link state for this interface.
	LinkState string
	// LinkStatus shall contain the link status for this interface.
	LinkStatus string
	// LinkTransitionIndicator shall contain the number of link state transitions for this interface.
	LinkTransitionIndicator string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated port.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// MaxFrameSize shall contain the maximum frame size supported by the port.
	MaxFrameSize int
	// MaxSpeedGbps shall contain the maximum speed of which this port is capable of being configured. If capable of
	// autonegotiation, the system shall attempt to negotiate at the maximum speed set. This value includes overhead
	// associated with the protocol.
	MaxSpeedGbps float64
	// Metrics shall contain a link to the metrics associated with this port.
	Metrics PortMetrics
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PortId shall contain the name of the port as indicated on the device containing the port.
	PortId string
	// PortMedium shall contain the physical connection medium for this port.
	PortMedium PortMedium
	// PortProtocol shall contain the protocol being sent over this port.
	PortProtocol Protocol
	// PortType shall contain the port type for this port.
	PortType PortType
	// SFP shall contain data about the small form-factor pluggable (SFP) device currently occupying this port.
	SFP SFP
	// SignalDetected shall indicate whether a signal that is appropriate for this link technology is detected for this
	// port.
	SignalDetected bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Width shall contain the number of physical transport links that this port contains.
	Width int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Port object from the raw JSON.
func (port *Port) UnmarshalJSON(b []byte) error {
	type temp Port
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*port = Port(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	port.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (port *Port) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Port)
	original.UnmarshalJSON(port.rawData)

	readWriteFields := []string{
		"Enabled",
		"InterfaceEnabled",
		"LinkState",
		"LinkTransitionIndicator",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(port).Elem()

	return port.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPort will get a Port instance from the service.
func GetPort(c common.Client, uri string) (*Port, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var port Port
	err = json.NewDecoder(resp.Body).Decode(&port)
	if err != nil {
		return nil, err
	}

	port.SetClient(c)
	return &port, nil
}

// ListReferencedPorts gets the collection of Port from
// a provided reference.
func ListReferencedPorts(c common.Client, link string) ([]*Port, error) {
	var result []*Port
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, portLink := range links.ItemLinks {
		port, err := GetPort(c, portLink)
		if err != nil {
			collectionError.Failures[portLink] = err
		} else {
			result = append(result, port)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// SFP shall describe a small form-factor pluggable (SFP) device attached to a port.
type SFP struct {
	// FiberConnectionType shall contain the fiber connection type used by the SFP.
	FiberConnectionType FiberConnectionType
	// Manufacturer shall contain the name of the organization responsible for producing the SFP. This organization may
	// be the entity from which the SFP is purchased, but this is not necessarily true.
	Manufacturer string
	// MediumType shall contain the medium type used by the SFP.
	MediumType MediumType
	// PartNumber shall contain the manufacturer-provided part number for the SFP.
	PartNumber string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the SFP.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedSFPTypes shall contain an array of SFP device types supported by this port.
	SupportedSFPTypes []SFPType
	// Type shall contain the SFP device type currently attached to this port.
	Type SFPType
}

// UnmarshalJSON unmarshals a SFP object from the raw JSON.
func (sfp *SFP) UnmarshalJSON(b []byte) error {
	type temp SFP
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sfp = SFP(t.temp)

	// Extract the links to other entities for later

	return nil
}
