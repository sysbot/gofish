//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// BatteryMetrics shall be used to represent the metrics of a battery unit for a Redfish implementation.
type BatteryMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// CellVoltages shall contain the cell voltages, in volt units, for this battery. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Voltage'.
	CellVoltages []SensorVoltageExcerpt
	// CellVoltages@odata.count
	CellVoltagesCount int `json:"CellVoltages@odata.count"`
	// ChargePercent shall contain the amount of charge available, in percent units, in this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Percent'.
	ChargePercent string
	// Description provides a description of this resource.
	Description string
	// DischargeCycles shall contain the number of discharges this battery sustained.
	DischargeCycles float64
	// InputCurrentAmps shall contain the input current, in ampere units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Current'.
	InputCurrentAmps string
	// InputVoltage shall contain the input voltage, in volt units, for this battery. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Voltage'.
	InputVoltage string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutputCurrentAmps shall contain the output currents, in ampere units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Current'. The sensors shall appear in the same array order as the OutputVoltages property.
	OutputCurrentAmps []SensorCurrentExcerpt
	// OutputCurrentAmps@odata.count
	OutputCurrentAmpsCount int `json:"OutputCurrentAmps@odata.count"`
	// OutputVoltages shall contain the output voltages, in volt units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. The sensors shall appear in the same array order as the OutputCurrentAmps
	// property.
	OutputVoltages []SensorVoltageExcerpt
	// OutputVoltages@odata.count
	OutputVoltagesCount int `json:"OutputVoltages@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StoredChargeAmpHours shall contain the stored charge, in ampere-hours units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'ChargeAh'.
	StoredChargeAmpHours string
	// StoredEnergyWattHours shall contain the stored energy, in watt-hour units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'EnergyWh'.
	StoredEnergyWattHours string
	// TemperatureCelsius shall contain the temperature, in degrees Celsius units, for this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Temperature'.
	TemperatureCelsius string
}

// UnmarshalJSON unmarshals a BatteryMetrics object from the raw JSON.
func (batterymetrics *BatteryMetrics) UnmarshalJSON(b []byte) error {
	type temp BatteryMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*batterymetrics = BatteryMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetBatteryMetrics will get a BatteryMetrics instance from the service.
func GetBatteryMetrics(c common.Client, uri string) (*BatteryMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var batterymetrics BatteryMetrics
	err = json.NewDecoder(resp.Body).Decode(&batterymetrics)
	if err != nil {
		return nil, err
	}

	batterymetrics.SetClient(c)
	return &batterymetrics, nil
}

// ListReferencedBatteryMetricss gets the collection of BatteryMetrics from
// a provided reference.
func ListReferencedBatteryMetricss(c common.Client, link string) ([]*BatteryMetrics, error) {
	var result []*BatteryMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, batterymetricsLink := range links.ItemLinks {
		batterymetrics, err := GetBatteryMetrics(c, batterymetricsLink)
		if err != nil {
			collectionError.Failures[batterymetricsLink] = err
		} else {
			result = append(result, batterymetrics)
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
