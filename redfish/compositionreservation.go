//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CompositionReservation This resource represents the composition reservation of the composition service for a
// Redfish implementation.
type CompositionReservation struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Client shall contain the client that owns the reservation. The service shall determine this value based on the
	// client that invoked the Compose action that resulted in the creation of this reservation.
	Client string
	// Description provides a description of this resource.
	Description string
	// Manifest shall contain the manifest document processed by the service that resulted in this reservation. This
	// property shall be required if the RequestFormat parameter in the Compose action request contained the value
	// 'Manifest'.
	Manifest string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ReservationTime shall indicate the date and time when the reservation was created by the service.
	ReservationTime string
	// ReservedResourceBlocks shall contain an array of links to resources of type ResourceBlock that represent the
	// reserved resource blocks for this reservation. Upon deletion of the reservation or when the reservation is
	// applied, the Reserved property in the referenced resource blocks shall change to 'false'.
	ReservedResourceBlocks []ResourceBlock
	// ReservedResourceBlocks@odata.count
	ReservedResourceBlocksCount int `json:"ReservedResourceBlocks@odata.count"`
}

// UnmarshalJSON unmarshals a CompositionReservation object from the raw JSON.
func (compositionreservation *CompositionReservation) UnmarshalJSON(b []byte) error {
	type temp CompositionReservation
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*compositionreservation = CompositionReservation(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetCompositionReservation will get a CompositionReservation instance from the service.
func GetCompositionReservation(c common.Client, uri string) (*CompositionReservation, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var compositionreservation CompositionReservation
	err = json.NewDecoder(resp.Body).Decode(&compositionreservation)
	if err != nil {
		return nil, err
	}

	compositionreservation.SetClient(c)
	return &compositionreservation, nil
}

// ListReferencedCompositionReservations gets the collection of CompositionReservation from
// a provided reference.
func ListReferencedCompositionReservations(c common.Client, link string) ([]*CompositionReservation, error) {
	var result []*CompositionReservation
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, compositionreservationLink := range links.ItemLinks {
		compositionreservation, err := GetCompositionReservation(c, compositionreservationLink)
		if err != nil {
			collectionError.Failures[compositionreservationLink] = err
		} else {
			result = append(result, compositionreservation)
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
