//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AdditionalVersions shall contain additional versions.
type AdditionalVersions struct {
	// Bootloader shall contain the bootloader version contained in this software.
	Bootloader string
	// Kernel shall contain the kernel version contained in this software. For strict POSIX software, the value shall
	// contain the output of 'uname -srm'. For Microsoft Windows, the value shall contain the output of 'ver'.
	Kernel string
	// Microcode shall contain the microcode version contained in this software.
	Microcode string
	// OSDistribution shall contain the operating system name of this software.
	OSDistribution string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a AdditionalVersions object from the raw JSON.
func (additionalversions *AdditionalVersions) UnmarshalJSON(b []byte) error {
	type temp AdditionalVersions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*additionalversions = AdditionalVersions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// MeasurementBlock shall describe a DSP0274-defined measurement block.
type MeasurementBlock struct {
	// Measurement shall contain the value of the hexadecimal string representation of the numeric value of the
	// DSP0274-defined Measurement field of the measurement block.
	Measurement string
	// MeasurementIndex shall contain the value of DSP0274-defined Index field of the measurement block.
	MeasurementIndex int
	// MeasurementSize shall contain the value of DSP0274-defined MeasurementSize field of the measurement block.
	MeasurementSize int
	// MeasurementSpecification shall contain the value of DSP0274-defined MeasurementSpecification field of the
	// measurement block.
	MeasurementSpecification int
}

// UnmarshalJSON unmarshals a MeasurementBlock object from the raw JSON.
func (measurementblock *MeasurementBlock) UnmarshalJSON(b []byte) error {
	type temp MeasurementBlock
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*measurementblock = MeasurementBlock(t.temp)

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

// SoftwareInventory This Resource contains a single software component that this Redfish Service manages.
type SoftwareInventory struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// AdditionalVersions shall contain the additional versions of this software.
	AdditionalVersions string
	// Description provides a description of this resource.
	Description string
	// LowestSupportedVersion shall represent the lowest supported version of this software. This string is formatted
	// using the same format used for the Version property.
	LowestSupportedVersion string
	// Manufacturer shall represent the name of the manufacturer or producer of this software.
	Manufacturer string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RelatedItem shall contain an array of IDs for pointers consistent with JSON Pointer syntax to the Resource that
	// is associated with this software inventory item.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// ReleaseDate shall contain the date of release or production for this software. If the time of day is unknown,
	// the time of day portion of the property shall contain '00:00:00Z'.
	ReleaseDate string
	// SoftwareId shall represent an implementation-specific label that identifies this software. This string
	// correlates with a component repository or database.
	SoftwareId string
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
	// UefiDevicePaths shall contain a list UEFI device paths of the components associated with this software inventory
	// item. The UEFI device paths shall be formatted as defined by the UEFI Specification.
	UefiDevicePaths []string
	// Updateable shall indicate whether the Update Service can update this software. If 'true', the Service can update
	// this software. If 'false', the Service cannot update this software and the software is for reporting purposes
	// only.
	Updateable bool
	// Version shall contain the version of this software.
	Version string
	// WriteProtected shall indicate whether the software image can be overwritten, where a value 'true' shall indicate
	// that the software cannot be altered or overwritten.
	WriteProtected bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SoftwareInventory object from the raw JSON.
func (softwareinventory *SoftwareInventory) UnmarshalJSON(b []byte) error {
	type temp SoftwareInventory
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*softwareinventory = SoftwareInventory(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	softwareinventory.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (softwareinventory *SoftwareInventory) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(SoftwareInventory)
	original.UnmarshalJSON(softwareinventory.rawData)

	readWriteFields := []string{
		"WriteProtected",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(softwareinventory).Elem()

	return softwareinventory.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSoftwareInventory will get a SoftwareInventory instance from the service.
func GetSoftwareInventory(c common.Client, uri string) (*SoftwareInventory, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var softwareinventory SoftwareInventory
	err = json.NewDecoder(resp.Body).Decode(&softwareinventory)
	if err != nil {
		return nil, err
	}

	softwareinventory.SetClient(c)
	return &softwareinventory, nil
}

// ListReferencedSoftwareInventorys gets the collection of SoftwareInventory from
// a provided reference.
func ListReferencedSoftwareInventorys(c common.Client, link string) ([]*SoftwareInventory, error) {
	var result []*SoftwareInventory
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, softwareinventoryLink := range links.ItemLinks {
		softwareinventory, err := GetSoftwareInventory(c, softwareinventoryLink)
		if err != nil {
			collectionError.Failures[softwareinventoryLink] = err
		} else {
			result = append(result, softwareinventory)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
