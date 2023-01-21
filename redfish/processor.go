//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// BaseSpeedPriorityState is
type BaseSpeedPriorityState string

const (
	// EnabledBaseSpeedPriorityState Base speed priority is enabled.
	EnabledBaseSpeedPriorityState BaseSpeedPriorityState = "Enabled"
	// DisabledBaseSpeedPriorityState Base speed priority is disabled.
	DisabledBaseSpeedPriorityState BaseSpeedPriorityState = "Disabled"
)

// FpgaType is
type FpgaType string

const (
	// IntegratedFpgaType The FPGA device integrated with other processor in the single chip.
	IntegratedFpgaType FpgaType = "Integrated"
	// DiscreteFpgaType The discrete FPGA device.
	DiscreteFpgaType FpgaType = "Discrete"
)

// InstructionSet is
type InstructionSet string

const (
	// x86InstructionSet x86 32-bit.
	x86InstructionSet InstructionSet = "x86"
	// x8664InstructionSet x86 64-bit.
	x8664InstructionSet InstructionSet = "x86-64"
	// IA64InstructionSet Intel IA-64.
	IA64InstructionSet InstructionSet = "IA-64"
	// ARMA32InstructionSet ARM 32-bit.
	ARMA32InstructionSet InstructionSet = "ARM-A32"
	// ARMA64InstructionSet ARM 64-bit.
	ARMA64InstructionSet InstructionSet = "ARM-A64"
	// MIPS32InstructionSet MIPS 32-bit.
	MIPS32InstructionSet InstructionSet = "MIPS32"
	// MIPS64InstructionSet MIPS 64-bit.
	MIPS64InstructionSet InstructionSet = "MIPS64"
	// PowerISAInstructionSet PowerISA-64 or PowerISA-32.
	PowerISAInstructionSet InstructionSet = "PowerISA"
	// OEMInstructionSet OEM-defined.
	OEMInstructionSet InstructionSet = "OEM"
)

// ProcessorArchitecture is
type ProcessorArchitecture string

const (
	// x86ProcessorArchitecture x86 or x86-64.
	x86ProcessorArchitecture ProcessorArchitecture = "x86"
	// IA64ProcessorArchitecture Intel Itanium.
	IA64ProcessorArchitecture ProcessorArchitecture = "IA-64"
	// ARMProcessorArchitecture ARM.
	ARMProcessorArchitecture ProcessorArchitecture = "ARM"
	// MIPSProcessorArchitecture MIPS.
	MIPSProcessorArchitecture ProcessorArchitecture = "MIPS"
	// PowerProcessorArchitecture Power.
	PowerProcessorArchitecture ProcessorArchitecture = "Power"
	// OEMProcessorArchitecture OEM-defined.
	OEMProcessorArchitecture ProcessorArchitecture = "OEM"
)

// ProcessorMemoryType is
type ProcessorMemoryType string

const (
	// L1CacheProcessorMemoryType L1 cache.
	L1CacheProcessorMemoryType ProcessorMemoryType = "L1Cache"
	// L2CacheProcessorMemoryType L2 cache.
	L2CacheProcessorMemoryType ProcessorMemoryType = "L2Cache"
	// L3CacheProcessorMemoryType L3 cache.
	L3CacheProcessorMemoryType ProcessorMemoryType = "L3Cache"
	// L4CacheProcessorMemoryType L4 cache.
	L4CacheProcessorMemoryType ProcessorMemoryType = "L4Cache"
	// L5CacheProcessorMemoryType L5 cache.
	L5CacheProcessorMemoryType ProcessorMemoryType = "L5Cache"
	// L6CacheProcessorMemoryType L6 cache.
	L6CacheProcessorMemoryType ProcessorMemoryType = "L6Cache"
	// L7CacheProcessorMemoryType L7 cache.
	L7CacheProcessorMemoryType ProcessorMemoryType = "L7Cache"
	// HBM1ProcessorMemoryType High Bandwidth Memory.
	HBM1ProcessorMemoryType ProcessorMemoryType = "HBM1"
	// HBM2ProcessorMemoryType The second generation of High Bandwidth Memory.
	HBM2ProcessorMemoryType ProcessorMemoryType = "HBM2"
	// HBM3ProcessorMemoryType The third generation of High Bandwidth Memory.
	HBM3ProcessorMemoryType ProcessorMemoryType = "HBM3"
	// SGRAMProcessorMemoryType Synchronous graphics RAM.
	SGRAMProcessorMemoryType ProcessorMemoryType = "SGRAM"
	// GDDRProcessorMemoryType Synchronous graphics random-access memory.
	GDDRProcessorMemoryType ProcessorMemoryType = "GDDR"
	// GDDR2ProcessorMemoryType Double data rate type two synchronous graphics random-access memory.
	GDDR2ProcessorMemoryType ProcessorMemoryType = "GDDR2"
	// GDDR3ProcessorMemoryType Double data rate type three synchronous graphics random-access memory.
	GDDR3ProcessorMemoryType ProcessorMemoryType = "GDDR3"
	// GDDR4ProcessorMemoryType Double data rate type four synchronous graphics random-access memory.
	GDDR4ProcessorMemoryType ProcessorMemoryType = "GDDR4"
	// GDDR5ProcessorMemoryType Double data rate type five synchronous graphics random-access memory.
	GDDR5ProcessorMemoryType ProcessorMemoryType = "GDDR5"
	// GDDR5XProcessorMemoryType Double data rate type five X synchronous graphics random-access memory.
	GDDR5XProcessorMemoryType ProcessorMemoryType = "GDDR5X"
	// GDDR6ProcessorMemoryType Double data rate type six synchronous graphics random-access memory.
	GDDR6ProcessorMemoryType ProcessorMemoryType = "GDDR6"
	// DDRProcessorMemoryType Double data rate synchronous dynamic random-access memory.
	DDRProcessorMemoryType ProcessorMemoryType = "DDR"
	// DDR2ProcessorMemoryType Double data rate type two synchronous dynamic random-access memory.
	DDR2ProcessorMemoryType ProcessorMemoryType = "DDR2"
	// DDR3ProcessorMemoryType Double data rate type three synchronous dynamic random-access memory.
	DDR3ProcessorMemoryType ProcessorMemoryType = "DDR3"
	// DDR4ProcessorMemoryType Double data rate type four synchronous dynamic random-access memory.
	DDR4ProcessorMemoryType ProcessorMemoryType = "DDR4"
	// DDR5ProcessorMemoryType Double data rate type five synchronous dynamic random-access memory.
	DDR5ProcessorMemoryType ProcessorMemoryType = "DDR5"
	// SDRAMProcessorMemoryType Synchronous dynamic random-access memory.
	SDRAMProcessorMemoryType ProcessorMemoryType = "SDRAM"
	// SRAMProcessorMemoryType Static random-access memory.
	SRAMProcessorMemoryType ProcessorMemoryType = "SRAM"
	// FlashProcessorMemoryType Flash memory.
	FlashProcessorMemoryType ProcessorMemoryType = "Flash"
	// OEMProcessorMemoryType OEM-defined.
	OEMProcessorMemoryType ProcessorMemoryType = "OEM"
)

// ProcessorType is
type ProcessorType string

const (
	// CPUProcessorType A CPU.
	CPUProcessorType ProcessorType = "CPU"
	// GPUProcessorType A GPU.
	GPUProcessorType ProcessorType = "GPU"
	// FPGAProcessorType An FPGA.
	FPGAProcessorType ProcessorType = "FPGA"
	// DSPProcessorType A DSP.
	DSPProcessorType ProcessorType = "DSP"
	// AcceleratorProcessorType An accelerator.
	AcceleratorProcessorType ProcessorType = "Accelerator"
	// CoreProcessorType A core in a processor.
	CoreProcessorType ProcessorType = "Core"
	// ThreadProcessorType A thread in a processor.
	ThreadProcessorType ProcessorType = "Thread"
	// OEMProcessorType An OEM-defined processing unit.
	OEMProcessorType ProcessorType = "OEM"
)

// SystemInterfaceType is
type SystemInterfaceType string

const (
	// QPISystemInterfaceType The Intel QuickPath Interconnect.
	QPISystemInterfaceType SystemInterfaceType = "QPI"
	// UPISystemInterfaceType The Intel UltraPath Interconnect.
	UPISystemInterfaceType SystemInterfaceType = "UPI"
	// PCIeSystemInterfaceType A PCI Express interface.
	PCIeSystemInterfaceType SystemInterfaceType = "PCIe"
	// EthernetSystemInterfaceType An Ethernet interface.
	EthernetSystemInterfaceType SystemInterfaceType = "Ethernet"
	// AMBASystemInterfaceType The Arm Advanced Microcontroller Bus Architecture interface.
	AMBASystemInterfaceType SystemInterfaceType = "AMBA"
	// CCIXSystemInterfaceType The Cache Coherent Interconnect for Accelerators interface.
	CCIXSystemInterfaceType SystemInterfaceType = "CCIX"
	// CXLSystemInterfaceType The Compute Express Link interface.
	CXLSystemInterfaceType SystemInterfaceType = "CXL"
	// OEMSystemInterfaceType An OEM-defined interface.
	OEMSystemInterfaceType SystemInterfaceType = "OEM"
)

// ThrottleCause is
type ThrottleCause string

const (
	// PowerLimitThrottleCause The cause of the processor being throttled is a power limit.
	PowerLimitThrottleCause ThrottleCause = "PowerLimit"
	// ThermalLimitThrottleCause The cause of the processor being throttled is a thermal limit.
	ThermalLimitThrottleCause ThrottleCause = "ThermalLimit"
	// ClockLimitThrottleCause The cause of the processor being throttled is a clock limit.
	ClockLimitThrottleCause ThrottleCause = "ClockLimit"
	// UnknownThrottleCause The cause of the processor being throttled is not known.
	UnknownThrottleCause ThrottleCause = "Unknown"
	// OEMThrottleCause The cause of the processor being throttled is OEM-specific.
	OEMThrottleCause ThrottleCause = "OEM"
)

// TurboState is
type TurboState string

const (
	// EnabledTurboState Turbo is enabled.
	EnabledTurboState TurboState = "Enabled"
	// DisabledTurboState Turbo is disabled.
	DisabledTurboState TurboState = "Disabled"
)

// EthernetInterface shall contain the definition for an Ethernet interface for a Redfish implementation.
type EthernetInterface struct {
	// MaxLanes shall contain the maximum number of lanes supported by this interface.
	MaxLanes int
	// MaxSpeedMbps shall contain the maximum speed supported by this interface.
	MaxSpeedMbps int
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a EthernetInterface object from the raw JSON.
func (ethernetinterface *EthernetInterface) UnmarshalJSON(b []byte) error {
	type temp EthernetInterface
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ethernetinterface = EthernetInterface(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FPGA shall contain the properties of the FPGA device represented by a processor.
type FPGA struct {
	// ExternalInterfaces shall contain an array of objects that describe the external connectivity of the FPGA.
	ExternalInterfaces []ProcessorInterface
	// FirmwareId shall contain a string describing the FPGA firmware identifier.
	FirmwareId string
	// FirmwareManufacturer shall contain a string describing the FPGA firmware manufacturer.
	FirmwareManufacturer string
	// FpgaType shall contain a type of the FPGA device.
	FpgaType string
	// Model shall contain a model of the FPGA device.
	Model string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeVirtualFunctions shall contain an integer that describes the number of PCIe Virtual Functions configured
	// within the FPGA.
	PCIeVirtualFunctions string
	// ProgrammableFromHost shall indicate whether the FPGA firmware can be reprogrammed from the host by using system
	// software. If 'false', system software shall not be able to program the FPGA firmware from the system interface.
	// In either state, a management controller may be able to program the FPGA firmware by using the sideband
	// interface.
	ProgrammableFromHost bool
	// ReconfigurationSlots shall contain an array of the structures that describe the FPGA reconfiguration slots that
	// the acceleration functions can program.
	ReconfigurationSlots []FpgaReconfigurationSlot
}

// UnmarshalJSON unmarshals a FPGA object from the raw JSON.
func (fpga *FPGA) UnmarshalJSON(b []byte) error {
	type temp FPGA
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fpga = FPGA(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FpgaReconfigurationSlot shall contain information about the FPGA reconfiguration slot.
type FpgaReconfigurationSlot struct {
	// AccelerationFunction shall contain a link to a resource of type AccelerationFunction that represents the code
	// programmed into this reconfiguration slot.
	AccelerationFunction string
	// ProgrammableFromHost shall indicate whether the reconfiguration slot can be reprogrammed from the host by using
	// system software. If 'false', system software shall not be able to program the reconfiguration slot from the
	// system interface. In either state, a management controller may be able to program the reconfiguration slot by
	// using the sideband interface.
	ProgrammableFromHost bool
	// SlotId shall contain the FPGA reconfiguration slot identifier.
	SlotId string
	// UUID shall contain a universal unique identifier number for the reconfiguration slot.
	UUID string
}

// UnmarshalJSON unmarshals a FpgaReconfigurationSlot object from the raw JSON.
func (fpgareconfigurationslot *FpgaReconfigurationSlot) UnmarshalJSON(b []byte) error {
	type temp FpgaReconfigurationSlot
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fpgareconfigurationslot = FpgaReconfigurationSlot(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// Chassis shall contain a link to a resource of type Chassis that represents the physical container associated
	// with this processor.
	Chassis string
	// ConnectedProcessors shall contain an array of links to resources of type Processor that are directly connected
	// to this processor.
	ConnectedProcessors []Processor
	// ConnectedProcessors@odata.count
	ConnectedProcessorsCount int `json:"ConnectedProcessors@odata.count"`
	// Endpoints shall contain an array of links to resources of type Endpoint that represent endpoints associated with
	// this processor.
	Endpoints []Endpoint
	// Endpoints@odata.count
	EndpointsCount int `json:"Endpoints@odata.count"`
	// GraphicsController shall contain a link to resource of type GraphicsController that is associated with this
	// processor.
	GraphicsController GraphicsController
	// Memory shall contain an array of links to resources of type Memory that are associated with this processor.
	Memory []Memory
	// Memory@odata.count
	MemoryCount int `json:"Memory@odata.count"`
	// NetworkDeviceFunctions shall contain an array of links to resources of type NetworkDeviceFunction that represent
	// the network device functions to which this processor performs offload computation, such as with a SmartNIC.
	NetworkDeviceFunctions []NetworkDeviceFunction
	// NetworkDeviceFunctions@odata.count
	NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevice shall contain a link to a resource of type PCIeDevice that represents the PCIe device associated with
	// this processor.
	PCIeDevice string
	// PCIeFunctions shall contain an array of links to resources of type PCIeFunction that represent the PCIe
	// functions associated with this processor.
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

// MemorySummary shall contain properties that describe the summary of all memory that are associated with a
// processor.
type MemorySummary struct {
	// ECCModeEnabled shall indicate if memory ECC mode is enabled for this processor. This value shall not affect
	// system memory ECC mode.
	ECCModeEnabled bool
	// Metrics shall contain a link to a resource of type MemoryMetrics that contains the metrics associated with all
	// memory of this processor.
	Metrics string
	// TotalCacheSizeMiB shall contain the total size of cache memory of this processor.
	TotalCacheSizeMiB int
	// TotalMemorySizeMiB shall contain the total size of non-cache, volatile memory attached to this processor. This
	// value indicates the size of memory directly attached or with strong affinity to this processor, not the total
	// memory accessible by the processor. This property shall not be present for implementations where all processors
	// have equal memory performance or access characteristics, such as hop count, for all system memory.
	TotalMemorySizeMiB int
}

// UnmarshalJSON unmarshals a MemorySummary object from the raw JSON.
func (memorysummary *MemorySummary) UnmarshalJSON(b []byte) error {
	type temp MemorySummary
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorysummary = MemorySummary(t.temp)

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

// Processor shall represent a single processor that a system contains. A processor includes both performance
// characteristics, clock speed, architecture, core count, and so on, and compatibility, such as the CPU ID
// instruction results.
type Processor struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccelerationFunctions shall contain a link to a resource collection of type AccelerationFunctionCollection.
	AccelerationFunctions string
	// Actions shall contain the available actions for this resource.
	Actions string
	// AdditionalFirmwareVersions shall contain the additional firmware versions of the processor.
	AdditionalFirmwareVersions string
	// AppliedOperatingConfig shall contain a link to a resource of type OperatingConfig that specifies the
	// configuration is applied to this processor.
	AppliedOperatingConfig string
	// Assembly shall contain a link to a resource of type Assembly.
	Assembly string
	// BaseSpeedMHz shall contain the base (nominal) clock speed of the processor in MHz.
	BaseSpeedMHz int
	// BaseSpeedPriorityState shall contain the state of the base frequency settings of the operating configuration
	// applied to this processor.
	BaseSpeedPriorityState BaseSpeedPriorityState
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	Certificates string
	// Description provides a description of this resource.
	Description string
	// Enabled shall indicate if this processor is enabled.
	Enabled string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this processor.
	EnvironmentMetrics string
	// FPGA shall contain an object containing properties for processors of type FPGA.
	FPGA string
	// Family shall contain a string that identifies the processor family, as specified by the combination of the
	// EffectiveFamily and EffectiveModel properties.
	Family string
	// FirmwareVersion shall contain a string describing the firmware version of the processor as provided by the
	// manufacturer.
	FirmwareVersion string
	// HighSpeedCoreIDs shall contain an array of core identifiers corresponding to the cores that have been configured
	// with the higher clock speed from the operating configuration applied to this processor.
	HighSpeedCoreIDs []string
	// InstructionSet shall contain the string that identifies the instruction set of the processor contained in this
	// socket.
	InstructionSet InstructionSet
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Location shall contain location information of the associated processor.
	Location string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// Manufacturer shall contain a string that identifies the manufacturer of the processor.
	Manufacturer string
	// MaxSpeedMHz shall indicate the maximum rated clock speed of the processor in MHz.
	MaxSpeedMHz int
	// MaxTDPWatts shall contain the maximum Thermal Design Power (TDP) in watts.
	MaxTDPWatts int
	// MemorySummary shall contain properties that describe the summary of all memory that are associated with this
	// processor.
	MemorySummary string
	// Metrics shall contain a link to a resource of type ProcessorMetrics that contains the metrics associated with
	// this processor.
	Metrics string
	// MinSpeedMHz shall indicate the minimum rated clock speed of the processor in MHz.
	MinSpeedMHz int
	// Model shall indicate the model information as provided by the manufacturer of this processor.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingConfigs shall contain a link to a resource collection of type OperatingConfigCollection.
	OperatingConfigs string
	// OperatingSpeedMHz shall contain the operating speed of the processor in MHz. The operating speed of the
	// processor may change more frequently than the manager is able to monitor.
	OperatingSpeedMHz int
	// OperatingSpeedRangeMHz shall contain the operating speed control, measured in megahertz units, for this
	// resource. The value of the DataSourceUri property, if present, shall reference a resource of type Control with
	// the ControlType property containing the value of 'FrequencyMHz'.
	OperatingSpeedRangeMHz ControlRangeExcerpt
	// PartNumber shall contain a part number assigned by the organization that is responsible for producing or
	// manufacturing the processor.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection. It shall contain the interconnect
	// ports of this processor. It shall not contain ports of for GraphicsController resources, USBController
	// resources, or other adapter-related type of resources.
	Ports string
	// ProcessorArchitecture shall contain the string that identifies the architecture of the processor contained in
	// this socket.
	ProcessorArchitecture ProcessorArchitecture
	// ProcessorId shall contain identification information for this processor.
	ProcessorId string
	// ProcessorIndex shall contain the zero-based index of the processor, indexed within the next unit of containment.
	// The value of this property shall match the ordering in the operating system topology interfaces, with offset
	// adjustments, if needed.
	ProcessorIndex int
	// ProcessorMemory shall contain the memory directly attached or integrated within this processor.
	ProcessorMemory []ProcessorMemory
	// ProcessorType shall contain the string that identifies the type of processor contained in this socket.
	ProcessorType ProcessorType
	// Replaceable shall indicate whether this component can be independently replaced as allowed by the vendor's
	// replacement policy. A value of 'false' indicates the component needs to be replaced by policy, as part of
	// another component. If the 'LocationType' property of this component contains 'Embedded', this property shall
	// contain 'false'.
	Replaceable bool
	// SerialNumber shall contain a manufacturer-allocated number that identifies the processor.
	SerialNumber string
	// Socket shall contain the string that identifies the physical location or socket of the processor.
	Socket string
	// SparePartNumber shall contain the spare part number of the processor.
	SparePartNumber string
	// SpeedLimitMHz shall contain the clock limit of the processor in MHz. This value shall be within the range of
	// MinSpeedMHz and MaxSpeedMHz as provided by the manufacturer of this processor.
	SpeedLimitMHz int
	// SpeedLocked shall indicate whether the clock speed of the processor is fixed, where a value 'true' shall
	// indicate that the clock speed is fixed at the value specified in the SpeedLimitMHz property.
	SpeedLocked bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SubProcessors shall contain a link to a resource collection of type ProcessorCollection.
	SubProcessors string
	// SystemInterface shall contain an object that describes the connectivity between the host system and the
	// processor.
	SystemInterface string
	// TDPWatts shall contain the nominal Thermal Design Power (TDP) in watts.
	TDPWatts int
	// ThrottleCauses shall contain the causes of the processor being throttled. If Throttled contains 'false', this
	// property shall contain an empty array.
	ThrottleCauses []ThrottleCause
	// Throttled shall indicate whether the processor is throttled.
	Throttled bool
	// TotalCores shall indicate the total count of independent processor cores contained within this processor.
	TotalCores int
	// TotalEnabledCores shall indicate the total count of enabled independent processor cores contained within this
	// processor.
	TotalEnabledCores int
	// TotalThreads shall indicate the total count of independent execution threads that this processor supports.
	TotalThreads int
	// TurboState shall contain the state of turbo for this processor.
	TurboState TurboState
	// UUID shall contain a universal unique identifier number for the processor. RFC4122 describes methods to use to
	// create the value. The value should be considered to be opaque. Client software should only treat the overall
	// value as a universally unique identifier and should not interpret any sub-fields within the UUID.
	UUID string
	// Version shall contain the hardware version of the processor as determined by the vendor or supplier.
	Version string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Processor object from the raw JSON.
func (processor *Processor) UnmarshalJSON(b []byte) error {
	type temp Processor
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processor = Processor(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	processor.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (processor *Processor) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Processor)
	original.UnmarshalJSON(processor.rawData)

	readWriteFields := []string{
		"AppliedOperatingConfig",
		"Enabled",
		"LocationIndicatorActive",
		"OperatingSpeedRangeMHz",
		"SpeedLimitMHz",
		"SpeedLocked",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(processor).Elem()

	return processor.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetProcessor will get a Processor instance from the service.
func GetProcessor(c common.Client, uri string) (*Processor, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var processor Processor
	err = json.NewDecoder(resp.Body).Decode(&processor)
	if err != nil {
		return nil, err
	}

	processor.SetClient(c)
	return &processor, nil
}

// ListReferencedProcessors gets the collection of Processor from
// a provided reference.
func ListReferencedProcessors(c common.Client, link string) ([]*Processor, error) {
	var result []*Processor
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, processorLink := range links.ItemLinks {
		processor, err := GetProcessor(c, processorLink)
		if err != nil {
			collectionError.Failures[processorLink] = err
		} else {
			result = append(result, processor)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// ProcessorId shall contain identification information for a processor.
type ProcessorId struct {
	// EffectiveFamily shall indicate the effective Family information as provided by the manufacturer of this
	// processor.
	EffectiveFamily string
	// EffectiveModel shall indicate the effective Model information as provided by the manufacturer of this processor.
	EffectiveModel string
	// IdentificationRegisters shall contain the raw manufacturer-provided processor-specific identification registers
	// of this processor's features.
	IdentificationRegisters string
	// MicrocodeInfo shall indicate the microcode information as provided by the manufacturer of this processor.
	MicrocodeInfo string
	// ProtectedIdentificationNumber shall contain the Protected Processor Identification Number (PPIN) for this
	// processor.
	ProtectedIdentificationNumber string
	// Step shall indicate the Step or revision string information as provided by the manufacturer of this processor.
	Step string
	// VendorId shall indicate the vendor Identification string information as provided by the manufacturer of this
	// processor.
	VendorId string
}

// UnmarshalJSON unmarshals a ProcessorId object from the raw JSON.
func (processorid *ProcessorId) UnmarshalJSON(b []byte) error {
	type temp ProcessorId
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processorid = ProcessorId(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ProcessorInterface shall contain information about the system interface, or external connection, to the
// processor.
type ProcessorInterface struct {
	// Ethernet shall contain an object the describes the Ethernet-related information for this interface.
	Ethernet string
	// InterfaceType shall contain an enumerated value that describes the type of interface between the system, or
	// external connection, and the processor.
	InterfaceType SystemInterfaceType
	// PCIe shall contain an object the describes the PCIe-related information for this interface.
	PCIe string
}

// UnmarshalJSON unmarshals a ProcessorInterface object from the raw JSON.
func (processorinterface *ProcessorInterface) UnmarshalJSON(b []byte) error {
	type temp ProcessorInterface
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processorinterface = ProcessorInterface(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ProcessorMemory shall contain information about memory directly attached or integrated within a processor.
type ProcessorMemory struct {
	// CapacityMiB shall contain the memory capacity in MiB.
	CapacityMiB int
	// IntegratedMemory shall indicate whether this memory is integrated within the processor. Otherwise, it is
	// discrete memory attached to the processor.
	IntegratedMemory bool
	// MemoryType shall contain a type of the processor memory type.
	MemoryType ProcessorMemoryType
	// SpeedMHz shall contain the operating speed of the memory in MHz.
	SpeedMHz int
}

// UnmarshalJSON unmarshals a ProcessorMemory object from the raw JSON.
func (processormemory *ProcessorMemory) UnmarshalJSON(b []byte) error {
	type temp ProcessorMemory
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processormemory = ProcessorMemory(t.temp)

	// Extract the links to other entities for later

	return nil
}
