//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// EnvironmentMetrics shall represent the environmental metrics for a Redfish implementation.
type EnvironmentMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AbsoluteHumidity shall contain the absolute (volumetric) humidity sensor reading, in grams/cubic meter units,
	// for this resource. The value of the DataSourceUri property, if present, shall reference a resource of type
	// Sensor with the ReadingType property containing the value 'AbsoluteHumidity'.
	AbsoluteHumidity SensorExcerpt
	// Actions shall contain the available actions for this resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// DewPointCelsius shall contain the dew point, in degrees Celsius, based on the temperature and humidity values
	// for this resource. The value of the DataSourceUri property, if present, shall reference a resource of type
	// Sensor with the ReadingType property containing the value 'Temperature'.
	DewPointCelsius SensorExcerpt
	// EnergyJoules shall contain the total energy, in joules, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'EnergyJoules'. This property is used for reporting device-level energy consumption measurements, while
	// EnergykWh is used for large-scale consumption measurements.
	EnergyJoules SensorExcerpt
	// EnergykWh shall contain the total energy, in kilowatt-hours, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// FanSpeedsPercent shall contain the fan speeds, in percent units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Percent'.
	FanSpeedsPercent []SensorFanArrayExcerpt
	// FanSpeedsPercent@odata.count
	FanSpeedsPercentCount int `json:"FanSpeedsPercent@odata.count"`
	// HumidityPercent shall contain the humidity, in percent units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Humidity'.
	HumidityPercent SensorExcerpt
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerLimitWatts shall contain the power limit control, in watt units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Control with the ControlType property
	// containing the value of 'Power'.
	PowerLimitWatts ControlSingleExcerpt
	// PowerLoadPercent shall contain the power load, in percent units, for this device, that represents the 'Total'
	// ElectricalContext for this device. The value of the DataSourceUri property, if present, shall reference a
	// resource of type Sensor with the ReadingType property containing the value 'Percent'.
	PowerLoadPercent SensorExcerpt
	// PowerWatts shall contain the total power, in watt units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'.
	PowerWatts SensorPowerExcerpt
	// TemperatureCelsius shall contain the temperature, in degrees Celsius units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EnvironmentMetrics object from the raw JSON.
func (environmentmetrics *EnvironmentMetrics) UnmarshalJSON(b []byte) error {
	type temp EnvironmentMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*environmentmetrics = EnvironmentMetrics(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	environmentmetrics.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (environmentmetrics *EnvironmentMetrics) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(EnvironmentMetrics)
	original.UnmarshalJSON(environmentmetrics.rawData)

	readWriteFields := []string{
		"PowerLimitWatts",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(environmentmetrics).Elem()

	return environmentmetrics.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEnvironmentMetrics will get a EnvironmentMetrics instance from the service.
func GetEnvironmentMetrics(c common.Client, uri string) (*EnvironmentMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var environmentmetrics EnvironmentMetrics
	err = json.NewDecoder(resp.Body).Decode(&environmentmetrics)
	if err != nil {
		return nil, err
	}

	environmentmetrics.SetClient(c)
	return &environmentmetrics, nil
}

// ListReferencedEnvironmentMetricss gets the collection of EnvironmentMetrics from
// a provided reference.
func ListReferencedEnvironmentMetricss(c common.Client, link string) ([]*EnvironmentMetrics, error) {
	var result []*EnvironmentMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, environmentmetricsLink := range links.ItemLinks {
		environmentmetrics, err := GetEnvironmentMetrics(c, environmentmetricsLink)
		if err != nil {
			collectionError.Failures[environmentmetricsLink] = err
		} else {
			result = append(result, environmentmetrics)
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
