//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CurrentPeriod shall describe the memory metrics since last reset or ClearCurrentPeriod action for a switch.
type CurrentPeriod struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors of memory since reset.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors of memory since reset.
	UncorrectableECCErrorCount int
}

// UnmarshalJSON unmarshals a CurrentPeriod object from the raw JSON.
func (currentperiod *CurrentPeriod) UnmarshalJSON(b []byte) error {
	type temp CurrentPeriod
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*currentperiod = CurrentPeriod(t.temp)

	// Extract the links to other entities for later

	return nil
}

// InternalMemoryMetrics shall contain properties that describe the memory metrics for a switch.
type InternalMemoryMetrics struct {
	// CurrentPeriod shall contain properties that describe the metrics for the current period of memory for this
	// switch.
	CurrentPeriod string
	// LifeTime shall contain properties that describe the metrics for the lifetime of memory for this switch.
	LifeTime string
}

// UnmarshalJSON unmarshals a InternalMemoryMetrics object from the raw JSON.
func (internalmemorymetrics *InternalMemoryMetrics) UnmarshalJSON(b []byte) error {
	type temp InternalMemoryMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*internalmemorymetrics = InternalMemoryMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// LifeTime shall describe the memory metrics since manufacturing for a switch.
type LifeTime struct {
	// CorrectableECCErrorCount shall contain the number of the correctable errors for the lifetime of memory.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of the uncorrectable errors for the lifetime of memory.
	UncorrectableECCErrorCount int
}

// UnmarshalJSON unmarshals a LifeTime object from the raw JSON.
func (lifetime *LifeTime) UnmarshalJSON(b []byte) error {
	type temp LifeTime
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*lifetime = LifeTime(t.temp)

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

// SwitchMetrics shall represent the metrics for a switch device in a Redfish implementation.
type SwitchMetrics struct {
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
	// InternalMemoryMetrics shall contain properties that describe the memory metrics for a switch.
	InternalMemoryMetrics string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeErrors shall contain the PCIe errors associated with this switch.
	PCIeErrors string
}

// UnmarshalJSON unmarshals a SwitchMetrics object from the raw JSON.
func (switchmetrics *SwitchMetrics) UnmarshalJSON(b []byte) error {
	type temp SwitchMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*switchmetrics = SwitchMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetSwitchMetrics will get a SwitchMetrics instance from the service.
func GetSwitchMetrics(c common.Client, uri string) (*SwitchMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var switchmetrics SwitchMetrics
	err = json.NewDecoder(resp.Body).Decode(&switchmetrics)
	if err != nil {
		return nil, err
	}

	switchmetrics.SetClient(c)
	return &switchmetrics, nil
}

// ListReferencedSwitchMetricss gets the collection of SwitchMetrics from
// a provided reference.
func ListReferencedSwitchMetricss(c common.Client, link string) ([]*SwitchMetrics, error) {
	var result []*SwitchMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, switchmetricsLink := range links.ItemLinks {
		switchmetrics, err := GetSwitchMetrics(c, switchmetricsLink)
		if err != nil {
			collectionError.Failures[switchmetricsLink] = err
		} else {
			result = append(result, switchmetrics)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
