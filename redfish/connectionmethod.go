//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ConnectionMethodType is
type ConnectionMethodType string

const (
	// RedfishConnectionMethodType shall indicate the connection method is Redfish.
	RedfishConnectionMethodType ConnectionMethodType = "Redfish"
	// SNMPConnectionMethodType shall indicate the connection method is SNMP.
	SNMPConnectionMethodType ConnectionMethodType = "SNMP"
	// IPMI15ConnectionMethodType shall indicate the connection method is IPMI 1.5.
	IPMI15ConnectionMethodType ConnectionMethodType = "IPMI15"
	// IPMI20ConnectionMethodType shall indicate the connection method is IPMI 2.0.
	IPMI20ConnectionMethodType ConnectionMethodType = "IPMI20"
	// NETCONFConnectionMethodType shall indicate the connection method is NETCONF.
	NETCONFConnectionMethodType ConnectionMethodType = "NETCONF"
	// OEMConnectionMethodType shall indicate the connection method is OEM. The ConnectionMethodVariant property shall
	// contain further identification information.
	OEMConnectionMethodType ConnectionMethodType = "OEM"
)

// ConnectionMethod shall represent a connection method for a Redfish implementation.
type ConnectionMethod struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ConnectionMethodType shall contain an identifier of the connection method.
	ConnectionMethodType ConnectionMethodType
	// ConnectionMethodVariant shall contain an additional identifier of the connection method. This property shall be
	// present if ConnectionMethodType is 'OEM'.
	ConnectionMethodVariant string
	// Description provides a description of this resource.
	Description string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a ConnectionMethod object from the raw JSON.
func (connectionmethod *ConnectionMethod) UnmarshalJSON(b []byte) error {
	type temp ConnectionMethod
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*connectionmethod = ConnectionMethod(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetConnectionMethod will get a ConnectionMethod instance from the service.
func GetConnectionMethod(c common.Client, uri string) (*ConnectionMethod, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var connectionmethod ConnectionMethod
	err = json.NewDecoder(resp.Body).Decode(&connectionmethod)
	if err != nil {
		return nil, err
	}

	connectionmethod.SetClient(c)
	return &connectionmethod, nil
}

// ListReferencedConnectionMethods gets the collection of ConnectionMethod from
// a provided reference.
func ListReferencedConnectionMethods(c common.Client, link string) ([]*ConnectionMethod, error) {
	var result []*ConnectionMethod
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, connectionmethodLink := range links.ItemLinks {
		connectionmethod, err := GetConnectionMethod(c, connectionmethodLink)
		if err != nil {
			collectionError.Failures[connectionmethodLink] = err
		} else {
			result = append(result, connectionmethod)
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
	// AggregationSources shall contain an array of links to resources of type AggregationSource that are using this
	// connection method.
	AggregationSources []AggregationSource
	// AggregationSources@odata.count
	AggregationSourcesCount int `json:"AggregationSources@odata.count"`
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
