//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Ethernet shall describe the Ethernet related network function metrics.
type Ethernet struct {
	// NumOffloadedIPv4Conns shall contain the total number of offloaded TCP/IPv4 connections.
	NumOffloadedIPv4Conns int
	// NumOffloadedIPv6Conns shall contain the total number of offloaded TCP/IPv6 connections.
	NumOffloadedIPv6Conns int
}

// UnmarshalJSON unmarshals a Ethernet object from the raw JSON.
func (ethernet *Ethernet) UnmarshalJSON(b []byte) error {
	type temp Ethernet
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ethernet = Ethernet(t.temp)

	// Extract the links to other entities for later

	return nil
}

// FibreChannel shall describe the Fibre Channel related network function metrics.
type FibreChannel struct {
	// PortLoginAccepts shall contain the total number of PLOGI ACC responses received by this Fibre Channel function.
	PortLoginAccepts int
	// PortLoginRejects shall contain the total number of PLOGI RJT responses received by this Fibre Channel function.
	PortLoginRejects int
	// PortLoginRequests shall contain the total number of PLOGI requests sent by this function.
	PortLoginRequests int
	// RXCongestionFPINs shall contain the total number of Congestion FPINs received by this Fibre Channel function.
	RXCongestionFPINs int
	// RXDeliveryFPINs shall contain the total number of Delivery FPINs received by this Fibre Channel function.
	RXDeliveryFPINs int
	// RXExchanges shall contain the total number of Fibre Channel exchanges received.
	RXExchanges int
	// RXLinkIntegrityFPINs shall contain the total number of Link Integrity FPINs received by this Fibre Channel
	// function.
	RXLinkIntegrityFPINs int
	// RXPeerCongestionFPINs shall contain the total number of Peer Congestion FPINs received by this Fibre Channel
	// function.
	RXPeerCongestionFPINs int
	// RXSequences shall contain the total number of Fibre Channel sequences received.
	RXSequences int
	// TXCongestionFPINs shall contain the total number of Congestion FPINs sent by this Fibre Channel function.
	TXCongestionFPINs int
	// TXDeliveryFPINs shall contain the total number of Delivery FPINs sent by this Fibre Channel function.
	TXDeliveryFPINs int
	// TXExchanges shall contain the total number of Fibre Channel exchanges transmitted.
	TXExchanges int
	// TXLinkIntegrityFPINs shall contain the total number of Link Integrity FPINs sent by this Fibre Channel function.
	TXLinkIntegrityFPINs int
	// TXPeerCongestionFPINs shall contain the total number of Peer Congestion FPINs sent by this Fibre Channel
	// function.
	TXPeerCongestionFPINs int
	// TXSequences shall contain the total number of Fibre Channel sequences transmitted.
	TXSequences int
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

// NetworkDeviceFunctionMetrics shall represent the network metrics for a single network function of a network
// adapter in a Redfish implementation.
type NetworkDeviceFunctionMetrics struct {
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
	// Ethernet shall contain network function metrics specific to Ethernet adapters.
	Ethernet string
	// FibreChannel shall contain network function metrics specific to Fibre Channel adapters.
	FibreChannel string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RXAvgQueueDepthPercent shall contain the average RX queue depth as the percentage.
	RXAvgQueueDepthPercent float64
	// RXBytes shall contain the total number of bytes received on a network function, inclusive of all protocol
	// overhead.
	RXBytes int
	// RXFrames shall contain the total number of frames received on a network function.
	RXFrames int
	// RXMulticastFrames shall contain the total number of good multicast frames received on a network function since
	// reset, including host and remote management passthrough traffic.
	RXMulticastFrames int
	// RXQueuesEmpty shall indicate whether nothing is in a network function's RX queues to DMA.
	RXQueuesEmpty bool
	// RXQueuesFull shall contain the number of RX queues that are full.
	RXQueuesFull int
	// RXUnicastFrames shall contain the total number of good unicast frames received on a network function since
	// reset.
	RXUnicastFrames int
	// TXAvgQueueDepthPercent shall contain the average TX queue depth as the percentage.
	TXAvgQueueDepthPercent float64
	// TXBytes shall contain the total number of bytes sent on a network function, inclusive of all protocol overhead.
	TXBytes int
	// TXFrames shall contain the total number of frames sent on a network function.
	TXFrames int
	// TXMulticastFrames shall contain the total number of good multicast frames transmitted on a network function
	// since reset, including host and remote management passthrough traffic.
	TXMulticastFrames int
	// TXQueuesEmpty shall indicate whether all TX queues for a network function are empty.
	TXQueuesEmpty bool
	// TXQueuesFull shall contain the number of TX queues that are full.
	TXQueuesFull int
	// TXUnicastFrames shall contain the total number of good unicast frames transmitted on a network function since
	// reset, including host and remote management passthrough traffic.
	TXUnicastFrames int
}

// UnmarshalJSON unmarshals a NetworkDeviceFunctionMetrics object from the raw JSON.
func (networkdevicefunctionmetrics *NetworkDeviceFunctionMetrics) UnmarshalJSON(b []byte) error {
	type temp NetworkDeviceFunctionMetrics
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*networkdevicefunctionmetrics = NetworkDeviceFunctionMetrics(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetNetworkDeviceFunctionMetrics will get a NetworkDeviceFunctionMetrics instance from the service.
func GetNetworkDeviceFunctionMetrics(c common.Client, uri string) (*NetworkDeviceFunctionMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var networkdevicefunctionmetrics NetworkDeviceFunctionMetrics
	err = json.NewDecoder(resp.Body).Decode(&networkdevicefunctionmetrics)
	if err != nil {
		return nil, err
	}

	networkdevicefunctionmetrics.SetClient(c)
	return &networkdevicefunctionmetrics, nil
}

// ListReferencedNetworkDeviceFunctionMetricss gets the collection of NetworkDeviceFunctionMetrics from
// a provided reference.
func ListReferencedNetworkDeviceFunctionMetricss(c common.Client, link string) ([]*NetworkDeviceFunctionMetrics, error) {
	var result []*NetworkDeviceFunctionMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, networkdevicefunctionmetricsLink := range links.ItemLinks {
		networkdevicefunctionmetrics, err := GetNetworkDeviceFunctionMetrics(c, networkdevicefunctionmetricsLink)
		if err != nil {
			collectionError.Failures[networkdevicefunctionmetricsLink] = err
		} else {
			result = append(result, networkdevicefunctionmetrics)
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
