//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Links Add ability to manage spare capacity.
type Links struct {
	// CacheDataVolumes shall be a pointer to the cache data volumes this volume serves as a cache volume. The
	// corresponding VolumeUsage property shall be set to CacheOnly when this property is used.
	CacheDataVolumes []Volume
	// CacheDataVolumes@odata.count
	CacheDataVolumesCount int `json:"CacheDataVolumes@odata.count"`
	// CacheVolumeSource shall be a pointer to the cache volume source for this volume. The corresponding VolumeUsage
	// property shall be set to Data when this property is used.
	CacheVolumeSource Volume
	// ClassOfService shall contain a reference to the ClassOfService that this storage volume conforms to.
	ClassOfService string
	// ClientEndpoints shall be references to the client Endpoints this volume is associated with.
	ClientEndpoints []Endpoint
	// ClientEndpoints@odata.count
	ClientEndpointsCount int `json:"ClientEndpoints@odata.count"`
	// ConsistencyGroups shall be references to the ConsistencyGroups this volume is associated with.
	ConsistencyGroups []ConsistencyGroup
	// ConsistencyGroups@odata.count
	ConsistencyGroupsCount int `json:"ConsistencyGroups@odata.count"`
	// DedicatedSpareDrives shall be a reference to the resources that this volume is associated with and shall
	// reference resources of type Drive. This property shall only contain references to Drive entities which are
	// currently assigned as a dedicated spare and are able to support this Volume.
	DedicatedSpareDrives []Drive
	// DedicatedSpareDrives@odata.count
	DedicatedSpareDrivesCount int `json:"DedicatedSpareDrives@odata.count"`
	// Drives shall be a reference to the resources that this volume is associated with and shall reference resources
	// of type Drive. This property shall only contain references to Drive entities which are currently members of the
	// Volume, not hot spare Drives which are not currently a member of the volume.
	Drives []Drive
	// Drives@odata.count
	DrivesCount int `json:"Drives@odata.count"`
	// JournalingMedia shall be a pointer to the journaling media used for this Volume to address the write hole issue.
	// Valid when WriteHoleProtectionPolicy property is set to 'Journaling'.
	JournalingMedia Resource
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OwningStorageResource shall be a pointer to the Storage resource that owns or contains this volume.
	OwningStorageResource string
	// OwningStorageService shall be a pointer to the StorageService that owns or contains this volume.
	OwningStorageService string
	// ServerEndpoints shall be references to the server Endpoints this volume is associated with.
	ServerEndpoints []Endpoint
	// ServerEndpoints@odata.count
	ServerEndpointsCount int `json:"ServerEndpoints@odata.count"`
	// SpareResourceSets shall contain resources that may be utilized to replace the capacity provided by a failed
	// resource having a compatible type.
	SpareResourceSets []SpareResourceSet
	// SpareResourceSets@odata.count
	SpareResourceSetsCount int `json:"SpareResourceSets@odata.count"`
	// StorageGroups shall be references to the StorageGroups this volume is associated with.
	StorageGroups []StorageGroup
	// StorageGroups@odata.count
	StorageGroupsCount int `json:"StorageGroups@odata.count"`
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

// NVMeNamespaceProperties This contains properties to use when Volume is used to describe an NVMe Namespace.
type NVMeNamespaceProperties struct {
	// FormattedLBASize shall contain the LBA data size and metadata size combination that the namespace has been
	// formatted with. This is a 4-bit data structure.
	FormattedLBASize string
	// IsShareable shall indicate whether the namespace is shareable.
	IsShareable bool
	// LBAFormatsSupported shall be a list of the LBA formats supported for the namespace, or potential namespaces.
	LBAFormatsSupported []LBAFormatType
	// MetadataTransferredAtEndOfDataLBA shall indicate whether or not the metadata is transferred at the end of the
	// LBA creating an extended data LBA.
	MetadataTransferredAtEndOfDataLBA bool
	// NVMeVersion shall contain the version of the NVMe Base Specification supported.
	NVMeVersion string
	// NamespaceFeatures shall contain a set of Namespace Features.
	NamespaceFeatures NamespaceFeatures
	// NamespaceId shall contain the NVMe Namespace Identifier for this namespace. This property shall be a hex value.
	// Namespace identifiers are not durable and do not have meaning outside the scope of the NVMe subsystem. NSID 0x0,
	// 0xFFFFFFFF, 0xFFFFFFFE are special purpose values.
	NamespaceId string
	// NumberLBAFormats shall contain the number of LBA data size and metadata size combinations supported by this
	// namespace. The value of this property is between 0 and 16. LBA formats with an index set beyond this value will
	// not be supported.
	NumberLBAFormats int
}

// UnmarshalJSON unmarshals a NVMeNamespaceProperties object from the raw JSON.
func (nvmenamespaceproperties *NVMeNamespaceProperties) UnmarshalJSON(b []byte) error {
	type temp NVMeNamespaceProperties
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*nvmenamespaceproperties = NVMeNamespaceProperties(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NamespaceFeatures
type NamespaceFeatures struct {
	// SupportsAtomicTransactionSize shall indicate whether or not the NVM fields for Namespace preferred write
	// granularity (NPWG), write alignment (NPWA), deallocate granularity (NPDG), deallocate alignment (NPDA) and
	// optimal write size (NOWS) are defined for this namespace and should be used by the host for I/O optimization.
	SupportsAtomicTransactionSize bool
	// SupportsDeallocatedOrUnwrittenLBError shall indicate that the controller supports deallocated or unwritten
	// logical block error for this namespace.
	SupportsDeallocatedOrUnwrittenLBError bool
	// SupportsIOPerformanceHints shall indicate that the Namespace Atomic Write Unit Normal (NAWUN), Namespace Atomic
	// Write Unit Power Fail (NAWUPF), and Namespace Atomic Compare and Write Unit (NACWU) fields are defined for this
	// namespace and should be used by the host for this namespace instead of the controller-level properties AWUN,
	// AWUPF, and ACWU.
	SupportsIOPerformanceHints bool
	// SupportsNGUIDReuse shall indicate that the namespace supports the use of an NGUID (namespace globally unique
	// identifier) value.
	SupportsNGUIDReuse bool
	// SupportsThinProvisioning shall indicate whether or not the NVMe Namespace supports thin provisioning.
	// Specifically, the namespace capacity reported may be less than the namespace size.
	SupportsThinProvisioning bool
}

// UnmarshalJSON unmarshals a NamespaceFeatures object from the raw JSON.
func (namespacefeatures *NamespaceFeatures) UnmarshalJSON(b []byte) error {
	type temp NamespaceFeatures
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*namespacefeatures = NamespaceFeatures(t.temp)

	// Extract the links to other entities for later

	return nil
}

// OemActions
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

// Operation
type Operation struct {
	// AssociatedFeaturesRegistry A reference to the task associated with the operation if any.
	AssociatedFeaturesRegistry string
	// OperationName The name of the operation.
	OperationName string
	// PercentageComplete The percentage of the operation that has been completed.
	PercentageComplete int
}

// UnmarshalJSON unmarshals a Operation object from the raw JSON.
func (operation *Operation) UnmarshalJSON(b []byte) error {
	type temp Operation
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*operation = Operation(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Volume shall be used to represent a volume, virtual disk, logical disk, LUN, or other logical storage for a
// Redfish implementation.
type Volume struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccessCapabilities shall specify a current storage access capability.
	AccessCapabilities []StorageAccessCapability
	// Actions shall contain the available actions for this resource.
	Actions string
	// AllocatedPools shall contain references to all storage pools allocated from this volume.
	AllocatedPools string
	// BlockSizeBytes shall contain size of the smallest addressable unit of the associated volume.
	BlockSizeBytes int
	// Capacity Information about the utilization of capacity allocated to this storage volume.
	Capacity string
	// CapacityBytes shall contain the size in bytes of the associated volume.
	CapacityBytes int
	// CapacitySources Fully or partially consumed storage from a source resource. Each entry provides capacity
	// allocation information from a named source resource.
	CapacitySources []CapacitySource
	// CapacitySources@odata.count
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// Compressed shall contain a boolean indicator if the Volume is currently utilizing compression or not.
	Compressed bool
	// Deduplicated shall contain a boolean indicator if the Volume is currently utilizing deduplication or not.
	Deduplicated bool
	// Description provides a description of this resource.
	Description string
	// DisplayName shall contain a user-configurable string to name the volume.
	DisplayName string
	// Encrypted shall contain a boolean indicator if the Volume is currently utilizing encryption or not.
	Encrypted bool
	// EncryptionTypes shall contain the types of encryption used by this Volume.
	EncryptionTypes []EncryptionTypes
	// IOPerfModeEnabled shall indicate whether IO performance mode is enabled for the volume.
	IOPerfModeEnabled bool
	// IOStatistics shall represent IO statistics for this volume.
	IOStatistics string
	// Identifiers shall contain a list of all known durable names for the associated volume.
	Identifiers []Identifier
	// InitializeMethod shall indicate the initialization method used for this volume. If InitializeMethod is not
	// specified, the InitializeMethod should be Foreground. This value reflects the most recently used Initialization
	// Method, and may be changed using the Initialize Action.
	InitializeMethod InitializeMethod
	// IsBootCapable shall indicate whether or not the Volume contains a boot image and is capable of booting. This
	// property may be settable by an admin or client with visibility into the contents of the volume. This property
	// should only be set to true when VolumeUsage is either not specified, or when VolumeUsage is set to Data or
	// SystemData.
	IsBootCapable bool
	// Links shall contain references to resources that are related to, but not contained by (subordinate to), this
	// resource.
	Links string
	// LogicalUnitNumber shall contain host-visible LogicalUnitNumber assigned to this Volume. This property shall only
	// be used when in a single connect configuration and no StorageGroup configuration is used.
	LogicalUnitNumber int
	// LowSpaceWarningThresholdPercents shall be triggered: Across all CapacitySources entries, percent =
	// (SUM(AllocatedBytes) - SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []string
	// Manufacturer shall contain a value that represents the manufacturer or implementer of the storage volume.
	Manufacturer string
	// MaxBlockSizeBytes shall contain size of the largest addressable unit of this storage volume.
	MaxBlockSizeBytes int
	// MediaSpanCount shall indicate the number of media elements used per span in the secondary RAID for a
	// hierarchical RAID type.
	MediaSpanCount int
	// Model shall represents a specific storage volume implementation.
	Model string
	// NVMeNamespaceProperties shall contain properties to use when Volume is used to describe an NVMe Namespace.
	NVMeNamespaceProperties NVMeNamespaceProperties
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Operations shall contain a list of all currently running on the Volume.
	Operations []Operation
	// OptimumIOSizeBytes shall contain the optimum IO size to use when performing IO on this volume. For logical
	// disks, this is the stripe size. For physical disks, this describes the physical sector size.
	OptimumIOSizeBytes int
	// ProvisioningPolicy shall specify the volume's supported storage allocation policy.
	ProvisioningPolicy ProvisioningPolicy
	// RAIDType shall contain the RAID type of the associated Volume.
	RAIDType RAIDType
	// ReadCachePolicy shall contain a boolean indicator of the read cache policy for the Volume.
	ReadCachePolicy ReadCachePolicyType
	// RecoverableCapacitySourceCount The value is the number of available capacity source resources currently
	// available in the event that an equivalent capacity source resource fails.
	RecoverableCapacitySourceCount int
	// RemainingCapacityPercent shall return {[(SUM(AllocatedBytes) - SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100
	// represented as an integer value.
	RemainingCapacityPercent int
	// ReplicaInfo shall describe the replica relationship between this storage volume and a corresponding source
	// volume.
	ReplicaInfo string
	// ReplicaTargets shall reference the target replicas that are sourced by this replica.
	ReplicaTargets []idRef
	// ReplicaTargets@odata.count
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// Status shall contain the status of the Volume.
	Status common.Status
	// StorageGroups shall contain references to all storage groups that include this volume.
	StorageGroups string
	// StripSizeBytes The number of consecutively addressed virtual disk blocks (bytes) mapped to consecutively
	// addressed blocks on a single member extent of a disk array. Synonym for stripe depth and chunk size.
	StripSizeBytes int
	// VolumeUsage shall contain the volume usage type for the Volume.
	VolumeUsage VolumeUsageType
	// WriteCachePolicy shall contain a boolean indicator of the write cache policy for the Volume.
	WriteCachePolicy WriteCachePolicyType
	// WriteCacheState shall contain the WriteCacheState policy setting for the Volume.
	WriteCacheState WriteCacheStateType
	// WriteHoleProtectionPolicy shall be set to 'Off'.
	WriteHoleProtectionPolicy string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Volume object from the raw JSON.
func (volume *Volume) UnmarshalJSON(b []byte) error {
	type temp Volume
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*volume = Volume(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	volume.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (volume *Volume) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Volume)
	original.UnmarshalJSON(volume.rawData)

	readWriteFields := []string{
		"AccessCapabilities",
		"CapacityBytes",
		"CapacitySources",
		"Compressed",
		"Deduplicated",
		"DisplayName",
		"Encrypted",
		"EncryptionTypes",
		"IOPerfModeEnabled",
		"IsBootCapable",
		"LowSpaceWarningThresholdPercents",
		"ProvisioningPolicy",
		"ReadCachePolicy",
		"RecoverableCapacitySourceCount",
		"StripSizeBytes",
		"WriteCachePolicy",
		"WriteHoleProtectionPolicy",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(volume).Elem()

	return volume.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetVolume will get a Volume instance from the service.
func GetVolume(c common.Client, uri string) (*Volume, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var volume Volume
	err = json.NewDecoder(resp.Body).Decode(&volume)
	if err != nil {
		return nil, err
	}

	volume.SetClient(c)
	return &volume, nil
}

// ListReferencedVolumes gets the collection of Volume from
// a provided reference.
func ListReferencedVolumes(c common.Client, link string) ([]*Volume, error) {
	var result []*Volume
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, volumeLink := range links.ItemLinks {
		volume, err := GetVolume(c, volumeLink)
		if err != nil {
			collectionError.Failures[volumeLink] = err
		} else {
			result = append(result, volume)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
