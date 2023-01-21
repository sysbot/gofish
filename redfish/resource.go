//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DurableNameFormat is
type DurableNameFormat string

const (
	// NAADurableNameFormat shall contain a hexadecimal representation of the Name Address Authority structure, as
	// defined in the T11 Fibre Channel - Framing and Signaling - 3 (FC-FS-3) specification. The DurableName property
	// shall follow the regular expression pattern '^(([0-9A-Fa-f]{2}){8}){1,2}$', where the most significant octet is
	// first.
	NAADurableNameFormat DurableNameFormat = "NAA"
	// iQNDurableNameFormat shall be in the iSCSI Qualified Name (iQN) format, as defined in RFC3720 and RFC3721.
	iQNDurableNameFormat DurableNameFormat = "iQN"
	// FCWWNDurableNameFormat shall contain a hexadecimal representation of the World-Wide Name (WWN) format, as
	// defined in the T11 Fibre Channel Physical and Signaling Interface Specification. The DurableName property shall
	// follow the regular expression pattern '^([0-9A-Fa-f]{2}[:-]){7}([0-9A-Fa-f]{2})$', where the most significant
	// octet is first.
	FCWWNDurableNameFormat DurableNameFormat = "FC_WWN"
	// UUIDDurableNameFormat shall contain the hexadecimal representation of the UUID, as defined by RFC4122. The
	// DurableName property shall follow the regular expression pattern
	// '([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})'.
	UUIDDurableNameFormat DurableNameFormat = "UUID"
	// EUIDurableNameFormat shall contain the hexadecimal representation of the IEEE-defined 64-bit Extended Unique
	// Identifier (EUI), as defined in the IEEE's Guidelines for 64-bit Global Identifier (EUI-64) Specification. The
	// DurableName property shall follow the regular expression pattern '^([0-9A-Fa-f]{2}[:-]){7}([0-9A-Fa-f]{2})$',
	// where the most significant octet is first.
	EUIDurableNameFormat DurableNameFormat = "EUI"
	// NQNDurableNameFormat shall be in the NVMe Qualified Name (NQN) format, as defined in the NVN Express over Fabric
	// Specification.
	NQNDurableNameFormat DurableNameFormat = "NQN"
	// NSIDDurableNameFormat shall be in the NVM Namespace Identifier (NSID) format, as defined in the NVN Express
	// Specification.
	NSIDDurableNameFormat DurableNameFormat = "NSID"
	// NGUIDDurableNameFormat shall be in the Namespace Globally Unique Identifier (NGUID), as defined in the NVN
	// Express Specification. The DurableName property shall follow the regular expression pattern
	// '^([0-9A-Fa-f]{2}){16}$', where the most significant octet is first.
	NGUIDDurableNameFormat DurableNameFormat = "NGUID"
	// MACAddressDurableNameFormat shall be a media access control address (MAC address), which is a unique identifier
	// assigned to a network interface controller (NIC) for use as a network address. This value should not be used if
	// a more specific type of identifier is available. The DurableName property shall follow the regular expression
	// pattern '^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$', where the most significant octet is first.
	MACAddressDurableNameFormat DurableNameFormat = "MACAddress"
)

// LocationType is This enumeration shall list the types of locations for a part within an enclosure.
type LocationType string

const (
	// SlotLocationType shall indicate the part is located in a slot.
	SlotLocationType LocationType = "Slot"
	// BayLocationType shall indicate the part is located in a bay.
	BayLocationType LocationType = "Bay"
	// ConnectorLocationType shall indicate the part is located in a connector or port.
	ConnectorLocationType LocationType = "Connector"
	// SocketLocationType shall indicate the part is located in a socket.
	SocketLocationType LocationType = "Socket"
	// BackplaneLocationType shall indicate the part is a backplane in an enclosure.
	BackplaneLocationType LocationType = "Backplane"
	// EmbeddedLocationType shall indicate the part is embedded or otherwise permanently incorporated into a larger
	// part or device. This value shall not be used for parts that can be removed by a user or are considered field-
	// replaceable.
	EmbeddedLocationType LocationType = "Embedded"
)

// Orientation is This enumeration shall list the orientations for the ordering of the LocationOrdinalValue
// property.
type Orientation string

const (
	// FrontToBackOrientation shall indicate the ordering for LocationOrdinalValue is front to back.
	FrontToBackOrientation Orientation = "FrontToBack"
	// BackToFrontOrientation shall indicate the ordering for LocationOrdinalValue is back to front.
	BackToFrontOrientation Orientation = "BackToFront"
	// TopToBottomOrientation shall indicate the ordering for LocationOrdinalValue is top to bottom.
	TopToBottomOrientation Orientation = "TopToBottom"
	// BottomToTopOrientation shall indicate the ordering for LocationOrdinalValue is bottom to top.
	BottomToTopOrientation Orientation = "BottomToTop"
	// LeftToRightOrientation shall indicate the ordering for LocationOrdinalValue is left to right.
	LeftToRightOrientation Orientation = "LeftToRight"
	// RightToLeftOrientation shall indicate the ordering for LocationOrdinalValue is right to left.
	RightToLeftOrientation Orientation = "RightToLeft"
)

// RackUnits is Enumeration literals shall name the type of rack unit in use.
type RackUnits string

const (
	// OpenURackUnits shall be specified in terms of the Open Compute Open Rack Specification.
	OpenURackUnits RackUnits = "OpenU"
	// EIA310RackUnits shall conform to the EIA-310 standard.
	EIA310RackUnits RackUnits = "EIA_310"
)

// Reference is This enumeration shall list the reference areas for the location of the part within an enclosure.
type Reference string

const (
	// TopReference shall indicate the part is in the top of the unit.
	TopReference Reference = "Top"
	// BottomReference shall indicate the part is in the bottom of the unit.
	BottomReference Reference = "Bottom"
	// FrontReference shall indicate the part is in the front of the unit.
	FrontReference Reference = "Front"
	// RearReference shall indicate the part is in the rear of the unit.
	RearReference Reference = "Rear"
	// LeftReference shall indicate the part is on the left side of of the unit.
	LeftReference Reference = "Left"
	// RightReference shall indicate the part is on the right side of the unit.
	RightReference Reference = "Right"
	// MiddleReference shall indicate the part is in the middle of the unit.
	MiddleReference Reference = "Middle"
)

// ContactInfo shall contain contact information for an individual or organization responsible for this resource.
type ContactInfo struct {
	// ContactName shall contain the name of a person or organization to contact for information about this resource.
	ContactName string
	// EmailAddress shall contain the email address for a person or organization to contact for information about this
	// resource.
	EmailAddress string
	// PhoneNumber shall contain the phone number for a person or organization to contact for information about this
	// resource.
	PhoneNumber string
}

// UnmarshalJSON unmarshals a ContactInfo object from the raw JSON.
func (contactinfo *ContactInfo) UnmarshalJSON(b []byte) error {
	type temp ContactInfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*contactinfo = ContactInfo(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Identifier shall be unique within the managed ecosystem.
type Identifier struct {
	// DurableName shall contain the world-wide unique identifier for the resource. The string shall be in the
	// Identifier.DurableNameFormat property value format.
	DurableName string
	// DurableNameFormat shall represent the format of the DurableName property.
	DurableNameFormat DurableNameFormat
}

// UnmarshalJSON unmarshals a Identifier object from the raw JSON.
func (identifier *Identifier) UnmarshalJSON(b []byte) error {
	type temp Identifier
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*identifier = Identifier(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Location shall describe the location of a resource.
type Location struct {
	// AltitudeMeters shall contain the altitude of the resource, in meters units, defined as the elevation above sea
	// level.
	AltitudeMeters float64
	// Contacts shall contain an array of contact information for an individual or organization responsible for this
	// resource.
	Contacts []ContactInfo
	// Latitude shall contain the latitude of the resource specified in degrees using a decimal format and not minutes
	// or seconds.
	Latitude float64
	// Longitude shall contain the longitude of the resource specified in degrees using a decimal format and not
	// minutes or seconds.
	Longitude float64
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartLocation shall contain the part location for a resource within an enclosure. This representation shall
	// indicate the location of a part within a location specified by the Placement property.
	PartLocation string
	// Placement shall contain a place within the addressed location.
	Placement string
	// PostalAddress shall contain a postal address of the resource.
	PostalAddress string
}

// UnmarshalJSON unmarshals a Location object from the raw JSON.
func (location *Location) UnmarshalJSON(b []byte) error {
	type temp Location
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*location = Location(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PartLocation shall describe a location for a resource within an enclosure.
type PartLocation struct {
	// LocationOrdinalValue shall contain the number that represents the location of the part based on the
	// LocationType. LocationOrdinalValue shall be measured based on the Orientation value starting with '0'.
	LocationOrdinalValue int
	// LocationType shall contain the type of location of the part.
	LocationType LocationType
	// Orientation shall contain the orientation for the ordering used by the LocationOrdinalValue property.
	Orientation Orientation
	// Reference shall contain the general location within the unit of the part.
	Reference Reference
	// ServiceLabel shall contain the label assigned for service at the part location.
	ServiceLabel string
}

// UnmarshalJSON unmarshals a PartLocation object from the raw JSON.
func (partlocation *PartLocation) UnmarshalJSON(b []byte) error {
	type temp PartLocation
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*partlocation = PartLocation(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Placement shall describe a location within a resource. Examples include a shelf in a rack.
type Placement struct {
	// AdditionalInfo shall contain additional information, such as Tile, Column (Post), Wall, or other designation
	// that describes a location that cannot be conveyed with other properties defined for the Placement object.
	AdditionalInfo string
	// Rack shall contain the name of the rack within a row.
	Rack string
	// RackOffset shall be measured from bottom to top, starting with 0.
	RackOffset int
	// RackOffsetUnits shall contain a RackUnit enumeration literal that indicates the type of rack units in use.
	RackOffsetUnits RackUnits
	// Row shall contain the name of the row.
	Row string
}

// UnmarshalJSON unmarshals a Placement object from the raw JSON.
func (placement *Placement) UnmarshalJSON(b []byte) error {
	type temp Placement
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*placement = Placement(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PostalAddress shall describe a postal address for a resource. For more information, see RFC5139. Depending on
// use, the instance can represent a past, current, or future location.
type PostalAddress struct {
	common.Entity
	// AdditionalCode shall conform to the RFC5139-defined requirements of the ADDCODE field.
	AdditionalCode string
	// AdditionalInfo shall conform to the requirements of the LOC field as defined in RFC5139. Provides additional
	// information.
	AdditionalInfo string
	// Building shall conform to the RFC5139-defined requirements of the BLD field. Names the building.
	Building string
	// City shall conform to the RFC5139-defined requirements of the A3 field. Names a city, township, or shi (JP).
	City string
	// Community shall conform to the RFC5139-defined requirements of the PCN field. A postal community name.
	Community string
	// Country shall conform to the RFC5139-defined requirements of the Country field.
	Country string
	// District shall conform to the RFC5139-defined requirements of the A2 field. Names a county, parish, gun (JP), or
	// district (IN).
	District string
	// Division shall conform to the RFC5139-defined requirements of the A4 field. Names a city division, borough, city
	// district, ward, or chou (JP).
	Division string
	// Floor shall conform to the RFC5139-defined requirements of the FLR field. Provides a floor designation.
	Floor string
	// HouseNumber shall conform to the RFC5139-defined requirements of the HNO field. The numeric portion of the house
	// number.
	HouseNumber int
	// HouseNumberSuffix shall conform to the RFC5139-defined requirements of the HNS field. Provides a suffix to a
	// house number, (F, B, or 1/2).
	HouseNumberSuffix string
	// Landmark shall conform to the RFC5139-defined requirements of the LMK field. Identifies a landmark or vanity
	// address.
	Landmark string
	// LeadingStreetDirection shall conform to the requirements of the PRD field as defined in RFC5139. Names a leading
	// street direction, (N, W, or SE).
	LeadingStreetDirection string
	// Neighborhood shall conform to the RFC5139-defined requirements of the A5 field. Names a neighborhood or block.
	Neighborhood string
	// POBox shall conform to the RFC5139-defined requirements of the POBOX field. A post office box (PO box).
	POBox string
	// PlaceType shall conform to the RFC5139-defined requirements of the PLC field. Examples include office and
	// residence.
	PlaceType string
	// PostalCode shall conform to the RFC5139-defined requirements of the PC field. A postal code (or zip code).
	PostalCode string
	// Road shall conform to the RFC5139-defined requirements of the RD field. Designates a primary road or street.
	Road string
	// RoadBranch shall conform to the RFC5139-defined requirements of the RDBR field. Shall contain a post office box
	// (PO box) road branch.
	RoadBranch string
	// RoadPostModifier shall conform to the RFC5139-defined requirements of the POM field. For example, Extended.
	RoadPostModifier string
	// RoadPreModifier shall conform to the RFC5139-defined requirements of the PRM field. For example, Old or New.
	RoadPreModifier string
	// RoadSection shall conform to the RFC5139-defined requirements of the RDSEC field. A road section.
	RoadSection string
	// RoadSubBranch shall conform to the RFC5139-defined requirements of the RDSUBBR field.
	RoadSubBranch string
	// Room shall conform to the RFC5139-defined requirements of the ROOM field. A name or number of a room to locate
	// the resource within the unit.
	Room string
	// Seat shall conform to the RFC5139-defined requirements of the SEAT field. A name or number of a seat, such as
	// the desk, cubicle, or workstation.
	Seat string
	// Street shall conform to the RFC5139-defined requirements of the A6 field. Names a street.
	Street string
	// StreetSuffix shall conform to the RFC5139-defined requirements of the STS field. Names a street suffix.
	StreetSuffix string
	// Territory shall conform to the RFC5139-defined requirements of the A1 field when it names a territory, state,
	// region, province, or prefecture within a country.
	Territory string
	// TrailingStreetSuffix shall conform to the RFC5139-defined requirements of the POD field. Names a trailing street
	// suffix.
	TrailingStreetSuffix string
	// Unit shall conform to the RFC5139-defined requirements of the UNIT field. The name or number of a unit, such as
	// the apartment or suite, to locate the resource.
	Unit string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PostalAddress object from the raw JSON.
func (postaladdress *PostalAddress) UnmarshalJSON(b []byte) error {
	type temp PostalAddress
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*postaladdress = PostalAddress(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	postaladdress.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (postaladdress *PostalAddress) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(PostalAddress)
	original.UnmarshalJSON(postaladdress.rawData)

	readWriteFields := []string{
		"AdditionalCode",
		"AdditionalInfo",
		"Building",
		"City",
		"Community",
		"Country",
		"District",
		"Division",
		"Floor",
		"HouseNumber",
		"HouseNumberSuffix",
		"Landmark",
		"LeadingStreetDirection",
		"Neighborhood",
		"POBox",
		"PlaceType",
		"PostalCode",
		"Road",
		"RoadBranch",
		"RoadPostModifier",
		"RoadPreModifier",
		"RoadSection",
		"RoadSubBranch",
		"Room",
		"Seat",
		"Street",
		"StreetSuffix",
		"Territory",
		"TrailingStreetSuffix",
		"Unit",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(postaladdress).Elem()

	return postaladdress.Entity.Update(originalElement, currentElement, readWriteFields)
}

// ReferenceableMember shall contain the location of this element within an item.
type ReferenceableMember struct {
	common.Entity
	// MemberId shall contain the unique identifier for this member within an array. For services supporting Redfish
	// v1.6 or higher, this value shall contain the zero-based array index.
	MemberId string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a ReferenceableMember object from the raw JSON.
func (referenceablemember *ReferenceableMember) UnmarshalJSON(b []byte) error {
	type temp ReferenceableMember
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*referenceablemember = ReferenceableMember(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Resource The base type for resources and members that can be linked to.
type Resource struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a Resource object from the raw JSON.
func (resource *Resource) UnmarshalJSON(b []byte) error {
	type temp Resource
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*resource = Resource(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetResource will get a Resource instance from the service.
func GetResource(c common.Client, uri string) (*Resource, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resource Resource
	err = json.NewDecoder(resp.Body).Decode(&resource)
	if err != nil {
		return nil, err
	}

	resource.SetClient(c)
	return &resource, nil
}

// ListReferencedResources gets the collection of Resource from
// a provided reference.
func ListReferencedResources(c common.Client, link string) ([]*Resource, error) {
	var result []*Resource
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, resourceLink := range links.ItemLinks {
		resource, err := GetResource(c, resourceLink)
		if err != nil {
			collectionError.Failures[resourceLink] = err
		} else {
			result = append(result, resource)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// ResourceCollection
type ResourceCollection struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a ResourceCollection object from the raw JSON.
func (resourcecollection *ResourceCollection) UnmarshalJSON(b []byte) error {
	type temp ResourceCollection
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*resourcecollection = ResourceCollection(t.temp)

	// Extract the links to other entities for later

	return nil
}
