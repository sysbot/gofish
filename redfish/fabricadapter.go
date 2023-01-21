//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// FabricAdapter A FabricAdapter represents the physical Fabric adapter capable of connecting to an interconnect
// fabric. Examples include but are not limited to Ethernet, NVMe over Fabrics, Gen-Z, and SAS fabric adapters.
type FabricAdapter struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ASICManufacturer shall contain the manufacturer name of the ASIC for the fabric adapter as defined by the
	// manufacturer.
	ASICManufacturer string
	// ASICPartNumber shall contain the part number of the ASIC for the fabric adapter as defined by the manufacturer.
	ASICPartNumber string
	// ASICRevisionIdentifier shall contain the revision identifier of the ASIC for the fabric adapter as defined by
	// the manufacturer.
	ASICRevisionIdentifier string
	// Actions shall contain the available actions for this Resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// FabricType shall contain the configured fabric type of this fabric adapter.
	FabricType string
	// FabricTypeCapabilities shall contain an array of fabric types supported by this fabric adapter.
	FabricTypeCapabilities []Protocol
	// FirmwareVersion shall contain the firmware version for the fabric adapter as defined by the manufacturer.
	FirmwareVersion string
	// GenZ shall contain the Gen-Z specific properties for this fabric adapter.
	GenZ string
	// Links shall contain links to Resources related to but not subordinate to this Resource.
	Links string
	// Location shall contain location information for the fabric adapter.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain a value that represents the manufacturer of the fabric adapter.
	Manufacturer string
	// Model shall contain the information about how the manufacturer refers to this fabric adapter.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeInterface shall contain details on the PCIe interface that connects this PCIe-based fabric adapter to its
	// host.
	PCIeInterface string
	// PartNumber shall contain the part number for the fabric adapter as defined by the manufacturer.
	PartNumber string
	// Ports shall contain a link to a Resource Collection of type PortCollection.
	Ports string
	// SKU shall contain the SKU for the fabric adapter.
	SKU string
	// SerialNumber shall contain the serial number for the fabric adapter.
	SerialNumber string
	// SparePartNumber shall contain the spare part number for the fabric adapter as defined by the manufacturer.
	SparePartNumber string
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
	// UUID shall contain a universal unique identifier number for the fabric adapter.
	UUID string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a FabricAdapter object from the raw JSON.
func (fabricadapter *FabricAdapter) UnmarshalJSON(b []byte) error {
	type temp FabricAdapter
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fabricadapter = FabricAdapter(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	fabricadapter.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (fabricadapter *FabricAdapter) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(FabricAdapter)
	original.UnmarshalJSON(fabricadapter.rawData)

	readWriteFields := []string{
		"FabricType",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fabricadapter).Elem()

	return fabricadapter.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetFabricAdapter will get a FabricAdapter instance from the service.
func GetFabricAdapter(c common.Client, uri string) (*FabricAdapter, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fabricadapter FabricAdapter
	err = json.NewDecoder(resp.Body).Decode(&fabricadapter)
	if err != nil {
		return nil, err
	}

	fabricadapter.SetClient(c)
	return &fabricadapter, nil
}

// ListReferencedFabricAdapters gets the collection of FabricAdapter from
// a provided reference.
func ListReferencedFabricAdapters(c common.Client, link string) ([]*FabricAdapter, error) {
	var result []*FabricAdapter
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, fabricadapterLink := range links.ItemLinks {
		fabricadapter, err := GetFabricAdapter(c, fabricadapterLink)
		if err != nil {
			collectionError.Failures[fabricadapterLink] = err
		} else {
			result = append(result, fabricadapter)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// GenZ shall contain Gen-Z related properties for a fabric adapter.
type GenZ struct {
	// MSDT shall contain a link to a Resource Collection of type RouteEntryCollection, and shall represent the Gen-Z
	// Core Specification-defined MSDT structure.
	MSDT string
	// PIDT shall contain an array of table entry values for the Gen-Z Core Specification-defined Packet Injection
	// Delay Table for the component.
	PIDT []string
	// RITable shall contain an array of table entry values for the Gen-Z Core Specification-defined Responder
	// Interface Table for the component.
	RITable []string
	// RequestorVCAT shall contain a link to a Resource Collection of type VCATEntryCollection, and shall represent the
	// Gen-Z Core Specification-defined REQ-VCAT structure.
	RequestorVCAT string
	// ResponderVCAT shall contain a link to a Resource Collection of type VCATEntryCollection, and shall represent the
	// Gen-Z Core Specification-defined RSP-VCAT structure.
	ResponderVCAT string
	// SSDT shall contain a link to a Resource Collection of type RouteEntryCollection, and shall represent the Gen-Z
	// Core Specification-defined SSDT structure.
	SSDT string
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

// Links shall contain links to Resources related to but not subordinate to this Resource.
type Links struct {
	// Endpoints shall contain an array of links to Resources of type Endpoint that represents the logical fabric
	// connection associated with this fabric adapter.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// MemoryDomains shall contain an array of links to resources of type MemoryDomain that represent the memory
	// domains associated with this fabric adapter.
	MemoryDomains []MemoryDomain
	// MemoryDomains@odata.count
	MemoryDomainsCount int `json:"MemoryDomains@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevices shall contain an array of links to resources of type PCIeDevice that represent the PCIe devices
	// associated with this fabric adapter.
	PCIeDevices []PCIeDevice
	// PCIeDevices@odata.count
	PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
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
