//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ResetToDefaultsType is
type ResetToDefaultsType string

const (
	// ResetAllResetToDefaultsType Reset all settings to factory defaults and remove all volumes.
	ResetAllResetToDefaultsType ResetToDefaultsType = "ResetAll"
	// PreserveVolumesResetToDefaultsType Reset all settings to factory defaults but preserve the configured volumes on
	// the controllers.
	PreserveVolumesResetToDefaultsType ResetToDefaultsType = "PreserveVolumes"
)

// CacheSummary shall contain properties that describe the cache memory for a storage controller.
type CacheSummary struct {
	// PersistentCacheSizeMiB shall contain the amount of cache memory that is persistent as measured in mebibytes.
	// This size shall be less than or equal to the TotalCacheSizeMiB.
	PersistentCacheSizeMiB int
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TotalCacheSizeMiB shall contain the amount of configured cache memory as measured in mebibytes.
	TotalCacheSizeMiB int
}

// UnmarshalJSON unmarshals a CacheSummary object from the raw JSON.
func (cachesummary *CacheSummary) UnmarshalJSON(b []byte) error {
	type temp CacheSummary
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*cachesummary = CacheSummary(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Enclosures shall contain an array of links to resources of type Chassis that represent the physical containers
	// attached to this resource.
	Enclosures []Chassis
	// Enclosures@odata.count
	EnclosuresCount int `json:"Enclosures@odata.count"`
	// HostingStorageSystems shall contain an array of links to resources of type ComputerSystem that represent the
	// storage systems that host this storage subsystem. The members of this array shall be in the StorageSystems
	// resource collection off the service root.
	HostingStorageSystems []ComputerSystem
	// HostingStorageSystems@odata.count
	HostingStorageSystemsCount int `json:"HostingStorageSystems@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SimpleStorage shall contain a link to a resource of type SimpleStorage that represents the same storage
	// subsystem as this resource.
	SimpleStorage string
	// StorageServices shall contain an array of links to resources of type StorageService with which this storage
	// subsystem is associated.
	StorageServices []StorageService
	// StorageServices@odata.count
	StorageServicesCount int `json:"StorageServices@odata.count"`
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

// Rates shall contain all the rate settings available on the controller.
type Rates struct {
	// ConsistencyCheckRatePercent shall contain the percentage of controller resources used for checking data
	// consistency on volumes.
	ConsistencyCheckRatePercent int
	// RebuildRatePercent shall contain the percentage of controller resources used for rebuilding volumes.
	RebuildRatePercent int
	// TransformationRatePercent shall contain the percentage of controller resources used for transforming volumes.
	TransformationRatePercent int
}

// UnmarshalJSON unmarshals a Rates object from the raw JSON.
func (rates *Rates) UnmarshalJSON(b []byte) error {
	type temp Rates
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*rates = Rates(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Storage shall represent a storage subsystem in the Redfish Specification.
type Storage struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ConsistencyGroups shall contain a link to a resource collection of type ConsistencyGroupCollection. The property
	// shall be used when groups of volumes are treated as a single resource by an application or set of applications.
	ConsistencyGroups string
	// Controllers shall contain a link to a resource collection of type StorageControllerCollection that contains the
	// set of storage controllers allocated to this storage subsystem.
	Controllers string
	// Description provides a description of this resource.
	Description string
	// Drives shall contain a set of the drives attached to the storage controllers that this resource represents.
	Drives []Drive
	// Drives@odata.count
	DrivesCount int `json:"Drives@odata.count"`
	// EndpointGroups shall contain a link to a resource collection of type EndpointGroupCollection. This property
	// shall be implemented when atomic control is needed to perform mapping, masking and zoning operations.
	EndpointGroups string
	// FileSystems shall contain a link to a resource collection of type FileSystemCollection. This property shall be
	// used when file systems are shared or exported by the storage subsystem.
	FileSystems string
	// Identifiers shall contain a list of all known durable names for the storage subsystem.
	Identifiers []Identifier
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Redundancy shall contain redundancy information for the storage subsystem.
	Redundancy []Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StorageControllers@odata.count
	StorageControllersCount int `json:"StorageControllers@odata.count"`
	// StorageGroups shall contain a link to a resource collection of type StorageGroupsCollection. This property shall
	// be used when implementing mapping and masking.
	StorageGroups string
	// StoragePools shall contain a link to a resource collection of type StoragePoolCollection. This property shall be
	// used when an abstraction of media, rather than references to individual media, are used as the storage data
	// source.
	StoragePools string
	// Volumes shall contain a link to a resource collection of type VolumeCollection.
	Volumes string
}

// UnmarshalJSON unmarshals a Storage object from the raw JSON.
func (storage *Storage) UnmarshalJSON(b []byte) error {
	type temp Storage
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storage = Storage(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetStorage will get a Storage instance from the service.
func GetStorage(c common.Client, uri string) (*Storage, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storage Storage
	err = json.NewDecoder(resp.Body).Decode(&storage)
	if err != nil {
		return nil, err
	}

	storage.SetClient(c)
	return &storage, nil
}

// ListReferencedStorages gets the collection of Storage from
// a provided reference.
func ListReferencedStorages(c common.Client, link string) ([]*Storage, error) {
	var result []*Storage
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, storageLink := range links.ItemLinks {
		storage, err := GetStorage(c, storageLink)
		if err != nil {
			collectionError.Failures[storageLink] = err
		} else {
			result = append(result, storage)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// StorageController shall represent a resource that represents a storage controller in the Redfish Specification.
type StorageController struct {
	common.Entity
	// Actions shall contain the available actions for this resource.
	Actions string
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// AssetTag shall track the storage controller for inventory purposes.
	AssetTag string
	// CacheSummary shall contain properties that describe the cache memory for this resource.
	CacheSummary string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	Certificates string
	// ControllerRates shall contain all the rate settings available on the controller.
	ControllerRates string
	// FirmwareVersion shall contain the firmware version as defined by the manufacturer for the associated storage
	// controller.
	FirmwareVersion string
	// Identifiers shall contain a list of all known durable names for the associated storage controller.
	Identifiers []Identifier
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated storage controller.
	Location string
	// Manufacturer shall contain the name of the organization responsible for producing the storage controller. This
	// organization may be the entity from which the storage controller is purchased, but this is not necessarily true.
	Manufacturer string
	// MemberId shall contain the unique identifier for this member within an array. For services supporting Redfish
	// v1.6 or higher, this value shall contain the zero-based array index.
	MemberId string
	// Model shall contain the name by which the manufacturer generally refers to the storage controller.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeInterface shall contain details on the PCIe interface that connects this PCIe-based controller to its host.
	PCIeInterface string
	// PartNumber shall contain a part number assigned by the organization that is responsible for producing or
	// manufacturing the storage controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	Ports string
	// SKU shall contain the stock-keeping unit number for this storage storage controller.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the storage controller.
	SerialNumber string
	// SpeedGbps shall represent the maximum supported speed of the storage bus interface, in Gbit/s. The specified
	// interface connects the controller to the storage devices, not the controller to a host. For example, SAS bus not
	// PCIe host bus.
	SpeedGbps float64
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedControllerProtocols shall contain the supported set of protocols for communicating to this storage
	// controller.
	SupportedControllerProtocols []Protocol
	// SupportedDeviceProtocols shall contain the set of protocols this storage controller can use to communicate with
	// attached devices.
	SupportedDeviceProtocols []Protocol
	// SupportedRAIDTypes shall contain an array of all the RAID types supported by this controller.
	SupportedRAIDTypes []RAIDType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a StorageController object from the raw JSON.
func (storagecontroller *StorageController) UnmarshalJSON(b []byte) error {
	type temp StorageController
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storagecontroller = StorageController(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	storagecontroller.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (storagecontroller *StorageController) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(StorageController)
	original.UnmarshalJSON(storagecontroller.rawData)

	readWriteFields := []string{
		"AssetTag",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(storagecontroller).Elem()

	return storagecontroller.Entity.Update(originalElement, currentElement, readWriteFields)
}

// StorageControllerActions shall contain the available actions for this resource.
type StorageControllerActions struct {
	// Oem shall contain the available OEM-specific actions for this resource.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a StorageControllerActions object from the raw JSON.
func (storagecontrolleractions *StorageControllerActions) UnmarshalJSON(b []byte) error {
	type temp StorageControllerActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storagecontrolleractions = StorageControllerActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// StorageControllerLinks shall contain links to resources that are related to but are not contained by, or
// subordinate to, this resource.
type StorageControllerLinks struct {
	// Endpoints shall contain an array of links to resources of type Endpoint with which this controller is
	// associated.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeFunctions shall contain an array of links to resources of type PCIeFunction that represents the PCIe
	// functions associated with this resource.
	PCIeFunctions []PCIeFunction
	// PCIeFunctions@odata.count
	PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	// StorageServices@odata.count
	StorageServicesCount int `json:"StorageServices@odata.count"`
}

// UnmarshalJSON unmarshals a StorageControllerLinks object from the raw JSON.
func (storagecontrollerlinks *StorageControllerLinks) UnmarshalJSON(b []byte) error {
	type temp StorageControllerLinks
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storagecontrollerlinks = StorageControllerLinks(t.temp)

	// Extract the links to other entities for later

	return nil
}

// StorageControllerOemActions shall contain the available OEM-specific actions for this resource.
type StorageControllerOemActions struct {
}

// UnmarshalJSON unmarshals a StorageControllerOemActions object from the raw JSON.
func (storagecontrolleroemactions *StorageControllerOemActions) UnmarshalJSON(b []byte) error {
	type temp StorageControllerOemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storagecontrolleroemactions = StorageControllerOemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}
