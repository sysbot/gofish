//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CircuitType is
type CircuitType string

const (
	// MainsCircuitType A mains input or utility circuit.
	MainsCircuitType CircuitType = "Mains"
	// BranchCircuitType A branch (output) circuit.
	BranchCircuitType CircuitType = "Branch"
	// SubfeedCircuitType A subfeed (output) circuit.
	SubfeedCircuitType CircuitType = "Subfeed"
	// FeederCircuitType A feeder (output) circuit.
	FeederCircuitType CircuitType = "Feeder"
	// BusCircuitType An electrical bus circuit.
	BusCircuitType CircuitType = "Bus"
)

// VoltageType is
type VoltageType string

const (
	// ACVoltageType Alternating Current (AC) circuit.
	ACVoltageType VoltageType = "AC"
	// DCVoltageType Direct Current (DC) circuit.
	DCVoltageType VoltageType = "DC"
)

// Circuit shall be used to represent an electrical circuit for a Redfish implementation.
type Circuit struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// BreakerState shall contain the state of the over current protection device.
	BreakerState BreakerStates
	// CircuitType shall contain the type of circuit.
	CircuitType CircuitType
	// ConfigurationLocked shall indicate whether modification requests to this resource are not permitted. If 'true',
	// services shall reject modification requests to other properties in this resource.
	ConfigurationLocked string
	// CriticalCircuit shall indicate whether the circuit is designated as a critical circuit, and therefore is
	// excluded from autonomous logic that could affect the state of the circuit. The value shall be 'true' if the
	// circuit is deemed critical, and 'false' if the circuit is not critical.
	CriticalCircuit bool
	// CurrentAmps shall contain the current, in ampere units, for this single phase circuit. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Current'. This property shall not appear in resource instances representing poly-phase
	// circuits.
	CurrentAmps SensorCurrentExcerpt
	// Description provides a description of this resource.
	Description string
	// ElectricalConsumerNames shall contain an array of user-assigned identifying strings that describe downstream
	// devices that are powered by this circuit.
	ElectricalConsumerNames []string
	// ElectricalContext shall contain the combination of current-carrying conductors that distribute power.
	ElectricalContext ElectricalContext
	// ElectricalSourceManagerURI shall contain a URI to the management application or device that provides monitoring
	// or control of the upstream electrical source that provide power to this circuit. If a value has not been
	// assigned by a user, the value of this property shall be an empty string.
	ElectricalSourceManagerURI string
	// ElectricalSourceName shall contain a string that identifies the upstream electrical source, such as the name of
	// a circuit or outlet, that provides power to this circuit. If a value has not been assigned by a user, the value
	// of this property shall be an empty string.
	ElectricalSourceName string
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this circuit, that represents the 'Total'
	// ElectricalContext sensor when multiple energy sensors exist for this circuit. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// FrequencyHz shall contain the frequency, in hertz units, for this circuit. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Frequency'.
	FrequencyHz SensorExcerpt
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// NominalVoltage shall contain the nominal voltage for this circuit, in Volts.
	NominalVoltage NominalVoltageType
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhaseWiringType shall contain the number of ungrounded current-carrying conductors (phases) and the total number
	// of conductors (wires).
	PhaseWiringType PhaseWiringType
	// PlugType shall contain the type of physical plug used for this circuit, as defined by IEC, NEMA, or regional
	// standard.
	PlugType PlugType
	// PolyPhaseCurrentAmps shall contain the current sensors for this circuit. For single phase circuits this property
	// shall contain a duplicate copy of the current sensor referenced in the CurrentAmps property, if present. For
	// poly-phase circuits this property should contain multiple current sensor readings used to fully describe the
	// circuit.
	PolyPhaseCurrentAmps CurrentSensors
	// PolyPhaseEnergykWh shall contain the energy sensors for this circuit. For single phase circuits this property
	// shall contain a duplicate copy of the energy sensor referenced in the EnergykWh property, if present. For poly-
	// phase circuits this property should contain multiple energy sensor readings used to fully describe the circuit.
	PolyPhaseEnergykWh EnergySensors
	// PolyPhasePowerWatts shall contain the power sensors for this circuit. For single phase circuits this property
	// shall contain a duplicate copy of the power sensor referenced in the PowerWatts property, if present. For poly-
	// phase circuits this property should contain multiple power sensor readings used to fully describe the circuit.
	PolyPhasePowerWatts PowerSensors
	// PolyPhaseVoltage shall contain the voltage sensors for this circuit. For single phase circuits this property
	// shall contain a duplicate copy of the voltage sensor referenced in the Voltage property, if present. For poly-
	// phase circuits this property should contain multiple voltage sensor readings used to fully describe the circuit.
	PolyPhaseVoltage VoltageSensors
	// PowerControlLocked shall indicate whether requests to the PowerControl action are locked. If 'true', services
	// shall reject requests to the PowerControl action.
	PowerControlLocked string
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on after a PowerControl action to
	// cycle power. The value '0' shall indicate no delay to power on.
	PowerCycleDelaySeconds float64
	// PowerEnabled shall indicate the power enable state of the circuit. The value 'true' shall indicate that the
	// circuit can be powered on, and 'false' shall indicate that the circuit cannot be powered.
	PowerEnabled bool
	// PowerLoadPercent shall contain the power load, in percent units, for this circuit, that represents the 'Total'
	// ElectricalContext for this circuit.
	PowerLoadPercent SensorExcerpt
	// PowerOffDelaySeconds shall contain the number of seconds to delay power off after a PowerControl action. The
	// value '0' shall indicate no delay to power off.
	PowerOffDelaySeconds float64
	// PowerOnDelaySeconds shall contain the number of seconds to delay power up after a power cycle or a PowerControl
	// action. The value '0' shall indicate no delay to power up.
	PowerOnDelaySeconds float64
	// PowerRestoreDelaySeconds shall contain the number of seconds to delay power on after a power fault. The value
	// '0' shall indicate no delay to power on.
	PowerRestoreDelaySeconds float64
	// PowerRestorePolicy shall contain the desired PowerState of the circuit when power is applied. The value
	// 'LastState' shall return the circuit to the PowerState it was in when power was lost.
	PowerRestorePolicy string
	// PowerState shall contain the power state of the circuit.
	PowerState PowerState
	// PowerStateInTransition shall indicate whether the PowerState property will undergo a transition between on and
	// off states due to a configured delay. The transition may be due to the configuration of the power on, off, or
	// restore delay properties. If 'true', the PowerState property will transition at the conclusion of a configured
	// delay.
	PowerStateInTransition string
	// PowerWatts shall contain the total power, in watt units, for this circuit, that represents the 'Total'
	// ElectricalContext sensor when multiple power sensors exist for this circuit. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'.
	PowerWatts SensorPowerExcerpt
	// RatedCurrentAmps shall contain the rated maximum current for this circuit, in ampere units, after any required
	// de-rating, due to safety agency or other regulatory requirements, has been applied.
	RatedCurrentAmps float64
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UnbalancedCurrentPercent shall contain the current imbalance, in percent units, between phases in a poly-phase
	// circuit. The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Percent'.
	UnbalancedCurrentPercent SensorExcerpt
	// UnbalancedVoltagePercent shall contain the voltage imbalance, in percent units, between phases in a poly-phase
	// circuit. The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Percent'.
	UnbalancedVoltagePercent SensorExcerpt
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
	// Voltage shall contain the voltage, in volt units, for this single phase circuit. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Voltage'. This property shall not appear in resource instances representing poly-phase circuits.
	Voltage SensorVoltageExcerpt
	// VoltageType shall contain the type of voltage applied to the circuit.
	VoltageType VoltageType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Circuit object from the raw JSON.
func (circuit *Circuit) UnmarshalJSON(b []byte) error {
	type temp Circuit
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*circuit = Circuit(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	circuit.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (circuit *Circuit) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Circuit)
	original.UnmarshalJSON(circuit.rawData)

	readWriteFields := []string{
		"ConfigurationLocked",
		"CriticalCircuit",
		"ElectricalConsumerNames",
		"ElectricalSourceManagerURI",
		"ElectricalSourceName",
		"LocationIndicatorActive",
		"PowerControlLocked",
		"PowerCycleDelaySeconds",
		"PowerOffDelaySeconds",
		"PowerOnDelaySeconds",
		"PowerRestoreDelaySeconds",
		"PowerRestorePolicy",
		"UserLabel",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(circuit).Elem()

	return circuit.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetCircuit will get a Circuit instance from the service.
func GetCircuit(c common.Client, uri string) (*Circuit, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var circuit Circuit
	err = json.NewDecoder(resp.Body).Decode(&circuit)
	if err != nil {
		return nil, err
	}

	circuit.SetClient(c)
	return &circuit, nil
}

// ListReferencedCircuits gets the collection of Circuit from
// a provided reference.
func ListReferencedCircuits(c common.Client, link string) ([]*Circuit, error) {
	var result []*Circuit
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, circuitLink := range links.ItemLinks {
		circuit, err := GetCircuit(c, circuitLink)
		if err != nil {
			collectionError.Failures[circuitLink] = err
		} else {
			result = append(result, circuit)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// CurrentSensors shall contain properties that describe current sensor readings for a circuit.
type CurrentSensors struct {
	// Line1 shall contain the line current, in ampere units, for L1. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value 'Current'.
	// This property shall not be present if the equipment does not include an L1 measurement.
	Line1 SensorCurrentExcerpt
	// Line2 shall contain the line current, in ampere units, for L2. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value 'Current'.
	// This property shall not be present if the equipment does not include an L2 measurement.
	Line2 SensorCurrentExcerpt
	// Line3 shall contain the line current, in ampere units, for L3. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value 'Current'.
	// This property shall not be present if the equipment does not include an L3 measurement.
	Line3 SensorCurrentExcerpt
	// Neutral shall contain the line current, in ampere units, for the Neutral line. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Current'. This property shall not be present if the equipment does not include a Neutral line
	// measurement.
	Neutral SensorCurrentExcerpt
}

// UnmarshalJSON unmarshals a CurrentSensors object from the raw JSON.
func (currentsensors *CurrentSensors) UnmarshalJSON(b []byte) error {
	type temp CurrentSensors
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*currentsensors = CurrentSensors(t.temp)

	// Extract the links to other entities for later

	return nil
}

// EnergySensors shall contain properties that describe energy sensor readings for a circuit.
type EnergySensors struct {
	// Line1ToLine2 shall contain the energy, in kilowatt-hour units, between L1 and L2. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'EnergykWh'. This property shall not be present if the equipment does not include an L1-L2 measurement.
	Line1ToLine2 SensorEnergykWhExcerpt
	// Line1ToNeutral shall contain the energy, in kilowatt-hour units, between L1 and Neutral. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'EnergykWh'. This property shall not be present if the equipment does not include an
	// L1-Neutral measurement.
	Line1ToNeutral SensorEnergykWhExcerpt
	// Line2ToLine3 shall contain the energy, in kilowatt-hour units, between L2 and L3. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'EnergykWh'. This property shall not be present if the equipment does not include an L2-L3 measurement.
	Line2ToLine3 SensorEnergykWhExcerpt
	// Line2ToNeutral shall contain the energy, in kilowatt-hour units, between L2 and Neutral. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'EnergykWh'. This property shall not be present if the equipment does not include an
	// L2-Neutral measurement.
	Line2ToNeutral SensorEnergykWhExcerpt
	// Line3ToLine1 shall contain the energy, in kilowatt-hour units, between L3 and L1. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'EnergykWh'. This property shall not be present if the equipment does not include an L3-L1 measurement.
	Line3ToLine1 SensorEnergykWhExcerpt
	// Line3ToNeutral shall contain the energy, in kilowatt-hour units, between L3 and Neutral. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'EnergykWh'. This property shall not be present if the equipment does not include an
	// L3-Neutral measurement.
	Line3ToNeutral SensorEnergykWhExcerpt
}

// UnmarshalJSON unmarshals a EnergySensors object from the raw JSON.
func (energysensors *EnergySensors) UnmarshalJSON(b []byte) error {
	type temp EnergySensors
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*energysensors = EnergySensors(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// BranchCircuit shall contain a link to a resource of type Circuit that represents the branch circuit associated
	// with this circuit.
	BranchCircuit Circuit
	// DistributionCircuits shall contain an array of links to resources of type Circuit that represent the circuits
	// powered by this circuit.
	DistributionCircuits []Circuit
	// DistributionCircuits@odata.count
	DistributionCircuitsCount int `json:"DistributionCircuits@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Outlets shall contain an array of links to resources of type Outlet that represent the outlets associated with
	// this circuit.
	Outlets []Outlet
	// Outlets@odata.count
	OutletsCount int `json:"Outlets@odata.count"`
	// PowerOutlet shall contain a link to a resource of type Outlet that represents the outlet that provides power to
	// this circuit.
	PowerOutlet Outlet
	// SourceCircuit shall contain a link to a resource of type Circuit that represents the circuit that provides power
	// to this circuit. This property should be used when the power source is not represented by an Outlet resource,
	// such as a feeder circuit.
	SourceCircuit Circuit
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

// PowerSensors shall contain properties that describe power sensor readings for a circuit.
type PowerSensors struct {
	// Line1ToLine2 shall contain the power, in watt units, between L1 and L2. The value of the DataSourceUri property,
	// if present, shall reference a resource of type Sensor with the ReadingType property containing the value
	// 'Power'. This property shall not be present if the equipment does not include an L1-L2 measurement.
	Line1ToLine2 SensorPowerExcerpt
	// Line1ToNeutral shall contain the power, in watt units, between L1 and Neutral. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'. This property shall not be present if the equipment does not include an L1-Neutral measurement.
	Line1ToNeutral SensorPowerExcerpt
	// Line2ToLine3 shall contain the power, in watt units, between L2 and L3. The value of the DataSourceUri property,
	// if present, shall reference a resource of type Sensor with the ReadingType property containing the value
	// 'Power'. This property shall not be present if the equipment does not include an L2-L3 measurement.
	Line2ToLine3 SensorPowerExcerpt
	// Line2ToNeutral shall contain the power, in watt units, between L2 and Neutral. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'. This property shall not be present if the equipment does not include an L2-Neutral measurement.
	Line2ToNeutral SensorPowerExcerpt
	// Line3ToLine1 shall contain the power, in watt units, between L3 and L1. The value of the DataSourceUri property,
	// if present, shall reference a resource of type Sensor with the ReadingType property containing the value
	// 'Power'. This property shall not be present if the equipment does not include an L3-L1 measurement.
	Line3ToLine1 SensorPowerExcerpt
	// Line3ToNeutral shall contain the power, in watt units, between L3 and Neutral. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'. This property shall not be present if the equipment does not include an L3-Neutral measurement.
	Line3ToNeutral SensorPowerExcerpt
}

// UnmarshalJSON unmarshals a PowerSensors object from the raw JSON.
func (powersensors *PowerSensors) UnmarshalJSON(b []byte) error {
	type temp PowerSensors
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powersensors = PowerSensors(t.temp)

	// Extract the links to other entities for later

	return nil
}

// VoltageSensors shall contain properties that describe voltage sensor readings for a circuit.
type VoltageSensors struct {
	// Line1ToLine2 shall contain the line-to-line voltage, in volt units, between L1 and L2. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an L1-L2
	// measurement.
	Line1ToLine2 SensorVoltageExcerpt
	// Line1ToNeutral shall contain the line-to-line voltage, in volt units, between L1 and Neutral. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an
	// L1-Neutral measurement.
	Line1ToNeutral SensorVoltageExcerpt
	// Line2ToLine3 shall contain the line-to-line voltage, in volt units, between L2 and L3. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an L2-L3
	// measurement.
	Line2ToLine3 SensorVoltageExcerpt
	// Line2ToNeutral shall contain the line-to-line voltage, in volt units, between L2 and Neutral. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an
	// L2-Neutral measurement.
	Line2ToNeutral SensorVoltageExcerpt
	// Line3ToLine1 shall contain the line-to-line voltage, in volt units, between L3 and L1. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an L3-L1
	// measurement.
	Line3ToLine1 SensorVoltageExcerpt
	// Line3ToNeutral shall contain the line-to-line voltage, in volt units, between L3 and Neutral. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an
	// L3-Neutral measurement.
	Line3ToNeutral SensorVoltageExcerpt
}

// UnmarshalJSON unmarshals a VoltageSensors object from the raw JSON.
func (voltagesensors *VoltageSensors) UnmarshalJSON(b []byte) error {
	type temp VoltageSensors
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*voltagesensors = VoltageSensors(t.temp)

	// Extract the links to other entities for later

	return nil
}
