//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CertificateLocations shall represent the Certificate Location Properties for a Redfish implementation.
type CertificateLocations struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// Links shall contain links to Resources that are related to but are not contained by or subordinate to this
	// Resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a CertificateLocations object from the raw JSON.
func (certificatelocations *CertificateLocations) UnmarshalJSON(b []byte) error {
	type temp CertificateLocations
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*certificatelocations = CertificateLocations(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetCertificateLocations will get a CertificateLocations instance from the service.
func GetCertificateLocations(c common.Client, uri string) (*CertificateLocations, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var certificatelocations CertificateLocations
	err = json.NewDecoder(resp.Body).Decode(&certificatelocations)
	if err != nil {
		return nil, err
	}

	certificatelocations.SetClient(c)
	return &certificatelocations, nil
}

// ListReferencedCertificateLocationss gets the collection of CertificateLocations from
// a provided reference.
func ListReferencedCertificateLocationss(c common.Client, link string) ([]*CertificateLocations, error) {
	var result []*CertificateLocations
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, certificatelocationsLink := range links.ItemLinks {
		certificatelocations, err := GetCertificateLocations(c, certificatelocationsLink)
		if err != nil {
			collectionError.Failures[certificatelocationsLink] = err
		} else {
			result = append(result, certificatelocations)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Links shall contain links to Resources that are related to but are not contained by or subordinate to this
// Resource.
type Links struct {
	// Certificates shall contain an array of links to Certificate Resources that are installed on this service.
	Certificates []Certificate
	// Certificates@odata.count
	CertificatesCount int `json:"Certificates@odata.count"`
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

// OemActions shall contain the available OEM-specific actions for this Resource.
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
