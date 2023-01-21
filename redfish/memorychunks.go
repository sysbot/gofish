//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AddressRangeType is
type AddressRangeType string

const (
	// VolatileAddressRangeType Volatile memory.
	VolatileAddressRangeType AddressRangeType = "Volatile"
	// PMEMAddressRangeType Byte accessible persistent memory.
	PMEMAddressRangeType AddressRangeType = "PMEM"
	// BlockAddressRangeType Block accessible memory.
	BlockAddressRangeType AddressRangeType = "Block"
)

// InterleaveSet shall describe an interleave set of which the memory chunk is a part.
type InterleaveSet struct {
	// Memory shall contain the memory device to which these settings apply.
	Memory string
	// MemoryLevel shall contain the level of this interleave set for multi-level tiered memory.
	MemoryLevel int
	// OffsetMiB shall contain the offset within the DIMM that corresponds to the start of this memory region, with
	// units in MiB.
	OffsetMiB int
	// RegionId shall contain the DIMM region identifier.
	RegionId string
	// SizeMiB shall contain the size of this memory region, with units in MiB.
	SizeMiB int
}

// UnmarshalJSON unmarshals a InterleaveSet object from the raw JSON.
func (interleaveset *InterleaveSet) UnmarshalJSON(b []byte) error {
	type temp InterleaveSet
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*interleaveset = InterleaveSet(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Endpoints shall contain a link to the resources of type Endpoint with which this memory chunk is associated.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
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

// MemoryChunks shall represent memory chunks and interleave sets in a Redfish implementation.
type MemoryChunks struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AddressRangeOffsetMiB shall be the offset of the memory chunk in the address range in MiB.
	AddressRangeOffsetMiB int
	// AddressRangeType shall contain the type of memory chunk.
	AddressRangeType AddressRangeType
	// Description provides a description of this resource.
	Description string
	// DisplayName shall contain a user-configurable string to name the memory chunk.
	DisplayName string
	// InterleaveSets shall represent the interleave sets for the memory chunk. If not specified by the client during a
	// create operation, the memory chunk shall be created across all available memory within the memory domain.
	InterleaveSets []InterleaveSet
	// IsMirrorEnabled shall indicate whether memory mirroring is enabled for this memory chunk.
	IsMirrorEnabled bool
	// IsSpare shall indicate whether sparing is enabled for this memory chunk.
	IsSpare bool
	// Links shall contain links to resources that are related to but are not contained by or subordinate to this
	// resource.
	Links string
	// MemoryChunkSizeMiB shall contain the size of the memory chunk in MiB.
	MemoryChunkSizeMiB int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a MemoryChunks object from the raw JSON.
func (memorychunks *MemoryChunks) UnmarshalJSON(b []byte) error {
	type temp MemoryChunks
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorychunks = MemoryChunks(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	memorychunks.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (memorychunks *MemoryChunks) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(MemoryChunks)
	original.UnmarshalJSON(memorychunks.rawData)

	readWriteFields := []string{
		"DisplayName",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(memorychunks).Elem()

	return memorychunks.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMemoryChunks will get a MemoryChunks instance from the service.
func GetMemoryChunks(c common.Client, uri string) (*MemoryChunks, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memorychunks MemoryChunks
	err = json.NewDecoder(resp.Body).Decode(&memorychunks)
	if err != nil {
		return nil, err
	}

	memorychunks.SetClient(c)
	return &memorychunks, nil
}

// ListReferencedMemoryChunkss gets the collection of MemoryChunks from
// a provided reference.
func ListReferencedMemoryChunkss(c common.Client, link string) ([]*MemoryChunks, error) {
	var result []*MemoryChunks
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, memorychunksLink := range links.ItemLinks {
		memorychunks, err := GetMemoryChunks(c, memorychunksLink)
		if err != nil {
			collectionError.Failures[memorychunksLink] = err
		} else {
			result = append(result, memorychunks)
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
