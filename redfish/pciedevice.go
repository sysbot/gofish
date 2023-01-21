//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DeviceType is
type DeviceType string

const (
	// SingleFunctionDeviceType A single-function PCIe device.
	SingleFunctionDeviceType DeviceType = "SingleFunction"
	// MultiFunctionDeviceType A multi-function PCIe device.
	MultiFunctionDeviceType DeviceType = "MultiFunction"
	// SimulatedDeviceType A PCIe device that is not currently physically present, but is being simulated by the PCIe
	// infrastructure.
	SimulatedDeviceType DeviceType = "Simulated"
	// RetimerDeviceType A PCIe retimer device.
	RetimerDeviceType DeviceType = "Retimer"
)

// LaneSplittingType is
type LaneSplittingType string

const (
	// NoneLaneSplittingType The slot has no lane splitting.
	NoneLaneSplittingType LaneSplittingType = "None"
	// BridgedLaneSplittingType The slot has a bridge to share the lanes with associated devices.
	BridgedLaneSplittingType LaneSplittingType = "Bridged"
	// BifurcatedLaneSplittingType The slot is bifurcated to split the lanes with associated devices.
	BifurcatedLaneSplittingType LaneSplittingType = "Bifurcated"
)

// SlotType is
type SlotType string

const (
	// FullLengthSlotType Full-Length PCIe slot.
	FullLengthSlotType SlotType = "FullLength"
	// HalfLengthSlotType Half-Length PCIe slot.
	HalfLengthSlotType SlotType = "HalfLength"
	// LowProfileSlotType Low-Profile or Slim PCIe slot.
	LowProfileSlotType SlotType = "LowProfile"
	// MiniSlotType Mini PCIe slot.
	MiniSlotType SlotType = "Mini"
	// M2SlotType PCIe M.2 slot.
	M2SlotType SlotType = "M2"
	// OEMSlotType An OEM-specific slot.
	OEMSlotType SlotType = "OEM"
	// OCP3SmallSlotType Open Compute Project 3.0 small form factor slot.
	OCP3SmallSlotType SlotType = "OCP3Small"
	// OCP3LargeSlotType Open Compute Project 3.0 large form factor slot.
	OCP3LargeSlotType SlotType = "OCP3Large"
	// U2SlotType U.2 / SFF-8639 slot or bay.
	U2SlotType SlotType = "U2"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Chassis shall contain an array of links to resources of type Chassis that represent the physical containers
	// associated with this resource.
	Chassis []Chassis
	// Chassis@odata.count
	ChassisCount int `json:"Chassis@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeFunctions@odata.count
	PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	// Switch shall contain a link to a resource of type Switch that is associated with this PCIe device.
	Switch Switch
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

// PCIeDevice shall represent a PCIe device in a Redfish implementation.
type PCIeDevice struct {
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
	// AssetTag shall contain an identifying string that tracks the PCIe device for inventory purposes.
	AssetTag string
	// Description provides a description of this resource.
	Description string
	// DeviceType shall contain the device type of the PCIe device such as 'SingleFunction' or 'MultiFunction'.
	DeviceType string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this PCIe device.
	EnvironmentMetrics string
	// FirmwareVersion shall contain the firmware version of the PCIe device.
	FirmwareVersion string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Manufacturer shall contain the name of the organization responsible for producing the PCIe device. This
	// organization may be the entity from whom the PCIe device is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to the PCIe device.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeFunctions shall contain a link to a Resource Collection of type PCIeFunctionCollection. This property should
	// not be present if DeviceType contains 'Retimer'.
	PCIeFunctions string
	// PCIeInterface shall contain details for the PCIe interface that connects this PCIe device to its host or
	// upstream switch.
	PCIeInterface string
	// PartNumber shall contain a part number assigned by the organization that is responsible for producing or
	// manufacturing the PCIe device.
	PartNumber string
	// ReadyToRemove shall indicate whether the PCIe device is ready for removal. Setting the value to 'true' shall
	// cause the service to perform appropriate actions to quiesce the device. A task may spawn while the device is
	// quiescing.
	ReadyToRemove bool
	// SKU shall contain the stock-keeping unit number for this PCIe device.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the PCIe device.
	SerialNumber string
	// Slot shall contain information about the PCIe slot for this PCIe device.
	Slot Slot
	// SparePartNumber shall contain the spare part number of the PCIe device.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UUID shall contain the universal unique identifier number for this PCIe device.
	UUID string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PCIeDevice object from the raw JSON.
func (pciedevice *PCIeDevice) UnmarshalJSON(b []byte) error {
	type temp PCIeDevice
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pciedevice = PCIeDevice(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	pciedevice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (pciedevice *PCIeDevice) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(PCIeDevice)
	original.UnmarshalJSON(pciedevice.rawData)

	readWriteFields := []string{
		"AssetTag",
		"ReadyToRemove",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(pciedevice).Elem()

	return pciedevice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPCIeDevice will get a PCIeDevice instance from the service.
func GetPCIeDevice(c common.Client, uri string) (*PCIeDevice, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pciedevice PCIeDevice
	err = json.NewDecoder(resp.Body).Decode(&pciedevice)
	if err != nil {
		return nil, err
	}

	pciedevice.SetClient(c)
	return &pciedevice, nil
}

// ListReferencedPCIeDevices gets the collection of PCIeDevice from
// a provided reference.
func ListReferencedPCIeDevices(c common.Client, link string) ([]*PCIeDevice, error) {
	var result []*PCIeDevice
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, pciedeviceLink := range links.ItemLinks {
		pciedevice, err := GetPCIeDevice(c, pciedeviceLink)
		if err != nil {
			collectionError.Failures[pciedeviceLink] = err
		} else {
			result = append(result, pciedevice)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// PCIeErrors shall contain properties that describe the PCIe errors associated with this device.
type PCIeErrors struct {
	// CorrectableErrorCount shall contain the total number of the PCIe correctable errors for this device.
	CorrectableErrorCount int
	// FatalErrorCount shall contain the total number of the PCIe fatal errors for this device.
	FatalErrorCount int
	// L0ToRecoveryCount shall contain the total number of times the PCIe link transitioned from L0 to the recovery
	// state for this device.
	L0ToRecoveryCount int
	// NAKReceivedCount shall contain the total number of NAKs issued on the PCIe link by the receiver. A NAK is issued
	// by the receiver when it detects that a TLP from this device was missed. This could be because this device did
	// not transmit it, or because the receiver could not properly decode the packet.
	NAKReceivedCount int
	// NAKSentCount shall contain the total number of NAKs issued on the PCIe link by this device. A NAK is issued by
	// the device when it detects that a TLP from the receiver was missed. This could be because the receiver did not
	// transmit it, or because this device could not properly decode the packet.
	NAKSentCount int
	// NonFatalErrorCount shall contain the total number of the PCIe non-fatal errors for this device.
	NonFatalErrorCount int
	// ReplayCount shall contain the total number of the replays issued on the PCIe link by this device. A replay is a
	// retransmission of a TLP and occurs because the ACK timer is expired, which means that the receiver did not send
	// the ACK or this device did not properly decode the ACK.
	ReplayCount int
	// ReplayRolloverCount shall contain the total number of the replay rollovers issued on the PCIe link by this
	// device. A replay rollover occurs when consecutive replays failed to resolve the errors on the link, which means
	// that this device forced the link into the recovery state.
	ReplayRolloverCount int
}

// UnmarshalJSON unmarshals a PCIeErrors object from the raw JSON.
func (pcieerrors *PCIeErrors) UnmarshalJSON(b []byte) error {
	type temp PCIeErrors
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pcieerrors = PCIeErrors(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PCIeInterface shall contain the definition for a PCIe interface for a Redfish implementation.
type PCIeInterface struct {
	// LanesInUse shall contain the number of PCIe lanes in use by this device, which shall be equal to or less than
	// the MaxLanes property value.
	LanesInUse int
	// MaxLanes shall contain the maximum number of PCIe lanes supported by this device.
	MaxLanes int
	// MaxPCIeType shall contain the maximum PCIe specification that this device supports.
	MaxPCIeType PCIeTypes
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeType shall contain the negotiated PCIe interface version in use by this device.
	PCIeType PCIeTypes
}

// UnmarshalJSON unmarshals a PCIeInterface object from the raw JSON.
func (pcieinterface *PCIeInterface) UnmarshalJSON(b []byte) error {
	type temp PCIeInterface
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pcieinterface = PCIeInterface(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Slot shall contain properties that describe the PCIe slot associated with a PCIe device.
type Slot struct {
	// LaneSplitting shall contain lane splitting information of the associated PCIe slot.
	LaneSplitting LaneSplittingType
	// Lanes shall contain the maximum number of PCIe lanes supported by the slot.
	Lanes int
	// Location shall contain part location information, including a ServiceLabel property, of the associated PCIe
	// slot.
	Location string
	// PCIeType shall contain the maximum PCIe specification that this slot supports.
	PCIeType PCIeTypes
	// SlotType shall contain the PCIe slot type.
	SlotType SlotType
}

// UnmarshalJSON unmarshals a Slot object from the raw JSON.
func (slot *Slot) UnmarshalJSON(b []byte) error {
	type temp Slot
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*slot = Slot(t.temp)

	// Extract the links to other entities for later

	return nil
}
