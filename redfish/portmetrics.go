//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// FibreChannel shall describe Fibre Channel-specific metrics for network ports.
type FibreChannel struct {
	// CorrectableFECErrors shall contain the total number of times this port has received traffic with correctable
	// forward error correction (FEC) errors.
	CorrectableFECErrors int
	// InvalidCRCs shall contain the total number of invalid cyclic redundancy checks (CRCs) observed on this port.
	InvalidCRCs int
	// InvalidTXWords shall contain the total number of times this port has received invalid transmission words.
	InvalidTXWords int
	// LinkFailures shall contain the total number of link failures observed on this port.
	LinkFailures int
	// LossesOfSignal shall contain the total number of times this port has lost signal.
	LossesOfSignal int
	// LossesOfSync shall contain the total number of times this port has lost sync.
	LossesOfSync int
	// RXBBCreditZero shall contain the number of times the receive buffer-to-buffer credit count transitioned to zero
	// since last counter reset.
	RXBBCreditZero int
	// RXExchanges shall contain the total number of Fibre Channel exchanges received.
	RXExchanges int
	// RXSequences shall contain the total number of Fibre Channel sequences received.
	RXSequences int
	// TXBBCreditZero shall contain the number of times the transmit buffer-to-buffer credit count transitioned to zero
	// since last counter reset.
	TXBBCreditZero int
	// TXBBCreditZeroDurationMilliseconds shall contain the total amount of time in milliseconds the port has been
	// blocked from transmitting due to lack of buffer credits since the last counter reset.
	TXBBCreditZeroDurationMilliseconds int
	// TXBBCredits shall contain the number of transmit buffer-to-buffer credits the port is configured to use.
	TXBBCredits int
	// TXExchanges shall contain the total number of Fibre Channel exchanges transmitted.
	TXExchanges int
	// TXSequences shall contain the total number of Fibre Channel sequences transmitted.
	TXSequences int
	// UncorrectableFECErrors shall contain the total number of times this port has received traffic with uncorrectable
	// forward error correction (FEC) errors.
	UncorrectableFECErrors int
}

// UnmarshalJSON unmarshals a FibreChannel object from the raw JSON.
func (fibrechannel *FibreChannel) UnmarshalJSON(b []byte) error {
	type temp FibreChannel
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fibrechannel = FibreChannel(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GenZ shall describe the Gen-Z related port metrics.
type GenZ struct {
	// AccessKeyViolations shall contain the total number of Access Key Violations detected for packets received or
	// transmitted on this interface.
	AccessKeyViolations int
	// EndToEndCRCErrors shall contain total number of ECRC transient errors detected in received link-local and end-
	// to-end packets.
	EndToEndCRCErrors int
	// LLRRecovery shall contain the total number of times Link-level Reliability (LLR) recovery has been initiated by
	// this interface. This is not to be confused with the number of packets retransmitted due to initiating LLR
	// recovery.
	LLRRecovery int
	// LinkNTE shall contain the total number of link-local non-transient errors detected on this interface.
	LinkNTE int
	// MarkedECN shall contain the number of packets that the component set the Congestion ECN bit prior to
	// transmission through this interface.
	MarkedECN int
	// NonCRCTransientErrors shall contain the total number of transient errors detected that are unrelated to CRC
	// validation, which covers link-local and end-to-end packets, such as malformed Link Idle packets or PLA signal
	// errors.
	NonCRCTransientErrors int
	// PacketCRCErrors shall contain the total number of PCRC transient errors detected in received link-local and end-
	// to-end packets.
	PacketCRCErrors int
	// PacketDeadlineDiscards shall contain the number of packets discarded by this interface due to the Congestion
	// Deadline sub-field reaching zero prior to packet transmission.
	PacketDeadlineDiscards int
	// RXStompedECRC shall contain the total number of packets that this interface received with a stomped ECRC field.
	RXStompedECRC int
	// ReceivedECN shall contain the number of packets received on this interface with the Congestion ECN bit set.
	ReceivedECN int
	// TXStompedECRC shall contain the total number of packets that this interfaced stomped the ECRC field.
	TXStompedECRC int
}

// UnmarshalJSON unmarshals a GenZ object from the raw JSON.
func (genz *GenZ) UnmarshalJSON(b []byte) error {
	type temp GenZ
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*genz = GenZ(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Networking shall describe the metrics for network ports, including Ethernet, Fibre Channel, and InfiniBand, that
// are not specific to one of these protocols.
type Networking struct {
	// RDMAProtectionErrors shall contain the total number of RDMA protection errors.
	RDMAProtectionErrors int
	// RDMAProtocolErrors shall contain the total number of RDMA protocol errors.
	RDMAProtocolErrors int
	// RDMARXBytes shall contain the total number of RDMA bytes received on a port since reset.
	RDMARXBytes int
	// RDMARXRequests shall contain the total number of RDMA requests received on a port since reset.
	RDMARXRequests int
	// RDMATXBytes shall contain the total number of RDMA bytes transmitted on a port since reset.
	RDMATXBytes int
	// RDMATXReadRequests shall contain the total number of RDMA read requests transmitted on a port since reset.
	RDMATXReadRequests int
	// RDMATXRequests shall contain the total number of RDMA requests transmitted on a port since reset.
	RDMATXRequests int
	// RDMATXSendRequests shall contain the total number of RDMA send requests transmitted on a port since reset.
	RDMATXSendRequests int
	// RDMATXWriteRequests shall contain the total number of RDMA write requests transmitted on a port since reset.
	RDMATXWriteRequests int
	// RXBroadcastFrames shall contain the total number of valid broadcast frames received on a port since reset,
	// including host and remote management passthrough traffic.
	RXBroadcastFrames int
	// RXDiscards shall contain the total number of frames discarded in a port's receive path since reset.
	RXDiscards int
	// RXFCSErrors shall contain the total number of frames received with frame check sequence (FCS) errors on a port
	// since reset.
	RXFCSErrors int
	// RXFalseCarrierErrors shall contain the total number of false carrier errors received from phy on a port since
	// reset.
	RXFalseCarrierErrors int
	// RXFrameAlignmentErrors shall contain the total number of frames received with alignment errors on a port since
	// reset.
	RXFrameAlignmentErrors int
	// RXFrames shall contain the total number of frames received on a port since reset.
	RXFrames int
	// RXMulticastFrames shall contain the total number of valid multicast frames received on a port since reset,
	// including host and remote management passthrough traffic.
	RXMulticastFrames int
	// RXOversizeFrames shall contain the total number of frames that exceed the maximum frame size.
	RXOversizeFrames int
	// RXPFCFrames shall contain the total number of priority flow control (PFC) frames received on a port since reset.
	RXPFCFrames int
	// RXPauseXOFFFrames shall contain the total number of flow control frames from the network to pause transmission.
	RXPauseXOFFFrames int
	// RXPauseXONFrames shall contain the total number of flow control frames from the network to resume transmission.
	RXPauseXONFrames int
	// RXUndersizeFrames shall contain the total number of frames that are smaller than the minimum frame size of 64
	// bytes.
	RXUndersizeFrames int
	// RXUnicastFrames shall contain the total number of valid unicast frames received on a port since reset.
	RXUnicastFrames int
	// TXBroadcastFrames shall contain the total number of good broadcast frames transmitted on a port since reset,
	// including host and remote management passthrough traffic.
	TXBroadcastFrames int
	// TXDiscards shall contain the total number of frames discarded in a port's transmit path since reset.
	TXDiscards int
	// TXExcessiveCollisions shall contain the number of times a single transmitted frame encountered more than 15
	// collisions.
	TXExcessiveCollisions int
	// TXFrames shall contain the total number of frames transmitted on a port since reset.
	TXFrames int
	// TXLateCollisions shall contain the total number of collisions that occurred after one slot time as defined by
	// IEEE 802.3.
	TXLateCollisions int
	// TXMulticastFrames shall contain the total number of good multicast frames transmitted on a port since reset,
	// including host and remote management passthrough traffic.
	TXMulticastFrames int
	// TXMultipleCollisions shall contain the times that a transmitted frame encountered 2-15 collisions.
	TXMultipleCollisions int
	// TXPFCFrames shall contain the total number of priority flow control (PFC) frames sent on a port since reset.
	TXPFCFrames int
	// TXPauseXOFFFrames shall contain the total number of XOFF frames transmitted to the network.
	TXPauseXOFFFrames int
	// TXPauseXONFrames shall contain the total number of XON frames transmitted to the network.
	TXPauseXONFrames int
	// TXSingleCollisions shall contain the times that a successfully transmitted frame encountered a single collision.
	TXSingleCollisions int
	// TXUnicastFrames shall contain the total number of good unicast frames transmitted on a port since reset,
	// including host and remote management passthrough traffic.
	TXUnicastFrames int
}

// UnmarshalJSON unmarshals a Networking object from the raw JSON.
func (networking *Networking) UnmarshalJSON(b []byte) error {
	type temp Networking
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*networking = Networking(t.temp)

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

// PortMetrics shall represent the port metrics for a switch device or component port summary in a Redfish
// implementation.
type PortMetrics struct {
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
	// FibreChannel shall contain Fibre Channel-specific port metrics for network ports.
	FibreChannel string
	// GenZ shall contain the port metrics specific to Gen-Z ports.
	GenZ string
	// Networking shall contain port metrics for network ports, including Ethernet, Fibre Channel, and InfiniBand, that
	// are not specific to one of these protocols.
	Networking string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeErrors shall contain the PCIe errors associated with this port.
	PCIeErrors string
	// RXBytes shall contain the total number of bytes received on a port since reset, including host and remote
	// management passthrough traffic, and inclusive of all protocol overhead.
	RXBytes int
	// RXErrors shall contain the total number of received errors on a port since reset.
	RXErrors int
	// SAS shall contain an array of physical related metrics for Serial Attached SCSI (SAS). Each member in the array
	// shall represent a single phy.
	SAS []SAS
	// TXBytes shall contain the total number of bytes transmitted on a port since reset, including host and remote
	// management passthrough traffic, and inclusive of all protocol overhead.
	TXBytes int
	// TXErrors shall contain the total number of transmission errors on a port since reset.
	TXErrors int
	// Transceivers shall contain an array of transceiver related metrics for this port. Each member in the array shall
	// represent a single transceiver.
	Transceivers []Transceiver
}

// UnmarshalJSON unmarshals a PortMetrics object from the raw JSON.
func (portmetrics *PortMetrics) UnmarshalJSON(b []byte) error {
	type temp PortMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*portmetrics = PortMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetPortMetrics will get a PortMetrics instance from the service.
func GetPortMetrics(c common.Client, uri string) (*PortMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var portmetrics PortMetrics
	err = json.NewDecoder(resp.Body).Decode(&portmetrics)
	if err != nil {
		return nil, err
	}

	portmetrics.SetClient(c)
	return &portmetrics, nil
}

// ListReferencedPortMetricss gets the collection of PortMetrics from
// a provided reference.
func ListReferencedPortMetricss(c common.Client, link string) ([]*PortMetrics, error) {
	var result []*PortMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, portmetricsLink := range links.ItemLinks {
		portmetrics, err := GetPortMetrics(c, portmetricsLink)
		if err != nil {
			collectionError.Failures[portmetricsLink] = err
		} else {
			result = append(result, portmetrics)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// SAS shall describe physical (phy) related metrics for Serial Attached SCSI (SAS).
type SAS struct {
	// InvalidDwordCount shall contain the number of invalid dwords that have been received by the phy outside of phy
	// reset sequences.
	InvalidDwordCount int
	// LossOfDwordSynchronizationCount shall contain the number of times the phy has restarted the link reset sequence
	// because it lost dword synchronization.
	LossOfDwordSynchronizationCount int
	// RunningDisparityErrorCount shall contain the number of dwords containing running disparity errors that have been
	// received by the phy outside of phy reset sequences.
	RunningDisparityErrorCount int
}

// UnmarshalJSON unmarshals a SAS object from the raw JSON.
func (sas *SAS) UnmarshalJSON(b []byte) error {
	type temp SAS
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sas = SAS(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Transceiver shall describe the transceiver related metrics.
type Transceiver struct {
	// RXInputPowerMilliWatts shall contain the RX input power value of a small form-factor pluggable (SFP)
	// transceiver.
	RXInputPowerMilliWatts float64
	// SupplyVoltage shall contain the supply voltage of a small form-factor pluggable (SFP) transceiver.
	SupplyVoltage float64
	// TXBiasCurrentMilliAmps shall contain the TX bias current value of a small form-factor pluggable (SFP)
	// transceiver.
	TXBiasCurrentMilliAmps float64
	// TXOutputPowerMilliWatts shall contain the TX output power value of a small form-factor pluggable (SFP)
	// transceiver.
	TXOutputPowerMilliWatts float64
}

// UnmarshalJSON unmarshals a Transceiver object from the raw JSON.
func (transceiver *Transceiver) UnmarshalJSON(b []byte) error {
	type temp Transceiver
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*transceiver = Transceiver(t.temp)

	// Extract the links to other entities for later

	return nil
}
