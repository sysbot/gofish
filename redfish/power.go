//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// InputType is
type InputType string

const (
	// ACInputType Alternating Current (AC) input range.
	ACInputType InputType = "AC"
	// DCInputType Direct Current (DC) input range.
	DCInputType InputType = "DC"
)

// LineInputVoltageType is
type LineInputVoltageType string

const (
	// UnknownLineInputVoltageType The power supply line input voltage type cannot be determined.
	UnknownLineInputVoltageType LineInputVoltageType = "Unknown"
	// ACLowLineLineInputVoltageType 100-127V AC input.
	ACLowLineLineInputVoltageType LineInputVoltageType = "ACLowLine"
	// ACMidLineLineInputVoltageType 200-240V AC input.
	ACMidLineLineInputVoltageType LineInputVoltageType = "ACMidLine"
	// ACHighLineLineInputVoltageType 277V AC input.
	ACHighLineLineInputVoltageType LineInputVoltageType = "ACHighLine"
	// DCNeg48VLineInputVoltageType -48V DC input.
	DCNeg48VLineInputVoltageType LineInputVoltageType = "DCNeg48V"
	// DC380VLineInputVoltageType High Voltage DC input (380V).
	DC380VLineInputVoltageType LineInputVoltageType = "DC380V"
	// AC120VLineInputVoltageType AC 120V nominal input.
	AC120VLineInputVoltageType LineInputVoltageType = "AC120V"
	// AC240VLineInputVoltageType AC 240V nominal input.
	AC240VLineInputVoltageType LineInputVoltageType = "AC240V"
	// AC277VLineInputVoltageType AC 277V nominal input.
	AC277VLineInputVoltageType LineInputVoltageType = "AC277V"
	// ACandDCWideRangeLineInputVoltageType Wide range AC or DC input.
	ACandDCWideRangeLineInputVoltageType LineInputVoltageType = "ACandDCWideRange"
	// ACWideRangeLineInputVoltageType Wide range AC input.
	ACWideRangeLineInputVoltageType LineInputVoltageType = "ACWideRange"
	// DC240VLineInputVoltageType DC 240V nominal input.
	DC240VLineInputVoltageType LineInputVoltageType = "DC240V"
)

// PowerLimitException is
type PowerLimitException string

const (
	// NoActionPowerLimitException Take no action when the limit is exceeded.
	NoActionPowerLimitException PowerLimitException = "NoAction"
	// HardPowerOffPowerLimitException Turn the power off immediately when the limit is exceeded.
	HardPowerOffPowerLimitException PowerLimitException = "HardPowerOff"
	// LogEventOnlyPowerLimitException Log an event when the limit is exceeded, but take no further action.
	LogEventOnlyPowerLimitException PowerLimitException = "LogEventOnly"
	// OemPowerLimitException Take an OEM-defined action.
	OemPowerLimitException PowerLimitException = "Oem"
)

// PowerSupplyType is
type PowerSupplyType string

const (
	// UnknownPowerSupplyType The power supply type cannot be determined.
	UnknownPowerSupplyType PowerSupplyType = "Unknown"
	// ACPowerSupplyType Alternating Current (AC) power supply.
	ACPowerSupplyType PowerSupplyType = "AC"
	// DCPowerSupplyType Direct Current (DC) power supply.
	DCPowerSupplyType PowerSupplyType = "DC"
	// ACorDCPowerSupplyType The power supply supports both DC or AC.
	ACorDCPowerSupplyType PowerSupplyType = "ACorDC"
)

// InputRange shall describe an input range that the associated power supply can utilize.
type InputRange struct {
	// InputType shall contain the input type (AC or DC) of the associated range.
	InputType InputType
	// MaximumFrequencyHz shall contain the value, in Hertz, of the maximum line input frequency that the power supply
	// is capable of consuming for this range.
	MaximumFrequencyHz float64
	// MaximumVoltage shall contain the value, in volts, of the maximum line input voltage that the power supply is
	// capable of consuming for this range.
	MaximumVoltage float64
	// MinimumFrequencyHz shall contain the value, in Hertz, of the minimum line input frequency that the power supply
	// is capable of consuming for this range.
	MinimumFrequencyHz float64
	// MinimumVoltage shall contain the value, in volts, of the minimum line input voltage that the power supply is
	// capable of consuming for this range.
	MinimumVoltage float64
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutputWattage shall contain the maximum amount of power, in watts, that the associated power supply is rated to
	// deliver while operating in this input range.
	OutputWattage float64
}

// UnmarshalJSON unmarshals a InputRange object from the raw JSON.
func (inputrange *InputRange) UnmarshalJSON(b []byte) error {
	type temp InputRange
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*inputrange = InputRange(t.temp)

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

// Power shall contain the power metrics for a Redfish implementation.
type Power struct {
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
	// PowerControl shall contain the set of power control readings and settings.
	PowerControl []PowerControl
	// PowerControl@odata.count
	PowerControlCount int `json:"PowerControl@odata.count"`
	// PowerSupplies shall contain the set of power supplies associated with this system or device.
	PowerSupplies []PowerSupply
	// PowerSupplies@odata.count
	PowerSuppliesCount int `json:"PowerSupplies@odata.count"`
	// Redundancy shall contain redundancy information for the set of power supplies in this system or device.
	Redundancy []Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Voltages shall contain the set of voltage sensors for this chassis.
	Voltages []Voltage
	// Voltages@odata.count
	VoltagesCount int `json:"Voltages@odata.count"`
}

// UnmarshalJSON unmarshals a Power object from the raw JSON.
func (power *Power) UnmarshalJSON(b []byte) error {
	type temp Power
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*power = Power(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetPower will get a Power instance from the service.
func GetPower(c common.Client, uri string) (*Power, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var power Power
	err = json.NewDecoder(resp.Body).Decode(&power)
	if err != nil {
		return nil, err
	}

	power.SetClient(c)
	return &power, nil
}

// ListReferencedPowers gets the collection of Power from
// a provided reference.
func ListReferencedPowers(c common.Client, link string) ([]*Power, error) {
	var result []*Power
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, powerLink := range links.ItemLinks {
		power, err := GetPower(c, powerLink)
		if err != nil {
			collectionError.Failures[powerLink] = err
		} else {
			result = append(result, power)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// PowerControl
type PowerControl struct {
	common.Entity
	// Actions shall contain the available actions for this resource.
	Actions string
	// MemberId shall uniquely identify the member within the collection. For services supporting Redfish v1.6 or
	// higher, this value shall contain the zero-based array index.
	MemberId string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalContext shall contain a description of the affected device(s) or region within the chassis to which this
	// power control applies.
	PhysicalContext string
	// PowerAllocatedWatts shall represent the total power currently allocated or budgeted to the chassis.
	PowerAllocatedWatts float64
	// PowerAvailableWatts shall represent the amount of reserve power capacity, in watts, that remains. This value is
	// the PowerCapacityWatts value minus the PowerAllocatedWatts value.
	PowerAvailableWatts float64
	// PowerCapacityWatts shall represent the total power capacity that can be allocated to the chassis.
	PowerCapacityWatts float64
	// PowerConsumedWatts shall represent the actual power that the chassis consumes, in watts.
	PowerConsumedWatts float64
	// PowerLimit shall contain power limit status and configuration information for this chassis.
	PowerLimit string
	// PowerMetrics shall contain power metrics for power readings, such as interval, minimum, maximum, and average
	// power consumption, for the chassis.
	PowerMetrics string
	// PowerRequestedWatts shall represent the amount of power, in watts, that the chassis currently requests to be
	// budgeted for future use.
	PowerRequestedWatts float64
	// RelatedItem shall contain an array of links to resources or objects associated with this power limit.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a PowerControl object from the raw JSON.
func (powercontrol *PowerControl) UnmarshalJSON(b []byte) error {
	type temp PowerControl
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powercontrol = PowerControl(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PowerControlActions shall contain the available actions for this resource.
type PowerControlActions struct {
	// Oem shall contain the available OEM-specific actions for this resource.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a PowerControlActions object from the raw JSON.
func (powercontrolactions *PowerControlActions) UnmarshalJSON(b []byte) error {
	type temp PowerControlActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powercontrolactions = PowerControlActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PowerControlOemActions shall contain the available OEM-specific actions for this resource.
type PowerControlOemActions struct {
}

// UnmarshalJSON unmarshals a PowerControlOemActions object from the raw JSON.
func (powercontroloemactions *PowerControlOemActions) UnmarshalJSON(b []byte) error {
	type temp PowerControlOemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powercontroloemactions = PowerControlOemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PowerLimit shall contain power limit status and configuration information for this chassis.
type PowerLimit struct {
	// CorrectionInMs shall represent the time interval in ms required for the limiting process to react and reduce the
	// power consumption below the limit.
	CorrectionInMs int
	// LimitException shall represent the action to be taken if the resource power consumption cannot be limited below
	// the specified limit after several correction time periods.
	LimitException PowerLimitException
	// LimitInWatts shall represent the power capping limit, in watts, for the resource. If 'null', power capping shall
	// be disabled.
	LimitInWatts float64
}

// UnmarshalJSON unmarshals a PowerLimit object from the raw JSON.
func (powerlimit *PowerLimit) UnmarshalJSON(b []byte) error {
	type temp PowerLimit
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powerlimit = PowerLimit(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PowerMetric shall contain power metrics for power readings, such as interval, minimum, maximum, and average
// power consumption, for a resource.
type PowerMetric struct {
	// AverageConsumedWatts shall represent the average power level that occurred over the last IntervalInMin minutes.
	AverageConsumedWatts float64
	// IntervalInMin shall represent the time interval or window, in minutes, over which the power metrics are
	// measured.
	IntervalInMin int
	// MaxConsumedWatts shall represent the maximum power level, in watts, that occurred within the last IntervalInMin
	// minutes.
	MaxConsumedWatts float64
	// MinConsumedWatts shall represent the minimum power level, in watts, that occurred within the last IntervalInMin
	// minutes.
	MinConsumedWatts float64
}

// UnmarshalJSON unmarshals a PowerMetric object from the raw JSON.
func (powermetric *PowerMetric) UnmarshalJSON(b []byte) error {
	type temp PowerMetric
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powermetric = PowerMetric(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PowerSupply Details of a power supplies associated with this system or device.
type PowerSupply struct {
	common.Entity
	// Actions shall contain the available actions for this resource.
	Actions string
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// EfficiencyPercent shall contain the measured power efficiency, as a percentage, of the associated power supply.
	EfficiencyPercent float64
	// FirmwareVersion shall contain the firmware version as defined by the manufacturer for the associated power
	// supply.
	FirmwareVersion string
	// HotPluggable shall indicate whether the device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Devices indicated as hot-pluggable shall allow the device to
	// become operable without altering the operational state of the underlying equipment. Devices that cannot be
	// inserted or removed from equipment in operation, or devices that cannot become operable without affecting the
	// operational state of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// IndicatorLED shall contain the indicator light state for the indicator light associated with this power supply.
	IndicatorLED IndicatorLED
	// InputRanges shall contain a collection of ranges usable by the power supply unit.
	InputRanges []InputRange
	// LastPowerOutputWatts shall contain the average power output, measured in watts, of the associated power supply.
	LastPowerOutputWatts float64
	// LineInputVoltage shall contain the value in Volts of the line input voltage (measured or configured for) that
	// the power supply has been configured to operate with or is currently receiving.
	LineInputVoltage float64
	// LineInputVoltageType shall contain the type of input line voltage supported by the associated power supply.
	LineInputVoltageType LineInputVoltageType
	// Location shall contain location information of the associated power supply.
	Location string
	// Manufacturer shall contain the name of the organization responsible for producing the power supply. This
	// organization may be the entity from whom the power supply is purchased, but this is not necessarily true.
	Manufacturer string
	// MemberId shall uniquely identify the member within the collection. For services supporting Redfish v1.6 or
	// higher, this value shall contain the zero-based array index.
	MemberId string
	// Model shall contain the model information as defined by the manufacturer for the associated power supply.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for the associated power supply.
	PartNumber string
	// PowerCapacityWatts shall contain the maximum amount of power, in watts, that the associated power supply is
	// rated to deliver.
	PowerCapacityWatts float64
	// PowerInputWatts shall contain the measured input power, in watts, of the associated power supply.
	PowerInputWatts float64
	// PowerOutputWatts shall contain the measured output power, in watts, of the associated power supply.
	PowerOutputWatts float64
	// PowerSupplyType shall contain the input power type (AC or DC) of the associated power supply.
	PowerSupplyType PowerSupplyType
	// Redundancy shall contain an array of links to the redundancy groups to which this power supply belongs.
	Redundancy []Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// RelatedItem shall contain an array of links to resources or objects associated with this power supply.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SerialNumber shall contain the serial number as defined by the manufacturer for the associated power supply.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for the
	// associated power supply.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerSupply object from the raw JSON.
func (powersupply *PowerSupply) UnmarshalJSON(b []byte) error {
	type temp PowerSupply
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powersupply = PowerSupply(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	powersupply.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (powersupply *PowerSupply) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(PowerSupply)
	original.UnmarshalJSON(powersupply.rawData)

	readWriteFields := []string{
		"IndicatorLED",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(powersupply).Elem()

	return powersupply.Entity.Update(originalElement, currentElement, readWriteFields)
}

// PowerSupplyActions shall contain the available actions for this resource.
type PowerSupplyActions struct {
	// Oem shall contain the available OEM-specific actions for this resource.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a PowerSupplyActions object from the raw JSON.
func (powersupplyactions *PowerSupplyActions) UnmarshalJSON(b []byte) error {
	type temp PowerSupplyActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powersupplyactions = PowerSupplyActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PowerSupplyOemActions shall contain the available OEM-specific actions for this resource.
type PowerSupplyOemActions struct {
}

// UnmarshalJSON unmarshals a PowerSupplyOemActions object from the raw JSON.
func (powersupplyoemactions *PowerSupplyOemActions) UnmarshalJSON(b []byte) error {
	type temp PowerSupplyOemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powersupplyoemactions = PowerSupplyOemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Voltage
type Voltage struct {
	common.Entity
	// Actions shall contain the available actions for this resource.
	Actions string
	// LowerThresholdCritical shall contain the value at which the ReadingVolts property is below the normal range but
	// is not yet fatal. The value of the property shall use the same units as the ReadingVolts property.
	LowerThresholdCritical float64
	// LowerThresholdFatal shall contain the value at which the ReadingVolts property is below the normal range and is
	// fatal. The value of the property shall use the same units as the ReadingVolts property.
	LowerThresholdFatal float64
	// LowerThresholdNonCritical shall contain the value at which the ReadingVolts property is below normal range. The
	// value of the property shall use the same units as the ReadingVolts property.
	LowerThresholdNonCritical float64
	// MaxReadingRange shall indicate the highest possible value for the ReadingVolts property. The value of the
	// property shall use the same units as the ReadingVolts property.
	MaxReadingRange float64
	// MemberId shall uniquely identify the member within the collection. For services supporting Redfish v1.6 or
	// higher, this value shall contain the zero-based array index.
	MemberId string
	// MinReadingRange shall indicate the lowest possible value for the ReadingVolts property. The value of the
	// property shall use the same units as the ReadingVolts property.
	MinReadingRange float64
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalContext shall contain a description of the affected device or region within the chassis to which this
	// voltage measurement applies.
	PhysicalContext string
	// ReadingVolts shall contain the voltage sensor's reading.
	ReadingVolts float64
	// RelatedItem shall contain an array of links to resources or objects to which this voltage measurement applies.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall contain a numerical identifier for this voltage sensor that is unique within this resource.
	SensorNumber int
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UpperThresholdCritical shall contain the value at which the ReadingVolts property is above the normal range but
	// is not yet fatal. The value of the property shall use the same units as the ReadingVolts property.
	UpperThresholdCritical float64
	// UpperThresholdFatal shall contain the value at which the ReadingVolts property is above the normal range and is
	// fatal. The value of the property shall use the same units as the ReadingVolts property.
	UpperThresholdFatal float64
	// UpperThresholdNonCritical shall contain the value at which the ReadingVolts property is above the normal range.
	// The value of the property shall use the same units as the ReadingVolts property.
	UpperThresholdNonCritical float64
}

// UnmarshalJSON unmarshals a Voltage object from the raw JSON.
func (voltage *Voltage) UnmarshalJSON(b []byte) error {
	type temp Voltage
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*voltage = Voltage(t.temp)

	// Extract the links to other entities for later

	return nil
}

// VoltageActions shall contain the available actions for this resource.
type VoltageActions struct {
	// Oem shall contain the available OEM-specific actions for this resource.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a VoltageActions object from the raw JSON.
func (voltageactions *VoltageActions) UnmarshalJSON(b []byte) error {
	type temp VoltageActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*voltageactions = VoltageActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// VoltageOemActions shall contain the available OEM-specific actions for this resource.
type VoltageOemActions struct {
}

// UnmarshalJSON unmarshals a VoltageOemActions object from the raw JSON.
func (voltageoemactions *VoltageOemActions) UnmarshalJSON(b []byte) error {
	type temp VoltageOemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*voltageoemactions = VoltageOemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}
