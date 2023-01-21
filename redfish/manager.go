//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CommandConnectTypesSupported is
type CommandConnectTypesSupported string

const (
	// SSHCommandConnectTypesSupported The controller supports a command shell connection through the SSH protocol.
	SSHCommandConnectTypesSupported CommandConnectTypesSupported = "SSH"
	// TelnetCommandConnectTypesSupported The controller supports a command shell connection through the Telnet
	// protocol.
	TelnetCommandConnectTypesSupported CommandConnectTypesSupported = "Telnet"
	// IPMICommandConnectTypesSupported The controller supports a command shell connection through the IPMI Serial Over
	// LAN (SOL) protocol.
	IPMICommandConnectTypesSupported CommandConnectTypesSupported = "IPMI"
	// OemCommandConnectTypesSupported The controller supports a command shell connection through an OEM-specific
	// protocol.
	OemCommandConnectTypesSupported CommandConnectTypesSupported = "Oem"
)

// GraphicalConnectTypesSupported is
type GraphicalConnectTypesSupported string

const (
	// KVMIPGraphicalConnectTypesSupported The controller supports a graphical console connection through a KVM-IP
	// (redirection of Keyboard, Video, Mouse over IP) protocol.
	KVMIPGraphicalConnectTypesSupported GraphicalConnectTypesSupported = "KVMIP"
	// OemGraphicalConnectTypesSupported The controller supports a graphical console connection through an OEM-specific
	// protocol.
	OemGraphicalConnectTypesSupported GraphicalConnectTypesSupported = "Oem"
)

// ManagerType is
type ManagerType string

const (
	// ManagementControllerManagerType A controller that primarily monitors or manages the operation of a device or
	// system.
	ManagementControllerManagerType ManagerType = "ManagementController"
	// EnclosureManagerManagerType A controller that provides management functions for a chassis or group of devices or
	// systems.
	EnclosureManagerManagerType ManagerType = "EnclosureManager"
	// BMCManagerType A controller that provides management functions for a single computer system.
	BMCManagerType ManagerType = "BMC"
	// RackManagerManagerType A controller that provides management functions for a whole or part of a rack.
	RackManagerManagerType ManagerType = "RackManager"
	// AuxiliaryControllerManagerType A controller that provides management functions for a particular subsystem or
	// group of devices.
	AuxiliaryControllerManagerType ManagerType = "AuxiliaryController"
	// ServiceManagerType A software-based service that provides management functions.
	ServiceManagerType ManagerType = "Service"
)

// ResetToDefaultsType is
type ResetToDefaultsType string

const (
	// ResetAllResetToDefaultsType Reset all settings to factory defaults.
	ResetAllResetToDefaultsType ResetToDefaultsType = "ResetAll"
	// PreserveNetworkAndUsersResetToDefaultsType Reset all settings except network and local user names/passwords to
	// factory defaults.
	PreserveNetworkAndUsersResetToDefaultsType ResetToDefaultsType = "PreserveNetworkAndUsers"
	// PreserveNetworkResetToDefaultsType Reset all settings except network settings to factory defaults.
	PreserveNetworkResetToDefaultsType ResetToDefaultsType = "PreserveNetwork"
)

// SerialConnectTypesSupported is
type SerialConnectTypesSupported string

const (
	// SSHSerialConnectTypesSupported The controller supports a serial console connection through the SSH protocol.
	SSHSerialConnectTypesSupported SerialConnectTypesSupported = "SSH"
	// TelnetSerialConnectTypesSupported The controller supports a serial console connection through the Telnet
	// protocol.
	TelnetSerialConnectTypesSupported SerialConnectTypesSupported = "Telnet"
	// IPMISerialConnectTypesSupported The controller supports a serial console connection through the IPMI Serial Over
	// LAN (SOL) protocol.
	IPMISerialConnectTypesSupported SerialConnectTypesSupported = "IPMI"
	// OemSerialConnectTypesSupported The controller supports a serial console connection through an OEM-specific
	// protocol.
	OemSerialConnectTypesSupported SerialConnectTypesSupported = "Oem"
)

// CommandShell shall describe a command shell service for a manager.
type CommandShell struct {
	// ConnectTypesSupported shall contain an array of the enumerations. SSH shall be included if the Secure Shell
	// (SSH) protocol is supported. Telnet shall be included if the Telnet protocol is supported. IPMI shall be
	// included if the IPMI Serial Over LAN (SOL) protocol is supported.
	ConnectTypesSupported []CommandConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent service sessions that this implementation
	// supports.
	MaxConcurrentSessions string
	// ServiceEnabled shall indicate whether the protocol for the service is enabled.
	ServiceEnabled string
}

// UnmarshalJSON unmarshals a CommandShell object from the raw JSON.
func (commandshell *CommandShell) UnmarshalJSON(b []byte) error {
	type temp CommandShell
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*commandshell = CommandShell(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GraphicalConsole shall describe a graphical console service for a manager.
type GraphicalConsole struct {
	// ConnectTypesSupported shall contain an array of the enumerations. RDP shall be included if the Remote Desktop
	// (RDP) protocol is supported. KVMIP shall be included if a vendor-define KVM-IP protocol is supported.
	ConnectTypesSupported []GraphicalConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent service sessions that this implementation
	// supports.
	MaxConcurrentSessions string
	// ServiceEnabled shall indicate whether the protocol for the service is enabled.
	ServiceEnabled string
}

// UnmarshalJSON unmarshals a GraphicalConsole object from the raw JSON.
func (graphicalconsole *GraphicalConsole) UnmarshalJSON(b []byte) error {
	type temp GraphicalConsole
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*graphicalconsole = GraphicalConsole(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// ActiveSoftwareImage shall contain a link to a resource of type SoftwareInventory that represents the active
	// firmware image for this manager.
	ActiveSoftwareImage string
	// ManagedBy shall contain an array of links to resources of type Manager that represent the managers for this
	// manager.
	ManagedBy []Manager
	// ManagedBy@odata.count
	ManagedByCount int `json:"ManagedBy@odata.count"`
	// ManagerForChassis shall contain an array of links to chassis over which this manager instance has control.
	ManagerForChassis []Chassis
	// ManagerForChassis@odata.count
	ManagerForChassisCount int `json:"ManagerForChassis@odata.count"`
	// ManagerForManagers shall contain an array of links to resources of type Manager that represent the managers
	// being managed by this manager.
	ManagerForManagers []Manager
	// ManagerForManagers@odata.count
	ManagerForManagersCount int `json:"ManagerForManagers@odata.count"`
	// ManagerForServers shall contain an array of links to computer systems over which this manager instance has
	// control.
	ManagerForServers []ComputerSystem
	// ManagerForServers@odata.count
	ManagerForServersCount int `json:"ManagerForServers@odata.count"`
	// ManagerForSwitches shall contain an array of links to switches that this manager instance controls.
	ManagerForSwitches []Switch
	// ManagerForSwitches@odata.count
	ManagerForSwitchesCount int `json:"ManagerForSwitches@odata.count"`
	// ManagerInChassis shall contain a link to the chassis where this manager is located.
	ManagerInChassis string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SoftwareImages shall contain an array of links to resource of type SoftwareInventory that represent the firmware
	// images that apply to this manager.
	SoftwareImages []SoftwareInventory
	// SoftwareImages@odata.count
	SoftwareImagesCount int `json:"SoftwareImages@odata.count"`
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

// Manager shall represent a management subsystem for a Redfish implementation.
type Manager struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AdditionalFirmwareVersions shall contain the additional firmware versions of the manager.
	AdditionalFirmwareVersions string
	// AutoDSTEnabled shall indicate whether the manager is configured for automatic Daylight Saving Time (DST)
	// adjustment.
	AutoDSTEnabled string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	Certificates string
	// CommandShell shall contain information about the command shell service of this manager.
	CommandShell string
	// DateTime shall contain the current date and time with UTC offset of the manager.
	DateTime string
	// DateTimeLocalOffset shall contain the offset from UTC time that the DateTime property contains. If both DateTime
	// and DateTimeLocalOffset are provided in modification requests, services shall apply DateTimeLocalOffset after
	// DateTime is applied.
	DateTimeLocalOffset string
	// DedicatedNetworkPorts shall contain a link to a resource collection of type PortCollection that represent the
	// dedicated network ports of the manager.
	DedicatedNetworkPorts string
	// Description provides a description of this resource.
	Description string
	// EthernetInterfaces shall contain a link to a resource collection of type EthernetInterfaceCollection.
	EthernetInterfaces string
	// FirmwareVersion shall contain the firmware version as defined by the manufacturer for the associated manager.
	FirmwareVersion string
	// GraphicalConsole shall contain the information about the graphical console (KVM-IP) service of this manager.
	// This property should be used to describe a service for the manager's console or operating system, not a service
	// provided on behalf of a host operating system. Implementations representing host OS consoles, known generally as
	// a KVM-IP feature, should use the GraphicalConsole property in ComputerSystem.
	GraphicalConsole string
	// HostInterfaces shall contain a link to a resource collection of type HostInterfaceCollection.
	HostInterfaces string
	// LastResetTime shall contain the date and time when the manager last came out of a reset or was rebooted.
	LastResetTime string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated manager.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// LogServices shall contain a link to a resource collection of type LogServiceCollection that this manager uses.
	LogServices string
	// ManagerDiagnosticData shall contain a link to a resource of type ManagerDiagnosticData that represents the
	// diagnostic data for this manager.
	ManagerDiagnosticData ManagerDiagnosticData
	// ManagerType shall describe the function of this manager. The 'ManagementController' value shall be used if none
	// of the other enumerations apply.
	ManagerType string
	// Manufacturer shall contain the name of the organization responsible for producing the manager. This organization
	// may be the entity from whom the manager is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the information about how the manufacturer refers to this manager.
	Model string
	// NetworkProtocol shall contain a link to a resource of type ManagerNetworkProtocol, which represents the network
	// services for this manager.
	NetworkProtocol string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain a part number assigned by the organization that is responsible for producing or
	// manufacturing the manager.
	PartNumber string
	// PowerState shall contain the power state of the manager.
	PowerState PowerState
	// Redundancy shall show how this manager is grouped with other managers for form redundancy sets.
	Redundancy []Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// RemoteAccountService shall contain a link to the account service resource for the remote manager that this
	// resource represents. This property shall only be present when providing aggregation of a remote manager.
	RemoteAccountService string
	// RemoteRedfishServiceUri shall contain the URI of the Redfish service root for the remote manager that this
	// resource represents. This property shall only be present when providing aggregation of Redfish services.
	RemoteRedfishServiceUri string
	// SecurityPolicy shall contain a link to a resource of type SecurityPolicy that contains the security policy
	// settings for this manager.
	SecurityPolicy SecurityPolicy
	// SerialInterfaces shall contain a link to a resource collection of type SerialInterfaceCollection, which this
	// manager uses.
	SerialInterfaces string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the manager.
	SerialNumber string
	// ServiceEntryPointUUID shall contain the UUID of the Redfish service that is hosted by this manager. Each manager
	// providing an entry point to the same Redfish service shall report the same UUID value, even though the name of
	// the property may imply otherwise. This property shall not be present if this manager does not provide a Redfish
	// service entry point.
	ServiceEntryPointUUID UUID
	// ServiceIdentification shall contain a vendor or user-provided value that identifies and associates a discovered
	// Redfish service with a particular product instance. This property shall only be present if the manager provides
	// a ServiceRoot resource. The value of this property can be used during deployment processes to match user
	// credentials or other a priori product instance information to the appropriate Redfish service.
	ServiceIdentification string
	// SharedNetworkPorts shall contain a link to a resource collection of type PortCollection that represent the
	// shared network ports of the manager. The members of this collection shall reference Port resources subordinate
	// to NetworkAdapter resources.
	SharedNetworkPorts string
	// SparePartNumber shall contain the spare part number of the manager.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TimeZoneName shall contain the time zone of the manager. The time zone shall be either the 'Name' or the
	// 'Format' for the zone as defined in the IANA Time Zone Database. The value of this property is used for display
	// purposes, especially to enhance the display of time. A Redfish service may not be able to ensure accuracy and
	// consistency between the DateTimeOffset property and this property. Therefore, to specify the correct time zone
	// offset, see the DateTimeOffset property.
	TimeZoneName string
	// USBPorts shall contain a link to a resource collection of type PortCollection that represent the USB ports of
	// the manager.
	USBPorts string
	// UUID shall contain the UUID for the manager.
	UUID string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Manager object from the raw JSON.
func (manager *Manager) UnmarshalJSON(b []byte) error {
	type temp Manager
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*manager = Manager(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	manager.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (manager *Manager) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Manager)
	original.UnmarshalJSON(manager.rawData)

	readWriteFields := []string{
		"AutoDSTEnabled",
		"DateTime",
		"DateTimeLocalOffset",
		"LocationIndicatorActive",
		"ServiceIdentification",
		"TimeZoneName",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(manager).Elem()

	return manager.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetManager will get a Manager instance from the service.
func GetManager(c common.Client, uri string) (*Manager, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var manager Manager
	err = json.NewDecoder(resp.Body).Decode(&manager)
	if err != nil {
		return nil, err
	}

	manager.SetClient(c)
	return &manager, nil
}

// ListReferencedManagers gets the collection of Manager from
// a provided reference.
func ListReferencedManagers(c common.Client, link string) ([]*Manager, error) {
	var result []*Manager
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, managerLink := range links.ItemLinks {
		manager, err := GetManager(c, managerLink)
		if err != nil {
			collectionError.Failures[managerLink] = err
		} else {
			result = append(result, manager)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// ManagerService The manager services, such as serial console, command shell, or graphical console service.
type ManagerService struct {
	// MaxConcurrentSessions shall contain the maximum number of concurrent service sessions that this implementation
	// supports.
	MaxConcurrentSessions string
	// ServiceEnabled shall indicate whether the protocol for the service is enabled.
	ServiceEnabled string
}

// UnmarshalJSON unmarshals a ManagerService object from the raw JSON.
func (managerservice *ManagerService) UnmarshalJSON(b []byte) error {
	type temp ManagerService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*managerservice = ManagerService(t.temp)

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

// SerialConsole shall describe a serial console service for a manager.
type SerialConsole struct {
	// ConnectTypesSupported shall contain an array of the enumerations. SSH shall be included if the Secure Shell
	// (SSH) protocol is supported. Telnet shall be included if the Telnet protocol is supported. IPMI shall be
	// included if the IPMI Serial Over LAN (SOL) protocol is supported.
	ConnectTypesSupported []SerialConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent service sessions that this implementation
	// supports.
	MaxConcurrentSessions string
	// ServiceEnabled shall indicate whether the protocol for the service is enabled.
	ServiceEnabled string
}

// UnmarshalJSON unmarshals a SerialConsole object from the raw JSON.
func (serialconsole *SerialConsole) UnmarshalJSON(b []byte) error {
	type temp SerialConsole
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*serialconsole = SerialConsole(t.temp)

	// Extract the links to other entities for later

	return nil
}
