//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AuthenticationMethod is
type AuthenticationMethod string

const (
	// NoneAuthenticationMethod No iSCSI authentication is used.
	NoneAuthenticationMethod AuthenticationMethod = "None"
	// CHAPAuthenticationMethod iSCSI Challenge Handshake Authentication Protocol (CHAP) authentication is used.
	CHAPAuthenticationMethod AuthenticationMethod = "CHAP"
	// MutualCHAPAuthenticationMethod iSCSI Mutual Challenge Handshake Authentication Protocol (CHAP) authentication is
	// used.
	MutualCHAPAuthenticationMethod AuthenticationMethod = "MutualCHAP"
)

// BootMode is
type BootMode string

const (
	// DisabledBootMode Do not indicate to UEFI/BIOS that this device is bootable.
	DisabledBootMode BootMode = "Disabled"
	// PXEBootMode Boot this device by using the embedded PXE support. Only applicable if the NetDevFuncType is
	// 'Ethernet' or 'InfiniBand'.
	PXEBootMode BootMode = "PXE"
	// iSCSIBootMode Boot this device by using the embedded iSCSI boot support and configuration. Only applicable if
	// the NetDevFuncType is 'iSCSI' or 'Ethernet'.
	iSCSIBootMode BootMode = "iSCSI"
	// FibreChannelBootMode Boot this device by using the embedded Fibre Channel support and configuration. Only
	// applicable if the NetDevFuncType is 'FibreChannel'.
	FibreChannelBootMode BootMode = "FibreChannel"
	// FibreChannelOverEthernetBootMode Boot this device by using the embedded Fibre Channel over Ethernet (FCoE) boot
	// support and configuration. Only applicable if the NetDevFuncType is 'FibreChannelOverEthernet'.
	FibreChannelOverEthernetBootMode BootMode = "FibreChannelOverEthernet"
	// HTTPBootMode Boot this device by using the embedded HTTP/HTTPS support. Only applicable if the NetDevFuncType is
	// 'Ethernet'.
	HTTPBootMode BootMode = "HTTP"
)

// DataDirection is
type DataDirection string

const (
	// NoneDataDirection Indicates that this limit not enforced.
	NoneDataDirection DataDirection = "None"
	// IngressDataDirection Indicates that this limit is enforced on packets and bytes received by the network device
	// function.
	IngressDataDirection DataDirection = "Ingress"
	// EgressDataDirection Indicates that this limit is enforced on packets and bytes transmitted by the network device
	// function.
	EgressDataDirection DataDirection = "Egress"
)

// IPAddressType is
type IPAddressType string

const (
	// IPv4IPAddressType IPv4 addressing is used for all IP-fields in this object.
	IPv4IPAddressType IPAddressType = "IPv4"
	// IPv6IPAddressType IPv6 addressing is used for all IP-fields in this object.
	IPv6IPAddressType IPAddressType = "IPv6"
)

// NetworkDeviceTechnology is
type NetworkDeviceTechnology string

const (
	// DisabledNetworkDeviceTechnology Neither enumerated nor visible to the operating system.
	DisabledNetworkDeviceTechnology NetworkDeviceTechnology = "Disabled"
	// EthernetNetworkDeviceTechnology Appears to the operating system as an Ethernet device.
	EthernetNetworkDeviceTechnology NetworkDeviceTechnology = "Ethernet"
	// FibreChannelNetworkDeviceTechnology Appears to the operating system as a Fibre Channel device.
	FibreChannelNetworkDeviceTechnology NetworkDeviceTechnology = "FibreChannel"
	// iSCSINetworkDeviceTechnology Appears to the operating system as an iSCSI device.
	iSCSINetworkDeviceTechnology NetworkDeviceTechnology = "iSCSI"
	// FibreChannelOverEthernetNetworkDeviceTechnology Appears to the operating system as an FCoE device.
	FibreChannelOverEthernetNetworkDeviceTechnology NetworkDeviceTechnology = "FibreChannelOverEthernet"
	// InfiniBandNetworkDeviceTechnology Appears to the operating system as an InfiniBand device.
	InfiniBandNetworkDeviceTechnology NetworkDeviceTechnology = "InfiniBand"
)

// WWNSource is
type WWNSource string

const (
	// ConfiguredLocallyWWNSource The set of FC/FCoE boot targets was applied locally through API or UI.
	ConfiguredLocallyWWNSource WWNSource = "ConfiguredLocally"
	// ProvidedByFabricWWNSource The set of FC/FCoE boot targets was applied by the Fibre Channel fabric.
	ProvidedByFabricWWNSource WWNSource = "ProvidedByFabric"
)

// BootTargets shall describe a Fibre Channel boot target configured for a network device function.
type BootTargets struct {
	// BootPriority shall contain the relative priority for this entry in the boot targets array. Lower numbers shall
	// represent higher priority, with zero being the highest priority. The BootPriority shall be unique for all
	// entries of the BootTargets array.
	BootPriority int
	// LUNID shall contain the logical unit number (LUN) ID from which to boot on the device to which the corresponding
	// WWPN refers.
	LUNID string
	// WWPN shall contain World Wide Port Name (WWPN) from which to boot.
	WWPN string
}

// UnmarshalJSON unmarshals a BootTargets object from the raw JSON.
func (boottargets *BootTargets) UnmarshalJSON(b []byte) error {
	type temp BootTargets
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*boottargets = BootTargets(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Ethernet shall describe the Ethernet capabilities, status, and configuration values for a network device
// function.
type Ethernet struct {
	// EthernetInterfaces shall contain a link to a collection of type EthernetInterfaceCollection that represent the
	// Ethernet interfaces present on this network device function. This property shall not be present if this network
	// device function is not referenced by a NetworkInterface resource.
	EthernetInterfaces EthernetInterfaceCollection
	// MACAddress shall contain the effective current MAC address of this network device function. If an assignable MAC
	// address is not supported, this is a read-only alias of the PermanentMACAddress.
	MACAddress string
	// MTUSize The maximum transmission unit (MTU) configured for this network device function. This value serves as a
	// default for the OS driver when booting. The value only takes effect on boot.
	MTUSize int
	// MTUSizeMaximum shall contain the largest maximum transmission unit (MTU) size supported for this network device
	// function.
	MTUSizeMaximum int
	// PermanentMACAddress shall contain the permanent MAC Address of this function. Typically, this value is
	// programmed during manufacturing. This address is not assignable.
	PermanentMACAddress string
	// VLAN shall contain the VLAN for this interface. If this interface supports more than one VLAN, the VLAN property
	// shall not be present and the VLANs property shall be present instead.
	VLAN string
}

// UnmarshalJSON unmarshals a Ethernet object from the raw JSON.
func (ethernet *Ethernet) UnmarshalJSON(b []byte) error {
	type temp Ethernet
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ethernet = Ethernet(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FibreChannel shall describe the Fibre Channel capabilities, status, and configuration values for a network
// device function.
type FibreChannel struct {
	// AllowFIPVLANDiscovery shall indicate whether the FIP VLAN Discovery Protocol determines the FCoE VLAN ID
	// selected by the network device function for the FCoE connection. If 'true' and the FIP VLAN discovery succeeds,
	// the FCoEActiveVLANId property shall reflect the FCoE VLAN ID to use for all FCoE traffic. If 'false' or if the
	// FIP VLAN Discovery protocol fails, the FCoELocalVLANId shall be used for all FCoE traffic and the
	// FCoEActiveVLANId shall reflect the FCoELocalVLANId.
	AllowFIPVLANDiscovery bool
	// BootTargets shall contain an array of Fibre Channel boot targets configured for this network device function.
	BootTargets []BootTargets
	// FCoEActiveVLANId shall contain 'null' or a VLAN ID currently being used for FCoE traffic. When the FCoE link is
	// down this value shall be null. When the FCoE link is up this value shall be either the FCoELocalVLANId property
	// or a VLAN discovered through the FIP protocol.
	FCoEActiveVLANId int
	// FCoELocalVLANId shall contain the VLAN ID configured locally by setting this property. This value shall be used
	// for FCoE traffic to this network device function during boot unless AllowFIPVLANDiscovery is 'true' and a valid
	// FCoE VLAN ID is found through the FIP VLAN Discovery Protocol.
	FCoELocalVLANId int
	// FibreChannelId shall indicate the Fibre Channel ID that the switch assigns for this interface.
	FibreChannelId string
	// PermanentWWNN shall contain the permanent World Wide Node Name (WWNN) of this function. Typically, this value is
	// programmed during manufacturing. This address is not assignable.
	PermanentWWNN string
	// PermanentWWPN shall contain the permanent World Wide Port Name (WWPN) of this function. Typically, this value is
	// programmed during manufacturing. This address is not assignable.
	PermanentWWPN string
	// WWNN shall contain the effective current World Wide Node Name (WWNN) of this function. If an assignable WWNN is
	// not supported, this is a read-only alias of the permanent WWNN.
	WWNN string
	// WWNSource shall contain the configuration source of the World Wide Name (WWN) for this World Wide Node Name
	// (WWNN) and World Wide Port Name (WWPN) connection.
	WWNSource WWNSource
	// WWPN shall contain the effective current World Wide Port Name (WWPN) of this function. If an assignable WWPN is
	// not supported, this is a read-only alias of the permanent WWPN.
	WWPN string
}

// UnmarshalJSON unmarshals a FibreChannel object from the raw JSON.
func (fibrechannel *FibreChannel) UnmarshalJSON(b []byte) error {
	type temp FibreChannel
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fibrechannel = FibreChannel(t.temp)

	// Extract the links to other entities for later

	return nil
}

// HTTPBoot shall describe the HTTP and HTTPS boot capabilities, status, and configuration values for a network
// device function.
type HTTPBoot struct {
	// BootMediaURI shall contain the URI of the boot media loaded with this network device function. An empty string
	// shall indicate no boot media is configured. All other values shall begin with 'http://' or 'https://'.
	BootMediaURI string
}

// UnmarshalJSON unmarshals a HTTPBoot object from the raw JSON.
func (httpboot *HTTPBoot) UnmarshalJSON(b []byte) error {
	type temp HTTPBoot
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*httpboot = HTTPBoot(t.temp)

	// Extract the links to other entities for later

	return nil
}

// InfiniBand shall describe the InfiniBand capabilities, status, and configuration values for a network device
// function.
type InfiniBand struct {
	// MTUSize The maximum transmission unit (MTU) configured for this network device function.
	MTUSize int
	// NodeGUID shall contain the effective current node GUID of this virtual port of this network device function. If
	// an assignable node GUID is not supported, this is a read-only alias of the PermanentNodeGUID.
	NodeGUID string
	// PermanentNodeGUID shall contain the permanent node GUID of this network device function. Typically, this value
	// is programmed during manufacturing. This address is not assignable.
	PermanentNodeGUID string
	// PermanentPortGUID shall contain the permanent port GUID of this network device function. Typically, this value
	// is programmed during manufacturing. This address is not assignable.
	PermanentPortGUID string
	// PermanentSystemGUID shall contain the permanent system GUID of this network device function. Typically, this
	// value is programmed during manufacturing. This address is not assignable.
	PermanentSystemGUID string
	// PortGUID shall contain the effective current virtual port GUID of this network device function. If an assignable
	// port GUID is not supported, this is a read-only alias of the PermanentPortGUID.
	PortGUID string
	// SupportedMTUSizes shall contain an array of the maximum transmission unit (MTU) sizes supported for this network
	// device function.
	SupportedMTUSizes []string
	// SystemGUID shall contain the effective current system GUID of this virtual port of this network device function.
	// If an assignable system GUID is not supported, this is a read-only alias of the PermanentSystemGUID.
	SystemGUID string
}

// UnmarshalJSON unmarshals a InfiniBand object from the raw JSON.
func (infiniband *InfiniBand) UnmarshalJSON(b []byte) error {
	type temp InfiniBand
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*infiniband = InfiniBand(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Limit shall describe a single array element of the packet and byte limits of a network device function.
type Limit struct {
	// BurstBytesPerSecond shall contain the maximum number of bytes per second in a burst allowed for this network
	// device function.
	BurstBytesPerSecond int
	// BurstPacketsPerSecond shall contain the maximum number of packets per second in a burst allowed for this network
	// device function.
	BurstPacketsPerSecond int
	// Direction shall indicate the direction of the data to which this limit applies for this network device function.
	Direction DataDirection
	// SustainedBytesPerSecond shall contain the maximum number of sustained bytes per second allowed for this network
	// device function.
	SustainedBytesPerSecond int
	// SustainedPacketsPerSecond shall contain the maximum number of sustained packets per second allowed for this
	// network device function.
	SustainedPacketsPerSecond int
}

// UnmarshalJSON unmarshals a Limit object from the raw JSON.
func (limit *Limit) UnmarshalJSON(b []byte) error {
	type temp Limit
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*limit = Limit(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Endpoints shall contain an array of links to resources of type Endpoint that are associated with this network
	// device function.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// EthernetInterfaces shall contain an array of links to resources of type EthernetInterface that represent the
	// virtual interfaces that were created when one of the network device function VLANs is represented as a virtual
	// NIC for the purpose of showing the IP address associated with that VLAN.
	EthernetInterfaces []EthernetInterface
	// EthernetInterfaces@odata.count
	EthernetInterfacesCount int `json:"EthernetInterfaces@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OffloadProcessors shall contain an array of links to resources of type Processor that represent the processors
	// that performs offload computation for this network function, such as with a SmartNIC. This property shall not be
	// present if OffloadSystem is present.
	OffloadProcessors []Processor
	// OffloadProcessors@odata.count
	OffloadProcessorsCount int `json:"OffloadProcessors@odata.count"`
	// OffloadSystem shall contain a link to a resource of type ComputerSystem that represents the system that performs
	// offload computation for this network function, such as with a SmartNIC. The SystemType property contained in the
	// referenced ComputerSystem resource should contain the value 'DPU'. This property shall not be present if
	// OffloadProcessors is present.
	OffloadSystem string
	// PCIeFunction shall contain a link to a resource of type PCIeFunction that represents the PCIe function
	// associated with this network device function.
	PCIeFunction string
	// PhysicalNetworkPortAssignment shall contain a link to a resource of type Port to which this network device
	// function is currently assigned. This value shall be one of the AssignablePhysicalPorts array members.
	PhysicalNetworkPortAssignment string
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

// NetworkDeviceFunction shall represent a logical interface that a network adapter exposes in a Redfish
// implementation.
type NetworkDeviceFunction struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AllowDeny shall contain a link to a resource collection of type AllowDenyCollection that contains the
	// permissions for packets leaving and arriving to this network device function.
	AllowDeny string
	// AssignablePhysicalNetworkPorts shall contain an array of links to resources of type Port that are the physical
	// ports to which this network device function can be assigned.
	AssignablePhysicalNetworkPorts []Port
	// AssignablePhysicalNetworkPorts@odata.count
	AssignablePhysicalNetworkPortsCount int `json:"AssignablePhysicalNetworkPorts@odata.count"`
	// AssignablePhysicalPorts@odata.count
	AssignablePhysicalPortsCount int `json:"AssignablePhysicalPorts@odata.count"`
	// BootMode shall contain the boot mode configured for this network device function. If the value is not
	// 'Disabled', this network device function shall be configured for boot by using the specified technology.
	BootMode BootMode
	// Description provides a description of this resource.
	Description string
	// DeviceEnabled shall indicate whether the network device function is enabled. The operating system shall not
	// enumerate or see disabled network device functions.
	DeviceEnabled bool
	// Ethernet shall contain Ethernet capabilities, status, and configuration values for this network device function.
	Ethernet string
	// FibreChannel shall contain Fibre Channel capabilities, status, and configuration values for this network device
	// function.
	FibreChannel string
	// HTTPBoot shall contain HTTP and HTTPS boot capabilities, status, and configuration values for this network
	// device function.
	HTTPBoot string
	// InfiniBand shall contain InfiniBand capabilities, status, and configuration values for this network device
	// function.
	InfiniBand string
	// Limits shall contain an array of byte and packet limits for this network device function.
	Limits []Limit
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// MaxVirtualFunctions shall contain the number of virtual functions that are available for this network device
	// function.
	MaxVirtualFunctions int
	// Metrics shall contain a link to a resource of type NetworkDeviceFunctionMetrics that contains the metrics
	// associated with this network function.
	Metrics NetworkDeviceFunctionMetrics
	// NetDevFuncCapabilities shall contain an array of capabilities for this network device function.
	NetDevFuncCapabilities []NetworkDeviceTechnology
	// NetDevFuncType shall contain the configured capability of this network device function.
	NetDevFuncType NetworkDeviceTechnology
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SAVIEnabled shall indicate if the RFC7039-defined Source Address Validation Improvement (SAVI) is enabled for
	// this network device function.
	SAVIEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// VirtualFunctionsEnabled shall indicate whether single root input/output virtualization (SR-IOV) virtual
	// functions are enabled for this network device function.
	VirtualFunctionsEnabled bool
	// iSCSIBoot shall contain iSCSI boot capabilities, status, and configuration values for this network device
	// function.
	iSCSIBoot common.Link
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a NetworkDeviceFunction object from the raw JSON.
func (networkdevicefunction *NetworkDeviceFunction) UnmarshalJSON(b []byte) error {
	type temp NetworkDeviceFunction
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*networkdevicefunction = NetworkDeviceFunction(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	networkdevicefunction.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (networkdevicefunction *NetworkDeviceFunction) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(NetworkDeviceFunction)
	original.UnmarshalJSON(networkdevicefunction.rawData)

	readWriteFields := []string{
		"BootMode",
		"DeviceEnabled",
		"NetDevFuncType",
		"SAVIEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(networkdevicefunction).Elem()

	return networkdevicefunction.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetNetworkDeviceFunction will get a NetworkDeviceFunction instance from the service.
func GetNetworkDeviceFunction(c common.Client, uri string) (*NetworkDeviceFunction, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var networkdevicefunction NetworkDeviceFunction
	err = json.NewDecoder(resp.Body).Decode(&networkdevicefunction)
	if err != nil {
		return nil, err
	}

	networkdevicefunction.SetClient(c)
	return &networkdevicefunction, nil
}

// ListReferencedNetworkDeviceFunctions gets the collection of NetworkDeviceFunction from
// a provided reference.
func ListReferencedNetworkDeviceFunctions(c common.Client, link string) ([]*NetworkDeviceFunction, error) {
	var result []*NetworkDeviceFunction
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, networkdevicefunctionLink := range links.ItemLinks {
		networkdevicefunction, err := GetNetworkDeviceFunction(c, networkdevicefunctionLink)
		if err != nil {
			collectionError.Failures[networkdevicefunctionLink] = err
		} else {
			result = append(result, networkdevicefunction)
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

// iSCSIBoot shall describe the iSCSI boot capabilities, status, and configuration values for a network device
// function.
type iSCSIBoot struct {
	// AuthenticationMethod shall contain the iSCSI boot authentication method for this network device function.
	AuthenticationMethod AuthenticationMethod
	// CHAPSecret shall contain the shared secret for CHAP authentication.
	CHAPSecret string
	// CHAPUsername shall contain the user name for CHAP authentication.
	CHAPUsername string
	// IPAddressType shall contain the type of IP address being populated in the iSCSIBoot IP address fields. Mixing
	// IPv6 and IPv4 addresses on the same network device function shall not be permissible.
	IPAddressType IPAddressType
	// IPMaskDNSViaDHCP shall indicate whether the iSCSI boot initiator uses DHCP to obtain the initiator name, IP
	// address, and netmask.
	IPMaskDNSViaDHCP bool
	// InitiatorDefaultGateway shall contain the IPv6 or IPv4 iSCSI boot default gateway.
	InitiatorDefaultGateway string
	// InitiatorIPAddress shall contain the IPv6 or IPv4 address of the iSCSI boot initiator.
	InitiatorIPAddress string
	// InitiatorName shall contain the iSCSI boot initiator name. This property should match formats defined in RFC3720
	// or RFC3721.
	InitiatorName string
	// InitiatorNetmask shall contain the IPv6 or IPv4 netmask of the iSCSI boot initiator.
	InitiatorNetmask string
	// MutualCHAPSecret shall contain the CHAP secret for two-way CHAP authentication.
	MutualCHAPSecret string
	// MutualCHAPUsername shall contain the CHAP user name for two-way CHAP authentication.
	MutualCHAPUsername string
	// PrimaryDNS shall contain the IPv6 or IPv4 address of the primary DNS server for the iSCSI boot initiator.
	PrimaryDNS string
	// PrimaryLUN shall contain the logical unit number (LUN) for the primary iSCSI boot target.
	PrimaryLUN int
	// PrimaryTargetIPAddress shall contain the IPv4 or IPv6 address for the primary iSCSI boot target.
	PrimaryTargetIPAddress string
	// PrimaryTargetName shall contain the name of the primary iSCSI boot target. This property should match formats
	// defined in RFC3720 or RFC3721.
	PrimaryTargetName string
	// PrimaryTargetTCPPort shall contain the TCP port for the primary iSCSI boot target.
	PrimaryTargetTCPPort int
	// PrimaryVLANEnable shall indicate whether this VLAN is enabled for the primary iSCSI boot target.
	PrimaryVLANEnable bool
	// PrimaryVLANId shall contain the 802.1q VLAN ID to use for iSCSI boot from the primary target. This VLAN ID is
	// only used if PrimaryVLANEnable is true.
	PrimaryVLANId int
	// RouterAdvertisementEnabled shall indicate whether IPv6 router advertisement is enabled for the iSCSI boot
	// target. This setting shall apply to only IPv6 configurations.
	RouterAdvertisementEnabled bool
	// SecondaryDNS shall contain the IPv6 or IPv4 address of the secondary DNS server for the iSCSI boot initiator.
	SecondaryDNS string
	// SecondaryLUN shall contain the logical unit number (LUN) for the secondary iSCSI boot target.
	SecondaryLUN int
	// SecondaryTargetIPAddress shall contain the IPv4 or IPv6 address for the secondary iSCSI boot target.
	SecondaryTargetIPAddress string
	// SecondaryTargetName shall contain the name of the secondary iSCSI boot target. This property should match
	// formats defined in RFC3720 or RFC3721.
	SecondaryTargetName string
	// SecondaryTargetTCPPort shall contain the TCP port for the secondary iSCSI boot target.
	SecondaryTargetTCPPort int
	// SecondaryVLANEnable shall indicate whether this VLAN is enabled for the secondary iSCSI boot target.
	SecondaryVLANEnable bool
	// SecondaryVLANId shall contain the 802.1q VLAN ID to use for iSCSI boot from the secondary target. This VLAN ID
	// is only used if SecondaryVLANEnable is 'true'.
	SecondaryVLANId int
	// TargetInfoViaDHCP shall indicate whether the iSCSI boot target name, LUN, IP address, and netmask should be
	// obtained from DHCP.
	TargetInfoViaDHCP bool
}

// UnmarshalJSON unmarshals a iSCSIBoot object from the raw JSON.
func (iscsiboot *iSCSIBoot) UnmarshalJSON(b []byte) error {
	type temp iSCSIBoot
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*iscsiboot = iSCSIBoot(t.temp)

	// Extract the links to other entities for later

	return nil
}
