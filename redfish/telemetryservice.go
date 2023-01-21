//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CollectionFunction is If present, the metric value shall be computed according to this function.
type CollectionFunction string

const (
	// AverageCollectionFunction An averaging function.
	AverageCollectionFunction CollectionFunction = "Average"
	// MaximumCollectionFunction A maximum function.
	MaximumCollectionFunction CollectionFunction = "Maximum"
	// MinimumCollectionFunction A minimum function.
	MinimumCollectionFunction CollectionFunction = "Minimum"
	// SummationCollectionFunction A summation function.
	SummationCollectionFunction CollectionFunction = "Summation"
)

// MetricValue shall contain properties that capture a metric value and other associated information.
type MetricValue struct {
	// MetricDefinition shall contain a link to a resource of type MetricDefinition that describes what this metric
	// value captures.
	MetricDefinition string
	// MetricId shall contain the same value as the Id property of the source metric within the associated metric
	// definition.
	MetricId string
	// MetricProperty shall be URI to the property following the JSON fragment notation, as defined by RFC6901, to
	// identify an individual property in a Redfish resource.
	MetricProperty string
	// MetricValue shall contain the metric value, as a string.
	MetricValue string
	// Timestamp shall time when the metric value was obtained. Note that this value may be different from the time
	// when this instance is created.
	Timestamp string
}

// UnmarshalJSON unmarshals a MetricValue object from the raw JSON.
func (metricvalue *MetricValue) UnmarshalJSON(b []byte) error {
	type temp MetricValue
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*metricvalue = MetricValue(t.temp)

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

// TelemetryService This resource contains a telemetry service for a Redfish implementation.
type TelemetryService struct {
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
	// LogService shall contain a link to a resource of type LogService that this telemetry service uses.
	LogService string
	// MaxReports shall contain the maximum number of metric reports that this service supports.
	MaxReports int
	// MetricDefinitions shall contain a link to a resource collection of type MetricDefinitionCollection.
	MetricDefinitions string
	// MetricReportDefinitions shall contain a link to a resource collection of type MetricReportDefinitionCollection.
	MetricReportDefinitions string
	// MetricReports shall contain a link to a resource collection of type MetricReportCollection.
	MetricReports string
	// MinCollectionInterval shall contain the minimum time interval between gathering metric data that this service
	// allows.
	MinCollectionInterval string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedCollectionFunctions shall contain the function to apply over the collection duration.
	SupportedCollectionFunctions []CollectionFunction
	// Triggers shall contain a link to a resource collection of type TriggersCollection.
	Triggers string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a TelemetryService object from the raw JSON.
func (telemetryservice *TelemetryService) UnmarshalJSON(b []byte) error {
	type temp TelemetryService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*telemetryservice = TelemetryService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	telemetryservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (telemetryservice *TelemetryService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(TelemetryService)
	original.UnmarshalJSON(telemetryservice.rawData)

	readWriteFields := []string{
		"ServiceEnabled",
		"SupportedCollectionFunctions",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(telemetryservice).Elem()

	return telemetryservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetTelemetryService will get a TelemetryService instance from the service.
func GetTelemetryService(c common.Client, uri string) (*TelemetryService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var telemetryservice TelemetryService
	err = json.NewDecoder(resp.Body).Decode(&telemetryservice)
	if err != nil {
		return nil, err
	}

	telemetryservice.SetClient(c)
	return &telemetryservice, nil
}

// ListReferencedTelemetryServices gets the collection of TelemetryService from
// a provided reference.
func ListReferencedTelemetryServices(c common.Client, link string) ([]*TelemetryService, error) {
	var result []*TelemetryService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, telemetryserviceLink := range links.ItemLinks {
		telemetryservice, err := GetTelemetryService(c, telemetryserviceLink)
		if err != nil {
			collectionError.Failures[telemetryserviceLink] = err
		} else {
			result = append(result, telemetryservice)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
