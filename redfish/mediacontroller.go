//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// MediaControllerType is
type MediaControllerType string

const (
	// MemoryMediaControllerType shall indicate the media controller is for memory.
	MemoryMediaControllerType MediaControllerType = "Memory"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Endpoints shall contain an array of links to resources of type Endpoint with which this media controller is
	// associated.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// MemoryDomains shall contain an array of links to resources of type MemoryDomain that represent the memory
	// domains associated with this memory controller.
	MemoryDomains []MemoryDomain
	// MemoryDomains@odata.count
	MemoryDomainsCount int `json:"MemoryDomains@odata.count"`
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

// MediaController This resource contains the media controller in a Redfish implementation.
type MediaController struct {
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
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this media controller.
	EnvironmentMetrics string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Manufacturer shall contain the manufacturer of the media controller.
	Manufacturer string
	// MediaControllerType shall contain the type of media controller.
	MediaControllerType MediaControllerType
	// Model shall contain the model of the media controller.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall indicate the part number as provided by the manufacturer of this media controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	Ports string
	// SerialNumber shall indicate the serial number as provided by the manufacturer of this media controller.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UUID shall contain a universal unique identifier number for the media controller.
	UUID string
}

// UnmarshalJSON unmarshals a MediaController object from the raw JSON.
func (mediacontroller *MediaController) UnmarshalJSON(b []byte) error {
	type temp MediaController
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*mediacontroller = MediaController(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetMediaController will get a MediaController instance from the service.
func GetMediaController(c common.Client, uri string) (*MediaController, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mediacontroller MediaController
	err = json.NewDecoder(resp.Body).Decode(&mediacontroller)
	if err != nil {
		return nil, err
	}

	mediacontroller.SetClient(c)
	return &mediacontroller, nil
}

// ListReferencedMediaControllers gets the collection of MediaController from
// a provided reference.
func ListReferencedMediaControllers(c common.Client, link string) ([]*MediaController, error) {
	var result []*MediaController
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, mediacontrollerLink := range links.ItemLinks {
		mediacontroller, err := GetMediaController(c, mediacontrollerLink)
		if err != nil {
			collectionError.Failures[mediacontrollerLink] = err
		} else {
			result = append(result, mediacontroller)
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
