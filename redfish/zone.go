//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ExternalAccessibility is
type ExternalAccessibility string

const (
	// GloballyAccessibleExternalAccessibility shall indicate that any external entity with the correct access details,
	// which may include authorization information, can access the endpoints that this zone lists, regardless of zone.
	GloballyAccessibleExternalAccessibility ExternalAccessibility = "GloballyAccessible"
	// NonZonedAccessibleExternalAccessibility shall indicate that any external entity that another zone does not
	// explicitly list can access the endpoints that this zone lists.
	NonZonedAccessibleExternalAccessibility ExternalAccessibility = "NonZonedAccessible"
	// ZoneOnlyExternalAccessibility shall indicate that endpoints in this zone are only accessible by endpoints that
	// this zone explicitly lists.
	ZoneOnlyExternalAccessibility ExternalAccessibility = "ZoneOnly"
	// NoInternalRoutingExternalAccessibility shall indicate that implicit routing within this zone is not defined.
	NoInternalRoutingExternalAccessibility ExternalAccessibility = "NoInternalRouting"
)

// ZoneType is
type ZoneType string

const (
	// DefaultZoneType shall indicate a zone in which all endpoints are added by default when instantiated. This value
	// shall only be used for zones subordinate to the fabric collection.
	DefaultZoneType ZoneType = "Default"
	// ZoneOfEndpointsZoneType shall indicate a zone that contains resources of type Endpoint. This value shall only be
	// used for zones subordinate to the fabric collection.
	ZoneOfEndpointsZoneType ZoneType = "ZoneOfEndpoints"
	// ZoneOfZonesZoneType shall indicate a zone that contains resources of type Zone. This value shall only be used
	// for zones subordinate to the fabric collection.
	ZoneOfZonesZoneType ZoneType = "ZoneOfZones"
	// ZoneOfResourceBlocksZoneType shall indicate a zone that contains resources of type ResourceBlock. This value
	// shall only be used for zones subordinate to the composition service.
	ZoneOfResourceBlocksZoneType ZoneType = "ZoneOfResourceBlocks"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// AddressPools shall contain an array of links to resources of type AddressPool with which this zone is
	// associated.
	AddressPools []AddressPool
	// AddressPools@odata.count
	AddressPoolsCount int `json:"AddressPools@odata.count"`
	// ContainedByZones shall contain an array of links to resources of type Zone that represent the zones that contain
	// this zone. The zones referenced by this property shall not be contained by other zones.
	ContainedByZones []Zone
	// ContainedByZones@odata.count
	ContainedByZonesCount int `json:"ContainedByZones@odata.count"`
	// ContainsZones shall contain an array of links to resources of type Zone that represent the zones that are
	// contained by this zone. The zones referenced by this property shall not contain other zones.
	ContainsZones []Zone
	// ContainsZones@odata.count
	ContainsZonesCount int `json:"ContainsZones@odata.count"`
	// Endpoints shall contain an array of links to resources of type Endpoint that this zone contains.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// InvolvedSwitches shall contain an array of links to resources of type Switch in this zone.
	InvolvedSwitches []Switch
	// InvolvedSwitches@odata.count
	InvolvedSwitchesCount int `json:"InvolvedSwitches@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ResourceBlocks shall contain an array of links to resources of type ResourceBlock with which this zone is
	// associated.
	ResourceBlocks []ResourceBlock
	// ResourceBlocks@odata.count
	ResourceBlocksCount int `json:"ResourceBlocks@odata.count"`
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

// Zone shall represent a simple fabric zone for a Redfish implementation.
type Zone struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// DefaultRoutingEnabled shall indicate whether routing within this zone is enabled.
	DefaultRoutingEnabled bool
	// Description provides a description of this resource.
	Description string
	// ExternalAccessibility shall contain and indication of accessibility of endpoints in this zone to endpoints
	// outside of this zone.
	ExternalAccessibility ExternalAccessibility
	// Identifiers shall contain a list of all known durable names for the associated zone.
	Identifiers []Identifier
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// ZoneType shall contain the type of zone that this zone represents.
	ZoneType ZoneType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Zone object from the raw JSON.
func (zone *Zone) UnmarshalJSON(b []byte) error {
	type temp Zone
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*zone = Zone(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	zone.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (zone *Zone) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Zone)
	original.UnmarshalJSON(zone.rawData)

	readWriteFields := []string{
		"DefaultRoutingEnabled",
		"ExternalAccessibility",
		"ZoneType",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(zone).Elem()

	return zone.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetZone will get a Zone instance from the service.
func GetZone(c common.Client, uri string) (*Zone, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var zone Zone
	err = json.NewDecoder(resp.Body).Decode(&zone)
	if err != nil {
		return nil, err
	}

	zone.SetClient(c)
	return &zone, nil
}

// ListReferencedZones gets the collection of Zone from
// a provided reference.
func ListReferencedZones(c common.Client, link string) ([]*Zone, error) {
	var result []*Zone
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, zoneLink := range links.ItemLinks {
		zone, err := GetZone(c, zoneLink)
		if err != nil {
			collectionError.Failures[zoneLink] = err
		} else {
			result = append(result, zone)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
