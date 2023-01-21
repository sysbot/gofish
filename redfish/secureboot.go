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
	// ResetAllKeysToDefaultResetKeysType Reset the contents of all UEFI Secure Boot key databases, including the PK
	// key database, to the default values.
	ResetAllKeysToDefaultResetKeysType ResetKeysType = "ResetAllKeysToDefault"
	// DeleteAllKeysResetKeysType Delete the contents of all UEFI Secure Boot key databases, including the PK key
	// database. This puts the system in Setup Mode.
	DeleteAllKeysResetKeysType ResetKeysType = "DeleteAllKeys"
	// DeletePKResetKeysType Delete the contents of the PK UEFI Secure Boot database. This puts the system in Setup
	// Mode.
	DeletePKResetKeysType ResetKeysType = "DeletePK"
)

// SecureBootCurrentBootType is
type SecureBootCurrentBootType string

const (
	// EnabledSecureBootCurrentBootType UEFI Secure Boot is currently enabled.
	EnabledSecureBootCurrentBootType SecureBootCurrentBootType = "Enabled"
	// DisabledSecureBootCurrentBootType UEFI Secure Boot is currently disabled.
	DisabledSecureBootCurrentBootType SecureBootCurrentBootType = "Disabled"
)

// SecureBootModeType is
type SecureBootModeType string

const (
	// SetupModeSecureBootModeType UEFI Secure Boot is currently in Setup Mode.
	SetupModeSecureBootModeType SecureBootModeType = "SetupMode"
	// UserModeSecureBootModeType UEFI Secure Boot is currently in User Mode.
	UserModeSecureBootModeType SecureBootModeType = "UserMode"
	// AuditModeSecureBootModeType UEFI Secure Boot is currently in Audit Mode.
	AuditModeSecureBootModeType SecureBootModeType = "AuditMode"
	// DeployedModeSecureBootModeType UEFI Secure Boot is currently in Deployed Mode.
	DeployedModeSecureBootModeType SecureBootModeType = "DeployedMode"
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

// SecureBoot This resource contains UEFI Secure Boot information for a Redfish implementation.
type SecureBoot struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SecureBootCurrentBoot shall indicate the UEFI Secure Boot state during the current boot cycle.
	SecureBootCurrentBoot SecureBootCurrentBootType
	// SecureBootDatabases shall be a link to a resource collection of type SecureBootDatabaseCollection.
	SecureBootDatabases string
	// SecureBootEnable shall indicate whether the UEFI Secure Boot takes effect on next boot. This property can be
	// enabled in UEFI boot mode only.
	SecureBootEnable bool
	// SecureBootMode shall contain the current UEFI Secure Boot mode, as defined in the UEFI Specification.
	SecureBootMode SecureBootModeType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SecureBoot object from the raw JSON.
func (secureboot *SecureBoot) UnmarshalJSON(b []byte) error {
	type temp SecureBoot
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*secureboot = SecureBoot(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	secureboot.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (secureboot *SecureBoot) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(SecureBoot)
	original.UnmarshalJSON(secureboot.rawData)

	readWriteFields := []string{
		"SecureBootEnable",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(secureboot).Elem()

	return secureboot.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSecureBoot will get a SecureBoot instance from the service.
func GetSecureBoot(c common.Client, uri string) (*SecureBoot, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var secureboot SecureBoot
	err = json.NewDecoder(resp.Body).Decode(&secureboot)
	if err != nil {
		return nil, err
	}

	secureboot.SetClient(c)
	return &secureboot, nil
}

// ListReferencedSecureBoots gets the collection of SecureBoot from
// a provided reference.
func ListReferencedSecureBoots(c common.Client, link string) ([]*SecureBoot, error) {
	var result []*SecureBoot
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, securebootLink := range links.ItemLinks {
		secureboot, err := GetSecureBoot(c, securebootLink)
		if err != nil {
			collectionError.Failures[securebootLink] = err
		} else {
			result = append(result, secureboot)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
