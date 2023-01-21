//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// BootTimeStatistics shall contain the boot time statistics of a manager.
type BootTimeStatistics struct {
	// FirmwareTimeSeconds shall contain the number of seconds the manager spent in the firmware stage.
	FirmwareTimeSeconds float64
	// InitrdTimeSeconds shall contain the number of seconds the manager spent in the initrd boot stage.
	InitrdTimeSeconds float64
	// KernelTimeSeconds shall contain the number of seconds the manager spent in the kernel stage.
	KernelTimeSeconds float64
	// LoaderTimeSeconds shall contain the number of seconds the manager spent in the loader stage.
	LoaderTimeSeconds float64
	// UserSpaceTimeSeconds shall contain the number of seconds the manager spent in the user space boot stage.
	UserSpaceTimeSeconds float64
}

// UnmarshalJSON unmarshals a BootTimeStatistics object from the raw JSON.
func (boottimestatistics *BootTimeStatistics) UnmarshalJSON(b []byte) error {
	type temp BootTimeStatistics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*boottimestatistics = BootTimeStatistics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// I2CBusStatistics shall contain statistics of an I2C bus.
type I2CBusStatistics struct {
	// BusErrorCount shall contain the number of bus errors on this I2C bus. Bus errors include, but are not limited
	// to, an SDA rising or falling edge while SCL is high or a stuck bus signal.
	BusErrorCount int
	// I2CBusName shall contain the name of the I2C bus.
	I2CBusName string
	// NACKCount shall contain the number of NACKs on this I2C bus.
	NACKCount int
	// TotalTransactionCount shall contain the total number of transactions on this I2C bus. The count shall include
	// the number of I2C transactions initiated by the manager and the number of I2C transactions where the manager is
	// the target device.
	TotalTransactionCount int
}

// UnmarshalJSON unmarshals a I2CBusStatistics object from the raw JSON.
func (i2cbusstatistics *I2CBusStatistics) UnmarshalJSON(b []byte) error {
	type temp I2CBusStatistics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i2cbusstatistics = I2CBusStatistics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ManagerDiagnosticData shall represent internal diagnostic data for a manager for a Redfish implementation.
// Clients should not make decisions for raising alerts, creating service events, or other actions based on
// information in this resource.
type ManagerDiagnosticData struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// BootTimeStatistics shall contain the boot time statistics of the manager.
	BootTimeStatistics string
	// Description provides a description of this resource.
	Description string
	// FreeStorageSpaceKiB shall contain the available storage space on this manager in kibibytes (KiB).
	FreeStorageSpaceKiB int
	// I2CBuses shall contain the statistics of the I2C buses. Services may subdivide a physical bus into multiple
	// entries in this property based on how the manager tracks bus segments, virtual buses from a controller, and
	// other segmentation capabilities.
	I2CBuses []I2CBusStatistics
	// MemoryECCStatistics shall contain the memory ECC statistics of the manager.
	MemoryECCStatistics string
	// MemoryStatistics shall contain the memory statistics of the manager.
	MemoryStatistics string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ProcessorStatistics shall contain the processor statistics of the manager.
	ProcessorStatistics string
	// TopProcesses shall contain the statistics of the top processes of this manager.
	TopProcesses []ProcessStatistics
}

// UnmarshalJSON unmarshals a ManagerDiagnosticData object from the raw JSON.
func (managerdiagnosticdata *ManagerDiagnosticData) UnmarshalJSON(b []byte) error {
	type temp ManagerDiagnosticData
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*managerdiagnosticdata = ManagerDiagnosticData(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetManagerDiagnosticData will get a ManagerDiagnosticData instance from the service.
func GetManagerDiagnosticData(c common.Client, uri string) (*ManagerDiagnosticData, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var managerdiagnosticdata ManagerDiagnosticData
	err = json.NewDecoder(resp.Body).Decode(&managerdiagnosticdata)
	if err != nil {
		return nil, err
	}

	managerdiagnosticdata.SetClient(c)
	return &managerdiagnosticdata, nil
}

// ListReferencedManagerDiagnosticDatas gets the collection of ManagerDiagnosticData from
// a provided reference.
func ListReferencedManagerDiagnosticDatas(c common.Client, link string) ([]*ManagerDiagnosticData, error) {
	var result []*ManagerDiagnosticData
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, managerdiagnosticdataLink := range links.ItemLinks {
		managerdiagnosticdata, err := GetManagerDiagnosticData(c, managerdiagnosticdataLink)
		if err != nil {
			collectionError.Failures[managerdiagnosticdataLink] = err
		} else {
			result = append(result, managerdiagnosticdata)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// MemoryECCStatistics shall contain the memory ECC statistics of a manager.
type MemoryECCStatistics struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors since reset.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors since reset.
	UncorrectableECCErrorCount int
}

// UnmarshalJSON unmarshals a MemoryECCStatistics object from the raw JSON.
func (memoryeccstatistics *MemoryECCStatistics) UnmarshalJSON(b []byte) error {
	type temp MemoryECCStatistics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memoryeccstatistics = MemoryECCStatistics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// MemoryStatistics shall contain the memory statistics of a manager.
type MemoryStatistics struct {
	// AvailableBytes shall contain the amount of memory available in bytes for starting new processes without
	// swapping. This includes free memory and reclaimable cache and buffers.
	AvailableBytes int
	// BuffersAndCacheBytes shall contain the amount of memory used in bytes by kernel buffers, page caches, and slabs.
	BuffersAndCacheBytes int
	// FreeBytes shall contain the amount of free memory in bytes.
	FreeBytes int
	// SharedBytes shall contain the amount of shared memory in bytes. This includes things such as memory consumed by
	// temporary filesystems.
	SharedBytes int
	// TotalBytes shall contain the total amount of memory in bytes.
	TotalBytes int
	// UsedBytes shall contain the amount of used memory in bytes. This value is calculated as TotalBytes minus
	// FreeBytes minus BuffersAndCacheBytes.
	UsedBytes int
}

// UnmarshalJSON unmarshals a MemoryStatistics object from the raw JSON.
func (memorystatistics *MemoryStatistics) UnmarshalJSON(b []byte) error {
	type temp MemoryStatistics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorystatistics = MemoryStatistics(t.temp)

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

// ProcessStatistics shall contain the statistics of a process running on a manager.
type ProcessStatistics struct {
	// CommandLine shall contain the command line with parameters of this process.
	CommandLine string
	// KernelTimeSeconds shall contain the number of seconds this process executed in kernel space.
	KernelTimeSeconds float64
	// ResidentSetSizeBytes shall contain the resident set size of this process in bytes, which is the amount of memory
	// allocated to the process and is in RAM.
	ResidentSetSizeBytes int
	// RestartAfterFailureCount shall contain the number of times this process has restarted unexpectedly, such as due
	// to unintentional failures, restarts, or shutdowns, with the same command line including arguments.
	RestartAfterFailureCount int
	// RestartCount shall contain the number of times this process has restarted with the same command line including
	// arguments.
	RestartCount int
	// UptimeSeconds shall contain the wall-clock time this process has been running in seconds.
	UptimeSeconds float64
	// UserTimeSeconds shall contain the number of seconds this process executed in user space.
	UserTimeSeconds float64
}

// UnmarshalJSON unmarshals a ProcessStatistics object from the raw JSON.
func (processstatistics *ProcessStatistics) UnmarshalJSON(b []byte) error {
	type temp ProcessStatistics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processstatistics = ProcessStatistics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ProcessorStatistics shall contain the processor statistics of a manager.
type ProcessorStatistics struct {
	// KernelPercent shall contain the percentage of CPU time spent in kernel mode.
	KernelPercent float64
	// UserPercent shall contain the percentage of CPU time spent in user mode.
	UserPercent float64
}

// UnmarshalJSON unmarshals a ProcessorStatistics object from the raw JSON.
func (processorstatistics *ProcessorStatistics) UnmarshalJSON(b []byte) error {
	type temp ProcessorStatistics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processorstatistics = ProcessorStatistics(t.temp)

	// Extract the links to other entities for later

	return nil
}
