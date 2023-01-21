//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// TrustedComponentType is
type TrustedComponentType string

const (
	// DiscreteTrustedComponentType shall indicate that the entity has a well-defined physical boundary within the
	// chassis.
	DiscreteTrustedComponentType TrustedComponentType = "Discrete"
	// IntegratedTrustedComponentType shall indicate that the entity is integrated into another device.
	IntegratedTrustedComponentType TrustedComponentType = "Integrated"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// ActiveSoftwareImage shall contain a link to a resource of type SoftwareInventory that represents the active
	// firmware image for this trusted component.
	ActiveSoftwareImage string
	// ComponentIntegrity shall contain an array of links to resources of type ComponentIntegrity that represent the
	// communication established with the trusted component by other resources. The TargetComponentURI property in the
	// referenced ComponentIntegrity resources shall reference this trusted component.
	ComponentIntegrity []ComponentIntegrity
	// ComponentIntegrity@odata.count
	ComponentIntegrityCount int `json:"ComponentIntegrity@odata.count"`
	// ComponentsProtected shall contain an array of links to resources whose integrity is measured or reported by the
	// trusted component.
	ComponentsProtected []idRef
	// ComponentsProtected@odata.count
	ComponentsProtectedCount int `json:"ComponentsProtected@odata.count"`
	// IntegratedInto shall contain a link to a resource to which this trusted component is physically integrated. This
	// property shall be present if TrustedComponentType contains 'Integrated'.
	IntegratedInto idRef
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SoftwareImages shall contain an array of links to resource of type SoftwareInventory that represent the firmware
	// images that apply to this trusted component.
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

// TrustedComponent shall represent a trusted component in a Redfish implementation.
type TrustedComponent struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains device
	// identity certificates of the trusted component.
	Certificates string
	// Description provides a description of this resource.
	Description string
	// FirmwareVersion shall contain a version number associated with the active software image on the trusted
	// component.
	FirmwareVersion string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Manufacturer shall contain the name of the organization responsible for producing the trusted component. This
	// organization may be the entity from whom the trusted component is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to the trusted component.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain a part number assigned by the organization that is responsible for producing or
	// manufacturing the trusted component.
	PartNumber string
	// SKU shall contain the stock-keeping unit number for this trusted component.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the trusted component.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TrustedComponentType shall contain the type of trusted component.
	TrustedComponentType string
	// UUID shall contain a universal unique identifier number for the trusted component.
	UUID string
}

// UnmarshalJSON unmarshals a TrustedComponent object from the raw JSON.
func (trustedcomponent *TrustedComponent) UnmarshalJSON(b []byte) error {
	type temp TrustedComponent
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*trustedcomponent = TrustedComponent(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetTrustedComponent will get a TrustedComponent instance from the service.
func GetTrustedComponent(c common.Client, uri string) (*TrustedComponent, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var trustedcomponent TrustedComponent
	err = json.NewDecoder(resp.Body).Decode(&trustedcomponent)
	if err != nil {
		return nil, err
	}

	trustedcomponent.SetClient(c)
	return &trustedcomponent, nil
}

// ListReferencedTrustedComponents gets the collection of TrustedComponent from
// a provided reference.
func ListReferencedTrustedComponents(c common.Client, link string) ([]*TrustedComponent, error) {
	var result []*TrustedComponent
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, trustedcomponentLink := range links.ItemLinks {
		trustedcomponent, err := GetTrustedComponent(c, trustedcomponentLink)
		if err != nil {
			collectionError.Failures[trustedcomponentLink] = err
		} else {
			result = append(result, trustedcomponent)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
