//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DataSanitizationType is
type DataSanitizationType string

const (
	// BlockEraseDataSanitizationType shall indicate sanitization is performed by deleting all logical block addresses,
	// including those that are not currently mapping to active addresses, but leaving the data on the drive.
	BlockEraseDataSanitizationType DataSanitizationType = "BlockErase"
	// CryptographicEraseDataSanitizationType shall indicate sanitization is performed by erasing the target data's
	// encryption key leaving only the ciphertext on the drive. For more information, see NIST800-88 and ISO/IEC 27040.
	CryptographicEraseDataSanitizationType DataSanitizationType = "CryptographicErase"
	// OverwriteDataSanitizationType shall indicate sanitization is performed by overwriting data by writing an
	// implementation specific pattern onto all sectors of the drive.
	OverwriteDataSanitizationType DataSanitizationType = "Overwrite"
)

// EncryptionAbility is
type EncryptionAbility string

const (
	// NoneEncryptionAbility The drive is not capable of self-encryption.
	NoneEncryptionAbility EncryptionAbility = "None"
	// SelfEncryptingDriveEncryptionAbility The drive is capable of self-encryption per the Trusted Computing Group's
	// Self Encrypting Drive Standard.
	SelfEncryptingDriveEncryptionAbility EncryptionAbility = "SelfEncryptingDrive"
	// OtherEncryptionAbility The drive is capable of self-encryption through some other means.
	OtherEncryptionAbility EncryptionAbility = "Other"
)

// EncryptionStatus is
type EncryptionStatus string

const (
	// UnecryptedEncryptionStatus The drive is not currently encrypted.
	UnecryptedEncryptionStatus EncryptionStatus = "Unecrypted"
	// UnlockedEncryptionStatus The drive is currently encrypted but the data is accessible to the user in unencrypted
	// form.
	UnlockedEncryptionStatus EncryptionStatus = "Unlocked"
	// LockedEncryptionStatus The drive is currently encrypted and the data is not accessible to the user. However, the
	// system can unlock the drive automatically.
	LockedEncryptionStatus EncryptionStatus = "Locked"
	// ForeignEncryptionStatus The drive is currently encrypted, the data is not accessible to the user, and the system
	// requires user intervention to expose the data.
	ForeignEncryptionStatus EncryptionStatus = "Foreign"
	// UnencryptedEncryptionStatus The drive is not currently encrypted.
	UnencryptedEncryptionStatus EncryptionStatus = "Unencrypted"
)

// HotspareReplacementModeType is
type HotspareReplacementModeType string

const (
	// RevertibleHotspareReplacementModeType The hot spare drive that is commissioned due to a drive failure reverts to
	// a hot spare after the failed drive is replaced and rebuilt.
	RevertibleHotspareReplacementModeType HotspareReplacementModeType = "Revertible"
	// NonRevertibleHotspareReplacementModeType The hot spare drive that is commissioned due to a drive failure remains
	// as a data drive and does not revert to a hot spare if the failed drive is replaced.
	NonRevertibleHotspareReplacementModeType HotspareReplacementModeType = "NonRevertible"
)

// HotspareType is
type HotspareType string

const (
	// NoneHotspareType The drive is not a hot spare.
	NoneHotspareType HotspareType = "None"
	// GlobalHotspareType The drive is serving as a hot spare for all other drives in this storage domain.
	GlobalHotspareType HotspareType = "Global"
	// ChassisHotspareType The drive is serving as a hot spare for all other drives in this storage domain that are
	// contained in the same chassis.
	ChassisHotspareType HotspareType = "Chassis"
	// DedicatedHotspareType The drive is serving as a hot spare for a user-defined set of drives or volumes. Clients
	// cannot specify this value when modifying the HotspareType property. This value is reported as a result of
	// configuring the spare drives within a volume.
	DedicatedHotspareType HotspareType = "Dedicated"
)

// MediaType is
type MediaType string

const (
	// HDDMediaType The drive media type is traditional magnetic platters.
	HDDMediaType MediaType = "HDD"
	// SSDMediaType The drive media type is solid state or flash memory.
	SSDMediaType MediaType = "SSD"
	// SMRMediaType The drive media type is shingled magnetic recording.
	SMRMediaType MediaType = "SMR"
)

// StatusIndicator is
type StatusIndicator string

const (
	// OKStatusIndicator The drive is OK.
	OKStatusIndicator StatusIndicator = "OK"
	// FailStatusIndicator The drive has failed.
	FailStatusIndicator StatusIndicator = "Fail"
	// RebuildStatusIndicator The drive is being rebuilt.
	RebuildStatusIndicator StatusIndicator = "Rebuild"
	// PredictiveFailureAnalysisStatusIndicator The drive still works but is predicted to fail soon.
	PredictiveFailureAnalysisStatusIndicator StatusIndicator = "PredictiveFailureAnalysis"
	// HotspareStatusIndicator The drive has been marked to automatically rebuild and replace a failed drive.
	HotspareStatusIndicator StatusIndicator = "Hotspare"
	// InACriticalArrayStatusIndicator The array to which this drive belongs has been degraded.
	InACriticalArrayStatusIndicator StatusIndicator = "InACriticalArray"
	// InAFailedArrayStatusIndicator The array to which this drive belongs has failed.
	InAFailedArrayStatusIndicator StatusIndicator = "InAFailedArray"
)

// Drive shall represent a drive or other physical storage medium for a Redfish implementation.
type Drive struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// AssetTag shall track the drive for inventory purposes.
	AssetTag string
	// BlockSizeBytes shall contain size of the smallest addressable unit of the associated drive.
	BlockSizeBytes int
	// CapableSpeedGbs shall contain fastest capable bus speed, in gigabit per second (Gbit/s), of the associated
	// drive.
	CapableSpeedGbs float64
	// CapacityBytes shall contain the raw size, in bytes, of the associated drive.
	CapacityBytes int
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	Certificates string
	// Description provides a description of this resource.
	Description string
	// EncryptionAbility shall contain the encryption ability for the associated drive.
	EncryptionAbility EncryptionAbility
	// EncryptionStatus shall contain the encryption status for the associated drive.
	EncryptionStatus EncryptionStatus
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this drive.
	EnvironmentMetrics string
	// FailurePredicted shall indicate whether this drive currently predicts a manufacturer-defined failure.
	FailurePredicted bool
	// HotspareReplacementMode shall indicate whether a commissioned hot spare continues to serve as a hot spare after
	// the failed drive is replaced.
	HotspareReplacementMode HotspareReplacementModeType
	// HotspareType shall contain the hot spare type for the associated drive. If the drive currently serves as a hot
	// spare, its Status.State field shall be 'StandbySpare' and 'Enabled' when it is part of a volume.
	HotspareType HotspareType
	// Identifiers shall contain a list of all known durable names for the associated drive.
	Identifiers []Identifier
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for producing the drive. This organization
	// may be the entity from whom the drive is purchased, but this is not necessarily true.
	Manufacturer string
	// MediaType shall contain the type of media contained in the associated drive.
	MediaType MediaType
	// Model shall contain the name by which the manufacturer generally refers to the drive.
	Model string
	// Multipath shall indicate whether the drive is accessible by an initiator from multiple paths allowing for
	// failover capabilities upon a path failure.
	Multipath bool
	// NegotiatedSpeedGbs shall contain current bus speed, in gigabit per second (Gbit/s), of the associated drive.
	NegotiatedSpeedGbs float64
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Operations shall contain a list of all operations currently running on the Drive.
	Operations []Operations
	// PartNumber shall contain the part number assigned by the organization that is responsible for producing or
	// manufacturing the drive.
	PartNumber string
	// PhysicalLocation shall contain location information of the associated drive.
	PhysicalLocation string
	// PredictedMediaLifeLeftPercent shall contain an indicator of the percentage of life remaining in the drive's
	// media.
	PredictedMediaLifeLeftPercent float64
	// Protocol shall contain the protocol that the associated drive currently uses to communicate to the storage
	// controller for this system.
	Protocol Protocol
	// ReadyToRemove shall indicate whether the system is prepared for the removal of this drive.
	ReadyToRemove bool
	// Revision shall contain the manufacturer-defined revision for the associated drive.
	Revision string
	// RotationSpeedRPM shall contain the rotation speed, in revolutions per minute (RPM), of the associated drive.
	RotationSpeedRPM float64
	// SKU shall contain the stock-keeping unit (SKU) number for this drive.
	SKU string
	// SerialNumber shall contain the manufacturer-allocated number that identifies the drive.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StatusIndicator shall contain the status indicator state for the status indicator associated with this drive.
	// The Redfish.AllowableValues annotation specifies the valid values for this property.
	StatusIndicator StatusIndicator
	// WriteCacheEnabled shall indicate whether the drive write cache is enabled.
	WriteCacheEnabled bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Drive object from the raw JSON.
func (drive *Drive) UnmarshalJSON(b []byte) error {
	type temp Drive
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*drive = Drive(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	drive.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (drive *Drive) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Drive)
	original.UnmarshalJSON(drive.rawData)

	readWriteFields := []string{
		"AssetTag",
		"HotspareReplacementMode",
		"HotspareType",
		"LocationIndicatorActive",
		"ReadyToRemove",
		"StatusIndicator",
		"WriteCacheEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(drive).Elem()

	return drive.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetDrive will get a Drive instance from the service.
func GetDrive(c common.Client, uri string) (*Drive, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var drive Drive
	err = json.NewDecoder(resp.Body).Decode(&drive)
	if err != nil {
		return nil, err
	}

	drive.SetClient(c)
	return &drive, nil
}

// ListReferencedDrives gets the collection of Drive from
// a provided reference.
func ListReferencedDrives(c common.Client, link string) ([]*Drive, error) {
	var result []*Drive
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, driveLink := range links.ItemLinks {
		drive, err := GetDrive(c, driveLink)
		if err != nil {
			collectionError.Failures[driveLink] = err
		} else {
			result = append(result, drive)
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
	// Chassis shall contain a link to a resource of type Chassis that represents the physical container associated
	// with this drive.
	Chassis string
	// Endpoints shall contain an array of links to resources of type Endpoint with which this drive is associated.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// NetworkDeviceFunctions shall contain the array of links to resources of type NetworkDeviceFunction. This
	// property should only be present for drives with network connectivity, such as Ethernet attached drives.
	NetworkDeviceFunctions []NetworkDeviceFunction
	// NetworkDeviceFunctions@odata.count
	NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeFunctions shall link to a resource of type PCIeFunction that represents the PCIe functions associated with
	// this resource.
	PCIeFunctions []PCIeFunction
	// PCIeFunctions@odata.count
	PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	// Storage shall contain a link to a resource of type Storage that represents the storage subsystem to which this
	// drive belongs.
	Storage string
	// StoragePools shall contain an array of links of type StoragePool to which this drive belongs.
	StoragePools []StoragePool
	// StoragePools@odata.count
	StoragePoolsCount int `json:"StoragePools@odata.count"`
	// Volumes shall contain an array of links to resources of type Volume with which this drive is associated. This
	// property shall include all volume resources of which this drive is a member and all volumes for which this drive
	// acts as a spare if the hot spare type is 'Dedicated'.
	Volumes []Volume
	// Volumes@odata.count
	VolumesCount int `json:"Volumes@odata.count"`
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

// Operations shall describe a currently running operation on the resource.
type Operations struct {
	// AssociatedTask shall contain a link to a resource of type Task that represents the task associated with the
	// operation.
	AssociatedTask string
	// OperationName shall contain a string of the name of the operation.
	OperationName string
	// PercentageComplete shall contain an integer of the percentage of the operation that has been completed.
	PercentageComplete int
}

// UnmarshalJSON unmarshals a Operations object from the raw JSON.
func (operations *Operations) UnmarshalJSON(b []byte) error {
	type temp Operations
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*operations = Operations(t.temp)

	// Extract the links to other entities for later

	return nil
}
