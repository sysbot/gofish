//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AccelerationFunctionType is
type AccelerationFunctionType string

const (
	// EncryptionAccelerationFunctionType An encryption function.
	EncryptionAccelerationFunctionType AccelerationFunctionType = "Encryption"
	// CompressionAccelerationFunctionType A compression function.
	CompressionAccelerationFunctionType AccelerationFunctionType = "Compression"
	// PacketInspectionAccelerationFunctionType A packet inspection function.
	PacketInspectionAccelerationFunctionType AccelerationFunctionType = "PacketInspection"
	// PacketSwitchAccelerationFunctionType A packet switch function.
	PacketSwitchAccelerationFunctionType AccelerationFunctionType = "PacketSwitch"
	// SchedulerAccelerationFunctionType A scheduler function.
	SchedulerAccelerationFunctionType AccelerationFunctionType = "Scheduler"
	// AudioProcessingAccelerationFunctionType An audio processing function.
	AudioProcessingAccelerationFunctionType AccelerationFunctionType = "AudioProcessing"
	// VideoProcessingAccelerationFunctionType A video processing function.
	VideoProcessingAccelerationFunctionType AccelerationFunctionType = "VideoProcessing"
	// OEMAccelerationFunctionType An OEM-defined acceleration function.
	OEMAccelerationFunctionType AccelerationFunctionType = "OEM"
)

// AccelerationFunction shall represent the acceleration function that a processor implements in a Redfish
// implementation. This can include functions such as audio processing, compression, encryption, packet inspection,
// packet switching, scheduling, or video processing.
type AccelerationFunction struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccelerationFunctionType shall contain the string that identifies the acceleration function type.
	AccelerationFunctionType AccelerationFunctionType
	// Actions shall contain the available actions for this Resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// FpgaReconfigurationSlots shall contain an array of the FPGA reconfiguration slot identifiers that this
	// acceleration function occupies.
	FpgaReconfigurationSlots []string
	// Links shall contain links to Resources that are related to but are not contained by, or subordinate to, this
	// Resource.
	Links string
	// Manufacturer shall contain a string that identifies the manufacturer of the acceleration function.
	Manufacturer string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerWatts shall contain the total acceleration function power consumption, in watts.
	PowerWatts string
	// Status shall contain any status or health properties of the Resource.
	Status common.Status
	// UUID shall contain a UUID for the acceleration function. RFC4122 describes methods that can create the value.
	// The value should be considered to be opaque. Client software should only treat the overall value as a UUID and
	// should not interpret any sub-fields within the UUID.
	UUID string
	// Version shall describe the acceleration function version.
	Version string
}

// UnmarshalJSON unmarshals a AccelerationFunction object from the raw JSON.
func (accelerationfunction *AccelerationFunction) UnmarshalJSON(b []byte) error {
	type temp AccelerationFunction
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*accelerationfunction = AccelerationFunction(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetAccelerationFunction will get a AccelerationFunction instance from the service.
func GetAccelerationFunction(c common.Client, uri string) (*AccelerationFunction, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var accelerationfunction AccelerationFunction
	err = json.NewDecoder(resp.Body).Decode(&accelerationfunction)
	if err != nil {
		return nil, err
	}

	accelerationfunction.SetClient(c)
	return &accelerationfunction, nil
}

// ListReferencedAccelerationFunctions gets the collection of AccelerationFunction from
// a provided reference.
func ListReferencedAccelerationFunctions(c common.Client, link string) ([]*AccelerationFunction, error) {
	var result []*AccelerationFunction
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, accelerationfunctionLink := range links.ItemLinks {
		accelerationfunction, err := GetAccelerationFunction(c, accelerationfunctionLink)
		if err != nil {
			collectionError.Failures[accelerationfunctionLink] = err
		} else {
			result = append(result, accelerationfunction)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Links shall contain links to Resources that are related to but are not contained by, or subordinate to, this
// Resource.
type Links struct {
	// Endpoints shall contain an array of links to Resources of the Endpoint type that are associated with this
	// acceleration function.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeFunctions shall contain an array of links of the PCIeFunction type that represent the PCIe functions
	// associated with this acceleration function.
	PCIeFunctions []PCIeFunction
	// PCIeFunctions@odata.count
	PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
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

// OemActions shall contain the available OEM-specific actions for this Resource.
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
