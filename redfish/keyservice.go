//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// KeyService shall represent the key service properties for a Redfish implementation.
type KeyService struct {
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
	// NVMeoFKeyPolicies shall contain a link to a resource collection of type KeyPolicyCollection that contains the
	// NVMe-oF key policies maintained by this service. The KeyPolicyType property for all members of this collection
	// shall contain the value 'NVMeoF'.
	NVMeoFKeyPolicies string
	// NVMeoFSecrets shall contain a link to a resource collection of type KeyCollection that contains the NVMe-oF keys
	// maintained by this service. The KeyType property for all members of this collection shall contain the value
	// 'NVMeoF'.
	NVMeoFSecrets string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a KeyService object from the raw JSON.
func (keyservice *KeyService) UnmarshalJSON(b []byte) error {
	type temp KeyService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*keyservice = KeyService(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetKeyService will get a KeyService instance from the service.
func GetKeyService(c common.Client, uri string) (*KeyService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var keyservice KeyService
	err = json.NewDecoder(resp.Body).Decode(&keyservice)
	if err != nil {
		return nil, err
	}

	keyservice.SetClient(c)
	return &keyservice, nil
}

// ListReferencedKeyServices gets the collection of KeyService from
// a provided reference.
func ListReferencedKeyServices(c common.Client, link string) ([]*KeyService, error) {
	var result []*KeyService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, keyserviceLink := range links.ItemLinks {
		keyservice, err := GetKeyService(c, keyserviceLink)
		if err != nil {
			collectionError.Failures[keyserviceLink] = err
		} else {
			result = append(result, keyservice)
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
