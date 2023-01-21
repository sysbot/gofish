//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// ElectricalBuses shall contain an array of links to resources of type PowerDistribution that represent the
	// electrical buses in this power domain.
	ElectricalBuses []PowerDistribution
	// ElectricalBuses@odata.count
	ElectricalBusesCount int `json:"ElectricalBuses@odata.count"`
	// FloorPDUs shall contain an array of links to resources of type PowerDistribution that represents the floor power
	// distribution units in this power domain.
	FloorPDUs []PowerDistribution
	// FloorPDUs@odata.count
	FloorPDUsCount int `json:"FloorPDUs@odata.count"`
	// ManagedBy shall contain an array of links to resources of type Manager that represent the managers that manage
	// this power domain.
	ManagedBy []Manager
	// ManagedBy@odata.count
	ManagedByCount int `json:"ManagedBy@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerShelves shall contain an array of links to resources of type PowerDistribution that represents the power
	// shelves in this power domain.
	PowerShelves []PowerDistribution
	// PowerShelves@odata.count
	PowerShelvesCount int `json:"PowerShelves@odata.count"`
	// RackPDUs shall contain an array of links to resources of type PowerDistribution that represents the rack-level
	// power distribution units in this power domain.
	RackPDUs []PowerDistribution
	// RackPDUs@odata.count
	RackPDUsCount int `json:"RackPDUs@odata.count"`
	// Switchgear shall contain an array of links to resources of type PowerDistribution that represents the switchgear
	// in this power domain.
	Switchgear []PowerDistribution
	// Switchgear@odata.count
	SwitchgearCount int `json:"Switchgear@odata.count"`
	// TransferSwitches shall contain an array of links to resources of type PowerDistribution that represents the
	// transfer switches in this power domain.
	TransferSwitches []PowerDistribution
	// TransferSwitches@odata.count
	TransferSwitchesCount int `json:"TransferSwitches@odata.count"`
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

// PowerDomain shall be used to represent a DCIM power domain for a Redfish implementation.
type PowerDomain struct {
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
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a PowerDomain object from the raw JSON.
func (powerdomain *PowerDomain) UnmarshalJSON(b []byte) error {
	type temp PowerDomain
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powerdomain = PowerDomain(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetPowerDomain will get a PowerDomain instance from the service.
func GetPowerDomain(c common.Client, uri string) (*PowerDomain, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var powerdomain PowerDomain
	err = json.NewDecoder(resp.Body).Decode(&powerdomain)
	if err != nil {
		return nil, err
	}

	powerdomain.SetClient(c)
	return &powerdomain, nil
}

// ListReferencedPowerDomains gets the collection of PowerDomain from
// a provided reference.
func ListReferencedPowerDomains(c common.Client, link string) ([]*PowerDomain, error) {
	var result []*PowerDomain
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, powerdomainLink := range links.ItemLinks {
		powerdomain, err := GetPowerDomain(c, powerdomainLink)
		if err != nil {
			collectionError.Failures[powerdomainLink] = err
		} else {
			result = append(result, powerdomain)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
