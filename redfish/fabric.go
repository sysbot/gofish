//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Fabric shall represent a simple switchable fabric for a Redfish implementation.
type Fabric struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// AddressPools shall contain a link to a resource collection of type AddressPoolCollection.
	AddressPools string
	// Connections shall contain a link to a resource collection of type ConnectionCollection.
	Connections string
	// Description provides a description of this resource.
	Description string
	// EndpointGroups shall contain a link to a resource collection of type EndpointGroupCollection.
	EndpointGroups string
	// Endpoints shall contain a link to a resource collection of type EndpointCollection.
	Endpoints string
	// FabricType shall contain the type of fabric being represented by this simple fabric.
	FabricType Protocol
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// MaxZones shall contain the maximum number of zones the switch can currently configure. Changes in the logical or
	// physical configuration of the system can change this value.
	MaxZones int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Switches shall contain a link to a resource collection of type SwitchCollection.
	Switches string
	// UUID shall contain a universal unique identifier number for the fabric.
	UUID string
	// Zones shall contain a link to a resource collection of type ZoneCollection.
	Zones string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Fabric object from the raw JSON.
func (fabric *Fabric) UnmarshalJSON(b []byte) error {
	type temp Fabric
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fabric = Fabric(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	fabric.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (fabric *Fabric) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Fabric)
	original.UnmarshalJSON(fabric.rawData)

	readWriteFields := []string{
		"UUID",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fabric).Elem()

	return fabric.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetFabric will get a Fabric instance from the service.
func GetFabric(c common.Client, uri string) (*Fabric, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fabric Fabric
	err = json.NewDecoder(resp.Body).Decode(&fabric)
	if err != nil {
		return nil, err
	}

	fabric.SetClient(c)
	return &fabric, nil
}

// ListReferencedFabrics gets the collection of Fabric from
// a provided reference.
func ListReferencedFabrics(c common.Client, link string) ([]*Fabric, error) {
	var result []*Fabric
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, fabricLink := range links.ItemLinks {
		fabric, err := GetFabric(c, fabricLink)
		if err != nil {
			collectionError.Failures[fabricLink] = err
		} else {
			result = append(result, fabric)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
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
