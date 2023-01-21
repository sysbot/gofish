//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// EventSeverity is
type EventSeverity string

const (
	// OKEventSeverity Informational or operating normally.
	OKEventSeverity EventSeverity = "OK"
	// WarningEventSeverity A condition that requires attention.
	WarningEventSeverity EventSeverity = "Warning"
	// CriticalEventSeverity A critical condition that requires immediate attention.
	CriticalEventSeverity EventSeverity = "Critical"
)

// LogDiagnosticDataTypes is
type LogDiagnosticDataTypes string

const (
	// ManagerLogDiagnosticDataTypes Manager diagnostic data.
	ManagerLogDiagnosticDataTypes LogDiagnosticDataTypes = "Manager"
	// PreOSLogDiagnosticDataTypes Pre-OS diagnostic data.
	PreOSLogDiagnosticDataTypes LogDiagnosticDataTypes = "PreOS"
	// OSLogDiagnosticDataTypes Operating system (OS) diagnostic data.
	OSLogDiagnosticDataTypes LogDiagnosticDataTypes = "OS"
	// OEMLogDiagnosticDataTypes OEM diagnostic data.
	OEMLogDiagnosticDataTypes LogDiagnosticDataTypes = "OEM"
	// CPERLogDiagnosticDataTypes shall indicate the data provided at the URI specified by the AdditionalDataURI
	// property is a complete UEFI Specification-defined Common Platform Error Record. The CPER data shall contain a
	// Record Header and at least one Section as defined by the UEFI Specification.
	CPERLogDiagnosticDataTypes LogDiagnosticDataTypes = "CPER"
	// CPERSectionLogDiagnosticDataTypes shall indicate the data provided at the URI specified by the AdditionalDataURI
	// property is a single Section of a UEFI Specification-defined Common Platform Error Record. The CPER data shall
	// contain one Section as defined by the UEFI Specification, with no Record Header.
	CPERSectionLogDiagnosticDataTypes LogDiagnosticDataTypes = "CPERSection"
)

// LogEntryCode is
type LogEntryCode string

const (
	// AssertLogEntryCode The condition has been asserted.
	AssertLogEntryCode LogEntryCode = "Assert"
	// DeassertLogEntryCode The condition has been deasserted.
	DeassertLogEntryCode LogEntryCode = "Deassert"
	// LowerNoncriticalgoinglowLogEntryCode The reading crossed the Lower Non-critical threshold while going low.
	LowerNoncriticalgoinglowLogEntryCode LogEntryCode = "Lower Non-critical - going low"
	// LowerNoncriticalgoinghighLogEntryCode The reading crossed the Lower Non-critical threshold while going high.
	LowerNoncriticalgoinghighLogEntryCode LogEntryCode = "Lower Non-critical - going high"
	// LowerCriticalgoinglowLogEntryCode The reading crossed the Lower Critical threshold while going low.
	LowerCriticalgoinglowLogEntryCode LogEntryCode = "Lower Critical - going low"
	// LowerCriticalgoinghighLogEntryCode The reading crossed the Lower Critical threshold while going high.
	LowerCriticalgoinghighLogEntryCode LogEntryCode = "Lower Critical - going high"
	// LowerNonrecoverablegoinglowLogEntryCode The reading crossed the Lower Non-recoverable threshold while going low.
	LowerNonrecoverablegoinglowLogEntryCode LogEntryCode = "Lower Non-recoverable - going low"
	// LowerNonrecoverablegoinghighLogEntryCode The reading crossed the Lower Non-recoverable threshold while going
	// high.
	LowerNonrecoverablegoinghighLogEntryCode LogEntryCode = "Lower Non-recoverable - going high"
	// UpperNoncriticalgoinglowLogEntryCode The reading crossed the Upper Non-critical threshold while going low.
	UpperNoncriticalgoinglowLogEntryCode LogEntryCode = "Upper Non-critical - going low"
	// UpperNoncriticalgoinghighLogEntryCode The reading crossed the Upper Non-critical threshold while going high.
	UpperNoncriticalgoinghighLogEntryCode LogEntryCode = "Upper Non-critical - going high"
	// UpperCriticalgoinglowLogEntryCode The reading crossed the Upper Critical threshold while going low.
	UpperCriticalgoinglowLogEntryCode LogEntryCode = "Upper Critical - going low"
	// UpperCriticalgoinghighLogEntryCode The reading crossed the Upper Critical threshold while going high.
	UpperCriticalgoinghighLogEntryCode LogEntryCode = "Upper Critical - going high"
	// UpperNonrecoverablegoinglowLogEntryCode The reading crossed the Upper Non-recoverable threshold while going low.
	UpperNonrecoverablegoinglowLogEntryCode LogEntryCode = "Upper Non-recoverable - going low"
	// UpperNonrecoverablegoinghighLogEntryCode The reading crossed the Upper Non-recoverable threshold while going
	// high.
	UpperNonrecoverablegoinghighLogEntryCode LogEntryCode = "Upper Non-recoverable - going high"
	// TransitiontoIdleLogEntryCode The state transitioned to idle.
	TransitiontoIdleLogEntryCode LogEntryCode = "Transition to Idle"
	// TransitiontoActiveLogEntryCode The state transitioned to active.
	TransitiontoActiveLogEntryCode LogEntryCode = "Transition to Active"
	// TransitiontoBusyLogEntryCode The state transitioned to busy.
	TransitiontoBusyLogEntryCode LogEntryCode = "Transition to Busy"
	// StateDeassertedLogEntryCode The state has been deasserted.
	StateDeassertedLogEntryCode LogEntryCode = "State Deasserted"
	// StateAssertedLogEntryCode The state has been asserted.
	StateAssertedLogEntryCode LogEntryCode = "State Asserted"
	// PredictiveFailuredeassertedLogEntryCode A Predictive Failure is no longer present.
	PredictiveFailuredeassertedLogEntryCode LogEntryCode = "Predictive Failure deasserted"
	// PredictiveFailureassertedLogEntryCode A Predictive Failure has been detected.
	PredictiveFailureassertedLogEntryCode LogEntryCode = "Predictive Failure asserted"
	// LimitNotExceededLogEntryCode A limit has not been exceeded.
	LimitNotExceededLogEntryCode LogEntryCode = "Limit Not Exceeded"
	// LimitExceededLogEntryCode A limit has been exceeded.
	LimitExceededLogEntryCode LogEntryCode = "Limit Exceeded"
	// PerformanceMetLogEntryCode Performance meets expectations.
	PerformanceMetLogEntryCode LogEntryCode = "Performance Met"
	// PerformanceLagsLogEntryCode Performance does not meet expectations.
	PerformanceLagsLogEntryCode LogEntryCode = "Performance Lags"
	// TransitiontoOKLogEntryCode A state has changed to OK.
	TransitiontoOKLogEntryCode LogEntryCode = "Transition to OK"
	// TransitiontoNonCriticalfromOKLogEntryCode A state has changed to Non-Critical from OK.
	TransitiontoNonCriticalfromOKLogEntryCode LogEntryCode = "Transition to Non-Critical from OK"
	// TransitiontoCriticalfromlesssevereLogEntryCode A state has changed to Critical from less severe.
	TransitiontoCriticalfromlesssevereLogEntryCode LogEntryCode = "Transition to Critical from less severe"
	// TransitiontoNonrecoverablefromlesssevereLogEntryCode A state has changed to Non-recoverable from less severe.
	TransitiontoNonrecoverablefromlesssevereLogEntryCode LogEntryCode = "Transition to Non-recoverable from less severe"
	// TransitiontoNonCriticalfrommoresevereLogEntryCode A state has changed to Non-Critical from more severe.
	TransitiontoNonCriticalfrommoresevereLogEntryCode LogEntryCode = "Transition to Non-Critical from more severe"
	// TransitiontoCriticalfromNonrecoverableLogEntryCode A state has changed to Critical from Non-recoverable.
	TransitiontoCriticalfromNonrecoverableLogEntryCode LogEntryCode = "Transition to Critical from Non-recoverable"
	// TransitiontoNonrecoverableLogEntryCode A state has changed to Non-recoverable.
	TransitiontoNonrecoverableLogEntryCode LogEntryCode = "Transition to Non-recoverable"
	// MonitorLogEntryCode A monitor event.
	MonitorLogEntryCode LogEntryCode = "Monitor"
	// InformationalLogEntryCode An informational event.
	InformationalLogEntryCode LogEntryCode = "Informational"
	// DeviceRemovedDeviceAbsentLogEntryCode A device has been removed or is absent.
	DeviceRemovedDeviceAbsentLogEntryCode LogEntryCode = "Device Removed / Device Absent"
	// DeviceInsertedDevicePresentLogEntryCode A device has been inserted or is present.
	DeviceInsertedDevicePresentLogEntryCode LogEntryCode = "Device Inserted / Device Present"
	// DeviceDisabledLogEntryCode A device has been disabled.
	DeviceDisabledLogEntryCode LogEntryCode = "Device Disabled"
	// DeviceEnabledLogEntryCode A device has been enabled.
	DeviceEnabledLogEntryCode LogEntryCode = "Device Enabled"
	// TransitiontoRunningLogEntryCode A state has transitioned to Running.
	TransitiontoRunningLogEntryCode LogEntryCode = "Transition to Running"
	// TransitiontoInTestLogEntryCode A state has transitioned to In Test.
	TransitiontoInTestLogEntryCode LogEntryCode = "Transition to In Test"
	// TransitiontoPowerOffLogEntryCode A state has transitioned to Power Off.
	TransitiontoPowerOffLogEntryCode LogEntryCode = "Transition to Power Off"
	// TransitiontoOnLineLogEntryCode A state has transitioned to On Line.
	TransitiontoOnLineLogEntryCode LogEntryCode = "Transition to On Line"
	// TransitiontoOffLineLogEntryCode A state has transitioned to Off Line.
	TransitiontoOffLineLogEntryCode LogEntryCode = "Transition to Off Line"
	// TransitiontoOffDutyLogEntryCode A state has transitioned to Off Duty.
	TransitiontoOffDutyLogEntryCode LogEntryCode = "Transition to Off Duty"
	// TransitiontoDegradedLogEntryCode A state has transitioned to Degraded.
	TransitiontoDegradedLogEntryCode LogEntryCode = "Transition to Degraded"
	// TransitiontoPowerSaveLogEntryCode A state has transitioned to Power Save.
	TransitiontoPowerSaveLogEntryCode LogEntryCode = "Transition to Power Save"
	// InstallErrorLogEntryCode An install error has been detected.
	InstallErrorLogEntryCode LogEntryCode = "Install Error"
	// FullyRedundantLogEntryCode Indicates that full redundancy has been regained.
	FullyRedundantLogEntryCode LogEntryCode = "Fully Redundant"
	// RedundancyLostLogEntryCode Entered any non-redundant state, including Non-redundant: Insufficient Resources.
	RedundancyLostLogEntryCode LogEntryCode = "Redundancy Lost"
	// RedundancyDegradedLogEntryCode Redundancy still exists, but at less than full level.
	RedundancyDegradedLogEntryCode LogEntryCode = "Redundancy Degraded"
	// NonredundantSufficientResourcesfromRedundantLogEntryCode Redundancy has been lost but unit is functioning with
	// minimum resources needed for normal operation.
	NonredundantSufficientResourcesfromRedundantLogEntryCode LogEntryCode = "Non-redundant:Sufficient Resources from Redundant"
	// NonredundantSufficientResourcesfromInsufficientResourcesLogEntryCode Unit has regained minimum resources needed
	// for normal operation.
	NonredundantSufficientResourcesfromInsufficientResourcesLogEntryCode LogEntryCode = "Non-redundant:Sufficient Resources from Insufficient Resources"
	// NonredundantInsufficientResourcesLogEntryCode Unit is non-redundant and has insufficient resources to maintain
	// normal operation.
	NonredundantInsufficientResourcesLogEntryCode LogEntryCode = "Non-redundant:Insufficient Resources"
	// RedundancyDegradedfromFullyRedundantLogEntryCode Unit has lost some redundant resource(s) but is still in a
	// redundant state.
	RedundancyDegradedfromFullyRedundantLogEntryCode LogEntryCode = "Redundancy Degraded from Fully Redundant"
	// RedundancyDegradedfromNonredundantLogEntryCode Unit has regained some resource(s) and is redundant but not fully
	// redundant.
	RedundancyDegradedfromNonredundantLogEntryCode LogEntryCode = "Redundancy Degraded from Non-redundant"
	// D0PowerStateLogEntryCode The ACPI-defined D0 power state.
	D0PowerStateLogEntryCode LogEntryCode = "D0 Power State"
	// D1PowerStateLogEntryCode The ACPI-defined D1 power state.
	D1PowerStateLogEntryCode LogEntryCode = "D1 Power State"
	// D2PowerStateLogEntryCode The ACPI-defined D2 power state.
	D2PowerStateLogEntryCode LogEntryCode = "D2 Power State"
	// D3PowerStateLogEntryCode The ACPI-defined D3 power state.
	D3PowerStateLogEntryCode LogEntryCode = "D3 Power State"
	// OEMLogEntryCode An OEM-defined event.
	OEMLogEntryCode LogEntryCode = "OEM"
)

// LogEntryType is
type LogEntryType string

const (
	// EventLogEntryType A Redfish-defined message.
	EventLogEntryType LogEntryType = "Event"
	// SELLogEntryType A legacy IPMI System Event Log (SEL) entry.
	SELLogEntryType LogEntryType = "SEL"
	// OemLogEntryType An entry in an OEM-defined format.
	OemLogEntryType LogEntryType = "Oem"
)

// OriginatorTypes is
type OriginatorTypes string

const (
	// ClientOriginatorTypes A client of the service created this log entry.
	ClientOriginatorTypes OriginatorTypes = "Client"
	// InternalOriginatorTypes A process running on the service created this log entry.
	InternalOriginatorTypes OriginatorTypes = "Internal"
	// SupportingServiceOriginatorTypes A process not running on the service but running on a supporting service, such
	// as RDE implementations, UEFI, or host processes, created this log entry.
	SupportingServiceOriginatorTypes OriginatorTypes = "SupportingService"
)

// SensorType is
type SensorType string

const (
	// PlatformSecurityViolationAttemptSensorType A platform security sensor.
	PlatformSecurityViolationAttemptSensorType SensorType = "Platform Security Violation Attempt"
	// TemperatureSensorType A temperature sensor.
	TemperatureSensorType SensorType = "Temperature"
	// VoltageSensorType A voltage sensor.
	VoltageSensorType SensorType = "Voltage"
	// CurrentSensorType A current sensor.
	CurrentSensorType SensorType = "Current"
	// FanSensorType A fan sensor.
	FanSensorType SensorType = "Fan"
	// PhysicalChassisSecuritySensorType A physical security sensor.
	PhysicalChassisSecuritySensorType SensorType = "Physical Chassis Security"
	// ProcessorSensorType A sensor for a processor.
	ProcessorSensorType SensorType = "Processor"
	// PowerSupplyConverterSensorType A sensor for a power supply or DC-to-DC converter.
	PowerSupplyConverterSensorType SensorType = "Power Supply / Converter"
	// PowerUnitSensorType A sensor for a power unit.
	PowerUnitSensorType SensorType = "PowerUnit"
	// CoolingDeviceSensorType A sensor for a cooling device.
	CoolingDeviceSensorType SensorType = "CoolingDevice"
	// OtherUnitsbasedSensorSensorType A sensor for a miscellaneous analog sensor.
	OtherUnitsbasedSensorSensorType SensorType = "Other Units-based Sensor"
	// MemorySensorType A sensor for a memory device.
	MemorySensorType SensorType = "Memory"
	// DriveSlotBaySensorType A sensor for a drive slot or bay.
	DriveSlotBaySensorType SensorType = "Drive Slot/Bay"
	// POSTMemoryResizeSensorType A sensor for a POST memory resize event.
	POSTMemoryResizeSensorType SensorType = "POST Memory Resize"
	// SystemFirmwareProgressSensorType A sensor for a system firmware progress event.
	SystemFirmwareProgressSensorType SensorType = "System Firmware Progress"
	// EventLoggingDisabledSensorType A sensor for the event log.
	EventLoggingDisabledSensorType SensorType = "Event Logging Disabled"
	// SystemEventSensorType A sensor for a system event.
	SystemEventSensorType SensorType = "System Event"
	// CriticalInterruptSensorType A sensor for a critical interrupt event.
	CriticalInterruptSensorType SensorType = "Critical Interrupt"
	// ButtonSwitchSensorType A sensor for a button or switch.
	ButtonSwitchSensorType SensorType = "Button/Switch"
	// ModuleBoardSensorType A sensor for a module or board.
	ModuleBoardSensorType SensorType = "Module/Board"
	// MicrocontrollerCoprocessorSensorType A sensor for a microcontroller or coprocessor.
	MicrocontrollerCoprocessorSensorType SensorType = "Microcontroller/Coprocessor"
	// AddinCardSensorType A sensor for an add-in card.
	AddinCardSensorType SensorType = "Add-in Card"
	// ChassisSensorType A sensor for a chassis.
	ChassisSensorType SensorType = "Chassis"
	// ChipSetSensorType A sensor for a chipset.
	ChipSetSensorType SensorType = "ChipSet"
	// OtherFRUSensorType A sensor for another type of FRU.
	OtherFRUSensorType SensorType = "Other FRU"
	// CableInterconnectSensorType A sensor for a cable or interconnect device type.
	CableInterconnectSensorType SensorType = "Cable/Interconnect"
	// TerminatorSensorType A sensor for a terminator.
	TerminatorSensorType SensorType = "Terminator"
	// SystemBootRestartSensorType A sensor for a system boot or restart event.
	SystemBootRestartSensorType SensorType = "SystemBoot/Restart"
	// BootErrorSensorType A sensor for a boot error event.
	BootErrorSensorType SensorType = "Boot Error"
	// BaseOSBootInstallationStatusSensorType A sensor for a base OS boot or installation status event.
	BaseOSBootInstallationStatusSensorType SensorType = "BaseOSBoot/InstallationStatus"
	// OSStopShutdownSensorType A sensor for an OS stop or shutdown event
	OSStopShutdownSensorType SensorType = "OS Stop/Shutdown"
	// SlotConnectorSensorType A sensor for a slot or connector.
	SlotConnectorSensorType SensorType = "Slot/Connector"
	// SystemACPIPowerStateSensorType A sensor for an ACPI power state event.
	SystemACPIPowerStateSensorType SensorType = "System ACPI PowerState"
	// WatchdogSensorType A sensor for a watchdog event.
	WatchdogSensorType SensorType = "Watchdog"
	// PlatformAlertSensorType A sensor for a platform alert event.
	PlatformAlertSensorType SensorType = "Platform Alert"
	// EntityPresenceSensorType A sensor for an entity presence event.
	EntityPresenceSensorType SensorType = "Entity Presence"
	// MonitorASICICSensorType A sensor for a monitor ASIC or IC.
	MonitorASICICSensorType SensorType = "Monitor ASIC/IC"
	// LANSensorType A sensor for a LAN device.
	LANSensorType SensorType = "LAN"
	// ManagementSubsystemHealthSensorType A sensor for a management subsystem health event.
	ManagementSubsystemHealthSensorType SensorType = "Management Subsystem Health"
	// BatterySensorType A sensor for a battery.
	BatterySensorType SensorType = "Battery"
	// SessionAuditSensorType A sensor for a session audit event.
	SessionAuditSensorType SensorType = "Session Audit"
	// VersionChangeSensorType A sensor for a version change event.
	VersionChangeSensorType SensorType = "Version Change"
	// FRUStateSensorType A sensor for a FRU state event.
	FRUStateSensorType SensorType = "FRUState"
	// OEMSensorType An OEM-defined sensor.
	OEMSensorType SensorType = "OEM"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OriginOfCondition shall contain a link to the resource that caused the log entry. For log entries that represent
	// the creation or deletion of a resource, this property should reference the created or deleted resource and not
	// the collection that contains the resource.
	OriginOfCondition string
	// RelatedItem shall contain an array of links to resources that are related to this log entry. It shall not
	// contain links to LogEntry resources. RelatedLogEntries is used to reference related log entries. This property
	// shall not contain the value of the OriginOfCondition property.
	RelatedItem []idRef
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// RelatedLogEntries shall contain an array of links to resources of type LogEntry in this or other log services
	// that are related to this log entry.
	RelatedLogEntries []LogEntry
	// RelatedLogEntries@odata.count
	RelatedLogEntriesCount int `json:"RelatedLogEntries@odata.count"`
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

// LogEntry shall represent the log format for log services in a Redfish implementation.
type LogEntry struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AdditionalDataSizeBytes shall contain the size of the additional data referenced by the AdditionalDataURI
	// property for the log entry.
	AdditionalDataSizeBytes int
	// AdditionalDataURI shall contain the URI at which to access the additional data for the log entry, using the
	// Redfish protocol and authentication methods.
	AdditionalDataURI string
	// Created shall contain the date and time when the log entry was created.
	Created string
	// Description provides a description of this resource.
	Description string
	// DiagnosticDataType shall contain the type of diagnostic data contained at the URI referenced by
	// AdditionalDataURI.
	DiagnosticDataType LogDiagnosticDataTypes
	// EntryCode shall contain the entry code for the log entry if the EntryType is 'SEL'. Tables 42-1 and 42-2 of the
	// IPMI Specification v2.0 revision 1.1 describe these enumerations.
	EntryCode LogEntryCode
	// EntryType shall represent the type of log entry. If the resource represents an IPMI SEL entry, the value shall
	// contain 'SEL'. If the resource represents a Redfish event log entry, the value shall contain 'Event'. If the
	// resource represents an OEM log entry format, the value shall contain 'Oem'.
	EntryType string
	// EventGroupId shall indicate that events are related and shall have the same value in the case where multiple
	// event messages are produced by the same root cause. Implementations shall use separate values for events with
	// separate root cause. There shall not be ordering of events implied by this property's value.
	EventGroupId int
	// EventId shall indicate a unique identifier for the event, the format of which is implementation dependent.
	EventId string
	// EventTimestamp shall contain the date and time when the event occurred.
	EventTimestamp string
	// GeneratorId shall contain the 'Generator ID' field of the IPMI SEL Event Record. If EntryType is not 'SEL', this
	// property should not be present.
	GeneratorId string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Message shall contain the message of the log entry. This property decodes from the entry type. If the entry type
	// is 'Event', this property contains a message. If the entry type is 'SEL', this property contains an SEL-specific
	// message, following the format specified in Table 32-1, SEL Event Records, in the IPMI Specification v2.0
	// revision 1.1. Otherwise, this property contains an OEM-specific log entry. In most cases, this property contains
	// the actual log entry.
	Message string
	// MessageArgs shall contain message arguments to substitute into the included or looked-up message.
	MessageArgs []string
	// MessageId shall contain the MessageId, event data, or OEM-specific information. This property decodes from the
	// entry type. If the entry type is 'Event', this property contains a Redfish Specification-defined MessageId
	// property of the event. If the entry type is 'SEL', the format should follow the pattern
	// '^0[xX](([a-fA-F]|[0-9]){2}){4}$', which results in a string in the form '0xNNaabbcc', where 'NN' is the
	// EventDir/EventType byte, 'aa' is the Event Data 1 byte, 'bb' is Event Data 2 byte, 'cc' is Event Data 3 byte,
	// corresponding with bytes 13-16 in the IPMI SEL Event Record. Otherwise, this property contains OEM-specific
	// information.
	MessageId string
	// Modified shall contain the date and time when the log entry was last modified. This property shall not appear if
	// the log entry has not been modified since it was created.
	Modified string
	// OEMDiagnosticDataType shall contain the OEM-defined type of diagnostic data contained at the URI referenced by
	// AdditionalDataURI. This property shall be present if DiagnosticDataType is 'OEM'.
	OEMDiagnosticDataType string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OemLogEntryCode shall represent the OEM-specific Log Entry Code type of the Entry. This property shall only be
	// present if EntryType is 'SEL' and LogEntryCode is 'OEM'.
	OemLogEntryCode string
	// OemRecordFormat shall represent the OEM-specific format of the entry. This property shall be required if the
	// EntryType value is 'Oem'.
	OemRecordFormat string
	// OemSensorType shall represent the OEM-specific sensor type of the entry. This property shall only be used if
	// EntryType is 'SEL' and SensorType is 'OEM'.
	OemSensorType string
	// Originator shall contain the source of the log entry.
	Originator string
	// OriginatorType shall contain the type of originator data.
	OriginatorType string
	// Resolution shall contain the resolution of the log entry. Services should replace the resolution defined in the
	// message registry with a more specific resolution in a log entry.
	Resolution string
	// Resolved shall contain an indication if the cause of the log entry has been resolved or repaired. The value
	// 'true' shall indicate if the cause of the log entry has been resolved or repaired. This property shall contain
	// the value 'false' if the log entry is still active. The value 'false' shall be the initial state.
	Resolved bool
	// SensorNumber shall contain the IPMI sensor number if the value of the EntryType property is 'SEL'. This property
	// should not appear in the resource for other values of EntryType.
	SensorNumber int
	// SensorType shall contain the sensor type to which the log entry pertains if the entry type is 'SEL'. Table 42-3,
	// Sensor Type Codes, in the IPMI Specification v2.0 revision 1.1 describes these enumerations.
	SensorType SensorType
	// ServiceProviderNotified shall contain an indication if the log entry has been sent to the service provider.
	ServiceProviderNotified bool
	// Severity shall contain the severity of the condition that created the log entry, as defined in the Status
	// section of the Redfish Specification.
	Severity EventSeverity
	// SpecificEventExistsInGroup shall indicate that this log entry is equivalent to another log entry, with a more
	// specific definition, within the same EventGroupId. For example, the 'DriveFailed' message from the Storage
	// Device Message Registry is more specific than the 'ResourceStatusChangedCritical' message from the Resource
	// Event Message Registry, when both occur with the same EventGroupId. This property shall contain 'true' if a more
	// specific event is available, and shall contain 'false' if no equivalent event exists in the same EventGroupId.
	// If this property is absent, the value shall be assumed to be 'false'.
	SpecificEventExistsInGroup string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a LogEntry object from the raw JSON.
func (logentry *LogEntry) UnmarshalJSON(b []byte) error {
	type temp LogEntry
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*logentry = LogEntry(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	logentry.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (logentry *LogEntry) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(LogEntry)
	original.UnmarshalJSON(logentry.rawData)

	readWriteFields := []string{
		"Resolved",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(logentry).Elem()

	return logentry.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetLogEntry will get a LogEntry instance from the service.
func GetLogEntry(c common.Client, uri string) (*LogEntry, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var logentry LogEntry
	err = json.NewDecoder(resp.Body).Decode(&logentry)
	if err != nil {
		return nil, err
	}

	logentry.SetClient(c)
	return &logentry, nil
}

// ListReferencedLogEntrys gets the collection of LogEntry from
// a provided reference.
func ListReferencedLogEntrys(c common.Client, link string) ([]*LogEntry, error) {
	var result []*LogEntry
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, logentryLink := range links.ItemLinks {
		logentry, err := GetLogEntry(c, logentryLink)
		if err != nil {
			collectionError.Failures[logentryLink] = err
		} else {
			result = append(result, logentry)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
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
