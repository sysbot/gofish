//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Mapping shall describe a mapping between a Resource type and the relevant privileges that accesses the Resource.
type Mapping struct {
	// Entity shall contain the Resource name, such as 'Manager'.
	Entity string
	// OperationMap shall list the mapping between HTTP methods and the privilege required for the Resource.
	OperationMap string
	// PropertyOverrides shall contain the privilege overrides of properties, such as the 'Password' property in the
	// 'ManagerAccount' Resource.
	PropertyOverrides []Target_PrivilegeMap
	// ResourceURIOverrides shall contain the privilege overrides of Resource URIs. The target lists the Resource URI
	// and the new privileges.
	ResourceURIOverrides []Target_PrivilegeMap
	// SubordinateOverrides shall contain the privilege overrides of the subordinate Resource. The target lists are
	// identified by Resource type.
	SubordinateOverrides []Target_PrivilegeMap
}

// UnmarshalJSON unmarshals a Mapping object from the raw JSON.
func (mapping *Mapping) UnmarshalJSON(b []byte) error {
	type temp Mapping
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*mapping = Mapping(t.temp)

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

// OperationMap shall describe the specific privileges required to complete a set of HTTP operations.
type OperationMap struct {
	// DELETE shall contain the privilege required to complete an HTTP DELETE operation.
	DELETE []OperationPrivilege
	// GET shall contain the privilege required to complete an HTTP GET operation.
	GET []OperationPrivilege
	// HEAD shall contain the privilege required to complete an HTTP HEAD operation.
	HEAD []OperationPrivilege
	// PATCH shall contain the privilege required to complete an HTTP PATCH operation.
	PATCH []OperationPrivilege
	// POST shall contain the privilege required to complete an HTTP POST operation.
	POST []OperationPrivilege
	// PUT shall contain the privilege required to complete an HTTP PUT operation.
	PUT []OperationPrivilege
}

// UnmarshalJSON unmarshals a OperationMap object from the raw JSON.
func (operationmap *OperationMap) UnmarshalJSON(b []byte) error {
	type temp OperationMap
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*operationmap = OperationMap(t.temp)

	// Extract the links to other entities for later

	return nil
}

// OperationPrivilege shall describe the privileges required to complete a specific HTTP operation.
type OperationPrivilege struct {
	// Privilege shall contain an array of privileges that are required to complete a specific HTTP operation on a
	// Resource. This set of strings match zero or more strings in the PrivilegesUsed and OEMPrivilegesUsed properties.
	Privilege []string
}

// UnmarshalJSON unmarshals a OperationPrivilege object from the raw JSON.
func (operationprivilege *OperationPrivilege) UnmarshalJSON(b []byte) error {
	type temp OperationPrivilege
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*operationprivilege = OperationPrivilege(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PrivilegeRegistry This Resource contains operation-to-privilege mappings.
type PrivilegeRegistry struct {
	common.Entity
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// Mappings shall describe the mappings between entities and the relevant privileges that access those entities.
	Mappings []Mapping
	// OEMPrivilegesUsed shall contain an array of OEM privileges used in this mapping.
	OEMPrivilegesUsed []string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PrivilegesUsed shall contain an array of Redfish standard privileges used in this mapping.
	PrivilegesUsed []PrivilegeType
}

// UnmarshalJSON unmarshals a PrivilegeRegistry object from the raw JSON.
func (privilegeregistry *PrivilegeRegistry) UnmarshalJSON(b []byte) error {
	type temp PrivilegeRegistry
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*privilegeregistry = PrivilegeRegistry(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetPrivilegeRegistry will get a PrivilegeRegistry instance from the service.
func GetPrivilegeRegistry(c common.Client, uri string) (*PrivilegeRegistry, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var privilegeregistry PrivilegeRegistry
	err = json.NewDecoder(resp.Body).Decode(&privilegeregistry)
	if err != nil {
		return nil, err
	}

	privilegeregistry.SetClient(c)
	return &privilegeregistry, nil
}

// ListReferencedPrivilegeRegistrys gets the collection of PrivilegeRegistry from
// a provided reference.
func ListReferencedPrivilegeRegistrys(c common.Client, link string) ([]*PrivilegeRegistry, error) {
	var result []*PrivilegeRegistry
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, privilegeregistryLink := range links.ItemLinks {
		privilegeregistry, err := GetPrivilegeRegistry(c, privilegeregistryLink)
		if err != nil {
			collectionError.Failures[privilegeregistryLink] = err
		} else {
			result = append(result, privilegeregistry)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// TargetPrivilegeMap shall describe a mapping between one or more targets and the HTTP operations associated with
// them.
type TargetPrivilegeMap struct {
	// OperationMap shall contain the mapping between the HTTP operation and the privilege required to complete the
	// operation.
	OperationMap string
	// Targets shall contain the array of URIs, Resource types, or properties. For example, '/redfish/v1/Systems/1',
	// 'Manager', or 'Password'. When the Targets property is not present, no override is specified.
	Targets []string
}

// UnmarshalJSON unmarshals a TargetPrivilegeMap object from the raw JSON.
func (targetprivilegemap *TargetPrivilegeMap) UnmarshalJSON(b []byte) error {
	type temp TargetPrivilegeMap
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*targetprivilegemap = TargetPrivilegeMap(t.temp)

	// Extract the links to other entities for later

	return nil
}
