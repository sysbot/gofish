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

// TemperatureSummary shall contain properties that describe temperature sensor for a subsystem.
type TemperatureSummary struct {
	// Ambient shall contain the temperature, in degrees Celsius units, for the ambient temperature of this subsystem.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Temperature'.
	Ambient SensorExcerpt
	// Exhaust shall contain the temperature, in degrees Celsius units, for the exhaust temperature of this subsystem.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Temperature'.
	Exhaust SensorExcerpt
	// Intake shall contain the temperature, in degrees Celsius units, for the intake temperature of this subsystem.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Temperature'.
	Intake SensorExcerpt
	// Internal shall contain the temperature, in degrees Celsius units, for the internal temperature of this
	// subsystem. The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with
	// the ReadingType property containing the value 'Temperature'.
	Internal SensorExcerpt
}

// UnmarshalJSON unmarshals a TemperatureSummary object from the raw JSON.
func (temperaturesummary *TemperatureSummary) UnmarshalJSON(b []byte) error {
	type temp TemperatureSummary
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*temperaturesummary = TemperatureSummary(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ThermalMetrics shall represent the thermal metrics of a chassis for a Redfish implementation.
type ThermalMetrics struct {
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
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// TemperatureReadingsCelsius shall contain the temperatures, in degrees Celsius units, for this subsystem. The
	// value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType
	// property containing the value 'Temperature'.
	TemperatureReadingsCelsius []SensorArrayExcerpt
	// TemperatureReadingsCelsius@odata.count
	TemperatureReadingsCelsiusCount int `json:"TemperatureReadingsCelsius@odata.count"`
	// TemperatureSummaryCelsius shall contain the temperature sensor readings for this subsystem.
	TemperatureSummaryCelsius TemperatureSummary
}

// UnmarshalJSON unmarshals a ThermalMetrics object from the raw JSON.
func (thermalmetrics *ThermalMetrics) UnmarshalJSON(b []byte) error {
	type temp ThermalMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermalmetrics = ThermalMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetThermalMetrics will get a ThermalMetrics instance from the service.
func GetThermalMetrics(c common.Client, uri string) (*ThermalMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var thermalmetrics ThermalMetrics
	err = json.NewDecoder(resp.Body).Decode(&thermalmetrics)
	if err != nil {
		return nil, err
	}

	thermalmetrics.SetClient(c)
	return &thermalmetrics, nil
}

// ListReferencedThermalMetricss gets the collection of ThermalMetrics from
// a provided reference.
func ListReferencedThermalMetricss(c common.Client, link string) ([]*ThermalMetrics, error) {
	var result []*ThermalMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, thermalmetricsLink := range links.ItemLinks {
		thermalmetrics, err := GetThermalMetrics(c, thermalmetricsLink)
		if err != nil {
			collectionError.Failures[thermalmetricsLink] = err
		} else {
			result = append(result, thermalmetrics)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
