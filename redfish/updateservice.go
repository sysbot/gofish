//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ApplyTime is
type ApplyTime string

const (
	// ImmediateApplyTime shall indicate the HttpPushUri-provided software is applied immediately.
	ImmediateApplyTime ApplyTime = "Immediate"
	// OnResetApplyTime shall indicate the HttpPushUri-provided software is applied when the system or service is
	// reset.
	OnResetApplyTime ApplyTime = "OnReset"
	// AtMaintenanceWindowStartApplyTime shall indicate the HttpPushUri-provided software is applied during the
	// maintenance window specified by the MaintenanceWindowStartTime and MaintenanceWindowDurationInSeconds
	// properties. A service may perform resets during this maintenance window.
	AtMaintenanceWindowStartApplyTime ApplyTime = "AtMaintenanceWindowStart"
	// InMaintenanceWindowOnResetApplyTime shall indicate the HttpPushUri-provided software is applied during the
	// maintenance window specified by the MaintenanceWindowStartTime and MaintenanceWindowDurationInSeconds
	// properties, and if a reset occurs within the maintenance window.
	InMaintenanceWindowOnResetApplyTime ApplyTime = "InMaintenanceWindowOnReset"
	// OnStartUpdateRequestApplyTime shall indicate the HttpPushUri-provided software is applied when the StartUpdate
	// action of the update service is invoked.
	OnStartUpdateRequestApplyTime ApplyTime = "OnStartUpdateRequest"
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
	// NSFTransferProtocolType Network File System (NFS).
	NSFTransferProtocolType TransferProtocolType = "NSF"
	// SCPTransferProtocolType Secure Copy Protocol (SCP).
	SCPTransferProtocolType TransferProtocolType = "SCP"
	// TFTPTransferProtocolType Trivial File Transfer Protocol (TFTP).
	TFTPTransferProtocolType TransferProtocolType = "TFTP"
	// OEMTransferProtocolType A manufacturer-defined protocol.
	OEMTransferProtocolType TransferProtocolType = "OEM"
	// NFSTransferProtocolType Network File System (NFS).
	NFSTransferProtocolType TransferProtocolType = "NFS"
)

// HttpPushUriApplyTime shall contain settings for when to apply HttpPushUri-provided software.
type HttpPushUriApplyTime struct {
	// ApplyTime shall indicate the time when to apply the HttpPushUri-provided software update.
	ApplyTime string
	// MaintenanceWindowDurationInSeconds shall indicate the end of the maintenance window as the number of seconds
	// after the time specified by the MaintenanceWindowStartTime property. This property shall be required if the
	// HttpPushUriApplyTime property value is 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowDurationInSeconds string
	// MaintenanceWindowStartTime shall indicate the date and time when the service can start to apply the HttpPushUri-
	// provided software as part of a maintenance window. This property shall be required if the HttpPushUriApplyTime
	// property value is 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowStartTime string
}

// UnmarshalJSON unmarshals a HttpPushUriApplyTime object from the raw JSON.
func (httppushuriapplytime *HttpPushUriApplyTime) UnmarshalJSON(b []byte) error {
	type temp HttpPushUriApplyTime
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*httppushuriapplytime = HttpPushUriApplyTime(t.temp)

	// Extract the links to other entities for later

	return nil
}

// HttpPushUriOptions shall contain settings and requirements of the service for HttpPushUri-provided software
// updates.
type HttpPushUriOptions struct {
	// ForceUpdate shall indicate whether the service should bypass update policies when applying the HttpPushUri-
	// provided image, such as allowing a component to be downgraded. Services may contain update policies that are
	// never bypassed, such as minimum version enforcement. If this property is not present, it shall be assumed to be
	// 'false'.
	ForceUpdate string
	// HttpPushUriApplyTime shall contain settings for when to apply HttpPushUri-provided firmware.
	HttpPushUriApplyTime string
}

// UnmarshalJSON unmarshals a HttpPushUriOptions object from the raw JSON.
func (httppushurioptions *HttpPushUriOptions) UnmarshalJSON(b []byte) error {
	type temp HttpPushUriOptions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*httppushurioptions = HttpPushUriOptions(t.temp)

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

// UpdateParameters shall contain the update parameters when passing the update image when using the URI specified
// by the MultipartHttpPushUri property to push a software image.
type UpdateParameters struct {
	// ForceUpdate shall indicate whether the service should bypass update policies when applying the provided image,
	// such as allowing a component to be downgraded. Services may contain update policies that are never bypassed,
	// such as minimum version enforcement. If the client does not provide this parameter, the service shall default
	// this value to 'false'.
	ForceUpdate string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Targets shall contain zero or more URIs that indicate where to apply the update image when using the URI
	// specified by the MultipartHttpPushUri property to push a software image. These targets should correspond to
	// software inventory instances or their related items. If this property is not present or contains no targets, the
	// service shall apply the software image to all applicable targets, as determined by the service.
	Targets []string
}

// UnmarshalJSON unmarshals a UpdateParameters object from the raw JSON.
func (updateparameters *UpdateParameters) UnmarshalJSON(b []byte) error {
	type temp UpdateParameters
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*updateparameters = UpdateParameters(t.temp)

	// Extract the links to other entities for later

	return nil
}

// UpdateService shall represent an update service and the properties that affect the service itself for a Redfish
// implementation.
type UpdateService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ClientCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the client identity certificates that are provided to the server referenced by the ImageURI property in
	// SimpleUpdate as part of TLS handshaking.
	ClientCertificates string
	// Description provides a description of this resource.
	Description string
	// FirmwareInventory shall contain a link to a resource collection of type SoftwareInventoryCollection. The
	// resource collection should contain the set of software components generally referred to as platform firmware or
	// that does not execute within a host operating system. Software in this collection is generally updated using
	// platform-specific methods or utilities.
	FirmwareInventory string
	// HttpPushUri shall contain a URI at which the update service supports an HTTP or HTTPS POST of a software image
	// for the purpose of installing software contained within the image. Access to this URI shall require the same
	// privilege as access to the update service. If the service requires the 'Content-Length' header for POST requests
	// to this URI, the service should return HTTP 411 if the client does not include this header in the POST request.
	// The value of this property should not contain a URI of a Redfish resource.
	HttpPushUri string
	// HttpPushUriOptions shall contain options and requirements of the service for HttpPushUri-provided software
	// updates.
	HttpPushUriOptions string
	// HttpPushUriOptionsBusy shall indicate whether a client uses the HttpPushUriOptions properties for software
	// updates. When a client uses any HttpPushUriOptions properties for software updates, it should set this property
	// to 'true'. When a client no longer uses HttpPushUriOptions properties for software updates, it should set this
	// property to 'false'. This property can provide multiple clients a way to negotiate ownership of
	// HttpPushUriOptions properties. Clients can use this property to determine whether another client uses
	// HttpPushUriOptions properties for software updates. This property has no functional requirements for the
	// service.
	HttpPushUriOptionsBusy bool
	// HttpPushUriTargets shall contain zero or more URIs that indicate where to apply the update image when using the
	// URI specified by the HttpPushUri property to push a software image. These targets should correspond to
	// SoftwareInventory instances or their related items. If this property is not present or contains no targets, the
	// service shall apply the software image to all applicable targets, as determined by the service.
	HttpPushUriTargets []string
	// HttpPushUriTargetsBusy shall indicate whether any client has reserved the HttpPushUriTargets property for
	// firmware updates. A client should set this property to 'true' when it uses HttpPushUriTargets for firmware
	// updates. A client should set it to 'false' when it is no longer uses HttpPushUriTargets for updates. The
	// property can provide multiple clients a way to negotiate ownership of HttpPushUriTargets and helps clients
	// determine whether another client is using HttpPushUriTargets to make firmware updates. This property has no
	// functional requirements for the service.
	HttpPushUriTargetsBusy bool
	// MaxImageSizeBytes shall indicate the maximum size of the software update image that clients can send to this
	// update service.
	MaxImageSizeBytes int
	// MultipartHttpPushUri shall contain a URI used to perform a Redfish Specification-defined Multipart HTTP or HTTPS
	// POST of a software image for the purpose of installing software contained within the image. The value of this
	// property should not contain a URI of a Redfish resource.
	MultipartHttpPushUri string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RemoteServerCertificates shall contain a link to a resource collection of type CertificateCollection that
	// represents the server certificates for the server referenced by the ImageURI property in SimpleUpdate. If
	// VerifyRemoteServerCertificate is 'true', services shall compare the certificates in this collection with the
	// certificate obtained during handshaking with the image server in order to verify the identify of the image
	// server prior to transferring the image. If the server cannot be verified, the service shall not send the
	// transfer request. If VerifyRemoteServerCertificate is 'false', the service shall not perform certificate
	// verification.
	RemoteServerCertificates string
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// SoftwareInventory shall contain a link to a resource collection of type SoftwareInventoryCollection. The
	// resource collection should contain the set of software components executed in the context of a host operating
	// system. This can include device drivers, applications, or offload workloads. Software in this collection is
	// generally updated using operating system-centric methods.
	SoftwareInventory string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// VerifyRemoteServerCertificate shall indicate whether whether the service will verify the certificate of the
	// server referenced by the ImageURI property in SimpleUpdate prior to sending the transfer request. If this
	// property is not supported by the service, it shall be assumed to be 'false'. This property should default to
	// 'false' in order to maintain compatibility with older clients.
	VerifyRemoteServerCertificate bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a UpdateService object from the raw JSON.
func (updateservice *UpdateService) UnmarshalJSON(b []byte) error {
	type temp UpdateService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*updateservice = UpdateService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	updateservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (updateservice *UpdateService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(UpdateService)
	original.UnmarshalJSON(updateservice.rawData)

	readWriteFields := []string{
		"HttpPushUriOptionsBusy",
		"HttpPushUriTargets",
		"HttpPushUriTargetsBusy",
		"ServiceEnabled",
		"VerifyRemoteServerCertificate",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(updateservice).Elem()

	return updateservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetUpdateService will get a UpdateService instance from the service.
func GetUpdateService(c common.Client, uri string) (*UpdateService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updateservice UpdateService
	err = json.NewDecoder(resp.Body).Decode(&updateservice)
	if err != nil {
		return nil, err
	}

	updateservice.SetClient(c)
	return &updateservice, nil
}

// ListReferencedUpdateServices gets the collection of UpdateService from
// a provided reference.
func ListReferencedUpdateServices(c common.Client, link string) ([]*UpdateService, error) {
	var result []*UpdateService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, updateserviceLink := range links.ItemLinks {
		updateservice, err := GetUpdateService(c, updateserviceLink)
		if err != nil {
			collectionError.Failures[updateserviceLink] = err
		} else {
			result = append(result, updateservice)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
