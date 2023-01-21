//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// PowerEquipmentType is
type PowerEquipmentType string

const (
	// RackPDUPowerEquipmentType A power distribution unit providing outlets for a rack or similar quantity of devices.
	RackPDUPowerEquipmentType PowerEquipmentType = "RackPDU"
	// FloorPDUPowerEquipmentType A power distribution unit providing feeder circuits for further power distribution.
	FloorPDUPowerEquipmentType PowerEquipmentType = "FloorPDU"
	// ManualTransferSwitchPowerEquipmentType A manual power transfer switch.
	ManualTransferSwitchPowerEquipmentType PowerEquipmentType = "ManualTransferSwitch"
	// AutomaticTransferSwitchPowerEquipmentType An automatic power transfer switch.
	AutomaticTransferSwitchPowerEquipmentType PowerEquipmentType = "AutomaticTransferSwitch"
	// SwitchgearPowerEquipmentType Electrical switchgear.
	SwitchgearPowerEquipmentType PowerEquipmentType = "Switchgear"
	// PowerShelfPowerEquipmentType A power shelf.
	PowerShelfPowerEquipmentType PowerEquipmentType = "PowerShelf"
	// BusPowerEquipmentType An electrical bus.
	BusPowerEquipmentType PowerEquipmentType = "Bus"
)

// TransferSensitivityType is
type TransferSensitivityType string

const (
	// HighTransferSensitivityType High sensitivity for initiating a transfer.
	HighTransferSensitivityType TransferSensitivityType = "High"
	// MediumTransferSensitivityType Medium sensitivity for initiating a transfer.
	MediumTransferSensitivityType TransferSensitivityType = "Medium"
	// LowTransferSensitivityType Low sensitivity for initiating a transfer.
	LowTransferSensitivityType TransferSensitivityType = "Low"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Chassis shall contain an array of links to resources of type Chassis that represents the physical container
	// associated with this resource.
	Chassis []Chassis
	// Chassis@odata.count
	ChassisCount int `json:"Chassis@odata.count"`
	// Facility shall contain a link to a resource of type Facility that represents the facility that contains this
	// equipment.
	Facility string
	// ManagedBy shall contain an array of links to resources of type Manager that represent the managers that manage
	// this equipment.
	ManagedBy []Manager
	// ManagedBy@odata.count
	ManagedByCount int `json:"ManagedBy@odata.count"`
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

// PowerDistribution shall be used to represent a power distribution component or unit for a Redfish
// implementation.
type PowerDistribution struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AssetTag shall contain the user-assigned asset tag, which is an identifying string that tracks the equipment for
	// inventory purposes.
	AssetTag string
	// Branches shall contain a link to a resource collection of type CircuitCollection that contains the branch
	// circuits for this equipment.
	Branches string
	// Description provides a description of this resource.
	Description string
	// EquipmentType shall contain the type of equipment this resource represents.
	EquipmentType string
	// Feeders shall contain a link to a resource collection of type CircuitCollection that contains the feeder
	// circuits for this equipment.
	Feeders string
	// FirmwareVersion shall contain a string describing the firmware version of this equipment as provided by the
	// manufacturer.
	FirmwareVersion string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated equipment.
	Location string
	// Mains shall contain a link to a resource collection of type CircuitCollection that contains the power input
	// circuits for this equipment.
	Mains string
	// MainsRedundancy shall contain redundancy information for the mains (input) circuits for this equipment. The
	// values of the RedundancyGroup array shall reference resources of type Circuit.
	MainsRedundancy string
	// Manufacturer shall contain the name of the organization responsible for producing the equipment. This
	// organization may be the entity from which the equipment is purchased, but this is not necessarily true.
	Manufacturer string
	// Metrics shall contain a link to a resource of type PowerDistributionMetrics.
	Metrics string
	// Model shall contain the manufacturer-provided model information of this equipment.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutletGroups shall contain a link to a resource collection of type OutletCollection that contains the outlet
	// groups for this equipment.
	OutletGroups string
	// Outlets shall contain a link to a resource collection of type OutletCollection that contains the outlets for
	// this equipment.
	Outlets string
	// PartNumber shall contain the manufacturer-provided part number for the equipment.
	PartNumber string
	// PowerSupplies shall contain a link to a resource collection of type PowerSupplyCollection.
	PowerSupplies string
	// PowerSupplyRedundancy shall contain redundancy information for the set of power supplies for this equipment. The
	// values of the RedundancyGroup array shall reference resources of type PowerSupply.
	PowerSupplyRedundancy []RedundantGroup
	// ProductionDate shall contain the date of production or manufacture for this equipment.
	ProductionDate string
	// Sensors shall be a link to a resource collection of type SensorCollection that contains the sensors located in
	// the equipment and sub-components.
	Sensors string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the equipment.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Subfeeds shall contain a link to a resource collection of type CircuitCollection that contains the subfeed
	// circuits for this equipment.
	Subfeeds string
	// TransferConfiguration shall contain the configuration information regarding an automatic transfer switch
	// function for this resource.
	TransferConfiguration TransferConfiguration
	// TransferCriteria shall contain the criteria for initiating a transfer within an automatic transfer switch
	// function for this resource.
	TransferCriteria TransferCriteria
	// UUID shall contain the UUID for the equipment.
	UUID string
	// Version shall contain the hardware version of this equipment as determined by the vendor or supplier.
	Version string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerDistribution object from the raw JSON.
func (powerdistribution *PowerDistribution) UnmarshalJSON(b []byte) error {
	type temp PowerDistribution
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powerdistribution = PowerDistribution(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	powerdistribution.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (powerdistribution *PowerDistribution) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(PowerDistribution)
	original.UnmarshalJSON(powerdistribution.rawData)

	readWriteFields := []string{
		"AssetTag",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(powerdistribution).Elem()

	return powerdistribution.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPowerDistribution will get a PowerDistribution instance from the service.
func GetPowerDistribution(c common.Client, uri string) (*PowerDistribution, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var powerdistribution PowerDistribution
	err = json.NewDecoder(resp.Body).Decode(&powerdistribution)
	if err != nil {
		return nil, err
	}

	powerdistribution.SetClient(c)
	return &powerdistribution, nil
}

// ListReferencedPowerDistributions gets the collection of PowerDistribution from
// a provided reference.
func ListReferencedPowerDistributions(c common.Client, link string) ([]*PowerDistribution, error) {
	var result []*PowerDistribution
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, powerdistributionLink := range links.ItemLinks {
		powerdistribution, err := GetPowerDistribution(c, powerdistributionLink)
		if err != nil {
			collectionError.Failures[powerdistributionLink] = err
		} else {
			result = append(result, powerdistribution)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// TransferConfiguration shall contain the configuration information regarding an automatic transfer switch
// function for this resource.
type TransferConfiguration struct {
	// ActiveMainsId shall contain the mains circuit that is switched on and qualified to supply power to the output
	// circuit. The value shall be a string that matches the Id property value of a circuit contained in the collection
	// referenced by the Mains property.
	ActiveMainsId string
	// AutoTransferEnabled shall indicate if the qualified alternate mains circuit is automatically switched on when
	// the preferred mains circuit becomes unqualified and is automatically switched off.
	AutoTransferEnabled bool
	// ClosedTransitionAllowed shall indicate if a make-before-break switching sequence of the mains circuits is
	// permitted when they are both qualified and in synchronization.
	ClosedTransitionAllowed bool
	// ClosedTransitionTimeoutSeconds shall contain the time in seconds to wait for a closed transition to occur.
	ClosedTransitionTimeoutSeconds int
	// PreferredMainsId shall contain the preferred source for mains circuit to this equipment. The value shall be a
	// string that matches the Id property value of a circuit contained in the collection referenced by the Mains
	// property.
	PreferredMainsId string
	// RetransferDelaySeconds shall contain the time in seconds to delay the automatic transfer from the alternate
	// mains circuit back to the preferred mains circuit.
	RetransferDelaySeconds int
	// RetransferEnabled shall indicate if the automatic transfer is permitted from the alternate mains circuit back to
	// the preferred mains circuit after the preferred mains circuit is qualified again and the RetransferDelaySeconds
	// time has expired.
	RetransferEnabled bool
	// TransferDelaySeconds shall contain the time in seconds to delay the automatic transfer from the preferred mains
	// circuit to the alternate mains circuit when the preferred mains circuit is disqualified. A value of zero shall
	// mean it transfers as fast as possible.
	TransferDelaySeconds int
	// TransferInhibit shall indicate if any transfer is inhibited.
	TransferInhibit bool
}

// UnmarshalJSON unmarshals a TransferConfiguration object from the raw JSON.
func (transferconfiguration *TransferConfiguration) UnmarshalJSON(b []byte) error {
	type temp TransferConfiguration
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*transferconfiguration = TransferConfiguration(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TransferCriteria shall contain the criteria for initiating a transfer within an automatic transfer switch
// function for this resource.
type TransferCriteria struct {
	// OverNominalFrequencyHz shall contain the frequency in hertz over the nominal value that satisfies a criterion
	// for transfer.
	OverNominalFrequencyHz float64
	// OverVoltageRMSPercentage shall contain the positive percentage of voltage RMS over the nominal value that
	// satisfies a criterion for transfer.
	OverVoltageRMSPercentage float64
	// TransferSensitivity shall contain the setting that adjusts the analytical sensitivity of the detection of the
	// quality of voltage waveform that satisfies a criterion for transfer.
	TransferSensitivity TransferSensitivityType
	// UnderNominalFrequencyHz shall contain the frequency in hertz under the nominal value that satisfies a criterion
	// for transfer.
	UnderNominalFrequencyHz float64
	// UnderVoltageRMSPercentage shall contain the negative percentage of voltage RMS under the nominal value that
	// satisfies a criterion for transfer.
	UnderVoltageRMSPercentage float64
}

// UnmarshalJSON unmarshals a TransferCriteria object from the raw JSON.
func (transfercriteria *TransferCriteria) UnmarshalJSON(b []byte) error {
	type temp TransferCriteria
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*transfercriteria = TransferCriteria(t.temp)

	// Extract the links to other entities for later

	return nil
}
