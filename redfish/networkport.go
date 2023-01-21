//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// FlowControl is
type FlowControl string

const (
	// NoneFlowControl No IEEE 802.3x flow control is enabled on this port.
	NoneFlowControl FlowControl = "None"
	// TXFlowControl This station can initiate IEEE 802.3x flow control.
	TXFlowControl FlowControl = "TX"
	// RXFlowControl The link partner can initiate IEEE 802.3x flow control.
	RXFlowControl FlowControl = "RX"
	// TXRXFlowControl This station or the link partner can initiate IEEE 802.3x flow control.
	TXRXFlowControl FlowControl = "TX_RX"
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
)

// LinkStatus is
type LinkStatus string

const (
	// DownLinkStatus The port is enabled but link is down.
	DownLinkStatus LinkStatus = "Down"
	// UpLinkStatus The port is enabled and link is good (up).
	UpLinkStatus LinkStatus = "Up"
	// StartingLinkStatus This link on this interface is starting. A physical link has been established, but the port
	// is not able to transfer data.
	StartingLinkStatus LinkStatus = "Starting"
	// TrainingLinkStatus This physical link on this interface is training.
	TrainingLinkStatus LinkStatus = "Training"
)

// PortConnectionType is
type PortConnectionType string

const (
	// NotConnectedPortConnectionType This port is not connected.
	NotConnectedPortConnectionType PortConnectionType = "NotConnected"
	// NPortPortConnectionType This port connects through an N-port to a switch.
	NPortPortConnectionType PortConnectionType = "NPort"
	// PointToPointPortConnectionType This port connects in a point-to-point configuration.
	PointToPointPortConnectionType PortConnectionType = "PointToPoint"
	// PrivateLoopPortConnectionType This port connects in a private loop configuration.
	PrivateLoopPortConnectionType PortConnectionType = "PrivateLoop"
	// PublicLoopPortConnectionType This port connects in a public configuration.
	PublicLoopPortConnectionType PortConnectionType = "PublicLoop"
	// GenericPortConnectionType This port connection type is a generic fabric port.
	GenericPortConnectionType PortConnectionType = "Generic"
	// ExtenderFabricPortConnectionType This port connection type is an extender fabric port.
	ExtenderFabricPortConnectionType PortConnectionType = "ExtenderFabric"
)

// SupportedEthernetCapabilities is
type SupportedEthernetCapabilities string

const (
	// WakeOnLANSupportedEthernetCapabilities Wake on LAN (WoL) is supported on this port.
	WakeOnLANSupportedEthernetCapabilities SupportedEthernetCapabilities = "WakeOnLAN"
	// EEESupportedEthernetCapabilities IEEE 802.3az Energy-Efficient Ethernet (EEE) is supported on this port.
	EEESupportedEthernetCapabilities SupportedEthernetCapabilities = "EEE"
)

// NetDevFuncMaxBWAlloc shall describe a maximum bandwidth percentage allocation for a network device function
// associated with a port.
type NetDevFuncMaxBWAlloc struct {
	// MaxBWAllocPercent shall contain the maximum bandwidth percentage allocation for the associated network device
	// function.
	MaxBWAllocPercent int
	// NetworkDeviceFunction shall contain a link to a resource of type NetworkDeviceFunction that represents the
	// network device function associated with this bandwidth setting of this network port.
	NetworkDeviceFunction string
}

// UnmarshalJSON unmarshals a NetDevFuncMaxBWAlloc object from the raw JSON.
func (netdevfuncmaxbwalloc *NetDevFuncMaxBWAlloc) UnmarshalJSON(b []byte) error {
	type temp NetDevFuncMaxBWAlloc
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*netdevfuncmaxbwalloc = NetDevFuncMaxBWAlloc(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NetDevFuncMinBWAlloc shall describe a minimum bandwidth percentage allocation for a network device function
// associated with a port.
type NetDevFuncMinBWAlloc struct {
	// MinBWAllocPercent shall contain the minimum bandwidth percentage allocation for the associated network device
	// function. The sum total of all minimum percentages shall not exceed 100.
	MinBWAllocPercent int
	// NetworkDeviceFunction shall contain a link to a resource of type NetworkDeviceFunction that represents the
	// network device function associated with this bandwidth setting of this network port.
	NetworkDeviceFunction string
}

// UnmarshalJSON unmarshals a NetDevFuncMinBWAlloc object from the raw JSON.
func (netdevfuncminbwalloc *NetDevFuncMinBWAlloc) UnmarshalJSON(b []byte) error {
	type temp NetDevFuncMinBWAlloc
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*netdevfuncminbwalloc = NetDevFuncMinBWAlloc(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NetworkPort shall represent a discrete physical port that can connect to a network in a Redfish implementation.
type NetworkPort struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ActiveLinkTechnology shall contain the configured link technology of this port.
	ActiveLinkTechnology LinkNetworkTechnology
	// AssociatedNetworkAddresses shall contain an array of configured network addresses that are associated with this
	// network port, including the programmed address of the lowest numbered network device function, the configured
	// but not active address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedNetworkAddresses []string
	// CurrentLinkSpeedMbps shall contain the current configured link speed of this port.
	CurrentLinkSpeedMbps int
	// Description provides a description of this resource.
	Description string
	// EEEEnabled shall indicate whether IEEE 802.3az Energy-Efficient Ethernet (EEE) is enabled for this network port.
	EEEEnabled bool
	// FCFabricName shall indicate the FC Fabric Name provided by the switch.
	FCFabricName string
	// FCPortConnectionType shall contain the connection type for this port.
	FCPortConnectionType PortConnectionType
	// FlowControlConfiguration shall contain the locally configured 802.3x flow control setting for this network port.
	FlowControlConfiguration FlowControl
	// FlowControlStatus shall contain the 802.3x flow control behavior negotiated with the link partner for this
	// network port (Ethernet-only).
	FlowControlStatus FlowControl
	// LinkStatus shall contain the link status between this port and its link partner.
	LinkStatus LinkStatus
	// MaxFrameSize shall contain the maximum frame size supported by the port.
	MaxFrameSize int
	// NetDevFuncMaxBWAlloc shall contain an array of maximum bandwidth allocation percentages for the network device
	// functions associated with this port.
	NetDevFuncMaxBWAlloc []NetDevFuncMaxBWAlloc
	// NetDevFuncMinBWAlloc shall contain an array of minimum bandwidth percentage allocations for each of the network
	// device functions associated with this port.
	NetDevFuncMinBWAlloc []NetDevFuncMinBWAlloc
	// NumberDiscoveredRemotePorts shall contain the number of ports not on this adapter that this port has discovered.
	NumberDiscoveredRemotePorts int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalPortNumber shall contain the physical port number on the network adapter hardware that this network port
	// corresponds to. This value should match a value visible on the hardware.
	PhysicalPortNumber string
	// PortMaximumMTU shall contain the largest maximum transmission unit (MTU) that can be configured for this network
	// port.
	PortMaximumMTU int
	// SignalDetected shall indicate whether the port has detected enough signal on enough lanes to establish a link.
	SignalDetected bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedEthernetCapabilities shall contain an array of zero or more Ethernet capabilities supported by this
	// port.
	SupportedEthernetCapabilities []SupportedEthernetCapabilities
	// SupportedLinkCapabilities shall describe the static capabilities of the port, irrespective of transient
	// conditions such as cabling, interface module presence, or remote link partner status or configuration.
	SupportedLinkCapabilities []SupportedLinkCapabilities
	// VendorId shall indicate the vendor identification string information as provided by the manufacturer of this
	// port.
	VendorId string
	// WakeOnLANEnabled shall indicate whether Wake on LAN (WoL) is enabled for this network port.
	WakeOnLANEnabled bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a NetworkPort object from the raw JSON.
func (networkport *NetworkPort) UnmarshalJSON(b []byte) error {
	type temp NetworkPort
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*networkport = NetworkPort(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	networkport.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (networkport *NetworkPort) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(NetworkPort)
	original.UnmarshalJSON(networkport.rawData)

	readWriteFields := []string{
		"ActiveLinkTechnology",
		"CurrentLinkSpeedMbps",
		"EEEEnabled",
		"FlowControlConfiguration",
		"WakeOnLANEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(networkport).Elem()

	return networkport.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetNetworkPort will get a NetworkPort instance from the service.
func GetNetworkPort(c common.Client, uri string) (*NetworkPort, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var networkport NetworkPort
	err = json.NewDecoder(resp.Body).Decode(&networkport)
	if err != nil {
		return nil, err
	}

	networkport.SetClient(c)
	return &networkport, nil
}

// ListReferencedNetworkPorts gets the collection of NetworkPort from
// a provided reference.
func ListReferencedNetworkPorts(c common.Client, link string) ([]*NetworkPort, error) {
	var result []*NetworkPort
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, networkportLink := range links.ItemLinks {
		networkport, err := GetNetworkPort(c, networkportLink)
		if err != nil {
			collectionError.Failures[networkportLink] = err
		} else {
			result = append(result, networkport)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
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

// SupportedLinkCapabilities shall describe the static capabilities of an associated port, irrespective of
// transient conditions such as cabling, interface module presence, or remote link partner status or configuration.
type SupportedLinkCapabilities struct {
	// AutoSpeedNegotiation shall indicate whether the port is capable of autonegotiating speed.
	AutoSpeedNegotiation bool
	// CapableLinkSpeedMbps shall contain all of the possible network link speed capabilities of this port.
	CapableLinkSpeedMbps []string
	// LinkNetworkTechnology shall contain a network technology capability of this port.
	LinkNetworkTechnology LinkNetworkTechnology
}

// UnmarshalJSON unmarshals a SupportedLinkCapabilities object from the raw JSON.
func (supportedlinkcapabilities *SupportedLinkCapabilities) UnmarshalJSON(b []byte) error {
	type temp SupportedLinkCapabilities
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*supportedlinkcapabilities = SupportedLinkCapabilities(t.temp)

	// Extract the links to other entities for later

	return nil
}
