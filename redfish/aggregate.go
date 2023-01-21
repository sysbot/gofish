//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Aggregate shall represent an aggregation service grouping method for a Redfish implementation.
type Aggregate struct {
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
	// Elements shall contain an array of links to the elements of this aggregate.
	Elements []Resource
	// Elements@odata.count
	ElementsCount int `json:"Elements@odata.count"`
	// ElementsCount shall contain the number of entries in the Elements array.
	ElementsCount int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a Aggregate object from the raw JSON.
func (aggregate *Aggregate) UnmarshalJSON(b []byte) error {
	type temp Aggregate
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*aggregate = Aggregate(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetAggregate will get a Aggregate instance from the service.
func GetAggregate(c common.Client, uri string) (*Aggregate, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var aggregate Aggregate
	err = json.NewDecoder(resp.Body).Decode(&aggregate)
	if err != nil {
		return nil, err
	}

	aggregate.SetClient(c)
	return &aggregate, nil
}

// ListReferencedAggregates gets the collection of Aggregate from
// a provided reference.
func ListReferencedAggregates(c common.Client, link string) ([]*Aggregate, error) {
	var result []*Aggregate
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, aggregateLink := range links.ItemLinks {
		aggregate, err := GetAggregate(c, aggregateLink)
		if err != nil {
			collectionError.Failures[aggregateLink] = err
		} else {
			result = append(result, aggregate)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
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
