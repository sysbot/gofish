//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Event This resource contains an event for a Redfish implementation.
type Event struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Context shall contain a client supplied context for the event destination to which this event is being sent.
	Context string
	// Description provides a description of this resource.
	Description string
	// Events shall contain an array of objects that represent the occurrence of one or more events.
	Events []EventRecord
	// Events@odata.count
	EventsCount int `json:"Events@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a Event object from the raw JSON.
func (event *Event) UnmarshalJSON(b []byte) error {
	type temp Event
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*event = Event(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetEvent will get a Event instance from the service.
func GetEvent(c common.Client, uri string) (*Event, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var event Event
	err = json.NewDecoder(resp.Body).Decode(&event)
	if err != nil {
		return nil, err
	}

	event.SetClient(c)
	return &event, nil
}

// ListReferencedEvents gets the collection of Event from
// a provided reference.
func ListReferencedEvents(c common.Client, link string) ([]*Event, error) {
	var result []*Event
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, eventLink := range links.ItemLinks {
		event, err := GetEvent(c, eventLink)
		if err != nil {
			collectionError.Failures[eventLink] = err
		} else {
			result = append(result, event)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// EventRecord
type EventRecord struct {
	// Actions shall contain the available actions for this resource.
	Actions string
	// EventGroupId shall indicate that events are related and shall have the same value when multiple event messages
	// are produced by the same root cause. Implementations shall use separate values for events with a separate root
	// cause. This property value shall not imply an ordering of events. The '0' value shall indicate that this event
	// is not grouped with any other event.
	EventGroupId string
	// EventId shall contain a service-defined unique identifier for the event.
	EventId string
	// EventTimestamp shall indicate the time the event occurred where the value shall be consistent with the Redfish
	// service time that is also used for the values of the Modified property.
	EventTimestamp string
	// LogEntry shall contain a link to a resource of type LogEntry that represents the log entry created for this
	// event.
	LogEntry string
	// MemberId shall contain the unique identifier for this member within an array. For services supporting Redfish
	// v1.6 or higher, this value shall contain the zero-based array index.
	MemberId string
	// Message shall contain a human-readable event message.
	Message string
	// MessageArgs shall contain an array of message arguments that are substituted for the arguments in the message
	// when looked up in the message registry. It has the same semantics as the MessageArgs property in the Redfish
	// MessageRegistry schema.
	MessageArgs []string
	// MessageId shall contain a MessageId, as defined in the 'MessageId format' clause of the Redfish Specification.
	MessageId string
	// MessageSeverity shall contain the severity of the message in this event. Services can replace the value defined
	// in the message registry with a value more applicable to the implementation.
	MessageSeverity string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OriginOfCondition shall contain a link to the resource or object that originated the condition that caused the
	// event to be generated. If the event subscription has the IncludeOriginOfCondition property set to 'true', it
	// shall include the entire resource or object referenced by the link. For events that represent the creation or
	// deletion of a resource, this property should reference the created or deleted resource and not the collection
	// that contains the resource.
	OriginOfCondition string
	// SpecificEventExistsInGroup shall indicate that the event is equivalent to another event, with a more specific
	// definition, within the same EventGroupId. For example, the 'DriveFailed' message from the Storage Device Message
	// Registry is more specific than the 'ResourceStatusChangedCritical' message from the Resource Event Message
	// Registry, when both occur with the same EventGroupId. This property shall contain 'true' if a more specific
	// event is available, and shall contain 'false' if no equivalent event exists in the same EventGroupId. If this
	// property is absent, the value shall be assumed to be 'false'.
	SpecificEventExistsInGroup string
}

// UnmarshalJSON unmarshals a EventRecord object from the raw JSON.
func (eventrecord *EventRecord) UnmarshalJSON(b []byte) error {
	type temp EventRecord
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*eventrecord = EventRecord(t.temp)

	// Extract the links to other entities for later

	return nil
}

// EventRecordActions shall contain the available actions for this resource.
type EventRecordActions struct {
	// Oem shall contain the available OEM-specific actions for this resource.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a EventRecordActions object from the raw JSON.
func (eventrecordactions *EventRecordActions) UnmarshalJSON(b []byte) error {
	type temp EventRecordActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*eventrecordactions = EventRecordActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// EventRecordOemActions shall contain the available OEM-specific actions for this resource.
type EventRecordOemActions struct {
}

// UnmarshalJSON unmarshals a EventRecordOemActions object from the raw JSON.
func (eventrecordoemactions *EventRecordOemActions) UnmarshalJSON(b []byte) error {
	type temp EventRecordOemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*eventrecordoemactions = EventRecordOemActions(t.temp)

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
