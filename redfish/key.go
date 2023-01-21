//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// KeyType is
type KeyType string

const (
	// NVMeoFKeyType shall indicate the format of the key is defined by one of the NVMe specifications.
	NVMeoFKeyType KeyType = "NVMeoF"
	// SSHKeyType shall indicate the format of the key is defined by one of the SSH public key formats as defined in,
	// but not limited to, RFC4253, RFC4716, or RFC8709.
	SSHKeyType KeyType = "SSH"
)

// NVMeoFSecureHashType is This enumeration shall list the NVMe secure hash algorithms that a key is allowed to
// use.
type NVMeoFSecureHashType string

const (
	// SHA256NVMeoFSecureHashType shall indicate the SHA-256 hash function as defined by the 'DH-HMAC-CHAP hash
	// function identifiers' figure in the NVMe Base Specification.
	SHA256NVMeoFSecureHashType NVMeoFSecureHashType = "SHA256"
	// SHA384NVMeoFSecureHashType shall indicate the SHA-384 hash function as defined by the 'DH-HMAC-CHAP hash
	// function identifiers' figure in the NVMe Base Specification.
	SHA384NVMeoFSecureHashType NVMeoFSecureHashType = "SHA384"
	// SHA512NVMeoFSecureHashType shall indicate the SHA-512 hash function as defined by the 'DH-HMAC-CHAP hash
	// function identifiers' figure in the NVMe Base Specification.
	SHA512NVMeoFSecureHashType NVMeoFSecureHashType = "SHA512"
)

// NVMeoFSecurityProtocolType is This enumeration shall list the NVMe security protocols that a key protects.
type NVMeoFSecurityProtocolType string

const (
	// DHHCNVMeoFSecurityProtocolType shall indicate the Diffie-Hellman Hashed Message Authentication Code Challenge
	// Handshake Authentication Protocol (DH-HMAC-CHAP) as defined by the NVMe Base Specification.
	DHHCNVMeoFSecurityProtocolType NVMeoFSecurityProtocolType = "DHHC"
	// TLSPSKNVMeoFSecurityProtocolType shall indicate Transport Layer Security Pre-Shared Key (TLS PSK) as defined by
	// the NVMe TCP Transport Specification.
	TLSPSKNVMeoFSecurityProtocolType NVMeoFSecurityProtocolType = "TLS_PSK"
	// OEMNVMeoFSecurityProtocolType shall indicate an OEM-defined security protocol. The OEMSecurityProtocolType
	// property shall contain the specific OEM protocol.
	OEMNVMeoFSecurityProtocolType NVMeoFSecurityProtocolType = "OEM"
)

// Key shall represent a key for a Redfish implementation.
type Key struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// KeyString shall contain the key, and the format shall follow the requirements specified by the KeyType property
	// value.
	KeyString string
	// KeyType shall contain the format type for the key.
	KeyType KeyType
	// NVMeoF shall contain NVMe-oF specific properties for this key. This property shall be present if KeyType
	// contains the value 'NVMeoF'.
	NVMeoF NVMeoF
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// UserDescription shall contain a user-provided string that describes the key.
	UserDescription string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Key object from the raw JSON.
func (key *Key) UnmarshalJSON(b []byte) error {
	type temp Key
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*key = Key(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	key.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (key *Key) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Key)
	original.UnmarshalJSON(key.rawData)

	readWriteFields := []string{
		"UserDescription",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(key).Elem()

	return key.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetKey will get a Key instance from the service.
func GetKey(c common.Client, uri string) (*Key, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var key Key
	err = json.NewDecoder(resp.Body).Decode(&key)
	if err != nil {
		return nil, err
	}

	key.SetClient(c)
	return &key, nil
}

// ListReferencedKeys gets the collection of Key from
// a provided reference.
func ListReferencedKeys(c common.Client, link string) ([]*Key, error) {
	var result []*Key
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, keyLink := range links.ItemLinks {
		key, err := GetKey(c, keyLink)
		if err != nil {
			collectionError.Failures[keyLink] = err
		} else {
			result = append(result, key)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// NVMeoF shall contain NVMe-oF specific properties for a key.
type NVMeoF struct {
	// HostKeyId shall contain the value of the Id property of the Key resource representing the host key paired with
	// this target key. An empty string shall indicate the key is not paired. This property shall be absent for host
	// keys.
	HostKeyId string
	// NQN shall contain the NVMe Qualified Name (NQN) of the host or target subsystem associated with this key. The
	// value of this property shall follow the NQN format defined by the NVMe Base Specification.
	NQN string
	// OEMSecurityProtocolType shall contain the OEM-defined security protocol that this key uses. The value shall be
	// derived from the contents of the KeyString property. This property shall be present if SecurityProtocolType
	// contains the value 'OEM'.
	OEMSecurityProtocolType string
	// SecureHashAllowList shall contain the secure hash algorithms allowed with the usage of this key. An empty list
	// or the absence of this property shall indicate any secure hash algorithms are allowed with this key.
	SecureHashAllowList []NVMeoFSecureHashType
	// SecurityProtocolType shall contain the security protocol that this key uses. The value shall be derived from the
	// contents of the KeyString property.
	SecurityProtocolType NVMeoFSecurityProtocolType
}

// UnmarshalJSON unmarshals a NVMeoF object from the raw JSON.
func (nvmeof *NVMeoF) UnmarshalJSON(b []byte) error {
	type temp NVMeoF
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*nvmeof = NVMeoF(t.temp)

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
