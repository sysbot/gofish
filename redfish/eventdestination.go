//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DeliveryRetryPolicy is
type DeliveryRetryPolicy string

const (
	// TerminateAfterRetriesDeliveryRetryPolicy shall indicate the subscription is terminated after the maximum number
	// of retries is reached, specified by the DeliveryRetryAttempts property in the event service.
	TerminateAfterRetriesDeliveryRetryPolicy DeliveryRetryPolicy = "TerminateAfterRetries"
	// SuspendRetriesDeliveryRetryPolicy shall indicate the subscription is suspended after the maximum number of
	// retries is reached, specified by the DeliveryRetryAttempts property in the event service. The value of the State
	// property within Status shall contain 'Disabled' for a suspended subscription.
	SuspendRetriesDeliveryRetryPolicy DeliveryRetryPolicy = "SuspendRetries"
	// RetryForeverDeliveryRetryPolicy shall indicate the subscription is not suspended or terminated, and attempts at
	// delivery of future events shall continue regardless of the number of retries. The interval between retries
	// remains constant and is specified by the DeliveryRetryIntervalSeconds property in the event service.
	RetryForeverDeliveryRetryPolicy DeliveryRetryPolicy = "RetryForever"
	// RetryForeverWithBackoffDeliveryRetryPolicy shall indicate the subscription is not suspended or terminated, and
	// attempts at delivery of future events shall continue regardless of the number of retries. Retry attempts are
	// issued over time according to a service-defined backoff algorithm. The backoff algorithm may insert an
	// increasing amount of delay between retry attempts and may reach a maximum.
	RetryForeverWithBackoffDeliveryRetryPolicy DeliveryRetryPolicy = "RetryForeverWithBackoff"
)

// EventDestinationProtocol is
type EventDestinationProtocol string

const (
	// RedfishEventDestinationProtocol shall indicate the destination follows the Redfish Specification for event
	// notifications. Destinations requesting EventFormatType of 'Event' shall receive a Redfish resource of type
	// Event. Destinations requesting EventFormatType of 'MetricReport' shall receive a Redfish resource of type
	// MetricReport.
	RedfishEventDestinationProtocol EventDestinationProtocol = "Redfish"
	// SNMPv1EventDestinationProtocol shall indicate the destination follows the RFC1157-defined SNMPv1 protocol.
	SNMPv1EventDestinationProtocol EventDestinationProtocol = "SNMPv1"
	// SNMPv2cEventDestinationProtocol shall indicate the destination follows the SNMPv2c protocol as defined by
	// RFC1441 and RFC1452.
	SNMPv2cEventDestinationProtocol EventDestinationProtocol = "SNMPv2c"
	// SNMPv3EventDestinationProtocol shall indicate the destination follows the SNMPv3 protocol as defined by RFC3411
	// and RFC3418.
	SNMPv3EventDestinationProtocol EventDestinationProtocol = "SNMPv3"
	// SMTPEventDestinationProtocol shall indicate the destination follows the RFC5321-defined SMTP specification.
	SMTPEventDestinationProtocol EventDestinationProtocol = "SMTP"
	// SyslogTLSEventDestinationProtocol shall indicate the destination follows the TLS-based transport for syslog as
	// defined in RFC5424.
	SyslogTLSEventDestinationProtocol EventDestinationProtocol = "SyslogTLS"
	// SyslogTCPEventDestinationProtocol shall indicate the destination follows the TCP-based transport for syslog as
	// defined in RFC6587.
	SyslogTCPEventDestinationProtocol EventDestinationProtocol = "SyslogTCP"
	// SyslogUDPEventDestinationProtocol shall indicate the destination follows the UDP-based transport for syslog as
	// defined in RFC5424.
	SyslogUDPEventDestinationProtocol EventDestinationProtocol = "SyslogUDP"
	// SyslogRELPEventDestinationProtocol shall indicate the destination follows the Reliable Event Logging Protocol
	// (RELP) transport for syslog as defined by www.rsyslog.com.
	SyslogRELPEventDestinationProtocol EventDestinationProtocol = "SyslogRELP"
	// OEMEventDestinationProtocol shall indicate an OEM specific protocol. The OEMProtocol property shall contain the
	// specific OEM event destination protocol.
	OEMEventDestinationProtocol EventDestinationProtocol = "OEM"
)

// SNMPAuthenticationProtocols is
type SNMPAuthenticationProtocols string

const (
	// NoneSNMPAuthenticationProtocols shall indicate authentication is not required.
	NoneSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "None"
	// CommunityStringSNMPAuthenticationProtocols shall indicate authentication using SNMP community strings and the
	// value of TrapCommunity.
	CommunityStringSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "CommunityString"
	// HMACMD5SNMPAuthenticationProtocols shall indicate authentication conforms to the RFC3414-defined HMAC-MD5-96
	// authentication protocol.
	HMACMD5SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC_MD5"
	// HMACSHA96SNMPAuthenticationProtocols shall indicate authentication conforms to the RFC3414-defined HMAC-SHA-96
	// authentication protocol.
	HMACSHA96SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC_SHA96"
	// HMAC128SHA224SNMPAuthenticationProtocols shall indicate authentication for SNMPv3 access conforms to the
	// RFC7860-defined usmHMAC128SHA224AuthProtocol.
	HMAC128SHA224SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC128_SHA224"
	// HMAC192SHA256SNMPAuthenticationProtocols shall indicate authentication for SNMPv3 access conforms to the
	// RFC7860-defined usmHMAC192SHA256AuthProtocol.
	HMAC192SHA256SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC192_SHA256"
	// HMAC256SHA384SNMPAuthenticationProtocols shall indicate authentication for SNMPv3 access conforms to the
	// RFC7860-defined usmHMAC256SHA384AuthProtocol.
	HMAC256SHA384SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC256_SHA384"
	// HMAC384SHA512SNMPAuthenticationProtocols shall indicate authentication for SNMPv3 access conforms to the
	// RFC7860-defined usmHMAC384SHA512AuthProtocol.
	HMAC384SHA512SNMPAuthenticationProtocols SNMPAuthenticationProtocols = "HMAC384_SHA512"
)

// SNMPEncryptionProtocols is
type SNMPEncryptionProtocols string

const (
	// NoneSNMPEncryptionProtocols shall indicate there is no encryption.
	NoneSNMPEncryptionProtocols SNMPEncryptionProtocols = "None"
	// CBCDESSNMPEncryptionProtocols shall indicate encryption conforms to the RFC3414-defined CBC-DES encryption
	// protocol.
	CBCDESSNMPEncryptionProtocols SNMPEncryptionProtocols = "CBC_DES"
	// CFB128AES128SNMPEncryptionProtocols shall indicate encryption conforms to the RFC3414-defined CFB128-AES-128
	// encryption protocol.
	CFB128AES128SNMPEncryptionProtocols SNMPEncryptionProtocols = "CFB128_AES128"
)

// SubscriptionType is
type SubscriptionType string

const (
	// RedfishEventSubscriptionType The subscription follows the Redfish Specification for event notifications. To send
	// an event notification, a service sends an HTTP POST to the subscriber's destination URI.
	RedfishEventSubscriptionType SubscriptionType = "RedfishEvent"
	// SSESubscriptionType The subscription follows the HTML5 server-sent event definition for event notifications.
	SSESubscriptionType SubscriptionType = "SSE"
	// SNMPTrapSubscriptionType shall indicate the subscription follows the various versions of SNMP Traps for event
	// notifications. Protocol shall specify the appropriate version of SNMP.
	SNMPTrapSubscriptionType SubscriptionType = "SNMPTrap"
	// SNMPInformSubscriptionType shall indicate the subscription follows versions 2 and 3 of SNMP Inform for event
	// notifications. Protocol shall specify the appropriate version of SNMP.
	SNMPInformSubscriptionType SubscriptionType = "SNMPInform"
	// SyslogSubscriptionType shall indicate the subscription forwards syslog messages to the event destination.
	// Protocol shall specify the appropriate syslog protocol.
	SyslogSubscriptionType SubscriptionType = "Syslog"
	// OEMSubscriptionType shall indicate an OEM subscription type. The OEMSubscriptionType property shall contain the
	// specific OEM subscription type.
	OEMSubscriptionType SubscriptionType = "OEM"
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

// EventDestination shall represent the target of an event subscription, including the event types and context to
// provide to the target in the event payload.
type EventDestination struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that represent the
	// server certificates for the server referenced by the Destination property. If VerifyCertificate is 'true',
	// services shall compare the certificates in this collection with the certificate obtained during handshaking with
	// the event destination in order to verify the identify of the event destination prior to sending an event. If the
	// server cannot be verified, the service shall not send the event. If VerifyCertificate is 'false', the service
	// shall not perform certificate verification.
	Certificates string
	// ClientCertificates shall contain a link to a resource collection of type CertificateCollection that represents
	// the client identity certificates that are provided to the server referenced by the Destination property as part
	// of TLS handshaking.
	ClientCertificates string
	// Context shall contain a client-supplied context that remains with the connection through the connection's
	// lifetime.
	Context string
	// DeliveryRetryPolicy shall indicate the subscription delivery retry policy for events where the subscription type
	// is RedfishEvent.
	DeliveryRetryPolicy DeliveryRetryPolicy
	// Description provides a description of this resource.
	Description string
	// Destination shall contain a URI to the destination where the events are sent. If Protocol is 'SMTP', the URI
	// shall follow the RFC6068-described format. SNMP URIs shall be consistent with RFC4088. Specifically, for SNMPv3,
	// if a username is specified in the SNMP URI, the SNMPv3 authentication and encryption configuration associated
	// with that user shall be utilized in the SNMPv3 traps. Syslog URIs shall be consistent with RFC3986 and contain
	// the scheme 'syslog://'. Server-sent event destinations shall be in the form 'redfish-sse://ip:port' where 'ip'
	// and 'port' are the IP address and the port of the client with the open SSE connection. For other URIs, such as
	// HTTP or HTTPS, they shall be consistent with RFC3986.
	Destination string
	// EventFormatType shall indicate the content types of the message that this service sends to the EventDestination.
	// If this property is not present, the EventFormatType shall be assumed to be Event.
	EventFormatType EventFormatType
	// ExcludeMessageIds shall contain an array of exculded MessageIds that are not allowed values for the MessageId
	// property within an event sent to the subscriber. The MessageId shall be in the 'MessageRegistry.MessageId'
	// format. If included, the MessageId major and minor version details should be ignored. Events with a MessageId
	// that is contained in this array shall not be sent to the subscriber. If this property is an empty array or is
	// absent, no exclusive filtering based upon the MessageId of an event is performed.
	ExcludeMessageIds []string
	// ExcludeRegistryPrefixes shall contain an array of prefixes of excluded message registries that contain the
	// MessageIds that are not allowed values for the MessageId property within an event sent to the subscriber. Events
	// with a MessageId that is from a message registry contained in this array shall not be sent to the subscriber. If
	// this property is an empty array or is absent, no exclusive filtering based upon message registry of the
	// MessageId of an event is performed.
	ExcludeRegistryPrefixes []string
	// HeartbeatIntervalMinutes shall indicate the interval for sending periodic heartbeat events to the subscriber.
	// The value shall be the interval, in minutes, between each periodic event. This property shall not be present if
	// the SendHeartbeat property is not present.
	HeartbeatIntervalMinutes int
	// HttpHeaders shall contain an object consisting of the names and values of of HTTP header to be included with
	// every event POST to the event destination. This object shall be null or an empty array in responses. An empty
	// array is the preferred return value in responses.
	HttpHeaders []HttpHeaderProperty
	// IncludeOriginOfCondition shall indicate whether the event payload sent to the subscription destination will
	// expand the OriginOfCondition property to include the resource or object referenced by the OriginOfCondition
	// property.
	IncludeOriginOfCondition bool
	// MessageIds shall contain an array of MessageIds that are the allowable values for the MessageId property within
	// an event sent to the subscriber. The MessageId should be in the 'MessageRegistry.MessageId' format. If included,
	// the MessageId major and minor version details should be ignored. Events with a MessageId that is not contained
	// in this array and is not from a message registry contained in RegistryPrefixes shall not be sent to the
	// subscriber. If this property is an empty array or is absent, no inclusive filtering based upon the MessageId of
	// an event is performed.
	MessageIds []string
	// MetricReportDefinitions shall specify an array of metric report definitions that are the only allowable
	// generators of metric reports for this subscription. Metric reports originating from metric report definitions
	// not contained in this array shall not be sent to the subscriber. If this property is absent or the array is
	// empty, the service shall send metric reports originating from any metric report definition to the subscriber.
	MetricReportDefinitions []MetricReportDefinition
	// MetricReportDefinitions@odata.count
	MetricReportDefinitionsCount int `json:"MetricReportDefinitions@odata.count"`
	// OEMProtocol shall contain the protocol type that the event uses to send the event to the destination. This
	// property shall be present if Protocol is 'OEM'.
	OEMProtocol string
	// OEMSubscriptionType shall indicate the OEM-defined type of subscription for events. This property shall be
	// present if SubscriptionType is 'OEM'.
	OEMSubscriptionType string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OriginResources shall specify an array of resources, resource collections, or referenceable members that are the
	// only allowable values for the OriginOfCondition property within an event that the service sends to the
	// subscriber. Events with an OriginOfCondition that is not contained in this array, and is not subordinate to
	// members of this array if SubordinateResources contains the value 'true', shall not be sent to the subscriber. If
	// this property is an empty array or is absent, no filtering based upon the URI of the OriginOfCondition of an
	// event is performed.
	OriginResources []idRef
	// OriginResources@odata.count
	OriginResourcesCount int `json:"OriginResources@odata.count"`
	// Protocol shall contain the protocol type that the event uses to send the event to the destination. A 'Redfish'
	// value shall indicate that the event type shall adhere to the type defined in the Redfish Specification.
	Protocol string
	// RegistryPrefixes shall contain an array the prefixes of message registries that contain the MessageIds that are
	// the allowable values for the MessageId property within an event sent to the subscriber. Events with a MessageId
	// that is not from a message registry contained in this array and is not contained by MessageIds shall not be sent
	// to the subscriber. If this property is an empty array or is absent, no inclusive filtering based upon message
	// registry of the MessageId of an event is performed.
	RegistryPrefixes []string
	// ResourceTypes shall specify an array of resource type values that contain the allowable resource types for the
	// resource referenced by the OriginOfCondition property. Events with the resource type of the resource referenced
	// by the OriginOfCondition property that is not contained in this array shall not be sent to the subscriber. If
	// this property is an empty array or is absent, no filtering based upon the resource type of the OriginOfCondition
	// of an event is performed. This property shall contain only the general namespace for the type and not the
	// versioned value. For example, it shall not contain 'Task.v1_2_0.Task' and instead shall contain 'Task'. To
	// specify that a client is subscribing to metric reports, the EventTypes property should include 'MetricReport'.
	ResourceTypes []string
	// SNMP shall contain the settings for an SNMP event destination.
	SNMP string
	// SendHeartbeat shall indicate that the service shall periodically send the 'RedfishServiceFunctional' message
	// defined in the Heartbeat Event Message Registry to the subscriber. If this property is not present, no periodic
	// event shall be sent. This property shall not apply to event destinations if the SubscriptionType property
	// contains the value 'SSE'.
	SendHeartbeat bool
	// Status shall contain the status of the subscription.
	Status common.Status
	// SubordinateResources shall indicate whether the subscription is for events in the OriginResources array and its
	// subordinate resources. If 'true' and the OriginResources array is specified, the subscription is for events in
	// the OriginResources array and its subordinate resources. Note that resources associated through the Links
	// section are not considered subordinate. If 'false' and the OriginResources array is specified, the subscription
	// shall be for events in the OriginResources array only. If the OriginResources array is not present, this
	// property shall have no relevance.
	SubordinateResources bool
	// SubscriptionType shall indicate the type of subscription for events. If this property is not present, the
	// SubscriptionType shall be assumed to be RedfishEvent.
	SubscriptionType SubscriptionType
	// SyslogFilters shall describe all desired syslog messages to send to a remote syslog server. If this property
	// contains an empty array or is absent, all messages shall be sent.
	SyslogFilters []SyslogFilter
	// VerifyCertificate shall indicate whether whether the service will verify the certificate of the server
	// referenced by the Destination property prior to sending the event. If this property is not supported by the
	// service or specified by the client in the create request, it shall be assumed to be 'false'.
	VerifyCertificate bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EventDestination object from the raw JSON.
func (eventdestination *EventDestination) UnmarshalJSON(b []byte) error {
	type temp EventDestination
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*eventdestination = EventDestination(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	eventdestination.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (eventdestination *EventDestination) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(EventDestination)
	original.UnmarshalJSON(eventdestination.rawData)

	readWriteFields := []string{
		"Context",
		"DeliveryRetryPolicy",
		"VerifyCertificate",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(eventdestination).Elem()

	return eventdestination.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEventDestination will get a EventDestination instance from the service.
func GetEventDestination(c common.Client, uri string) (*EventDestination, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var eventdestination EventDestination
	err = json.NewDecoder(resp.Body).Decode(&eventdestination)
	if err != nil {
		return nil, err
	}

	eventdestination.SetClient(c)
	return &eventdestination, nil
}

// ListReferencedEventDestinations gets the collection of EventDestination from
// a provided reference.
func ListReferencedEventDestinations(c common.Client, link string) ([]*EventDestination, error) {
	var result []*EventDestination
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, eventdestinationLink := range links.ItemLinks {
		eventdestination, err := GetEventDestination(c, eventdestinationLink)
		if err != nil {
			collectionError.Failures[eventdestinationLink] = err
		} else {
			result = append(result, eventdestination)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// HttpHeaderProperty shall contain the HTTP header name and value to include with every event POST to the event
// destination.
type HttpHeaderProperty struct {
}

// UnmarshalJSON unmarshals a HttpHeaderProperty object from the raw JSON.
func (httpheaderproperty *HttpHeaderProperty) UnmarshalJSON(b []byte) error {
	type temp HttpHeaderProperty
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*httpheaderproperty = HttpHeaderProperty(t.temp)

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

// SNMPSettings shall contain the settings for an SNMP event destination.
type SNMPSettings struct {
	// AuthenticationKey shall contain the key for SNMPv3 authentication. The value shall be 'null' in responses. This
	// property accepts a passphrase or a hex-encoded key. If the string starts with 'Passphrase:', the remainder of
	// the string shall be the passphrase and shall be converted to the key as described in the 'Password to Key
	// Algorithm' section of RFC3414. If the string starts with 'Hex:', then the remainder of the string shall be the
	// key encoded in hexadecimal notation. If the string starts with neither, the full string shall be a passphrase
	// and shall be converted to the key as described in the 'Password to Key Algorithm' section of RFC3414. The
	// passphrase can contain any printable characters except for the double quotation mark.
	AuthenticationKey string
	// AuthenticationKeySet shall contain 'true' if a valid value was provided for the AuthenticationKey property.
	// Otherwise, the property shall contain 'false'.
	AuthenticationKeySet string
	// AuthenticationProtocol shall contain the SNMPv3 authentication protocol.
	AuthenticationProtocol SNMPAuthenticationProtocols
	// EncryptionKey shall contain the key for SNMPv3 encryption. The value shall be 'null' in responses. This property
	// accepts a passphrase or a hex-encoded key. If the string starts with 'Passphrase:', the remainder of the string
	// shall be the passphrase and shall be converted to the key as described in the 'Password to Key Algorithm'
	// section of RFC3414. If the string starts with 'Hex:', then the remainder of the string shall be the key encoded
	// in hexadecimal notation. If the string starts with neither, the full string shall be a passphrase and shall be
	// converted to the key as described in the 'Password to Key Algorithm' section of RFC3414. The passphrase can
	// contain any printable characters except for the double quotation mark.
	EncryptionKey string
	// EncryptionKeySet shall contain 'true' if a valid value was provided for the EncryptionKey property. Otherwise,
	// the property shall contain 'false'.
	EncryptionKeySet string
	// EncryptionProtocol shall contain the SNMPv3 encryption protocol.
	EncryptionProtocol SNMPEncryptionProtocols
	// TrapCommunity shall contain the SNMP trap community string. The value shall be 'null' in responses.
	TrapCommunity string
}

// UnmarshalJSON unmarshals a SNMPSettings object from the raw JSON.
func (snmpsettings *SNMPSettings) UnmarshalJSON(b []byte) error {
	type temp SNMPSettings
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*snmpsettings = SNMPSettings(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SyslogFilter shall contain the filter for a syslog message. The filter shall describe the desired syslog message
// to forward to a remote syslog server.
type SyslogFilter struct {
	// LogFacilities shall contain the types of programs that can log messages. If this property contains an empty
	// array or is absent, all facilities shall be indicated.
	LogFacilities []SyslogFacility
	// LowestSeverity shall contain the lowest syslog severity level that will be forwarded. The service shall forward
	// all messages equal to or greater than the value in this property. The value 'All' shall indicate all severities.
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
