//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// OemActions shall contain any additional OEM actions for this resource.
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

// PowerAllocation shall contain the set of properties describing the allocation of power for a subsystem.
type PowerAllocation struct {
	// AllocatedWatts shall represent the total power currently allocated or budgeted to this subsystem.
	AllocatedWatts float64
	// RequestedWatts shall represent the amount of power, in watts, that the subsystem currently requests to be
	// budgeted for future use.
	RequestedWatts float64
}

// UnmarshalJSON unmarshals a PowerAllocation object from the raw JSON.
func (powerallocation *PowerAllocation) UnmarshalJSON(b []byte) error {
	type temp PowerAllocation
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powerallocation = PowerAllocation(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PowerSubsystem shall represent a power subsystem for a Redfish implementation.
type PowerSubsystem struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Allocation shall contain the set of properties describing the allocation of power for this subsystem.
	Allocation string
	// Batteries shall contain a link to a resource collection of type BatteryCollection.
	Batteries string
	// CapacityWatts shall represent the total power capacity that can be allocated to this subsystem.
	CapacityWatts float64
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerSupplies shall contain a link to a resource collection of type PowerSupplyCollection.
	PowerSupplies string
	// PowerSupplyRedundancy shall contain redundancy information for the set of power supplies in this subsystem. The
	// values of the RedundancyGroup array shall reference resources of type PowerSupply.
	PowerSupplyRedundancy []RedundantGroup
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a PowerSubsystem object from the raw JSON.
func (powersubsystem *PowerSubsystem) UnmarshalJSON(b []byte) error {
	type temp PowerSubsystem
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powersubsystem = PowerSubsystem(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetPowerSubsystem will get a PowerSubsystem instance from the service.
func GetPowerSubsystem(c common.Client, uri string) (*PowerSubsystem, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var powersubsystem PowerSubsystem
	err = json.NewDecoder(resp.Body).Decode(&powersubsystem)
	if err != nil {
		return nil, err
	}

	powersubsystem.SetClient(c)
	return &powersubsystem, nil
}

// ListReferencedPowerSubsystems gets the collection of PowerSubsystem from
// a provided reference.
func ListReferencedPowerSubsystems(c common.Client, link string) ([]*PowerSubsystem, error) {
	var result []*PowerSubsystem
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, powersubsystemLink := range links.ItemLinks {
		powersubsystem, err := GetPowerSubsystem(c, powersubsystemLink)
		if err != nil {
			collectionError.Failures[powersubsystemLink] = err
		} else {
			result = append(result, powersubsystem)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
