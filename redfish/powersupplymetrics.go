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

// PowerSupplyMetrics shall be used to represent the metrics of a power supply unit for a Redfish implementation.
type PowerSupplyMetrics struct {
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
	// EnergykWh shall contain the total energy, in kilowatt-hours units, for this unit, that represents the 'Total'
	// ElectricalContext sensor when multiple energy sensors exist. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value
	// 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// FanSpeedPercent shall contain the fan speed, in percent units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Percent'.
	FanSpeedPercent SensorFanExcerpt
	// FrequencyHz shall contain the frequency, in hertz units, for this power supply.
	FrequencyHz SensorExcerpt
	// InputCurrentAmps shall contain the input current, in ampere units, for this power supply. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Current'.
	InputCurrentAmps SensorCurrentExcerpt
	// InputPowerWatts shall contain the input power, in watt units, for this power supply. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Power'.
	InputPowerWatts SensorPowerExcerpt
	// InputVoltage shall contain the input voltage, in volt units, for this power supply. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'.
	InputVoltage SensorVoltageExcerpt
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutputPowerWatts shall contain the total output power, in watt units, for this power supply. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Power'.
	OutputPowerWatts SensorPowerExcerpt
	// RailCurrentAmps shall contain the output currents, in ampere units, for this power supply. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Current'. The sensors shall appear in the same array order as the OutputRails property in
	// the associated PowerSupply resource.
	RailCurrentAmps []SensorCurrentExcerpt
	// RailCurrentAmps@odata.count
	RailCurrentAmpsCount int `json:"RailCurrentAmps@odata.count"`
	// RailPowerWatts shall contain the output power readings, in watt units, for this power supply. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Power'. The sensors shall appear in the same array order as the OutputRails property in
	// the associated PowerSupply resource.
	RailPowerWatts []SensorPowerExcerpt
	// RailPowerWatts@odata.count
	RailPowerWattsCount int `json:"RailPowerWatts@odata.count"`
	// RailVoltage shall contain the output voltages, in volt units, for this power supply. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. The sensors shall appear in the same array order as the OutputRails property in
	// the associated PowerSupply resource.
	RailVoltage []SensorVoltageExcerpt
	// RailVoltage@odata.count
	RailVoltageCount int `json:"RailVoltage@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TemperatureCelsius shall contain the temperature, in degrees Celsius units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
}

// UnmarshalJSON unmarshals a PowerSupplyMetrics object from the raw JSON.
func (powersupplymetrics *PowerSupplyMetrics) UnmarshalJSON(b []byte) error {
	type temp PowerSupplyMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powersupplymetrics = PowerSupplyMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetPowerSupplyMetrics will get a PowerSupplyMetrics instance from the service.
func GetPowerSupplyMetrics(c common.Client, uri string) (*PowerSupplyMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var powersupplymetrics PowerSupplyMetrics
	err = json.NewDecoder(resp.Body).Decode(&powersupplymetrics)
	if err != nil {
		return nil, err
	}

	powersupplymetrics.SetClient(c)
	return &powersupplymetrics, nil
}

// ListReferencedPowerSupplyMetricss gets the collection of PowerSupplyMetrics from
// a provided reference.
func ListReferencedPowerSupplyMetricss(c common.Client, link string) ([]*PowerSupplyMetrics, error) {
	var result []*PowerSupplyMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, powersupplymetricsLink := range links.ItemLinks {
		powersupplymetrics, err := GetPowerSupplyMetrics(c, powersupplymetricsLink)
		if err != nil {
			collectionError.Failures[powersupplymetricsLink] = err
		} else {
			result = append(result, powersupplymetrics)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
