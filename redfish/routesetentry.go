//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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

// RouteSetEntry This Resource contains the content of a route set in the Redfish Specification.
type RouteSetEntry struct {
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
	// EgressIdentifier shall contain the interface identifier corresponding to this route.
	EgressIdentifier string
	// HopCount shall contain the number of hops to the destination component from the indicated egress interface.
	HopCount int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// VCAction shall contain the index to the VCAT entry corresponding to this route.
	VCAction string
	// Valid shall indicate whether the entry is valid.
	Valid string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a RouteSetEntry object from the raw JSON.
func (routesetentry *RouteSetEntry) UnmarshalJSON(b []byte) error {
	type temp RouteSetEntry
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*routesetentry = RouteSetEntry(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	routesetentry.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (routesetentry *RouteSetEntry) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(RouteSetEntry)
	original.UnmarshalJSON(routesetentry.rawData)

	readWriteFields := []string{
		"EgressIdentifier",
		"HopCount",
		"VCAction",
		"Valid",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(routesetentry).Elem()

	return routesetentry.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetRouteSetEntry will get a RouteSetEntry instance from the service.
func GetRouteSetEntry(c common.Client, uri string) (*RouteSetEntry, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var routesetentry RouteSetEntry
	err = json.NewDecoder(resp.Body).Decode(&routesetentry)
	if err != nil {
		return nil, err
	}

	routesetentry.SetClient(c)
	return &routesetentry, nil
}

// ListReferencedRouteSetEntrys gets the collection of RouteSetEntry from
// a provided reference.
func ListReferencedRouteSetEntrys(c common.Client, link string) ([]*RouteSetEntry, error) {
	var result []*RouteSetEntry
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, routesetentryLink := range links.ItemLinks {
		routesetentry, err := GetRouteSetEntry(c, routesetentryLink)
		if err != nil {
			collectionError.Failures[routesetentryLink] = err
		} else {
			result = append(result, routesetentry)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
