//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CertificateUsageType is
type CertificateUsageType string

const (
	// UserCertificateUsageType This certificate is a user certificate like those associated with a manager account.
	UserCertificateUsageType CertificateUsageType = "User"
	// WebCertificateUsageType This certificate is a web or HTTPS certificate like those used for event destinations.
	WebCertificateUsageType CertificateUsageType = "Web"
	// SSHCertificateUsageType This certificate is used for SSH.
	SSHCertificateUsageType CertificateUsageType = "SSH"
	// DeviceCertificateUsageType This certificate is a device type certificate like those associated with SPDM and
	// other standards.
	DeviceCertificateUsageType CertificateUsageType = "Device"
	// PlatformCertificateUsageType This certificate is a platform type certificate like those associated with SPDM and
	// other standards.
	PlatformCertificateUsageType CertificateUsageType = "Platform"
	// BIOSCertificateUsageType This certificate is a BIOS certificate like those associated with UEFI.
	BIOSCertificateUsageType CertificateUsageType = "BIOS"
)

// Certificate shall represent a certificate for a Redfish implementation.
type Certificate struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// CertificateString shall contain the certificate, and the format shall follow the requirements specified by the
	// CertificateType property value. If the certificate contains any private keys, they shall be removed from the
	// string in responses. If the service does not know the private key for the certificate and is needed to use the
	// certificate, the client shall provide the private key as part of the string in the POST request.
	CertificateString string
	// CertificateType shall contain the format type for the certificate.
	CertificateType CertificateType
	// CertificateUsageTypes shall contain an array describing the types or purposes for this certificate.
	CertificateUsageTypes []CertificateUsageType
	// Description provides a description of this resource.
	Description string
	// Fingerprint shall be a string containing the ASCII representation of the fingerprint of the certificate. The
	// hash algorithm used to generate this fingerprint shall be specified by the FingerprintHashAlgorithm property.
	Fingerprint string
	// FingerprintHashAlgorithm shall be a string containing the hash algorithm used for generating the Fingerprint
	// property. The value shall be one of the strings in the 'Algorithm Name' field of the 'TPM_ALG_ID Constants'
	// table within the 'Trusted Computing Group Algorithm Registry'.
	FingerprintHashAlgorithm string
	// Issuer shall contain an object containing information about the issuer of the certificate.
	Issuer string
	// KeyUsage shall contain the key usage extension, which defines the purpose of the public keys in this
	// certificate.
	KeyUsage []KeyUsage
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SPDM shall contain SPDM-related information for the certificate. This property shall only be present for SPDM
	// certificates.
	SPDM SPDM
	// SerialNumber shall be a string containing the ASCII representation of the serial number of the certificate, as
	// defined by the RFC5280 'serialNumber' field.
	SerialNumber string
	// SignatureAlgorithm shall be a string containing the algorithm used for generating the signature of the
	// certificate, as defined by the RFC5280 'signatureAlgorithm' field. The value shall be a string representing the
	// ASN.1 OID of the signature algorithm as defined in, but not limited to, RFC3279, RFC4055, or RFC4491.
	SignatureAlgorithm string
	// Subject shall contain an object containing information about the subject of the certificate.
	Subject string
	// UefiSignatureOwner shall contain the GUID of the UEFI signature owner for this certificate as defined by the
	// UEFI Specification. This property shall only be present for certificates managed by UEFI.
	UefiSignatureOwner string
	// ValidNotAfter shall contain the date when the certificate validity period ends.
	ValidNotAfter string
	// ValidNotBefore shall contain the date when the certificate validity period begins.
	ValidNotBefore string
}

// UnmarshalJSON unmarshals a Certificate object from the raw JSON.
func (certificate *Certificate) UnmarshalJSON(b []byte) error {
	type temp Certificate
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*certificate = Certificate(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetCertificate will get a Certificate instance from the service.
func GetCertificate(c common.Client, uri string) (*Certificate, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var certificate Certificate
	err = json.NewDecoder(resp.Body).Decode(&certificate)
	if err != nil {
		return nil, err
	}

	certificate.SetClient(c)
	return &certificate, nil
}

// ListReferencedCertificates gets the collection of Certificate from
// a provided reference.
func ListReferencedCertificates(c common.Client, link string) ([]*Certificate, error) {
	var result []*Certificate
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, certificateLink := range links.ItemLinks {
		certificate, err := GetCertificate(c, certificateLink)
		if err != nil {
			collectionError.Failures[certificateLink] = err
		} else {
			result = append(result, certificate)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Identifier shall be unique within the managed ecosystem.
type Identifier struct {
	// AdditionalCommonNames shall contain an array of additional common names for the entity, as defined by the
	// RFC5280 'commonName' attribute, in array order as they appear in the certificate. This property shall not be
	// present if only one common name is found. The first common name shall not appear in this property.
	AdditionalCommonNames []string
	// AdditionalOrganizationalUnits shall contain an array of additional organizational units for the entity, as
	// defined by the RFC5280 'organizationalUnitName' attribute, in array order as they appear in the certificate.
	// This property shall not be present if only one organizational unit is found. The first organizational unit shall
	// not appear in this property.
	AdditionalOrganizationalUnits []string
	// City shall contain the city or locality of the organization of the entity, as defined by the RFC5280
	// 'localityName' attribute.
	City string
	// CommonName shall contain the common name of the entity, as defined by the RFC5280 'commonName' attribute.
	CommonName string
	// Country shall contain the two-letter ISO code for the country of the organization of the entity, as defined by
	// the RFC5280 'countryName' attribute.
	Country string
	// DisplayString shall contain a display string that represents the entire identifier. The string should be
	// formatted using industry conventions, such as the single-line human-readable string described by RFC2253 and
	// preserving the field order as shown in the certificate.
	DisplayString string
	// DomainComponents shall contain an array of domain component fields for the entity, as defined by the RFC4519
	// 'domainComponent' attribute, in array order as they appear in the certificate.
	DomainComponents []string
	// Email shall contain the email address of the contact within the organization of the entity, as defined by the
	// RFC2985 'emailAddress' attribute.
	Email string
	// Organization shall contain the name of the organization of the entity, as defined by the RFC5280
	// 'organizationName' attribute.
	Organization string
	// OrganizationalUnit shall contain the name of the unit or division of the organization of the entity, as defined
	// by the RFC5280 'organizationalUnitName' attribute.
	OrganizationalUnit string
	// State shall contain the state, province, or region of the organization of the entity, as defined by the RFC5280
	// 'stateOrProvinceName' attribute.
	State string
}

// UnmarshalJSON unmarshals a Identifier object from the raw JSON.
func (identifier *Identifier) UnmarshalJSON(b []byte) error {
	type temp Identifier
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*identifier = Identifier(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Issuer shall contain a link to a resources of type Certificate that represents the certificate of the CA that
	// issued this certificate.
	Issuer Certificate
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Subjects shall contain an array of links to resources of type Certificate that were issued by the CA that is
	// represented by this certificate.
	Subjects []Certificate
	// Subjects@odata.count
	SubjectsCount int `json:"Subjects@odata.count"`
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

// RekeyResponse shall contain the properties found in the response body for the Rekey action.
type RekeyResponse struct {
	// CSRString shall contain the certificate signing request as a PEM-encoded string, containing structures specified
	// by RFC2986. The private key should not be part of the string.
	CSRString string
	// Certificate shall contain a link to a resource of type Certificate that is replaced after the certificate
	// authority (CA) signs the certificate.
	Certificate string
}

// UnmarshalJSON unmarshals a RekeyResponse object from the raw JSON.
func (rekeyresponse *RekeyResponse) UnmarshalJSON(b []byte) error {
	type temp RekeyResponse
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*rekeyresponse = RekeyResponse(t.temp)

	// Extract the links to other entities for later

	return nil
}

// RenewResponse shall contain the properties found in the response body for the Renew action.
type RenewResponse struct {
	// CSRString shall contain the certificate signing request as a PEM-encoded string, containing structures specified
	// by RFC2986. The private key should not be part of the string.
	CSRString string
	// Certificate shall contain a link to a resource of type Certificate that is replaced after the certificate
	// authority (CA) signs the certificate.
	Certificate string
}

// UnmarshalJSON unmarshals a RenewResponse object from the raw JSON.
func (renewresponse *RenewResponse) UnmarshalJSON(b []byte) error {
	type temp RenewResponse
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*renewresponse = RenewResponse(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SPDM shall contain contain SPDM-related information for a certificate.
type SPDM struct {
	// SlotId shall contain an integer between 0 and 7, inclusive, that represents the slot identifier for an SPDM-
	// provided certificate.
	SlotId int
}

// UnmarshalJSON unmarshals a SPDM object from the raw JSON.
func (spdm *SPDM) UnmarshalJSON(b []byte) error {
	type temp SPDM
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdm = SPDM(t.temp)

	// Extract the links to other entities for later

	return nil
}
