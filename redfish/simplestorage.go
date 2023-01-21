//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Device shall describe a storage device visible to simple storage.
type Device struct {
	common.Entity
	// CapacityBytes shall represent the size, in bytes, of the storage device.
	CapacityBytes int
	// Manufacturer shall indicate the name of the manufacturer of this storage device.
	Manufacturer string
	// Model shall indicate the model information as provided by the manufacturer of this storage device.
	Model string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a Device object from the raw JSON.
func (device *Device) UnmarshalJSON(b []byte) error {
	type temp Device
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*device = Device(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to Resources related to but not subordinate to this Resource.
type Links struct {
	// Chassis shall contain a link to a Resource of type Chassis that represents the physical container associated
	// with this Resource.
	Chassis string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Storage shall contain a link to a Resource of type Storage that represents the same storage subsystem as this
	// Resource.
	Storage string
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

// SimpleStorage This Resource contains a storage controller and its directly-attached devices.
type SimpleStorage struct {
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
	// Devices shall contain a list of storage devices related to this Resource.
	Devices []Device
	// Links shall contain links to Resources related to but not subordinate to this Resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
	// UefiDevicePath shall contain the UEFI device path that identifies and locates the specific storage controller.
	UefiDevicePath string
}

// UnmarshalJSON unmarshals a SimpleStorage object from the raw JSON.
func (simplestorage *SimpleStorage) UnmarshalJSON(b []byte) error {
	type temp SimpleStorage
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*simplestorage = SimpleStorage(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetSimpleStorage will get a SimpleStorage instance from the service.
func GetSimpleStorage(c common.Client, uri string) (*SimpleStorage, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var simplestorage SimpleStorage
	err = json.NewDecoder(resp.Body).Decode(&simplestorage)
	if err != nil {
		return nil, err
	}

	simplestorage.SetClient(c)
	return &simplestorage, nil
}

// ListReferencedSimpleStorages gets the collection of SimpleStorage from
// a provided reference.
func ListReferencedSimpleStorages(c common.Client, link string) ([]*SimpleStorage, error) {
	var result []*SimpleStorage
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, simplestorageLink := range links.ItemLinks {
		simplestorage, err := GetSimpleStorage(c, simplestorageLink)
		if err != nil {
			collectionError.Failures[simplestorageLink] = err
		} else {
			result = append(result, simplestorage)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
