//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ConnectedVia is
type ConnectedVia string

const (
	// NotConnectedConnectedVia No current connection.
	NotConnectedConnectedVia ConnectedVia = "NotConnected"
	// URIConnectedVia Connected to a URI location.
	URIConnectedVia ConnectedVia = "URI"
	// AppletConnectedVia Connected to a client application.
	AppletConnectedVia ConnectedVia = "Applet"
	// OemConnectedVia Connected through an OEM-defined method.
	OemConnectedVia ConnectedVia = "Oem"
)

// MediaType is
type MediaType string

const (
	// CDMediaType A CD-ROM format (ISO) image.
	CDMediaType MediaType = "CD"
	// FloppyMediaType A floppy disk image.
	FloppyMediaType MediaType = "Floppy"
	// USBStickMediaType An emulation of a USB storage device.
	USBStickMediaType MediaType = "USBStick"
	// DVDMediaType A DVD-ROM format image.
	DVDMediaType MediaType = "DVD"
)

// TransferMethod is
type TransferMethod string

const (
	// StreamTransferMethod Stream image file data from the source URI.
	StreamTransferMethod TransferMethod = "Stream"
	// UploadTransferMethod Upload the entire image file from the source URI to the service.
	UploadTransferMethod TransferMethod = "Upload"
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
	// NFSTransferProtocolType Network File System (NFS).
	NFSTransferProtocolType TransferProtocolType = "NFS"
	// SCPTransferProtocolType Secure Copy Protocol (SCP).
	SCPTransferProtocolType TransferProtocolType = "SCP"
	// TFTPTransferProtocolType Trivial File Transfer Protocol (TFTP).
	TFTPTransferProtocolType TransferProtocolType = "TFTP"
	// OEMTransferProtocolType A manufacturer-defined protocol.
	OEMTransferProtocolType TransferProtocolType = "OEM"
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

// VirtualMedia shall represent a virtual media service for a Redfish implementation.
type VirtualMedia struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that represents the
	// server certificates for the server referenced by the Image property. If VerifyCertificate is 'true', services
	// shall compare the certificates in this collection with the certificate obtained during handshaking with the
	// image server in order to verify the identify of the image server prior to completing the remote media
	// connection. If the server cannot be verified, the service shall not complete the remote media connection. If
	// VerifyCertificate is 'false', the service shall not perform certificate verification.
	Certificates string
	// ClientCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the client identity certificates that are provided to the server referenced by the Image property as part of TLS
	// handshaking.
	ClientCertificates string
	// ConnectedVia shall contain the current connection method from a client to the virtual media that this resource
	// represents.
	ConnectedVia ConnectedVia
	// Description provides a description of this resource.
	Description string
	// Image shall contain the URI of the media attached to the virtual media. This value may specify an absolute URI
	// to remote media or a relative URI to media local to the implementation. A service may allow a relative URI to
	// reference a SoftwareInventory resource. The value 'null' shall indicates no image connection.
	Image string
	// ImageName shall contain the name of the image.
	ImageName string
	// Inserted shall indicate whether media is present in the virtual media device.
	Inserted bool
	// MediaTypes shall contain an array of the supported media types for this connection.
	MediaTypes []MediaType
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Password shall represent the password to access the Image parameter-specified URI. The value shall be null in
	// responses.
	Password string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TransferMethod shall describe how the image transfer occurs.
	TransferMethod TransferMethod
	// TransferProtocolType shall represent the network protocol to use with the specified image URI.
	TransferProtocolType TransferProtocolType
	// UserName shall represent the user name to access the Image parameter-specified URI.
	UserName string
	// VerifyCertificate shall indicate whether whether the service will verify the certificate of the server
	// referenced by the Image property prior to completing the remote media connection. If this property is not
	// supported by the service, it shall be assumed to be 'false'. This property should default to 'false' in order to
	// maintain compatibility with older clients.
	VerifyCertificate bool
	// WriteProtected shall indicate whether the remote device media prevents writing to that media.
	WriteProtected bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a VirtualMedia object from the raw JSON.
func (virtualmedia *VirtualMedia) UnmarshalJSON(b []byte) error {
	type temp VirtualMedia
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*virtualmedia = VirtualMedia(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	virtualmedia.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (virtualmedia *VirtualMedia) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(VirtualMedia)
	original.UnmarshalJSON(virtualmedia.rawData)

	readWriteFields := []string{
		"Image",
		"Inserted",
		"Password",
		"TransferMethod",
		"TransferProtocolType",
		"UserName",
		"VerifyCertificate",
		"WriteProtected",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(virtualmedia).Elem()

	return virtualmedia.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetVirtualMedia will get a VirtualMedia instance from the service.
func GetVirtualMedia(c common.Client, uri string) (*VirtualMedia, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var virtualmedia VirtualMedia
	err = json.NewDecoder(resp.Body).Decode(&virtualmedia)
	if err != nil {
		return nil, err
	}

	virtualmedia.SetClient(c)
	return &virtualmedia, nil
}

// ListReferencedVirtualMedias gets the collection of VirtualMedia from
// a provided reference.
func ListReferencedVirtualMedias(c common.Client, link string) ([]*VirtualMedia, error) {
	var result []*VirtualMedia
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, virtualmediaLink := range links.ItemLinks {
		virtualmedia, err := GetVirtualMedia(c, virtualmediaLink)
		if err != nil {
			collectionError.Failures[virtualmediaLink] = err
		} else {
			result = append(result, virtualmedia)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
