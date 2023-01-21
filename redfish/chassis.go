//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ChassisType is
type ChassisType string

const (
	// RackChassisType An equipment rack, typically a 19-inch wide freestanding unit.
	RackChassisType ChassisType = "Rack"
	// BladeChassisType An enclosed or semi-enclosed, typically vertically-oriented, system chassis that must be
	// plugged into a multi-system chassis to function normally.
	BladeChassisType ChassisType = "Blade"
	// EnclosureChassisType A generic term for a chassis that does not fit any other description.
	EnclosureChassisType ChassisType = "Enclosure"
	// StandAloneChassisType A single, free-standing system, commonly called a tower or desktop chassis.
	StandAloneChassisType ChassisType = "StandAlone"
	// RackMountChassisType A single-system chassis designed specifically for mounting in an equipment rack.
	RackMountChassisType ChassisType = "RackMount"
	// CardChassisType A loose device or circuit board intended to be installed in a system or other enclosure.
	CardChassisType ChassisType = "Card"
	// CartridgeChassisType A small self-contained system intended to be plugged into a multi-system chassis.
	CartridgeChassisType ChassisType = "Cartridge"
	// RowChassisType A collection of equipment racks.
	RowChassisType ChassisType = "Row"
	// PodChassisType A collection of equipment racks in a large, likely transportable, container.
	PodChassisType ChassisType = "Pod"
	// ExpansionChassisType A chassis that expands the capabilities or capacity of another chassis.
	ExpansionChassisType ChassisType = "Expansion"
	// SidecarChassisType A chassis that mates mechanically with another chassis to expand its capabilities or
	// capacity.
	SidecarChassisType ChassisType = "Sidecar"
	// ZoneChassisType A logical division or portion of a physical chassis that contains multiple devices or systems
	// that cannot be physically separated.
	ZoneChassisType ChassisType = "Zone"
	// SledChassisType An enclosed or semi-enclosed, system chassis that must be plugged into a multi-system chassis to
	// function normally similar to a blade type chassis.
	SledChassisType ChassisType = "Sled"
	// ShelfChassisType An enclosed or semi-enclosed, typically horizontally-oriented, system chassis that must be
	// plugged into a multi-system chassis to function normally.
	ShelfChassisType ChassisType = "Shelf"
	// DrawerChassisType An enclosed or semi-enclosed, typically horizontally-oriented, system chassis that can be slid
	// into a multi-system chassis.
	DrawerChassisType ChassisType = "Drawer"
	// ModuleChassisType A small, typically removable, chassis or card that contains devices for a particular subsystem
	// or function.
	ModuleChassisType ChassisType = "Module"
	// ComponentChassisType A small chassis, card, or device that contains devices for a particular subsystem or
	// function.
	ComponentChassisType ChassisType = "Component"
	// IPBasedDriveChassisType A chassis in a drive form factor with IP-based network connections.
	IPBasedDriveChassisType ChassisType = "IPBasedDrive"
	// RackGroupChassisType A group of racks that form a single entity or share infrastructure.
	RackGroupChassisType ChassisType = "RackGroup"
	// StorageEnclosureChassisType A chassis that encloses storage.
	StorageEnclosureChassisType ChassisType = "StorageEnclosure"
	// OtherChassisType A chassis that does not fit any of these definitions.
	OtherChassisType ChassisType = "Other"
)

// EnvironmentalClass is
type EnvironmentalClass string

const (
	// A1EnvironmentalClass ASHRAE Environmental Class 'A1'.
	A1EnvironmentalClass EnvironmentalClass = "A1"
	// A2EnvironmentalClass ASHRAE Environmental Class 'A2'.
	A2EnvironmentalClass EnvironmentalClass = "A2"
	// A3EnvironmentalClass ASHRAE Environmental Class 'A3'.
	A3EnvironmentalClass EnvironmentalClass = "A3"
	// A4EnvironmentalClass ASHRAE Environmental Class 'A4'.
	A4EnvironmentalClass EnvironmentalClass = "A4"
)

// IndicatorLED is
type IndicatorLED string

const (
	// UnknownIndicatorLED shall represent the indicator LED is in an unknown state. The service shall reject PATCH or
	// PUT requests containing this value by returning the HTTP 400 (Bad Request) status code.
	UnknownIndicatorLED IndicatorLED = "Unknown"
	// LitIndicatorLED shall represent the indicator LED is in a solid on state. If the service does not support this
	// value, it shall return the HTTP 400 (Bad Request) status code to reject PATCH or PUT requests that contain this
	// value.
	LitIndicatorLED IndicatorLED = "Lit"
	// BlinkingIndicatorLED shall represent the indicator LED is in a blinking state where the LED is being turned on
	// and off in repetition. If the service does not support this value, it shall reject PATCH or PUT requests
	// containing this value by returning the HTTP 400 (Bad Request) status code.
	BlinkingIndicatorLED IndicatorLED = "Blinking"
	// OffIndicatorLED shall represent the indicator LED is in a solid off state. If the service does not support this
	// value, it shall reject PATCH or PUT requests containing this value by returning the HTTP 400 (Bad Request)
	// status code.
	OffIndicatorLED IndicatorLED = "Off"
)

// IntrusionSensor is
type IntrusionSensor string

const (
	// NormalIntrusionSensor No abnormal physical security condition is detected at this time.
	NormalIntrusionSensor IntrusionSensor = "Normal"
	// HardwareIntrusionIntrusionSensor A door, lock, or other mechanism protecting the internal system hardware from
	// being accessed is detected to be in an insecure state.
	HardwareIntrusionIntrusionSensor IntrusionSensor = "HardwareIntrusion"
	// TamperingDetectedIntrusionSensor Physical tampering of the monitored entity is detected.
	TamperingDetectedIntrusionSensor IntrusionSensor = "TamperingDetected"
)

// IntrusionSensorReArm is
type IntrusionSensorReArm string

const (
	// ManualIntrusionSensorReArm A manual re-arm of this sensor restores it to the normal state.
	ManualIntrusionSensorReArm IntrusionSensorReArm = "Manual"
	// AutomaticIntrusionSensorReArm Because no abnormal physical security condition is detected, this sensor is
	// automatically restored to the normal state.
	AutomaticIntrusionSensorReArm IntrusionSensorReArm = "Automatic"
)

// PowerState is
type PowerState string

const (
	// OnPowerState The components within the chassis have power.
	OnPowerState PowerState = "On"
	// OffPowerState The components within the chassis have no power, except some components might continue to have AUX
	// power, such as the management controller.
	OffPowerState PowerState = "Off"
	// PoweringOnPowerState A temporary state between off and on. The components within the chassis can take time to
	// process the power on action.
	PoweringOnPowerState PowerState = "PoweringOn"
	// PoweringOffPowerState A temporary state between on and off. The components within the chassis can take time to
	// process the power off action.
	PoweringOffPowerState PowerState = "PoweringOff"
)

// ThermalDirection is
type ThermalDirection string

const (
	// FrontToBackThermalDirection shall indicate a chassis with the air intake generally from the front of the chassis
	// and the air exhaust out the back of the chassis.
	FrontToBackThermalDirection ThermalDirection = "FrontToBack"
	// BackToFrontThermalDirection shall indicate a chassis with the air intake generally from the back of the chassis
	// and the air exhaust out the front of the chassis.
	BackToFrontThermalDirection ThermalDirection = "BackToFront"
	// TopExhaustThermalDirection shall indicate a chassis with the air exhaust out the top of the chassis.
	TopExhaustThermalDirection ThermalDirection = "TopExhaust"
	// SealedThermalDirection shall indicate a sealed chassis with no air pathway through the chassis.
	SealedThermalDirection ThermalDirection = "Sealed"
)

// Chassis shall represent a chassis or other physical enclosure for a Redfish implementation.
type Chassis struct {
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
	// AssetTag shall contain an identifying string that tracks the chassis for inventory purposes.
	AssetTag string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	Certificates string
	// ChassisType shall indicate the physical form factor for the type of chassis.
	ChassisType string
	// Controls shall contain a link to a resource collection of type ControlCollection.
	Controls string
	// DepthMm shall represent the depth (length) of the chassis, in millimeters, as specified by the manufacturer.
	DepthMm float64
	// Description provides a description of this resource.
	Description string
	// Drives shall contain a link to a resource collection of type DriveCollection.
	Drives string
	// ElectricalSourceManagerURIs shall contain an array of URIs to the management applications or devices that
	// provide monitoring or control of the upstream electrical sources that provide power to this chassis.
	ElectricalSourceManagerURIs []string
	// ElectricalSourceNames shall contain an arrays of strings that identify the upstream electrical sources, such as
	// the names of circuits or outlets, that provide power to this chassis.
	ElectricalSourceNames []string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this chassis.
	EnvironmentMetrics string
	// EnvironmentalClass shall contain the ASHRAE Environmental Class for this chassis, as defined by ASHRAE Thermal
	// Guidelines for Data Processing Environments. These classes define respective environmental limits that include
	// temperature, relative humidity, dew point, and maximum allowable elevation.
	EnvironmentalClass EnvironmentalClass
	// FabricAdapters shall contain a link to a resource collection of type FabricAdapterCollection.
	FabricAdapters string
	// HeightMm shall represent the height of the chassis, in millimeters, as specified by the manufacturer.
	HeightMm float64
	// HotPluggable shall indicate whether the component can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Components indicated as hot-pluggable shall allow the
	// component to become operable without altering the operational state of the underlying equipment. Components that
	// cannot be inserted or removed from equipment in operation, or components that cannot become operable without
	// affecting the operational state of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated chassis.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// LogServices shall contain a link to a resource collection of type LogServiceCollection.
	LogServices string
	// Manufacturer shall contain the name of the organization responsible for producing the chassis. This organization
	// may be the entity from whom the chassis is purchased, but this is not necessarily true.
	Manufacturer string
	// MaxPowerWatts shall contain the upper bound of the total power consumed by the chassis.
	MaxPowerWatts float64
	// Memory shall contain a link to a resource collection of type MemoryCollection.
	Memory string
	// MemoryDomains shall contain a link to a resource collection of type MemoryDomainCollection.
	MemoryDomains string
	// MinPowerWatts shall contain the lower bound of the total power consumed by the chassis.
	MinPowerWatts float64
	// Model shall contain the name by which the manufacturer generally refers to the chassis.
	Model string
	// NetworkAdapters shall contain a link to a resource collection of type NetworkAdapterCollection.
	NetworkAdapters string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevices shall contain a link to a resource collection of type PCIeDeviceCollection.
	PCIeDevices string
	// PCIeSlots shall contain a link to the resource of type PCIeSlots that represents the PCIe slot information for
	// this chassis.
	PCIeSlots string
	// PartNumber shall contain a part number assigned by the organization that is responsible for producing or
	// manufacturing the chassis.
	PartNumber string
	// PhysicalSecurity shall contain the sensor state of the physical security.
	PhysicalSecurity string
	// PowerState shall contain the power state of the chassis.
	PowerState PowerState
	// PowerSubsystem shall contain a link to a resource of type PowerSubsystem that represents the power subsystem
	// information for this chassis.
	PowerSubsystem string
	// PoweredByParent shall indicate whether the chassis receives power from the chassis that contains it. The value
	// 'true' shall indicate that the containing chassis provides power. The value 'false' shall indicate the chassis
	// receives power from its own power subsystem, another chassis instance's power supplies, or outlets.
	PoweredByParent bool
	// Replaceable shall indicate whether this component can be independently replaced as allowed by the vendor's
	// replacement policy. A value of 'false' indicates the component needs to be replaced by policy, as part of
	// another component. If the 'LocationType' property of this component contains 'Embedded', this property shall
	// contain 'false'.
	Replaceable bool
	// SKU shall contain the stock-keeping unit number for this chassis.
	SKU string
	// Sensors shall contain a link to a resource collection of type SensorCollection that contains the sensors located
	// in the chassis and sub-components.
	Sensors string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the chassis.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the chassis.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// ThermalDirection shall indicate the general direction of the thermal management path through the chassis.
	ThermalDirection ThermalDirection
	// ThermalManagedByParent shall indicate whether the chassis relies on the containing chassis to provide thermal
	// management. The value 'true' shall indicate that the chassis relies on the containing chassis to provide thermal
	// management. The value 'false' shall indicate the chassis provides thermal management, and may provide details in
	// a ThermalSubsystem resource, or by populating the Fans property in Links.
	ThermalManagedByParent bool
	// ThermalSubsystem shall contain a link to a resource of type ThermalSubsystem that represents the thermal
	// subsystem information for this chassis.
	ThermalSubsystem string
	// TrustedComponents shall contain a link to a resource collection of type TrustedComponentCollection.
	TrustedComponents string
	// UUID shall contain the universal unique identifier number for this chassis.
	UUID string
	// Version shall contain the hardware version of this chassis as determined by the vendor or supplier.
	Version string
	// WeightKg shall represent the published mass, commonly referred to as weight, of the chassis, in kilograms.
	WeightKg float64
	// WidthMm shall represent the width of the chassis, in millimeters, as specified by the manufacturer.
	WidthMm float64
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Chassis object from the raw JSON.
func (chassis *Chassis) UnmarshalJSON(b []byte) error {
	type temp Chassis
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*chassis = Chassis(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	chassis.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (chassis *Chassis) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Chassis)
	original.UnmarshalJSON(chassis.rawData)

	readWriteFields := []string{
		"AssetTag",
		"ElectricalSourceManagerURIs",
		"ElectricalSourceNames",
		"EnvironmentalClass",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(chassis).Elem()

	return chassis.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetChassis will get a Chassis instance from the service.
func GetChassis(c common.Client, uri string) (*Chassis, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var chassis Chassis
	err = json.NewDecoder(resp.Body).Decode(&chassis)
	if err != nil {
		return nil, err
	}

	chassis.SetClient(c)
	return &chassis, nil
}

// ListReferencedChassiss gets the collection of Chassis from
// a provided reference.
func ListReferencedChassiss(c common.Client, link string) ([]*Chassis, error) {
	var result []*Chassis
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, chassisLink := range links.ItemLinks {
		chassis, err := GetChassis(c, chassisLink)
		if err != nil {
			collectionError.Failures[chassisLink] = err
		} else {
			result = append(result, chassis)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Cables shall contain an array of links to resources of type Cable that represent the cables connected to this
	// chassis.
	Cables []Cable
	// Cables@odata.count
	CablesCount int `json:"Cables@odata.count"`
	// ComputerSystems shall contain an array of links to resources of type ComputerSystem with which this physical
	// container is associated. If a chassis also links to a computer system to which this resource also links, this
	// chassis shall not link to that computer system.
	ComputerSystems []ComputerSystem
	// ComputerSystems@odata.count
	ComputerSystemsCount int `json:"ComputerSystems@odata.count"`
	// ContainedBy shall contain a link to a resource of type Chassis that represents the chassis that contains this
	// chassis.
	ContainedBy string
	// Contains shall contain an array of links to resources of type Chassis that represent the chassis instances that
	// this chassis contains.
	Contains []Chassis
	// Contains@odata.count
	ContainsCount int `json:"Contains@odata.count"`
	// CooledBy@odata.count
	CooledByCount int `json:"CooledBy@odata.count"`
	// Drives shall contain an array of links to resources of type Drive that are in this chassis.
	Drives []Drive
	// Drives@odata.count
	DrivesCount int `json:"Drives@odata.count"`
	// Facility shall contain a link to the resource of type Facility and shall represent the smallest facility that
	// contains this chassis. This property shall not appear in resources that include a ContainedBy property within
	// the Links property.
	Facility string
	// Fans shall contain an array of links to resources of type Fan that represent the fans that provide cooling to
	// this chassis. This property shall not be present if the ThermalManagedByParent property contains 'true' or if
	// the fans are contained in the ThermalSubsystem resource for this chassis.
	Fans []Fan
	// Fans@odata.count
	FansCount int `json:"Fans@odata.count"`
	// ManagedBy shall contain an array of links to resources of type Manager that manage this chassis.
	ManagedBy []Manager
	// ManagedBy@odata.count
	ManagedByCount int `json:"ManagedBy@odata.count"`
	// ManagersInChassis shall contain an array of links to resources of type Manager that are in this chassis.
	ManagersInChassis []Manager
	// ManagersInChassis@odata.count
	ManagersInChassisCount int `json:"ManagersInChassis@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevices@odata.count
	PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
	// PowerDistribution shall contain a link to a resource of type PowerDistribution that represents the power
	// distribution functionality contained within this chassis.
	PowerDistribution PowerDistribution
	// PowerOutlets shall contain an array of links to resources of type Outlet that represent the outlets that provide
	// power to this chassis. This property shall not be present if the PoweredByParent property contains 'true'.
	PowerOutlets []Outlet
	// PowerOutlets@odata.count
	PowerOutletsCount int `json:"PowerOutlets@odata.count"`
	// PowerSupplies shall contain an array of links to resources of type PowerSupply that represent the power supplies
	// that provide power to this chassis. This property shall not be present if the PoweredByParent property contains
	// 'true' or the power supplies are contained in the PowerSubsystem resource for this chassis.
	PowerSupplies []PowerSupply
	// PowerSupplies@odata.count
	PowerSuppliesCount int `json:"PowerSupplies@odata.count"`
	// PoweredBy@odata.count
	PoweredByCount int `json:"PoweredBy@odata.count"`
	// Processors shall contain an array of links to resources of type Processor type that this chassis contains.
	Processors []Processor
	// Processors@odata.count
	ProcessorsCount int `json:"Processors@odata.count"`
	// ResourceBlocks shall contain an array of links of to resources of type ResourceBlock that this chassis contains.
	ResourceBlocks []ResourceBlock
	// ResourceBlocks@odata.count
	ResourceBlocksCount int `json:"ResourceBlocks@odata.count"`
	// Storage shall contain an array of links to resources of type Storage that are connected to or contained in this
	// chassis.
	Storage []Storage
	// Storage@odata.count
	StorageCount int `json:"Storage@odata.count"`
	// Switches shall contain an array of links to resources of type Switch that this chassis contains.
	Switches []Switch
	// Switches@odata.count
	SwitchesCount int `json:"Switches@odata.count"`
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

// PhysicalSecurity shall describe the sensor state of the physical security.
type PhysicalSecurity struct {
	// IntrusionSensor shall represent the state of this physical security sensor. Hardware intrusion indicates the
	// internal hardware is detected as being accessed in an insecure state. Tampering detected indicates the physical
	// tampering of the monitored entity is detected.
	IntrusionSensor IntrusionSensor
	// IntrusionSensorNumber shall contain a numerical identifier for this physical security sensor that is unique
	// within this resource.
	IntrusionSensorNumber int
	// IntrusionSensorReArm shall represent the method that restores this physical security sensor to the normal state.
	// Manual indicates manual re-arm is needed. Automatic indicates the state is restored automatically because no
	// abnormal physical security conditions are detected.
	IntrusionSensorReArm IntrusionSensorReArm
}

// UnmarshalJSON unmarshals a PhysicalSecurity object from the raw JSON.
func (physicalsecurity *PhysicalSecurity) UnmarshalJSON(b []byte) error {
	type temp PhysicalSecurity
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*physicalsecurity = PhysicalSecurity(t.temp)

	// Extract the links to other entities for later

	return nil
}
