//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Links shall contain links to resources that are related to but are not contained by or subordinate to this
// resource.
type Links struct {
	// ManagedBy shall contain an array of links to resources of type Manager that represent the managers that manage
	// this power equipment.
	ManagedBy []Manager
	// ManagedBy@odata.count
	ManagedByCount int `json:"ManagedBy@odata.count"`
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

// PowerEquipment shall be used to represent the set of power equipment for a Redfish implementation.
type PowerEquipment struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// ElectricalBuses shall contain a link to a resource collection of type PowerDistributionCollection that contains
	// a set of electrical bus units.
	ElectricalBuses string
	// FloorPDUs shall contain a link to a resource collection of type PowerDistributionCollection that contains a set
	// of floor power distribution units.
	FloorPDUs string
	// Links shall contain links to resources that are related to but are not contained by or subordinate to this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerShelves shall contain a link to a resource collection of type PowerDistributionCollection that contains a
	// set of power shelves.
	PowerShelves string
	// RackPDUs shall contain a link to a resource collection of type PowerDistributionCollection that contains a set
	// of rack-level power distribution units.
	RackPDUs string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Switchgear shall contain a link to a resource collection of type PowerDistributionCollection that contains a set
	// of switchgear.
	Switchgear string
	// TransferSwitches shall contain a link to a resource collection of type PowerDistributionCollection that contains
	// a set of transfer switches.
	TransferSwitches string
}

// UnmarshalJSON unmarshals a PowerEquipment object from the raw JSON.
func (powerequipment *PowerEquipment) UnmarshalJSON(b []byte) error {
	type temp PowerEquipment
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powerequipment = PowerEquipment(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetPowerEquipment will get a PowerEquipment instance from the service.
func GetPowerEquipment(c common.Client, uri string) (*PowerEquipment, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var powerequipment PowerEquipment
	err = json.NewDecoder(resp.Body).Decode(&powerequipment)
	if err != nil {
		return nil, err
	}

	powerequipment.SetClient(c)
	return &powerequipment, nil
}

// ListReferencedPowerEquipments gets the collection of PowerEquipment from
// a provided reference.
func ListReferencedPowerEquipments(c common.Client, link string) ([]*PowerEquipment, error) {
	var result []*PowerEquipment
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, powerequipmentLink := range links.ItemLinks {
		powerequipment, err := GetPowerEquipment(c, powerequipmentLink)
		if err != nil {
			collectionError.Failures[powerequipmentLink] = err
		} else {
			result = append(result, powerequipment)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
