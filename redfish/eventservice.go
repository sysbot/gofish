//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// SMTPAuthenticationMethods is
type SMTPAuthenticationMethods string

const (
	// NoneSMTPAuthenticationMethods shall indicate authentication is not required.
	NoneSMTPAuthenticationMethods SMTPAuthenticationMethods = "None"
	// AutoDetectSMTPAuthenticationMethods shall indicate authentication is auto-detected.
	AutoDetectSMTPAuthenticationMethods SMTPAuthenticationMethods = "AutoDetect"
	// PlainSMTPAuthenticationMethods shall indicate authentication conforms to the RFC4954-defined AUTH PLAIN
	// mechanism.
	PlainSMTPAuthenticationMethods SMTPAuthenticationMethods = "Plain"
	// LoginSMTPAuthenticationMethods shall indicate authentication conforms to the RFC4954-defined AUTH LOGIN
	// mechanism.
	LoginSMTPAuthenticationMethods SMTPAuthenticationMethods = "Login"
	// CRAMMD5SMTPAuthenticationMethods shall indicate authentication conforms to the RFC4954-defined AUTH CRAM-MD5
	// mechanism.
	CRAMMD5SMTPAuthenticationMethods SMTPAuthenticationMethods = "CRAM_MD5"
)

// SMTPConnectionProtocol is
type SMTPConnectionProtocol string

const (
	// NoneSMTPConnectionProtocol shall indicate the connection is in clear text.
	NoneSMTPConnectionProtocol SMTPConnectionProtocol = "None"
	// AutoDetectSMTPConnectionProtocol shall indicate the connection is auto-detected.
	AutoDetectSMTPConnectionProtocol SMTPConnectionProtocol = "AutoDetect"
	// StartTLSSMTPConnectionProtocol shall indicate the connection conforms to the RFC3207-defined StartTLS extension.
	StartTLSSMTPConnectionProtocol SMTPConnectionProtocol = "StartTLS"
	// TLSSSLSMTPConnectionProtocol shall indicate the connection is TLS/SSL.
	TLSSSLSMTPConnectionProtocol SMTPConnectionProtocol = "TLS_SSL"
)

// EventService shall represent an event service for a Redfish implementation.
type EventService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// DeliveryRetryAttempts shall contain the number of times that the POST of an event is retried before the
	// subscription terminates. This retry occurs at the service level, which means that the HTTP POST to the event
	// destination fails with an HTTP '4XX' or '5XX' status code or an HTTP timeout occurs this many times before the
	// event destination subscription terminates.
	DeliveryRetryAttempts string
	// DeliveryRetryIntervalSeconds shall contain the interval, in seconds, between the retry attempts for any event
	// sent to the subscription destination.
	DeliveryRetryIntervalSeconds string
	// Description provides a description of this resource.
	Description string
	// EventFormatTypes shall contain the content types of the message that this service can send to the event
	// destination. If this property is not present, the EventFormatType shall be assumed to be 'Event'.
	EventFormatTypes []EventFormatType
	// ExcludeMessageId shall indicate whether this service supports filtering by the ExcludeMessageIds property.
	ExcludeMessageId string
	// ExcludeRegistryPrefix shall indicate whether this service supports filtering by the ExcludeRegistryPrefixes
	// property.
	ExcludeRegistryPrefix string
	// IncludeOriginOfConditionSupported shall indicate whether the service supports including the resource payload of
	// the origin of condition in the event payload. If 'true', event subscriptions are allowed to specify the
	// IncludeOriginOfCondition property.
	IncludeOriginOfConditionSupported bool
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RegistryPrefixes shall contain the array of the prefixes of the message registries that shall be allowed or
	// excluded for an event subscription.
	RegistryPrefixes []string
	// ResourceTypes shall specify an array of the valid @odata.type values that can be used for an event subscription.
	ResourceTypes []string
	// SMTP shall contain settings for SMTP event delivery.
	SMTP string
	// SSEFilterPropertiesSupported shall contain the properties that are supported in the '$filter' query parameter
	// for the URI indicated by the ServerSentEventUri property, as described by the Redfish Specification.
	SSEFilterPropertiesSupported string
	// ServerSentEventUri shall contain a URI that specifies an HTML5 Server-Sent Event-conformant endpoint.
	ServerSentEventUri string
	// ServiceEnabled shall indicate whether this service is enabled. If 'false', events are no longer published, new
	// SSE connections cannot be established, and existing SSE connections are terminated.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SubordinateResourcesSupported shall indicate whether the service supports the SubordinateResources property on
	// both event subscriptions and generated events.
	SubordinateResourcesSupported bool
	// Subscriptions shall contain the link to a resource collection of type EventDestinationCollection.
	Subscriptions string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a EventService object from the raw JSON.
func (eventservice *EventService) UnmarshalJSON(b []byte) error {
	type temp EventService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*eventservice = EventService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	eventservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (eventservice *EventService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(EventService)
	original.UnmarshalJSON(eventservice.rawData)

	readWriteFields := []string{
		"DeliveryRetryAttempts",
		"DeliveryRetryIntervalSeconds",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(eventservice).Elem()

	return eventservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetEventService will get a EventService instance from the service.
func GetEventService(c common.Client, uri string) (*EventService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var eventservice EventService
	err = json.NewDecoder(resp.Body).Decode(&eventservice)
	if err != nil {
		return nil, err
	}

	eventservice.SetClient(c)
	return &eventservice, nil
}

// ListReferencedEventServices gets the collection of EventService from
// a provided reference.
func ListReferencedEventServices(c common.Client, link string) ([]*EventService, error) {
	var result []*EventService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, eventserviceLink := range links.ItemLinks {
		eventservice, err := GetEventService(c, eventserviceLink)
		if err != nil {
			collectionError.Failures[eventserviceLink] = err
		} else {
			result = append(result, eventservice)
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

// SMTP shall contain settings for SMTP event delivery.
type SMTP struct {
	// Authentication shall contain the authentication method for the SMTP server.
	Authentication SMTPAuthenticationMethods
	// ConnectionProtocol shall contain the connection type to the outgoing SMTP server.
	ConnectionProtocol SMTPConnectionProtocol
	// FromAddress shall contain the email address to use for the 'from' field in an outgoing email.
	FromAddress string
	// Password shall contain the password for authentication with the SMTP server. The value shall be 'null' in
	// responses.
	Password string
	// Port shall contain the destination port for the SMTP server.
	Port int
	// ServerAddress shall contain the address of the SMTP server for outgoing email.
	ServerAddress string
	// ServiceEnabled shall indicate if SMTP for event delivery is enabled.
	ServiceEnabled bool
	// Username shall contain the username for authentication with the SMTP server.
	Username string
}

// UnmarshalJSON unmarshals a SMTP object from the raw JSON.
func (smtp *SMTP) UnmarshalJSON(b []byte) error {
	type temp SMTP
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*smtp = SMTP(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SSEFilterPropertiesSupported shall contain a set of properties that are supported in the '$filter' query
// parameter for the URI indicated by the ServerSentEventUri property, as described by the Redfish Specification.
type SSEFilterPropertiesSupported struct {
	// EventFormatType shall indicate whether this service supports filtering by the EventFormatType property.
	EventFormatType string
	// MessageId shall indicate whether this service supports filtering by the MessageIds property.
	MessageId string
	// MetricReportDefinition shall indicate whether this service supports filtering by the MetricReportDefinitions
	// property.
	MetricReportDefinition string
	// OriginResource shall indicate whether this service supports filtering by the OriginResources property.
	OriginResource string
	// RegistryPrefix shall indicate whether this service supports filtering by the RegistryPrefixes property.
	RegistryPrefix string
	// ResourceType shall indicate whether this service supports filtering by the ResourceTypes property.
	ResourceType string
	// SubordinateResources shall indicate whether this service supports filtering by the SubordinateResources
	// property.
	SubordinateResources string
}

// UnmarshalJSON unmarshals a SSEFilterPropertiesSupported object from the raw JSON.
func (ssefilterpropertiessupported *SSEFilterPropertiesSupported) UnmarshalJSON(b []byte) error {
	type temp SSEFilterPropertiesSupported
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ssefilterpropertiessupported = SSEFilterPropertiesSupported(t.temp)

	// Extract the links to other entities for later

	return nil
}
