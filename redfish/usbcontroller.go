//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevice shall contain a link to a resource of type PCIeDevice that represents this USB controller.
	PCIeDevice PCIeDevice
	// Processors shall contain an array of links to resources of type Processor that represent processors that can
	// utilize this USB controller.
	Processors []Processor
	// Processors@odata.count
	ProcessorsCount int `json:"Processors@odata.count"`
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

// USBController shall represent a USB controller in a Redfish implementation.
type USBController struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Manufacturer shall contain the name of the organization responsible for producing the USB controller. This
	// organization may be the entity from which the USB controller is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the manufacturer-provided model information of this USB controller.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the USB controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	Ports string
	// SKU shall contain the SKU number for this USB controller.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the USB controller.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the USB controller.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a USBController object from the raw JSON.
func (usbcontroller *USBController) UnmarshalJSON(b []byte) error {
	type temp USBController
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*usbcontroller = USBController(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetUSBController will get a USBController instance from the service.
func GetUSBController(c common.Client, uri string) (*USBController, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var usbcontroller USBController
	err = json.NewDecoder(resp.Body).Decode(&usbcontroller)
	if err != nil {
		return nil, err
	}

	usbcontroller.SetClient(c)
	return &usbcontroller, nil
}

// ListReferencedUSBControllers gets the collection of USBController from
// a provided reference.
func ListReferencedUSBControllers(c common.Client, link string) ([]*USBController, error) {
	var result []*USBController
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, usbcontrollerLink := range links.ItemLinks {
		usbcontroller, err := GetUSBController(c, usbcontrollerLink)
		if err != nil {
			collectionError.Failures[usbcontrollerLink] = err
		} else {
			result = append(result, usbcontroller)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
