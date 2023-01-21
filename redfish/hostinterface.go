//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AuthenticationMode is
type AuthenticationMode string

const (
	// AuthNoneAuthenticationMode Requests without any sort of authentication are allowed.
	AuthNoneAuthenticationMode AuthenticationMode = "AuthNone"
	// BasicAuthAuthenticationMode Requests using HTTP Basic Authentication are allowed.
	BasicAuthAuthenticationMode AuthenticationMode = "BasicAuth"
	// RedfishSessionAuthAuthenticationMode Requests using Redfish Session Authentication are allowed.
	RedfishSessionAuthAuthenticationMode AuthenticationMode = "RedfishSessionAuth"
	// OemAuthAuthenticationMode Requests using OEM authentication mechanisms are allowed.
	OemAuthAuthenticationMode AuthenticationMode = "OemAuth"
)

// HostInterfaceType is
type HostInterfaceType string

const (
	// NetworkHostInterfaceHostInterfaceType This interface is a Network Host Interface.
	NetworkHostInterfaceHostInterfaceType HostInterfaceType = "NetworkHostInterface"
)

// CredentialBootstrapping shall contain settings for the Redfish Host Interface Specification-defined 'credential
// bootstrapping via IPMI commands' feature for this interface.
type CredentialBootstrapping struct {
	// EnableAfterReset shall indicate whether credential bootstrapping is enabled after a reset for this interface. If
	// 'true', services shall set the Enabled property to 'true' after a reset of the host or the service.
	EnableAfterReset bool
	// Enabled shall indicate whether credential bootstrapping is enabled for this interface.
	Enabled bool
	// RoleId shall contain the Id property of the role resource that is used for the bootstrap account created for
	// this interface.
	RoleId string
}

// UnmarshalJSON unmarshals a CredentialBootstrapping object from the raw JSON.
func (credentialbootstrapping *CredentialBootstrapping) UnmarshalJSON(b []byte) error {
	type temp CredentialBootstrapping
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*credentialbootstrapping = CredentialBootstrapping(t.temp)

	// Extract the links to other entities for later

	return nil
}

// HostInterface shall represent a Host Interface as part of the Redfish Specification.
type HostInterface struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this Resource.
	Actions string
	// AuthNoneRoleId shall contain the Id property of the Role Resource that is used when no authentication on this
	// interface is performed. This property shall contain absent if AuthNone is not supported by the service for the
	// AuthenticationModes property.
	AuthNoneRoleId string
	// AuthenticationModes shall contain an array consisting of the authentication modes allowed on this interface.
	AuthenticationModes []AuthenticationMode
	// CredentialBootstrapping shall contain settings for the Redfish Host Interface Specification-defined 'credential
	// bootstrapping via IPMI commands' feature for this interface. This property shall be absent if credential
	// bootstrapping is not supported by the service.
	CredentialBootstrapping string
	// Description provides a description of this resource.
	Description string
	// ExternallyAccessible shall indicate whether external entities can access this interface. External entities are
	// non-host entities. For example, if the host and manager are connected through a switch and the switch also
	// exposes an external port on the system, external clients can also use the interface, and this property value is
	// 'true'.
	ExternallyAccessible bool
	// HostEthernetInterfaces shall contain a link to a Resource Collection of type EthernetInterface that computer
	// systems use as the Host Interface to this manager.
	HostEthernetInterfaces string
	// HostInterfaceType shall contain an enumeration that describes the type of the interface.
	HostInterfaceType HostInterfaceType
	// InterfaceEnabled shall indicate whether this interface is enabled.
	InterfaceEnabled bool
	// Links shall contain links to Resources related to but not subordinate to this Resource.
	Links string
	// ManagerEthernetInterface shall contain a link to a Resource of type EthernetInterface that represents the
	// network interface that this manager uses as the Host Interface.
	ManagerEthernetInterface string
	// NetworkProtocol shall contain a link to a Resource of type ManagerNetworkProtocol that represents the network
	// services for this manager.
	NetworkProtocol string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a HostInterface object from the raw JSON.
func (hostinterface *HostInterface) UnmarshalJSON(b []byte) error {
	type temp HostInterface
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*hostinterface = HostInterface(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	hostinterface.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (hostinterface *HostInterface) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(HostInterface)
	original.UnmarshalJSON(hostinterface.rawData)

	readWriteFields := []string{
		"AuthNoneRoleId",
		"AuthenticationModes",
		"InterfaceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(hostinterface).Elem()

	return hostinterface.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetHostInterface will get a HostInterface instance from the service.
func GetHostInterface(c common.Client, uri string) (*HostInterface, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var hostinterface HostInterface
	err = json.NewDecoder(resp.Body).Decode(&hostinterface)
	if err != nil {
		return nil, err
	}

	hostinterface.SetClient(c)
	return &hostinterface, nil
}

// ListReferencedHostInterfaces gets the collection of HostInterface from
// a provided reference.
func ListReferencedHostInterfaces(c common.Client, link string) ([]*HostInterface, error) {
	var result []*HostInterface
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, hostinterfaceLink := range links.ItemLinks {
		hostinterface, err := GetHostInterface(c, hostinterfaceLink)
		if err != nil {
			collectionError.Failures[hostinterfaceLink] = err
		} else {
			result = append(result, hostinterface)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Links shall contain links to Resources related to but not subordinate to this Resource.
type Links struct {
	// AuthNoneRole shall contain a link to a Resource of type Role, and should link to the Resource identified by
	// property AuthNoneRoleId. This property shall be absent if AuthNone is not supported by the service for the
	// AuthenticationModes property.
	AuthNoneRole string
	// ComputerSystems shall contain an array of links to Resources of the ComputerSystem type that are connected to
	// this Host Interface.
	ComputerSystems []ComputerSystem
	// ComputerSystems@odata.count
	ComputerSystemsCount int `json:"ComputerSystems@odata.count"`
	// CredentialBootstrappingRole shall contain a link to a resource of type Role, and should link to the resource
	// identified by the RoleId property within CredentialBootstrapping. This property shall be absent if the Redfish
	// Host Interface Specification-defined 'credential bootstrapping via IPMI commands' feature is not supported by
	// the service.
	CredentialBootstrappingRole string
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
