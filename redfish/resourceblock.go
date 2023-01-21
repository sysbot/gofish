//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CompositionState is
type CompositionState string

const (
	// ComposingCompositionState Intermediate state indicating composition is in progress.
	ComposingCompositionState CompositionState = "Composing"
	// ComposedAndAvailableCompositionState The resource block is currently participating in one or more compositions,
	// and is available to use in more compositions.
	ComposedAndAvailableCompositionState CompositionState = "ComposedAndAvailable"
	// ComposedCompositionState Final successful state of a resource block that has participated in composition.
	ComposedCompositionState CompositionState = "Composed"
	// UnusedCompositionState The resource block is free and can participate in composition.
	UnusedCompositionState CompositionState = "Unused"
	// FailedCompositionState The final composition resulted in failure and manual intervention might be required to
	// fix it.
	FailedCompositionState CompositionState = "Failed"
	// UnavailableCompositionState The resource block has been made unavailable by the service, such as due to
	// maintenance being performed on the resource block.
	UnavailableCompositionState CompositionState = "Unavailable"
)

// PoolType is
type PoolType string

const (
	// FreePoolType This resource block is in the free pool and is not contributing to any composed resources.
	FreePoolType PoolType = "Free"
	// ActivePoolType This resource block is in the active pool and is contributing to at least one composed resource
	// as a result of a composition request.
	ActivePoolType PoolType = "Active"
	// UnassignedPoolType This resource block is not assigned to any pools.
	UnassignedPoolType PoolType = "Unassigned"
)

// ResourceBlockType is
type ResourceBlockType string

const (
	// ComputeResourceBlockType This resource block contains resources of type 'Processor' and 'Memory' in a manner
	// that creates a compute complex.
	ComputeResourceBlockType ResourceBlockType = "Compute"
	// ProcessorResourceBlockType This resource block contains resources of type 'Processor'.
	ProcessorResourceBlockType ResourceBlockType = "Processor"
	// MemoryResourceBlockType This resource block contains resources of type 'Memory'.
	MemoryResourceBlockType ResourceBlockType = "Memory"
	// NetworkResourceBlockType This resource block contains network resources, such as resource of type
	// 'EthernetInterface' and 'NetworkInterface'.
	NetworkResourceBlockType ResourceBlockType = "Network"
	// StorageResourceBlockType This resource block contains storage resources, such as resources of type 'Storage' and
	// 'SimpleStorage'.
	StorageResourceBlockType ResourceBlockType = "Storage"
	// ComputerSystemResourceBlockType This resource block contains resources of type 'ComputerSystem'.
	ComputerSystemResourceBlockType ResourceBlockType = "ComputerSystem"
	// ExpansionResourceBlockType This resource block is capable of changing over time based on its configuration.
	// Different types of devices within this resource block can be added and removed over time.
	ExpansionResourceBlockType ResourceBlockType = "Expansion"
	// IndependentResourceResourceBlockType This resource block is capable of being consumed as a standalone component.
	// This resource block can represent things such as a software platform on one or more computer systems or an
	// appliance that provides composable resources and other services, and can be managed independently of the Redfish
	// service.
	IndependentResourceResourceBlockType ResourceBlockType = "IndependentResource"
)

// CompositionStatus shall contain properties that describe the high level composition status of the resource
// block.
type CompositionStatus struct {
	// CompositionState shall contain an enumerated value that describes the composition state of the resource block.
	CompositionState CompositionState
	// MaxCompositions shall contain a number indicating the maximum number of compositions in which this resource
	// block can participate simultaneously. Services can have additional constraints that prevent this value from
	// being achieved, such as due to system topology and current composed resource utilization. If SharingCapable is
	// 'false', this value shall be set to '1'. The service shall support this property if SharingCapable supported.
	MaxCompositions int
	// NumberOfCompositions shall contain the number of compositions in which this resource block is currently
	// participating.
	NumberOfCompositions int
	// Reserved shall indicate whether any client has reserved the resource block. A client sets this property after
	// the resource block is identified as composed. It shall provide a way for multiple clients to negotiate the
	// ownership of the resource block.
	Reserved bool
	// SharingCapable shall indicate whether this resource block can participate in multiple compositions
	// simultaneously. If this property is not provided, it shall be assumed that this resource block is not capable of
	// being shared.
	SharingCapable bool
	// SharingEnabled shall indicate whether this resource block can participate in multiple compositions
	// simultaneously. The service shall reject modifications of this property with HTTP 400 Bad Request if this
	// resource block is already being used as part of a composed resource. If 'false', the service shall not use the
	// 'ComposedAndAvailable' state for this resource block.
	SharingEnabled bool
}

// UnmarshalJSON unmarshals a CompositionStatus object from the raw JSON.
func (compositionstatus *CompositionStatus) UnmarshalJSON(b []byte) error {
	type temp CompositionStatus
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*compositionstatus = CompositionStatus(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Chassis shall contain an array of links to resources of type Chassis that represent the physical container
	// associated with this resource block.
	Chassis []Chassis
	// Chassis@odata.count
	ChassisCount int `json:"Chassis@odata.count"`
	// ComputerSystems shall contain an array of links to resources of type ComputerSystem that represent the computer
	// systems composed from this resource block.
	ComputerSystems []ComputerSystem
	// ComputerSystems@odata.count
	ComputerSystemsCount int `json:"ComputerSystems@odata.count"`
	// ConsumingResourceBlocks shall contain an array of links to resources of type ResourceBlock that represent the
	// resource blocks that depend on this resource block as a component.
	ConsumingResourceBlocks []ResourceBlock
	// ConsumingResourceBlocks@odata.count
	ConsumingResourceBlocksCount int `json:"ConsumingResourceBlocks@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SupplyingResourceBlocks shall contain an array of links to resources of type ResourceBlock that represent the
	// resource blocks that this resource block depends on as components.
	SupplyingResourceBlocks []ResourceBlock
	// SupplyingResourceBlocks@odata.count
	SupplyingResourceBlocksCount int `json:"SupplyingResourceBlocks@odata.count"`
	// Zones shall contain an array of links to resources of type Zone that represent the binding constraints
	// associated with this resource block.
	Zones []Zone
	// Zones@odata.count
	ZonesCount int `json:"Zones@odata.count"`
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

// ResourceBlock shall represent a resource block for a Redfish implementation.
type ResourceBlock struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Client shall contain the client to which this resource block is assigned.
	Client string
	// CompositionStatus shall contain composition status information about this resource block.
	CompositionStatus string
	// ComputerSystems shall contain an array of links to resource of type ComputerSystem that this resource block
	// contains.
	ComputerSystems []ComputerSystem
	// ComputerSystems@odata.count
	ComputerSystemsCount int `json:"ComputerSystems@odata.count"`
	// Description provides a description of this resource.
	Description string
	// Drives shall contain an array of links to resource of type Drive that this resource block contains.
	Drives []Drive
	// Drives@odata.count
	DrivesCount int `json:"Drives@odata.count"`
	// EthernetInterfaces shall contain an array of links to resource of type EthernetInterface that this resource
	// block contains.
	EthernetInterfaces []EthernetInterface
	// EthernetInterfaces@odata.count
	EthernetInterfacesCount int `json:"EthernetInterfaces@odata.count"`
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Memory shall contain an array of links to resource of type Memory that this resource block contains.
	Memory []Memory
	// Memory@odata.count
	MemoryCount int `json:"Memory@odata.count"`
	// NetworkInterfaces shall contain an array of links to resource of type NetworkInterface that this resource block
	// contains.
	NetworkInterfaces []NetworkInterface
	// NetworkInterfaces@odata.count
	NetworkInterfacesCount int `json:"NetworkInterfaces@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Pool shall contain the pool to which this resource block belongs. If this resource block is not assigned to a
	// client, this property shall contain the value 'Unassigned'. If this resource block is assigned to a client, this
	// property shall not contain the value 'Unassigned'.
	Pool PoolType
	// Processors shall contain an array of links to resource of type Processor that this resource block contains.
	Processors []Processor
	// Processors@odata.count
	ProcessorsCount int `json:"Processors@odata.count"`
	// ResourceBlockType shall contain an array of enumerated values that describe the type of resources available.
	ResourceBlockType []ResourceBlockType
	// SimpleStorage shall contain an array of links to resource of type SimpleStorage that this resource block
	// contains.
	SimpleStorage []SimpleStorage
	// SimpleStorage@odata.count
	SimpleStorageCount int `json:"SimpleStorage@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Storage shall contain an array of links to resource of type Storage that this resource block contains.
	Storage []Storage
	// Storage@odata.count
	StorageCount int `json:"Storage@odata.count"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ResourceBlock object from the raw JSON.
func (resourceblock *ResourceBlock) UnmarshalJSON(b []byte) error {
	type temp ResourceBlock
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*resourceblock = ResourceBlock(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	resourceblock.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (resourceblock *ResourceBlock) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ResourceBlock)
	original.UnmarshalJSON(resourceblock.rawData)

	readWriteFields := []string{
		"Client",
		"Pool",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(resourceblock).Elem()

	return resourceblock.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetResourceBlock will get a ResourceBlock instance from the service.
func GetResourceBlock(c common.Client, uri string) (*ResourceBlock, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resourceblock ResourceBlock
	err = json.NewDecoder(resp.Body).Decode(&resourceblock)
	if err != nil {
		return nil, err
	}

	resourceblock.SetClient(c)
	return &resourceblock, nil
}

// ListReferencedResourceBlocks gets the collection of ResourceBlock from
// a provided reference.
func ListReferencedResourceBlocks(c common.Client, link string) ([]*ResourceBlock, error) {
	var result []*ResourceBlock
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, resourceblockLink := range links.ItemLinks {
		resourceblock, err := GetResourceBlock(c, resourceblockLink)
		if err != nil {
			collectionError.Failures[resourceblockLink] = err
		} else {
			result = append(result, resourceblock)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// ResourceBlockLimits shall specify the allowable quantities of types of resource blocks for a given composition
// request.
type ResourceBlockLimits struct {
	// MaxCompute shall contain an integer that specifies the maximum number of resource blocks of type 'Compute'
	// allowed for the composition request.
	MaxCompute int
	// MaxComputerSystem shall contain an integer that specifies the maximum number of resource blocks of type
	// 'ComputerSystem' allowed for the composition request.
	MaxComputerSystem int
	// MaxExpansion shall contain an integer that specifies the maximum number of resource blocks of type 'Expansion'
	// allowed for the composition request.
	MaxExpansion int
	// MaxMemory shall contain an integer that specifies the maximum number of resource blocks of type 'Memory' allowed
	// for the composition request.
	MaxMemory int
	// MaxNetwork shall contain an integer that specifies the maximum number of resource blocks of type 'Network'
	// allowed for the composition request.
	MaxNetwork int
	// MaxProcessor shall contain an integer that specifies the maximum number of resource blocks of type 'Processor'
	// allowed for the composition request.
	MaxProcessor int
	// MaxStorage shall contain an integer that specifies the maximum number of resource blocks of type 'Storage'
	// allowed for the composition request.
	MaxStorage int
	// MinCompute shall contain an integer that specifies the minimum number of resource blocks of type 'Compute'
	// required for the composition request.
	MinCompute int
	// MinComputerSystem shall contain an integer that specifies the minimum number of resource blocks of type
	// 'ComputerSystem' required for the composition request.
	MinComputerSystem int
	// MinExpansion shall contain an integer that specifies the minimum number of resource blocks of type 'Expansion'
	// required for the composition request.
	MinExpansion int
	// MinMemory shall contain an integer that specifies the minimum number of resource blocks of type 'Memory'
	// required for the composition request.
	MinMemory int
	// MinNetwork shall contain an integer that specifies the minimum number of resource blocks of type 'Network'
	// required for the composition request.
	MinNetwork int
	// MinProcessor shall contain an integer that specifies the minimum number of resource blocks of type 'Processor'
	// required for the composition request.
	MinProcessor int
	// MinStorage shall contain an integer that specifies the minimum number of resource blocks of type 'Storage'
	// required for the composition request.
	MinStorage int
}

// UnmarshalJSON unmarshals a ResourceBlockLimits object from the raw JSON.
func (resourceblocklimits *ResourceBlockLimits) UnmarshalJSON(b []byte) error {
	type temp ResourceBlockLimits
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*resourceblocklimits = ResourceBlockLimits(t.temp)

	// Extract the links to other entities for later

	return nil
}
