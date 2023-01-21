//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Location shall contain the location information for a Message Registry file.
type Location struct {
	// ArchiveFile shall contain the file name of the individual Message Registry file within the archive file
	// specified by the ArchiveUri property. The file name shall conform to the Redfish Specification-specified syntax.
	ArchiveFile string
	// ArchiveUri shall contain a URI that is colocated with the Redfish Service that specifies the location of the
	// Message Registry file, which can be retrieved using the Redfish protocol and authentication methods. This
	// property shall be used for only ZIP or other archive files. The ArchiveFile property shall contain the file name
	// of the individual Message Registry file within the archive file.
	ArchiveUri string
	// Language shall contain an RFC5646-conformant language code or 'default'.
	Language string
	// PublicationUri shall contain a URI not colocated with the Redfish Service that specifies the canonical location
	// of the Message Registry file. This property shall be used for only individual Message Registry files.
	PublicationUri string
	// Uri shall contain a URI colocated with the Redfish Service that specifies the location of the Message Registry
	// file, which can be retrieved using the Redfish protocol and authentication methods. This property shall be used
	// for only individual Message Registry files. The file name portion of the URI shall conform to Redfish
	// Specification-specified syntax.
	Uri string
}

// UnmarshalJSON unmarshals a Location object from the raw JSON.
func (location *Location) UnmarshalJSON(b []byte) error {
	type temp Location
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*location = Location(t.temp)

	// Extract the links to other entities for later

	return nil
}

// MessageRegistryFile shall represent the Message Registry file locator for a Redfish implementation.
type MessageRegistryFile struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// Languages This property contains a set of RFC5646-conformant language codes.
	Languages []string
	// Location shall contain the location information for this Message Registry file.
	Location []Location
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Registry shall contain the Message Registry name and it major and minor versions, as defined by the Redfish
	// Specification. This registry can be any type of registry, such as Message Registry, Privilege Registry, or
	// Attribute Registry.
	Registry string
}

// UnmarshalJSON unmarshals a MessageRegistryFile object from the raw JSON.
func (messageregistryfile *MessageRegistryFile) UnmarshalJSON(b []byte) error {
	type temp MessageRegistryFile
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*messageregistryfile = MessageRegistryFile(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetMessageRegistryFile will get a MessageRegistryFile instance from the service.
func GetMessageRegistryFile(c common.Client, uri string) (*MessageRegistryFile, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var messageregistryfile MessageRegistryFile
	err = json.NewDecoder(resp.Body).Decode(&messageregistryfile)
	if err != nil {
		return nil, err
	}

	messageregistryfile.SetClient(c)
	return &messageregistryfile, nil
}

// ListReferencedMessageRegistryFiles gets the collection of MessageRegistryFile from
// a provided reference.
func ListReferencedMessageRegistryFiles(c common.Client, link string) ([]*MessageRegistryFile, error) {
	var result []*MessageRegistryFile
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, messageregistryfileLink := range links.ItemLinks {
		messageregistryfile, err := GetMessageRegistryFile(c, messageregistryfileLink)
		if err != nil {
			collectionError.Failures[messageregistryfileLink] = err
		} else {
			result = append(result, messageregistryfile)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
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
