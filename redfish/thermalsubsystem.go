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

// ThermalSubsystem shall represent a thermal subsystem for a Redfish implementation.
type ThermalSubsystem struct {
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
	// FanRedundancy shall contain redundancy information for the groups of fans in this subsystem.
	FanRedundancy []RedundantGroup
	// Fans shall contain a link to a resource collection of type FanCollection.
	Fans string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// ThermalMetrics shall contain a link to a resource collection of type ThermalMetrics.
	ThermalMetrics string
}

// UnmarshalJSON unmarshals a ThermalSubsystem object from the raw JSON.
func (thermalsubsystem *ThermalSubsystem) UnmarshalJSON(b []byte) error {
	type temp ThermalSubsystem
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermalsubsystem = ThermalSubsystem(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetThermalSubsystem will get a ThermalSubsystem instance from the service.
func GetThermalSubsystem(c common.Client, uri string) (*ThermalSubsystem, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var thermalsubsystem ThermalSubsystem
	err = json.NewDecoder(resp.Body).Decode(&thermalsubsystem)
	if err != nil {
		return nil, err
	}

	thermalsubsystem.SetClient(c)
	return &thermalsubsystem, nil
}

// ListReferencedThermalSubsystems gets the collection of ThermalSubsystem from
// a provided reference.
func ListReferencedThermalSubsystems(c common.Client, link string) ([]*ThermalSubsystem, error) {
	var result []*ThermalSubsystem
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, thermalsubsystemLink := range links.ItemLinks {
		thermalsubsystem, err := GetThermalSubsystem(c, thermalsubsystemLink)
		if err != nil {
			collectionError.Failures[thermalsubsystemLink] = err
		} else {
			result = append(result, thermalsubsystem)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
