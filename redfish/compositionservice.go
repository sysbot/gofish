//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ComposeRequestFormat is
type ComposeRequestFormat string

const (
	// ManifestComposeRequestFormat shall indicate that the request contains a manifest as defined by the Redfish
	// Manifest schema.
	ManifestComposeRequestFormat ComposeRequestFormat = "Manifest"
)

// ComposeRequestType is
type ComposeRequestType string

const (
	// PreviewComposeRequestType shall indicate that the request is to preview the outcome of the operations specified
	// by the manifest to show what the service will do based on the contents of the request, and not affect any
	// resources within the service.
	PreviewComposeRequestType ComposeRequestType = "Preview"
	// PreviewReserveComposeRequestType shall indicate that the request is to preview the outcome of the operations
	// specified by the manifest to show what the service will do based on the contents of the request. Resources that
	// would have been affected by this request shall be marked as reserved, but otherwise shall not be affected.
	PreviewReserveComposeRequestType ComposeRequestType = "PreviewReserve"
	// ApplyComposeRequestType shall indicate that the request is to apply the requested operations specified by the
	// manifest and modify resources as needed.
	ApplyComposeRequestType ComposeRequestType = "Apply"
)

// ComposeResponse shall contain the properties found in the response body for the Compose action.
type ComposeResponse struct {
	// Manifest shall contain the manifest containing the compose operation response. This property shall be required
	// if RequestFormat contains the value 'Manifest'.
	Manifest string
	// RequestFormat shall contain the format of the request.
	RequestFormat string
	// RequestType shall contain the type of request.
	RequestType string
	// ReservationId shall contain the value of the Id property of the CompositionReservation resource that was
	// created. This property shall be required if RequestType contains the value 'PreviewReserve'.
	ReservationId string
}

// UnmarshalJSON unmarshals a ComposeResponse object from the raw JSON.
func (composeresponse *ComposeResponse) UnmarshalJSON(b []byte) error {
	type temp ComposeResponse
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*composeresponse = ComposeResponse(t.temp)

	// Extract the links to other entities for later

	return nil
}

// CompositionService shall represent the composition service and its properties for a Redfish implementation.
type CompositionService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ActivePool shall contain a link to a resource collection of type ResourceBlockCollection. The members of this
	// collection shall represent the resource blocks in the active pool. Services shall filter members of this
	// collection based on the requesting client.
	ActivePool string
	// AllowOverprovisioning shall indicate whether this service is allowed to overprovision a composition relative to
	// the composition request.
	AllowOverprovisioning bool
	// AllowZoneAffinity shall indicate whether a client can request that a specific resource zone fulfill a
	// composition request.
	AllowZoneAffinity bool
	// CompositionReservations shall contain a link to a resource collection of type CompositionReservationCollection.
	// The members of this collection shall contain links to reserved resource blocks and the related document that
	// caused the reservations. Services shall filter members of this collection based on the requesting client.
	CompositionReservations string
	// Description provides a description of this resource.
	Description string
	// FreePool shall contain a link to a resource collection of type ResourceBlockCollection. The members of this
	// collection shall represent the resource blocks in the free pool. Services shall filter members of this
	// collection based on the requesting client.
	FreePool string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ReservationDuration shall contain the length of time a composition reservation is held before the service
	// deletes the reservation marks any related resource blocks as no longer reserved.
	ReservationDuration string
	// ResourceBlocks shall contain a link to a resource collection of type ResourceBlockCollection.
	ResourceBlocks string
	// ResourceZones shall contain a link to a resource collection of type ZoneCollection.
	ResourceZones string
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a CompositionService object from the raw JSON.
func (compositionservice *CompositionService) UnmarshalJSON(b []byte) error {
	type temp CompositionService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*compositionservice = CompositionService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	compositionservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (compositionservice *CompositionService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(CompositionService)
	original.UnmarshalJSON(compositionservice.rawData)

	readWriteFields := []string{
		"AllowOverprovisioning",
		"ReservationDuration",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(compositionservice).Elem()

	return compositionservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetCompositionService will get a CompositionService instance from the service.
func GetCompositionService(c common.Client, uri string) (*CompositionService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var compositionservice CompositionService
	err = json.NewDecoder(resp.Body).Decode(&compositionservice)
	if err != nil {
		return nil, err
	}

	compositionservice.SetClient(c)
	return &compositionservice, nil
}

// ListReferencedCompositionServices gets the collection of CompositionService from
// a provided reference.
func ListReferencedCompositionServices(c common.Client, link string) ([]*CompositionService, error) {
	var result []*CompositionService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, compositionserviceLink := range links.ItemLinks {
		compositionservice, err := GetCompositionService(c, compositionserviceLink)
		if err != nil {
			collectionError.Failures[compositionserviceLink] = err
		} else {
			result = append(result, compositionservice)
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
