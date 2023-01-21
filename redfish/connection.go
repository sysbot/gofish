//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AccessCapability is
type AccessCapability string

const (
	// ReadAccessCapability Endpoints are allowed to perform reads from the specified resource.
	ReadAccessCapability AccessCapability = "Read"
	// WriteAccessCapability Endpoints are allowed to perform writes to the specified resource.
	WriteAccessCapability AccessCapability = "Write"
)

// AccessState is This type shall describe the access to the associated resource in this connection.
type AccessState string

const (
	// OptimizedAccessState shall indicate the resource is in an active and optimized state.
	OptimizedAccessState AccessState = "Optimized"
	// NonOptimizedAccessState shall indicate the resource is in an active and non-optimized state.
	NonOptimizedAccessState AccessState = "NonOptimized"
	// StandbyAccessState shall indicate the resource is in a standby state.
	StandbyAccessState AccessState = "Standby"
	// UnavailableAccessState shall indicate the resource is in an unavailable state.
	UnavailableAccessState AccessState = "Unavailable"
	// TransitioningAccessState shall indicate the resource is transitioning to a new state.
	TransitioningAccessState AccessState = "Transitioning"
)

// ConnectionType is
type ConnectionType string

const (
	// StorageConnectionType A connection to storage related resources, such as volumes.
	StorageConnectionType ConnectionType = "Storage"
	// MemoryConnectionType A connection to memory related resources.
	MemoryConnectionType ConnectionType = "Memory"
)

// Connection shall represent a connection information in the Redfish Specification.
type Connection struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ConnectionKeys shall contain the permissions keys required to access the specified resources for this
	// connection. Some fabrics require permission checks on transactions from authorized initiators.
	ConnectionKeys string
	// ConnectionType shall contain the type of resources this connection specifies.
	ConnectionType ConnectionType
	// Description provides a description of this resource.
	Description string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// MemoryChunkInfo shall contain the set of memory chunks and access capabilities specified for this connection.
	MemoryChunkInfo []MemoryChunkInfo
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// VolumeInfo shall contain the set of volumes and access capabilities specified for this connection.
	VolumeInfo []VolumeInfo
}

// UnmarshalJSON unmarshals a Connection object from the raw JSON.
func (connection *Connection) UnmarshalJSON(b []byte) error {
	type temp Connection
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*connection = Connection(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetConnection will get a Connection instance from the service.
func GetConnection(c common.Client, uri string) (*Connection, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var connection Connection
	err = json.NewDecoder(resp.Body).Decode(&connection)
	if err != nil {
		return nil, err
	}

	connection.SetClient(c)
	return &connection, nil
}

// ListReferencedConnections gets the collection of Connection from
// a provided reference.
func ListReferencedConnections(c common.Client, link string) ([]*Connection, error) {
	var result []*Connection
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, connectionLink := range links.ItemLinks {
		connection, err := GetConnection(c, connectionLink)
		if err != nil {
			collectionError.Failures[connectionLink] = err
		} else {
			result = append(result, connection)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// ConnectionKey shall contain the permission key information required to access the target resources for a
// connection.
type ConnectionKey struct {
	// GenZ shall contain the Gen-Z-specific permission key information for this connection.
	GenZ GenZConnectionKey
}

// UnmarshalJSON unmarshals a ConnectionKey object from the raw JSON.
func (connectionkey *ConnectionKey) UnmarshalJSON(b []byte) error {
	type temp ConnectionKey
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*connectionkey = ConnectionKey(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GenZConnectionKey shall contain the Gen-Z-specific permission key information for a connection.
type GenZConnectionKey struct {
	// AccessKey shall contain the Gen-Z Core Specification-defined Access Key for this connection.
	AccessKey string
	// RKeyDomainCheckingEnabled shall indicate whether Region Key domain checking is enabled for this connection.
	RKeyDomainCheckingEnabled string
	// RKeyReadOnlyKey shall contain the Gen-Z Core Specification-defined read-only Region Key for this connection.
	RKeyReadOnlyKey string
	// RKeyReadWriteKey shall contain the Gen-Z Core Specification-defined read-write Region Key for this connection.
	RKeyReadWriteKey string
}

// UnmarshalJSON unmarshals a GenZConnectionKey object from the raw JSON.
func (genzconnectionkey *GenZConnectionKey) UnmarshalJSON(b []byte) error {
	type temp GenZConnectionKey
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*genzconnectionkey = GenZConnectionKey(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// InitiatorEndpointGroups shall contain an array of links to resources of type EndpointGroup that are the
	// initiator endpoint groups associated with this connection. If the referenced endpoint groups contain the
	// GroupType property, the GroupType property shall contain the value 'Initiator' or 'Client'. This property shall
	// not be present if InitiatorEndpoints is present.
	InitiatorEndpointGroups []EndpointGroup
	// InitiatorEndpointGroups@odata.count
	InitiatorEndpointGroupsCount int `json:"InitiatorEndpointGroups@odata.count"`
	// InitiatorEndpoints shall contain an array of links to resources of type Endpoint that are the initiator
	// endpoints associated with this connection. If the referenced endpoints contain the EntityRole property, the
	// EntityRole property shall contain the value 'Initiator' or 'Both'. This property shall not be present if
	// InitiatorEndpointGroups is present.
	InitiatorEndpoints []Endpoint
	// InitiatorEndpoints@odata.count
	InitiatorEndpointsCount int `json:"InitiatorEndpoints@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// TargetEndpointGroups shall contain an array of links to resources of type EndpointGroup that are the target
	// endpoint groups associated with this connection. If the referenced endpoint groups contain the GroupType
	// property, the GroupType property shall contain the value 'Target' or 'Server'. This property shall not be
	// present if TargetEndpoints is present.
	TargetEndpointGroups []EndpointGroup
	// TargetEndpointGroups@odata.count
	TargetEndpointGroupsCount int `json:"TargetEndpointGroups@odata.count"`
	// TargetEndpoints shall contain an array of links to resources of type Endpoint that are the target endpoints
	// associated with this connection. If the referenced endpoints contain the EntityRole property, the EntityRole
	// property shall contain the value 'Target' or 'Both'. This property shall not be present if TargetEndpointGroups
	// is present.
	TargetEndpoints []Endpoint
	// TargetEndpoints@odata.count
	TargetEndpointsCount int `json:"TargetEndpoints@odata.count"`
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

// MemoryChunkInfo shall contain the combination of permissions and memory chunk information.
type MemoryChunkInfo struct {
	// AccessCapabilities shall specify a current memory access capability.
	AccessCapabilities []AccessCapability
	// AccessState shall contain the access state for the associated resource in this connection.
	AccessState AccessState
	// MemoryChunk shall contain a link to a resource of type MemoryChunk. The endpoints referenced by the
	// InitiatorEndpoints or InitiatorEndpointGroups properties shall be given access to this memory chunk as described
	// by this object. If TargetEndpoints or TargetEndpointGroups is present, the referenced initiator endpoints shall
	// be required to access the referenced memory chunk through one of the referenced target endpoints.
	MemoryChunk MemoryChunks
}

// UnmarshalJSON unmarshals a MemoryChunkInfo object from the raw JSON.
func (memorychunkinfo *MemoryChunkInfo) UnmarshalJSON(b []byte) error {
	type temp MemoryChunkInfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorychunkinfo = MemoryChunkInfo(t.temp)

	// Extract the links to other entities for later

	return nil
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

// VolumeInfo shall contain the combination of permissions and volume information.
type VolumeInfo struct {
	// AccessCapabilities shall specify a current storage access capability.
	AccessCapabilities []AccessCapability
	// AccessState shall contain the access state for the associated resource in this connection.
	AccessState AccessState
	// Volume shall contain a link to a resource of type Volume. The endpoints referenced by the InitiatorEndpoints or
	// InitiatorEndpointGroups properties shall be given access to this volume as described by this object. If
	// TargetEndpoints or TargetEndpointGroups is present, the referenced initiator endpoints shall be required to
	// access the referenced volume through one of the referenced target endpoints.
	Volume string
}

// UnmarshalJSON unmarshals a VolumeInfo object from the raw JSON.
func (volumeinfo *VolumeInfo) UnmarshalJSON(b []byte) error {
	type temp VolumeInfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*volumeinfo = VolumeInfo(t.temp)

	// Extract the links to other entities for later

	return nil
}
