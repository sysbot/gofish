//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ReadingUnits is
type ReadingUnits string

const (
	// RPMReadingUnits The fan reading and thresholds are measured in revolutions per minute.
	RPMReadingUnits ReadingUnits = "RPM"
	// PercentReadingUnits The fan reading and thresholds are measured as a percentage.
	PercentReadingUnits ReadingUnits = "Percent"
)

// Fan
type Fan struct {
	common.Entity
	// Actions shall contain the available actions for this resource.
	Actions string
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// HotPluggable shall indicate whether the device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Hot-pluggable devices can become operable without altering
	// the operational state of the underlying equipment. Devices that cannot be inserted or removed from equipment in
	// operation, or devices that cannot become operable without affecting the operational state of that equipment,
	// shall be not hot-pluggable.
	HotPluggable bool
	// IndicatorLED shall contain the state of the indicator light associated with this fan.
	IndicatorLED IndicatorLED
	// Location shall contain location information of the associated fan.
	Location string
	// LowerThresholdCritical shall contain the value at which the Reading property is below the normal range but is
	// not yet fatal. The value of the property shall use the same units as the Reading property.
	LowerThresholdCritical int
	// LowerThresholdFatal shall contain the value at which the Reading property is below the normal range and is
	// fatal. The value of the property shall use the same units as the Reading property.
	LowerThresholdFatal int
	// LowerThresholdNonCritical shall contain the value at which the Reading property is below normal range. The value
	// of the property shall use the same units as the Reading property.
	LowerThresholdNonCritical int
	// Manufacturer shall contain the name of the organization responsible for producing the fan. This organization may
	// be the entity from whom the fan is purchased, but this is not necessarily true.
	Manufacturer string
	// MaxReadingRange shall indicate the highest possible value for the Reading property. The value of the property
	// shall use the same units as the Reading property.
	MaxReadingRange int
	// MemberId shall uniquely identify the member within the collection. For services supporting Redfish v1.6 or
	// higher, this value shall contain the zero-based array index.
	MemberId string
	// MinReadingRange shall indicate the lowest possible value for the Reading property. The value of the property
	// shall use the same units as the Reading property.
	MinReadingRange int
	// Model shall contain the model information as defined by the manufacturer for the associated fan.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for the associated fan.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region within the chassis with which this
	// fan is associated.
	PhysicalContext string
	// Reading shall contain the fan sensor reading.
	Reading int
	// ReadingUnits shall contain the units in which the fan reading and thresholds are measured.
	ReadingUnits ReadingUnits
	// Redundancy shall contain an array of links to the redundancy groups to which this fan belongs.
	Redundancy []Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// RelatedItem shall contain an array of links to resources or objects that this fan services.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall contain a numerical identifier for this fan speed sensor that is unique within this resource.
	SensorNumber int
	// SerialNumber shall contain the serial number as defined by the manufacturer for the associated fan.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for the
	// associated fan.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UpperThresholdCritical shall contain the value at which the Reading property is above the normal range but is
	// not yet fatal. The value of the property shall use the same units as the Reading property.
	UpperThresholdCritical int
	// UpperThresholdFatal shall contain the value at which the Reading property is above the normal range and is
	// fatal. The value of the property shall use the same units as the Reading property.
	UpperThresholdFatal int
	// UpperThresholdNonCritical shall contain the value at which the Reading property is above the normal range. The
	// value of the property shall use the same units as the Reading property.
	UpperThresholdNonCritical int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Fan object from the raw JSON.
func (fan *Fan) UnmarshalJSON(b []byte) error {
	type temp Fan
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fan = Fan(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	fan.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (fan *Fan) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Fan)
	original.UnmarshalJSON(fan.rawData)

	readWriteFields := []string{
		"IndicatorLED",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fan).Elem()

	return fan.Entity.Update(originalElement, currentElement, readWriteFields)
}

// FanActions shall contain the available actions for this resource.
type FanActions struct {
	// Oem shall contain the available OEM-specific actions for this resource.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a FanActions object from the raw JSON.
func (fanactions *FanActions) UnmarshalJSON(b []byte) error {
	type temp FanActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fanactions = FanActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FanOemActions shall contain the available OEM-specific actions for this resource.
type FanOemActions struct {
}

// UnmarshalJSON unmarshals a FanOemActions object from the raw JSON.
func (fanoemactions *FanOemActions) UnmarshalJSON(b []byte) error {
	type temp FanOemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fanoemactions = FanOemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Temperature
type Temperature struct {
	common.Entity
	// Actions shall contain the available actions for this resource.
	Actions string
	// AdjustedMaxAllowableOperatingValue shall indicate the adjusted maximum allowable operating temperature for the
	// equipment monitored by this temperature sensor, as specified by a standards body, manufacturer, or a
	// combination, and adjusted based on environmental conditions present. For example, liquid inlet temperature can
	// be adjusted based on the available liquid pressure.
	AdjustedMaxAllowableOperatingValue int
	// AdjustedMinAllowableOperatingValue shall indicate the adjusted minimum allowable operating temperature for the
	// equipment monitored by this temperature sensor, as specified by a standards body, manufacturer, or a
	// combination, and adjusted based on environmental conditions present. For example, liquid inlet temperature can
	// be adjusted based on the available liquid pressure.
	AdjustedMinAllowableOperatingValue int
	// DeltaPhysicalContext shall contain a description of the affected device or region within the chassis to which
	// the DeltaReadingCelsius temperature measurement applies, relative to PhysicalContext.
	DeltaPhysicalContext string
	// DeltaReadingCelsius shall contain the delta of the values of the temperature readings across this sensor and the
	// sensor at DeltaPhysicalContext.
	DeltaReadingCelsius float64
	// LowerThresholdCritical shall contain the value at which the ReadingCelsius property is below the normal range
	// but is not yet fatal. The value of the property shall use the same units as the ReadingCelsius property.
	LowerThresholdCritical float64
	// LowerThresholdFatal shall contain the value at which the ReadingCelsius property is below the normal range and
	// is fatal. The value of the property shall use the same units as the ReadingCelsius property.
	LowerThresholdFatal float64
	// LowerThresholdNonCritical shall contain the value at which the ReadingCelsius property is below normal range.
	// The value of the property shall use the same units as the ReadingCelsius property.
	LowerThresholdNonCritical float64
	// LowerThresholdUser shall contain the value at which the ReadingCelsius property is below the user-defined range.
	// The value of the property shall use the same units as the ReadingCelsius property. The value shall be equal to
	// the value of LowerThresholdNonCritical, LowerThresholdCritical, or LowerThresholdFatal, unless set by a user.
	LowerThresholdUser int
	// MaxAllowableOperatingValue shall indicate the maximum allowable operating temperature for the equipment
	// monitored by this temperature sensor, as specified by a standards body, manufacturer, or a combination.
	MaxAllowableOperatingValue int
	// MaxReadingRangeTemp shall indicate the highest possible value for the ReadingCelsius property. The value of the
	// property shall use the same units as the ReadingCelsius property.
	MaxReadingRangeTemp float64
	// MemberId shall uniquely identify the member within the collection. For services supporting Redfish v1.6 or
	// higher, this value shall contain the zero-based array index.
	MemberId string
	// MinAllowableOperatingValue shall indicate the minimum allowable operating temperature for the equipment
	// monitored by this temperature sensor, as specified by a standards body, manufacturer, or a combination.
	MinAllowableOperatingValue int
	// MinReadingRangeTemp shall indicate the lowest possible value for the ReadingCelsius property. The value of the
	// property shall use the same units as the ReadingCelsius property.
	MinReadingRangeTemp float64
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalContext shall contain a description of the affected device or region within the chassis to which this
	// temperature applies.
	PhysicalContext string
	// ReadingCelsius shall contain the temperature in Celsius degrees.
	ReadingCelsius float64
	// RelatedItem shall contain an array of links to resources or objects that represent areas or devices to which
	// this temperature applies.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall contain a numerical identifier for this temperature sensor that is unique within this
	// resource.
	SensorNumber int
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UpperThresholdCritical shall contain the value at which the ReadingCelsius property is above the normal range
	// but is not yet fatal. The value of the property shall use the same units as the ReadingCelsius property.
	UpperThresholdCritical float64
	// UpperThresholdFatal shall contain the value at which the ReadingCelsius property is above the normal range and
	// is fatal. The value of the property shall use the same units as the ReadingCelsius property.
	UpperThresholdFatal float64
	// UpperThresholdNonCritical shall contain the value at which the ReadingCelsius property is above the normal
	// range. The value of the property shall use the same units as the ReadingCelsius property.
	UpperThresholdNonCritical float64
	// UpperThresholdUser shall contain the value at which the ReadingCelsius property is above the user-defined range.
	// The value of the property shall use the same units as the ReadingCelsius property. The value shall be equal to
	// the value of UpperThresholdNonCritical, UpperThresholdCritical, or UpperThresholdFatal, unless set by a user.
	UpperThresholdUser int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Temperature object from the raw JSON.
func (temperature *Temperature) UnmarshalJSON(b []byte) error {
	type temp Temperature
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*temperature = Temperature(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	temperature.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (temperature *Temperature) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Temperature)
	original.UnmarshalJSON(temperature.rawData)

	readWriteFields := []string{
		"LowerThresholdUser",
		"UpperThresholdUser",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(temperature).Elem()

	return temperature.Entity.Update(originalElement, currentElement, readWriteFields)
}

// TemperatureActions shall contain the available actions for this resource.
type TemperatureActions struct {
	// Oem shall contain the available OEM-specific actions for this resource.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a TemperatureActions object from the raw JSON.
func (temperatureactions *TemperatureActions) UnmarshalJSON(b []byte) error {
	type temp TemperatureActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*temperatureactions = TemperatureActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TemperatureOemActions shall contain the available OEM-specific actions for this resource.
type TemperatureOemActions struct {
}

// UnmarshalJSON unmarshals a TemperatureOemActions object from the raw JSON.
func (temperatureoemactions *TemperatureOemActions) UnmarshalJSON(b []byte) error {
	type temp TemperatureOemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*temperatureoemactions = TemperatureOemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Thermal shall contain the thermal management properties for temperature monitoring and management of cooling
// fans for a Redfish implementation.
type Thermal struct {
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
	// Fans shall contain the set of fans for this chassis.
	Fans []Fan
	// Fans@odata.count
	FansCount int `json:"Fans@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Redundancy shall contain redundancy information for the fans in this chassis.
	Redundancy []Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Temperatures shall contain the set of temperature sensors for this chassis.
	Temperatures []Temperature
	// Temperatures@odata.count
	TemperaturesCount int `json:"Temperatures@odata.count"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Thermal object from the raw JSON.
func (thermal *Thermal) UnmarshalJSON(b []byte) error {
	type temp Thermal
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermal = Thermal(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	thermal.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (thermal *Thermal) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Thermal)
	original.UnmarshalJSON(thermal.rawData)

	readWriteFields := []string{
		"Fans",
		"Temperatures",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(thermal).Elem()

	return thermal.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetThermal will get a Thermal instance from the service.
func GetThermal(c common.Client, uri string) (*Thermal, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var thermal Thermal
	err = json.NewDecoder(resp.Body).Decode(&thermal)
	if err != nil {
		return nil, err
	}

	thermal.SetClient(c)
	return &thermal, nil
}

// ListReferencedThermals gets the collection of Thermal from
// a provided reference.
func ListReferencedThermals(c common.Client, link string) ([]*Thermal, error) {
	var result []*Thermal
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, thermalLink := range links.ItemLinks {
		thermal, err := GetThermal(c, thermalLink)
		if err != nil {
			collectionError.Failures[thermalLink] = err
		} else {
			result = append(result, thermal)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// ThermalActions shall contain the available actions for this resource.
type ThermalActions struct {
	// Oem shall contain the available OEM-specific actions for this resource.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a ThermalActions object from the raw JSON.
func (thermalactions *ThermalActions) UnmarshalJSON(b []byte) error {
	type temp ThermalActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermalactions = ThermalActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ThermalOemActions shall contain the available OEM-specific actions for this resource.
type ThermalOemActions struct {
}

// UnmarshalJSON unmarshals a ThermalOemActions object from the raw JSON.
func (thermaloemactions *ThermalOemActions) UnmarshalJSON(b []byte) error {
	type temp ThermalOemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermaloemactions = ThermalOemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}
