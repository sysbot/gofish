//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ClearingType is The conditions when an event is cleared.
type ClearingType string

const (
	// SameOriginOfConditionClearingType shall describe when the message for an event is cleared by the other messages
	// in the ClearingLogic property, provided the OriginOfCondition for both events are the same.
	SameOriginOfConditionClearingType ClearingType = "SameOriginOfCondition"
)

// ParamType is
type ParamType string

const (
	// stringParamType The argument is a string.
	stringParamType ParamType = "string"
	// numberParamType The argument is a number.
	numberParamType ParamType = "number"
)

// ClearingLogic shall contain the available actions for this resource.
type ClearingLogic struct {
	// ClearsAll shall indicate whether all prior conditions and messages are cleared, provided the ClearsIf condition
	// is met.
	ClearsAll bool
	// ClearsIf shall contain the condition the event is cleared.
	ClearsIf ClearingType
	// ClearsMessage shall contain an array of MessageIds that this message clears when the other conditions are met.
	// The MessageIds shall not include the message registry name or version and shall contain only the MessageId
	// portion. MessageIds shall not refer to other message registries.
	ClearsMessage []string
}

// UnmarshalJSON unmarshals a ClearingLogic object from the raw JSON.
func (clearinglogic *ClearingLogic) UnmarshalJSON(b []byte) error {
	type temp ClearingLogic
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*clearinglogic = ClearingLogic(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Message shall represent how a message is defined within a message registry.
type Message struct {
	// ArgDescriptions shall contain an ordered array of text describing each argument used as substitution in the
	// message.
	ArgDescriptions []string
	// ArgLongDescriptions shall contain an ordered array of normative language for each argument used as substitution
	// in the message.
	ArgLongDescriptions []string
	// ClearingLogic shall contain the available actions for this resource.
	ClearingLogic string
	// Deprecated shall indicate that a message is deprecated. The value of the string should explain the deprecation,
	// including reference to new message or messages to be used. The message can be supported in new and existing
	// implementations, but usage in new implementations is discouraged. Deprecated messages are likely to be removed
	// in a future major version of the message registry.
	Deprecated string
	// Description provides a description of this resource.
	Description string
	// LongDescription shall contain the normative language that describes this message's usage in a Redfish
	// implementation.
	LongDescription string
	// Message shall contain the message to display. If a %integer is included in part of the string, it shall
	// represent a string substitution for any MessageArgs that accompany the message, in order.
	Message string
	// MessageSeverity shall contain the severity of the message. Services can replace the severity defined in the
	// message registry with a value more applicable to the implementation in message payloads and event payloads.
	MessageSeverity Health
	// NumberOfArgs shall contain the number of arguments that are substituted for the locations marked with %<integer>
	// in the message.
	NumberOfArgs string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ParamTypes shall contain an ordered array of argument data types that match the data types of the MessageArgs.
	ParamTypes []ParamType
	// Resolution shall contain the resolution of the message. Services can replace the resolution defined in the
	// message registry with a more specific resolution in message payloads.
	Resolution string
	// VersionAdded shall contain the version of the message registry when the message was added. This property shall
	// not appear for messages created at version '1.0.0' of a message registry.
	VersionAdded string
	// VersionDeprecated shall contain the version of the registry when the message was deprecated. This property shall
	// not appear if the message has not been deprecated.
	VersionDeprecated string
}

// UnmarshalJSON unmarshals a Message object from the raw JSON.
func (message *Message) UnmarshalJSON(b []byte) error {
	type temp Message
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*message = Message(t.temp)

	// Extract the links to other entities for later

	return nil
}

// MessageProperty shall contain the message keys contained in the message registry. The message keys are the
// suffix of the MessageId and shall be unique within this message registry.
type MessageProperty struct {
}

// UnmarshalJSON unmarshals a MessageProperty object from the raw JSON.
func (messageproperty *MessageProperty) UnmarshalJSON(b []byte) error {
	type temp MessageProperty
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*messageproperty = MessageProperty(t.temp)

	// Extract the links to other entities for later

	return nil
}

// MessageRegistry shall represent a message registry for a Redfish implementation.
type MessageRegistry struct {
	common.Entity
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// Language shall contain an RFC5646-conformant language code.
	Language string
	// Messages shall contain the message keys contained in the message registry. The message keys are the suffix of
	// the MessageId and shall be unique within this message registry.
	Messages string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OwningEntity shall represent the publisher of this message registry.
	OwningEntity string
	// RegistryPrefix shall contain the Redfish Specification-defined prefix used in forming and decoding MessageIds
	// that uniquely identifies all messages that belong to this message registry.
	RegistryPrefix string
	// RegistryVersion shall contain the version of this message registry.
	RegistryVersion string
}

// UnmarshalJSON unmarshals a MessageRegistry object from the raw JSON.
func (messageregistry *MessageRegistry) UnmarshalJSON(b []byte) error {
	type temp MessageRegistry
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*messageregistry = MessageRegistry(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetMessageRegistry will get a MessageRegistry instance from the service.
func GetMessageRegistry(c common.Client, uri string) (*MessageRegistry, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var messageregistry MessageRegistry
	err = json.NewDecoder(resp.Body).Decode(&messageregistry)
	if err != nil {
		return nil, err
	}

	messageregistry.SetClient(c)
	return &messageregistry, nil
}

// ListReferencedMessageRegistrys gets the collection of MessageRegistry from
// a provided reference.
func ListReferencedMessageRegistrys(c common.Client, link string) ([]*MessageRegistry, error) {
	var result []*MessageRegistry
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, messageregistryLink := range links.ItemLinks {
		messageregistry, err := GetMessageRegistry(c, messageregistryLink)
		if err != nil {
			collectionError.Failures[messageregistryLink] = err
		} else {
			result = append(result, messageregistry)
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
