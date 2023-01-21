//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)


// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
    // Chassis shall contain a link to a resource of type Chassis with which this switch is associated.
    Chassis  string
    // Endpoints shall contain an array of links to resources of type Endpoint with which this switch is associated.
    Endpoints  []Endpoint
    // Endpoints@odata.count
    EndpointsCount  int `json:"Endpoints@odata.count"`
    // ManagedBy shall contain an array of links to resources of type Manager with which this switch is associated.
    ManagedBy  []Manager
    // ManagedBy@odata.count
    ManagedByCount  int `json:"ManagedBy@odata.count"`
    // Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
// Redfish Specification-described requirements.
    OEM  json.RawMessage `json:"Oem"`
    // PCIeDevice shall contain a link to a resource of type PCIeDevice that represents the PCIe device providing this
// switch.
    PCIeDevice  PCIeDevice
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


// Switch This resource contains a switch for a Redfish implementation.
type Switch struct {
    common.Entity
    // ODataContext is the odata context.
    ODataContext  string `json:"@odata.context"`
    // ODataEtag is the odata etag.
    ODataEtag  string `json:"@odata.etag"`
    // ODataType is the odata type.
    ODataType  string `json:"@odata.type"`
    // Actions shall contain the available actions for this resource.
    Actions  string
    // AssetTag shall contain the user-assigned asset tag, which is an identifying string that tracks the drive for
// inventory purposes.
    AssetTag  string
    // Certificates shall contain a link to a resource collection of type CertificateCollection that contains
// certificates for device identity and attestation.
    Certificates  string
    // CurrentBandwidthGbps shall contain the internal bandwidth of this switch currently negotiated and running.
    CurrentBandwidthGbps  float64
    // Description provides a description of this resource.
    Description  string
    // DomainID shall contain The domain ID for this switch. This property has a scope of uniqueness within the fabric
// of which the switch is a member.
    DomainID  int
    // Enabled shall indicate if this switch is enabled.
    Enabled  string
    // EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
// metrics for this switch.
    EnvironmentMetrics  string
    // FirmwareVersion shall contain the firmware version as defined by the manufacturer for the associated switch.
    FirmwareVersion  string
    // IsManaged shall indicate whether this switch is in a managed or unmanaged state.
    IsManaged  bool
    // Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
    Links  string
    // Location shall contain location information of the associated switch.
    Location  string
    // LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
// reflect the implementation of the locating function.
    LocationIndicatorActive  bool
    // LogServices shall contain a link to a resource collection of type LogServiceCollection.
    LogServices  string
    // Manufacturer shall contain the name of the organization responsible for producing the switch. This organization
// may be the entity from which the switch is purchased, but this is not necessarily true.
    Manufacturer  string
    // MaxBandwidthGbps shall contain the maximum internal bandwidth this switch is capable of being configured. If
// capable of autonegotiation, the switch shall attempt to negotiate to the specified maximum bandwidth.
    MaxBandwidthGbps  float64
    // Metrics shall contain a link to the metrics associated with this switch.
    Metrics  string
    // Model shall contain the manufacturer-provided model information of this switch.
    Model  string
    // Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
// Redfish Specification-described requirements.
    OEM  json.RawMessage `json:"Oem"`
    // PartNumber shall contain the manufacturer-provided part number for the switch.
    PartNumber  string
    // Ports shall contain a link to a resource collection of type PortCollection.
    Ports  string
    // PowerState shall contain the power state of the switch.
    PowerState  PowerState
    // Redundancy shall contain an array that shows how this switch is grouped with other switches for form redundancy
// sets.
    Redundancy  []Redundancy
    // Redundancy@odata.count
    RedundancyCount  int `json:"Redundancy@odata.count"`
    // SKU shall contain the SKU number for this switch.
    SKU  string
    // SerialNumber shall contain a manufacturer-allocated number that identifies the switch.
    SerialNumber  string
    // Status shall contain any status or health properties of the resource.
    Status  common.Status
    // SupportedProtocols shall contain an array of protocols this switch supports. If the value of SwitchType is
// 'MultiProtocol', this property shall be required.
    SupportedProtocols  []Protocol
    // SwitchType shall contain the protocol being sent over this switch. For a switch that supports multiple
// protocols, the value should be 'MultiProtocol' and the SupportedProtocols property should be used to describe
// the supported protocols.
    SwitchType  Protocol
    // TotalSwitchWidth shall contain the number of physical transport lanes, phys, or other physical transport links
// that this switch contains. For PCIe, this value shall be the lane count.
    TotalSwitchWidth  int
    // UUID shall contain a universal unique identifier number for the switch.
    UUID  string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Switch object from the raw JSON.
func (switch *Switch) UnmarshalJSON(b []byte) error {
    type temp Switch
    var t struct {
        temp
    }

    err := json.Unmarshal(b, &t)
    if err != nil {
        return err
    }

    *switch = Switch(t.temp)

    // Extract the links to other entities for later


	// This is a read/write object, so we need to save the raw object data for later
	switch.rawData = b

    return nil
}
// Update commits updates to this object's properties to the running system.
func (switch *Switch) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Switch)
	original.UnmarshalJSON(switch.rawData)

	readWriteFields := []string{
        "AssetTag",
        "Enabled",
        "IsManaged",
        "LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(switch).Elem()

	return switch.Entity.Update(originalElement, currentElement, readWriteFields)
}


// GetSwitch will get a Switch instance from the service.
func GetSwitch(c common.Client, uri string) (*Switch, error) {
    resp, err := c.Get(uri)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var switch Switch
    err = json.NewDecoder(resp.Body).Decode(&switch)
    if err != nil {
        return nil, err
    }

    switch.SetClient(c)
    return &switch, nil
}

// ListReferencedSwitchs gets the collection of Switch from
// a provided reference.
func ListReferencedSwitchs(c common.Client, link string) ([]*Switch, error) {
    var result []*Switch
    if link == "" {
        return result, nil
    }

    links, err := common.GetCollection(c, link)
    if err != nil {
        return result, err
    }

    collectionError := common.NewCollectionError()
    for _, switchLink := range links.ItemLinks {
        switch, err := GetSwitch(c, switchLink)
        if err != nil {
            collectionError.Failures[switchLink] = err
        } else {
            result = append(result, switch)
        }
    }

    if collectionError.Empty() {
        return result, nil
    } else {
        return result, collectionError
    }
}



