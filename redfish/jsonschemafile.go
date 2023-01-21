//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// JsonSchemaFile shall represent the schema file locator Resource for a Redfish implementation.
type JsonSchemaFile struct {
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
	// Location shall contain the location information for this schema file.
	Location []Location
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Schema shall contain the @odata.type property value for that schema and shall conform to the Redfish
	// Specification-specified syntax for the Type property.
	Schema string
}

// UnmarshalJSON unmarshals a JsonSchemaFile object from the raw JSON.
func (jsonschemafile *JsonSchemaFile) UnmarshalJSON(b []byte) error {
	type temp JsonSchemaFile
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*jsonschemafile = JsonSchemaFile(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetJsonSchemaFile will get a JsonSchemaFile instance from the service.
func GetJsonSchemaFile(c common.Client, uri string) (*JsonSchemaFile, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jsonschemafile JsonSchemaFile
	err = json.NewDecoder(resp.Body).Decode(&jsonschemafile)
	if err != nil {
		return nil, err
	}

	jsonschemafile.SetClient(c)
	return &jsonschemafile, nil
}

// ListReferencedJsonSchemaFiles gets the collection of JsonSchemaFile from
// a provided reference.
func ListReferencedJsonSchemaFiles(c common.Client, link string) ([]*JsonSchemaFile, error) {
	var result []*JsonSchemaFile
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, jsonschemafileLink := range links.ItemLinks {
		jsonschemafile, err := GetJsonSchemaFile(c, jsonschemafileLink)
		if err != nil {
			collectionError.Failures[jsonschemafileLink] = err
		} else {
			result = append(result, jsonschemafile)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Location shall describe location information for a schema file.
type Location struct {
	// ArchiveFile shall contain the file name of the individual schema file within the archive file that the
	// ArchiveUri property specifies. The file name shall conform to the Redfish Specification-described format.
	ArchiveFile string
	// ArchiveUri shall contain a URI colocated with the Redfish Service that specifies the location of the schema
	// file, which can be retrieved using the Redfish protocol and authentication methods. This property shall be used
	// for only archive files, in zip or other formats. The ArchiveFile value shall be the individual schema file name
	// within the archive file.
	ArchiveUri string
	// Language shall contain an RFC5646-conformant language code or the 'default' string.
	Language string
	// PublicationUri shall contain a URI not colocated with the Redfish Service that specifies the canonical location
	// of the schema file. This property shall be used for only individual schema files.
	PublicationUri string
	// Uri shall contain a URI colocated with the Redfish Service that specifies the location of the schema file, which
	// can be retrieved using the Redfish protocol and authentication methods. This property shall be used for only
	// individual schema files. The file name portion of the URI shall conform to the format specified in the Redfish
	// Specification.
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
