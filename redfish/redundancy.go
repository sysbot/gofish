//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// RedundancyMode is
type RedundancyMode string

const (
	// FailoverRedundancyMode Failure of one unit automatically causes a standby or offline unit in the redundancy set
	// to take over its functions.
	FailoverRedundancyMode RedundancyMode = "Failover"
	// NmRedundancyMode Multiple units are available and active such that normal operation will continue if one or more
	// units fail.
	NmRedundancyMode RedundancyMode = "N+m"
	// SharingRedundancyMode Multiple units contribute or share such that operation will continue, but at a reduced
	// capacity, if one or more units fail.
	SharingRedundancyMode RedundancyMode = "Sharing"
	// SparingRedundancyMode One or more spare units are available to take over the function of a failed unit, but
	// takeover is not automatic.
	SparingRedundancyMode RedundancyMode = "Sparing"
	// NotRedundantRedundancyMode The subsystem is not configured in a redundancy mode, either due to configuration or
	// the functionality has been disabled by the user.
	NotRedundantRedundancyMode RedundancyMode = "NotRedundant"
)

// RedundancyType is
type RedundancyType string

const (
	// FailoverRedundancyType shall indicate that a failure of one unit automatically causes a standby or offline unit
	// in the redundancy set to take over its functions.
	FailoverRedundancyType RedundancyType = "Failover"
	// NPlusMRedundancyType shall indicate that the capacity or services provided by the set of N+M devices can
	// withstand failure of up to M units, with all units in the group normally providing capacity or service.
	NPlusMRedundancyType RedundancyType = "NPlusM"
	// SharingRedundancyType Multiple units contribute or share such that operation will continue, but at a reduced
	// capacity, if one or more units fail.
	SharingRedundancyType RedundancyType = "Sharing"
	// SparingRedundancyType One or more spare units are available to take over the function of a failed unit, but
	// takeover is not automatic.
	SparingRedundancyType RedundancyType = "Sparing"
	// NotRedundantRedundancyType The subsystem is not configured in a redundancy mode, either due to configuration or
	// the functionality has been disabled by the user.
	NotRedundantRedundancyType RedundancyType = "NotRedundant"
)

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

// Redundancy This object represents the redundancy element property.
type Redundancy struct {
	common.Entity
	// Actions shall contain the available actions for this resource.
	Actions string
	// MaxNumSupported shall contain the maximum number of members allowed in the redundancy group.
	MaxNumSupported int
	// MemberId shall contain the unique identifier for this member within an array. For services supporting Redfish
	// v1.6 or higher, this value shall contain the zero-based array index.
	MemberId string
	// MinNumNeeded shall contain the minimum number of members allowed in the redundancy group for the current
	// redundancy mode to still be fault tolerant.
	MinNumNeeded int
	// Mode shall contain the information about the redundancy mode of this subsystem.
	Mode RedundancyMode
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RedundancyEnabled shall indicate whether the redundancy is enabled.
	RedundancyEnabled bool
	// RedundancySet shall contain the links to components that are part of this redundancy set.
	RedundancySet []idRef
	// RedundancySet@odata.count
	RedundancySetCount int `json:"RedundancySet@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Redundancy object from the raw JSON.
func (redundancy *Redundancy) UnmarshalJSON(b []byte) error {
	type temp Redundancy
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*redundancy = Redundancy(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	redundancy.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (redundancy *Redundancy) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Redundancy)
	original.UnmarshalJSON(redundancy.rawData)

	readWriteFields := []string{
		"Mode",
		"RedundancyEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(redundancy).Elem()

	return redundancy.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetRedundancy will get a Redundancy instance from the service.
func GetRedundancy(c common.Client, uri string) (*Redundancy, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var redundancy Redundancy
	err = json.NewDecoder(resp.Body).Decode(&redundancy)
	if err != nil {
		return nil, err
	}

	redundancy.SetClient(c)
	return &redundancy, nil
}

// ListReferencedRedundancys gets the collection of Redundancy from
// a provided reference.
func ListReferencedRedundancys(c common.Client, link string) ([]*Redundancy, error) {
	var result []*Redundancy
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, redundancyLink := range links.ItemLinks {
		redundancy, err := GetRedundancy(c, redundancyLink)
		if err != nil {
			collectionError.Failures[redundancyLink] = err
		} else {
			result = append(result, redundancy)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// RedundantGroup shall contain redundancy information for the set of devices in this redundancy group.
type RedundantGroup struct {
	// MaxSupportedInGroup shall contain the maximum number of devices allowed in the redundancy group.
	MaxSupportedInGroup int
	// MinNeededInGroup shall contain the minimum number of functional devices needed in the redundancy group for the
	// current redundancy mode to be fault tolerant.
	MinNeededInGroup int
	// RedundancyGroup shall contain the links to the resources that represent the devices that are part of this
	// redundancy group.
	RedundancyGroup []Resource
	// RedundancyGroup@odata.count
	RedundancyGroupCount int `json:"RedundancyGroup@odata.count"`
	// RedundancyType shall contain the information about the redundancy mode of this redundancy group.
	RedundancyType RedundancyType
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a RedundantGroup object from the raw JSON.
func (redundantgroup *RedundantGroup) UnmarshalJSON(b []byte) error {
	type temp RedundantGroup
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*redundantgroup = RedundantGroup(t.temp)

	// Extract the links to other entities for later

	return nil
}
