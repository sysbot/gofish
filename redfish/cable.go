//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CableClass is
type CableClass string

const (
	// PowerCableClass This cable is used for connecting to a power system.
	PowerCableClass CableClass = "Power"
	// NetworkCableClass This cable is used for connecting to a networking system.
	NetworkCableClass CableClass = "Network"
	// StorageCableClass This cable is used for connecting to a storage system.
	StorageCableClass CableClass = "Storage"
	// FanCableClass This cable is used for connecting to a fan system.
	FanCableClass CableClass = "Fan"
	// PCIeCableClass This cable is used for connecting to a PCIe endpoint.
	PCIeCableClass CableClass = "PCIe"
	// USBCableClass This cable is used for connecting to a USB endpoint.
	USBCableClass CableClass = "USB"
	// VideoCableClass This cable is used for connecting to a video system.
	VideoCableClass CableClass = "Video"
	// FabricCableClass This cable is used for connecting to a fabric.
	FabricCableClass CableClass = "Fabric"
	// SerialCableClass This cable is used for connecting to a serial endpoint.
	SerialCableClass CableClass = "Serial"
	// GeneralCableClass This cable is used for providing general connectivity.
	GeneralCableClass CableClass = "General"
)

// CableStatus is
type CableStatus string

const (
	// NormalCableStatus shall indicate the cable is operating normally. The State property in Status shall contain the
	// value 'Enabled' and The Health property in Status shall contain the value 'OK'.
	NormalCableStatus CableStatus = "Normal"
	// DegradedCableStatus shall indicate the cable is degraded. The State property in Status shall contain the value
	// 'Enabled' and The Health property in Status shall contain the value 'Warning'.
	DegradedCableStatus CableStatus = "Degraded"
	// FailedCableStatus shall indicate the cable has failed. The State property in Status shall contain the value
	// 'Enabled' and The Health property in Status shall contain the value 'Critical'.
	FailedCableStatus CableStatus = "Failed"
	// TestingCableStatus shall indicate the cable is under test. The State property in Status shall contain the value
	// 'InTest'.
	TestingCableStatus CableStatus = "Testing"
	// DisabledCableStatus shall indicate the cable is disabled. The State property in Status shall contain the value
	// 'Disabled'.
	DisabledCableStatus CableStatus = "Disabled"
	// SetByServiceCableStatus shall indicate the status for the cable is not defined by the user. If implemented, the
	// service shall determine the value of the State and Health properties in Status.
	SetByServiceCableStatus CableStatus = "SetByService"
)

// ConnectorType is
type ConnectorType string

const (
	// ACPowerConnectorType This cable connects to a AC power connector.
	ACPowerConnectorType ConnectorType = "ACPower"
	// DB9ConnectorType This cable connects to a DB9 connector.
	DB9ConnectorType ConnectorType = "DB9"
	// DCPowerConnectorType This cable connects to a DC power connector.
	DCPowerConnectorType ConnectorType = "DCPower"
	// DisplayPortConnectorType This cable connects to a DisplayPort power connector.
	DisplayPortConnectorType ConnectorType = "DisplayPort"
	// HDMIConnectorType This cable connects to an HDMI connector.
	HDMIConnectorType ConnectorType = "HDMI"
	// ICIConnectorType This cable connects to an ICI connector.
	ICIConnectorType ConnectorType = "ICI"
	// IPASSConnectorType This cable connects to an IPASS connector.
	IPASSConnectorType ConnectorType = "IPASS"
	// PCIeConnectorType This cable connects to a PCIe connector.
	PCIeConnectorType ConnectorType = "PCIe"
	// ProprietaryConnectorType This cable connects to a proprietary connector.
	ProprietaryConnectorType ConnectorType = "Proprietary"
	// RJ45ConnectorType This cable connects to an RJ45 connector.
	RJ45ConnectorType ConnectorType = "RJ45"
	// SATAConnectorType This cable connects to a SATA connector.
	SATAConnectorType ConnectorType = "SATA"
	// SCSIConnectorType This cable connects to a SCSI connector.
	SCSIConnectorType ConnectorType = "SCSI"
	// SlimSASConnectorType This cable connects to a SlimSAS connector.
	SlimSASConnectorType ConnectorType = "SlimSAS"
	// SFPConnectorType This cable connects to a SFP connector.
	SFPConnectorType ConnectorType = "SFP"
	// SFPPlusConnectorType This cable connects to a SFPPlus connector.
	SFPPlusConnectorType ConnectorType = "SFPPlus"
	// USBAConnectorType This cable connects to a USB-A connector.
	USBAConnectorType ConnectorType = "USBA"
	// USBCConnectorType This cable connects to a USB-C connector.
	USBCConnectorType ConnectorType = "USBC"
	// QSFPConnectorType This cable connects to a QSFP connector.
	QSFPConnectorType ConnectorType = "QSFP"
	// CDFPConnectorType This cable connects to a CDFP connector.
	CDFPConnectorType ConnectorType = "CDFP"
	// OSFPConnectorType This cable connects to a OSFP connector.
	OSFPConnectorType ConnectorType = "OSFP"
)

// Cable This resource contains a simple cable for a Redfish implementation.
type Cable struct {
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
	// AssetTag shall track the cable for inventory purposes.
	AssetTag string
	// CableClass shall contain the cable class for this cable.
	CableClass CableClass
	// CableStatus shall contain the user-reported status of this resource.
	CableStatus string
	// CableType shall contain a user-defined type for this cable.
	CableType string
	// Description provides a description of this resource.
	Description string
	// DownstreamConnectorTypes shall contain an array of connector types this cable supports.
	DownstreamConnectorTypes []ConnectorType
	// DownstreamName shall contain any identifier for a downstream resource.
	DownstreamName string
	// LengthMeters shall contain the length of the cable in meters.
	LengthMeters float64
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated assembly.
	Location string
	// Manufacturer shall contain the name of the organization responsible for producing the cable. This organization
	// might be the entity from whom the cable is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to the cable.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number assigned by the organization that is responsible for producing or
	// manufacturing the cable.
	PartNumber string
	// SKU shall contain the stock-keeping unit (SKU) number for this cable.
	SKU string
	// SerialNumber shall contain the manufacturer-allocated number that identifies the cable.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UpstreamConnectorTypes shall contain an array of connector types this cable supports.
	UpstreamConnectorTypes []ConnectorType
	// UpstreamName shall contain any identifier for an upstream resource.
	UpstreamName string
	// UserDescription shall contain a user-defined description for this cable.
	UserDescription string
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
	// Vendor shall contain the name of the company that provides the final product that includes this cable.
	Vendor string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Cable object from the raw JSON.
func (cable *Cable) UnmarshalJSON(b []byte) error {
	type temp Cable
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*cable = Cable(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	cable.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (cable *Cable) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Cable)
	original.UnmarshalJSON(cable.rawData)

	readWriteFields := []string{
		"AssetTag",
		"CableClass",
		"CableStatus",
		"CableType",
		"DownstreamConnectorTypes",
		"DownstreamName",
		"LengthMeters",
		"Manufacturer",
		"Model",
		"PartNumber",
		"SKU",
		"SerialNumber",
		"UpstreamConnectorTypes",
		"UpstreamName",
		"UserDescription",
		"UserLabel",
		"Vendor",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(cable).Elem()

	return cable.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetCable will get a Cable instance from the service.
func GetCable(c common.Client, uri string) (*Cable, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cable Cable
	err = json.NewDecoder(resp.Body).Decode(&cable)
	if err != nil {
		return nil, err
	}

	cable.SetClient(c)
	return &cable, nil
}

// ListReferencedCables gets the collection of Cable from
// a provided reference.
func ListReferencedCables(c common.Client, link string) ([]*Cable, error) {
	var result []*Cable
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, cableLink := range links.ItemLinks {
		cable, err := GetCable(c, cableLink)
		if err != nil {
			collectionError.Failures[cableLink] = err
		} else {
			result = append(result, cable)
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
	// DownstreamChassis shall contain an array of links to resources of type Chassis that represent the physical
	// downstream containers connected to this cable.
	DownstreamChassis []Chassis
	// DownstreamChassis@odata.count
	DownstreamChassisCount int `json:"DownstreamChassis@odata.count"`
	// DownstreamPorts shall contain an array of links to resources of type Port that represent the physical downstream
	// connections connected to this cable.
	DownstreamPorts []Port
	// DownstreamPorts@odata.count
	DownstreamPortsCount int `json:"DownstreamPorts@odata.count"`
	// DownstreamResources shall contain an array of links to resources that represent the physical downstream
	// connections connected to this cable. Even if the resource is already referenced in another property within
	// Links, such as DownstreamPorts or DownstreamChassis, it shall also be referenced in this property.
	DownstreamResources []Resource
	// DownstreamResources@odata.count
	DownstreamResourcesCount int `json:"DownstreamResources@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// UpstreamChassis shall contain an array of links to resources of type Chassis that represent the physical
	// upstream containers connected to this cable.
	UpstreamChassis []Chassis
	// UpstreamChassis@odata.count
	UpstreamChassisCount int `json:"UpstreamChassis@odata.count"`
	// UpstreamPorts shall contain an array of links to resources of type Port that represent the physical upstream
	// connections connected to this cable.
	UpstreamPorts []Port
	// UpstreamPorts@odata.count
	UpstreamPortsCount int `json:"UpstreamPorts@odata.count"`
	// UpstreamResources shall contain an array of links to resources that represent the physical upstream connections
	// connected to this cable. Even if the resource is already referenced in another property within Links, such as
	// UpstreamPorts or UpstreamChassis, it shall also be referenced in this property.
	UpstreamResources []Resource
	// UpstreamResources@odata.count
	UpstreamResourcesCount int `json:"UpstreamResources@odata.count"`
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
