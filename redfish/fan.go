//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Fan shall represent the management properties for monitoring and management of cooling fans for a Redfish
// implementation.
type Fan struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// Description provides a description of this resource.
	Description string
	// HotPluggable shall indicate whether the device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Hot-pluggable devices can become operable without altering
	// the operational state of the underlying equipment. Devices that cannot be inserted or removed from equipment in
	// operation, or devices that cannot become operable without affecting the operational state of that equipment,
	// shall be not hot-pluggable.
	HotPluggable bool
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of this fan.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for producing the fan. This organization may
	// be the entity from whom the fan is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the model information as defined by the manufacturer for this fan.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for this fan.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region within the chassis with which this
	// fan is associated.
	PhysicalContext string
	// PowerWatts shall contain the total power, in watt units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'.
	PowerWatts SensorPowerExcerpt
	// Replaceable shall indicate whether this component can be independently replaced as allowed by the vendor's
	// replacement policy. A value of 'false' indicates the component needs to be replaced by policy, as part of
	// another component. If the 'LocationType' property of this component contains 'Embedded', this property shall
	// contain 'false'.
	Replaceable bool
	// SerialNumber shall contain the serial number as defined by the manufacturer for this fan.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for this fan.
	SparePartNumber string
	// SpeedPercent shall contain the fan speed, in percent units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Percent'.
	SpeedPercent SensorFanExcerpt
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Fan object from the raw JSON.
func (fan *Fan) UnmarshalJSON(b []byte) error {
	type temp Fan
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fan = Fan(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	fan.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (fan *Fan) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Fan)
	original.UnmarshalJSON(fan.rawData)

	readWriteFields := []string{
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fan).Elem()

	return fan.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetFan will get a Fan instance from the service.
func GetFan(c common.Client, uri string) (*Fan, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fan Fan
	err = json.NewDecoder(resp.Body).Decode(&fan)
	if err != nil {
		return nil, err
	}

	fan.SetClient(c)
	return &fan, nil
}

// ListReferencedFans gets the collection of Fan from
// a provided reference.
func ListReferencedFans(c common.Client, link string) ([]*Fan, error) {
	var result []*Fan
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, fanLink := range links.ItemLinks {
		fan, err := GetFan(c, fanLink)
		if err != nil {
			collectionError.Failures[fanLink] = err
		} else {
			result = append(result, fan)
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
	// CoolingChassis shall contain an array of links to resources of type Chassis that represent the chassis directly
	// cooled by this fan. This property shall not be present if the fan is only providing cooling to its containing
	// chassis.
	CoolingChassis []Chassis
	// CoolingChassis@odata.count
	CoolingChassisCount int `json:"CoolingChassis@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
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
