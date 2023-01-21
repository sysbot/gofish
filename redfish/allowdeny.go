//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AllowType is
type AllowType string

const (
	// AllowAllowType shall be permitted.
	AllowAllowType AllowType = "Allow"
	// DenyAllowType shall not be permitted.
	DenyAllowType AllowType = "Deny"
)

// DataDirection is
type DataDirection string

const (
	// IngressDataDirection Indicates that this limit is enforced on packets and bytes received by the network device
	// function.
	IngressDataDirection DataDirection = "Ingress"
	// EgressDataDirection Indicates that this limit is enforced on packets and bytes transmitted by the network device
	// function.
	EgressDataDirection DataDirection = "Egress"
)

// IPAddressType is
type IPAddressType string

const (
	// IPv4IPAddressType IPv4 addressing is used for all IP-fields in this object.
	IPv4IPAddressType IPAddressType = "IPv4"
	// IPv6IPAddressType IPv6 addressing is used for all IP-fields in this object.
	IPv6IPAddressType IPAddressType = "IPv6"
)

// AllowDeny shall represent an AllowDeny resource in a Redfish implementation.
type AllowDeny struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AllowType shall indicate the type of permission.
	AllowType AllowType
	// Description provides a description of this resource.
	Description string
	// DestinationPortLower shall contain the TCP, UDP, or other destination port to which this rule begins
	// application, inclusive.
	DestinationPortLower int
	// DestinationPortUpper shall contain the TCP, UDP, or other destination port to which this rule ends application,
	// inclusive.
	DestinationPortUpper int
	// Direction shall indicate the direction of the data to which this permission applies for this network device
	// function.
	Direction DataDirection
	// IANAProtocolNumber shall contain the IANA protocol number to which this permission applies.
	IANAProtocolNumber int
	// IPAddressLower shall contain the lower IP address to which this permission applies.
	IPAddressLower string
	// IPAddressType shall contain the type of IP address populated in the IPAddressLower and IPAddressUpper
	// properties. Services shall not permit mixing IPv6 and IPv4 addresses on the same resource.
	IPAddressType IPAddressType
	// IPAddressUpper shall contain the upper IP address to which this permission applies.
	IPAddressUpper string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SourcePortLower shall contain the TCP, UDP, or other source port to which this rule begins application,
	// inclusive.
	SourcePortLower int
	// SourcePortUpper shall contain the TCP, UDP, or other source port to which this rule ends application, inclusive.
	SourcePortUpper int
	// StatefulSession shall indicate if this permission only applies to stateful connection, which are those using
	// SYN, ACK, and FIN.
	StatefulSession bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AllowDeny object from the raw JSON.
func (allowdeny *AllowDeny) UnmarshalJSON(b []byte) error {
	type temp AllowDeny
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*allowdeny = AllowDeny(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	allowdeny.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (allowdeny *AllowDeny) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AllowDeny)
	original.UnmarshalJSON(allowdeny.rawData)

	readWriteFields := []string{
		"AllowType",
		"DestinationPortLower",
		"DestinationPortUpper",
		"Direction",
		"IANAProtocolNumber",
		"IPAddressLower",
		"IPAddressType",
		"IPAddressUpper",
		"SourcePortLower",
		"SourcePortUpper",
		"StatefulSession",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(allowdeny).Elem()

	return allowdeny.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetAllowDeny will get a AllowDeny instance from the service.
func GetAllowDeny(c common.Client, uri string) (*AllowDeny, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var allowdeny AllowDeny
	err = json.NewDecoder(resp.Body).Decode(&allowdeny)
	if err != nil {
		return nil, err
	}

	allowdeny.SetClient(c)
	return &allowdeny, nil
}

// ListReferencedAllowDenys gets the collection of AllowDeny from
// a provided reference.
func ListReferencedAllowDenys(c common.Client, link string) ([]*AllowDeny, error) {
	var result []*AllowDeny
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, allowdenyLink := range links.ItemLinks {
		allowdeny, err := GetAllowDeny(c, allowdenyLink)
		if err != nil {
			collectionError.Failures[allowdenyLink] = err
		} else {
			result = append(result, allowdeny)
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
