//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Expand is
type Expand string

const (
	// NoneExpand shall indicate that references in the manifest response will not be expanded.
	NoneExpand Expand = "None"
	// AllExpand shall indicate that all subordinate references in the manifest response will be expanded.
	AllExpand Expand = "All"
	// RelevantExpand shall indicate that relevant subordinate references in the manifest response will be expanded.
	RelevantExpand Expand = "Relevant"
)

// StanzaType is
type StanzaType string

const (
	// ComposeSystemStanzaType shall indicate a stanza that describes the specific, constrained, or mixed resources
	// required to compose a computer system. The resource blocks assigned to the computer system shall be moved to the
	// active pool. The Request property of the stanza shall contain a resource of type ComputerSystem that represents
	// the composition request. The Response property of the stanza shall contain a resource of type ComputerSystem
	// that represents the composed system or a Redfish Specification-defined error response.
	ComposeSystemStanzaType StanzaType = "ComposeSystem"
	// DecomposeSystemStanzaType shall indicate a stanza that references a computer system to decompose and return the
	// resource blocks to the free pool that are no longer contributing to composed resources. The Request property of
	// the stanza shall be a Redfish Specification-defined reference object containing a reference to the resource of
	// type ComputerSystem to decompose. The Response property of the stanza shall contain a resource of type
	// ComputerSystem that represents the decomposed system or a Redfish Specification-defined error response.
	DecomposeSystemStanzaType StanzaType = "DecomposeSystem"
	// ComposeResourceStanzaType shall indicate a stanza that describes a composed resource block. The resource blocks
	// assigned to the composed resource block shall be moved to the active pool. The Request property of the stanza
	// shall contain a resource of type ResourceBlock that represents the composition request. The Response property of
	// the stanza shall contain a resource of type ResourceBlock that represents the composed resource block or a
	// Redfish Specification-defined error response.
	ComposeResourceStanzaType StanzaType = "ComposeResource"
	// DecomposeResourceStanzaType shall indicate a stanza that references a composed resource block to decompose and
	// return the resource blocks to the free pool that are no longer contributing to composed resources. The Request
	// property of the stanza shall be a reference object as defined by the 'Reference properties' clause of the
	// Redfish Specification containing a reference to the resource of type ResourceBlock to decompose. The Response
	// property of the stanza shall contain a resource of type ResourceBlock that represents the decomposed resource
	// block or a Redfish Specification-defined error response.
	DecomposeResourceStanzaType StanzaType = "DecomposeResource"
	// OEMStanzaType shall indicate a stanza that describes an OEM-specific request. The OEMStanzaType property shall
	// contain the specific OEM stanza type.
	OEMStanzaType StanzaType = "OEM"
	// RegisterResourceBlockStanzaType shall indicate a stanza that references a resource to create a resource block
	// that references the resource and add it to the free pool. The Request property of the stanza shall contain a
	// resource of type ResourceBlock that represents the registration request. The Response property of the stanza
	// shall contain a resource of type ResourceBlock that represents the composed system or a Redfish Specification-
	// defined error response.
	RegisterResourceBlockStanzaType StanzaType = "RegisterResourceBlock"
)

// Manifest shall describe a manifest containing a set of requests to be fulfilled.
type Manifest struct {
	// Description provides a description of this resource.
	Description string
	// Expand shall contain the expansion control for references in manifest responses.
	Expand Expand
	// Stanzas shall contain an array of stanzas that describe the requests specified by this manifest.
	Stanzas []Stanza
	// Timestamp shall contain the date and time when the manifest was created.
	Timestamp string
}

// UnmarshalJSON unmarshals a Manifest object from the raw JSON.
func (manifest *Manifest) UnmarshalJSON(b []byte) error {
	type temp Manifest
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*manifest = Manifest(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetManifest will get a Manifest instance from the service.
func GetManifest(c common.Client, uri string) (*Manifest, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var manifest Manifest
	err = json.NewDecoder(resp.Body).Decode(&manifest)
	if err != nil {
		return nil, err
	}

	manifest.SetClient(c)
	return &manifest, nil
}

// ListReferencedManifests gets the collection of Manifest from
// a provided reference.
func ListReferencedManifests(c common.Client, link string) ([]*Manifest, error) {
	var result []*Manifest
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, manifestLink := range links.ItemLinks {
		manifest, err := GetManifest(c, manifestLink)
		if err != nil {
			collectionError.Failures[manifestLink] = err
		} else {
			result = append(result, manifest)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Request shall describe the request details of a stanza within a manifest. Its contents vary depending on the
// value of the StanzaType property of the stanza.
type Request struct {
}

// UnmarshalJSON unmarshals a Request object from the raw JSON.
func (request *Request) UnmarshalJSON(b []byte) error {
	type temp Request
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*request = Request(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Response shall describe the response details of a stanza within a manifest. Its contents vary depending on the
// value of the StanzaType property of the stanza.
type Response struct {
}

// UnmarshalJSON unmarshals a Response object from the raw JSON.
func (response *Response) UnmarshalJSON(b []byte) error {
	type temp Response
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*response = Response(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Stanza shall contain properties that describe a request to be fulfilled within a manifest.
type Stanza struct {
	// OEMStanzaType shall contain the OEM-defined type of stanza. This property shall be present if StanzaType is
	// 'OEM'.
	OEMStanzaType string
	// Request shall contain the request details for the stanza and the contents vary based depending on the value of
	// the StanzaType property.
	Request Request
	// Response shall contain the response details for the stanza and the contents vary based depending on the value of
	// the StanzaType property.
	Response Response
	// StanzaId shall contain the identifier of the stanza.
	StanzaId string
	// StanzaType shall contain the type of stanza.
	StanzaType StanzaType
}

// UnmarshalJSON unmarshals a Stanza object from the raw JSON.
func (stanza *Stanza) UnmarshalJSON(b []byte) error {
	type temp Stanza
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*stanza = Stanza(t.temp)

	// Extract the links to other entities for later

	return nil
}
