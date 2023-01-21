//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DeepOperations shall contain information about deep operations that the service supports.
type DeepOperations struct {
	// DeepPATCH shall indicate whether this service supports the Redfish Specification-defined deep PATCH operation.
	DeepPATCH string
	// DeepPOST shall indicate whether this service supports the Redfish Specification-defined deep POST operation.
	DeepPOST string
	// MaxLevels shall contain the maximum levels of resources allowed in deep operations.
	MaxLevels string
}

// UnmarshalJSON unmarshals a DeepOperations object from the raw JSON.
func (deepoperations *DeepOperations) UnmarshalJSON(b []byte) error {
	type temp DeepOperations
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*deepoperations = DeepOperations(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Expand shall contain information about the support of the $expand query parameter by the service.
type Expand struct {
	// ExpandAll shall indicate whether this service supports the asterisk ('*') option of the $expand query parameter.
	ExpandAll string
	// Levels shall indicate whether the service supports the $levels option of the $expand query parameter.
	Levels string
	// Links shall indicate whether this service supports the supports the tilde (~) option of the $expand query
	// parameter.
	Links string
	// MaxLevels shall contain the maximum $levels option value in the $expand query parameter. Shall be included only
	// if $levels is true.
	MaxLevels string
	// NoLinks shall indicate whether the service supports the period ('.') option of the $expand query parameter.
	NoLinks string
}

// UnmarshalJSON unmarshals a Expand object from the raw JSON.
func (expand *Expand) UnmarshalJSON(b []byte) error {
	type temp Expand
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*expand = Expand(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to Resources related to but not subordinate to this Resource.
type Links struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Sessions shall contain a link to a Resource Collection of type SessionCollection.
	Sessions string
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

// ProtocolFeaturesSupported shall contain information about protocol features that the service supports.
type ProtocolFeaturesSupported struct {
	// DeepOperations shall contain information about deep operations that the service supports.
	DeepOperations string
	// ExcerptQuery shall indicate whether this service supports the excerpt query parameter.
	ExcerptQuery string
	// ExpandQuery shall contain information about the support of the $expand query parameter by the service.
	ExpandQuery string
	// FilterQuery shall indicate whether this service supports the $filter query parameter.
	FilterQuery string
	// MultipleHTTPRequests shall indicate whether this service supports multiple outstanding HTTP requests.
	MultipleHTTPRequests string
	// OnlyMemberQuery shall indicate whether this service supports the only query parameter.
	OnlyMemberQuery string
	// SelectQuery shall indicate whether this service supports the $select query parameter.
	SelectQuery string
}

// UnmarshalJSON unmarshals a ProtocolFeaturesSupported object from the raw JSON.
func (protocolfeaturessupported *ProtocolFeaturesSupported) UnmarshalJSON(b []byte) error {
	type temp ProtocolFeaturesSupported
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*protocolfeaturessupported = ProtocolFeaturesSupported(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ServiceRoot shall comply with the Redfish Specification-described requirements.
type ServiceRoot struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccountService shall contain a link to a Resource of type AccountService.
	AccountService string
	// AggregationService shall contain a link to a resource of type AggregationService.
	AggregationService string
	// Cables shall contain a link to a resource collection of type CableCollection.
	Cables string
	// CertificateService shall contain a link to a Resource of type CertificateService.
	CertificateService string
	// Chassis shall contain a link to a Resource Collection of type ChassisCollection.
	Chassis string
	// ComponentIntegrity shall contain a link to a resource collection of type ComponentIntegrityCollection.
	ComponentIntegrity string
	// CompositionService shall contain a link to a Resource of type CompositionService.
	CompositionService string
	// Description provides a description of this resource.
	Description string
	// EventService shall contain a link to a Resource of type EventService.
	EventService string
	// Fabrics shall contain a link to a Resource Collection of type FabricCollection.
	Fabrics string
	// Facilities shall contain a link to a resource collection of type FacilityCollection.
	Facilities string
	// JobService shall contain a link to a Resource of type JobService.
	JobService string
	// JsonSchemas shall contain a link to a Resource Collection of type JsonSchemaFileCollection.
	JsonSchemas string
	// KeyService shall contain a link to a resource of type KeyService.
	KeyService string
	// LicenseService shall contain a link to a resource of type LicenseService.
	LicenseService string
	// Links shall contain links to Resources related to but not subordinate to this Resource.
	Links string
	// Managers shall contain a link to a Resource Collection of type ManagerCollection.
	Managers string
	// NVMeDomains shall contain a link to a resource collection of type NVMeDomainCollection.
	NVMeDomains string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerEquipment shall contain a link to a resource of type PowerEquipment.
	PowerEquipment string
	// Product shall include the name of the product represented by this Redfish Service.
	Product string
	// ProtocolFeaturesSupported shall contain information about protocol features that the service supports.
	ProtocolFeaturesSupported string
	// RedfishVersion shall represent the Redfish protocol version, as specified in the Protocol Version clause of the
	// Redfish Specification, to which this Service conforms.
	RedfishVersion string
	// RegisteredClients shall contain a link to a resource collection of type RegisteredClientCollection.
	RegisteredClients string
	// Registries shall contain a link to a Resource Collection of type MessageRegistryFileCollection.
	Registries string
	// ResourceBlocks shall contain a link to a Resource Collection of type ResourceBlockCollection.
	ResourceBlocks string
	// ServiceConditions shall contain a link to a resource of type ServiceConditions.
	ServiceConditions string
	// ServiceIdentification shall contain a vendor or user-provided value that identifies and associates a discovered
	// Redfish service with a particular product instance. The value of the property shall contain the value of the
	// 'ServiceIdentification' property in the Manager resource providing the Redfish service root resource. The value
	// of this property can be used during deployment processes to match user credentials or other a priori product
	// instance information to the appropriate Redfish service.
	ServiceIdentification string
	// SessionService shall contain a link to a Resource of type SessionService.
	SessionService string
	// Storage shall contain a link to a resource collection of type StorageCollection.
	Storage string
	// StorageServices shall contain a link to a Resource Collection of type StorageServiceCollection.
	StorageServices string
	// StorageSystems shall contain a link to a Resource Collection of type StorageSystemCollection. This collection
	// shall contain computer systems that act as storage servers. The HostingRoles attribute of each such computer
	// system shall have a StorageServer entry.
	StorageSystems string
	// Systems shall contain a link to a Resource Collection of type ComputerSystemCollection.
	Systems string
	// Tasks shall contain a link to a Resource of type TaskService.
	Tasks string
	// TelemetryService shall contain a link to a Resource of type TelemetryService.
	TelemetryService string
	// UUID shall represent the id of the Redfish Service instance. The format of this string shall contain a 32-byte
	// value in the form 8-4-4-4-12. If SSDP is used, this value shall be an exact match of the UUID value returned in
	// a 200 OK from an SSDP M-SEARCH request during discovery. RFC4122 describes methods to use to create a UUID
	// value. The value should be considered to be opaque. Client software should only treat the overall value as a
	// universally unique identifier and should not interpret any sub-fields within the UUID.
	UUID string
	// UpdateService shall contain a link to a Resource of type UpdateService.
	UpdateService string
	// Vendor shall include the name of the manufacturer or vendor represented by this Redfish Service. If this
	// property is supported, the vendor name shall not be included in the Product property value.
	Vendor string
}

// UnmarshalJSON unmarshals a ServiceRoot object from the raw JSON.
func (serviceroot *ServiceRoot) UnmarshalJSON(b []byte) error {
	type temp ServiceRoot
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*serviceroot = ServiceRoot(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetServiceRoot will get a ServiceRoot instance from the service.
func GetServiceRoot(c common.Client, uri string) (*ServiceRoot, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var serviceroot ServiceRoot
	err = json.NewDecoder(resp.Body).Decode(&serviceroot)
	if err != nil {
		return nil, err
	}

	serviceroot.SetClient(c)
	return &serviceroot, nil
}

// ListReferencedServiceRoots gets the collection of ServiceRoot from
// a provided reference.
func ListReferencedServiceRoots(c common.Client, link string) ([]*ServiceRoot, error) {
	var result []*ServiceRoot
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, servicerootLink := range links.ItemLinks {
		serviceroot, err := GetServiceRoot(c, servicerootLink)
		if err != nil {
			collectionError.Failures[servicerootLink] = err
		} else {
			result = append(result, serviceroot)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
