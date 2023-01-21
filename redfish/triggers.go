//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DirectionOfCrossingEnum is The value shall indicate the direction of crossing that corresponds to a trigger.
type DirectionOfCrossingEnum string

const (
	// IncreasingDirectionOfCrossingEnum A trigger condition is met when the metric value crosses the trigger value
	// while increasing.
	IncreasingDirectionOfCrossingEnum DirectionOfCrossingEnum = "Increasing"
	// DecreasingDirectionOfCrossingEnum A trigger is met when the metric value crosses the trigger value while
	// decreasing.
	DecreasingDirectionOfCrossingEnum DirectionOfCrossingEnum = "Decreasing"
)

// DiscreteTriggerConditionEnum is This type shall specify the condition, in relationship to the discrete trigger
// values, which constitutes a trigger.
type DiscreteTriggerConditionEnum string

const (
	// SpecifiedDiscreteTriggerConditionEnum A discrete trigger condition is met when the metric value becomes one of
	// the values that the DiscreteTriggers property lists.
	SpecifiedDiscreteTriggerConditionEnum DiscreteTriggerConditionEnum = "Specified"
	// ChangedDiscreteTriggerConditionEnum A discrete trigger condition is met whenever the metric value changes.
	ChangedDiscreteTriggerConditionEnum DiscreteTriggerConditionEnum = "Changed"
)

// MetricTypeEnum is This type shall specify the type of metric for which the trigger is configured.
type MetricTypeEnum string

const (
	// NumericMetricTypeEnum The trigger is for numeric sensor.
	NumericMetricTypeEnum MetricTypeEnum = "Numeric"
	// DiscreteMetricTypeEnum The trigger is for a discrete sensor.
	DiscreteMetricTypeEnum MetricTypeEnum = "Discrete"
)

// ThresholdActivation is
type ThresholdActivation string

const (
	// IncreasingThresholdActivation This threshold is activated when the reading changes from a value lower than the
	// threshold to a value higher than the threshold.
	IncreasingThresholdActivation ThresholdActivation = "Increasing"
	// DecreasingThresholdActivation This threshold is activated when the reading changes from a value higher than the
	// threshold to a value lower than the threshold.
	DecreasingThresholdActivation ThresholdActivation = "Decreasing"
	// EitherThresholdActivation This threshold is activated when either the Increasing or Decreasing conditions are
	// met.
	EitherThresholdActivation ThresholdActivation = "Either"
)

// TriggerActionEnum is This type shall specify the actions to perform when a trigger condition is met.
type TriggerActionEnum string

const (
	// LogToLogServiceTriggerActionEnum shall log the occurrence of the condition to the log that the LogService
	// property in the telemetry service resource describes.
	LogToLogServiceTriggerActionEnum TriggerActionEnum = "LogToLogService"
	// RedfishEventTriggerActionEnum shall send an event to subscribers.
	RedfishEventTriggerActionEnum TriggerActionEnum = "RedfishEvent"
	// RedfishMetricReportTriggerActionEnum shall force the metric reports managed by the MetricReportDefinitions
	// specified by the MetricReportDefinitions property to be updated, regardless of the MetricReportDefinitionType
	// property value. The actions specified in the ReportActions property of each MetricReportDefinition shall be
	// performed.
	RedfishMetricReportTriggerActionEnum TriggerActionEnum = "RedfishMetricReport"
)

// DiscreteTrigger shall contain the characteristics of the discrete trigger.
type DiscreteTrigger struct {
	common.Entity
	// DwellTime shall contain the amount of time that a trigger event persists before the TriggerActions are
	// performed.
	DwellTime string
	// Severity shall contain the Severity property to be used in the event message.
	Severity Health
	// Value shall contain the value discrete metric that constitutes a trigger event. The DwellTime shall be measured
	// from this point in time.
	Value string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a DiscreteTrigger object from the raw JSON.
func (discretetrigger *DiscreteTrigger) UnmarshalJSON(b []byte) error {
	type temp DiscreteTrigger
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*discretetrigger = DiscreteTrigger(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	discretetrigger.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (discretetrigger *DiscreteTrigger) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(DiscreteTrigger)
	original.UnmarshalJSON(discretetrigger.rawData)

	readWriteFields := []string{
		"DwellTime",
		"Severity",
		"Value",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(discretetrigger).Elem()

	return discretetrigger.Entity.Update(originalElement, currentElement, readWriteFields)
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// MetricReportDefinitions shall contain a set of links to metric report definitions that generate new metric
	// reports when a trigger condition is met and when the TriggerActions property contains 'RedfishMetricReport'.
	MetricReportDefinitions []MetricReportDefinition
	// MetricReportDefinitions@odata.count
	MetricReportDefinitionsCount int `json:"MetricReportDefinitions@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
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

// Threshold shall contain the properties for an individual threshold for this sensor.
type Threshold struct {
	// Activation shall indicate the direction of crossing of the reading for this sensor that activates the threshold.
	Activation ThresholdActivation
	// DwellTime shall indicate the duration the sensor value violates the threshold before the threshold is activated.
	DwellTime string
	// Reading shall indicate the reading for this sensor that activates the threshold. The value of the property shall
	// use the same units as the MetricProperties property.
	Reading float64
}

// UnmarshalJSON unmarshals a Threshold object from the raw JSON.
func (threshold *Threshold) UnmarshalJSON(b []byte) error {
	type temp Threshold
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*threshold = Threshold(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Thresholds shall contain a set of thresholds for a sensor.
type Thresholds struct {
	// LowerCritical shall contain the value at which the MetricProperties property is below the normal range and may
	// require attention. The value of the property shall use the same units as the MetricProperties property.
	LowerCritical string
	// LowerWarning shall contain the value at which the MetricProperties property is below the normal range. The value
	// of the property shall use the same units as the MetricProperties property.
	LowerWarning string
	// UpperCritical shall contain the value at which the MetricProperties property is above the normal range and may
	// require attention. The value of the property shall use the same units as the MetricProperties property.
	UpperCritical string
	// UpperWarning shall contain the value at which the MetricProperties property is above the normal range. The value
	// of the property shall use the same units as the MetricProperties property.
	UpperWarning string
}

// UnmarshalJSON unmarshals a Thresholds object from the raw JSON.
func (thresholds *Thresholds) UnmarshalJSON(b []byte) error {
	type temp Thresholds
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thresholds = Thresholds(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Triggers shall contain a trigger that applies to metrics.
type Triggers struct {
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
	// DiscreteTriggerCondition shall contain the conditions when a discrete metric triggers.
	DiscreteTriggerCondition DiscreteTriggerConditionEnum
	// DiscreteTriggers shall contain a list of values to which to compare a metric reading. This property shall be
	// present when the DiscreteTriggerCondition property is 'Specified'.
	DiscreteTriggers []DiscreteTrigger
	// EventTriggers shall contain an array of MessageIds that specify when a trigger condition is met based on an
	// event. When the service generates an event and if it contains a MessageId within this array, a trigger condition
	// shall be met. The MetricType property should not be present if this resource is configured for event-based
	// triggers.
	EventTriggers []string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// MetricIds shall contain the labels for the metric definitions that contain the property identifiers for this
	// trigger. This property shall match the value of the Id property of the corresponding metric definitions.
	MetricIds []string
	// MetricProperties shall contain an array of URIs with wildcards and property identifiers for this trigger. Use a
	// set of curly braces to delimit each wildcard in the URI. Replace each wildcard with its corresponding entry in
	// the Wildcard array property. A URI that contains wildcards shall link to a resource property to which the metric
	// definition applies after all wildcards are replaced with their corresponding entries in the Wildcard array
	// property. The property identifiers portion of the URI shall follow the RFC6901-defined JSON fragment notation
	// rules.
	MetricProperties []string
	// MetricType shall contain the metric type of the trigger.
	MetricType MetricTypeEnum
	// NumericThresholds shall contain the list of thresholds to which to compare a numeric metric value.
	NumericThresholds string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TriggerActions shall contain the actions that the trigger initiates.
	TriggerActions []TriggerActionEnum
	// Wildcards shall contain the wildcards and their substitution values for the entries in the MetricProperties
	// array property. Each wildcard shall have a corresponding entry in this array property.
	Wildcards []Wildcard
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Triggers object from the raw JSON.
func (triggers *Triggers) UnmarshalJSON(b []byte) error {
	type temp Triggers
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*triggers = Triggers(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	triggers.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (triggers *Triggers) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Triggers)
	original.UnmarshalJSON(triggers.rawData)

	readWriteFields := []string{
		"EventTriggers",
		"MetricIds",
		"MetricProperties",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(triggers).Elem()

	return triggers.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetTriggers will get a Triggers instance from the service.
func GetTriggers(c common.Client, uri string) (*Triggers, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var triggers Triggers
	err = json.NewDecoder(resp.Body).Decode(&triggers)
	if err != nil {
		return nil, err
	}

	triggers.SetClient(c)
	return &triggers, nil
}

// ListReferencedTriggerss gets the collection of Triggers from
// a provided reference.
func ListReferencedTriggerss(c common.Client, link string) ([]*Triggers, error) {
	var result []*Triggers
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, triggersLink := range links.ItemLinks {
		triggers, err := GetTriggers(c, triggersLink)
		if err != nil {
			collectionError.Failures[triggersLink] = err
		} else {
			result = append(result, triggers)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Wildcard shall contain a wildcard and its substitution values.
type Wildcard struct {
	common.Entity
	// Values shall contain the list of values to substitute for the wildcard.
	Values []string
}

// UnmarshalJSON unmarshals a Wildcard object from the raw JSON.
func (wildcard *Wildcard) UnmarshalJSON(b []byte) error {
	type temp Wildcard
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*wildcard = Wildcard(t.temp)

	// Extract the links to other entities for later

	return nil
}
