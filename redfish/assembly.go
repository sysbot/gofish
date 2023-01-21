//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Assembly shall represent an assembly for a Redfish implementation. Assembly information contains details about a
// device, such as part number, serial number, manufacturer, and production date. It also provides access to the
// original data for the assembly.
type Assembly struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// Assemblies shall define assembly records for a Redfish implementation.
	Assemblies []AssemblyData
	// Assemblies@odata.count
	AssembliesCount int `json:"Assemblies@odata.count"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Assembly object from the raw JSON.
func (assembly *Assembly) UnmarshalJSON(b []byte) error {
	type temp Assembly
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*assembly = Assembly(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	assembly.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (assembly *Assembly) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Assembly)
	original.UnmarshalJSON(assembly.rawData)

	readWriteFields := []string{
		"Assemblies",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(assembly).Elem()

	return assembly.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetAssembly will get a Assembly instance from the service.
func GetAssembly(c common.Client, uri string) (*Assembly, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var assembly Assembly
	err = json.NewDecoder(resp.Body).Decode(&assembly)
	if err != nil {
		return nil, err
	}

	assembly.SetClient(c)
	return &assembly, nil
}

// ListReferencedAssemblys gets the collection of Assembly from
// a provided reference.
func ListReferencedAssemblys(c common.Client, link string) ([]*Assembly, error) {
	var result []*Assembly
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, assemblyLink := range links.ItemLinks {
		assembly, err := GetAssembly(c, assemblyLink)
		if err != nil {
			collectionError.Failures[assemblyLink] = err
		} else {
			result = append(result, assembly)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// AssemblyData
type AssemblyData struct {
	common.Entity
	// Actions shall contain the available actions for this Resource.
	Actions string
	// BinaryDataURI shall contain the URI at which to access an image of the assembly information, using the Redfish
	// protocol and authentication methods. The Service provides this URI for the download of the OEM-specific binary
	// image of the assembly data. An HTTP GET from this URI shall return a response payload of MIME time
	// 'application/octet-stream'. If the service supports it, an HTTP PUT to this URI shall replace the binary image
	// of the assembly.
	BinaryDataURI string
	// Description provides a description of this resource.
	Description string
	// EngineeringChangeLevel shall contain the engineering change level or revision of the assembly.
	EngineeringChangeLevel string
	// Location shall contain location information of the associated assembly.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// MemberId shall contain the unique identifier for this member within an array. For services supporting Redfish
	// v1.6 or higher, this value shall contain the zero-based array index.
	MemberId string
	// Model shall contain the name by which the manufacturer generally refers to the assembly.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number of the assembly.
	PartNumber string
	// PhysicalContext shall contain a description of the physical context for the assembly data.
	PhysicalContext string
	// Producer shall contain the name of the company that produced or manufactured the assembly. This value shall be
	// equal to the 'Manufacturer' field value in a PLDM FRU structure, if applicable, for the assembly.
	Producer string
	// ProductionDate shall contain the date of production or manufacture for the assembly. The time of day portion of
	// the property shall be '00:00:00Z', if the time of day is unknown.
	ProductionDate string
	// Replaceable shall indicate whether the component associated this assembly can be independently replaced as
	// allowed by the vendor's replacement policy. A value of 'false' indicates the component needs to be replaced by
	// policy, as part of another component. If the 'LocationType' property of this assembly contains 'Embedded', this
	// property shall contain 'false'.
	Replaceable bool
	// SKU shall contain the SKU of the assembly.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the assembly.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the assembly.
	SparePartNumber string
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
	// Vendor shall contain the name of the company that provides the final product that includes this assembly. This
	// value shall be equal to the 'Vendor' field value in a PLDM FRU structure, if applicable, for the assembly.
	Vendor string
	// Version shall contain the hardware version of the assembly as determined by the vendor or supplier.
	Version string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AssemblyData object from the raw JSON.
func (assemblydata *AssemblyData) UnmarshalJSON(b []byte) error {
	type temp AssemblyData
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*assemblydata = AssemblyData(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	assemblydata.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (assemblydata *AssemblyData) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AssemblyData)
	original.UnmarshalJSON(assemblydata.rawData)

	readWriteFields := []string{
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(assemblydata).Elem()

	return assemblydata.Entity.Update(originalElement, currentElement, readWriteFields)
}

// AssemblyDataActions shall contain the available actions for this Resource.
type AssemblyDataActions struct {
	// Oem shall contain the available OEM-specific actions for this Resource.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a AssemblyDataActions object from the raw JSON.
func (assemblydataactions *AssemblyDataActions) UnmarshalJSON(b []byte) error {
	type temp AssemblyDataActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*assemblydataactions = AssemblyDataActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// AssemblyDataOemActions shall contain the available OEM-specific actions for this Resource.
type AssemblyDataOemActions struct {
}

// UnmarshalJSON unmarshals a AssemblyDataOemActions object from the raw JSON.
func (assemblydataoemactions *AssemblyDataOemActions) UnmarshalJSON(b []byte) error {
	type temp AssemblyDataOemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*assemblydataoemactions = AssemblyDataOemActions(t.temp)

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
