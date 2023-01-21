//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CStateResidency shall contain properties that describe the C-state residency of the processor or core.
type CStateResidency struct {
	// Level shall contain the C-state level, such as C0, C1, or C2. When this resource is subordinate to the
	// ProcessorSummary object, this property is not applicable.
	Level string
	// ResidencyPercent shall contain the percentage of time that the processor or core has spent in this particular
	// level of C-state. When this resource is subordinate to the ProcessorSummary object, this property is not
	// applicable.
	ResidencyPercent float64
}

// UnmarshalJSON unmarshals a CStateResidency object from the raw JSON.
func (cstateresidency *CStateResidency) UnmarshalJSON(b []byte) error {
	type temp CStateResidency
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*cstateresidency = CStateResidency(t.temp)

	// Extract the links to other entities for later

	return nil
}

// CacheMetrics shall contain properties that describe cache metrics of a processor or core.
type CacheMetrics struct {
	// CacheMiss shall contain the number of cache line misses of the processor or core in millions.
	CacheMiss float64
	// CacheMissesPerInstruction shall contain the number of cache misses per instruction of the processor or core.
	CacheMissesPerInstruction float64
	// HitRatio shall contain the cache hit ratio of the processor or core.
	HitRatio float64
	// Level shall contain the level of the cache in the processor or core.
	Level string
	// OccupancyBytes shall contain the total cache occupancy of the processor or core in bytes.
	OccupancyBytes int
	// OccupancyPercent shall contain the total cache occupancy percentage of the processor or core.
	OccupancyPercent float64
}

// UnmarshalJSON unmarshals a CacheMetrics object from the raw JSON.
func (cachemetrics *CacheMetrics) UnmarshalJSON(b []byte) error {
	type temp CacheMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*cachemetrics = CacheMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// CacheMetricsTotal shall contain properties that describe the metrics for all of the cache memory for a
// processor.
type CacheMetricsTotal struct {
	// CurrentPeriod shall contain properties that describe the metrics for the current period of cache memory for this
	// processor.
	CurrentPeriod string
	// LifeTime shall contain properties that describe the metrics for the lifetime of cache memory for this processor.
	LifeTime string
}

// UnmarshalJSON unmarshals a CacheMetricsTotal object from the raw JSON.
func (cachemetricstotal *CacheMetricsTotal) UnmarshalJSON(b []byte) error {
	type temp CacheMetricsTotal
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*cachemetricstotal = CacheMetricsTotal(t.temp)

	// Extract the links to other entities for later

	return nil
}

// CoreMetrics shall contain properties that describe the cores of a processor.
type CoreMetrics struct {
	// CStateResidency shall contain properties that describe the C-state residency of this core in the processor.
	CStateResidency []CStateResidency
	// CoreCache shall contain properties that describe the cache metrics of this core in the processor.
	CoreCache []CacheMetrics
	// CoreId shall contain the identifier of the core within the processor.
	CoreId string
	// CorrectableCoreErrorCount shall contain the number of correctable core errors, such as TLB or cache errors. When
	// this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableCoreErrorCount over all processors.
	CorrectableCoreErrorCount int
	// CorrectableOtherErrorCount shall contain the number of the correctable errors of all other components. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableOtherErrorCount over all processors.
	CorrectableOtherErrorCount int
	// IOStallCount shall contain the number of stalled cycles due to I/O operations of this core in the processor.
	IOStallCount int
	// InstructionsPerCycle shall contain the number of instructions per clock cycle of this core in the processor.
	InstructionsPerCycle float64
	// MemoryStallCount shall contain the number of stalled cycles due to memory operations of this core in the
	// processor.
	MemoryStallCount int
	// UncorrectableCoreErrorCount shall contain the number of uncorrectable core errors, such as TLB or cache errors.
	// When this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableCoreErrorCount over all processors.
	UncorrectableCoreErrorCount int
	// UncorrectableOtherErrorCount shall contain the number of the uncorrectable errors of all other components. When
	// this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableOtherErrorCount over all processors.
	UncorrectableOtherErrorCount int
	// UnhaltedCycles shall contain the number of unhalted cycles of this core in the processor.
	UnhaltedCycles float64
}

// UnmarshalJSON unmarshals a CoreMetrics object from the raw JSON.
func (coremetrics *CoreMetrics) UnmarshalJSON(b []byte) error {
	type temp CoreMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*coremetrics = CoreMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// CurrentPeriod shall describe the cache memory metrics since last system reset or ClearCurrentPeriod action for a
// processor.
type CurrentPeriod struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors of cache memory since reset or
	// ClearCurrentPeriod action for this processor. When this resource is subordinate to the ProcessorSummary object,
	// this property shall be the sum of CorrectableECCErrorCount over all processors.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors of cache memory since reset or
	// ClearCurrentPeriod action for this processor. When this resource is subordinate to the ProcessorSummary object,
	// this property shall be the sum of UncorrectableECCErrorCount over all processors.
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

// LifeTime shall describe the cache memory metrics since manufacturing for a processor.
type LifeTime struct {
	// CorrectableECCErrorCount shall contain the number of the correctable errors for the lifetime of cache memory.
	// When this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableECCErrorCount over all processors.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of the uncorrectable errors for the lifetime of cache
	// memory. When this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableECCErrorCount over all processors.
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

// ProcessorMetrics This resource contains the processor metrics for a single processor in a Redfish
// implementation.
type ProcessorMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// BandwidthPercent shall contain the bandwidth usage of the processor as a percentage. When this resource is
	// subordinate to the ProcessorSummary object, this property shall be the CPU utilization over all processors as a
	// percentage.
	BandwidthPercent float64
	// Cache shall contain properties that describe this processor's cache. When this resource is subordinate to the
	// ProcessorSummary object, this property is not applicable.
	Cache []CacheMetrics
	// CacheMetricsTotal shall contain properties that describe the metrics for all of the cache memory of this
	// processor.
	CacheMetricsTotal string
	// CoreMetrics shall contain properties that describe the cores of this processor. When this resource is
	// subordinate to the ProcessorSummary object, this property is not applicable.
	CoreMetrics []CoreMetrics
	// CoreVoltage shall contain the core voltage, in volt units, of this processor. The core voltage of the processor
	// may change more frequently than the manager is able to monitor. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value 'Voltage'.
	CoreVoltage SensorVoltageExcerpt
	// CorrectableCoreErrorCount shall contain the number of correctable core errors, such as TLB or cache errors. When
	// this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableCoreErrorCount over all processors.
	CorrectableCoreErrorCount int
	// CorrectableOtherErrorCount shall contain the number of the correctable errors of all other components. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableOtherErrorCount over all processors.
	CorrectableOtherErrorCount int
	// Description provides a description of this resource.
	Description string
	// FrequencyRatio shall contain the frequency relative to the nominal processor frequency ratio of this processor.
	// When this resource is subordinate to the ProcessorSummary object, this property shall be the average
	// FrequencyRatio over all processors.
	FrequencyRatio float64
	// KernelPercent shall contain total percentage of time the processor has spent in kernel mode. When this resource
	// is subordinate to the ProcessorSummary object, this property shall be the average KernelPercent over all
	// processors.
	KernelPercent float64
	// LocalMemoryBandwidthBytes shall contain the local memory bandwidth usage of this processor in bytes. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// LocalMemoryBandwidthBytes over all processors.
	LocalMemoryBandwidthBytes int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingSpeedMHz shall contain the operating speed of the processor in MHz. The operating speed of the
	// processor may change more frequently than the manager is able to monitor.
	OperatingSpeedMHz int
	// PCIeErrors shall contain the PCIe errors associated with this processor.
	PCIeErrors string
	// PowerLimitThrottleDuration shall contain the total duration of throttling caused by a power limit of the
	// processor since reset.
	PowerLimitThrottleDuration string
	// RemoteMemoryBandwidthBytes shall contain the remote memory bandwidth usage of this processor in bytes. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// RemoteMemoryBandwidthBytes over all processors.
	RemoteMemoryBandwidthBytes int
	// ThermalLimitThrottleDuration shall contain the total duration of throttling caused by a thermal limit of the
	// processor since reset.
	ThermalLimitThrottleDuration string
	// ThrottlingCelsius shall contain the CPU margin to throttle based on an offset between the maximum temperature in
	// which the processor can operate, and the processor's current temperature. When this resource is subordinate to
	// the ProcessorSummary object, this property is not applicable.
	ThrottlingCelsius float64
	// UncorrectableCoreErrorCount shall contain the number of uncorrectable core errors, such as TLB or cache errors.
	// When this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableCoreErrorCount over all processors.
	UncorrectableCoreErrorCount int
	// UncorrectableOtherErrorCount shall contain the number of the uncorrectable errors of all other components. When
	// this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableOtherErrorCount over all processors.
	UncorrectableOtherErrorCount int
	// UserPercent shall contain total percentage of time the processor has spent in user mode. When this resource is
	// subordinate to the ProcessorSummary object, this property shall be the average UserPercent over all processors.
	UserPercent float64
}

// UnmarshalJSON unmarshals a ProcessorMetrics object from the raw JSON.
func (processormetrics *ProcessorMetrics) UnmarshalJSON(b []byte) error {
	type temp ProcessorMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processormetrics = ProcessorMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetProcessorMetrics will get a ProcessorMetrics instance from the service.
func GetProcessorMetrics(c common.Client, uri string) (*ProcessorMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var processormetrics ProcessorMetrics
	err = json.NewDecoder(resp.Body).Decode(&processormetrics)
	if err != nil {
		return nil, err
	}

	processormetrics.SetClient(c)
	return &processormetrics, nil
}

// ListReferencedProcessorMetricss gets the collection of ProcessorMetrics from
// a provided reference.
func ListReferencedProcessorMetricss(c common.Client, link string) ([]*ProcessorMetrics, error) {
	var result []*ProcessorMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, processormetricsLink := range links.ItemLinks {
		processormetrics, err := GetProcessorMetrics(c, processormetricsLink)
		if err != nil {
			collectionError.Failures[processormetricsLink] = err
		} else {
			result = append(result, processormetrics)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
