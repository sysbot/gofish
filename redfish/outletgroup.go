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
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Outlets shall be an array of links to resources of type Outlet that represent the outlets in this outlet group.
	Outlets []Outlet
	// Outlets@odata.count
	OutletsCount int `json:"Outlets@odata.count"`
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

// OutletGroup shall be used to represent an electrical outlet group for a Redfish implementation.
type OutletGroup struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ConfigurationLocked shall indicate whether modification requests to this resource are not permitted. If 'true',
	// services shall reject modification requests to other properties in this resource.
	ConfigurationLocked string
	// CreatedBy shall contain the name of the person or application that created this outlet group.
	CreatedBy string
	// Description provides a description of this resource.
	Description string
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this outlet group, that represents the
	// 'Total' ElectricalContext sensor when multiple energy sensors exist for this outlet group. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerControlLocked shall indicate whether requests to the PowerControl action are locked. If 'true', services
	// shall reject requests to the PowerControl action.
	PowerControlLocked string
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on after a PowerControl action to
	// cycle power. The value '0' shall indicate no delay to power on.
	PowerCycleDelaySeconds float64
	// PowerEnabled shall contain the power enable state of the outlet group. True shall indicate that the group can be
	// powered on, and false shall indicate that the group cannot be powered.
	PowerEnabled bool
	// PowerOffDelaySeconds shall contain the number of seconds to delay power off after a PowerControl action. The
	// value '0' shall indicate no delay to power off.
	PowerOffDelaySeconds float64
	// PowerOnDelaySeconds shall contain the number of seconds to delay power up after a power cycle or a PowerControl
	// action. The value '0' shall indicate no delay to power up.
	PowerOnDelaySeconds float64
	// PowerRestoreDelaySeconds shall contain the number of seconds to delay power on after a power fault. The value
	// '0' shall indicate no delay to power on.
	PowerRestoreDelaySeconds float64
	// PowerRestorePolicy shall contain the desired PowerState of the outlet group when power is applied. The value
	// 'LastState' shall return the outlet group to the PowerState it was in when power was lost.
	PowerRestorePolicy string
	// PowerState shall contain the power state of the outlet group.
	PowerState PowerState
	// PowerStateInTransition shall indicate whether the PowerState property will undergo a transition between on and
	// off states due to a configured delay. The transition may be due to the configuration of the power on, off, or
	// restore delay properties. If 'true', the PowerState property will transition at the conclusion of a configured
	// delay.
	PowerStateInTransition string
	// PowerWatts shall contain the total power, in watt units, for this outlet group, that represents the 'Total'
	// ElectricalContext sensor when multiple power sensors exist for this outlet group. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'.
	PowerWatts SensorPowerExcerpt
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a OutletGroup object from the raw JSON.
func (outletgroup *OutletGroup) UnmarshalJSON(b []byte) error {
	type temp OutletGroup
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*outletgroup = OutletGroup(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	outletgroup.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (outletgroup *OutletGroup) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(OutletGroup)
	original.UnmarshalJSON(outletgroup.rawData)

	readWriteFields := []string{
		"ConfigurationLocked",
		"CreatedBy",
		"PowerControlLocked",
		"PowerCycleDelaySeconds",
		"PowerOffDelaySeconds",
		"PowerOnDelaySeconds",
		"PowerRestoreDelaySeconds",
		"PowerRestorePolicy",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(outletgroup).Elem()

	return outletgroup.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetOutletGroup will get a OutletGroup instance from the service.
func GetOutletGroup(c common.Client, uri string) (*OutletGroup, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var outletgroup OutletGroup
	err = json.NewDecoder(resp.Body).Decode(&outletgroup)
	if err != nil {
		return nil, err
	}

	outletgroup.SetClient(c)
	return &outletgroup, nil
}

// ListReferencedOutletGroups gets the collection of OutletGroup from
// a provided reference.
func ListReferencedOutletGroups(c common.Client, link string) ([]*OutletGroup, error) {
	var result []*OutletGroup
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, outletgroupLink := range links.ItemLinks {
		outletgroup, err := GetOutletGroup(c, outletgroupLink)
		if err != nil {
			collectionError.Failures[outletgroupLink] = err
		} else {
			result = append(result, outletgroup)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
