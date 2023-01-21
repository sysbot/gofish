//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// BaseSpeedPrioritySettings shall specify the clock speed for a set of cores.
type BaseSpeedPrioritySettings struct {
	// BaseSpeedMHz shall contain the clock speed to configure the set of cores in MHz.
	BaseSpeedMHz int
	// CoreCount shall contain the number of cores to configure with the speed specified by the BaseSpeedMHz property.
	// The sum of all CoreCount properties shall equal the value of the TotalAvailableCoreCount property.
	CoreCount int
	// CoreIDs shall contain an array identifying the cores to configure with the speed specified by the BaseSpeedMHz
	// property. The length of the array shall equal the value of the CoreCount property.
	CoreIDs []string
}

// UnmarshalJSON unmarshals a BaseSpeedPrioritySettings object from the raw JSON.
func (basespeedprioritysettings *BaseSpeedPrioritySettings) UnmarshalJSON(b []byte) error {
	type temp BaseSpeedPrioritySettings
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*basespeedprioritysettings = BaseSpeedPrioritySettings(t.temp)

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

// OperatingConfig shall represent an operational configuration for a processor in the Redfish Specification.
type OperatingConfig struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// BaseSpeedMHz shall contain the base (nominal) clock speed of the processor in MHz.
	BaseSpeedMHz int
	// BaseSpeedPrioritySettings shall contain an array of objects that specify the clock speed for sets of cores when
	// the configuration is operational.
	BaseSpeedPrioritySettings []BaseSpeedPrioritySettings
	// Description provides a description of this resource.
	Description string
	// MaxJunctionTemperatureCelsius shall contain the maximum temperature of the junction in degrees Celsius.
	MaxJunctionTemperatureCelsius int
	// MaxSpeedMHz shall contain the maximum clock speed to which the processor can be configured in MHz.
	MaxSpeedMHz int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// TDPWatts shall contain the thermal design point of the processor in watts.
	TDPWatts int
	// TotalAvailableCoreCount shall contain the number of cores in the processor that can be configured.
	TotalAvailableCoreCount int
	// TurboProfile shall contain an array of objects that specify the turbo profile for a set of active cores.
	TurboProfile []TurboProfileDatapoint
}

// UnmarshalJSON unmarshals a OperatingConfig object from the raw JSON.
func (operatingconfig *OperatingConfig) UnmarshalJSON(b []byte) error {
	type temp OperatingConfig
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*operatingconfig = OperatingConfig(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetOperatingConfig will get a OperatingConfig instance from the service.
func GetOperatingConfig(c common.Client, uri string) (*OperatingConfig, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var operatingconfig OperatingConfig
	err = json.NewDecoder(resp.Body).Decode(&operatingconfig)
	if err != nil {
		return nil, err
	}

	operatingconfig.SetClient(c)
	return &operatingconfig, nil
}

// ListReferencedOperatingConfigs gets the collection of OperatingConfig from
// a provided reference.
func ListReferencedOperatingConfigs(c common.Client, link string) ([]*OperatingConfig, error) {
	var result []*OperatingConfig
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, operatingconfigLink := range links.ItemLinks {
		operatingconfig, err := GetOperatingConfig(c, operatingconfigLink)
		if err != nil {
			collectionError.Failures[operatingconfigLink] = err
		} else {
			result = append(result, operatingconfig)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// TurboProfileDatapoint shall specify the turbo profile for a set of active cores.
type TurboProfileDatapoint struct {
	// ActiveCoreCount shall contain the number of cores to be configured with the maximum turbo clock speed. The value
	// shall be less than or equal the TotalAvailableCoreCount property.
	ActiveCoreCount int
	// MaxSpeedMHz shall contain the maximum turbo clock speed that correspond to the number of active cores in MHz.
	MaxSpeedMHz int
}

// UnmarshalJSON unmarshals a TurboProfileDatapoint object from the raw JSON.
func (turboprofiledatapoint *TurboProfileDatapoint) UnmarshalJSON(b []byte) error {
	type temp TurboProfileDatapoint
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*turboprofiledatapoint = TurboProfileDatapoint(t.temp)

	// Extract the links to other entities for later

	return nil
}
