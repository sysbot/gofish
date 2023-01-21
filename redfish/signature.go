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

// Signature This resource contains a signature for a Redfish implementation.
type Signature struct {
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
	// SignatureString shall contain the string of the signature, and the format shall follow the requirements
	// specified by the value of the SignatureType property. If the signature contains any private keys, they shall be
	// removed from the string in responses. If the private key for the signature is not known by the service and is
	// needed to use the signature, the client shall provide the private key as part of the string in the POST request.
	SignatureString string
	// SignatureType shall contain the format type for the signature. The format is qualified by the value of the
	// SignatureTypeRegistry property.
	SignatureType string
	// SignatureTypeRegistry shall contain the type for the signature.
	SignatureTypeRegistry SignatureTypeRegistry
	// UefiSignatureOwner shall contain the GUID of the UEFI signature owner for this signature as defined by the UEFI
	// Specification. This property shall only be present if the SignatureTypeRegistry property is 'UEFI'.
	UefiSignatureOwner string
}

// UnmarshalJSON unmarshals a Signature object from the raw JSON.
func (signature *Signature) UnmarshalJSON(b []byte) error {
	type temp Signature
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*signature = Signature(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetSignature will get a Signature instance from the service.
func GetSignature(c common.Client, uri string) (*Signature, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var signature Signature
	err = json.NewDecoder(resp.Body).Decode(&signature)
	if err != nil {
		return nil, err
	}

	signature.SetClient(c)
	return &signature, nil
}

// ListReferencedSignatures gets the collection of Signature from
// a provided reference.
func ListReferencedSignatures(c common.Client, link string) ([]*Signature, error) {
	var result []*Signature
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, signatureLink := range links.ItemLinks {
		signature, err := GetSignature(c, signatureLink)
		if err != nil {
			collectionError.Failures[signatureLink] = err
		} else {
			result = append(result, signature)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
