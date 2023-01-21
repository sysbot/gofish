//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AutomaticRetryConfig is
type AutomaticRetryConfig string

const (
	// DisabledAutomaticRetryConfig shall indicate that automatic retrying of booting is disabled.
	DisabledAutomaticRetryConfig AutomaticRetryConfig = "Disabled"
	// RetryAttemptsAutomaticRetryConfig shall indicate that the number of retries of booting is based on the
	// AutomaticRetryAttempts property, and the RemainingAutomaticRetryAttempts property indicates the number of
	// remaining attempts.
	RetryAttemptsAutomaticRetryConfig AutomaticRetryConfig = "RetryAttempts"
	// RetryAlwaysAutomaticRetryConfig shall indicate that the system will always automatically retry booting.
	RetryAlwaysAutomaticRetryConfig AutomaticRetryConfig = "RetryAlways"
)

// BootOrderTypes is The enumerations of BootOrderTypes specify the choice of boot order property to use when
// controller the persistent boot order for this computer system.
type BootOrderTypes string

const (
	// BootOrderBootOrderTypes The system uses the BootOrder property to specify the persistent boot order.
	BootOrderBootOrderTypes BootOrderTypes = "BootOrder"
	// AliasBootOrderBootOrderTypes The system uses the AliasBootOrder property to specify the persistent boot order.
	AliasBootOrderBootOrderTypes BootOrderTypes = "AliasBootOrder"
)

// BootProgressTypes is
type BootProgressTypes string

const (
	// NoneBootProgressTypes shall indicate that the system is not booting or running, such as the system is powered
	// off.
	NoneBootProgressTypes BootProgressTypes = "None"
	// PrimaryProcessorInitializationStartedBootProgressTypes shall indicate that the system has started to initialize
	// the primary processor.
	PrimaryProcessorInitializationStartedBootProgressTypes BootProgressTypes = "PrimaryProcessorInitializationStarted"
	// BusInitializationStartedBootProgressTypes shall indicate that the system has started to initialize the buses.
	BusInitializationStartedBootProgressTypes BootProgressTypes = "BusInitializationStarted"
	// MemoryInitializationStartedBootProgressTypes shall indicate that the system has started to initialize memory.
	MemoryInitializationStartedBootProgressTypes BootProgressTypes = "MemoryInitializationStarted"
	// SecondaryProcessorInitializationStartedBootProgressTypes shall indicate that the system has started to
	// initialize the secondary processors.
	SecondaryProcessorInitializationStartedBootProgressTypes BootProgressTypes = "SecondaryProcessorInitializationStarted"
	// PCIResourceConfigStartedBootProgressTypes shall indicate that the system has started to initialize PCI
	// resources.
	PCIResourceConfigStartedBootProgressTypes BootProgressTypes = "PCIResourceConfigStarted"
	// SystemHardwareInitializationCompleteBootProgressTypes shall indicate that the system has completed initializing
	// all hardware.
	SystemHardwareInitializationCompleteBootProgressTypes BootProgressTypes = "SystemHardwareInitializationComplete"
	// SetupEnteredBootProgressTypes shall indicate that the system has entered the setup utility.
	SetupEnteredBootProgressTypes BootProgressTypes = "SetupEntered"
	// OSBootStartedBootProgressTypes shall indicate that the operating system has started to boot.
	OSBootStartedBootProgressTypes BootProgressTypes = "OSBootStarted"
	// OSRunningBootProgressTypes shall indicate that the operating system is running and shall indicate the final boot
	// progress state.
	OSRunningBootProgressTypes BootProgressTypes = "OSRunning"
	// OEMBootProgressTypes shall indicate an OEM-defined boot progress state.
	OEMBootProgressTypes BootProgressTypes = "OEM"
)

// BootSourceOverrideEnabled is
type BootSourceOverrideEnabled string

const (
	// DisabledBootSourceOverrideEnabled The system boots normally.
	DisabledBootSourceOverrideEnabled BootSourceOverrideEnabled = "Disabled"
	// OnceBootSourceOverrideEnabled On its next boot cycle, the system boots one time to the boot source override
	// target. Then, the BootSourceOverrideEnabled value is reset to 'Disabled'.
	OnceBootSourceOverrideEnabled BootSourceOverrideEnabled = "Once"
	// ContinuousBootSourceOverrideEnabled The system boots to the target specified in the BootSourceOverrideTarget
	// property until this property is 'Disabled'.
	ContinuousBootSourceOverrideEnabled BootSourceOverrideEnabled = "Continuous"
)

// BootSourceOverrideMode is
type BootSourceOverrideMode string

const (
	// LegacyBootSourceOverrideMode The system boots in non-UEFI boot mode to the boot source override target.
	LegacyBootSourceOverrideMode BootSourceOverrideMode = "Legacy"
	// UEFIBootSourceOverrideMode The system boots in UEFI boot mode to the boot source override target.
	UEFIBootSourceOverrideMode BootSourceOverrideMode = "UEFI"
)

// CompositionUseCase is
type CompositionUseCase string

const (
	// ResourceBlockCapableCompositionUseCase shall indicate the computer system supports being registered as a
	// resource block in order for it to participate in composition requests.
	ResourceBlockCapableCompositionUseCase CompositionUseCase = "ResourceBlockCapable"
	// ExpandableSystemCompositionUseCase shall indicate the computer system supports expandable system composition and
	// is associated with a resource block.
	ExpandableSystemCompositionUseCase CompositionUseCase = "ExpandableSystem"
)

// GraphicalConnectTypesSupported is
type GraphicalConnectTypesSupported string

const (
	// KVMIPGraphicalConnectTypesSupported The controller supports a graphical console connection through a KVM-IP
	// (redirection of Keyboard, Video, Mouse over IP) protocol.
	KVMIPGraphicalConnectTypesSupported GraphicalConnectTypesSupported = "KVMIP"
	// OEMGraphicalConnectTypesSupported The controller supports a graphical console connection through an OEM-specific
	// protocol.
	OEMGraphicalConnectTypesSupported GraphicalConnectTypesSupported = "OEM"
)

// HostingRole is The enumerations of HostingRoles specify different features that the hosting ComputerSystem
// supports.
type HostingRole string

const (
	// ApplicationServerHostingRole The system hosts functionality that supports general purpose applications.
	ApplicationServerHostingRole HostingRole = "ApplicationServer"
	// StorageServerHostingRole The system hosts functionality that supports the system acting as a storage server.
	StorageServerHostingRole HostingRole = "StorageServer"
	// SwitchHostingRole The system hosts functionality that supports the system acting as a switch.
	SwitchHostingRole HostingRole = "Switch"
	// ApplianceHostingRole The system hosts functionality that supports the system acting as an appliance.
	ApplianceHostingRole HostingRole = "Appliance"
	// BareMetalServerHostingRole The system hosts functionality that supports the system acting as a bare metal
	// server.
	BareMetalServerHostingRole HostingRole = "BareMetalServer"
	// VirtualMachineServerHostingRole The system hosts functionality that supports the system acting as a virtual
	// machine server.
	VirtualMachineServerHostingRole HostingRole = "VirtualMachineServer"
	// ContainerServerHostingRole The system hosts functionality that supports the system acting as a container server.
	ContainerServerHostingRole HostingRole = "ContainerServer"
)

// IndicatorLED is
type IndicatorLED string

const (
	// UnknownIndicatorLED shall represent that the indicator LED is in an unknown state. The service shall reject
	// PATCH or PUT requests containing this value by returning the HTTP 400 (Bad Request) status code.
	UnknownIndicatorLED IndicatorLED = "Unknown"
	// LitIndicatorLED shall represent that the indicator LED is in a solid on state. If the service does not support
	// this value, it shall reject PATCH or PUT requests containing this value by returning the HTTP 400 (Bad Request)
	// status code.
	LitIndicatorLED IndicatorLED = "Lit"
	// BlinkingIndicatorLED shall represent that the indicator LED is in a blinking state where the LED is being turned
	// on and off in repetition. If the service does not support this value, it shall reject PATCH or PUT requests
	// containing this value by returning the HTTP 400 (Bad Request) status code.
	BlinkingIndicatorLED IndicatorLED = "Blinking"
	// OffIndicatorLED shall represent that the indicator LED is in a solid off state. If the service does not support
	// this value, it shall reject PATCH or PUT requests containing this value by returning the HTTP 400 (Bad Request)
	// status code.
	OffIndicatorLED IndicatorLED = "Off"
)

// InterfaceType is
type InterfaceType string

const (
	// TPM12InterfaceType Trusted Platform Module (TPM) 1.2.
	TPM12InterfaceType InterfaceType = "TPM1_2"
	// TPM20InterfaceType Trusted Platform Module (TPM) 2.0.
	TPM20InterfaceType InterfaceType = "TPM2_0"
	// TCM10InterfaceType Trusted Cryptography Module (TCM) 1.0.
	TCM10InterfaceType InterfaceType = "TCM1_0"
)

// InterfaceTypeSelection is The enumerations of InterfaceTypeSelection specify the method for switching the
// TrustedModule InterfaceType, for instance between TPM1_2 and TPM2_0, if supported.
type InterfaceTypeSelection string

const (
	// NoneInterfaceTypeSelection The TrustedModule does not support switching the InterfaceType.
	NoneInterfaceTypeSelection InterfaceTypeSelection = "None"
	// FirmwareUpdateInterfaceTypeSelection The TrustedModule supports switching InterfaceType through a firmware
	// update.
	FirmwareUpdateInterfaceTypeSelection InterfaceTypeSelection = "FirmwareUpdate"
	// BiosSettingInterfaceTypeSelection The TrustedModule supports switching InterfaceType through platform software,
	// such as a BIOS configuration attribute.
	BiosSettingInterfaceTypeSelection InterfaceTypeSelection = "BiosSetting"
	// OemMethodInterfaceTypeSelection The TrustedModule supports switching InterfaceType through an OEM proprietary
	// mechanism.
	OemMethodInterfaceTypeSelection InterfaceTypeSelection = "OemMethod"
)

// MemoryMirroring is
type MemoryMirroring string

const (
	// SystemMemoryMirroring The system supports DIMM mirroring at the system level. Individual DIMMs are not paired
	// for mirroring in this mode.
	SystemMemoryMirroring MemoryMirroring = "System"
	// DIMMMemoryMirroring The system supports DIMM mirroring at the DIMM level. Individual DIMMs can be mirrored.
	DIMMMemoryMirroring MemoryMirroring = "DIMM"
	// HybridMemoryMirroring The system supports a hybrid mirroring at the system and DIMM levels. Individual DIMMs can
	// be mirrored.
	HybridMemoryMirroring MemoryMirroring = "Hybrid"
	// NoneMemoryMirroring The system does not support DIMM mirroring.
	NoneMemoryMirroring MemoryMirroring = "None"
)

// PowerMode is
type PowerMode string

const (
	// MaximumPerformancePowerMode shall indicate the system performs at the highest speeds possible. This mode should
	// be used when performance is the top priority.
	MaximumPerformancePowerMode PowerMode = "MaximumPerformance"
	// BalancedPerformancePowerMode shall indicate the system performs at the highest speeds possible when the
	// utilization is high and performs at reduced speeds when the utilization is low to save power. This mode is a
	// compromise between 'MaximumPerformance' and 'PowerSaving'.
	BalancedPerformancePowerMode PowerMode = "BalancedPerformance"
	// PowerSavingPowerMode shall indicate the system performs at reduced speeds to save power. This mode should be
	// used when power saving is the top priority.
	PowerSavingPowerMode PowerMode = "PowerSaving"
	// StaticPowerMode shall indicate the system performs at a static base speed.
	StaticPowerMode PowerMode = "Static"
	// OSControlledPowerMode shall indicate the system performs at a operating system controlled power mode.
	OSControlledPowerMode PowerMode = "OSControlled"
	// OEMPowerMode shall indicate the system performs at an OEM-defined power mode.
	OEMPowerMode PowerMode = "OEM"
)

// PowerRestorePolicyTypes is The enumerations of PowerRestorePolicyTypes specify the choice of power state for the
// system when power is applied.
type PowerRestorePolicyTypes string

const (
	// AlwaysOnPowerRestorePolicyTypes The system always powers on when power is applied.
	AlwaysOnPowerRestorePolicyTypes PowerRestorePolicyTypes = "AlwaysOn"
	// AlwaysOffPowerRestorePolicyTypes The system always remains powered off when power is applied.
	AlwaysOffPowerRestorePolicyTypes PowerRestorePolicyTypes = "AlwaysOff"
	// LastStatePowerRestorePolicyTypes The system returns to its last on or off power state when power is applied.
	LastStatePowerRestorePolicyTypes PowerRestorePolicyTypes = "LastState"
)

// PowerState is
type PowerState string

const (
	// OnPowerState The system is powered on.
	OnPowerState PowerState = "On"
	// OffPowerState The system is powered off, although some components might continue to have AUX power such as
	// management controller.
	OffPowerState PowerState = "Off"
	// PoweringOnPowerState A temporary state between off and on. This temporary state can be very short.
	PoweringOnPowerState PowerState = "PoweringOn"
	// PoweringOffPowerState A temporary state between on and off. The power off action can take time while the OS is
	// in the shutdown process.
	PoweringOffPowerState PowerState = "PoweringOff"
)

// StopBootOnFault is
type StopBootOnFault string

const (
	// NeverStopBootOnFault shall indicate the system will continue to attempt to boot if a fault occurs.
	NeverStopBootOnFault StopBootOnFault = "Never"
	// AnyFaultStopBootOnFault shall indicate the system will stop the boot if a fault occurs. This includes, but is
	// not limited to, faults that affect performance, fault tolerance, or capacity.
	AnyFaultStopBootOnFault StopBootOnFault = "AnyFault"
)

// SystemType is
type SystemType string

const (
	// PhysicalSystemType A SystemType of Physical typically represents the hardware aspects of a system, such as a
	// management controller.
	PhysicalSystemType SystemType = "Physical"
	// VirtualSystemType A SystemType of Virtual typically represents a system that is actually a virtual machine
	// instance.
	VirtualSystemType SystemType = "Virtual"
	// OSSystemType A SystemType of OS typically represents an OS or hypervisor view of the system.
	OSSystemType SystemType = "OS"
	// PhysicallyPartitionedSystemType A SystemType of PhysicallyPartitioned typically represents a single system
	// constructed from one or more physical systems through a firmware or hardware-based service.
	PhysicallyPartitionedSystemType SystemType = "PhysicallyPartitioned"
	// VirtuallyPartitionedSystemType A SystemType of VirtuallyPartitioned typically represents a single system
	// constructed from one or more virtual systems through a software-based service.
	VirtuallyPartitionedSystemType SystemType = "VirtuallyPartitioned"
	// ComposedSystemType A SystemType of Composed typically represents a single system constructed from disaggregated
	// resources through the Redfish composition service.
	ComposedSystemType SystemType = "Composed"
	// DPUSystemType A SystemType of DPU typically represents a single system that performs offload computation as a
	// data processing unit, such as a SmartNIC.
	DPUSystemType SystemType = "DPU"
)

// TrustedModuleRequiredToBoot is
type TrustedModuleRequiredToBoot string

const (
	// DisabledTrustedModuleRequiredToBoot shall indicate a Trusted Module is not required to boot.
	DisabledTrustedModuleRequiredToBoot TrustedModuleRequiredToBoot = "Disabled"
	// RequiredTrustedModuleRequiredToBoot shall indicate a functioning Trusted Module is required to boot.
	RequiredTrustedModuleRequiredToBoot TrustedModuleRequiredToBoot = "Required"
)

// WatchdogTimeoutActions is The enumerations of WatchdogTimeoutActions specify the choice of action to take when
// the host watchdog timer reaches its timeout value.
type WatchdogTimeoutActions string

const (
	// NoneWatchdogTimeoutActions No action taken.
	NoneWatchdogTimeoutActions WatchdogTimeoutActions = "None"
	// ResetSystemWatchdogTimeoutActions Reset the system.
	ResetSystemWatchdogTimeoutActions WatchdogTimeoutActions = "ResetSystem"
	// PowerCycleWatchdogTimeoutActions Power cycle the system.
	PowerCycleWatchdogTimeoutActions WatchdogTimeoutActions = "PowerCycle"
	// PowerDownWatchdogTimeoutActions Power down the system.
	PowerDownWatchdogTimeoutActions WatchdogTimeoutActions = "PowerDown"
	// OEMWatchdogTimeoutActions Perform an OEM-defined action.
	OEMWatchdogTimeoutActions WatchdogTimeoutActions = "OEM"
)

// WatchdogWarningActions is The enumerations of WatchdogWarningActions specify the choice of action to take when
// the host watchdog timer is close (typically 3-10 seconds) to reaching its timeout value.
type WatchdogWarningActions string

const (
	// NoneWatchdogWarningActions No action taken.
	NoneWatchdogWarningActions WatchdogWarningActions = "None"
	// DiagnosticInterruptWatchdogWarningActions Raise a (typically non-maskable) Diagnostic Interrupt.
	DiagnosticInterruptWatchdogWarningActions WatchdogWarningActions = "DiagnosticInterrupt"
	// SMIWatchdogWarningActions Raise a Systems Management Interrupt (SMI).
	SMIWatchdogWarningActions WatchdogWarningActions = "SMI"
	// MessagingInterruptWatchdogWarningActions Raise a legacy IPMI messaging interrupt.
	MessagingInterruptWatchdogWarningActions WatchdogWarningActions = "MessagingInterrupt"
	// SCIWatchdogWarningActions Raise an interrupt using the ACPI System Control Interrupt (SCI).
	SCIWatchdogWarningActions WatchdogWarningActions = "SCI"
	// OEMWatchdogWarningActions Perform an OEM-defined action.
	OEMWatchdogWarningActions WatchdogWarningActions = "OEM"
)

// Boot shall contain properties that describe boot information for a system.
type Boot struct {
	// AliasBootOrder shall contain an ordered array of boot source aliases of the BootSource type that represents the
	// persistent boot order of this computer system. This array shall not contain duplicate values. Virtual devices
	// for an alias should take precedence over a physical device. Systems may attempt to boot from multiple devices
	// that share an alias.
	AliasBootOrder []BootSource
	// AutomaticRetryAttempts shall contain the number of attempts the system will automatically retry booting in the
	// event the system enters an error state on boot.
	AutomaticRetryAttempts int
	// AutomaticRetryConfig shall contain the configuration of how the system retry booting automatically.
	AutomaticRetryConfig AutomaticRetryConfig
	// BootNext shall contain the BootOptionReference of the UEFI boot option for one time boot, as defined by the UEFI
	// Specification. The valid values for this property are specified in the values of the BootOrder array.
	// BootSourceOverrideEnabled set to 'Continuous' is not supported for BootSourceOverrideTarget set to
	// 'UefiBootNext' because this setting is defined in UEFI as a one-time boot setting.
	BootNext string
	// BootOptions shall contain a link to a resource collection of type BootOptionCollection.
	BootOptions string
	// BootOrder shall contain an array of BootOptionReference strings that represent the persistent boot order for
	// this computer system. For UEFI systems, this is the UEFI Specification-defined UEFI BootOrder.
	BootOrder []string
	// BootOrderPropertySelection shall indicate which boot order property the system uses for the persistent boot
	// order.
	BootOrderPropertySelection BootOrderTypes
	// BootSourceOverrideEnabled shall contain 'Once' for a one-time boot override, and 'Continuous' for a remain-
	// active-until-cancelled override. If set to 'Once', the value is reset to 'Disabled' after the
	// BootSourceOverrideTarget actions have completed successfully. Changes to this property do not alter the BIOS
	// persistent boot order configuration.
	BootSourceOverrideEnabled BootSourceOverrideEnabled
	// BootSourceOverrideMode shall contain the BIOS boot mode to use when the system boots from the
	// BootSourceOverrideTarget boot source.
	BootSourceOverrideMode BootSourceOverrideMode
	// BootSourceOverrideTarget shall contain the source to boot the system from, overriding the normal boot order. The
	// @Redfish.AllowableValues annotation specifies the valid values for this property. 'UefiTarget' indicates to boot
	// from the UEFI device path found in UefiTargetBootSourceOverride. 'UefiBootNext' indicates to boot from the UEFI
	// BootOptionReference found in BootNext. Virtual devices for a target should take precedence over a physical
	// device. Systems may attempt to boot from multiple devices that share a target identifier. Changes to this
	// property do not alter the BIOS persistent boot order configuration.
	BootSourceOverrideTarget BootSource
	// Certificates shall contain a link to a resource collection of type CertificateCollection.
	Certificates string
	// HttpBootUri shall contain the URI to perform an HTTP or HTTPS boot when BootSourceOverrideTarget is set to
	// 'UefiHttp'.
	HttpBootUri string
	// RemainingAutomaticRetryAttempts shall contain the number of attempts remaining the system will retry booting in
	// the event the system enters an error state on boot. If '0', the system has no remaining automatic boot retry
	// attempts and shall not automatically retry booting if the system enters an error state. This property shall be
	// reset to the value of AutomaticRetryAttempts upon a successful boot attempt.
	RemainingAutomaticRetryAttempts int
	// StopBootOnFault shall contain the setting if the boot should stop on a fault.
	StopBootOnFault StopBootOnFault
	// TrustedModuleRequiredToBoot shall contain the Trusted Module boot requirement.
	TrustedModuleRequiredToBoot TrustedModuleRequiredToBoot
	// UefiTargetBootSourceOverride shall contain the UEFI device path of the override boot target. Changes to this
	// property do not alter the BIOS persistent boot order configuration.
	UefiTargetBootSourceOverride string
}

// UnmarshalJSON unmarshals a Boot object from the raw JSON.
func (boot *Boot) UnmarshalJSON(b []byte) error {
	type temp Boot
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*boot = Boot(t.temp)

	// Extract the links to other entities for later

	return nil
}

// BootProgress shall contain the last boot progress state and time.
type BootProgress struct {
	// LastBootTimeSeconds shall contain the number of seconds that elapsed between system reset or power on and
	// LastState transitioning to 'OSRunning'. If LastState contains 'OSRunning', this property shall contain the most
	// recent boot time. For other values of LastState, this property shall contain the boot time for the previous
	// boot.
	LastBootTimeSeconds float64
	// LastState shall contain the last boot progress state.
	LastState BootProgressTypes
	// LastStateTime shall contain the date and time when the last boot state was updated.
	LastStateTime string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OemLastState shall represent the OEM-specific LastState of the BootProgress. This property shall only be present
	// if LastState is 'OEM'.
	OemLastState string
}

// UnmarshalJSON unmarshals a BootProgress object from the raw JSON.
func (bootprogress *BootProgress) UnmarshalJSON(b []byte) error {
	type temp BootProgress
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*bootprogress = BootProgress(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Composition shall contain information about the composition capabilities and state of a computer system.
type Composition struct {
	// UseCases shall contain the composition use cases in which this computer system can participate.
	UseCases []CompositionUseCase
}

// UnmarshalJSON unmarshals a Composition object from the raw JSON.
func (composition *Composition) UnmarshalJSON(b []byte) error {
	type temp Composition
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*composition = Composition(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ComputerSystem shall represent a computing system in the Redfish Specification.
type ComputerSystem struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AssetTag shall contain the system asset tag value.
	AssetTag string
	// Bios shall contain a link to a resource of type Bios that lists the BIOS settings for this system.
	Bios string
	// BiosVersion shall contain the version string of the currently installed and running BIOS for x86 systems. For
	// other systems, the property may contain a version string that represents the primary system firmware.
	BiosVersion string
	// Boot shall contain the boot settings for this system.
	Boot string
	// BootProgress shall contain the last boot progress state and time.
	BootProgress BootProgress
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	Certificates string
	// Composition shall contain information about the composition capabilities and state of the computer system.
	Composition Composition
	// Description provides a description of this resource.
	Description string
	// EthernetInterfaces shall contain a link to a resource collection of type EthernetInterfaceCollection.
	EthernetInterfaces string
	// FabricAdapters shall contain a link to a resource collection of type FabricAdapterCollection.
	FabricAdapters string
	// GraphicalConsole shall contain the information about the graphical console (KVM-IP) service of this system.
	GraphicalConsole string
	// GraphicsControllers shall contain a link to a resource collection of type GraphicsControllerCollection that
	// contains graphics controllers that can output video for this system.
	GraphicsControllers string
	// HostName shall contain the host name for this system, as reported by the operating system or hypervisor. A
	// service running in the host operating system typically reports this value to the manager.
	HostName string
	// HostWatchdogTimer shall contain properties that describe the host watchdog timer functionality for this
	// ComputerSystem.
	HostWatchdogTimer string
	// HostedServices shall describe services that this computer system supports.
	HostedServices string
	// HostingRoles shall contain the hosting roles that this computer system supports.
	HostingRoles []HostingRole
	// IdlePowerSaver shall contain the idle power saver settings of the computer system.
	IdlePowerSaver IdlePowerSaver
	// KeyManagement shall contain the key management settings of the computer system.
	KeyManagement KeyManagement
	// LastResetTime shall contain the date and time when the system last came out of a reset or was rebooted.
	LastResetTime string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// LogServices shall contain a link to a resource collection of type LogServiceCollection.
	LogServices string
	// Manufacturer shall contain a value that represents the manufacturer of the system.
	Manufacturer string
	// ManufacturingMode shall indicate whether the system is in manufacturing mode. If the system supports SMBIOS, the
	// value shall match the 'Manufacturing mode is enabled' setting from the 'BIOS Characteristics' entry.
	ManufacturingMode bool
	// Memory shall contain a link to a resource collection of type MemoryCollection.
	Memory string
	// MemoryDomains shall contain a link to a resource collection of type MemoryDomainCollection.
	MemoryDomains string
	// MemorySummary shall describe the central memory for this resource.
	MemorySummary string
	// Model shall describe how the manufacturer refers to this system. Typically, this value is the product name for
	// this system without the manufacturer name.
	Model string
	// NetworkInterfaces shall contain a link to a resource collection of type NetworkInterfaceCollection.
	NetworkInterfaces string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevices shall contain an array of links of the PCIeDevice type.
	PCIeDevices []PCIeDevice
	// PCIeDevices@odata.count
	PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
	// PCIeFunctions shall contain an array of links of the PCIeFunction type.
	PCIeFunctions []PCIeFunction
	// PCIeFunctions@odata.count
	PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	// PartNumber shall contain the manufacturer-defined part number for the system.
	PartNumber string
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on after a 'Reset' action requesting
	// 'PowerCycle'. The value '0' shall indicate no delay to power on.
	PowerCycleDelaySeconds float64
	// PowerMode shall contain the computer system power mode setting.
	PowerMode PowerMode
	// PowerOffDelaySeconds shall contain the number of seconds to delay power off during a reset. The value '0' shall
	// indicate no delay to power off.
	PowerOffDelaySeconds float64
	// PowerOnDelaySeconds shall contain the number of seconds to delay power on after a power cycle or during a reset.
	// The value '0' shall indicate no delay to power on.
	PowerOnDelaySeconds float64
	// PowerRestorePolicy shall indicate the desired PowerState of the system when power is applied to the system. The
	// 'LastState' value shall return the system to the PowerState it was in when power was lost.
	PowerRestorePolicy string
	// PowerState shall contain the power state of the system.
	PowerState PowerState
	// ProcessorSummary shall describe the central processors for this resource. Processors described by this property
	// shall be limited to the processors that execute system code, and shall not include processors used for offload
	// functionality.
	ProcessorSummary string
	// Processors shall contain a link to a resource collection of type ProcessorCollection.
	Processors string
	// Redundancy shall contain a set of redundancy entities. Each entity specifies a kind and level of redundancy and
	// a collection, or redundancy set, of other computer systems that provide the specified redundancy to this
	// computer system.
	Redundancy []Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// SKU shall contain the SKU for the system.
	SKU string
	// SecureBoot shall contain a link to a resource of type SecureBoot.
	SecureBoot string
	// SerialConsole shall contain information about the serial console services of this system.
	SerialConsole string
	// SerialNumber shall contain the serial number for the system.
	SerialNumber string
	// SimpleStorage shall contain a link to a resource collection of type SimpleStorageCollection.
	SimpleStorage string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Storage shall contain a link to a resource collection of type StorageCollection.
	Storage string
	// SubModel shall contain the information about the sub-model (or configuration) of the system. This shall not
	// include the model/product name or the manufacturer name.
	SubModel string
	// SystemType An enumeration that indicates the kind of system that this resource represents.
	SystemType string
	// USBControllers shall contain a link to a resource collection of type USBControllerCollection that contains USB
	// controllers for this system.
	USBControllers string
	// UUID shall contain the universal unique identifier number for this system. RFC4122 describes methods to create
	// this value. The value should be considered to be opaque. Client software should only treat the overall value as
	// a UUID and should not interpret any sub-fields within the UUID. If the system supports SMBIOS, the property
	// value should follow the SMBIOS 2.6 and later recommendation for converting the SMBIOS 16-byte UUID structure
	// into the Redfish canonical 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' string format, so that the property value
	// matches the byte order presented by current OS APIs, such as WMI and dmidecode.
	UUID string
	// VirtualMedia shall contain a link to a resource collection of type VirtualMediaCollection that this system uses.
	VirtualMedia string
	// VirtualMediaConfig shall contain the information about the virtual media service of this system.
	VirtualMediaConfig string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ComputerSystem object from the raw JSON.
func (computersystem *ComputerSystem) UnmarshalJSON(b []byte) error {
	type temp ComputerSystem
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*computersystem = ComputerSystem(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	computersystem.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (computersystem *ComputerSystem) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ComputerSystem)
	original.UnmarshalJSON(computersystem.rawData)

	readWriteFields := []string{
		"AssetTag",
		"HostName",
		"LocationIndicatorActive",
		"PowerCycleDelaySeconds",
		"PowerMode",
		"PowerOffDelaySeconds",
		"PowerOnDelaySeconds",
		"PowerRestorePolicy",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(computersystem).Elem()

	return computersystem.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetComputerSystem will get a ComputerSystem instance from the service.
func GetComputerSystem(c common.Client, uri string) (*ComputerSystem, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var computersystem ComputerSystem
	err = json.NewDecoder(resp.Body).Decode(&computersystem)
	if err != nil {
		return nil, err
	}

	computersystem.SetClient(c)
	return &computersystem, nil
}

// ListReferencedComputerSystems gets the collection of ComputerSystem from
// a provided reference.
func ListReferencedComputerSystems(c common.Client, link string) ([]*ComputerSystem, error) {
	var result []*ComputerSystem
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, computersystemLink := range links.ItemLinks {
		computersystem, err := GetComputerSystem(c, computersystemLink)
		if err != nil {
			collectionError.Failures[computersystemLink] = err
		} else {
			result = append(result, computersystem)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// HostGraphicalConsole shall describe a graphical console service for a computer system.
type HostGraphicalConsole struct {
	// ConnectTypesSupported shall contain an array of the enumerations. KVMIP shall be included if a vendor-define
	// KVM-IP protocol is supported.
	ConnectTypesSupported []GraphicalConnectTypesSupported
	// MaxConcurrentSessions shall contain the maximum number of concurrent service sessions that this implementation
	// supports.
	MaxConcurrentSessions string
	// Port shall contain the port assigned to the service.
	Port int
	// ServiceEnabled shall indicate whether the protocol for the service is enabled.
	ServiceEnabled string
}

// UnmarshalJSON unmarshals a HostGraphicalConsole object from the raw JSON.
func (hostgraphicalconsole *HostGraphicalConsole) UnmarshalJSON(b []byte) error {
	type temp HostGraphicalConsole
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*hostgraphicalconsole = HostGraphicalConsole(t.temp)

	// Extract the links to other entities for later

	return nil
}

// HostSerialConsole shall describe the serial console services for a computer system.
type HostSerialConsole struct {
	// IPMI shall contain connection details for a serial console service that uses the IPMI Serial-over-LAN (SOL)
	// protocol.
	IPMI string
	// MaxConcurrentSessions shall contain the maximum number of concurrent service sessions that this implementation
	// supports.
	MaxConcurrentSessions string
	// SSH shall contain connection details for a serial console service that uses the Secure Shell (SSH) protocol.
	SSH string
	// Telnet shall contain connection details for a serial console service that uses the Telnet protocol.
	Telnet string
}

// UnmarshalJSON unmarshals a HostSerialConsole object from the raw JSON.
func (hostserialconsole *HostSerialConsole) UnmarshalJSON(b []byte) error {
	type temp HostSerialConsole
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*hostserialconsole = HostSerialConsole(t.temp)

	// Extract the links to other entities for later

	return nil
}

// HostedServices shall describe services that a computer system supports.
type HostedServices struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// StorageServices shall contain a link to a resource collection of type HostedStorageServices.
	StorageServices string
}

// UnmarshalJSON unmarshals a HostedServices object from the raw JSON.
func (hostedservices *HostedServices) UnmarshalJSON(b []byte) error {
	type temp HostedServices
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*hostedservices = HostedServices(t.temp)

	// Extract the links to other entities for later

	return nil
}

// IdlePowerSaver shall contain the idle power saver settings of a computer system.
type IdlePowerSaver struct {
	// Enabled shall indicate if idle power saver is enabled.
	Enabled string
	// EnterDwellTimeSeconds shall contain the duration in seconds the computer system is below the
	// EnterUtilizationPercent value before the idle power save is activated.
	EnterDwellTimeSeconds int
	// EnterUtilizationPercent shall contain the percentage of utilization that the computer system shall be lower than
	// to enter idle power save.
	EnterUtilizationPercent float64
	// ExitDwellTimeSeconds shall contain the duration in seconds the computer system is above the
	// ExitUtilizationPercent value before the idle power save is stopped.
	ExitDwellTimeSeconds int
	// ExitUtilizationPercent shall contain the percentage of utilization that the computer system shall be higher than
	// to exit idle power save.
	ExitUtilizationPercent float64
}

// UnmarshalJSON unmarshals a IdlePowerSaver object from the raw JSON.
func (idlepowersaver *IdlePowerSaver) UnmarshalJSON(b []byte) error {
	type temp IdlePowerSaver
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*idlepowersaver = IdlePowerSaver(t.temp)

	// Extract the links to other entities for later

	return nil
}

// KMIPServer shall contain the KMIP server settings for a computer system.
type KMIPServer struct {
	// Address shall contain the KMIP server address.
	Address string
	// Password shall contain the password to access the KMIP server. The value shall be 'null' in responses.
	Password string
	// Port shall contain the KMIP server port.
	Port int
	// Username shall contain the username to access the KMIP server.
	Username string
}

// UnmarshalJSON unmarshals a KMIPServer object from the raw JSON.
func (kmipserver *KMIPServer) UnmarshalJSON(b []byte) error {
	type temp KMIPServer
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*kmipserver = KMIPServer(t.temp)

	// Extract the links to other entities for later

	return nil
}

// KeyManagement shall contain the key management settings of a computer system.
type KeyManagement struct {
	// KMIPCertificates shall contain a link to a resource collection of type CertificateCollection that represents the
	// server certificates for the servers referenced by the KMIPServers property.
	KMIPCertificates string
	// KMIPServers shall contain the KMIP servers to which this computer system is subscribed for key management.
	KMIPServers []KMIPServer
}

// UnmarshalJSON unmarshals a KeyManagement object from the raw JSON.
func (keymanagement *KeyManagement) UnmarshalJSON(b []byte) error {
	type temp KeyManagement
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*keymanagement = KeyManagement(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Chassis shall contain an array of links to resources of type Chassis that represent the physical containers
	// associated with this resource.
	Chassis []Chassis
	// Chassis@odata.count
	ChassisCount int `json:"Chassis@odata.count"`
	// ConsumingComputerSystems shall be an array of links to ComputerSystems that are realized, in whole or in part,
	// from this ComputerSystem.
	ConsumingComputerSystems []ComputerSystem
	// ConsumingComputerSystems@odata.count
	ConsumingComputerSystemsCount int `json:"ConsumingComputerSystems@odata.count"`
	// CooledBy shall contain an array of links to resources or objects that cool this computer system.
	CooledBy []idRef
	// CooledBy@odata.count
	CooledByCount int `json:"CooledBy@odata.count"`
	// Endpoints shall contain an array of links to resources of type Endpoint with which this system is associated.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// ManagedBy shall contain an array of link to resources of type Manager that represent the resources with
	// management responsibility for this resource.
	ManagedBy []Manager
	// ManagedBy@odata.count
	ManagedByCount int `json:"ManagedBy@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OffloadedNetworkDeviceFunctions shall contain an array of links to resources of type NetworkDeviceFunction that
	// represent the network device functions to which this system performs offload computation, such as with a
	// SmartNIC. This property shall not be present if the SystemType property does not contain 'DPU'.
	OffloadedNetworkDeviceFunctions []NetworkDeviceFunction
	// OffloadedNetworkDeviceFunctions@odata.count
	OffloadedNetworkDeviceFunctionsCount int `json:"OffloadedNetworkDeviceFunctions@odata.count"`
	// PoweredBy shall contain an array of links to resources or objects that power this computer system.
	PoweredBy []idRef
	// PoweredBy@odata.count
	PoweredByCount int `json:"PoweredBy@odata.count"`
	// ResourceBlocks shall contain an array of links to resources of type ResourceBlock that show the resource blocks
	// that are used in this computer system.
	ResourceBlocks []ResourceBlock
	// ResourceBlocks@odata.count
	ResourceBlocksCount int `json:"ResourceBlocks@odata.count"`
	// SupplyingComputerSystems shall be an array of links to ComputerSystems that contribute, in whole or in part, to
	// the implementation of this ComputerSystem.
	SupplyingComputerSystems []ComputerSystem
	// SupplyingComputerSystems@odata.count
	SupplyingComputerSystemsCount int `json:"SupplyingComputerSystems@odata.count"`
	// TrustedComponents shall contain an array of link to resources of type TrustedComponent.
	TrustedComponents []TrustedComponent
	// TrustedComponents@odata.count
	TrustedComponentsCount int `json:"TrustedComponents@odata.count"`
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

// MemorySummary shall contain properties that describe the central memory for a system.
type MemorySummary struct {
	// MemoryMirroring shall contain the ability and type of memory mirroring that this computer system supports.
	MemoryMirroring MemoryMirroring
	// Metrics shall contain a link to the metrics associated with all memory in this system.
	Metrics string
	// TotalSystemMemoryGiB shall contain the amount of configured system general purpose volatile (RAM) memory as
	// measured in gibibytes.
	TotalSystemMemoryGiB float64
	// TotalSystemPersistentMemoryGiB shall contain the total amount of configured persistent memory available to the
	// system as measured in gibibytes.
	TotalSystemPersistentMemoryGiB float64
}

// UnmarshalJSON unmarshals a MemorySummary object from the raw JSON.
func (memorysummary *MemorySummary) UnmarshalJSON(b []byte) error {
	type temp MemorySummary
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorysummary = MemorySummary(t.temp)

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

// ProcessorSummary shall contain properties that describe the central processors for a system. Processors
// described by this type shall be limited to the processors that execute system code, and shall not include
// processors used for offload functionality.
type ProcessorSummary struct {
	// CoreCount shall contain the total number of central processor cores in in the system.
	CoreCount int
	// Count shall contain the total number of physical central processors in the system.
	Count int
	// LogicalProcessorCount shall contain the total number of logical central processors in the system.
	LogicalProcessorCount int
	// Metrics shall contain a link to the metrics associated with all processors in this system.
	Metrics string
	// Model shall contain the processor model for the central processors in the system, per the description in the
	// Processor Information - Processor Family section of the SMBIOS Specification DSP0134 2.8 or later.
	Model string
	// ThreadingEnabled shall indicate that all Processor resources in this system where the ProcessorType property
	// contains 'CPU' have multiple threading support enabled.
	ThreadingEnabled string
}

// UnmarshalJSON unmarshals a ProcessorSummary object from the raw JSON.
func (processorsummary *ProcessorSummary) UnmarshalJSON(b []byte) error {
	type temp ProcessorSummary
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processorsummary = ProcessorSummary(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SerialConsoleProtocol shall describe a serial console service for a computer system.
type SerialConsoleProtocol struct {
	// ConsoleEntryCommand shall contain a command string that can be provided by a client to select or enter the
	// system's serial console, when the console is shared among several systems or a manager CLI.
	ConsoleEntryCommand string
	// HotKeySequenceDisplay shall contain a string that can be provided to a user to describe the hotkey sequence used
	// to exit the serial console session, or, if shared with a manager CLI, to return to the CLI.
	HotKeySequenceDisplay string
	// Port shall contain the port assigned to the protocol.
	Port int
	// ServiceEnabled shall indicate whether the protocol for the service is enabled.
	ServiceEnabled string
	// SharedWithManagerCLI shall indicate whether the serial console service is shared with access to the manager's
	// command-line interface (CLI).
	SharedWithManagerCLI string
}

// UnmarshalJSON unmarshals a SerialConsoleProtocol object from the raw JSON.
func (serialconsoleprotocol *SerialConsoleProtocol) UnmarshalJSON(b []byte) error {
	type temp SerialConsoleProtocol
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*serialconsoleprotocol = SerialConsoleProtocol(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TrustedModules shall describe a Trusted Module for a system.
type TrustedModules struct {
	// FirmwareVersion shall contain the firmware version as defined by the manufacturer for the Trusted Module.
	FirmwareVersion string
	// FirmwareVersion2 shall contain the 2nd firmware version, if applicable, as defined by the manufacturer for the
	// Trusted Module.
	FirmwareVersion2 string
	// InterfaceType shall contain the interface type of the installed Trusted Module.
	InterfaceType InterfaceType
	// InterfaceTypeSelection shall contain the interface type Selection method (for example to switch between TPM1_2
	// and TPM2_0) that is supported by this TrustedModule.
	InterfaceTypeSelection InterfaceTypeSelection
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a TrustedModules object from the raw JSON.
func (trustedmodules *TrustedModules) UnmarshalJSON(b []byte) error {
	type temp TrustedModules
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*trustedmodules = TrustedModules(t.temp)

	// Extract the links to other entities for later

	return nil
}

// VirtualMediaConfig shall describe a virtual media service service for a computer system.
type VirtualMediaConfig struct {
	// Port shall contain the port assigned to the service.
	Port int
	// ServiceEnabled shall indicate whether the protocol for the service is enabled.
	ServiceEnabled string
}

// UnmarshalJSON unmarshals a VirtualMediaConfig object from the raw JSON.
func (virtualmediaconfig *VirtualMediaConfig) UnmarshalJSON(b []byte) error {
	type temp VirtualMediaConfig
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*virtualmediaconfig = VirtualMediaConfig(t.temp)

	// Extract the links to other entities for later

	return nil
}

// WatchdogTimer shall contain properties that describe the host watchdog timer functionality for this
// ComputerSystem.
type WatchdogTimer struct {
	// FunctionEnabled shall indicate whether a user has enabled the host watchdog timer functionality. This property
	// indicates only that a user has enabled the timer. To activate the timer, installation of additional host-based
	// software is necessary; an update to this property does not initiate the timer.
	FunctionEnabled bool
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TimeoutAction shall contain the action to perform when the watchdog timer reaches its timeout value.
	TimeoutAction WatchdogTimeoutActions
	// WarningAction shall contain the action to perform before the watchdog timer expires. This action typically
	// occurs from three to ten seconds before to the timeout value, but the exact timing is dependent on the
	// implementation.
	WarningAction WatchdogWarningActions
}

// UnmarshalJSON unmarshals a WatchdogTimer object from the raw JSON.
func (watchdogtimer *WatchdogTimer) UnmarshalJSON(b []byte) error {
	type temp WatchdogTimer
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*watchdogtimer = WatchdogTimer(t.temp)

	// Extract the links to other entities for later

	return nil
}
