//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)



// BitRate is
type BitRate string

const (
    // 1200BitRate A bit rate of 1200 bit/s.
    1200BitRate BitRate = "1200"
    // 2400BitRate A bit rate of 2400 bit/s.
    2400BitRate BitRate = "2400"
    // 4800BitRate A bit rate of 4800 bit/s.
    4800BitRate BitRate = "4800"
    // 9600BitRate A bit rate of 9600 bit/s.
    9600BitRate BitRate = "9600"
    // 19200BitRate A bit rate of 19200 bit/s.
    19200BitRate BitRate = "19200"
    // 38400BitRate A bit rate of 38400 bit/s.
    38400BitRate BitRate = "38400"
    // 57600BitRate A bit rate of 57600 bit/s.
    57600BitRate BitRate = "57600"
    // 115200BitRate A bit rate of 115200 bit/s.
    115200BitRate BitRate = "115200"
    // 230400BitRate A bit rate of 230400 bit/s.
    230400BitRate BitRate = "230400"
)

// ConnectorType is
type ConnectorType string

const (
    // RJ45ConnectorType An RJ45 connector.
    RJ45ConnectorType ConnectorType = "RJ45"
    // RJ11ConnectorType An RJ11 connector.
    RJ11ConnectorType ConnectorType = "RJ11"
    // DB9FemaleConnectorType A DB9 Female connector.
    DB9FemaleConnectorType ConnectorType = "DB9 Female"
    // DB9MaleConnectorType A DB9 Male connector.
    DB9MaleConnectorType ConnectorType = "DB9 Male"
    // DB25FemaleConnectorType A DB25 Female connector.
    DB25FemaleConnectorType ConnectorType = "DB25 Female"
    // DB25MaleConnectorType A DB25 Male connector.
    DB25MaleConnectorType ConnectorType = "DB25 Male"
    // USBConnectorType A USB connector.
    USBConnectorType ConnectorType = "USB"
    // mUSBConnectorType A mUSB connector.
    mUSBConnectorType ConnectorType = "mUSB"
    // uUSBConnectorType A uUSB connector.
    uUSBConnectorType ConnectorType = "uUSB"
)

// DataBits is
type DataBits string

const (
    // 5DataBits Five bits of data following the start bit.
    5DataBits DataBits = "5"
    // 6DataBits Six bits of data following the start bit.
    6DataBits DataBits = "6"
    // 7DataBits Seven bits of data following the start bit.
    7DataBits DataBits = "7"
    // 8DataBits Eight bits of data following the start bit.
    8DataBits DataBits = "8"
)

// FlowControl is
type FlowControl string

const (
    // NoneFlowControl No flow control imposed.
    NoneFlowControl FlowControl = "None"
    // SoftwareFlowControl XON/XOFF in-band flow control imposed.
    SoftwareFlowControl FlowControl = "Software"
    // HardwareFlowControl Out-of-band flow control imposed.
    HardwareFlowControl FlowControl = "Hardware"
)

// Parity is
type Parity string

const (
    // NoneParity No parity bit.
    NoneParity Parity = "None"
    // EvenParity An even parity bit.
    EvenParity Parity = "Even"
    // OddParity An odd parity bit.
    OddParity Parity = "Odd"
    // MarkParity A mark parity bit.
    MarkParity Parity = "Mark"
    // SpaceParity A space parity bit.
    SpaceParity Parity = "Space"
)

// PinOut is
type PinOut string

const (
    // CiscoPinOut The Cisco pinout configuration.
    CiscoPinOut PinOut = "Cisco"
    // CycladesPinOut The Cyclades pinout configuration.
    CycladesPinOut PinOut = "Cyclades"
    // DigiPinOut The Digi pinout configuration.
    DigiPinOut PinOut = "Digi"
)

// SignalType is
type SignalType string

const (
    // Rs232SignalType The serial interface follows RS232.
    Rs232SignalType SignalType = "Rs232"
    // Rs485SignalType The serial interface follows RS485.
    Rs485SignalType SignalType = "Rs485"
)

// StopBits is
type StopBits string

const (
    // 1StopBits One stop bit following the data bits.
    1StopBits StopBits = "1"
    // 2StopBits Two stop bits following the data bits.
    2StopBits StopBits = "2"
)
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


// SerialInterface shall represent a serial interface as part of the Redfish Specification.
type SerialInterface struct {
    common.Entity
    // ODataContext is the odata context.
    ODataContext  string `json:"@odata.context"`
    // ODataEtag is the odata etag.
    ODataEtag  string `json:"@odata.etag"`
    // ODataType is the odata type.
    ODataType  string `json:"@odata.type"`
    // Actions shall contain the available actions for this resource.
    Actions  string
    // BitRate shall indicate the transmit and receive speed of the serial connection.
    BitRate  string
    // ConnectorType shall indicate the type of physical connector used for this serial connection.
    ConnectorType  string
    // DataBits shall indicate number of data bits for the serial connection.
    DataBits  string
    // Description provides a description of this resource.
    Description  string
    // FlowControl shall indicate the flow control mechanism for the serial connection.
    FlowControl  string
    // InterfaceEnabled shall indicate whether this interface is enabled.
    InterfaceEnabled  bool
    // Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
// Redfish Specification-described requirements.
    OEM  json.RawMessage `json:"Oem"`
    // Parity shall indicate parity information for a serial connection.
    Parity  string
    // PinOut shall indicate the physical pinout for the serial connector.
    PinOut  PinOut
    // SignalType shall contain the type of serial signaling in use for the serial connection.
    SignalType  string
    // StopBits shall indicate the stop bits for the serial connection.
    StopBits  string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SerialInterface object from the raw JSON.
func (serialinterface *SerialInterface) UnmarshalJSON(b []byte) error {
    type temp SerialInterface
    var t struct {
        temp
    }

    err := json.Unmarshal(b, &t)
    if err != nil {
        return err
    }

    *serialinterface = SerialInterface(t.temp)

    // Extract the links to other entities for later


	// This is a read/write object, so we need to save the raw object data for later
	serialinterface.rawData = b

    return nil
}
// Update commits updates to this object's properties to the running system.
func (serialinterface *SerialInterface) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(SerialInterface)
	original.UnmarshalJSON(serialinterface.rawData)

	readWriteFields := []string{
        "BitRate",
        "DataBits",
        "FlowControl",
        "InterfaceEnabled",
        "Parity",
        "StopBits",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(serialinterface).Elem()

	return serialinterface.Entity.Update(originalElement, currentElement, readWriteFields)
}


// GetSerialInterface will get a SerialInterface instance from the service.
func GetSerialInterface(c common.Client, uri string) (*SerialInterface, error) {
    resp, err := c.Get(uri)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var serialinterface SerialInterface
    err = json.NewDecoder(resp.Body).Decode(&serialinterface)
    if err != nil {
        return nil, err
    }

    serialinterface.SetClient(c)
    return &serialinterface, nil
}

// ListReferencedSerialInterfaces gets the collection of SerialInterface from
// a provided reference.
func ListReferencedSerialInterfaces(c common.Client, link string) ([]*SerialInterface, error) {
    var result []*SerialInterface
    if link == "" {
        return result, nil
    }

    links, err := common.GetCollection(c, link)
    if err != nil {
        return result, err
    }

    collectionError := common.NewCollectionError()
    for _, serialinterfaceLink := range links.ItemLinks {
        serialinterface, err := GetSerialInterface(c, serialinterfaceLink)
        if err != nil {
            collectionError.Failures[serialinterfaceLink] = err
        } else {
            result = append(result, serialinterface)
        }
    }

    if collectionError.Empty() {
        return result, nil
    } else {
        return result, collectionError
    }
}



