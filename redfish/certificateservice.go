//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CertificateService shall represent the certificate service properties for a Redfish implementation.
type CertificateService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// CertificateLocations shall contain a link to a resource of type CertificateLocations.
	CertificateLocations string
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a CertificateService object from the raw JSON.
func (certificateservice *CertificateService) UnmarshalJSON(b []byte) error {
	type temp CertificateService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*certificateservice = CertificateService(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetCertificateService will get a CertificateService instance from the service.
func GetCertificateService(c common.Client, uri string) (*CertificateService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var certificateservice CertificateService
	err = json.NewDecoder(resp.Body).Decode(&certificateservice)
	if err != nil {
		return nil, err
	}

	certificateservice.SetClient(c)
	return &certificateservice, nil
}

// ListReferencedCertificateServices gets the collection of CertificateService from
// a provided reference.
func ListReferencedCertificateServices(c common.Client, link string) ([]*CertificateService, error) {
	var result []*CertificateService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, certificateserviceLink := range links.ItemLinks {
		certificateservice, err := GetCertificateService(c, certificateserviceLink)
		if err != nil {
			collectionError.Failures[certificateserviceLink] = err
		} else {
			result = append(result, certificateservice)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// GenerateCSRResponse shall contain the properties found in the response body for the GenerateCSR action.
type GenerateCSRResponse struct {
	// CSRString shall contain the Privacy Enhanced Mail (PEM)-encoded string, which contains RFC2986-specified
	// structures, of the certificate signing request. The private key should not be part of the string.
	CSRString string
	// CertificateCollection shall contain a link to a resource collection of type CertificateCollection where the
	// certificate is installed after the certificate authority (CA) has signed the certificate.
	CertificateCollection string
}

// UnmarshalJSON unmarshals a GenerateCSRResponse object from the raw JSON.
func (generatecsrresponse *GenerateCSRResponse) UnmarshalJSON(b []byte) error {
	type temp GenerateCSRResponse
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*generatecsrresponse = GenerateCSRResponse(t.temp)

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
