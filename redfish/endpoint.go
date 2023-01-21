//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// EntityRole is
type EntityRole string

const (
	// InitiatorEntityRole The entity sends commands, messages, or other types of requests to other entities on the
	// fabric, but cannot receive commands from other entities.
	InitiatorEntityRole EntityRole = "Initiator"
	// TargetEntityRole The entity receives commands, messages, or other types of requests from other entities on the
	// fabric, but cannot send commands to other entities.
	TargetEntityRole EntityRole = "Target"
	// BothEntityRole The entity can both send and receive commands, messages, and other requests to or from other
	// entities on the fabric.
	BothEntityRole EntityRole = "Both"
)

// EntityType is
type EntityType string

const (
	// StorageInitiatorEntityType shall indicate the entity this endpoint represents is a storage initiator. The
	// EntityLink property, if present, should be of type StorageController.
	StorageInitiatorEntityType EntityType = "StorageInitiator"
	// RootComplexEntityType shall indicate the entity this endpoint represents is a PCI(e) root complex. The
	// EntityLink property, if present, should be of type ComputerSystem.
	RootComplexEntityType EntityType = "RootComplex"
	// NetworkControllerEntityType shall indicate the entity this endpoint represents is a network controller. The
	// EntityLink property, if present, should be of type NetworkDeviceFunction or EthernetInterface.
	NetworkControllerEntityType EntityType = "NetworkController"
	// DriveEntityType shall indicate the entity this endpoint represents is a drive. The EntityLink property, if
	// present, should be of type Drive.
	DriveEntityType EntityType = "Drive"
	// StorageExpanderEntityType shall indicate the entity this endpoint represents is a storage expander. The
	// EntityLink property, if present, should be of type Chassis.
	StorageExpanderEntityType EntityType = "StorageExpander"
	// DisplayControllerEntityType shall indicate the entity this endpoint represents is a display controller.
	DisplayControllerEntityType EntityType = "DisplayController"
	// BridgeEntityType shall indicate the entity this endpoint represents is a PCI(e) bridge.
	BridgeEntityType EntityType = "Bridge"
	// ProcessorEntityType shall indicate the entity this endpoint represents is a processor. The EntityLink property,
	// if present, should be of type Processor.
	ProcessorEntityType EntityType = "Processor"
	// VolumeEntityType shall indicate the entity this endpoint represents is a volume. The EntityLink property, if
	// present, should be of type Volume.
	VolumeEntityType EntityType = "Volume"
	// AccelerationFunctionEntityType shall indicate the entity this endpoint represents is an acceleration function.
	// The EntityLink property, if present, should be of type AccelerationFunction.
	AccelerationFunctionEntityType EntityType = "AccelerationFunction"
	// MediaControllerEntityType shall indicate the entity this endpoint represents is a media controller. The
	// EntityLink property, if present, should be of type MediaController.
	MediaControllerEntityType EntityType = "MediaController"
	// MemoryChunkEntityType shall indicate the entity this endpoint represents is a memory chunk. The EntityLink
	// property, if present, should be of type MemoryChunk.
	MemoryChunkEntityType EntityType = "MemoryChunk"
	// SwitchEntityType shall indicate the entity this endpoint represents is a switch and not an expander. The
	// EntityLink property, if present, should be of type Switch.
	SwitchEntityType EntityType = "Switch"
	// FabricBridgeEntityType shall indicate the entity this endpoint represents is a fabric bridge. The EntityLink
	// property, if present, should be of type FabricAdapter.
	FabricBridgeEntityType EntityType = "FabricBridge"
	// ManagerEntityType shall indicate the entity this endpoint represents is a manager. The EntityLink property, if
	// present, should be of type Manager.
	ManagerEntityType EntityType = "Manager"
	// StorageSubsystemEntityType shall indicate the entity this endpoint represents is a storage subsystem. The
	// EntityLink property, if present, should be of type Storage.
	StorageSubsystemEntityType EntityType = "StorageSubsystem"
)

// ConnectedEntity shall represent a remote resource that is connected to a network accessible to an endpoint.
type ConnectedEntity struct {
	// EntityLink shall contain a link to an entity of the type specified by the description of the EntityType property
	// value.
	EntityLink string
	// EntityPciId shall contain the PCI ID of the connected PCIe entity.
	EntityPciId string
	// EntityRole shall indicate if the specified entity is an initiator, target, or both.
	EntityRole EntityRole
	// EntityType shall indicate if type of connected entity.
	EntityType EntityType
	// GenZ shall contain the Gen-Z related properties for the entity.
	GenZ GenZ
	// Identifiers shall be unique in the context of other resources that can reached over the connected network.
	Identifiers []Identifier
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a ConnectedEntity object from the raw JSON.
func (connectedentity *ConnectedEntity) UnmarshalJSON(b []byte) error {
	type temp ConnectedEntity
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*connectedentity = ConnectedEntity(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Endpoint This resource contains a fabric endpoint for a Redfish implementation.
type Endpoint struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ConnectedEntities shall contain all entities to which this endpoint allows access.
	ConnectedEntities []ConnectedEntity
	// Description provides a description of this resource.
	Description string
	// EndpointProtocol shall contain the protocol this endpoint uses to communicate with other endpoints on this
	// fabric.
	EndpointProtocol Protocol
	// HostReservationMemoryBytes shall contain the amount of memory in bytes that the host should allocate to connect
	// to this endpoint.
	HostReservationMemoryBytes int
	// IPTransportDetails shall contain the details for each IP transport supported by this endpoint.
	IPTransportDetails []IPTransportDetails
	// Identifiers shall be unique in the context of other endpoints that can reached over the connected network.
	Identifiers []Identifier
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PciId shall contain the PCI ID of the endpoint.
	PciId string
	// Redundancy shall show how this endpoint is grouped with other endpoints for form redundancy sets.
	Redundancy []Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a Endpoint object from the raw JSON.
func (endpoint *Endpoint) UnmarshalJSON(b []byte) error {
	type temp Endpoint
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*endpoint = Endpoint(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetEndpoint will get a Endpoint instance from the service.
func GetEndpoint(c common.Client, uri string) (*Endpoint, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var endpoint Endpoint
	err = json.NewDecoder(resp.Body).Decode(&endpoint)
	if err != nil {
		return nil, err
	}

	endpoint.SetClient(c)
	return &endpoint, nil
}

// ListReferencedEndpoints gets the collection of Endpoint from
// a provided reference.
func ListReferencedEndpoints(c common.Client, link string) ([]*Endpoint, error) {
	var result []*Endpoint
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, endpointLink := range links.ItemLinks {
		endpoint, err := GetEndpoint(c, endpointLink)
		if err != nil {
			collectionError.Failures[endpointLink] = err
		} else {
			result = append(result, endpoint)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// GCID shall contain the Gen-Z Core Specification-defined Global Component ID.
type GCID struct {
	// CID shall contain the 12 bit component identifier portion of the GCID of the entity.
	CID string
	// SID shall contain the 16 bit subnet identifier portion of the GCID of the entity.
	SID string
}

// UnmarshalJSON unmarshals a GCID object from the raw JSON.
func (gcid *GCID) UnmarshalJSON(b []byte) error {
	type temp GCID
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*gcid = GCID(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GenZ shall contain the Gen-Z related properties for an entity.
type GenZ struct {
	// GCID shall contain the Gen-Z Core Specification-defined Global Component ID for the entity.
	GCID GCID
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

// IPTransportDetails shall contain properties that specify the details of the transport supported by the endpoint.
type IPTransportDetails struct {
	// IPv4Address shall contain the IPv4Address.
	IPv4Address string
	// IPv6Address shall contain the IPv6Address.
	IPv6Address string
	// Port shall contain an specify UDP or TCP port number used for communication with the endpoint.
	Port string
	// TransportProtocol shall contain the protocol used by the connection entity.
	TransportProtocol string
}

// UnmarshalJSON unmarshals a IPTransportDetails object from the raw JSON.
func (iptransportdetails *IPTransportDetails) UnmarshalJSON(b []byte) error {
	type temp IPTransportDetails
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*iptransportdetails = IPTransportDetails(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// AddressPools shall contain an array of links to resources of type AddressPool with which this endpoint is
	// associated.
	AddressPools []AddressPool
	// AddressPools@odata.count
	AddressPoolsCount int `json:"AddressPools@odata.count"`
	// ConnectedPorts shall contain an array of links to resources of type Port that represent the switch ports or
	// remote device ports to which this endpoint is connected.
	ConnectedPorts []Port
	// ConnectedPorts@odata.count
	ConnectedPortsCount int `json:"ConnectedPorts@odata.count"`
	// Connections shall contain an array of links to resources of type Connection that represent the connections to
	// which this endpoint belongs.
	Connections []Connection
	// Connections@odata.count
	ConnectionsCount int `json:"Connections@odata.count"`
	// LocalPorts shall contain an array of links to resources of type Port that represent the device ports that this
	// endpoint represents.
	LocalPorts []Port
	// LocalPorts@odata.count
	LocalPortsCount int `json:"LocalPorts@odata.count"`
	// MutuallyExclusiveEndpoints shall contain an array of links to resources of type Endpoint that cannot be used in
	// a zone if this endpoint is in a zone.
	MutuallyExclusiveEndpoints []Endpoint
	// MutuallyExclusiveEndpoints@odata.count
	MutuallyExclusiveEndpointsCount int `json:"MutuallyExclusiveEndpoints@odata.count"`
	// NetworkDeviceFunction shall contain an array of links to resources of type NetworkDeviceFunction with which this
	// endpoint is associated.
	NetworkDeviceFunction []NetworkDeviceFunction
	// NetworkDeviceFunction@odata.count
	NetworkDeviceFunctionCount int `json:"NetworkDeviceFunction@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Ports@odata.count
	PortsCount int `json:"Ports@odata.count"`
	// Zones shall contain an array of links to resources of type Zone that represent the zones to which this endpoint
	// belongs.
	Zones []Zone
	// Zones@odata.count
	ZonesCount int `json:"Zones@odata.count"`
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

// PciId shall describe a PCI ID.
type PciId struct {
	// ClassCode shall contain the PCI Class Code, Subclass, and Programming Interface of the PCIe device function.
	ClassCode string
	// DeviceId shall contain the PCI Device ID of the PCIe device function.
	DeviceId string
	// FunctionNumber shall contain the PCI Function Number of the connected PCIe entity.
	FunctionNumber int
	// SubsystemId shall contain the PCI Subsystem ID of the PCIe device function.
	SubsystemId string
	// SubsystemVendorId shall contain the PCI Subsystem Vendor ID of the PCIe device function.
	SubsystemVendorId string
	// VendorId shall contain the PCI Vendor ID of the PCIe device function.
	VendorId string
}

// UnmarshalJSON unmarshals a PciId object from the raw JSON.
func (pciid *PciId) UnmarshalJSON(b []byte) error {
	type temp PciId
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pciid = PciId(t.temp)

	// Extract the links to other entities for later

	return nil
}
