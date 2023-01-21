//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// TransferProtocolType is
type TransferProtocolType string

const (
	// CIFSTransferProtocolType Common Internet File System (CIFS).
	CIFSTransferProtocolType TransferProtocolType = "CIFS"
	// FTPTransferProtocolType File Transfer Protocol (FTP).
	FTPTransferProtocolType TransferProtocolType = "FTP"
	// SFTPTransferProtocolType Secure File Transfer Protocol (SFTP).
	SFTPTransferProtocolType TransferProtocolType = "SFTP"
	// HTTPTransferProtocolType Hypertext Transfer Protocol (HTTP).
	HTTPTransferProtocolType TransferProtocolType = "HTTP"
	// HTTPSTransferProtocolType Hypertext Transfer Protocol Secure (HTTPS).
	HTTPSTransferProtocolType TransferProtocolType = "HTTPS"
	// SCPTransferProtocolType Secure Copy Protocol (SCP).
	SCPTransferProtocolType TransferProtocolType = "SCP"
	// TFTPTransferProtocolType Trivial File Transfer Protocol (TFTP).
	TFTPTransferProtocolType TransferProtocolType = "TFTP"
	// OEMTransferProtocolType A manufacturer-defined protocol.
	OEMTransferProtocolType TransferProtocolType = "OEM"
	// NFSTransferProtocolType Network File System (NFS).
	NFSTransferProtocolType TransferProtocolType = "NFS"
)

// LicenseService shall represent an license service and the properties that affect the service itself for a
// Redfish implementation.
type LicenseService struct {
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
	// LicenseExpirationWarningDays shall contain the number of days prior to a license expiration that the service
	// shall send the 'DaysBeforeExpiration' message from the License Message Registry at least once. A value of zero
	// shall indicate that no warning messages are sent prior to license expiration.
	LicenseExpirationWarningDays int
	// Licenses shall contain a link to a resource collection of type LicenseCollection. When installing a license with
	// a POST operation to this collection, the service may update an existing License resource instead of creating a
	// new resource. In these cases, the service shall respond with the HTTP '200 OK' status code or HTTP '204 No
	// Content' status code and the 'Location' header in the response shall contain the URI of the updated License
	// resource.
	Licenses LicenseCollection
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a LicenseService object from the raw JSON.
func (licenseservice *LicenseService) UnmarshalJSON(b []byte) error {
	type temp LicenseService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*licenseservice = LicenseService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	licenseservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (licenseservice *LicenseService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(LicenseService)
	original.UnmarshalJSON(licenseservice.rawData)

	readWriteFields := []string{
		"LicenseExpirationWarningDays",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(licenseservice).Elem()

	return licenseservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetLicenseService will get a LicenseService instance from the service.
func GetLicenseService(c common.Client, uri string) (*LicenseService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var licenseservice LicenseService
	err = json.NewDecoder(resp.Body).Decode(&licenseservice)
	if err != nil {
		return nil, err
	}

	licenseservice.SetClient(c)
	return &licenseservice, nil
}

// ListReferencedLicenseServices gets the collection of LicenseService from
// a provided reference.
func ListReferencedLicenseServices(c common.Client, link string) ([]*LicenseService, error) {
	var result []*LicenseService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, licenseserviceLink := range links.ItemLinks {
		licenseservice, err := GetLicenseService(c, licenseserviceLink)
		if err != nil {
			collectionError.Failures[licenseserviceLink] = err
		} else {
			result = append(result, licenseservice)
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
