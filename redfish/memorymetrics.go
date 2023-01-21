//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AlarmTrips shall contain properties that describe the types of alarms that have been raised by the memory. These
// alarms shall be reset when the system resets. Note that if they are re-discovered they can be reasserted.
type AlarmTrips struct {
	// AddressParityError shall indicate whether an address parity error was detected that a retry could not correct.
	AddressParityError bool
	// CorrectableECCError shall indicate whether the correctable error threshold crossing alarm trip was detected.
	CorrectableECCError bool
	// SpareBlock shall indicate whether the spare block capacity crossing alarm trip was detected.
	SpareBlock bool
	// Temperature shall indicates whether a temperature threshold alarm trip was detected.
	Temperature bool
	// UncorrectableECCError shall indicate whether the uncorrectable error threshold alarm trip was detected.
	UncorrectableECCError bool
}

// UnmarshalJSON unmarshals a AlarmTrips object from the raw JSON.
func (alarmtrips *AlarmTrips) UnmarshalJSON(b []byte) error {
	type temp AlarmTrips
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*alarmtrips = AlarmTrips(t.temp)

	// Extract the links to other entities for later

	return nil
}

// CurrentPeriod shall describe the memory metrics since last system reset or ClearCurrentPeriod action.
type CurrentPeriod struct {
	// BlocksRead shall contain the number of blocks read since reset. When this resource is subordinate to the
	// MemorySummary object, this property shall be the sum of BlocksRead over all memory.
	BlocksRead int
	// BlocksWritten shall contain the number of blocks written since reset. When this resource is subordinate to the
	// MemorySummary object, this property shall be the sum of BlocksWritten over all memory.
	BlocksWritten int
	// CorrectableECCErrorCount shall contain the number of correctable errors since reset. When this resource is
	// subordinate to the MemorySummary object, this property shall be the sum of CorrectableECCErrorCount over all
	// memory.
	CorrectableECCErrorCount int
	// IndeterminateCorrectableErrorCount shall contain the number of indeterminate correctable errors since reset.
	// Since the error origin is indeterminate, the same error can be duplicated across multiple MemoryMetrics
	// resources. When this resource is subordinate to the MemorySummary object, this property shall be the sum of
	// indeterminate correctable errors across all memory without duplication, which may not be the sum of all
	// IndeterminateCorrectableErrorCount properties over all memory.
	IndeterminateCorrectableErrorCount int
	// IndeterminateUncorrectableErrorCount shall contain the number of indeterminate uncorrectable errors since reset.
	// Since the error origin is indeterminate, the same error can be duplicated across multiple MemoryMetrics
	// resources. When this resource is subordinate to the MemorySummary object, this property shall be the sum of
	// indeterminate uncorrectable errors across all memory without duplication, which may not be the sum of all
	// IndeterminateUncorrectableErrorCount properties over all memory.
	IndeterminateUncorrectableErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors since reset. When this resource is
	// subordinate to the MemorySummary object, this property shall be the sum of UncorrectableECCErrorCount over all
	// memory.
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

// HealthData shall contain properties that describe the HealthData metrics for this resource.
type HealthData struct {
	// AlarmTrips shall contain properties describe the types of alarms that have been raised by the memory. When this
	// resource is subordinate to the MemorySummary object, this property shall indicate whether an alarm of a given
	// type have been raised by any area of memory.
	AlarmTrips string
	// DataLossDetected shall indicate whether data loss was detected. When this resource is subordinate to the
	// MemorySummary object, this property shall indicate whether any data loss was detected in any area of memory.
	DataLossDetected bool
	// LastShutdownSuccess shall indicate whether the last shutdown succeeded.
	LastShutdownSuccess bool
	// PerformanceDegraded shall indicate whether performance has degraded. When this resource is subordinate to the
	// MemorySummary object, this property shall indicate whether degraded performance mode status is detected in any
	// area of memory.
	PerformanceDegraded bool
	// PredictedMediaLifeLeftPercent shall contain an indicator of the percentage of life remaining in the media.
	PredictedMediaLifeLeftPercent float64
	// RemainingSpareBlockPercentage shall contain the remaining spare blocks as a percentage. When this resource is
	// subordinate to the MemorySummary object, this property shall be the RemainingSpareBlockPercentage over all
	// memory.
	RemainingSpareBlockPercentage float64
}

// UnmarshalJSON unmarshals a HealthData object from the raw JSON.
func (healthdata *HealthData) UnmarshalJSON(b []byte) error {
	type temp HealthData
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*healthdata = HealthData(t.temp)

	// Extract the links to other entities for later

	return nil
}

// LifeTime shall describe the memory metrics since manufacturing.
type LifeTime struct {
	// BlocksRead shall contain the number of blocks read for the lifetime of the memory. When this resource is
	// subordinate to the MemorySummary object, this property shall be the sum of BlocksRead over all memory.
	BlocksRead int
	// BlocksWritten shall contain the number of blocks written for the lifetime of the memory. When this resource is
	// subordinate to the MemorySummary object, this property shall be the sum of BlocksWritten over all memory.
	BlocksWritten int
	// CorrectableECCErrorCount shall contain the number of the correctable errors for the lifetime of the memory. When
	// this resource is subordinate to the MemorySummary object, this property shall be the sum of
	// CorrectableECCErrorCount over all memory.
	CorrectableECCErrorCount int
	// IndeterminateCorrectableErrorCount shall contain the number of indeterminate correctable errors for the lifetime
	// of the memory. Since the error origin is indeterminate, the same error can be duplicated across multiple
	// MemoryMetrics resources. When this resource is subordinate to the MemorySummary object, this property shall be
	// the sum of indeterminate correctable errors across all memory without duplication, which may not bey the sum of
	// all IndeterminateCorrectableErrorCount properties over all memory.
	IndeterminateCorrectableErrorCount int
	// IndeterminateUncorrectableErrorCount shall contain the number of indeterminate uncorrectable errors for the
	// lifetime of the memory. Since the error origin is indeterminate, the same error can be duplicated across
	// multiple MemoryMetrics resources. When this resource is subordinate to the MemorySummary object, this property
	// shall be the sum of indeterminate uncorrectable errors across all memory without duplication, which may not be
	// the sum of all IndeterminateUncorrectableErrorCount properties over all memory.
	IndeterminateUncorrectableErrorCount int
	// UncorrectableECCErrorCount shall contain the number of the uncorrectable errors for the lifetime of the memory.
	// When this resource is subordinate to the MemorySummary object, this property shall be the sum of
	// UncorrectableECCErrorCount over all memory.
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

// MemoryMetrics shall contain the memory metrics for a memory device or system memory summary in a Redfish
// implementation.
type MemoryMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// BandwidthPercent shall contain memory bandwidth utilization as a percentage. When this resource is subordinate
	// to the MemorySummary object, this property shall be the memory bandwidth utilization over all memory as a
	// percentage.
	BandwidthPercent float64
	// BlockSizeBytes shall contain the block size, in bytes, of all structure elements. When this resource is
	// subordinate to the MemorySummary object, this property is not applicable.
	BlockSizeBytes int
	// CurrentPeriod shall contain properties that describe the memory metrics for the current period.
	CurrentPeriod string
	// Description provides a description of this resource.
	Description string
	// HealthData shall contain properties that describe the health data memory metrics for the memory.
	HealthData string
	// LifeTime shall contain properties that describe the memory metrics for the lifetime of the memory.
	LifeTime string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingSpeedMHz shall contain the operating speed of memory in MHz or MT/s (mega-transfers per second) as
	// reported by the memory device. Memory devices that operate at their bus speed shall report the operating speed
	// in MHz (bus speed), while memory devices that transfer data faster than their bus speed, such as DDR memory,
	// shall report the operating speed in MT/s (mega-transfers/second). The reported value shall match the
	// conventionally reported values for the technology used by the memory device.
	OperatingSpeedMHz int
}

// UnmarshalJSON unmarshals a MemoryMetrics object from the raw JSON.
func (memorymetrics *MemoryMetrics) UnmarshalJSON(b []byte) error {
	type temp MemoryMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorymetrics = MemoryMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetMemoryMetrics will get a MemoryMetrics instance from the service.
func GetMemoryMetrics(c common.Client, uri string) (*MemoryMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memorymetrics MemoryMetrics
	err = json.NewDecoder(resp.Body).Decode(&memorymetrics)
	if err != nil {
		return nil, err
	}

	memorymetrics.SetClient(c)
	return &memorymetrics, nil
}

// ListReferencedMemoryMetricss gets the collection of MemoryMetrics from
// a provided reference.
func ListReferencedMemoryMetricss(c common.Client, link string) ([]*MemoryMetrics, error) {
	var result []*MemoryMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, memorymetricsLink := range links.ItemLinks {
		memorymetrics, err := GetMemoryMetrics(c, memorymetricsLink)
		if err != nil {
			collectionError.Failures[memorymetricsLink] = err
		} else {
			result = append(result, memorymetrics)
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
