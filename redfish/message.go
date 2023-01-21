//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Message shall contain a message that the Redfish service returns, as described in the Redfish Specification.
type Message struct {
	// Message shall contain a human-readable message.
	Message string
	// MessageArgs shall contain an array of message arguments that are substituted for the arguments in the message
	// when looked up in the message registry. It has the same semantics as the MessageArgs property in the Redfish
	// MessageRegistry schema.
	MessageArgs []string
	// MessageId shall contain a MessageId, as defined in the 'MessageId format' clause of the Redfish Specification.
	MessageId string
	// MessageSeverity shall contain the severity of the message. Services can replace the value defined in the message
	// registry with a value more applicable to the implementation.
	MessageSeverity string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RelatedProperties shall contain an array of RFC6901-defined JSON pointers indicating the properties described by
	// the message, if appropriate for the message.
	RelatedProperties []string
	// Resolution shall contain the resolution of the message. Services can replace the resolution defined in the
	// message registry with a more specific resolution in message payloads.
	Resolution string
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

// GetMessage will get a Message instance from the service.
func GetMessage(c common.Client, uri string) (*Message, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var message Message
	err = json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		return nil, err
	}

	message.SetClient(c)
	return &message, nil
}

// ListReferencedMessages gets the collection of Message from
// a provided reference.
func ListReferencedMessages(c common.Client, link string) ([]*Message, error) {
	var result []*Message
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, messageLink := range links.ItemLinks {
		message, err := GetMessage(c, messageLink)
		if err != nil {
			collectionError.Failures[messageLink] = err
		} else {
			result = append(result, message)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
