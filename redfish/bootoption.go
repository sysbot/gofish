//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// BootOption shall represent a single boot option within a system.
type BootOption struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Alias shall contain the string alias of this boot source that describes the type of boot.
	Alias BootSource
	// BootOptionEnabled shall indicate whether the boot option is enabled. If 'true', it is enabled. If 'false', the
	// boot option that the boot order array on the computer system contains shall be skipped. In the UEFI context,
	// this property shall influence the load option active flag for the boot option.
	BootOptionEnabled bool
	// BootOptionReference shall correspond to the boot option or device. For UEFI systems, this string shall match the
	// UEFI boot option variable name, such as 'Boot####'. The BootOrder array of a computer system resource contains
	// this value.
	BootOptionReference string
	// Description provides a description of this resource.
	Description string
	// DisplayName shall contain a user-readable boot option name, as it should appear in the boot order list in the
	// user interface.
	DisplayName string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RelatedItem shall contain an array of links to resources or objects that are associated with this boot option.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// UefiDevicePath shall contain the UEFI Specification-defined UEFI device path that identifies and locates the
	// device for this boot option.
	UefiDevicePath string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a BootOption object from the raw JSON.
func (bootoption *BootOption) UnmarshalJSON(b []byte) error {
	type temp BootOption
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*bootoption = BootOption(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	bootoption.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (bootoption *BootOption) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(BootOption)
	original.UnmarshalJSON(bootoption.rawData)

	readWriteFields := []string{
		"BootOptionEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(bootoption).Elem()

	return bootoption.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetBootOption will get a BootOption instance from the service.
func GetBootOption(c common.Client, uri string) (*BootOption, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bootoption BootOption
	err = json.NewDecoder(resp.Body).Decode(&bootoption)
	if err != nil {
		return nil, err
	}

	bootoption.SetClient(c)
	return &bootoption, nil
}

// ListReferencedBootOptions gets the collection of BootOption from
// a provided reference.
func ListReferencedBootOptions(c common.Client, link string) ([]*BootOption, error) {
	var result []*BootOption
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, bootoptionLink := range links.ItemLinks {
		bootoption, err := GetBootOption(c, bootoptionLink)
		if err != nil {
			collectionError.Failures[bootoptionLink] = err
		} else {
			result = append(result, bootoption)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
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
