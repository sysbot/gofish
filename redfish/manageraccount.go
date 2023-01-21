//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// SNMPAuthenticationProtocols is
type SNMPAuthenticationProtocols string

const (
	// NoneSNMPAuthenticationProtocols shall indicate authentication is not required.
	NoneSNMPAuthenticationProtocols SNMPAuthenticationProtocols = "None"
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
	// CFB128AES128SNMPEncryptionProtocols shall indicate encryption conforms to the RFC3826-defined CFB128-AES-128
	// encryption protocol.
	CFB128AES128SNMPEncryptionProtocols SNMPEncryptionProtocols = "CFB128_AES128"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Role shall contain a link to a resource of type Role, and should link to the resource identified by the RoleId
	// property.
	Role string
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

// ManagerAccount shall represent a user account for the manager in a Redfish implementation. The account shall
// indicate the allowed access to one of more services in the manager.
type ManagerAccount struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccountExpiration shall contain the date and time when this account expires. The service shall disable or delete
	// an account that has expired. This property shall not apply to accounts created by the Redfish Host Interface
	// Specification-defined credential bootstrapping. If the value is 'null', or the property is not present, the
	// account never expires.
	AccountExpiration string
	// AccountTypes shall contain an array of the various manager services that the account is allowed to access. This
	// shall not include functionality for receiving events or other notifications. If this property is not provided by
	// the client, the default value shall be an array that contains the value 'Redfish'. The service may add
	// additional values when this property is set or updated if allowed by the value of the StrictAccountTypes
	// property.
	AccountTypes []AccountTypes
	// Actions shall contain the available actions for this resource.
	Actions string
	// Certificates shall contain a link to a resource collection of type CertificateCollection.
	Certificates string
	// Description provides a description of this resource.
	Description string
	// Enabled shall indicate whether an account is enabled. If 'true', the account is enabled and the user can log in.
	// If 'false', the account is disabled and, in the future, the user cannot log in.
	Enabled string
	// HostBootstrapAccount shall indicate whether this account is a bootstrap account created by the Redfish Host
	// Interface Specification-defined credential bootstrapping.
	HostBootstrapAccount int
	// Keys shall contain a link to a resource collection of type KeyCollection that contains the keys that can be used
	// to authenticate this account.
	Keys string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Locked shall indicate whether the account service automatically locked the account because the
	// AccountLockoutThreshold was exceeded. To manually unlock the account before the lockout duration period, an
	// administrator shall be able to change the property to 'false' to clear the lockout condition.
	Locked string
	// OEMAccountTypes shall contain an array of the OEM account types for this account. This property shall be valid
	// when AccountTypes contains 'OEM'.
	OEMAccountTypes []string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Password shall contain the password for this account. The value shall be 'null' in responses.
	Password string
	// PasswordChangeRequired shall indicate whether the service requires that the password for this account be changed
	// before further access to the account is allowed. The implementation may deny access to the service if the
	// password has not been changed. A manager account created with an initial PasswordChangeRequired value of 'true'
	// may force a password change before first access of the account. When the Password property for this account is
	// updated, the service shall set this property to 'false'.
	PasswordChangeRequired bool
	// PasswordExpiration shall contain the date and time when this account password expires. If the value is 'null',
	// the account password never expires. If provided during account creation or password modification, this value
	// shall override the value of the PasswordExpirationDays property in the AccountService resource.
	PasswordExpiration string
	// RoleId shall contain the RoleId of the role resource configured for this account. The service shall reject POST,
	// PATCH, or PUT operations that provide a RoleId that does not exist by returning the HTTP 400 (Bad Request)
	// status code.
	RoleId string
	// SNMP shall contain the SNMP settings for this account when AccountTypes contains 'SNMP'.
	SNMP SNMPUserInfo
	// StrictAccountTypes shall indicate if the service needs to use the value of AccountTypes and OEMAccountTypes
	// values exactly as specified. A 'true' value shall indicate the service needs to either accept the value without
	// changes or reject the request. A 'false' value shall indicate the service may add additional 'AccountTypes' and
	// 'OEMAccountTypes' values as needed to support limitations it has in separately controlling access to individual
	// services. If this property is not present, the value shall be assumed to be 'false'. An update of the service
	// can cause account types to be added to or removed from the AccountTypes and OEMAccountTypes properties,
	// regardless of the value of this property. After a service update, clients should inspect all accounts where the
	// value of this property is 'true' and perform maintenance as needed.
	StrictAccountTypes bool
	// UserName shall contain the user name for this account.
	UserName string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ManagerAccount object from the raw JSON.
func (manageraccount *ManagerAccount) UnmarshalJSON(b []byte) error {
	type temp ManagerAccount
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*manageraccount = ManagerAccount(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	manageraccount.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (manageraccount *ManagerAccount) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ManagerAccount)
	original.UnmarshalJSON(manageraccount.rawData)

	readWriteFields := []string{
		"AccountExpiration",
		"AccountTypes",
		"Enabled",
		"Locked",
		"OEMAccountTypes",
		"Password",
		"PasswordChangeRequired",
		"PasswordExpiration",
		"RoleId",
		"StrictAccountTypes",
		"UserName",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(manageraccount).Elem()

	return manageraccount.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetManagerAccount will get a ManagerAccount instance from the service.
func GetManagerAccount(c common.Client, uri string) (*ManagerAccount, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var manageraccount ManagerAccount
	err = json.NewDecoder(resp.Body).Decode(&manageraccount)
	if err != nil {
		return nil, err
	}

	manageraccount.SetClient(c)
	return &manageraccount, nil
}

// ListReferencedManagerAccounts gets the collection of ManagerAccount from
// a provided reference.
func ListReferencedManagerAccounts(c common.Client, link string) ([]*ManagerAccount, error) {
	var result []*ManagerAccount
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, manageraccountLink := range links.ItemLinks {
		manageraccount, err := GetManagerAccount(c, manageraccountLink)
		if err != nil {
			collectionError.Failures[manageraccountLink] = err
		} else {
			result = append(result, manageraccount)
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

// SNMPUserInfo shall contain the SNMP settings for an account.
type SNMPUserInfo struct {
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
}

// UnmarshalJSON unmarshals a SNMPUserInfo object from the raw JSON.
func (snmpuserinfo *SNMPUserInfo) UnmarshalJSON(b []byte) error {
	type temp SNMPUserInfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*snmpuserinfo = SNMPUserInfo(t.temp)

	// Extract the links to other entities for later

	return nil
}
