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

// PowerDistributionMetrics shall be used to represent the metrics of a power distribution component or unit for a
// Redfish implementation.
type PowerDistributionMetrics struct {
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
	// EnergykWh shall contain the total energy, in kilowatt-hours, for this resource, that represents the 'Total'
	// ElectricalContext sensor when multiple energy sensors exist. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value
	// 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// HumidityPercent shall contain the humidity, in percent units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Humidity'.
	HumidityPercent SensorExcerpt
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerLoadPercent shall contain the power load, in percent units, for this device, that represents the 'Total'
	// ElectricalContext for this device. The value of the DataSourceUri property, if present, shall reference a
	// resource of type Sensor with the ReadingType property containing the value 'Percent'.
	PowerLoadPercent SensorExcerpt
	// PowerWatts shall contain the total power, in watt units, for this resource, that represents the 'Total'
	// ElectricalContext sensor when multiple power sensors exist. The value of the DataSourceUri property, if present,
	// shall reference a resource of type Sensor with the ReadingType property containing the value 'Power'.
	PowerWatts SensorPowerExcerpt
	// TemperatureCelsius shall contain the temperature, in degrees Celsius units, for this resource. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
}

// UnmarshalJSON unmarshals a PowerDistributionMetrics object from the raw JSON.
func (powerdistributionmetrics *PowerDistributionMetrics) UnmarshalJSON(b []byte) error {
	type temp PowerDistributionMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powerdistributionmetrics = PowerDistributionMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetPowerDistributionMetrics will get a PowerDistributionMetrics instance from the service.
func GetPowerDistributionMetrics(c common.Client, uri string) (*PowerDistributionMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var powerdistributionmetrics PowerDistributionMetrics
	err = json.NewDecoder(resp.Body).Decode(&powerdistributionmetrics)
	if err != nil {
		return nil, err
	}

	powerdistributionmetrics.SetClient(c)
	return &powerdistributionmetrics, nil
}

// ListReferencedPowerDistributionMetricss gets the collection of PowerDistributionMetrics from
// a provided reference.
func ListReferencedPowerDistributionMetricss(c common.Client, link string) ([]*PowerDistributionMetrics, error) {
	var result []*PowerDistributionMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, powerdistributionmetricsLink := range links.ItemLinks {
		powerdistributionmetrics, err := GetPowerDistributionMetrics(c, powerdistributionmetricsLink)
		if err != nil {
			collectionError.Failures[powerdistributionmetricsLink] = err
		} else {
			result = append(result, powerdistributionmetrics)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
