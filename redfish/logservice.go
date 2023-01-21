//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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
)

// LogEntryTypes is
type LogEntryTypes string

const (
	// EventLogEntryTypes The log contains Redfish-defined messages.
	EventLogEntryTypes LogEntryTypes = "Event"
	// SELLogEntryTypes The log contains legacy IPMI System Event Log (SEL) entries.
	SELLogEntryTypes LogEntryTypes = "SEL"
	// MultipleLogEntryTypes The log contains multiple log entry types and, therefore, the log service cannot guarantee
	// a single entry type.
	MultipleLogEntryTypes LogEntryTypes = "Multiple"
	// OEMLogEntryTypes The log contains entries in an OEM-defined format.
	OEMLogEntryTypes LogEntryTypes = "OEM"
)

// OverWritePolicy is
type OverWritePolicy string

const (
	// UnknownOverWritePolicy The overwrite policy is not known or is undefined.
	UnknownOverWritePolicy OverWritePolicy = "Unknown"
	// WrapsWhenFullOverWritePolicy When full, new entries to the log overwrite earlier entries.
	WrapsWhenFullOverWritePolicy OverWritePolicy = "WrapsWhenFull"
	// NeverOverWritesOverWritePolicy When full, new entries to the log are discarded.
	NeverOverWritesOverWritePolicy OverWritePolicy = "NeverOverWrites"
)

// SyslogFacility is This type shall specify the syslog facility codes as program types. Facility values are
// described in the RFC5424.
type SyslogFacility string

const (
	// KernSyslogFacility Kernel messages.
	KernSyslogFacility SyslogFacility = "Kern"
	// UserSyslogFacility User-level messages.
	UserSyslogFacility SyslogFacility = "User"
	// MailSyslogFacility Mail system.
	MailSyslogFacility SyslogFacility = "Mail"
	// DaemonSyslogFacility System daemons.
	DaemonSyslogFacility SyslogFacility = "Daemon"
	// AuthSyslogFacility Security/authentication messages.
	AuthSyslogFacility SyslogFacility = "Auth"
	// SyslogSyslogFacility Messages generated internally by syslogd.
	SyslogSyslogFacility SyslogFacility = "Syslog"
	// LPRSyslogFacility Line printer subsystem.
	LPRSyslogFacility SyslogFacility = "LPR"
	// NewsSyslogFacility Network news subsystem.
	NewsSyslogFacility SyslogFacility = "News"
	// UUCPSyslogFacility UUCP subsystem.
	UUCPSyslogFacility SyslogFacility = "UUCP"
	// CronSyslogFacility Clock daemon.
	CronSyslogFacility SyslogFacility = "Cron"
	// AuthprivSyslogFacility Security/authentication messages.
	AuthprivSyslogFacility SyslogFacility = "Authpriv"
	// FTPSyslogFacility FTP daemon.
	FTPSyslogFacility SyslogFacility = "FTP"
	// NTPSyslogFacility NTP subsystem.
	NTPSyslogFacility SyslogFacility = "NTP"
	// SecuritySyslogFacility Log audit.
	SecuritySyslogFacility SyslogFacility = "Security"
	// ConsoleSyslogFacility Log alert.
	ConsoleSyslogFacility SyslogFacility = "Console"
	// SolarisCronSyslogFacility Scheduling daemon.
	SolarisCronSyslogFacility SyslogFacility = "SolarisCron"
	// Local0SyslogFacility Locally used facility 0.
	Local0SyslogFacility SyslogFacility = "Local0"
	// Local1SyslogFacility Locally used facility 1.
	Local1SyslogFacility SyslogFacility = "Local1"
	// Local2SyslogFacility Locally used facility 2.
	Local2SyslogFacility SyslogFacility = "Local2"
	// Local3SyslogFacility Locally used facility 3.
	Local3SyslogFacility SyslogFacility = "Local3"
	// Local4SyslogFacility Locally used facility 4.
	Local4SyslogFacility SyslogFacility = "Local4"
	// Local5SyslogFacility Locally used facility 5.
	Local5SyslogFacility SyslogFacility = "Local5"
	// Local6SyslogFacility Locally used facility 6.
	Local6SyslogFacility SyslogFacility = "Local6"
	// Local7SyslogFacility Locally used facility 7.
	Local7SyslogFacility SyslogFacility = "Local7"
)

// SyslogSeverity is This type shall specify the syslog severity levels as an application-specific rating used to
// describe the urgency of the message. 'Emergency' should be reserved for messages indicating the system is
// unusable and 'Debug' should only be used when debugging a program. Severity values are described in RFC5424.
type SyslogSeverity string

const (
	// EmergencySyslogSeverity A panic condition.
	EmergencySyslogSeverity SyslogSeverity = "Emergency"
	// AlertSyslogSeverity A condition that should be corrected immediately, such as a corrupted system database.
	AlertSyslogSeverity SyslogSeverity = "Alert"
	// CriticalSyslogSeverity Hard device errors.
	CriticalSyslogSeverity SyslogSeverity = "Critical"
	// ErrorSyslogSeverity An Error.
	ErrorSyslogSeverity SyslogSeverity = "Error"
	// WarningSyslogSeverity A Warning.
	WarningSyslogSeverity SyslogSeverity = "Warning"
	// NoticeSyslogSeverity Conditions that are not error conditions, but that may require special handling.
	NoticeSyslogSeverity SyslogSeverity = "Notice"
	// InformationalSyslogSeverity Informational only.
	InformationalSyslogSeverity SyslogSeverity = "Informational"
	// DebugSyslogSeverity Messages that contain information normally of use only when debugging a program.
	DebugSyslogSeverity SyslogSeverity = "Debug"
	// AllSyslogSeverity A message of any severity.
	AllSyslogSeverity SyslogSeverity = "All"
)

// LogService shall represent a log service for a Redfish implementation. When the Id property contains
// 'DeviceLog', the log shall contain log entries that migrate with the device.
type LogService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AutoDSTEnabled shall indicate whether the log service is configured for automatic Daylight Saving Time (DST)
	// adjustment. DST adjustment shall not modify the timestamp of existing log entries.
	AutoDSTEnabled string
	// DateTime shall contain the current date and time with UTC offset of the log service.
	DateTime string
	// DateTimeLocalOffset shall contain the offset from UTC time that the DateTime property contains. If both DateTime
	// and DateTimeLocalOffset are provided in modification requests, services shall apply DateTimeLocalOffset after
	// DateTime is applied.
	DateTimeLocalOffset string
	// Description provides a description of this resource.
	Description string
	// Entries shall contain a link to a resource collection of type LogEntryCollection.
	Entries string
	// LogEntryType shall contain the value for the EntryType property of all LogEntry resources contained in the
	// LogEntryCollection resource for this log service. If the service cannot determine or guarantee a single
	// EntryType value for all LogEntry resources, this property shall contain the value 'Multiple'.
	LogEntryType LogEntryTypes
	// MaxNumberOfRecords shall contain the maximum number of LogEntry resources in the LogEntryCollection resource for
	// this service.
	MaxNumberOfRecords string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OverWritePolicy shall indicate the policy of the log service when the MaxNumberOfRecords has been reached.
	OverWritePolicy string
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SyslogFilters shall describe all desired syslog messages to be logged locally. If this property contains an
	// empty array, all messages shall be logged.
	SyslogFilters []SyslogFilter
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a LogService object from the raw JSON.
func (logservice *LogService) UnmarshalJSON(b []byte) error {
	type temp LogService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*logservice = LogService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	logservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (logservice *LogService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(LogService)
	original.UnmarshalJSON(logservice.rawData)

	readWriteFields := []string{
		"AutoDSTEnabled",
		"DateTime",
		"DateTimeLocalOffset",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(logservice).Elem()

	return logservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetLogService will get a LogService instance from the service.
func GetLogService(c common.Client, uri string) (*LogService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var logservice LogService
	err = json.NewDecoder(resp.Body).Decode(&logservice)
	if err != nil {
		return nil, err
	}

	logservice.SetClient(c)
	return &logservice, nil
}

// ListReferencedLogServices gets the collection of LogService from
// a provided reference.
func ListReferencedLogServices(c common.Client, link string) ([]*LogService, error) {
	var result []*LogService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, logserviceLink := range links.ItemLinks {
		logservice, err := GetLogService(c, logserviceLink)
		if err != nil {
			collectionError.Failures[logserviceLink] = err
		} else {
			result = append(result, logservice)
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

// SyslogFilter shall contain the filter for a syslog message. The filter shall describe the desired syslog message
// to be enabled locally.
type SyslogFilter struct {
	// LogFacilities shall contain the types of programs that can log messages. If this property contains an empty
	// array or is absent, all facilities shall be indicated.
	LogFacilities []SyslogFacility
	// LowestSeverity shall contain the lowest syslog severity level that will be logged. The service shall log all
	// messages equal to or greater than the value in this property. The value 'All' shall indicate all severities.
	LowestSeverity SyslogSeverity
}

// UnmarshalJSON unmarshals a SyslogFilter object from the raw JSON.
func (syslogfilter *SyslogFilter) UnmarshalJSON(b []byte) error {
	type temp SyslogFilter
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*syslogfilter = SyslogFilter(t.temp)

	// Extract the links to other entities for later

	return nil
}
