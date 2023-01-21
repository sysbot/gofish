//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

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

// VLAN shall contain any attributes of a VLAN.
type VLAN struct {
	// Tagged shall indicate whether this VLAN is tagged or untagged for this interface.
	Tagged bool
	// VLANEnable shall indicate whether this VLAN is enabled for this VLAN network interface.
	VLANEnable bool
	// VLANId shall contain the ID for this VLAN.
	VLANId VLANId
	// VLANPriority shall contain the priority for this VLAN.
	VLANPriority VLANPriority
}

// UnmarshalJSON unmarshals a VLAN object from the raw JSON.
func (vlan *VLAN) UnmarshalJSON(b []byte) error {
	type temp VLAN
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*vlan = VLAN(t.temp)

	// Extract the links to other entities for later

	return nil
}

// VLanNetworkInterface This resource contains information for a VLAN network instance that is available on a
// manager, system, or other device for a Redfish implementation.
type VLanNetworkInterface struct {
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
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// VLANEnable shall indicate whether this VLAN is enabled for this interface.
	VLANEnable bool
	// VLANId shall contain the ID for this VLAN.
	VLANId VLANId
	// VLANPriority shall contain the priority for this VLAN.
	VLANPriority VLANPriority
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a VLanNetworkInterface object from the raw JSON.
func (vlannetworkinterface *VLanNetworkInterface) UnmarshalJSON(b []byte) error {
	type temp VLanNetworkInterface
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*vlannetworkinterface = VLanNetworkInterface(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	vlannetworkinterface.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (vlannetworkinterface *VLanNetworkInterface) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(VLanNetworkInterface)
	original.UnmarshalJSON(vlannetworkinterface.rawData)

	readWriteFields := []string{
		"VLANEnable",
		"VLANId",
		"VLANPriority",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(vlannetworkinterface).Elem()

	return vlannetworkinterface.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetVLanNetworkInterface will get a VLanNetworkInterface instance from the service.
func GetVLanNetworkInterface(c common.Client, uri string) (*VLanNetworkInterface, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var vlannetworkinterface VLanNetworkInterface
	err = json.NewDecoder(resp.Body).Decode(&vlannetworkinterface)
	if err != nil {
		return nil, err
	}

	vlannetworkinterface.SetClient(c)
	return &vlannetworkinterface, nil
}

// ListReferencedVLanNetworkInterfaces gets the collection of VLanNetworkInterface from
// a provided reference.
func ListReferencedVLanNetworkInterfaces(c common.Client, link string) ([]*VLanNetworkInterface, error) {
	var result []*VLanNetworkInterface
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, vlannetworkinterfaceLink := range links.ItemLinks {
		vlannetworkinterface, err := GetVLanNetworkInterface(c, vlannetworkinterfaceLink)
		if err != nil {
			collectionError.Failures[vlannetworkinterfaceLink] = err
		} else {
			result = append(result, vlannetworkinterface)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
