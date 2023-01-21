//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Attributes shall contain the list of BIOS attributes and their values as determined by the manufacturer or
// provider. This object shall describe BIOS attribute settings as additional properties. If the object specifies a
// BIOS attribute registry, attributes shall be looked up in that attribute registry by their attribute name.
// Attributes in this attribute registry with the AttributeType of 'Enumeration' shall use valid ValueName values
// in this object, as listed in that attribute registry.
type Attributes struct {
}

// UnmarshalJSON unmarshals a Attributes object from the raw JSON.
func (attributes *Attributes) UnmarshalJSON(b []byte) error {
	type temp Attributes
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*attributes = Attributes(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Bios shall represent BIOS attributes for a Redfish implementation.
type Bios struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AttributeRegistry The link to the attribute registry that lists the metadata describing the BIOS attribute
	// settings in this resource.
	AttributeRegistry string
	// Attributes shall contain the list of BIOS attributes specific to the manufacturer or provider. BIOS attribute
	// settings appear as additional properties in this object, and can be looked up in the attribute registry by their
	// AttributeName.
	Attributes string
	// Description provides a description of this resource.
	Description string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ResetBiosToDefaultsPending shall indicate whether there is a pending request to reset the BIOS attributes to
	// default values. A successful completion of the ResetBios action shall set this property to 'true'. Applying the
	// default attribute values to this resource shall set this property to 'false'. Services may reject modification
	// requests to the settings resource if this property contains 'true'.
	ResetBiosToDefaultsPending bool
}

// UnmarshalJSON unmarshals a Bios object from the raw JSON.
func (bios *Bios) UnmarshalJSON(b []byte) error {
	type temp Bios
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*bios = Bios(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetBios will get a Bios instance from the service.
func GetBios(c common.Client, uri string) (*Bios, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bios Bios
	err = json.NewDecoder(resp.Body).Decode(&bios)
	if err != nil {
		return nil, err
	}

	bios.SetClient(c)
	return &bios, nil
}

// ListReferencedBioss gets the collection of Bios from
// a provided reference.
func ListReferencedBioss(c common.Client, link string) ([]*Bios, error) {
	var result []*Bios
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, biosLink := range links.ItemLinks {
		bios, err := GetBios(c, biosLink)
		if err != nil {
			collectionError.Failures[biosLink] = err
		} else {
			result = append(result, bios)
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
	// ActiveSoftwareImage shall contain a link a resource of type SoftwareInventory that represents the active BIOS
	// firmware image.
	ActiveSoftwareImage string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SoftwareImages shall contain an array of links to resources of type SoftwareInventory that represent the
	// firmware images that apply to this BIOS.
	SoftwareImages []SoftwareInventory
	// SoftwareImages@odata.count
	SoftwareImagesCount int `json:"SoftwareImages@odata.count"`
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
