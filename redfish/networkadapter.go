//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ControllerCapabilities shall describe the capabilities of a controller.
type ControllerCapabilities struct {
	// DataCenterBridging shall contain capability, status, and configuration values related to data center bridging
	// (DCB) for this controller.
	DataCenterBridging string
	// NPAR shall contain capability, status, and configuration values related to NIC partitioning for this controller.
	NPAR string
	// NPIV shall contain N_Port ID Virtualization (NPIV) capabilities for this controller.
	NPIV string
	// NetworkDeviceFunctionCount shall contain the number of physical functions available on this controller.
	NetworkDeviceFunctionCount int
	// NetworkPortCount shall contain the number of physical ports on this controller.
	NetworkPortCount int
	// VirtualizationOffload shall contain capability, status, and configuration values related to virtualization
	// offload for this controller.
	VirtualizationOffload string
}

// UnmarshalJSON unmarshals a ControllerCapabilities object from the raw JSON.
func (controllercapabilities *ControllerCapabilities) UnmarshalJSON(b []byte) error {
	type temp ControllerCapabilities
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controllercapabilities = ControllerCapabilities(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ControllerLinks shall contain links to resources that are related to but are not contained by, or subordinate
// to, this resource.
type ControllerLinks struct {
	// NetworkDeviceFunctions shall contain an array of links to resources of type NetworkDeviceFunction that represent
	// the network device functions associated with this network controller.
	NetworkDeviceFunctions []NetworkDeviceFunction
	// NetworkDeviceFunctions@odata.count
	NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
	// NetworkPorts@odata.count
	NetworkPortsCount int `json:"NetworkPorts@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevices shall contain an array of links to resources of type PCIeDevice that represent the PCIe devices
	// associated with this network controller.
	PCIeDevices []PCIeDevice
	// PCIeDevices@odata.count
	PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
	// Ports shall contain an array of links to resources of type Port that represent the ports associated with this
	// network controller.
	Ports []Port
	// Ports@odata.count
	PortsCount int `json:"Ports@odata.count"`
}

// UnmarshalJSON unmarshals a ControllerLinks object from the raw JSON.
func (controllerlinks *ControllerLinks) UnmarshalJSON(b []byte) error {
	type temp ControllerLinks
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controllerlinks = ControllerLinks(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Controllers shall describe a network controller ASIC that makes up part of a network adapter.
type Controllers struct {
	// ControllerCapabilities shall contain the capabilities of this controller.
	ControllerCapabilities string
	// FirmwarePackageVersion shall contain the version number of the user-facing firmware package.
	FirmwarePackageVersion string
	// Identifiers shall contain a list of all known durable names for the controller associated with the network
	// adapter.
	Identifiers []Identifier
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the controller associated with the network adapter.
	Location string
	// PCIeInterface shall contain details for the PCIe interface that connects this PCIe-based controller to its host.
	PCIeInterface string
}

// UnmarshalJSON unmarshals a Controllers object from the raw JSON.
func (controllers *Controllers) UnmarshalJSON(b []byte) error {
	type temp Controllers
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controllers = Controllers(t.temp)

	// Extract the links to other entities for later

	return nil
}

// DataCenterBridging shall describe the capability, status, and configuration values related to data center
// bridging (DCB) for a controller.
type DataCenterBridging struct {
	// Capable shall indicate whether this controller is capable of data center bridging (DCB).
	Capable bool
}

// UnmarshalJSON unmarshals a DataCenterBridging object from the raw JSON.
func (datacenterbridging *DataCenterBridging) UnmarshalJSON(b []byte) error {
	type temp DataCenterBridging
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*datacenterbridging = DataCenterBridging(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NPIV shall contain N_Port ID Virtualization (NPIV) capabilities for a controller.
type NPIV struct {
	// MaxDeviceLogins shall contain the maximum number of N_Port ID Virtualization (NPIV) logins allowed
	// simultaneously from all ports on this controller.
	MaxDeviceLogins int
	// MaxPortLogins shall contain the maximum number of N_Port ID Virtualization (NPIV) logins allowed per physical
	// port on this controller.
	MaxPortLogins int
}

// UnmarshalJSON unmarshals a NPIV object from the raw JSON.
func (npiv *NPIV) UnmarshalJSON(b []byte) error {
	type temp NPIV
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*npiv = NPIV(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NetworkAdapter shall represent a physical network adapter capable of connecting to a computer network in a
// Redfish implementation.
type NetworkAdapter struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	Certificates string
	// Controllers shall contain the set of network controllers ASICs that make up this network adapter.
	Controllers []Controllers
	// Description provides a description of this resource.
	Description string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this network adapter.
	EnvironmentMetrics string
	// Identifiers shall contain a list of all known durable names for the network adapter.
	Identifiers []Identifier
	// LLDPEnabled shall contain the state indicating whether LLDP is globally enabled on a network adapter. If set to
	// 'false', the LLDPEnabled value for the ports associated with this adapter shall be disregarded.
	LLDPEnabled string
	// Location shall contain location information of the network adapter.
	Location string
	// Manufacturer shall contain a value that represents the manufacturer of the network adapter.
	Manufacturer string
	// Metrics shall contain a link to a resource of type NetworkAdapterMetrics that contains the metrics associated
	// with this adapter.
	Metrics NetworkAdapterMetrics
	// Model shall contain the information about how the manufacturer refers to this network adapter.
	Model string
	// NetworkDeviceFunctions shall contain a link to a resource collection of type NetworkDeviceFunctionCollection.
	NetworkDeviceFunctions string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number for the network adapter as defined by the manufacturer.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	Ports string
	// Processors shall contain a link to a resource collection of type ProcessorCollection that represent the offload
	// processors contained in this network adapter.
	Processors string
	// SKU shall contain the SKU for the network adapter.
	SKU string
	// SerialNumber shall contain the serial number for the network adapter.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a NetworkAdapter object from the raw JSON.
func (networkadapter *NetworkAdapter) UnmarshalJSON(b []byte) error {
	type temp NetworkAdapter
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*networkadapter = NetworkAdapter(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	networkadapter.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (networkadapter *NetworkAdapter) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(NetworkAdapter)
	original.UnmarshalJSON(networkadapter.rawData)

	readWriteFields := []string{
		"LLDPEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(networkadapter).Elem()

	return networkadapter.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetNetworkAdapter will get a NetworkAdapter instance from the service.
func GetNetworkAdapter(c common.Client, uri string) (*NetworkAdapter, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var networkadapter NetworkAdapter
	err = json.NewDecoder(resp.Body).Decode(&networkadapter)
	if err != nil {
		return nil, err
	}

	networkadapter.SetClient(c)
	return &networkadapter, nil
}

// ListReferencedNetworkAdapters gets the collection of NetworkAdapter from
// a provided reference.
func ListReferencedNetworkAdapters(c common.Client, link string) ([]*NetworkAdapter, error) {
	var result []*NetworkAdapter
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, networkadapterLink := range links.ItemLinks {
		networkadapter, err := GetNetworkAdapter(c, networkadapterLink)
		if err != nil {
			collectionError.Failures[networkadapterLink] = err
		} else {
			result = append(result, networkadapter)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// NicPartitioning shall contain the capability, status, and configuration values for a controller.
type NicPartitioning struct {
	// NparCapable shall indicate whether the controller supports NIC function partitioning.
	NparCapable bool
	// NparEnabled shall indicate whether NIC function partitioning is active on this controller.
	NparEnabled bool
}

// UnmarshalJSON unmarshals a NicPartitioning object from the raw JSON.
func (nicpartitioning *NicPartitioning) UnmarshalJSON(b []byte) error {
	type temp NicPartitioning
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*nicpartitioning = NicPartitioning(t.temp)

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

// SRIOV shall contain single-root input/output virtualization (SR-IOV) capabilities.
type SRIOV struct {
	// SRIOVVEPACapable shall indicate whether this controller supports single root input/output virtualization (SR-
	// IOV) in Virtual Ethernet Port Aggregator (VEPA) mode.
	SRIOVVEPACapable bool
}

// UnmarshalJSON unmarshals a SRIOV object from the raw JSON.
func (sriov *SRIOV) UnmarshalJSON(b []byte) error {
	type temp SRIOV
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sriov = SRIOV(t.temp)

	// Extract the links to other entities for later

	return nil
}

// VirtualFunction shall describe the capability, status, and configuration values related to a virtual function
// for a controller.
type VirtualFunction struct {
	// DeviceMaxCount shall contain the maximum number of virtual functions supported by this controller.
	DeviceMaxCount int
	// MinAssignmentGroupSize shall contain the minimum number of virtual functions that can be allocated or moved
	// between physical functions for this controller.
	MinAssignmentGroupSize int
	// NetworkPortMaxCount shall contain the maximum number of virtual functions supported per network port for this
	// controller.
	NetworkPortMaxCount int
}

// UnmarshalJSON unmarshals a VirtualFunction object from the raw JSON.
func (virtualfunction *VirtualFunction) UnmarshalJSON(b []byte) error {
	type temp VirtualFunction
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*virtualfunction = VirtualFunction(t.temp)

	// Extract the links to other entities for later

	return nil
}

// VirtualizationOffload shall describe the capability, status, and configuration values related to a
// virtualization offload for a controller.
type VirtualizationOffload struct {
	// SRIOV shall contain single-root input/output virtualization (SR-IOV) capabilities.
	SRIOV string
	// VirtualFunction shall describe the capability, status, and configuration values related to the virtual function
	// for this controller.
	VirtualFunction string
}

// UnmarshalJSON unmarshals a VirtualizationOffload object from the raw JSON.
func (virtualizationoffload *VirtualizationOffload) UnmarshalJSON(b []byte) error {
	type temp VirtualizationOffload
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*virtualizationoffload = VirtualizationOffload(t.temp)

	// Extract the links to other entities for later

	return nil
}
