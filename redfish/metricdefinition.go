//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Calculable is The type shall describe the types of calculations that can be applied to the metric reading.
type Calculable string

const (
	// NonCalculatableCalculable No calculations should be performed on the metric reading.
	NonCalculatableCalculable Calculable = "NonCalculatable"
	// SummableCalculable The sum of the metric reading across multiple instances is meaningful.
	SummableCalculable Calculable = "Summable"
	// NonSummableCalculable The sum of the metric reading across multiple instances is not meaningful.
	NonSummableCalculable Calculable = "NonSummable"
)

// CalculationAlgorithmEnum is
type CalculationAlgorithmEnum string

const (
	// AverageCalculationAlgorithmEnum shall be calculated as the average metric reading over a sliding time interval.
	// The time interval shall contain the CalculationTimeInterval property value.
	AverageCalculationAlgorithmEnum CalculationAlgorithmEnum = "Average"
	// MaximumCalculationAlgorithmEnum shall be calculated as the maximum metric reading over a sliding time interval.
	// The time interval shall contain the CalculationTimeInterval property value.
	MaximumCalculationAlgorithmEnum CalculationAlgorithmEnum = "Maximum"
	// MinimumCalculationAlgorithmEnum shall be calculated as the minimum metric reading over a sliding time interval.
	// The time interval shall contain the CalculationTimeInterval property value.
	MinimumCalculationAlgorithmEnum CalculationAlgorithmEnum = "Minimum"
	// OEMCalculationAlgorithmEnum shall be calculated as specified by an OEM. The OEMCalculationAlgorithm property
	// shall contain the specific OEM calculation algorithm.
	OEMCalculationAlgorithmEnum CalculationAlgorithmEnum = "OEM"
)

// ImplementationType is
type ImplementationType string

const (
	// PhysicalSensorImplementationType The metric is implemented as a physical sensor.
	PhysicalSensorImplementationType ImplementationType = "PhysicalSensor"
	// CalculatedImplementationType The metric is implemented by applying a calculation on another metric property. The
	// calculation is specified in the CalculationAlgorithm property.
	CalculatedImplementationType ImplementationType = "Calculated"
	// SynthesizedImplementationType The metric is implemented by applying a calculation on one or more metric
	// properties. The calculation is not provided.
	SynthesizedImplementationType ImplementationType = "Synthesized"
	// DigitalMeterImplementationType The metric is implemented as digital meter.
	DigitalMeterImplementationType ImplementationType = "DigitalMeter"
)

// MetricDataType is This type shall describe the data type of the related metric values as defined by JSON data
// types.
type MetricDataType string

const (
	// BooleanMetricDataType The JSON boolean definition.
	BooleanMetricDataType MetricDataType = "Boolean"
	// DateTimeMetricDataType The JSON string definition with the date-time format.
	DateTimeMetricDataType MetricDataType = "DateTime"
	// DecimalMetricDataType The JSON decimal definition.
	DecimalMetricDataType MetricDataType = "Decimal"
	// IntegerMetricDataType The JSON integer definition.
	IntegerMetricDataType MetricDataType = "Integer"
	// StringMetricDataType The JSON string definition.
	StringMetricDataType MetricDataType = "String"
	// EnumerationMetricDataType The JSON string definition with a set of defined enumerations.
	EnumerationMetricDataType MetricDataType = "Enumeration"
)

// MetricType is This property shall contain the type of metric.
type MetricType string

const (
	// NumericMetricType The metric is a numeric metric. The metric value is any real number.
	NumericMetricType MetricType = "Numeric"
	// DiscreteMetricType shall indicate discrete states.
	DiscreteMetricType MetricType = "Discrete"
	// GaugeMetricType The metric is a gauge metric. The metric value is a real number. When the metric value reaches
	// the gauge's extrema, it stays at that value, until the reading falls within the extrema.
	GaugeMetricType MetricType = "Gauge"
	// CounterMetricType The metric is a counter metric. The metric reading is a non-negative integer that increases
	// monotonically. When a counter reaches its maximum, the value resets to 0 and resumes counting.
	CounterMetricType MetricType = "Counter"
	// CountdownMetricType The metric is a countdown metric. The metric reading is a non-negative integer that
	// decreases monotonically. When a counter reaches its minimum, the value resets to preset value and resumes
	// counting down.
	CountdownMetricType MetricType = "Countdown"
	// StringMetricType The metric is a non-discrete string metric. The metric reading is a non-discrete string that
	// displays some non-discrete, non-numeric data.
	StringMetricType MetricType = "String"
)

// CalculationParamsType shall contain the parameters for a metric calculation.
type CalculationParamsType struct {
	// ResultMetric shall contain a URI with wildcards and property identifiers of the metric property that stores the
	// result of the calculation. A set of curly braces shall delimit each wildcard in the URI. The corresponding entry
	// in the Wildcard property shall replace each wildcard. After each wildcard is replaced, it shall identify a
	// resource property to which the metric definition applies. The property identifiers portion of the URI shall
	// follow RFC6901-defined JSON pointer notation rules.
	ResultMetric string
	// SourceMetric shall contain a URI with wildcards and property identifiers of the metric property used as the
	// input into the calculation. A set of curly braces shall delimit each wildcard in the URI. The corresponding
	// entry in the Wildcard property shall replace each wildcard. After each wildcard is replaced, it shall identify a
	// resource property to which the metric definition applies. The property identifiers portion of the URI shall
	// follow RFC6901-defined JSON pointer notation rules.
	SourceMetric string
}

// UnmarshalJSON unmarshals a CalculationParamsType object from the raw JSON.
func (calculationparamstype *CalculationParamsType) UnmarshalJSON(b []byte) error {
	type temp CalculationParamsType
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*calculationparamstype = CalculationParamsType(t.temp)

	// Extract the links to other entities for later

	return nil
}

// MetricDefinition shall contain the metadata information for a metric in a Redfish implementation.
type MetricDefinition struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Accuracy shall contain the percent error +/- of the measured versus actual values. The property is not
	// meaningful when the MetricType property is 'Discrete'.
	Accuracy float64
	// Actions shall contain the available actions for this resource.
	Actions string
	// Calculable shall specify whether the metric can be used in a calculation.
	Calculable Calculable
	// CalculationAlgorithm shall contain the calculation performed to obtain the metric.
	CalculationAlgorithm CalculationAlgorithmEnum
	// CalculationParameters shall list the metric properties that are part of a calculation that this metric
	// definition defines. This property should be present if ImplementationType contains 'Synthesized' or
	// 'Calculated'.
	CalculationParameters []CalculationParamsType
	// CalculationTimeInterval shall specify the time interval over the metric calculation is performed.
	CalculationTimeInterval string
	// Calibration shall contain the calibration offset added to the metric reading. The value shall have the units
	// specified in the Units property. The property is not meaningful when the MetricType property is 'Discrete'.
	Calibration float64
	// Description provides a description of this resource.
	Description string
	// DiscreteValues shall specify the possible values of the discrete metric. This property shall have values when
	// the MetricType property is 'Discrete'.
	DiscreteValues []string
	// Implementation shall specify the implementation of the metric.
	Implementation ImplementationType
	// IsLinear shall indicate whether the metric values are linear versus non-linear. Linear metrics can use a greater
	// than relation to compared them. An example of linear metrics include performance metrics. Examples of non-linear
	// metrics include error codes.
	IsLinear bool
	// LogicalContexts shall contain the logical contexts related to the metric. This property should be present when
	// the PhysicalContext property does not provide complete information and additional context information is needed.
	// For example, if the metric refers to capacity or performance.
	LogicalContexts []LogicalContext
	// MaxReadingRange shall indicate the highest possible value for a related MetricValue. The value shall have the
	// units specified in the property Units. The property is not meaningful when the MetricType property is
	// 'Discrete'.
	MaxReadingRange float64
	// MetricDataType shall specify the data-type of the metric.
	MetricDataType MetricDataType
	// MetricProperties shall list the URIs with wildcards and property identifiers that this metric defines. A set of
	// curly braces shall delimit each wildcard in the URI. The corresponding entry in the Wildcard property shall
	// replace each wildcard. After each wildcard is replaced, it shall identify a resource property to which the
	// metric definition applies. The property identifiers portion of the URI shall follow RFC6901-defined JSON pointer
	// notation rules. This property should not be present if ImplementationType contains 'Synthesized' or
	// 'Calculated'.
	MetricProperties []string
	// MetricType shall specify the type of metric.
	MetricType MetricType
	// MinReadingRange shall contain the lowest possible value for the metric reading. The value shall have the units
	// specified in the property Units. The property is not meaningful when the MetricType property is 'Discrete'.
	MinReadingRange float64
	// OEMCalculationAlgorithm shall contain the OEM-defined calculation performed to obtain the metric. This property
	// shall be present if CalculationAlgorithm is 'OEM'.
	OEMCalculationAlgorithm string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalContext shall contain the physical context of the metric.
	PhysicalContext PhysicalContext
	// Precision shall specify the number of significant digits in the metric reading. The property is not meaningful
	// when the MetricType property is 'Discrete'.
	Precision int
	// SensingInterval shall specify the time interval between when a metric is updated.
	SensingInterval string
	// TimestampAccuracy shall specify the expected + or - variability of the timestamp.
	TimestampAccuracy string
	// Units shall specify the units of the metric. This property shall be consistent with the case-sensitive ('C/s'
	// column) Unified Code for Units of Measure. Note: Not all units of measured are covered by UCUM.
	Units string
	// Wildcards shall contain a list of wildcards and their replacement strings, which are applied to the
	// MetricProperties array property. Each wildcard shall have a corresponding entry in this array property.
	Wildcards []Wildcard
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a MetricDefinition object from the raw JSON.
func (metricdefinition *MetricDefinition) UnmarshalJSON(b []byte) error {
	type temp MetricDefinition
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*metricdefinition = MetricDefinition(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	metricdefinition.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (metricdefinition *MetricDefinition) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(MetricDefinition)
	original.UnmarshalJSON(metricdefinition.rawData)

	readWriteFields := []string{
		"Calculable",
		"CalculationTimeInterval",
		"DiscreteValues",
		"IsLinear",
		"MetricDataType",
		"MetricProperties",
		"MetricType",
		"SensingInterval",
		"Units",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(metricdefinition).Elem()

	return metricdefinition.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMetricDefinition will get a MetricDefinition instance from the service.
func GetMetricDefinition(c common.Client, uri string) (*MetricDefinition, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var metricdefinition MetricDefinition
	err = json.NewDecoder(resp.Body).Decode(&metricdefinition)
	if err != nil {
		return nil, err
	}

	metricdefinition.SetClient(c)
	return &metricdefinition, nil
}

// ListReferencedMetricDefinitions gets the collection of MetricDefinition from
// a provided reference.
func ListReferencedMetricDefinitions(c common.Client, link string) ([]*MetricDefinition, error) {
	var result []*MetricDefinition
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, metricdefinitionLink := range links.ItemLinks {
		metricdefinition, err := GetMetricDefinition(c, metricdefinitionLink)
		if err != nil {
			collectionError.Failures[metricdefinitionLink] = err
		} else {
			result = append(result, metricdefinition)
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
