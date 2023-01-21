//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ANAAccessState is
type ANAAccessState string

const (
	// OptimizedANAAccessState Commands processed by a controller provide optimized access to any namespace in the ANA
	// group.
	OptimizedANAAccessState ANAAccessState = "Optimized"
	// NonOptimizedANAAccessState Commands processed by a controller that reports this state for an ANA Group provide
	// non-optimized access characteristics, such as lower performance or non-optimal use of subsystem resources, to
	// any namespace in the ANA Group.
	NonOptimizedANAAccessState ANAAccessState = "NonOptimized"
	// InaccessibleANAAccessState Namespaces in this group are inaccessible. Commands are not able to access user data
	// of namespaces in the ANA Group.
	InaccessibleANAAccessState ANAAccessState = "Inaccessible"
	// PersistentLossANAAccessState The group is persistently inaccessible. Commands are persistently not able to
	// access user data of namespaces in the ANA Group.
	PersistentLossANAAccessState ANAAccessState = "PersistentLoss"
)

// NVMeControllerType is
type NVMeControllerType string

const (
	// AdminNVMeControllerType The NVMe controller is an admin controller.
	AdminNVMeControllerType NVMeControllerType = "Admin"
	// DiscoveryNVMeControllerType The NVMe controller is a discovery controller.
	DiscoveryNVMeControllerType NVMeControllerType = "Discovery"
	// IONVMeControllerType The NVMe controller is an IO controller.
	IONVMeControllerType NVMeControllerType = "IO"
)

// ANACharacteristics shall contain the ANA characteristics and volume information for a storage controller.
type ANACharacteristics struct {
	// AccessState shall contain the reported ANA access state.
	AccessState ANAAccessState
	// Volume shall contain a link to a resource of type Volume.
	Volume string
}

// UnmarshalJSON unmarshals a ANACharacteristics object from the raw JSON.
func (anacharacteristics *ANACharacteristics) UnmarshalJSON(b []byte) error {
	type temp ANACharacteristics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*anacharacteristics = ANACharacteristics(t.temp)

	// Extract the links to other entities for later

	return nil
}

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
	// AttachedVolumes shall contain an array of links to resources of type Volume that are attached to this instance
	// of storage controller.
	AttachedVolumes []Volume
	// AttachedVolumes@odata.count
	AttachedVolumesCount int `json:"AttachedVolumes@odata.count"`
	// Batteries shall contain an array of links to resources of type Battery that represent the batteries that provide
	// power to this storage controller during a power loss event, such as with battery-backed RAID controllers. This
	// property shall not be present if the batteries power the containing chassis as a whole rather than the
	// individual storage controller.
	Batteries []Battery
	// Batteries@odata.count
	BatteriesCount int `json:"Batteries@odata.count"`
	// Endpoints shall contain an array of links to resources of type Endpoint with which this controller is
	// associated.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// NetworkDeviceFunctions shall contain an array of links to resources of type NetworkDeviceFunction that represent
	// the devices that provide connectivity to this controller.
	NetworkDeviceFunctions []NetworkDeviceFunction
	// NetworkDeviceFunctions@odata.count
	NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeFunctions shall contain an array of links to resources of type PCIeFunction that represents the PCIe
	// functions associated with this resource.
	PCIeFunctions []PCIeFunction
	// PCIeFunctions@odata.count
	PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
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

// NVMeControllerAttributes shall contain NVMe controller attributes for a storage controller.
type NVMeControllerAttributes struct {
	// ReportsNamespaceGranularity shall indicate whether or not the controller supports reporting of Namespace
	// Granularity.
	ReportsNamespaceGranularity bool
	// ReportsUUIDList shall indicate whether or not the controller supports reporting of a UUID list.
	ReportsUUIDList bool
	// Supports128BitHostId shall indicate whether or not the controller supports a 128-bit Host Identifier.
	Supports128BitHostId bool
	// SupportsEnduranceGroups shall indicate whether or not the controller supports Endurance Groups.
	SupportsEnduranceGroups bool
	// SupportsExceedingPowerOfNonOperationalState shall indicate whether or not the controller supports exceeding
	// Power of Non-Operational State in order to execute controller initiated background operations in a non-
	// operational power state.
	SupportsExceedingPowerOfNonOperationalState bool
	// SupportsNVMSets shall indicate whether or not the controller supports NVM Sets.
	SupportsNVMSets bool
	// SupportsPredictableLatencyMode shall indicate whether or not the controller supports Predictable Latency Mode.
	SupportsPredictableLatencyMode bool
	// SupportsReadRecoveryLevels shall indicate whether or not the controller supports Read Recovery Levels.
	SupportsReadRecoveryLevels bool
	// SupportsReservations shall indicate if the controller supports reservations.
	SupportsReservations bool
	// SupportsSQAssociations shall indicate whether or not the controller supports SQ Associations.
	SupportsSQAssociations bool
	// SupportsTrafficBasedKeepAlive shall indicate whether or not the controller supports restarting Keep Alive Timer
	// if traffic is processed from an admin command or IO during a Keep Alive Timeout interval.
	SupportsTrafficBasedKeepAlive bool
}

// UnmarshalJSON unmarshals a NVMeControllerAttributes object from the raw JSON.
func (nvmecontrollerattributes *NVMeControllerAttributes) UnmarshalJSON(b []byte) error {
	type temp NVMeControllerAttributes
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*nvmecontrollerattributes = NVMeControllerAttributes(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NVMeControllerProperties shall contain NVMe related properties for a storage controller.
type NVMeControllerProperties struct {
	// ANACharacteristics shall contain the ANA characteristics and volume information.
	ANACharacteristics []ANACharacteristics
	// AllocatedCompletionQueues shall contain the number of I/O completion queues allocated to this NVMe I/O
	// controller.
	AllocatedCompletionQueues int
	// AllocatedSubmissionQueues shall contain the number of I/O submission queues allocated to this NVMe I/O
	// controller.
	AllocatedSubmissionQueues int
	// ControllerType shall contain the type of NVMe controller.
	ControllerType NVMeControllerType
	// MaxQueueSize shall contain the maximum individual queue entry size supported per queue. This is a zero-based
	// value, where the minimum value is one, indicating two entries. For PCIe, this applies to both submission and
	// completion queues. For NVMe-oF, this applies to only submission queues.
	MaxQueueSize int
	// NVMeControllerAttributes shall contain NVMe controller attributes.
	NVMeControllerAttributes NVMeControllerAttributes
	// NVMeSMARTCriticalWarnings shall contain the NVMe SMART Critical Warnings for this storage controller. This
	// property can contain possible triggers for the predictive drive failure warning for the corresponding drive.
	NVMeSMARTCriticalWarnings NVMeSMARTCriticalWarnings
	// NVMeVersion shall contain the version of the NVMe Base Specification supported.
	NVMeVersion string
}

// UnmarshalJSON unmarshals a NVMeControllerProperties object from the raw JSON.
func (nvmecontrollerproperties *NVMeControllerProperties) UnmarshalJSON(b []byte) error {
	type temp NVMeControllerProperties
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*nvmecontrollerproperties = NVMeControllerProperties(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NVMeSMARTCriticalWarnings shall contain the NVMe SMART Critical Warnings for a storage controller.
type NVMeSMARTCriticalWarnings struct {
	// MediaInReadOnly shall indicate the media has been placed in read only mode. This is not set when the read-only
	// condition on the media is a result of a change in the write protection state of a namespace.
	MediaInReadOnly bool
	// OverallSubsystemDegraded shall indicate that the NVM subsystem reliability has been compromised.
	OverallSubsystemDegraded bool
	// PMRUnreliable shall indicate that the Persistent Memory Region has become unreliable. PCIe memory reads can
	// return invalid data or generate poisoned PCIe TLP(s). Persistent Memory Region memory writes might not update
	// memory or might update memory with undefined data. The Persistent Memory Region might also have become non-
	// persistent.
	PMRUnreliable bool
	// PowerBackupFailed shall indicate that the volatile memory backup device has failed.
	PowerBackupFailed bool
	// SpareCapacityWornOut shall indicate that the available spare capacity has fallen below the threshold.
	SpareCapacityWornOut bool
}

// UnmarshalJSON unmarshals a NVMeSMARTCriticalWarnings object from the raw JSON.
func (nvmesmartcriticalwarnings *NVMeSMARTCriticalWarnings) UnmarshalJSON(b []byte) error {
	type temp NVMeSMARTCriticalWarnings
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*nvmesmartcriticalwarnings = NVMeSMARTCriticalWarnings(t.temp)

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

// StorageController shall represent a storage controller in the Redfish Specification.
type StorageController struct {
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
	// AssetTag shall track the storage controller for inventory purposes.
	AssetTag string
	// CacheSummary shall contain properties that describe the cache memory for this resource.
	CacheSummary string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	Certificates string
	// ControllerRates shall contain all the rate settings available on the controller.
	ControllerRates string
	// Description provides a description of this resource.
	Description string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this storage controller.
	EnvironmentMetrics string
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
	// Model shall contain the name by which the manufacturer generally refers to the storage controller.
	Model string
	// NVMeControllerProperties shall contain NVMe related properties for this storage controller.
	NVMeControllerProperties string
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

// GetStorageController will get a StorageController instance from the service.
func GetStorageController(c common.Client, uri string) (*StorageController, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storagecontroller StorageController
	err = json.NewDecoder(resp.Body).Decode(&storagecontroller)
	if err != nil {
		return nil, err
	}

	storagecontroller.SetClient(c)
	return &storagecontroller, nil
}

// ListReferencedStorageControllers gets the collection of StorageController from
// a provided reference.
func ListReferencedStorageControllers(c common.Client, link string) ([]*StorageController, error) {
	var result []*StorageController
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, storagecontrollerLink := range links.ItemLinks {
		storagecontroller, err := GetStorageController(c, storagecontrollerLink)
		if err != nil {
			collectionError.Failures[storagecontrollerLink] = err
		} else {
			result = append(result, storagecontroller)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
