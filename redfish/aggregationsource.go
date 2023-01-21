//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AggregationType is
type AggregationType string

const (
	// NotificationsOnlyAggregationType shall indicate that the aggregator is only aggregating notifications or events
	// from the aggregation source according to the connection method used. This value shall not be used with
	// connection methods that do not include notifications.
	NotificationsOnlyAggregationType AggregationType = "NotificationsOnly"
	// FullAggregationType shall indicate that the aggregator is performing full aggregation according to the
	// connection method without any limitation such as only receiving notifications.
	FullAggregationType AggregationType = "Full"
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

// AggregationSource shall represent an aggregation source for a Redfish implementation.
type AggregationSource struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AggregationType shall contain the type of aggregation used for the connection method towards the aggregation
	// source. If this property is not present, the value shall be assumed to be 'Complete'.
	AggregationType string
	// Description provides a description of this resource.
	Description string
	// HostName shall contain the URI of the system to be aggregated. This property shall not be required when the
	// aggregation source is configured to only receive notifications from the aggregated system and the
	// AggregationType property contains the value 'NotificationsOnly'.
	HostName string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Password shall contain a password for accessing the aggregation source. The value shall be 'null' in responses.
	Password string
	// SNMP shall contain the SNMP settings of the aggregation source.
	SNMP SNMPSettings
	// UserName shall contain the user name for accessing the aggregation source.
	UserName string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AggregationSource object from the raw JSON.
func (aggregationsource *AggregationSource) UnmarshalJSON(b []byte) error {
	type temp AggregationSource
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*aggregationsource = AggregationSource(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	aggregationsource.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (aggregationsource *AggregationSource) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AggregationSource)
	original.UnmarshalJSON(aggregationsource.rawData)

	readWriteFields := []string{
		"AggregationType",
		"HostName",
		"Password",
		"UserName",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(aggregationsource).Elem()

	return aggregationsource.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetAggregationSource will get a AggregationSource instance from the service.
func GetAggregationSource(c common.Client, uri string) (*AggregationSource, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var aggregationsource AggregationSource
	err = json.NewDecoder(resp.Body).Decode(&aggregationsource)
	if err != nil {
		return nil, err
	}

	aggregationsource.SetClient(c)
	return &aggregationsource, nil
}

// ListReferencedAggregationSources gets the collection of AggregationSource from
// a provided reference.
func ListReferencedAggregationSources(c common.Client, link string) ([]*AggregationSource, error) {
	var result []*AggregationSource
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, aggregationsourceLink := range links.ItemLinks {
		aggregationsource, err := GetAggregationSource(c, aggregationsourceLink)
		if err != nil {
			collectionError.Failures[aggregationsourceLink] = err
		} else {
			result = append(result, aggregationsource)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// ConnectionMethod shall contain an array of links to resources of type ConnectionMethod that are used to connect
	// to the aggregation source.
	ConnectionMethod ConnectionMethod
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ResourcesAccessed shall contain an array of links to the resources added to the service through the aggregation
	// source. It is recommended that this be the minimal number of properties needed to find the resources that would
	// be lost when the aggregation source is deleted. For example, this could be the pointers to the members of the
	// root level collections or the manager of a BMC.
	ResourcesAccessed []Resource
	// ResourcesAccessed@odata.count
	ResourcesAccessedCount int `json:"ResourcesAccessed@odata.count"`
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

// SNMPSettings shall contain the settings for an SNMP aggregation source.
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
	// TrapCommunity shall contain the SNMP trap community string. The value shall be 'null' in responses. Services may
	// provide a common trap community if not specified by the client when creating the aggregation source.
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
