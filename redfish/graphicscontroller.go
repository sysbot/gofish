//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// GraphicsController shall represent a graphics output device in a Redfish implementation.
type GraphicsController struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AssetTag shall contain the user-assigned asset tag, which is an identifying string that tracks the drive for
	// inventory purposes.
	AssetTag string
	// BiosVersion shall contain the version string of the currently installed and running BIOS or firmware for the
	// graphics controller.
	BiosVersion string
	// Description provides a description of this resource.
	Description string
	// DriverVersion shall contain the version string of the currently loaded driver for this graphics controller.
	DriverVersion string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated graphics controller.
	Location string
	// Manufacturer shall contain the name of the organization responsible for producing the graphics controller. This
	// organization may be the entity from which the graphics controller is purchased, but this is not necessarily
	// true.
	Manufacturer string
	// Model shall contain the manufacturer-provided model information of this graphics controller.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the graphics controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	Ports string
	// SKU shall contain the SKU number for this graphics controller.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the graphics controller.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the graphics controller.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a GraphicsController object from the raw JSON.
func (graphicscontroller *GraphicsController) UnmarshalJSON(b []byte) error {
	type temp GraphicsController
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*graphicscontroller = GraphicsController(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	graphicscontroller.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (graphicscontroller *GraphicsController) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(GraphicsController)
	original.UnmarshalJSON(graphicscontroller.rawData)

	readWriteFields := []string{
		"AssetTag",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(graphicscontroller).Elem()

	return graphicscontroller.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetGraphicsController will get a GraphicsController instance from the service.
func GetGraphicsController(c common.Client, uri string) (*GraphicsController, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var graphicscontroller GraphicsController
	err = json.NewDecoder(resp.Body).Decode(&graphicscontroller)
	if err != nil {
		return nil, err
	}

	graphicscontroller.SetClient(c)
	return &graphicscontroller, nil
}

// ListReferencedGraphicsControllers gets the collection of GraphicsController from
// a provided reference.
func ListReferencedGraphicsControllers(c common.Client, link string) ([]*GraphicsController, error) {
	var result []*GraphicsController
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, graphicscontrollerLink := range links.ItemLinks {
		graphicscontroller, err := GetGraphicsController(c, graphicscontrollerLink)
		if err != nil {
			collectionError.Failures[graphicscontrollerLink] = err
		} else {
			result = append(result, graphicscontroller)
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
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevice shall contain a link to a resource of type PCIeDevice that represents this graphics controller.
	PCIeDevice PCIeDevice
	// Processors shall contain an array of links to resources of type Processor that represent the processors that
	// this graphics controller contains.
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
