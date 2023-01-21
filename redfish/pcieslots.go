//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// SlotTypes is
type SlotTypes string

const (
	// FullLengthSlotTypes Full-Length PCIe slot.
	FullLengthSlotTypes SlotTypes = "FullLength"
	// HalfLengthSlotTypes Half-Length PCIe slot.
	HalfLengthSlotTypes SlotTypes = "HalfLength"
	// LowProfileSlotTypes Low-Profile or Slim PCIe slot.
	LowProfileSlotTypes SlotTypes = "LowProfile"
	// MiniSlotTypes Mini PCIe slot.
	MiniSlotTypes SlotTypes = "Mini"
	// M2SlotTypes PCIe M.2 slot.
	M2SlotTypes SlotTypes = "M2"
	// OEMSlotTypes An OEM-specific slot.
	OEMSlotTypes SlotTypes = "OEM"
	// OCP3SmallSlotTypes Open Compute Project 3.0 small form factor slot.
	OCP3SmallSlotTypes SlotTypes = "OCP3Small"
	// OCP3LargeSlotTypes Open Compute Project 3.0 large form factor slot.
	OCP3LargeSlotTypes SlotTypes = "OCP3Large"
	// U2SlotTypes U.2 / SFF-8639 slot or bay.
	U2SlotTypes SlotTypes = "U2"
)

// OemActions shall contain the available OEM-specific actions for this Resource.
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

// PCIeLinks shall contain links to Resources related to but not subordinate to this Resource.
type PCIeLinks struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevice shall contain an array of links to the Resources of the PCIeDevice type with which this physical slot
	// is associated. If the Status.State of this slot is 'Absent', this property shall not appear in the Resource.
	PCIeDevice []PCIeDevice
	// PCIeDevice@odata.count
	PCIeDeviceCount int `json:"PCIeDevice@odata.count"`
	// Processors shall contain an array of links to resources of type Processor that represent processors that are
	// directly connected or directly bridged to this PCIe slot.
	Processors []Processor
	// Processors@odata.count
	ProcessorsCount int `json:"Processors@odata.count"`
}

// UnmarshalJSON unmarshals a PCIeLinks object from the raw JSON.
func (pcielinks *PCIeLinks) UnmarshalJSON(b []byte) error {
	type temp PCIeLinks
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pcielinks = PCIeLinks(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PCIeSlot shall contain the definition for a PCIe Slot for a Redfish implementation.
type PCIeSlot struct {
	// HotPluggable shall contain indicating whether this PCIe slot supports hotplug.
	HotPluggable bool
	// Lanes shall contain the maximum number of PCIe lanes supported by the slot.
	Lanes int
	// Links shall contain links to Resources related to but not subordinate to this Resource.
	Links string
	// Location shall contain part location information, including a ServiceLabel of the associated PCIe Slot.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeType shall contain the maximum PCIe specification that this slot supports.
	PCIeType PCIeTypes
	// SlotType shall contain the slot type as specified by the PCIe specification.
	SlotType SlotTypes
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a PCIeSlot object from the raw JSON.
func (pcieslot *PCIeSlot) UnmarshalJSON(b []byte) error {
	type temp PCIeSlot
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pcieslot = PCIeSlot(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PCIeSlots shall represent a set of PCIe slot information for a Redfish implementation.
type PCIeSlots struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Slots shall contain an entry for each PCIe slot, including empty slots (with no device or card installed).
	Slots []PCIeSlot
}

// UnmarshalJSON unmarshals a PCIeSlots object from the raw JSON.
func (pcieslots *PCIeSlots) UnmarshalJSON(b []byte) error {
	type temp PCIeSlots
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pcieslots = PCIeSlots(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetPCIeSlots will get a PCIeSlots instance from the service.
func GetPCIeSlots(c common.Client, uri string) (*PCIeSlots, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pcieslots PCIeSlots
	err = json.NewDecoder(resp.Body).Decode(&pcieslots)
	if err != nil {
		return nil, err
	}

	pcieslots.SetClient(c)
	return &pcieslots, nil
}

// ListReferencedPCIeSlotss gets the collection of PCIeSlots from
// a provided reference.
func ListReferencedPCIeSlotss(c common.Client, link string) ([]*PCIeSlots, error) {
	var result []*PCIeSlots
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, pcieslotsLink := range links.ItemLinks {
		pcieslots, err := GetPCIeSlots(c, pcieslotsLink)
		if err != nil {
			collectionError.Failures[pcieslotsLink] = err
		} else {
			result = append(result, pcieslots)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
