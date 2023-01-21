//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// LineStatus is
type LineStatus string

const (
	// NormalLineStatus Line input is within normal operating range.
	NormalLineStatus LineStatus = "Normal"
	// LossOfInputLineStatus No power detected at line input.
	LossOfInputLineStatus LineStatus = "LossOfInput"
	// OutOfRangeLineStatus Line input voltage or current is outside of normal operating range.
	OutOfRangeLineStatus LineStatus = "OutOfRange"
)

// PowerSupplyType is
type PowerSupplyType string

const (
	// ACPowerSupplyType Alternating Current (AC) power supply.
	ACPowerSupplyType PowerSupplyType = "AC"
	// DCPowerSupplyType Direct Current (DC) power supply.
	DCPowerSupplyType PowerSupplyType = "DC"
	// ACorDCPowerSupplyType The power supply supports both DC or AC.
	ACorDCPowerSupplyType PowerSupplyType = "ACorDC"
	// DCRegulatorPowerSupplyType Direct Current (DC) voltage regulator.
	DCRegulatorPowerSupplyType PowerSupplyType = "DCRegulator"
)

// EfficiencyRating shall describe an efficiency rating for a power supply.
type EfficiencyRating struct {
	// EfficiencyPercent shall contain the rated efficiency, as a percentage, of this power supply at the specified
	// load.
	EfficiencyPercent float64
	// LoadPercent shall contain the load, as a percentage, of this power supply at which this efficiency rating is
	// valid.
	LoadPercent float64
}

// UnmarshalJSON unmarshals a EfficiencyRating object from the raw JSON.
func (efficiencyrating *EfficiencyRating) UnmarshalJSON(b []byte) error {
	type temp EfficiencyRating
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*efficiencyrating = EfficiencyRating(t.temp)

	// Extract the links to other entities for later

	return nil
}

// InputRange shall describe an input range that the associated power supply can utilize.
type InputRange struct {
	// CapacityWatts shall contain the maximum amount of power, in watts, that the associated power supply is rated to
	// deliver while operating in this input range.
	CapacityWatts float64
	// NominalVoltageType shall contain the input voltage type of the associated range.
	NominalVoltageType NominalVoltageType
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

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerOutlets shall contain an array of links to resources of type Outlet that represent the outlets that provide
	// power to this power supply.
	PowerOutlets []Outlet
	// PowerOutlets@odata.count
	PowerOutletsCount int `json:"PowerOutlets@odata.count"`
	// PoweringChassis shall contain an array of links to resources of type Chassis that represent the chassis directly
	// powered by this power supply. This property shall not be present if the power supply is only providing power to
	// its containing parent chassis.
	PoweringChassis []Chassis
	// PoweringChassis@odata.count
	PoweringChassisCount int `json:"PoweringChassis@odata.count"`
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

// OutputRail shall describe an output power rail provided by power supply.
type OutputRail struct {
	// NominalVoltage shall contain the nominal voltage of the associated output power rail.
	NominalVoltage float64
	// PhysicalContext shall contain a description of the device or region within the chassis to which this power rail
	// applies.
	PhysicalContext string
}

// UnmarshalJSON unmarshals a OutputRail object from the raw JSON.
func (outputrail *OutputRail) UnmarshalJSON(b []byte) error {
	type temp OutputRail
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*outputrail = OutputRail(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PowerSupply shall represent a power supply for a Redfish implementation.
type PowerSupply struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// Description provides a description of this resource.
	Description string
	// EfficiencyRatings shall contain an array of efficiency ratings for this power supply.
	EfficiencyRatings []EfficiencyRating
	// ElectricalSourceManagerURIs shall contain an array of URIs to the management applications or devices that
	// provide monitoring or control of the upstream electrical sources that provide power to this power supply.
	ElectricalSourceManagerURIs []string
	// ElectricalSourceNames shall contain an arrays of strings that identify the upstream electrical sources, such as
	// the names of circuits or outlets, that provide power to this power supply.
	ElectricalSourceNames []string
	// FirmwareVersion shall contain the firmware version as defined by the manufacturer for this power supply.
	FirmwareVersion string
	// HotPluggable shall indicate whether the device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Devices indicated as hot-pluggable shall allow the device to
	// become operable without altering the operational state of the underlying equipment. Devices that cannot be
	// inserted or removed from equipment in operation, or devices that cannot become operable without affecting the
	// operational state of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// InputNominalVoltageType shall contain the nominal voltage type of the input line voltage in use by this power
	// supply. This value shall be one of the values shown in the NominalVoltageType property in the InputRanges array,
	// if present.
	InputNominalVoltageType NominalVoltageType
	// InputRanges shall contain a collection of ranges usable by this power supply.
	InputRanges []InputRange
	// LineInputStatus shall contain the status of the power line input for this power supply.
	LineInputStatus LineStatus
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of this power supply.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for producing the power supply. This
	// organization may be the entity from whom the power supply is purchased, but this is not necessarily true.
	Manufacturer string
	// Metrics shall contain a link to a resource of type PowerSupplyMetrics.
	Metrics string
	// Model shall contain the model information as defined by the manufacturer for this power supply.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutputNominalVoltageType shall contain the nominal voltage type of the single output line of this power supply.
	// This property is intended to describe power supply types that connect to additional power infrastructure
	// components, such as a rectifier component in a modular power system. This property shall not be present for
	// power supplies not intended to connect to additional power infrastructure components.
	OutputNominalVoltageType NominalVoltageType
	// OutputRails shall contain an array of output power rails provided by this power supply. The elements shall be
	// ordered in ascending nominal voltage order. This ordering is necessary for consistency with Sensor properties in
	// an associated PowerSupplyMetrics resource.
	OutputRails []OutputRail
	// PartNumber shall contain the part number as defined by the manufacturer for this power supply.
	PartNumber string
	// PhaseWiringType shall contain the number of ungrounded current-carrying conductors (phases) and the total number
	// of conductors (wires) included in the input connector for the power supply.
	PhaseWiringType PhaseWiringType
	// PlugType shall contain the type of physical plug used for the input to this power supply, as defined by IEC,
	// NEMA, or regional standard.
	PlugType PlugType
	// PowerCapacityWatts shall contain the maximum amount of power, in watts, that this power supply is rated to
	// deliver.
	PowerCapacityWatts float64
	// PowerSupplyType shall contain the input power type (AC or DC) of this power supply.
	PowerSupplyType PowerSupplyType
	// ProductionDate shall contain the date of production or manufacture for this power supply.
	ProductionDate string
	// Replaceable shall indicate whether this component can be independently replaced as allowed by the vendor's
	// replacement policy. A value of 'false' indicates the component needs to be replaced by policy, as part of
	// another component. If the 'LocationType' property of this component contains 'Embedded', this property shall
	// contain 'false'.
	Replaceable bool
	// SerialNumber shall contain the serial number as defined by the manufacturer for this power supply.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for this power
	// supply.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Version shall contain the hardware version of this power supply as determined by the vendor or supplier.
	Version string
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
		"ElectricalSourceManagerURIs",
		"ElectricalSourceNames",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(powersupply).Elem()

	return powersupply.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPowerSupply will get a PowerSupply instance from the service.
func GetPowerSupply(c common.Client, uri string) (*PowerSupply, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var powersupply PowerSupply
	err = json.NewDecoder(resp.Body).Decode(&powersupply)
	if err != nil {
		return nil, err
	}

	powersupply.SetClient(c)
	return &powersupply, nil
}

// ListReferencedPowerSupplys gets the collection of PowerSupply from
// a provided reference.
func ListReferencedPowerSupplys(c common.Client, link string) ([]*PowerSupply, error) {
	var result []*PowerSupply
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, powersupplyLink := range links.ItemLinks {
		powersupply, err := GetPowerSupply(c, powersupplyLink)
		if err != nil {
			collectionError.Failures[powersupplyLink] = err
		} else {
			result = append(result, powersupply)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
