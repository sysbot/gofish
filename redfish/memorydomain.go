//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Links shall contain links to Resources related to but not subordinate to this Resource.
type Links struct {
	// FabricAdapters shall contain an array of links to resources of type FabricAdapter with which this memory domain
	// is associated.
	FabricAdapters []FabricAdapter
	// FabricAdapters@odata.count
	FabricAdaptersCount int `json:"FabricAdapters@odata.count"`
	// MediaControllers@odata.count
	MediaControllersCount int `json:"MediaControllers@odata.count"`
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

// MemoryDomain shall represent memory domains in a Redfish implementation.
type MemoryDomain struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// AllowsBlockProvisioning shall indicate whether this memory domain supports the creation of blocks of memory.
	AllowsBlockProvisioning bool
	// AllowsMemoryChunkCreation shall indicate whether this memory domain supports the creation of memory chunks.
	AllowsMemoryChunkCreation bool
	// AllowsMirroring shall indicate whether this memory domain supports the creation of memory chunks with mirroring
	// enabled.
	AllowsMirroring bool
	// AllowsSparing shall indicate whether this memory domain supports the creation of memory chunks with sparing
	// enabled.
	AllowsSparing bool
	// Description provides a description of this resource.
	Description string
	// InterleavableMemorySets shall represent the interleave sets for the memory chunk.
	InterleavableMemorySets []MemorySet
	// Links shall contain links to Resources related to but not subordinate to this Resource.
	Links string
	// MemoryChunks shall contain a link to a Resource Collection of type MemoryChunkCollection.
	MemoryChunks string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a MemoryDomain object from the raw JSON.
func (memorydomain *MemoryDomain) UnmarshalJSON(b []byte) error {
	type temp MemoryDomain
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorydomain = MemoryDomain(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetMemoryDomain will get a MemoryDomain instance from the service.
func GetMemoryDomain(c common.Client, uri string) (*MemoryDomain, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memorydomain MemoryDomain
	err = json.NewDecoder(resp.Body).Decode(&memorydomain)
	if err != nil {
		return nil, err
	}

	memorydomain.SetClient(c)
	return &memorydomain, nil
}

// ListReferencedMemoryDomains gets the collection of MemoryDomain from
// a provided reference.
func ListReferencedMemoryDomains(c common.Client, link string) ([]*MemoryDomain, error) {
	var result []*MemoryDomain
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, memorydomainLink := range links.ItemLinks {
		memorydomain, err := GetMemoryDomain(c, memorydomainLink)
		if err != nil {
			collectionError.Failures[memorydomainLink] = err
		} else {
			result = append(result, memorydomain)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// MemorySet shall represent the interleave sets for a memory chunk.
type MemorySet struct {
	// MemorySet shall be links to Resources of the Memory type.
	MemorySet []Memory
	// MemorySet@odata.count
	MemorySetCount int `json:"MemorySet@odata.count"`
}

// UnmarshalJSON unmarshals a MemorySet object from the raw JSON.
func (memoryset *MemorySet) UnmarshalJSON(b []byte) error {
	type temp MemorySet
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memoryset = MemorySet(t.temp)

	// Extract the links to other entities for later

	return nil
}

// OemActions shall contain the available OEM-specific actions for this Resource.
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
