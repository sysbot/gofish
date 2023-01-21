//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ComponentIntegrityType is
type ComponentIntegrityType string

const (
	// SPDMComponentIntegrityType shall indicate the integrity information is obtained through the Security Protocol
	// and Data Model (SPDM) protocol as defined in DMTF DSP0274.
	SPDMComponentIntegrityType ComponentIntegrityType = "SPDM"
	// TPMComponentIntegrityType shall indicate the integrity information is related to a Trusted Platform Module (TPM)
	// as defined by the Trusted Computing Group (TCG).
	TPMComponentIntegrityType ComponentIntegrityType = "TPM"
	// OEMComponentIntegrityType shall indicate the integrity information is OEM-specific and the OEM section may
	// include additional information.
	OEMComponentIntegrityType ComponentIntegrityType = "OEM"
)

// DMTFmeasurementTypes is
type DMTFmeasurementTypes string

const (
	// ImmutableROMDMTFmeasurementTypes Immutable ROM.
	ImmutableROMDMTFmeasurementTypes DMTFmeasurementTypes = "ImmutableROM"
	// MutableFirmwareDMTFmeasurementTypes Mutable firmware or any mutable code.
	MutableFirmwareDMTFmeasurementTypes DMTFmeasurementTypes = "MutableFirmware"
	// HardwareConfigurationDMTFmeasurementTypes Hardware configuration, such as straps.
	HardwareConfigurationDMTFmeasurementTypes DMTFmeasurementTypes = "HardwareConfiguration"
	// FirmwareConfigurationDMTFmeasurementTypes Firmware configuration, such as configurable firmware policy.
	FirmwareConfigurationDMTFmeasurementTypes DMTFmeasurementTypes = "FirmwareConfiguration"
	// MutableFirmwareVersionDMTFmeasurementTypes Mutable firmware version.
	MutableFirmwareVersionDMTFmeasurementTypes DMTFmeasurementTypes = "MutableFirmwareVersion"
	// MutableFirmwareSecurityVersionNumberDMTFmeasurementTypes Mutable firmware security version number.
	MutableFirmwareSecurityVersionNumberDMTFmeasurementTypes DMTFmeasurementTypes = "MutableFirmwareSecurityVersionNumber"
	// MeasurementManifestDMTFmeasurementTypes Measurement Manifest.
	MeasurementManifestDMTFmeasurementTypes DMTFmeasurementTypes = "MeasurementManifest"
)

// MeasurementSpecification is
type MeasurementSpecification string

const (
	// DMTFMeasurementSpecification shall indicate the measurement specification is defined by DMTF in DSP0274.
	DMTFMeasurementSpecification MeasurementSpecification = "DMTF"
)

// SPDMmeasurementSummaryType is
type SPDMmeasurementSummaryType string

const (
	// TCBSPDMmeasurementSummaryType The measurement summary covers the TCB.
	TCBSPDMmeasurementSummaryType SPDMmeasurementSummaryType = "TCB"
	// AllSPDMmeasurementSummaryType The measurement summary covers all measurements in SPDM.
	AllSPDMmeasurementSummaryType SPDMmeasurementSummaryType = "All"
)

// SecureSessionType is
type SecureSessionType string

const (
	// PlainSecureSessionType A plain text session without any protection.
	PlainSecureSessionType SecureSessionType = "Plain"
	// EncryptedAuthenticatedSecureSessionType An established session where both encryption and authentication are
	// protecting the communication.
	EncryptedAuthenticatedSecureSessionType SecureSessionType = "EncryptedAuthenticated"
	// AuthenticatedOnlySecureSessionType An established session where only authentication is protecting the
	// communication.
	AuthenticatedOnlySecureSessionType SecureSessionType = "AuthenticatedOnly"
)

// VerificationStatus is
type VerificationStatus string

const (
	// SuccessVerificationStatus Successful verification.
	SuccessVerificationStatus VerificationStatus = "Success"
	// FailedVerificationStatus Unsuccessful verification.
	FailedVerificationStatus VerificationStatus = "Failed"
)

// CommonAuthInfo shall contain common identity-related authentication information.
type CommonAuthInfo struct {
	// ComponentCertificate shall contain a link to a resource of type Certificate that represents the identify of the
	// component referenced by the TargetComponentURI property.
	ComponentCertificate string
	// VerificationStatus shall contain the status of the verification of the identity of the component referenced by
	// the TargetComponentURI property..
	VerificationStatus VerificationStatus
}

// UnmarshalJSON unmarshals a CommonAuthInfo object from the raw JSON.
func (commonauthinfo *CommonAuthInfo) UnmarshalJSON(b []byte) error {
	type temp CommonAuthInfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*commonauthinfo = CommonAuthInfo(t.temp)

	// Extract the links to other entities for later

	return nil
}

// CommunicationInfo shall contain information about communication between two components.
type CommunicationInfo struct {
	// Sessions shall contain an array of the active sessions or communication channels between two components The
	// active sessions or communication channels do not reflect how future sessions or communication channels are
	// established.
	Sessions []SingleSessionInfo
}

// UnmarshalJSON unmarshals a CommunicationInfo object from the raw JSON.
func (communicationinfo *CommunicationInfo) UnmarshalJSON(b []byte) error {
	type temp CommunicationInfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*communicationinfo = CommunicationInfo(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ComponentIntegrity shall represent critical and pertinent security information about a specific device, system,
// software element, or other managed entity.
type ComponentIntegrity struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// ComponentIntegrityEnabled shall indicate whether security protocols are enabled for the component. If
	// ComponentIntegrityType contains 'SPDM', a value of 'false' shall prohibit the SPDM Requester from using SPDM to
	// communicate with the component identified by the TargetComponentURI property. If ComponentIntegrityType contains
	// 'TPM', a value of 'false' shall disable the TPM component identified by the TargetComponentURI property
	// entirely. If 'false', services shall not provide the TPM and SPDM properties in response payloads for this
	// resource. If 'false', services shall reject action requests to this resource. If 'true', services shall allow
	// security protocols with the component identified by the TargetComponentURI property.
	ComponentIntegrityEnabled string
	// ComponentIntegrityType shall contain the underlying security technology providing integrity information for the
	// component.
	ComponentIntegrityType string
	// ComponentIntegrityTypeVersion shall contain the version of the security technology indicated by the
	// ComponentIntegrityType property. If ComponentIntegrityType contains 'SPDM', this property shall contain the
	// negotiated or selected SPDM protocol and shall follow the regular expression pattern '^\d+\.\d+\.\d+$'. If
	// ComponentIntegrityType contains 'TPM', this property shall contain the version of the TPM.
	ComponentIntegrityTypeVersion string
	// Description provides a description of this resource.
	Description string
	// LastUpdated shall contain the date and time when information for the component was last updated.
	LastUpdated string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SPDM shall contain integrity information about the SPDM Responder identified by the TargetComponentURI property
	// as reported by an SPDM Requester. This property shall be present if ComponentIntegrityType contains 'SPDM' and
	// 'ComponentIntegrityEnabled' contains 'true'. For other cases, this property shall be absent.
	SPDM string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TPM shall contain integrity information about the Trusted Platform Module (TPM) identified by the
	// TargetComponentURI property, This property shall be present if ComponentIntegrityType contains 'TPM' and
	// 'ComponentIntegrityEnabled' contains 'true'. For other cases, this property shall be absent.
	TPM string
	// TargetComponentURI shall contain a link to the resource whose integrity information is reported in this
	// resource. If ComponentIntegrityType contains 'SPDM', this property shall contain a URI to the resource that
	// represents the SPDM Responder. If ComponentIntegrityType contains 'TPM', this property shall contain a URI with
	// RFC6901-defined JSON fragment notation to a member of the TrustedModules array in a ComputerSystem resource that
	// represents the TPM or a resource of type TrustedComponent that represents the TPM.
	TargetComponentURI string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ComponentIntegrity object from the raw JSON.
func (componentintegrity *ComponentIntegrity) UnmarshalJSON(b []byte) error {
	type temp ComponentIntegrity
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*componentintegrity = ComponentIntegrity(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	componentintegrity.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (componentintegrity *ComponentIntegrity) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ComponentIntegrity)
	original.UnmarshalJSON(componentintegrity.rawData)

	readWriteFields := []string{
		"ComponentIntegrityEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(componentintegrity).Elem()

	return componentintegrity.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetComponentIntegrity will get a ComponentIntegrity instance from the service.
func GetComponentIntegrity(c common.Client, uri string) (*ComponentIntegrity, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var componentintegrity ComponentIntegrity
	err = json.NewDecoder(resp.Body).Decode(&componentintegrity)
	if err != nil {
		return nil, err
	}

	componentintegrity.SetClient(c)
	return &componentintegrity, nil
}

// ListReferencedComponentIntegritys gets the collection of ComponentIntegrity from
// a provided reference.
func ListReferencedComponentIntegritys(c common.Client, link string) ([]*ComponentIntegrity, error) {
	var result []*ComponentIntegrity
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, componentintegrityLink := range links.ItemLinks {
		componentintegrity, err := GetComponentIntegrity(c, componentintegrityLink)
		if err != nil {
			collectionError.Failures[componentintegrityLink] = err
		} else {
			result = append(result, componentintegrity)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type Links struct {
	// ComponentsProtected shall contain an array of links to resources that the component identified by the
	// TargetComponentURI property provides integrity protection. This property shall not contain the value of the
	// TargetComponentURI property.
	ComponentsProtected []idRef
	// ComponentsProtected@odata.count
	ComponentsProtectedCount int `json:"ComponentsProtected@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
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

// SPDMGetSignedMeasurementsResponse shall contain the SPDM signed measurements from an SPDM Responder.
type SPDMGetSignedMeasurementsResponse struct {
	// Certificate shall contain a link to a resource of type Certificate that represents the certificate corresponding
	// to the SPDM slot identifier that can be used to validate the signature. This property shall not be present if
	// the SlotId parameter contains the value '15'.
	Certificate string
	// HashingAlgorithm shall contain the hashing algorithm negotiated between the SPDM Requester and the SPDM
	// Responder. The allowable values for this property shall be the hash algorithm names found in the 'BaseHashAlgo'
	// field of the 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is an extended algorithm, this
	// property shall contain the value 'OEM'.
	HashingAlgorithm string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PublicKey shall contain a Privacy Enhanced Mail (PEM)-encoded public key, as defined in section 13 of RFC7468,
	// that can be used to validate the signature. This property shall only be present when the SPDM Requester was pre-
	// provisioned with the SPDM Responder's public key and the SlotId parameter contains the value '15'.
	PublicKey string
	// SignedMeasurements shall contain the cryptographic signed statement over the given nonce and measurement blocks
	// corresponding to the requested measurement indices. If the SPDM version is 1.2, this value shall be a
	// concatenation of SPDM 'VCA' and 'GET_MEASUREMENTS' requests and responses exchanged between the SPDM Requester
	// and the SPDM Responder. If SPDM version is 1.0 or 1.1, this value shall be a concatenation of SPDM
	// 'GET_MEASUREMENTS' requests and responses exchanged between the SPDM Requester and the SPDM Responder. The last
	// 'MEASUREMENTS' response shall contain a signature generated over the 'L2' string by the SPDM Responder.
	SignedMeasurements string
	// SigningAlgorithm shall contain the asymmetric signing algorithm negotiated between the SPDM Requester and the
	// SPDM Responder. The allowable values for this property shall be the asymmetric key signature algorithm names
	// found in the 'BaseAsymAlgo' field of the 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is
	// an extended algorithm, this property shall contain the value 'OEM'.
	SigningAlgorithm string
	// Version shall contain the SPDM version negotiated between the SPDM Requester and the SPDM Responder to generate
	// the cryptographic signed statement. For example, '1.0', '1.1', or '1.2'.
	Version string
}

// UnmarshalJSON unmarshals a SPDMGetSignedMeasurementsResponse object from the raw JSON.
func (spdmgetsignedmeasurementsresponse *SPDMGetSignedMeasurementsResponse) UnmarshalJSON(b []byte) error {
	type temp SPDMGetSignedMeasurementsResponse
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmgetsignedmeasurementsresponse = SPDMGetSignedMeasurementsResponse(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SPDMcommunication shall contain information about communication between two components.
type SPDMcommunication struct {
	// Sessions shall contain an array of the active sessions or communication channels between two components The
	// active sessions or communication channels do not reflect how future sessions or communication channels are
	// established.
	Sessions []SingleSessionInfo
}

// UnmarshalJSON unmarshals a SPDMcommunication object from the raw JSON.
func (spdmcommunication *SPDMcommunication) UnmarshalJSON(b []byte) error {
	type temp SPDMcommunication
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmcommunication = SPDMcommunication(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SPDMidentity shall contain identity authentication information about the SPDM Requester and SPDM Responder.
type SPDMidentity struct {
	// RequesterAuthentication shall contain authentication information of the identity of the SPDM Requester.
	RequesterAuthentication SPDMrequesterAuth
	// ResponderAuthentication shall contain authentication information of the identity of the SPDM Responder.
	ResponderAuthentication SPDMresponderAuth
}

// UnmarshalJSON unmarshals a SPDMidentity object from the raw JSON.
func (spdmidentity *SPDMidentity) UnmarshalJSON(b []byte) error {
	type temp SPDMidentity
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmidentity = SPDMidentity(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SPDMinfo shall contain integrity information about an SPDM Responder as reported by an SPDM Requester.
type SPDMinfo struct {
	// ComponentCommunication shall contain information about communication between the SPDM Requester and SPDM
	// Responder.
	ComponentCommunication SPDMcommunication
	// IdentityAuthentication shall contain identity authentication information about the SPDM Requester and SPDM
	// Responder.
	IdentityAuthentication SPDMidentity
	// MeasurementSet shall contain measurement information for the SPDM Responder.
	MeasurementSet SPDMmeasurementSet
	// Requester shall contain a link to the resource representing the SPDM Responder that is reporting the integrity
	// of the SPDM Responder identified by the TargetComponentURI property.
	Requester string
}

// UnmarshalJSON unmarshals a SPDMinfo object from the raw JSON.
func (spdminfo *SPDMinfo) UnmarshalJSON(b []byte) error {
	type temp SPDMinfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdminfo = SPDMinfo(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SPDMmeasurementSet shall contain SPDM Responder measurement information.
type SPDMmeasurementSet struct {
	// MeasurementSpecification shall contain the measurement specification negotiated between the SPDM Requester and
	// SPDM Responder.
	MeasurementSpecification MeasurementSpecification
	// MeasurementSummary shall contain the Base64-encoded measurement summary using the hash algorithm indicated by
	// the MeasurementSummaryHashAlgorithm property.
	MeasurementSummary string
	// MeasurementSummaryHashAlgorithm shall contain the hash algorithm used to compute the measurement summary. The
	// allowable values for this property shall be the hash algorithm names found in the 'BaseHashAlgo' field of the
	// 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is an extended algorithm, this property
	// shall contain the value 'OEM'.
	MeasurementSummaryHashAlgorithm string
	// MeasurementSummaryType shall contain the type of measurement summary.
	MeasurementSummaryType SPDMmeasurementSummaryType
	// Measurements shall contain measurements from an SPDM Responder.
	Measurements []SPDMsingleMeasurement
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// UnmarshalJSON unmarshals a SPDMmeasurementSet object from the raw JSON.
func (spdmmeasurementset *SPDMmeasurementSet) UnmarshalJSON(b []byte) error {
	type temp SPDMmeasurementSet
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmmeasurementset = SPDMmeasurementSet(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SPDMrequesterAuth shall contain authentication information of the identity of the SPDM Requester.
type SPDMrequesterAuth struct {
	// ProvidedCertificate shall contain a link to a resource of type Certificate that represents the identify of the
	// SPDM Requester provided in mutual authentication.
	ProvidedCertificate string
}

// UnmarshalJSON unmarshals a SPDMrequesterAuth object from the raw JSON.
func (spdmrequesterauth *SPDMrequesterAuth) UnmarshalJSON(b []byte) error {
	type temp SPDMrequesterAuth
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmrequesterauth = SPDMrequesterAuth(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SPDMresponderAuth shall contain common identity-related authentication information.
type SPDMresponderAuth struct {
	// ComponentCertificate shall contain a link to a resource of type Certificate that represents the identify of the
	// component referenced by the TargetComponentURI property.
	ComponentCertificate string
	// VerificationStatus shall contain the status of the verification of the identity of the component referenced by
	// the TargetComponentURI property..
	VerificationStatus VerificationStatus
}

// UnmarshalJSON unmarshals a SPDMresponderAuth object from the raw JSON.
func (spdmresponderauth *SPDMresponderAuth) UnmarshalJSON(b []byte) error {
	type temp SPDMresponderAuth
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmresponderauth = SPDMresponderAuth(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SPDMsingleMeasurement shall contain a single SPDM measurement for an SPDM Responder.
type SPDMsingleMeasurement struct {
	// LastUpdated shall contain the date and time when information for the measurement was last updated.
	LastUpdated string
	// Measurement shall contain the Base64-encoded measurement using the hash algorithm indicated by the
	// MeasurementHashAlgorithm property. This property shall not contain a raw bit stream as a measurement. If the
	// SPDM Responder provides a raw bit stream, the SPDM Requester may apply a hash algorithm to the raw bit stream in
	// order to report the measurement.
	Measurement string
	// MeasurementHashAlgorithm shall contain the hash algorithm used to compute the measurement. The allowable values
	// for this property shall be the hash algorithm names found in the 'BaseHashAlgo' field of the
	// 'NEGOTIATE_ALGORITHMS' request message in DSP0274. If the algorithm is an extended algorithm, this property
	// shall contain the value 'OEM'. This property shall not be present if MeasurementSpecification does not contain
	// 'DMTF'.
	MeasurementHashAlgorithm string
	// MeasurementIndex shall contain the index of the measurement.
	MeasurementIndex int
	// MeasurementType shall contain the type or characteristics of the data that this measurement represents. This
	// property shall not be present if MeasurementSpecification does not contain 'DMTF'.
	MeasurementType DMTFmeasurementTypes
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartofSummaryHash shall indicate if this measurement is part of the measurement summary in the
	// MeasurementSummary property. If this property is not present, it shall be assumed to be 'false'.
	PartofSummaryHash bool
	// SecurityVersionNumber shall contain an 8-byte hex-encoded string of the security version number the measurement
	// represents. This property shall only be present if MeasurementType contains the value
	// 'MutableFirmwareSecurityVersionNumber'.
	SecurityVersionNumber string
}

// UnmarshalJSON unmarshals a SPDMsingleMeasurement object from the raw JSON.
func (spdmsinglemeasurement *SPDMsingleMeasurement) UnmarshalJSON(b []byte) error {
	type temp SPDMsingleMeasurement
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*spdmsinglemeasurement = SPDMsingleMeasurement(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SingleSessionInfo shall contain information about a single communication channel or session between two
// components.
type SingleSessionInfo struct {
	// SessionId shall contain the unique identifier for the active session or communication channel between two
	// components.
	SessionId int
	// SessionType shall contain the type of session or communication channel between two components.
	SessionType SecureSessionType
}

// UnmarshalJSON unmarshals a SingleSessionInfo object from the raw JSON.
func (singlesessioninfo *SingleSessionInfo) UnmarshalJSON(b []byte) error {
	type temp SingleSessionInfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*singlesessioninfo = SingleSessionInfo(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TPMGetSignedMeasurementsResponse shall contain the TPM signed PCR measurements from an TPM.
type TPMGetSignedMeasurementsResponse struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SignedMeasurements shall contain a Base64-encoded cryptographic signed statement generated by the signer. This
	// value shall be the concatenation of the 'quoted' and 'signature' response values of the 'TPM2_Quote' command
	// defined in the Trusted Platform Module Library Specification.
	SignedMeasurements string
}

// UnmarshalJSON unmarshals a TPMGetSignedMeasurementsResponse object from the raw JSON.
func (tpmgetsignedmeasurementsresponse *TPMGetSignedMeasurementsResponse) UnmarshalJSON(b []byte) error {
	type temp TPMGetSignedMeasurementsResponse
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*tpmgetsignedmeasurementsresponse = TPMGetSignedMeasurementsResponse(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TPMauth shall contain common identity-related authentication information.
type TPMauth struct {
	// ComponentCertificate shall contain a link to a resource of type Certificate that represents the identify of the
	// component referenced by the TargetComponentURI property.
	ComponentCertificate string
	// VerificationStatus shall contain the status of the verification of the identity of the component referenced by
	// the TargetComponentURI property..
	VerificationStatus VerificationStatus
}

// UnmarshalJSON unmarshals a TPMauth object from the raw JSON.
func (tpmauth *TPMauth) UnmarshalJSON(b []byte) error {
	type temp TPMauth
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*tpmauth = TPMauth(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TPMcommunication shall contain information about communication between two components.
type TPMcommunication struct {
	// Sessions shall contain an array of the active sessions or communication channels between two components The
	// active sessions or communication channels do not reflect how future sessions or communication channels are
	// established.
	Sessions []SingleSessionInfo
}

// UnmarshalJSON unmarshals a TPMcommunication object from the raw JSON.
func (tpmcommunication *TPMcommunication) UnmarshalJSON(b []byte) error {
	type temp TPMcommunication
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*tpmcommunication = TPMcommunication(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TPMinfo shall contain integrity information about a Trusted Platform Module (TPM).
type TPMinfo struct {
	// ComponentCommunication shall contain information about communication with the TPM.
	ComponentCommunication TPMcommunication
	// IdentityAuthentication shall contain identity authentication information about the TPM.
	IdentityAuthentication TPMauth
	// MeasurementSet shall contain measurement information from the TPM.
	MeasurementSet TPMmeasurementSet
	// NonceSizeBytesMaximum shall contain the maximum number of bytes that can be specified in the Nonce parameter of
	// the TPMGetSignedMeasurements action.
	NonceSizeBytesMaximum int
}

// UnmarshalJSON unmarshals a TPMinfo object from the raw JSON.
func (tpminfo *TPMinfo) UnmarshalJSON(b []byte) error {
	type temp TPMinfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*tpminfo = TPMinfo(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TPMmeasurementSet shall contain Trusted Computing Group TPM measurement information.
type TPMmeasurementSet struct {
	// Measurements shall contain measurements from a TPM.
	Measurements []TPMsingleMeasurement
}

// UnmarshalJSON unmarshals a TPMmeasurementSet object from the raw JSON.
func (tpmmeasurementset *TPMmeasurementSet) UnmarshalJSON(b []byte) error {
	type temp TPMmeasurementSet
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*tpmmeasurementset = TPMmeasurementSet(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TPMsingleMeasurement shall contain a single Trusted Computing Group TPM measurement.
type TPMsingleMeasurement struct {
	// LastUpdated shall contain the date and time when information for the measurement was last updated.
	LastUpdated string
	// Measurement shall contain the Base64-encoded PCR digest using the hashing algorithm indicated by
	// MeasurementHashAlgorithm property.
	Measurement string
	// MeasurementHashAlgorithm shall contain the hash algorithm used to compute the measurement. The allowable values
	// for this property shall be the strings in the 'Algorithm Name' field of the 'TPM_ALG_ID Constants' table within
	// the 'Trusted Computing Group Algorithm Registry'.
	MeasurementHashAlgorithm string
	// PCR shall contain the Platform Configuration Register (PCR) bank of the measurement.
	PCR int
}

// UnmarshalJSON unmarshals a TPMsingleMeasurement object from the raw JSON.
func (tpmsinglemeasurement *TPMsingleMeasurement) UnmarshalJSON(b []byte) error {
	type temp TPMsingleMeasurement
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*tpmsinglemeasurement = TPMsingleMeasurement(t.temp)

	// Extract the links to other entities for later

	return nil
}
