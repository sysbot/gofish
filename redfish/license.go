//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AuthorizationScope is This property shall describe the authorization scope for this license.
type AuthorizationScope string

const (
	// DeviceAuthorizationScope shall indicate the license authorizes functionality for one or more specific device
	// instances, listed as values of the AuthorizedDevices property.
	DeviceAuthorizationScope AuthorizationScope = "Device"
	// CapacityAuthorizationScope shall indicate the license authorizes functionality for one or more device instances
	// limited to a maximum number of devices specified by the value of the MaxAuthorizedDevices property.
	CapacityAuthorizationScope AuthorizationScope = "Capacity"
	// ServiceAuthorizationScope shall indicate the license authorizes product-level or service-level functionality for
	// the Redfish service. This may include hardware or software features not tied to a specific device or subsystem.
	// License resources using this value shall not include the AuthorizedDevices nor the MaxAuthorizedDevices
	// properties.
	ServiceAuthorizationScope AuthorizationScope = "Service"
)

// LicenseOrigin is
type LicenseOrigin string

const (
	// BuiltInLicenseOrigin A license was provided with the product.
	BuiltInLicenseOrigin LicenseOrigin = "BuiltIn"
	// InstalledLicenseOrigin A license installed by user.
	InstalledLicenseOrigin LicenseOrigin = "Installed"
)

// LicenseType is
type LicenseType string

const (
	// ProductionLicenseType shall indicate a license purchased or obtained for use in production environments.
	ProductionLicenseType LicenseType = "Production"
	// PrototypeLicenseType shall indicate a license that is designed for the development or internal use.
	PrototypeLicenseType LicenseType = "Prototype"
	// TrialLicenseType shall indicate a trial version of a license.
	TrialLicenseType LicenseType = "Trial"
)

// ContactInfo shall contain contact information for an individual or organization responsible for this license.
type ContactInfo struct {
	// ContactName shall contain the name of a person or organization to contact for information about this license.
	ContactName string
	// EmailAddress shall contain the email address for a person or organization to contact for information about this
	// license.
	EmailAddress string
	// PhoneNumber shall contain the phone number for a person or organization to contact for information about this
	// license.
	PhoneNumber string
}

// UnmarshalJSON unmarshals a ContactInfo object from the raw JSON.
func (contactinfo *ContactInfo) UnmarshalJSON(b []byte) error {
	type temp ContactInfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*contactinfo = ContactInfo(t.temp)

	// Extract the links to other entities for later

	return nil
}

// License shall represent a license for a Redfish implementation.
type License struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AuthorizationScope shall contain the authorization scope of the license.
	AuthorizationScope AuthorizationScope
	// Contact shall contain an object containing information about the contact of the license.
	Contact string
	// Description provides a description of this resource.
	Description string
	// DownloadURI shall contain the URI at which to download the license file, using the Redfish protocol and
	// authentication methods. The service provides this URI for the download of the OEM-specific binary file of
	// license data. An HTTP GET from this URI shall return a response payload of MIME time 'application/octet-stream'.
	DownloadURI string
	// EntitlementId shall contain the entitlement identifier for this license, used to display a license key, partial
	// license key, or other value used to identify or differentiate license instances.
	EntitlementId string
	// ExpirationDate shall contain the date and time when the license expires.
	ExpirationDate string
	// GracePeriodDays shall contain the number of days that the license is still usable after the date and time
	// specified by the ExpirationDate property.
	GracePeriodDays int
	// InstallDate shall contain the date and time when the license was installed.
	InstallDate string
	// LicenseInfoURI shall contain the URI at which to provide more information about the license. The information
	// provided at the URI is intended to be general product-related and not tied to specific user, customer, or
	// license instance.
	LicenseInfoURI string
	// LicenseOrigin shall contain the origin for the license.
	LicenseOrigin LicenseOrigin
	// LicenseString shall contain the Base64-encoded string of the license. This property shall not appear in response
	// payloads.
	LicenseString string
	// LicenseType shall contain the type for the license.
	LicenseType LicenseType
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Manufacturer shall represent the name of the manufacturer or producer of this license.
	Manufacturer string
	// MaxAuthorizedDevices shall contain the maximum number of devices that are authorized by the license. This
	// property shall only be present if the AuthorizationScope property contains the value 'Capacity'.
	MaxAuthorizedDevices int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the license.
	PartNumber string
	// RemainingDuration shall contain the remaining usage duration before the license expires. This property shall
	// only be present for licenses that are based on usage time.
	RemainingDuration string
	// RemainingUseCount shall contain the remaining usage count before the license expires. This property shall only
	// be present for licenses that are based on usage count.
	RemainingUseCount int
	// Removable shall indicate whether a user can remove the license with an HTTP DELETE operation.
	Removable bool
	// SKU shall contain the SKU number for this license.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the license.
	SerialNumber string
	// Status shall contain the status of license.
	Status common.Status
}

// UnmarshalJSON unmarshals a License object from the raw JSON.
func (license *License) UnmarshalJSON(b []byte) error {
	type temp License
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*license = License(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetLicense will get a License instance from the service.
func GetLicense(c common.Client, uri string) (*License, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var license License
	err = json.NewDecoder(resp.Body).Decode(&license)
	if err != nil {
		return nil, err
	}

	license.SetClient(c)
	return &license, nil
}

// ListReferencedLicenses gets the collection of License from
// a provided reference.
func ListReferencedLicenses(c common.Client, link string) ([]*License, error) {
	var result []*License
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, licenseLink := range links.ItemLinks {
		license, err := GetLicense(c, licenseLink)
		if err != nil {
			collectionError.Failures[licenseLink] = err
		} else {
			result = append(result, license)
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
	// AuthorizedDevices shall contain an array of links to devices that are authorized by the license. Clients can
	// provide this property when installing a license to apply the license to specific devices. If not provided when
	// installing a license, the service may determine the devices to which the license applies. This property shall
	// not be present if the AuthorizationScope property contains the value 'Service'.
	AuthorizedDevices []idRef
	// AuthorizedDevices@odata.count
	AuthorizedDevicesCount int `json:"AuthorizedDevices@odata.count"`
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
