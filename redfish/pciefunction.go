//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DeviceClass is
type DeviceClass string

const (
	// UnclassifiedDeviceDeviceClass An unclassified device.
	UnclassifiedDeviceDeviceClass DeviceClass = "UnclassifiedDevice"
	// MassStorageControllerDeviceClass A mass storage controller.
	MassStorageControllerDeviceClass DeviceClass = "MassStorageController"
	// NetworkControllerDeviceClass A network controller.
	NetworkControllerDeviceClass DeviceClass = "NetworkController"
	// DisplayControllerDeviceClass A display controller.
	DisplayControllerDeviceClass DeviceClass = "DisplayController"
	// MultimediaControllerDeviceClass A multimedia controller.
	MultimediaControllerDeviceClass DeviceClass = "MultimediaController"
	// MemoryControllerDeviceClass A memory controller.
	MemoryControllerDeviceClass DeviceClass = "MemoryController"
	// BridgeDeviceClass A bridge.
	BridgeDeviceClass DeviceClass = "Bridge"
	// CommunicationControllerDeviceClass A communication controller.
	CommunicationControllerDeviceClass DeviceClass = "CommunicationController"
	// GenericSystemPeripheralDeviceClass A generic system peripheral.
	GenericSystemPeripheralDeviceClass DeviceClass = "GenericSystemPeripheral"
	// InputDeviceControllerDeviceClass An input device controller.
	InputDeviceControllerDeviceClass DeviceClass = "InputDeviceController"
	// DockingStationDeviceClass A docking station.
	DockingStationDeviceClass DeviceClass = "DockingStation"
	// ProcessorDeviceClass A processor.
	ProcessorDeviceClass DeviceClass = "Processor"
	// SerialBusControllerDeviceClass A serial bus controller.
	SerialBusControllerDeviceClass DeviceClass = "SerialBusController"
	// WirelessControllerDeviceClass A wireless controller.
	WirelessControllerDeviceClass DeviceClass = "WirelessController"
	// IntelligentControllerDeviceClass An intelligent controller.
	IntelligentControllerDeviceClass DeviceClass = "IntelligentController"
	// SatelliteCommunicationsControllerDeviceClass A satellite communications controller.
	SatelliteCommunicationsControllerDeviceClass DeviceClass = "SatelliteCommunicationsController"
	// EncryptionControllerDeviceClass An encryption controller.
	EncryptionControllerDeviceClass DeviceClass = "EncryptionController"
	// SignalProcessingControllerDeviceClass A signal processing controller.
	SignalProcessingControllerDeviceClass DeviceClass = "SignalProcessingController"
	// ProcessingAcceleratorsDeviceClass A processing accelerators.
	ProcessingAcceleratorsDeviceClass DeviceClass = "ProcessingAccelerators"
	// NonEssentialInstrumentationDeviceClass A non-essential instrumentation.
	NonEssentialInstrumentationDeviceClass DeviceClass = "NonEssentialInstrumentation"
	// CoprocessorDeviceClass A coprocessor.
	CoprocessorDeviceClass DeviceClass = "Coprocessor"
	// UnassignedClassDeviceClass An unassigned class.
	UnassignedClassDeviceClass DeviceClass = "UnassignedClass"
	// OtherDeviceClass A other class. The function Device Class Id needs to be verified.
	OtherDeviceClass DeviceClass = "Other"
)

// FunctionType is
type FunctionType string

const (
	// PhysicalFunctionType A physical PCIe function.
	PhysicalFunctionType FunctionType = "Physical"
	// VirtualFunctionType A virtual PCIe function.
	VirtualFunctionType FunctionType = "Virtual"
)

// Links shall contain links to Resources that are related to but are not contained by or subordinate to this
// Resource.
type Links struct {
	// Drives shall link to a Resource of type Drive that represents the storage drives associated with this Resource.
	Drives []Drive
	// Drives@odata.count
	DrivesCount int `json:"Drives@odata.count"`
	// EthernetInterfaces shall link to a Resource of type EthernetInterface that represents the network interfaces
	// associated with this Resource.
	EthernetInterfaces []EthernetInterface
	// EthernetInterfaces@odata.count
	EthernetInterfacesCount int `json:"EthernetInterfaces@odata.count"`
	// NetworkDeviceFunctions shall contain an array of links to Resources of the NetworkDeviceFunction type that
	// represents the network device functions associated with this Resource.
	NetworkDeviceFunctions []NetworkDeviceFunction
	// NetworkDeviceFunctions@odata.count
	NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevice shall contain a link to a Resource of type PCIeDevice of which this function is a part.
	PCIeDevice string
	// Processor shall link to a resource of type Processor that is hosted on this PCIe device function.
	Processor Processor
	// StorageControllers shall link to a Resource of type StorageController that represents the storage controllers
	// associated with this Resource.
	StorageControllers []StorageController
	// StorageControllers@odata.count
	StorageControllersCount int `json:"StorageControllers@odata.count"`
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

// OemActions shall contain the available OEM-specific actions for this Resource.
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

// PCIeFunction shall represent a PCIeFunction attached to a System.
type PCIeFunction struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// ClassCode shall contain the PCI Class Code of the PCIe device function.
	ClassCode string
	// Description provides a description of this resource.
	Description string
	// DeviceClass shall contain the device class of the PCIe device function, such as storage, network, or memory.
	DeviceClass string
	// DeviceId shall contain the PCI Device ID of the PCIe device function.
	DeviceId string
	// Enabled shall indicate if this PCIe device function is enabled.
	Enabled string
	// FunctionId shall contain the PCIe Function Number within a given PCIe device.
	FunctionId int
	// FunctionType shall contain the function type of the PCIe device function such as Physical or Virtual.
	FunctionType string
	// Links shall contain links to Resources that are related to but are not contained by, or subordinate to, this
	// Resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RevisionId shall contain the PCI Revision ID of the PCIe device function.
	RevisionId string
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
	// SubsystemId shall contain the PCI Subsystem ID of the PCIe device function.
	SubsystemId string
	// SubsystemVendorId shall contain the PCI Subsystem Vendor ID of the PCIe device function.
	SubsystemVendorId string
	// VendorId shall contain the PCI Vendor ID of the PCIe device function.
	VendorId string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PCIeFunction object from the raw JSON.
func (pciefunction *PCIeFunction) UnmarshalJSON(b []byte) error {
	type temp PCIeFunction
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pciefunction = PCIeFunction(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	pciefunction.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (pciefunction *PCIeFunction) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(PCIeFunction)
	original.UnmarshalJSON(pciefunction.rawData)

	readWriteFields := []string{
		"Enabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(pciefunction).Elem()

	return pciefunction.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPCIeFunction will get a PCIeFunction instance from the service.
func GetPCIeFunction(c common.Client, uri string) (*PCIeFunction, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pciefunction PCIeFunction
	err = json.NewDecoder(resp.Body).Decode(&pciefunction)
	if err != nil {
		return nil, err
	}

	pciefunction.SetClient(c)
	return &pciefunction, nil
}

// ListReferencedPCIeFunctions gets the collection of PCIeFunction from
// a provided reference.
func ListReferencedPCIeFunctions(c common.Client, link string) ([]*PCIeFunction, error) {
	var result []*PCIeFunction
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, pciefunctionLink := range links.ItemLinks {
		pciefunction, err := GetPCIeFunction(c, pciefunctionLink)
		if err != nil {
			collectionError.Failures[pciefunctionLink] = err
		} else {
			result = append(result, pciefunction)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
