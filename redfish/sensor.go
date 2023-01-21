//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ImplementationType is
type ImplementationType string

const (
	// PhysicalSensorImplementationType The reading is acquired from a physical sensor.
	PhysicalSensorImplementationType ImplementationType = "PhysicalSensor"
	// SynthesizedImplementationType The reading is obtained by applying a calculation on one or more properties or
	// multiple sensors. The calculation is not provided.
	SynthesizedImplementationType ImplementationType = "Synthesized"
	// ReportedImplementationType The reading is obtained from software or a device.
	ReportedImplementationType ImplementationType = "Reported"
)

// ReadingType is
type ReadingType string

const (
	// TemperatureReadingType shall indicate a temperature measurement, in degrees Celsius units, and the ReadingUnits
	// value shall be 'Cel'.
	TemperatureReadingType ReadingType = "Temperature"
	// HumidityReadingType shall indicate a relative humidity measurement, in percent units, and the ReadingUnits value
	// shall be '%'.
	HumidityReadingType ReadingType = "Humidity"
	// PowerReadingType shall indicate the arithmetic mean of product terms of instantaneous voltage and current values
	// measured over integer number of line cycles for a circuit, in watt units, and the ReadingUnits value shall be
	// 'W'.
	PowerReadingType ReadingType = "Power"
	// EnergykWhReadingType shall indicate the energy, integral of real power over time, of the monitored item. If
	// representing metered power consumption the value shall reflect the power consumption since the sensor metrics
	// were last reset. The value of the Reading property shall be in kilowatt-hour units and the ReadingUnits value
	// shall be 'kW.h'. This value is used for large-scale energy consumption measurements, while EnergyJoules and
	// EnergyWh are used for device-level consumption measurements.
	EnergykWhReadingType ReadingType = "EnergykWh"
	// EnergyJoulesReadingType shall indicate the energy, integral of real power over time, of the monitored item. If
	// representing metered power consumption the value shall reflect the power consumption since the sensor metrics
	// were last reset. The value of the Reading property shall be in joule units and the ReadingUnits value shall be
	// 'J'. This value is used for device-level energy consumption measurements, while EnergykWh is used for large-
	// scale consumption measurements.
	EnergyJoulesReadingType ReadingType = "EnergyJoules"
	// EnergyWhReadingType shall indicate the energy, integral of real power over time, of the monitored item. If
	// representing metered power consumption the value shall reflect the power consumption since the sensor metrics
	// were last reset. The value of the Reading property shall be in watt-hour units and the ReadingUnits value shall
	// be 'W.h'. This value is used for device-level energy consumption measurements, while EnergykWh is used for
	// large-scale consumption measurements.
	EnergyWhReadingType ReadingType = "EnergyWh"
	// ChargeAhReadingType shall indicate the amount of charge of the monitored item. If representing metered power
	// consumption, integral of real power over time, the value shall reflect the power consumption since the sensor
	// metrics were last reset. The value of the Reading property shall be in ampere-hour units and the ReadingUnits
	// value shall be 'A.h'.
	ChargeAhReadingType ReadingType = "ChargeAh"
	// VoltageReadingType shall indicate a measurement of the root mean square (RMS) of instantaneous voltage
	// calculated over an integer number of line cycles for a circuit. Voltage is expressed in volt units and the
	// ReadingUnits value shall be 'V'.
	VoltageReadingType ReadingType = "Voltage"
	// CurrentReadingType shall indicate a measurement of the root mean square (RMS) of instantaneous current
	// calculated over an integer number of line cycles for a circuit. Current is expressed in ampere units and the
	// ReadingUnits value shall be 'A'.
	CurrentReadingType ReadingType = "Current"
	// FrequencyReadingType shall indicate a frequency measurement, in hertz units, and the ReadingUnits value shall be
	// 'Hz'.
	FrequencyReadingType ReadingType = "Frequency"
	// PressureReadingType shall indicate a measurement of force, in pascal units, applied perpendicular to the surface
	// of an object per unit area over which that force is distributed. The ReadingUnits shall be 'Pa'.
	PressureReadingType ReadingType = "Pressure"
	// PressurekPaReadingType shall indicate a measurement of pressure, in kilopascal units, relative to atmospheric
	// pressure. The ReadingUnits value shall be 'kPa'.
	PressurekPaReadingType ReadingType = "PressurekPa"
	// LiquidLevelReadingType shall indicate a measurement of fluid height, in centimeter units, relative to a
	// specified vertical datum and the ReadingUnits value shall be 'cm'.
	LiquidLevelReadingType ReadingType = "LiquidLevel"
	// RotationalReadingType shall indicate a measurement of rotational frequency, in revolutions per minute unit, and
	// the ReadingUnits value shall be either '{rev}/min', which is preferred, or 'RPM', which is a deprecated value.
	RotationalReadingType ReadingType = "Rotational"
	// AirFlowReadingType shall indicate a measurement of a volume of gas per unit of time, in cubic feet per minute
	// units, that flows through a particular junction. The ReadingUnits shall be '[ft_i]3/min'.
	AirFlowReadingType ReadingType = "AirFlow"
	// LiquidFlowReadingType shall indicate a measurement of a volume of liquid per unit of time, in liters per second
	// units, that flows through a particular junction. The ReadingUnits shall be 'L/s'.
	LiquidFlowReadingType ReadingType = "LiquidFlow"
	// BarometricReadingType shall indicate a measurement of barometric pressure, in millimeters of a mercury column,
	// and the ReadingUnits value shall be 'mm[Hg]'.
	BarometricReadingType ReadingType = "Barometric"
	// AltitudeReadingType shall indicate a measurement of altitude, in meter units, defined as the elevation above sea
	// level. The ReadingUnits value shall be 'm'.
	AltitudeReadingType ReadingType = "Altitude"
	// PercentReadingType shall indicate a percentage measurement, in percent units, and the ReadingUnits value shall
	// be '%'.
	PercentReadingType ReadingType = "Percent"
	// AbsoluteHumidityReadingType shall indicate an absolute (volumetric) humidity measurement, in grams per cubic
	// meter units, and the ReadingUnits value shall be 'g/m3'.
	AbsoluteHumidityReadingType ReadingType = "AbsoluteHumidity"
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
	// EitherThresholdActivation This threshold is activated when either the increasing or decreasing conditions are
	// met.
	EitherThresholdActivation ThresholdActivation = "Either"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// AssociatedControls shall contain an array of links to resources of type Control that represent the controls that
	// can affect this sensor.
	AssociatedControls []Control
	// AssociatedControls@odata.count
	AssociatedControlsCount int `json:"AssociatedControls@odata.count"`
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

// Sensor shall represent a sensor for a Redfish implementation.
type Sensor struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Accuracy shall contain the percent error +/- of the measured versus actual values of the Reading property.
	Accuracy float64
	// Actions shall contain the available actions for this resource.
	Actions string
	// AdjustedMaxAllowableOperatingValue shall contain the adjusted maximum allowable operating value for the
	// equipment that this sensor monitors, as specified by a standards body, manufacturer, or both. The value is
	// adjusted based on environmental conditions. For example, liquid inlet temperature can be adjusted based on the
	// available liquid pressure.
	AdjustedMaxAllowableOperatingValue float64
	// AdjustedMinAllowableOperatingValue shall contain the adjusted minimum allowable operating value for the
	// equipment that this sensor monitors, as specified by a standards body, manufacturer, or both. This value is
	// adjusted based on environmental conditions. For example, liquid inlet temperature can be adjusted based on the
	// available liquid pressure.
	AdjustedMinAllowableOperatingValue float64
	// ApparentVA shall contain the product of voltage (RMS) multiplied by current (RMS) for a circuit. This property
	// can appear in sensors of the Power ReadingType, and shall not appear in sensors of other ReadingType values.
	ApparentVA float64
	// ApparentkVAh shall contain the apparent energy, in kilovolt-ampere-hour units, for an electrical energy
	// measurement. This property can appear in sensors with a ReadingType containing 'EnergykWh', and shall not appear
	// in sensors with other ReadingType values.
	ApparentkVAh float64
	// AverageReading shall contain the average sensor value over the time specified by the value of the
	// AveragingInterval property. The value shall be reset by the ResetMetrics action.
	AverageReading float64
	// AveragingInterval shall contain the interval over which the sensor value is averaged to produce the value of the
	// AverageReading property. This property shall only be present if the AverageReading property is present.
	AveragingInterval string
	// AveragingIntervalAchieved shall indicate that enough readings were collected to calculate the AverageReading
	// value over the interval specified by the AveragingInterval property. The value shall be reset by the
	// ResetMetrics action. This property shall only be present if the AveragingInterval property is present.
	AveragingIntervalAchieved bool
	// Calibration shall contain the offset applied to the raw sensor value to provide a calibrated value for the
	// sensor as returned by the Reading property. The value of this property shall follow the units of the Reading
	// property for this sensor instance. Updating the value of this property shall not affect the value of the
	// CalibrationTime property.
	Calibration float64
	// CalibrationTime shall contain the date and time that the sensor was last calibrated. This property is intended
	// to reflect the actual time the calibration occurred.
	CalibrationTime string
	// CrestFactor shall contain the ratio of the peak measurement divided by the RMS measurement and calculated over
	// same N line cycles. A sine wave would have a value of 1.414.
	CrestFactor float64
	// Description provides a description of this resource.
	Description string
	// ElectricalContext shall represent the combination of current-carrying conductors that distribute power.
	ElectricalContext ElectricalContext
	// Implementation shall contain the implementation of the sensor.
	Implementation ImplementationType
	// LifetimeReading shall contain the total accumulation of the Reading property over the sensor's life time. This
	// value shall not be reset by the ResetMetrics action.
	LifetimeReading float64
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall indicate the location information for this sensor.
	Location string
	// LowestReading shall contain the lowest sensor value since the last ResetMetrics action was performed or the
	// service last reset the time-based property values.
	LowestReading float64
	// LowestReadingTime shall contain the date and time when the lowest sensor value was observed.
	LowestReadingTime string
	// MaxAllowableOperatingValue shall contain the maximum allowable operating value for the equipment that this
	// sensor monitors, as specified by a standards body, manufacturer, or both.
	MaxAllowableOperatingValue float64
	// MinAllowableOperatingValue shall contain the minimum allowable operating value for the equipment that this
	// sensor monitors, as specified by a standards body, manufacturer, or both.
	MinAllowableOperatingValue float64
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PeakReading shall contain the peak sensor value since the last ResetMetrics action was performed or the service
	// last reset the time-based property values.
	PeakReading float64
	// PeakReadingTime shall contain the date and time when the peak sensor value was observed.
	PeakReadingTime string
	// PhaseAngleDegrees shall contain the phase angle, in degree units, between the current and voltage waveforms for
	// an electrical measurement. This property can appear in sensors with a ReadingType containing 'Power', and shall
	// not appear in sensors with other ReadingType values.
	PhaseAngleDegrees float64
	// PhysicalContext shall contain a description of the affected component or region within the equipment to which
	// this sensor measurement applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region within the equipment to which this
	// sensor measurement applies. This property generally differentiates multiple sensors within the same
	// PhysicalContext instance.
	PhysicalSubContext PhysicalSubContext
	// PowerFactor shall identify the quotient of real power (W) and apparent power (VA) for a circuit. PowerFactor is
	// expressed in unit-less 1/100ths. This property can appear in sensors containing a ReadingType value of 'Power',
	// and shall not appear in sensors of other ReadingType values.
	PowerFactor float64
	// Precision shall contain the number of significant digits in the Reading property.
	Precision float64
	// ReactiveVAR shall contain the arithmetic mean of product terms of instantaneous voltage and quadrature current
	// measurements calculated over an integer number of line cycles for a circuit. This property can appear in sensors
	// of the Power ReadingType, and shall not appear in sensors of other ReadingType values.
	ReactiveVAR float64
	// ReactivekVARh shall contain the reactive energy, in kilovolt-ampere-hours (reactive) units, for an electrical
	// energy measurement. This property can appear in sensors with a ReadingType containing 'EnergykWh', and shall not
	// appear in sensors with other ReadingType values.
	ReactivekVARh float64
	// Reading shall contain the sensor value.
	Reading float64
	// ReadingRangeMax shall indicate the maximum possible value of the Reading property for this sensor. This value is
	// the range of valid readings for this sensor. Values outside this range are discarded as reading errors.
	ReadingRangeMax float64
	// ReadingRangeMin shall indicate the minimum possible value of the Reading property for this sensor. This value is
	// the range of valid readings for this sensor. Values outside this range are discarded as reading errors.
	ReadingRangeMin float64
	// ReadingTime shall contain the date and time that the reading data was acquired from the sensor. This value is
	// used to synchronize readings from multiple sensors, and does not represent the time at which the resource was
	// accessed.
	ReadingTime string
	// ReadingType shall contain the type of the sensor.
	ReadingType ReadingType
	// ReadingUnits shall contain the units of the sensor's reading and thresholds.
	ReadingUnits string
	// RelatedItem shall contain an array of links to resources or objects that this sensor services.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensingInterval shall contain the time interval between readings of data from the sensor.
	SensingInterval string
	// SensorGroup shall contain information for a group of sensors that provide input for the value of this sensor's
	// reading. If this property is present, the Implementation property shall contain the value 'Synthesized'. The
	// group may be created for redundancy or to improve the accuracy of the reading through multiple sensor inputs.
	SensorGroup string
	// SensorResetTime shall contain the date and time when the ResetMetrics action was last performed or the service
	// last reset the time-based property values.
	SensorResetTime string
	// SpeedRPM shall contain a reading of the rotational speed of the device in revolutions per minute (RPM) units.
	SpeedRPM float64
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// THDPercent shall contain the total harmonic distortion of the Reading property in percent units.
	THDPercent float64
	// Thresholds shall contain the set of thresholds that derive a sensor's health and operational range.
	Thresholds string
	// VoltageType shall represent the type of input voltage the sensor monitors.
	VoltageType VoltageType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Sensor object from the raw JSON.
func (sensor *Sensor) UnmarshalJSON(b []byte) error {
	type temp Sensor
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensor = Sensor(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	sensor.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (sensor *Sensor) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Sensor)
	original.UnmarshalJSON(sensor.rawData)

	readWriteFields := []string{
		"AveragingInterval",
		"Calibration",
		"CalibrationTime",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(sensor).Elem()

	return sensor.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSensor will get a Sensor instance from the service.
func GetSensor(c common.Client, uri string) (*Sensor, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sensor Sensor
	err = json.NewDecoder(resp.Body).Decode(&sensor)
	if err != nil {
		return nil, err
	}

	sensor.SetClient(c)
	return &sensor, nil
}

// ListReferencedSensors gets the collection of Sensor from
// a provided reference.
func ListReferencedSensors(c common.Client, link string) ([]*Sensor, error) {
	var result []*Sensor
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, sensorLink := range links.ItemLinks {
		sensor, err := GetSensor(c, sensorLink)
		if err != nil {
			collectionError.Failures[sensorLink] = err
		} else {
			result = append(result, sensor)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// SensorArrayExcerpt shall represent a sensor for a Redfish implementation.
type SensorArrayExcerpt struct {
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceUri string
	// DeviceName shall contain the name of the device associated with this sensor. If the device is represented by a
	// resource, the value shall contain the value of the Name property of the associated resource.
	DeviceName string
	// PhysicalContext shall contain a description of the affected component or region within the equipment to which
	// this sensor measurement applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region within the equipment to which this
	// sensor measurement applies. This property generally differentiates multiple sensors within the same
	// PhysicalContext instance.
	PhysicalSubContext PhysicalSubContext
	// Reading shall contain the sensor value.
	Reading float64
}

// UnmarshalJSON unmarshals a SensorArrayExcerpt object from the raw JSON.
func (sensorarrayexcerpt *SensorArrayExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorArrayExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensorarrayexcerpt = SensorArrayExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SensorCurrentExcerpt shall represent a sensor for a Redfish implementation.
type SensorCurrentExcerpt struct {
	// CrestFactor shall contain the ratio of the peak measurement divided by the RMS measurement and calculated over
	// same N line cycles. A sine wave would have a value of 1.414.
	CrestFactor float64
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceUri string
	// Reading shall contain the sensor value.
	Reading float64
	// THDPercent shall contain the total harmonic distortion of the Reading property in percent units.
	THDPercent float64
}

// UnmarshalJSON unmarshals a SensorCurrentExcerpt object from the raw JSON.
func (sensorcurrentexcerpt *SensorCurrentExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorCurrentExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensorcurrentexcerpt = SensorCurrentExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SensorEnergykWhExcerpt shall represent a sensor for a Redfish implementation.
type SensorEnergykWhExcerpt struct {
	// ApparentkVAh shall contain the apparent energy, in kilovolt-ampere-hour units, for an electrical energy
	// measurement. This property can appear in sensors with a ReadingType containing 'EnergykWh', and shall not appear
	// in sensors with other ReadingType values.
	ApparentkVAh float64
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceUri string
	// LifetimeReading shall contain the total accumulation of the Reading property over the sensor's life time. This
	// value shall not be reset by the ResetMetrics action.
	LifetimeReading float64
	// ReactivekVARh shall contain the reactive energy, in kilovolt-ampere-hours (reactive) units, for an electrical
	// energy measurement. This property can appear in sensors with a ReadingType containing 'EnergykWh', and shall not
	// appear in sensors with other ReadingType values.
	ReactivekVARh float64
	// Reading shall contain the sensor value.
	Reading float64
	// SensorResetTime shall contain the date and time when the ResetMetrics action was last performed or the service
	// last reset the time-based property values.
	SensorResetTime string
}

// UnmarshalJSON unmarshals a SensorEnergykWhExcerpt object from the raw JSON.
func (sensorenergykwhexcerpt *SensorEnergykWhExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorEnergykWhExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensorenergykwhexcerpt = SensorEnergykWhExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SensorExcerpt shall represent a sensor for a Redfish implementation.
type SensorExcerpt struct {
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceUri string
	// Reading shall contain the sensor value.
	Reading float64
}

// UnmarshalJSON unmarshals a SensorExcerpt object from the raw JSON.
func (sensorexcerpt *SensorExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensorexcerpt = SensorExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SensorFanArrayExcerpt shall represent a sensor for a Redfish implementation.
type SensorFanArrayExcerpt struct {
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceUri string
	// DeviceName shall contain the name of the device associated with this sensor. If the device is represented by a
	// resource, the value shall contain the value of the Name property of the associated resource.
	DeviceName string
	// PhysicalContext shall contain a description of the affected component or region within the equipment to which
	// this sensor measurement applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region within the equipment to which this
	// sensor measurement applies. This property generally differentiates multiple sensors within the same
	// PhysicalContext instance.
	PhysicalSubContext PhysicalSubContext
	// Reading shall contain the sensor value.
	Reading float64
	// SpeedRPM shall contain a reading of the rotational speed of the device in revolutions per minute (RPM) units.
	SpeedRPM float64
}

// UnmarshalJSON unmarshals a SensorFanArrayExcerpt object from the raw JSON.
func (sensorfanarrayexcerpt *SensorFanArrayExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorFanArrayExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensorfanarrayexcerpt = SensorFanArrayExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SensorFanExcerpt shall represent a sensor for a Redfish implementation.
type SensorFanExcerpt struct {
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceUri string
	// Reading shall contain the sensor value.
	Reading float64
	// SpeedRPM shall contain a reading of the rotational speed of the device in revolutions per minute (RPM) units.
	SpeedRPM float64
}

// UnmarshalJSON unmarshals a SensorFanExcerpt object from the raw JSON.
func (sensorfanexcerpt *SensorFanExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorFanExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensorfanexcerpt = SensorFanExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SensorPowerArrayExcerpt shall represent a sensor for a Redfish implementation.
type SensorPowerArrayExcerpt struct {
	// ApparentVA shall contain the product of voltage (RMS) multiplied by current (RMS) for a circuit. This property
	// can appear in sensors of the Power ReadingType, and shall not appear in sensors of other ReadingType values.
	ApparentVA float64
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceUri string
	// PhaseAngleDegrees shall contain the phase angle, in degree units, between the current and voltage waveforms for
	// an electrical measurement. This property can appear in sensors with a ReadingType containing 'Power', and shall
	// not appear in sensors with other ReadingType values.
	PhaseAngleDegrees float64
	// PhysicalContext shall contain a description of the affected component or region within the equipment to which
	// this sensor measurement applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region within the equipment to which this
	// sensor measurement applies. This property generally differentiates multiple sensors within the same
	// PhysicalContext instance.
	PhysicalSubContext PhysicalSubContext
	// PowerFactor shall identify the quotient of real power (W) and apparent power (VA) for a circuit. PowerFactor is
	// expressed in unit-less 1/100ths. This property can appear in sensors containing a ReadingType value of 'Power',
	// and shall not appear in sensors of other ReadingType values.
	PowerFactor float64
	// ReactiveVAR shall contain the arithmetic mean of product terms of instantaneous voltage and quadrature current
	// measurements calculated over an integer number of line cycles for a circuit. This property can appear in sensors
	// of the Power ReadingType, and shall not appear in sensors of other ReadingType values.
	ReactiveVAR float64
	// Reading shall contain the sensor value.
	Reading float64
}

// UnmarshalJSON unmarshals a SensorPowerArrayExcerpt object from the raw JSON.
func (sensorpowerarrayexcerpt *SensorPowerArrayExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorPowerArrayExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensorpowerarrayexcerpt = SensorPowerArrayExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SensorPowerExcerpt shall represent a sensor for a Redfish implementation.
type SensorPowerExcerpt struct {
	// ApparentVA shall contain the product of voltage (RMS) multiplied by current (RMS) for a circuit. This property
	// can appear in sensors of the Power ReadingType, and shall not appear in sensors of other ReadingType values.
	ApparentVA float64
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceUri string
	// PhaseAngleDegrees shall contain the phase angle, in degree units, between the current and voltage waveforms for
	// an electrical measurement. This property can appear in sensors with a ReadingType containing 'Power', and shall
	// not appear in sensors with other ReadingType values.
	PhaseAngleDegrees float64
	// PowerFactor shall identify the quotient of real power (W) and apparent power (VA) for a circuit. PowerFactor is
	// expressed in unit-less 1/100ths. This property can appear in sensors containing a ReadingType value of 'Power',
	// and shall not appear in sensors of other ReadingType values.
	PowerFactor float64
	// ReactiveVAR shall contain the arithmetic mean of product terms of instantaneous voltage and quadrature current
	// measurements calculated over an integer number of line cycles for a circuit. This property can appear in sensors
	// of the Power ReadingType, and shall not appear in sensors of other ReadingType values.
	ReactiveVAR float64
	// Reading shall contain the sensor value.
	Reading float64
}

// UnmarshalJSON unmarshals a SensorPowerExcerpt object from the raw JSON.
func (sensorpowerexcerpt *SensorPowerExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorPowerExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensorpowerexcerpt = SensorPowerExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SensorVoltageExcerpt shall represent a sensor for a Redfish implementation.
type SensorVoltageExcerpt struct {
	// CrestFactor shall contain the ratio of the peak measurement divided by the RMS measurement and calculated over
	// same N line cycles. A sine wave would have a value of 1.414.
	CrestFactor float64
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceUri string
	// Reading shall contain the sensor value.
	Reading float64
	// THDPercent shall contain the total harmonic distortion of the Reading property in percent units.
	THDPercent float64
}

// UnmarshalJSON unmarshals a SensorVoltageExcerpt object from the raw JSON.
func (sensorvoltageexcerpt *SensorVoltageExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorVoltageExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensorvoltageexcerpt = SensorVoltageExcerpt(t.temp)

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
	// use the same units as the Reading property.
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

// Thresholds shall contain the set of thresholds that derive a sensor's health and operational range.
type Thresholds struct {
	// LowerCaution shall contain the value at which the Reading property is below normal range. The value of the
	// property shall use the same units as the Reading property.
	LowerCaution string
	// LowerCautionUser shall contain the value at which the Reading property is below normal range. The value of the
	// property shall use the same units as the Reading property.
	LowerCautionUser string
	// LowerCritical shall contain the value at which the Reading property is below the normal range but is not yet
	// fatal. The value of the property shall use the same units as the Reading property.
	LowerCritical string
	// LowerCriticalUser shall contain the value at which the Reading property is below the normal range but is not yet
	// fatal. The value of the property shall use the same units as the Reading property.
	LowerCriticalUser string
	// LowerFatal shall contain the value at which the Reading property is below the normal range and is fatal. The
	// value of the property shall use the same units as the Reading property.
	LowerFatal string
	// UpperCaution shall contain the value at which the Reading property is above the normal range. The value of the
	// property shall use the same units as the Reading property.
	UpperCaution string
	// UpperCautionUser shall contain the value at which the Reading property is above the normal range. The value of
	// the property shall use the same units as the Reading property.
	UpperCautionUser string
	// UpperCritical shall contain the value at which the Reading property is above the normal range but is not yet
	// fatal. The value of the property shall use the same units as the Reading property.
	UpperCritical string
	// UpperCriticalUser shall contain the value at which the Reading property is above the normal range but is not yet
	// fatal. The value of the property shall use the same units as the Reading property.
	UpperCriticalUser string
	// UpperFatal shall contain the value at which the Reading property is above the normal range and is fatal. The
	// value of the property shall use the same units as the Reading property.
	UpperFatal string
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
