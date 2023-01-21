//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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

// ServiceConditions shall be used to represent the overall conditions present in a service for a Redfish
// implementation.
type ServiceConditions struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Conditions shall represent a roll-up of the active conditions requiring attention in resources of this Redfish
	// service. The service may roll up any number of conditions originating from resources in the service, using the
	// 'ConditionInRelatedResource' message from Base Message Registry.
	Conditions []Condition
	// Description provides a description of this resource.
	Description string
	// HealthRollup shall contain the highest severity of any messages included in the Conditions property.
	HealthRollup string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a ServiceConditions object from the raw JSON.
func (serviceconditions *ServiceConditions) UnmarshalJSON(b []byte) error {
	type temp ServiceConditions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*serviceconditions = ServiceConditions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetServiceConditions will get a ServiceConditions instance from the service.
func GetServiceConditions(c common.Client, uri string) (*ServiceConditions, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var serviceconditions ServiceConditions
	err = json.NewDecoder(resp.Body).Decode(&serviceconditions)
	if err != nil {
		return nil, err
	}

	serviceconditions.SetClient(c)
	return &serviceconditions, nil
}

// ListReferencedServiceConditionss gets the collection of ServiceConditions from
// a provided reference.
func ListReferencedServiceConditionss(c common.Client, link string) ([]*ServiceConditions, error) {
	var result []*ServiceConditions
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, serviceconditionsLink := range links.ItemLinks {
		serviceconditions, err := GetServiceConditions(c, serviceconditionsLink)
		if err != nil {
			collectionError.Failures[serviceconditionsLink] = err
		} else {
			result = append(result, serviceconditions)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
