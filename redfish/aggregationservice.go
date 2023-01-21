//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AggregationService shall represent an aggregation service for a Redfish implementation.
type AggregationService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Aggregates shall contain a link to a resource collection of type AggregateCollection.
	Aggregates string
	// AggregationSources shall contain a link to a resource collection of type AggregationSourceCollection.
	AggregationSources string
	// ConnectionMethods shall contain a link to a resource collection of type ConnectionMethodCollection.
	ConnectionMethods string
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether the aggregation service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AggregationService object from the raw JSON.
func (aggregationservice *AggregationService) UnmarshalJSON(b []byte) error {
	type temp AggregationService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*aggregationservice = AggregationService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	aggregationservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (aggregationservice *AggregationService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AggregationService)
	original.UnmarshalJSON(aggregationservice.rawData)

	readWriteFields := []string{
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(aggregationservice).Elem()

	return aggregationservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetAggregationService will get a AggregationService instance from the service.
func GetAggregationService(c common.Client, uri string) (*AggregationService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var aggregationservice AggregationService
	err = json.NewDecoder(resp.Body).Decode(&aggregationservice)
	if err != nil {
		return nil, err
	}

	aggregationservice.SetClient(c)
	return &aggregationservice, nil
}

// ListReferencedAggregationServices gets the collection of AggregationService from
// a provided reference.
func ListReferencedAggregationServices(c common.Client, link string) ([]*AggregationService, error) {
	var result []*AggregationService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, aggregationserviceLink := range links.ItemLinks {
		aggregationservice, err := GetAggregationService(c, aggregationserviceLink)
		if err != nil {
			collectionError.Failures[aggregationserviceLink] = err
		} else {
			result = append(result, aggregationservice)
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
