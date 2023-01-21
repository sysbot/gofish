//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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

// Role This resource represents the Redfish role for the user account.
type Role struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AlternateRoleId shall contain a non-restricted 'RoleId' intended to be used in its place when the Restricted
	// property contains the value 'true'.
	AlternateRoleId string
	// AssignedPrivileges shall contain the Redfish privileges for this role. For predefined roles, this property shall
	// be read-only. For custom roles, some implementations may prevent writing to this property.
	AssignedPrivileges []PrivilegeType
	// Description provides a description of this resource.
	Description string
	// IsPredefined shall indicate whether the role is predefined by Redfish or an OEM as contrasted with a client-
	// defined role.
	IsPredefined string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OemPrivileges shall contain the OEM privileges for this role. For predefined roles, this property shall be read-
	// only. For custom roles, some implementations may prevent writing to this property.
	OemPrivileges []string
	// Restricted shall indicate whether use of the role is restricted by a service as defined by the 'Restricted roles
	// and restricted privileges' clause of the Redfish Specification. If this property is not present, the value shall
	// be assumed to be 'false'.
	Restricted string
	// RoleId shall contain the string name of the role. This property shall contain the same value as the Id property.
	RoleId string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Role object from the raw JSON.
func (role *Role) UnmarshalJSON(b []byte) error {
	type temp Role
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*role = Role(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	role.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (role *Role) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Role)
	original.UnmarshalJSON(role.rawData)

	readWriteFields := []string{
		"AssignedPrivileges",
		"OemPrivileges",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(role).Elem()

	return role.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetRole will get a Role instance from the service.
func GetRole(c common.Client, uri string) (*Role, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var role Role
	err = json.NewDecoder(resp.Body).Decode(&role)
	if err != nil {
		return nil, err
	}

	role.SetClient(c)
	return &role, nil
}

// ListReferencedRoles gets the collection of Role from
// a provided reference.
func ListReferencedRoles(c common.Client, link string) ([]*Role, error) {
	var result []*Role
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, roleLink := range links.ItemLinks {
		role, err := GetRole(c, roleLink)
		if err != nil {
			collectionError.Failures[roleLink] = err
		} else {
			result = append(result, role)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
