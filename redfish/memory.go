//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// BaseModuleType is
type BaseModuleType string

const (
	// RDIMMBaseModuleType Registered DIMM.
	RDIMMBaseModuleType BaseModuleType = "RDIMM"
	// UDIMMBaseModuleType UDIMM.
	UDIMMBaseModuleType BaseModuleType = "UDIMM"
	// SODIMMBaseModuleType SO_DIMM.
	SODIMMBaseModuleType BaseModuleType = "SO_DIMM"
	// LRDIMMBaseModuleType Load Reduced.
	LRDIMMBaseModuleType BaseModuleType = "LRDIMM"
	// MiniRDIMMBaseModuleType Mini_RDIMM.
	MiniRDIMMBaseModuleType BaseModuleType = "Mini_RDIMM"
	// MiniUDIMMBaseModuleType Mini_UDIMM.
	MiniUDIMMBaseModuleType BaseModuleType = "Mini_UDIMM"
	// SORDIMM72bBaseModuleType SO_RDIMM_72b.
	SORDIMM72bBaseModuleType BaseModuleType = "SO_RDIMM_72b"
	// SOUDIMM72bBaseModuleType SO_UDIMM_72b.
	SOUDIMM72bBaseModuleType BaseModuleType = "SO_UDIMM_72b"
	// SODIMM16bBaseModuleType SO_DIMM_16b.
	SODIMM16bBaseModuleType BaseModuleType = "SO_DIMM_16b"
	// SODIMM32bBaseModuleType SO_DIMM_32b.
	SODIMM32bBaseModuleType BaseModuleType = "SO_DIMM_32b"
	// DieBaseModuleType A die within a package.
	DieBaseModuleType BaseModuleType = "Die"
)

// ErrorCorrection is
type ErrorCorrection string

const (
	// NoECCErrorCorrection No ECC available.
	NoECCErrorCorrection ErrorCorrection = "NoECC"
	// SingleBitECCErrorCorrection Single bit data errors can be corrected by ECC.
	SingleBitECCErrorCorrection ErrorCorrection = "SingleBitECC"
	// MultiBitECCErrorCorrection Multibit data errors can be corrected by ECC.
	MultiBitECCErrorCorrection ErrorCorrection = "MultiBitECC"
	// AddressParityErrorCorrection Address parity errors can be corrected.
	AddressParityErrorCorrection ErrorCorrection = "AddressParity"
)

// MemoryClassification is
type MemoryClassification string

const (
	// VolatileMemoryClassification Volatile memory.
	VolatileMemoryClassification MemoryClassification = "Volatile"
	// ByteAccessiblePersistentMemoryClassification Byte-accessible persistent memory.
	ByteAccessiblePersistentMemoryClassification MemoryClassification = "ByteAccessiblePersistent"
	// BlockMemoryClassification Block-accessible memory.
	BlockMemoryClassification MemoryClassification = "Block"
)

// MemoryDeviceType is
type MemoryDeviceType string

const (
	// DDRMemoryDeviceType DDR.
	DDRMemoryDeviceType MemoryDeviceType = "DDR"
	// DDR2MemoryDeviceType DDR2.
	DDR2MemoryDeviceType MemoryDeviceType = "DDR2"
	// DDR3MemoryDeviceType DDR3.
	DDR3MemoryDeviceType MemoryDeviceType = "DDR3"
	// DDR4MemoryDeviceType DDR4.
	DDR4MemoryDeviceType MemoryDeviceType = "DDR4"
	// DDR4SDRAMMemoryDeviceType DDR4 SDRAM.
	DDR4SDRAMMemoryDeviceType MemoryDeviceType = "DDR4_SDRAM"
	// DDR4ESDRAMMemoryDeviceType DDR4E SDRAM.
	DDR4ESDRAMMemoryDeviceType MemoryDeviceType = "DDR4E_SDRAM"
	// LPDDR4SDRAMMemoryDeviceType LPDDR4 SDRAM.
	LPDDR4SDRAMMemoryDeviceType MemoryDeviceType = "LPDDR4_SDRAM"
	// DDR3SDRAMMemoryDeviceType DDR3 SDRAM.
	DDR3SDRAMMemoryDeviceType MemoryDeviceType = "DDR3_SDRAM"
	// LPDDR3SDRAMMemoryDeviceType LPDDR3 SDRAM.
	LPDDR3SDRAMMemoryDeviceType MemoryDeviceType = "LPDDR3_SDRAM"
	// DDR2SDRAMMemoryDeviceType DDR2 SDRAM.
	DDR2SDRAMMemoryDeviceType MemoryDeviceType = "DDR2_SDRAM"
	// DDR2SDRAMFBDIMMMemoryDeviceType DDR2 SDRAM FB_DIMM.
	DDR2SDRAMFBDIMMMemoryDeviceType MemoryDeviceType = "DDR2_SDRAM_FB_DIMM"
	// DDR2SDRAMFBDIMMPROBEMemoryDeviceType DDR2 SDRAM FB_DIMM PROBE.
	DDR2SDRAMFBDIMMPROBEMemoryDeviceType MemoryDeviceType = "DDR2_SDRAM_FB_DIMM_PROBE"
	// DDRSGRAMMemoryDeviceType DDR SGRAM.
	DDRSGRAMMemoryDeviceType MemoryDeviceType = "DDR_SGRAM"
	// DDRSDRAMMemoryDeviceType DDR SDRAM.
	DDRSDRAMMemoryDeviceType MemoryDeviceType = "DDR_SDRAM"
	// ROMMemoryDeviceType ROM.
	ROMMemoryDeviceType MemoryDeviceType = "ROM"
	// SDRAMMemoryDeviceType SDRAM.
	SDRAMMemoryDeviceType MemoryDeviceType = "SDRAM"
	// EDOMemoryDeviceType EDO.
	EDOMemoryDeviceType MemoryDeviceType = "EDO"
	// FastPageModeMemoryDeviceType Fast Page Mode.
	FastPageModeMemoryDeviceType MemoryDeviceType = "FastPageMode"
	// PipelinedNibbleMemoryDeviceType Pipelined Nibble.
	PipelinedNibbleMemoryDeviceType MemoryDeviceType = "PipelinedNibble"
	// LogicalMemoryDeviceType Logical Non-volatile device.
	LogicalMemoryDeviceType MemoryDeviceType = "Logical"
	// HBMMemoryDeviceType High Bandwidth Memory.
	HBMMemoryDeviceType MemoryDeviceType = "HBM"
	// HBM2MemoryDeviceType The second generation of High Bandwidth Memory.
	HBM2MemoryDeviceType MemoryDeviceType = "HBM2"
	// HBM3MemoryDeviceType The third generation of High Bandwidth Memory.
	HBM3MemoryDeviceType MemoryDeviceType = "HBM3"
	// GDDRMemoryDeviceType Synchronous graphics random-access memory.
	GDDRMemoryDeviceType MemoryDeviceType = "GDDR"
	// GDDR2MemoryDeviceType Double data rate type two synchronous graphics random-access memory.
	GDDR2MemoryDeviceType MemoryDeviceType = "GDDR2"
	// GDDR3MemoryDeviceType Double data rate type three synchronous graphics random-access memory.
	GDDR3MemoryDeviceType MemoryDeviceType = "GDDR3"
	// GDDR4MemoryDeviceType Double data rate type four synchronous graphics random-access memory.
	GDDR4MemoryDeviceType MemoryDeviceType = "GDDR4"
	// GDDR5MemoryDeviceType Double data rate type five synchronous graphics random-access memory.
	GDDR5MemoryDeviceType MemoryDeviceType = "GDDR5"
	// GDDR5XMemoryDeviceType Double data rate type five X synchronous graphics random-access memory.
	GDDR5XMemoryDeviceType MemoryDeviceType = "GDDR5X"
	// GDDR6MemoryDeviceType Double data rate type six synchronous graphics random-access memory.
	GDDR6MemoryDeviceType MemoryDeviceType = "GDDR6"
	// DDR5MemoryDeviceType Double data rate type five synchronous dynamic random-access memory.
	DDR5MemoryDeviceType MemoryDeviceType = "DDR5"
	// OEMMemoryDeviceType OEM-defined.
	OEMMemoryDeviceType MemoryDeviceType = "OEM"
)

// MemoryMedia is
type MemoryMedia string

const (
	// DRAMMemoryMedia DRAM media.
	DRAMMemoryMedia MemoryMedia = "DRAM"
	// NANDMemoryMedia NAND media.
	NANDMemoryMedia MemoryMedia = "NAND"
	// Intel3DXPointMemoryMedia Intel 3D XPoint media.
	Intel3DXPointMemoryMedia MemoryMedia = "Intel3DXPoint"
	// ProprietaryMemoryMedia Proprietary media.
	ProprietaryMemoryMedia MemoryMedia = "Proprietary"
)

// MemoryType is
type MemoryType string

const (
	// DRAMMemoryType shall represent a volatile DRAM memory device.
	DRAMMemoryType MemoryType = "DRAM"
	// NVDIMMNMemoryType shall represent an NVDIMM_N memory device as defined by JEDEC.
	NVDIMMNMemoryType MemoryType = "NVDIMM_N"
	// NVDIMMFMemoryType shall represent an NVDIMM_F memory device as defined by JEDEC.
	NVDIMMFMemoryType MemoryType = "NVDIMM_F"
	// NVDIMMPMemoryType shall represent an NVDIMM_P memory device as defined by JEDEC.
	NVDIMMPMemoryType MemoryType = "NVDIMM_P"
	// IntelOptaneMemoryType shall represent an Intel Optane Persistent Memory Module.
	IntelOptaneMemoryType MemoryType = "IntelOptane"
)

// OperatingMemoryModes is
type OperatingMemoryModes string

const (
	// VolatileOperatingMemoryModes Volatile memory.
	VolatileOperatingMemoryModes OperatingMemoryModes = "Volatile"
	// PMEMOperatingMemoryModes Persistent memory, byte-accessible through system address space.
	PMEMOperatingMemoryModes OperatingMemoryModes = "PMEM"
	// BlockOperatingMemoryModes Block-accessible system memory.
	BlockOperatingMemoryModes OperatingMemoryModes = "Block"
)

// SecurityStates is
type SecurityStates string

const (
	// EnabledSecurityStates Secure mode is enabled and access to the data is allowed.
	EnabledSecurityStates SecurityStates = "Enabled"
	// DisabledSecurityStates Secure mode is disabled.
	DisabledSecurityStates SecurityStates = "Disabled"
	// UnlockedSecurityStates Secure mode is enabled and access to the data is unlocked.
	UnlockedSecurityStates SecurityStates = "Unlocked"
	// LockedSecurityStates Secure mode is enabled and access to the data is locked.
	LockedSecurityStates SecurityStates = "Locked"
	// FrozenSecurityStates Secure state is frozen and cannot be modified until reset.
	FrozenSecurityStates SecurityStates = "Frozen"
	// PassphraselimitSecurityStates Number of attempts to unlock the memory exceeded limit.
	PassphraselimitSecurityStates SecurityStates = "Passphraselimit"
)

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Batteries shall contain an array of links to resources of type Battery that represent the batteries that provide
	// power to this memory device during a power loss event, such as with battery-backed NVDIMMs. This property shall
	// not be present if the batteries power the containing chassis as a whole rather than the individual memory
	// device.
	Batteries []Battery
	// Batteries@odata.count
	BatteriesCount int `json:"Batteries@odata.count"`
	// Chassis shall contain a link to a resource of type Chassis that represents the physical container associated
	// with this memory device.
	Chassis string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Processors shall contain an array of links to resources of type Processor that are associated with this memory
	// device.
	Processors []Processor
	// Processors@odata.count
	ProcessorsCount int `json:"Processors@odata.count"`
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

// Memory shall represent a memory device in a Redfish implementation.
type Memory struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// AllocationAlignmentMiB shall contain the alignment boundary on which memory regions are allocated, measured in
	// MiB.
	AllocationAlignmentMiB int
	// AllocationIncrementMiB shall contain the allocation increment for regions, measured in MiB.
	AllocationIncrementMiB int
	// AllowedSpeedsMHz shall contain the speed supported by this memory device.
	AllowedSpeedsMHz []string
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// BaseModuleType shall contain the base module type of the memory device.
	BaseModuleType BaseModuleType
	// BusWidthBits shall contain the bus width, in bits.
	BusWidthBits int
	// CacheSizeMiB shall contain the total size of the cache portion memory in MiB.
	CacheSizeMiB int
	// CapacityMiB shall contain the memory capacity in MiB.
	CapacityMiB int
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	Certificates string
	// ConfigurationLocked shall indicate whether the configuration of this memory device is locked and cannot be
	// altered.
	ConfigurationLocked bool
	// DataWidthBits shall contain the data width in bits.
	DataWidthBits int
	// Description provides a description of this resource.
	Description string
	// Enabled shall indicate if this memory is enabled.
	Enabled string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this memory.
	EnvironmentMetrics string
	// ErrorCorrection shall contain the error correction scheme supported for this memory device.
	ErrorCorrection ErrorCorrection
	// FirmwareApiVersion shall contain the version of API supported by the firmware.
	FirmwareApiVersion string
	// FirmwareRevision shall contain the revision of firmware on the memory controller.
	FirmwareRevision string
	// IsRankSpareEnabled shall indicate whether rank spare is enabled for this memory device.
	IsRankSpareEnabled bool
	// IsSpareDeviceEnabled shall indicate whether the spare device is enabled.
	IsSpareDeviceEnabled bool
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated memory device.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// Log shall contain a link to a resource of type LogService.
	Log string
	// LogicalSizeMiB shall contain the total size of the logical memory in MiB.
	LogicalSizeMiB int
	// Manufacturer shall contain the manufacturer of the memory device.
	Manufacturer string
	// MaxTDPMilliWatts shall contain an array of maximum power budgets supported by the memory device in milliwatts.
	MaxTDPMilliWatts []string
	// MemoryDeviceType shall contain the Memory Device Type as defined by SMBIOS.
	MemoryDeviceType MemoryDeviceType
	// MemoryLocation shall contain properties that describe the memory connection information to sockets and memory
	// controllers.
	MemoryLocation string
	// MemoryMedia shall contain the media types of this memory device.
	MemoryMedia []MemoryMedia
	// MemorySubsystemControllerManufacturerID shall contain the two byte manufacturer ID of the memory subsystem
	// controller of this memory device as defined by JEDEC in JEP-106.
	MemorySubsystemControllerManufacturerID string
	// MemorySubsystemControllerProductID shall contain the two byte product ID of the memory subsystem controller of
	// this memory device as defined by the manufacturer.
	MemorySubsystemControllerProductID string
	// MemoryType shall contain the type of memory device that this resource represents.
	MemoryType MemoryType
	// Metrics The link to the metrics associated with this memory device.
	Metrics string
	// Model shall indicate the model information as provided by the manufacturer of this memory.
	Model string
	// ModuleManufacturerID shall contain the two byte manufacturer ID of this memory device as defined by JEDEC in
	// JEP-106.
	ModuleManufacturerID string
	// ModuleProductID shall contain the two byte product ID of this memory device as defined by the manufacturer.
	ModuleProductID string
	// NonVolatileSizeMiB shall contain the total size of the non-volatile portion memory in MiB.
	NonVolatileSizeMiB int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingMemoryModes shall contain the memory modes supported by the memory device.
	OperatingMemoryModes []OperatingMemoryModes
	// OperatingSpeedMhz shall contain the operating speed of the memory device in MHz or MT/s (mega-transfers per
	// second) as reported by the memory device. Memory devices that operate at their bus speed shall report the
	// operating speed in MHz (bus speed), while memory devices that transfer data faster than their bus speed, such as
	// DDR memory, shall report the operating speed in MT/s (mega-transfers/second). The reported value shall match the
	// conventionally reported values for the technology used by the memory device.
	OperatingSpeedMhz int
	// OperatingSpeedRangeMHz shall contain the operating speed control, in megahertz units, for this resource. The
	// value of the DataSourceUri property, if present, shall reference a resource of type Control with the ControlType
	// property containing the value of 'FrequencyMHz'.
	OperatingSpeedRangeMHz ControlRangeExcerpt
	// PartNumber shall indicate the part number as provided by the manufacturer of this memory device.
	PartNumber string
	// PersistentRegionNumberLimit shall contain the total number of persistent regions this memory device can support.
	PersistentRegionNumberLimit int
	// PersistentRegionSizeLimitMiB shall contain the total size of persistent regions in MiB.
	PersistentRegionSizeLimitMiB int
	// PersistentRegionSizeMaxMiB shall contain the maximum size of a single persistent regions in MiB.
	PersistentRegionSizeMaxMiB int
	// PowerManagementPolicy shall contain properties that describe the power management policy for this resource.
	PowerManagementPolicy string
	// RankCount shall contain the number of ranks available in the memory device. The ranks could be used for spare or
	// interleave.
	RankCount int
	// Regions shall contain the memory region information within the memory device.
	Regions []RegionSet
	// SecurityCapabilities shall contain properties that describe the security capabilities of the memory device.
	SecurityCapabilities string
	// SecurityState shall contain the current security state of this memory device.
	SecurityState SecurityStates
	// SerialNumber shall indicate the serial number as provided by the manufacturer of this memory device.
	SerialNumber string
	// SpareDeviceCount shall contain the number of unused spare devices available in the memory device. If memory
	// devices fails, the spare device could be used.
	SpareDeviceCount int
	// SparePartNumber shall contain the spare part number of the memory.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// VolatileRegionNumberLimit shall contain the total number of volatile regions this memory device can support.
	VolatileRegionNumberLimit int
	// VolatileRegionSizeLimitMiB shall contain the total size of volatile regions in MiB.
	VolatileRegionSizeLimitMiB int
	// VolatileRegionSizeMaxMiB shall contain the maximum size of a single volatile regions in MiB.
	VolatileRegionSizeMaxMiB int
	// VolatileSizeMiB shall contain the total size of the volatile portion memory in MiB.
	VolatileSizeMiB int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Memory object from the raw JSON.
func (memory *Memory) UnmarshalJSON(b []byte) error {
	type temp Memory
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memory = Memory(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	memory.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (memory *Memory) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Memory)
	original.UnmarshalJSON(memory.rawData)

	readWriteFields := []string{
		"Enabled",
		"LocationIndicatorActive",
		"OperatingSpeedRangeMHz",
		"SecurityState",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(memory).Elem()

	return memory.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetMemory will get a Memory instance from the service.
func GetMemory(c common.Client, uri string) (*Memory, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memory Memory
	err = json.NewDecoder(resp.Body).Decode(&memory)
	if err != nil {
		return nil, err
	}

	memory.SetClient(c)
	return &memory, nil
}

// ListReferencedMemorys gets the collection of Memory from
// a provided reference.
func ListReferencedMemorys(c common.Client, link string) ([]*Memory, error) {
	var result []*Memory
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, memoryLink := range links.ItemLinks {
		memory, err := GetMemory(c, memoryLink)
		if err != nil {
			collectionError.Failures[memoryLink] = err
		} else {
			result = append(result, memory)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// MemoryLocation shall contain properties that describe the memory connection information to sockets and memory
// controllers.
type MemoryLocation struct {
	// Channel shall contain the channel number to which the memory device is connected.
	Channel int
	// MemoryController shall contain the memory controller number to which the memory device is connected.
	MemoryController int
	// Slot shall contain the slot number to which the memory device is connected.
	Slot int
	// Socket shall contain the socket number to which the memory device is connected.
	Socket int
}

// UnmarshalJSON unmarshals a MemoryLocation object from the raw JSON.
func (memorylocation *MemoryLocation) UnmarshalJSON(b []byte) error {
	type temp MemoryLocation
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorylocation = MemoryLocation(t.temp)

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

// PowerManagementPolicy shall contain properties that describe the power management policy for this resource.
type PowerManagementPolicy struct {
	// AveragePowerBudgetMilliWatts shall contain the average power budget, in milliwatts.
	AveragePowerBudgetMilliWatts int
	// MaxTDPMilliWatts shall contain the maximum TDP in milliwatts.
	MaxTDPMilliWatts int
	// PeakPowerBudgetMilliWatts shall contain the peak power budget, in milliwatts.
	PeakPowerBudgetMilliWatts int
	// PolicyEnabled shall indicate whether the power management policy is enabled.
	PolicyEnabled bool
}

// UnmarshalJSON unmarshals a PowerManagementPolicy object from the raw JSON.
func (powermanagementpolicy *PowerManagementPolicy) UnmarshalJSON(b []byte) error {
	type temp PowerManagementPolicy
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powermanagementpolicy = PowerManagementPolicy(t.temp)

	// Extract the links to other entities for later

	return nil
}

// RegionSet shall describe the memory region information within a memory device.
type RegionSet struct {
	// MemoryClassification shall contain the classification of memory that the memory region occupies.
	MemoryClassification MemoryClassification
	// OffsetMiB shall contain the offset within the memory that corresponds to the start of this memory region in MiB.
	OffsetMiB int
	// PassphraseEnabled shall indicate whether the passphrase is enabled for this region.
	PassphraseEnabled bool
	// RegionId shall contain the unique region ID representing a specific region within the memory device.
	RegionId string
	// SizeMiB shall contain the size of this memory region in MiB.
	SizeMiB int
}

// UnmarshalJSON unmarshals a RegionSet object from the raw JSON.
func (regionset *RegionSet) UnmarshalJSON(b []byte) error {
	type temp RegionSet
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*regionset = RegionSet(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SecurityCapabilities shall contain properties that describe the security capabilities of a memory device.
type SecurityCapabilities struct {
	// ConfigurationLockCapable shall indicate whether this memory device supports the locking, or freezing, of the
	// configuration.
	ConfigurationLockCapable bool
	// DataLockCapable shall indicate whether this memory device supports the locking of data access.
	DataLockCapable bool
	// MaxPassphraseCount shall contain the maximum number of passphrases supported for this memory device.
	MaxPassphraseCount int
	// PassphraseCapable shall indicate whether the memory device is passphrase capable.
	PassphraseCapable bool
	// PassphraseLockLimit shall contain the maximum number of incorrect passphrase access attempts allowed before
	// access to data is locked. If 0, the number of attempts is infinite.
	PassphraseLockLimit int
}

// UnmarshalJSON unmarshals a SecurityCapabilities object from the raw JSON.
func (securitycapabilities *SecurityCapabilities) UnmarshalJSON(b []byte) error {
	type temp SecurityCapabilities
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*securitycapabilities = SecurityCapabilities(t.temp)

	// Extract the links to other entities for later

	return nil
}
