//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// GroupType is
type GroupType string

const (
	// ClientGroupType shall indicate that the endpoint group contains client (initiator) endpoints. If the associated
	// endpoints contain the EntityRole property, the EntityRole property shall contain the value 'Initiator' or
	// 'Both'.
	ClientGroupType GroupType = "Client"
	// ServerGroupType shall indicate that the endpoint group contains server (target) endpoints. If the associated
	// endpoints contain the EntityRole property, the EntityRole property shall contain the value 'Target' or 'Both'.
	ServerGroupType GroupType = "Server"
	// InitiatorGroupType shall indicate that the endpoint group contains initiator endpoints. If the associated
	// endpoints contain the EntityRole property, the EntityRole property shall contain the value 'Initiator' or
	// 'Both'.
	InitiatorGroupType GroupType = "Initiator"
	// TargetGroupType shall indicate that the endpoint group contains target endpoints. If the associated endpoints
	// contain the EntityRole property, the EntityRole property shall contain the value 'Target' or 'Both'.
	TargetGroupType GroupType = "Target"
)

// EndpointGroup shall represent a group of endpoints that are managed as a unit for a Redfish implementation.
type EndpointGroup struct {
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
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// GroupType shall contain the endpoint group type. If this endpoint group represents a SCSI target group, the
	// value of this property shall contain 'Server' or 'Target'.
	GroupType GroupType
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// TargetEndpointGroupIdentifier shall contain a SCSI-defined identifier for this group that corresponds to the
	// TARGET PORT GROUP field in the REPORT TARGET PORT GROUPS response and the TARGET PORT GROUP field in an INQUIRY
	// VPD page 85 response, type 5h identifier. See the INCITS SAM-5 specification. This property may not be present
	// if the endpoint group does not represent a SCSI target group.
	TargetEndpointGroupIdentifier int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EndpointGroup object from the raw JSON.
func (endpointgroup *EndpointGroup) UnmarshalJSON(b []byte) error {
	type temp EndpointGroup
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*endpointgroup = EndpointGroup(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	endpointgroup.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (endpointgroup *EndpointGroup) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(EndpointGroup)
	original.UnmarshalJSON(endpointgroup.rawData)

	readWriteFields := []string{
		"GroupType",
		"TargetEndpointGroupIdentifier",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(endpointgroup).Elem()

	return endpointgroup.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEndpointGroup will get a EndpointGroup instance from the service.
func GetEndpointGroup(c common.Client, uri string) (*EndpointGroup, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var endpointgroup EndpointGroup
	err = json.NewDecoder(resp.Body).Decode(&endpointgroup)
	if err != nil {
		return nil, err
	}

	endpointgroup.SetClient(c)
	return &endpointgroup, nil
}

// ListReferencedEndpointGroups gets the collection of EndpointGroup from
// a provided reference.
func ListReferencedEndpointGroups(c common.Client, link string) ([]*EndpointGroup, error) {
	var result []*EndpointGroup
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, endpointgroupLink := range links.ItemLinks {
		endpointgroup, err := GetEndpointGroup(c, endpointgroupLink)
		if err != nil {
			collectionError.Failures[endpointgroupLink] = err
		} else {
			result = append(result, endpointgroup)
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
	// Connections shall contain an array of links to resources of type Connection that represent the connections to
	// which this endpoint group belongs.
	Connections []Connection
	// Connections@odata.count
	ConnectionsCount int `json:"Connections@odata.count"`
	// Endpoints shall contain an array of links to resources of type Endpoint that represent the endpoints that are in
	// this endpoint group.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
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
