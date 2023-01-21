//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ResetKeysType is
type ResetKeysType string

const (
	// ResetAllKeysToDefaultResetKeysType Reset the content of this UEFI Secure Boot key database to the default
	// values.
	ResetAllKeysToDefaultResetKeysType ResetKeysType = "ResetAllKeysToDefault"
	// DeleteAllKeysResetKeysType Delete the content of this UEFI Secure Boot key database.
	DeleteAllKeysResetKeysType ResetKeysType = "DeleteAllKeys"
)

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

// SecureBootDatabase shall be used to represent a UEFI Secure Boot database for a Redfish implementation.
type SecureBootDatabase struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Certificates shall be a link to a resource collection of type CertificateCollection.
	Certificates string
	// DatabaseId shall contain the name of the UEFI Secure Boot database. This property shall contain the same value
	// as the Id property. The value shall be one of the UEFI-defined Secure Boot databases: 'PK', 'KEK' 'db', 'dbx',
	// 'dbr', 'dbt', 'PKDefault', 'KEKDefault', 'dbDefault', 'dbxDefault', 'dbrDefault', or 'dbtDefault'.
	DatabaseId string
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Signatures shall be a link to a resource collection of type SignatureCollection.
	Signatures string
}

// UnmarshalJSON unmarshals a SecureBootDatabase object from the raw JSON.
func (securebootdatabase *SecureBootDatabase) UnmarshalJSON(b []byte) error {
	type temp SecureBootDatabase
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*securebootdatabase = SecureBootDatabase(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetSecureBootDatabase will get a SecureBootDatabase instance from the service.
func GetSecureBootDatabase(c common.Client, uri string) (*SecureBootDatabase, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var securebootdatabase SecureBootDatabase
	err = json.NewDecoder(resp.Body).Decode(&securebootdatabase)
	if err != nil {
		return nil, err
	}

	securebootdatabase.SetClient(c)
	return &securebootdatabase, nil
}

// ListReferencedSecureBootDatabases gets the collection of SecureBootDatabase from
// a provided reference.
func ListReferencedSecureBootDatabases(c common.Client, link string) ([]*SecureBootDatabase, error) {
	var result []*SecureBootDatabase
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, securebootdatabaseLink := range links.ItemLinks {
		securebootdatabase, err := GetSecureBootDatabase(c, securebootdatabaseLink)
		if err != nil {
			collectionError.Failures[securebootdatabaseLink] = err
		} else {
			result = append(result, securebootdatabase)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
