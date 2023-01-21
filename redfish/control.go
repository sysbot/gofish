//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ControlMode is
type ControlMode string

const (
	// AutomaticControlMode Automatically adjust control to meet the set point.
	AutomaticControlMode ControlMode = "Automatic"
	// OverrideControlMode User override of the automatic set point value.
	OverrideControlMode ControlMode = "Override"
	// ManualControlMode No automatic adjustments are made to the control.
	ManualControlMode ControlMode = "Manual"
	// DisabledControlMode The control has been disabled.
	DisabledControlMode ControlMode = "Disabled"
)

// ControlType is
type ControlType string

const (
	// TemperatureControlType shall indicate a control used to regulate temperature, in units of degrees Celsius,
	// either to a single set point or within a range, and the SetPointUnits property shall contain 'Cel'.
	TemperatureControlType ControlType = "Temperature"
	// PowerControlType shall indicate a control used to regulate or limit maximum power consumption, in watt units,
	// either to a single set point or within a range, and the SetPointUnits property shall contain 'W'.
	PowerControlType ControlType = "Power"
	// FrequencyControlType shall indicate a control used to limit the operating frequency, in hertz units, of a
	// device, either to a single set point or within a range, and the SetPointUnits property shall contain 'Hz'.
	FrequencyControlType ControlType = "Frequency"
	// FrequencyMHzControlType shall indicate a control used to limit the operating frequency, in megahertz units, of a
	// device, either to a single set point or within a range, and the SetPointUnits property shall contain 'MHz'.
	FrequencyMHzControlType ControlType = "FrequencyMHz"
	// PressureControlType shall indicate a control used to adjust pressure in a system, in kilopascal units, and the
	// SetPointUnits property shall contain 'kPa'.
	PressureControlType ControlType = "Pressure"
)

// ImplementationType is
type ImplementationType string

const (
	// ProgrammableImplementationType The set point can be adjusted through this interface.
	ProgrammableImplementationType ImplementationType = "Programmable"
	// DirectImplementationType The set point directly affects the control value.
	DirectImplementationType ImplementationType = "Direct"
	// MonitoredImplementationType A physical control that cannot be adjusted through this interface.
	MonitoredImplementationType ImplementationType = "Monitored"
)

// SetPointType is
type SetPointType string

const (
	// SingleSetPointType shall indicate the control utilizes a single set point for its operation. The SetPoint
	// property shall be present for this control type. The SettingMin and SettingMax properties shall not be present
	// for this control type.
	SingleSetPointType SetPointType = "Single"
	// RangeSetPointType shall indicate the control utilizes a set point range for its operation. The SettingMin and
	// SettingMax properties shall be present for this control type. The SetPoint property shall not be present for
	// this control type.
	RangeSetPointType SetPointType = "Range"
)

// Control shall represent a control point for a Redfish implementation.
type Control struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Accuracy shall contain the percent error of the measured versus actual values of the SetPoint property.
	Accuracy float64
	// Actions shall contain the available actions for this resource.
	Actions string
	// AllowableMax shall indicate the maximum possible value of the SetPoint or SettingMax properties for this
	// control. Services shall not accept values for SetPoint or SettingMax above this value.
	AllowableMax float64
	// AllowableMin shall indicate the minimum possible value of the SetPoint or SettingMin properties for this
	// control. Services shall not accept values for SetPoint or SettingMin below this value.
	AllowableMin float64
	// AllowableNumericValues shall contain the supported values for this control. The units shall follow the value of
	// SetPointUnits. This property should only be present when the set point or range has a limited set of supported
	// values that cannot be accurately described using the Increment property.
	AllowableNumericValues []string
	// AssociatedSensors shall contain an array of links to resources of type Sensor that represent the sensors related
	// to this control.
	AssociatedSensors []Sensor
	// AssociatedSensors@odata.count
	AssociatedSensorsCount int `json:"AssociatedSensors@odata.count"`
	// ControlDelaySeconds shall contain the time in seconds that will elapse after the control value deviates above or
	// below the value of SetPoint before the control will activate.
	ControlDelaySeconds float64
	// ControlLoop shall contain the details for the control loop described by this resource.
	ControlLoop ControlLoop
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// ControlType shall contain the type of the control.
	ControlType ControlType
	// DeadBand shall contain the maximum deviation value allowed above or below the value of SetPoint before the
	// control will activate.
	DeadBand float64
	// Description provides a description of this resource.
	Description string
	// Implementation shall contain the implementation of the control.
	Implementation ImplementationType
	// Increment shall contain the smallest change allowed to the value of the SetPoint, SettingMin, or SettingMax
	// properties. The units shall follow the value of SetPointUnits.
	Increment float64
	// Location shall indicate the location information for this control.
	Location string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalContext shall contain a description of the affected component or region within the equipment to which
	// this control applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region within the equipment to which this
	// control applies. This property generally differentiates multiple controls within the same PhysicalContext
	// instance.
	PhysicalSubContext PhysicalSubContext
	// RelatedItem shall contain an array of links to resources that this control services.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// Sensor shall contain the Sensor excerpt directly associated with this control. The value of the DataSourceUri
	// property shall reference a resource of type Sensor. This property shall not be present if multiple sensors are
	// associated with a single control.
	Sensor SensorExcerpt
	// SetPoint shall contain the desired set point control value. The units shall follow the value of SetPointUnits.
	SetPoint float64
	// SetPointType shall contain the type of set point definitions used to describe this control.
	SetPointType SetPointType
	// SetPointUnits shall contain the units of the control's set point.
	SetPointUnits string
	// SetPointUpdateTime shall contain the date and time that the value of SetPoint was last changed.
	SetPointUpdateTime string
	// SettingMax shall contain the maximum desired set point within the acceptable range. The service shall reject
	// values greater than the value of AllowableMax. The units shall follow the value of SetPointUnits.
	SettingMax float64
	// SettingMin shall contain the minimum desired set point within the acceptable range. The service shall reject
	// values less than the value of AllowableMin. The units shall follow the value of SetPointUnits.
	SettingMin float64
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Control object from the raw JSON.
func (control *Control) UnmarshalJSON(b []byte) error {
	type temp Control
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*control = Control(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	control.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (control *Control) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Control)
	original.UnmarshalJSON(control.rawData)

	readWriteFields := []string{
		"ControlDelaySeconds",
		"ControlMode",
		"DeadBand",
		"SetPoint",
		"SettingMax",
		"SettingMin",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(control).Elem()

	return control.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetControl will get a Control instance from the service.
func GetControl(c common.Client, uri string) (*Control, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var control Control
	err = json.NewDecoder(resp.Body).Decode(&control)
	if err != nil {
		return nil, err
	}

	control.SetClient(c)
	return &control, nil
}

// ListReferencedControls gets the collection of Control from
// a provided reference.
func ListReferencedControls(c common.Client, link string) ([]*Control, error) {
	var result []*Control
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, controlLink := range links.ItemLinks {
		control, err := GetControl(c, controlLink)
		if err != nil {
			collectionError.Failures[controlLink] = err
		} else {
			result = append(result, control)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// ControlExcerpt shall represent a control point for a Redfish implementation.
type ControlExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the SetPoint or SettingMax properties for this
	// control. Services shall not accept values for SetPoint or SettingMax above this value.
	AllowableMax float64
	// AllowableMin shall indicate the minimum possible value of the SetPoint or SettingMin properties for this
	// control. Services shall not accept values for SetPoint or SettingMin below this value.
	AllowableMin float64
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy. If no source resource is implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceUri string
	// Reading shall contain the value of the Reading property of the Sensor resource directly associated with this
	// control. This property shall not be present if multiple sensors are associated with a single control.
	Reading float64
	// ReadingUnits shall contain the units of the sensor's reading and thresholds. This property shall not be present
	// if multiple sensors are associated with a single control.
	ReadingUnits string
}

// UnmarshalJSON unmarshals a ControlExcerpt object from the raw JSON.
func (controlexcerpt *ControlExcerpt) UnmarshalJSON(b []byte) error {
	type temp ControlExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controlexcerpt = ControlExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ControlLoop shall describe the details of a control loop.
type ControlLoop struct {
	// CoefficientUpdateTime shall contain the date and time that any of the coefficients for the control loop were
	// last changed.
	CoefficientUpdateTime string
	// Differential shall contain the coefficient for the differential factor in a control loop.
	Differential float64
	// Integral shall contain the coefficient for the integral factor in a control loop.
	Integral float64
	// Proportional shall contain the coefficient for the proportional factor in a control loop.
	Proportional float64
}

// UnmarshalJSON unmarshals a ControlLoop object from the raw JSON.
func (controlloop *ControlLoop) UnmarshalJSON(b []byte) error {
	type temp ControlLoop
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controlloop = ControlLoop(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ControlRangeExcerpt shall represent a control point for a Redfish implementation.
type ControlRangeExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the SetPoint or SettingMax properties for this
	// control. Services shall not accept values for SetPoint or SettingMax above this value.
	AllowableMax float64
	// AllowableMin shall indicate the minimum possible value of the SetPoint or SettingMin properties for this
	// control. Services shall not accept values for SetPoint or SettingMin below this value.
	AllowableMin float64
	// AllowableNumericValues shall contain the supported values for this control. The units shall follow the value of
	// SetPointUnits. This property should only be present when the set point or range has a limited set of supported
	// values that cannot be accurately described using the Increment property.
	AllowableNumericValues []string
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy. If no source resource is implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceUri string
	// Reading shall contain the value of the Reading property of the Sensor resource directly associated with this
	// control. This property shall not be present if multiple sensors are associated with a single control.
	Reading float64
	// ReadingUnits shall contain the units of the sensor's reading and thresholds. This property shall not be present
	// if multiple sensors are associated with a single control.
	ReadingUnits string
	// SettingMax shall contain the maximum desired set point within the acceptable range. The service shall reject
	// values greater than the value of AllowableMax. The units shall follow the value of SetPointUnits.
	SettingMax float64
	// SettingMin shall contain the minimum desired set point within the acceptable range. The service shall reject
	// values less than the value of AllowableMin. The units shall follow the value of SetPointUnits.
	SettingMin float64
}

// UnmarshalJSON unmarshals a ControlRangeExcerpt object from the raw JSON.
func (controlrangeexcerpt *ControlRangeExcerpt) UnmarshalJSON(b []byte) error {
	type temp ControlRangeExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controlrangeexcerpt = ControlRangeExcerpt(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ControlSingleExcerpt shall represent a control point for a Redfish implementation.
type ControlSingleExcerpt struct {
	// AllowableMax shall indicate the maximum possible value of the SetPoint or SettingMax properties for this
	// control. Services shall not accept values for SetPoint or SettingMax above this value.
	AllowableMax float64
	// AllowableMin shall indicate the minimum possible value of the SetPoint or SettingMin properties for this
	// control. Services shall not accept values for SetPoint or SettingMin below this value.
	AllowableMin float64
	// ControlMode shall contain the operating mode of the control.
	ControlMode ControlMode
	// DataSourceUri shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy. If no source resource is implemented, meaning the excerpt represents the only available data, this
	// property shall not be present.
	DataSourceUri string
	// Reading shall contain the value of the Reading property of the Sensor resource directly associated with this
	// control. This property shall not be present if multiple sensors are associated with a single control.
	Reading float64
	// ReadingUnits shall contain the units of the sensor's reading and thresholds. This property shall not be present
	// if multiple sensors are associated with a single control.
	ReadingUnits string
	// SetPoint shall contain the desired set point control value. The units shall follow the value of SetPointUnits.
	SetPoint float64
}

// UnmarshalJSON unmarshals a ControlSingleExcerpt object from the raw JSON.
func (controlsingleexcerpt *ControlSingleExcerpt) UnmarshalJSON(b []byte) error {
	type temp ControlSingleExcerpt
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controlsingleexcerpt = ControlSingleExcerpt(t.temp)

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
