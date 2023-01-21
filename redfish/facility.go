//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// FacilityType is
type FacilityType string

const (
	// RoomFacilityType A room inside of a building or floor.
	RoomFacilityType FacilityType = "Room"
	// FloorFacilityType A floor inside of a building.
	FloorFacilityType FacilityType = "Floor"
	// BuildingFacilityType A structure with a roof and walls.
	BuildingFacilityType FacilityType = "Building"
	// SiteFacilityType A small area consisting of several buildings.
	SiteFacilityType FacilityType = "Site"
)

// Facility shall be used to represent a location containing equipment, such as a room, building, or campus, for a
// Redfish implementation.
type Facility struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AmbientMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the outdoor
	// environment metrics for this facility.
	AmbientMetrics string
	// Description provides a description of this resource.
	Description string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this facility.
	EnvironmentMetrics string
	// FacilityType shall contain the type of location this resource represents.
	FacilityType string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated facility.
	Location string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerDomains shall contain a link to a resource collection of type PowerDomainCollection that contains the power
	// domains associated with this facility.
	PowerDomains string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a Facility object from the raw JSON.
func (facility *Facility) UnmarshalJSON(b []byte) error {
	type temp Facility
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*facility = Facility(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetFacility will get a Facility instance from the service.
func GetFacility(c common.Client, uri string) (*Facility, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var facility Facility
	err = json.NewDecoder(resp.Body).Decode(&facility)
	if err != nil {
		return nil, err
	}

	facility.SetClient(c)
	return &facility, nil
}

// ListReferencedFacilitys gets the collection of Facility from
// a provided reference.
func ListReferencedFacilitys(c common.Client, link string) ([]*Facility, error) {
	var result []*Facility
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, facilityLink := range links.ItemLinks {
		facility, err := GetFacility(c, facilityLink)
		if err != nil {
			collectionError.Failures[facilityLink] = err
		} else {
			result = append(result, facility)
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
	// ContainedByFacility shall contain a link to a resource of type Facility that represents the facility that
	// contains this facility.
	ContainedByFacility string
	// ContainsChassis shall be an array of links to resources of type Chassis that represent the outermost chassis
	// that this facility contains. This array shall only contain chassis instances that do not include a ContainedBy
	// property within the Links property. That is, only chassis instances that are not contained by another chassis.
	ContainsChassis []Chassis
	// ContainsChassis@odata.count
	ContainsChassisCount int `json:"ContainsChassis@odata.count"`
	// ContainsFacilities shall be an array of links to resources of type Facility that represent the facilities that
	// this facility contains.
	ContainsFacilities []Facility
	// ContainsFacilities@odata.count
	ContainsFacilitiesCount int `json:"ContainsFacilities@odata.count"`
	// ElectricalBuses shall contain an array of links to resources of type PowerDistribution that represent the
	// electrical buses in this facility.
	ElectricalBuses []PowerDistribution
	// ElectricalBuses@odata.count
	ElectricalBusesCount int `json:"ElectricalBuses@odata.count"`
	// FloorPDUs shall be an array of links to resources of type PowerDistribution that represent the floor power
	// distribution units in this facility.
	FloorPDUs []PowerDistribution
	// FloorPDUs@odata.count
	FloorPDUsCount int `json:"FloorPDUs@odata.count"`
	// ManagedBy shall be an array of links to resources of type Manager that represent the managers that manager this
	// facility.
	ManagedBy []Manager
	// ManagedBy@odata.count
	ManagedByCount int `json:"ManagedBy@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerShelves shall be an array of links to resources of type PowerDistribution that represent the power shelves
	// in this facility.
	PowerShelves []PowerDistribution
	// PowerShelves@odata.count
	PowerShelvesCount int `json:"PowerShelves@odata.count"`
	// RackPDUs shall be an array of links to resources of type PowerDistribution that represent the rack-level power
	// distribution units in this facility.
	RackPDUs []PowerDistribution
	// RackPDUs@odata.count
	RackPDUsCount int `json:"RackPDUs@odata.count"`
	// Switchgear shall be an array of links to resources of type PowerDistribution that represent the switchgear in
	// this facility.
	Switchgear []PowerDistribution
	// Switchgear@odata.count
	SwitchgearCount int `json:"Switchgear@odata.count"`
	// TransferSwitches shall be an array of links to resources of type PowerDistribution that represent the transfer
	// switches in this facility.
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
