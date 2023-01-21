//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// SessionTypes is
type SessionTypes string

const (
	// HostConsoleSessionTypes The host's console, which could be connected through Telnet, SSH, or other protocol.
	HostConsoleSessionTypes SessionTypes = "HostConsole"
	// ManagerConsoleSessionTypes The manager's console, which could be connected through Telnet, SSH, SM CLP, or other
	// protocol.
	ManagerConsoleSessionTypes SessionTypes = "ManagerConsole"
	// IPMISessionTypes Intelligent Platform Management Interface.
	IPMISessionTypes SessionTypes = "IPMI"
	// KVMIPSessionTypes Keyboard-Video-Mouse over IP Session.
	KVMIPSessionTypes SessionTypes = "KVMIP"
	// OEMSessionTypes OEM type. For OEM session types, see the OemSessionType property.
	OEMSessionTypes SessionTypes = "OEM"
	// RedfishSessionTypes A Redfish session.
	RedfishSessionTypes SessionTypes = "Redfish"
	// VirtualMediaSessionTypes Virtual media.
	VirtualMediaSessionTypes SessionTypes = "VirtualMedia"
	// WebUISessionTypes A non-Redfish web user interface session, such as a graphical interface or another web-based
	// protocol.
	WebUISessionTypes SessionTypes = "WebUI"
)

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

// Session shall represent a session for a Redfish implementation.
type Session struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ClientOriginIPAddress shall contain the IP address of the client that created the session.
	ClientOriginIPAddress string
	// Context shall contain a client-supplied context that remains with the session through the session's lifetime.
	Context string
	// CreatedTime shall contain the date and time when the session was created.
	CreatedTime string
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OemSessionType shall contain the OEM-specific session type that is currently active if SessionType contains
	// 'OEM'.
	OemSessionType string
	// Password shall contain the password for this session. The value shall be 'null' in responses.
	Password string
	// SessionType shall represent the type of session that is currently active.
	SessionType SessionTypes
	// UserName shall contain the username that matches an account recognized by the account service.
	UserName string
}

// UnmarshalJSON unmarshals a Session object from the raw JSON.
func (session *Session) UnmarshalJSON(b []byte) error {
	type temp Session
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*session = Session(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetSession will get a Session instance from the service.
func GetSession(c common.Client, uri string) (*Session, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var session Session
	err = json.NewDecoder(resp.Body).Decode(&session)
	if err != nil {
		return nil, err
	}

	session.SetClient(c)
	return &session, nil
}

// ListReferencedSessions gets the collection of Session from
// a provided reference.
func ListReferencedSessions(c common.Client, link string) ([]*Session, error) {
	var result []*Session
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, sessionLink := range links.ItemLinks {
		session, err := GetSession(c, sessionLink)
		if err != nil {
			collectionError.Failures[sessionLink] = err
		} else {
			result = append(result, session)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
