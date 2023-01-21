//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AccountProviderTypes is
type AccountProviderTypes string

const (
	// RedfishServiceAccountProviderTypes shall be a DMTF Redfish Specification-conformant service. The
	// ServiceAddresses format shall contain a set of URIs that correspond to a Redfish account service.
	RedfishServiceAccountProviderTypes AccountProviderTypes = "RedfishService"
	// ActiveDirectoryServiceAccountProviderTypes shall be a Microsoft Active Directory Technical Specification-
	// conformant service. The ServiceAddresses format shall contain a set of fully qualified domain names (FQDN) or
	// NetBIOS names that links to the set of domain servers for the Active Directory service.
	ActiveDirectoryServiceAccountProviderTypes AccountProviderTypes = "ActiveDirectoryService"
	// LDAPServiceAccountProviderTypes shall be an RFC4511-conformant service. The ServiceAddresses format shall
	// contain a set of fully qualified domain names (FQDN) that links to the set of LDAP servers for the service.
	LDAPServiceAccountProviderTypes AccountProviderTypes = "LDAPService"
	// OEMAccountProviderTypes An OEM-specific external authentication or directory service.
	OEMAccountProviderTypes AccountProviderTypes = "OEM"
	// TACACSplusAccountProviderTypes shall be an RFC8907-conformant service. The ServiceAddresses format shall contain
	// a set of host:port that correspond to a TACACS+ service and where the format for host and port are defined in
	// RFC3986.
	TACACSplusAccountProviderTypes AccountProviderTypes = "TACACSplus"
	// OAuth2AccountProviderTypes shall be an RFC6749-conformant service. The ServiceAddresses format shall contain a
	// set of URIs that correspond to the RFC8414-defined metadata for the OAuth 2.0 service.
	OAuth2AccountProviderTypes AccountProviderTypes = "OAuth2"
)

// AuthenticationTypes is
type AuthenticationTypes string

const (
	// TokenAuthenticationTypes An opaque authentication token.
	TokenAuthenticationTypes AuthenticationTypes = "Token"
	// KerberosKeytabAuthenticationTypes A Kerberos keytab.
	KerberosKeytabAuthenticationTypes AuthenticationTypes = "KerberosKeytab"
	// UsernameAndPasswordAuthenticationTypes A user name and password combination.
	UsernameAndPasswordAuthenticationTypes AuthenticationTypes = "UsernameAndPassword"
	// OEMAuthenticationTypes An OEM-specific authentication mechanism.
	OEMAuthenticationTypes AuthenticationTypes = "OEM"
)

// LocalAccountAuth is
type LocalAccountAuth string

const (
	// EnabledLocalAccountAuth shall authenticate users based on the account service-defined manager accounts resource
	// collection.
	EnabledLocalAccountAuth LocalAccountAuth = "Enabled"
	// DisabledLocalAccountAuth shall never authenticate users based on the account service-defined manager accounts
	// resource collection.
	DisabledLocalAccountAuth LocalAccountAuth = "Disabled"
	// FallbackLocalAccountAuth shall authenticate users based on the account service-defined manager accounts resource
	// collection only if any external account providers are currently unreachable.
	FallbackLocalAccountAuth LocalAccountAuth = "Fallback"
	// LocalFirstLocalAccountAuth shall first authenticate users based on the account service-defined manager accounts
	// resource collection. If authentication fails, the service shall authenticate by using external account
	// providers.
	LocalFirstLocalAccountAuth LocalAccountAuth = "LocalFirst"
)

// OAuth2Mode is
type OAuth2Mode string

const (
	// DiscoveryOAuth2Mode shall indicate the service performs token validation from information found at the URIs
	// specified by the ServiceAddresses property. Services shall implement a caching method of this information so
	// it's not necessary to retrieve metadata and key information for every request containing a token.
	DiscoveryOAuth2Mode OAuth2Mode = "Discovery"
	// OfflineOAuth2Mode shall indicate the service performs token validation from properties configured by a client.
	// Clients should configure the Issuer and OAuthServiceSigningKeys properties for this mode.
	OfflineOAuth2Mode OAuth2Mode = "Offline"
)

// TACACSplusPasswordExchangeProtocol is
type TACACSplusPasswordExchangeProtocol string

const (
	// ASCIITACACSplusPasswordExchangeProtocol shall indicate the ASCII Login flow as described under section 5.4.2 of
	// RFC8907.
	ASCIITACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "ASCII"
	// PAPTACACSplusPasswordExchangeProtocol shall indicate the PAP Login flow as described under section 5.4.2 of
	// RFC8907.
	PAPTACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "PAP"
	// CHAPTACACSplusPasswordExchangeProtocol shall indicate the CHAP Login flow as described under section 5.4.2 of
	// RFC8907.
	CHAPTACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "CHAP"
	// MSCHAPv1TACACSplusPasswordExchangeProtocol shall indicate the MS-CHAP v1 Login flow as described under section
	// 5.4.2 of RFC8907.
	MSCHAPv1TACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "MSCHAPv1"
	// MSCHAPv2TACACSplusPasswordExchangeProtocol shall indicate the MS-CHAP v2 Login flow as described under section
	// 5.4.2 of RFC8907.
	MSCHAPv2TACACSplusPasswordExchangeProtocol TACACSplusPasswordExchangeProtocol = "MSCHAPv2"
)

// AccountService shall represent an account service for a Redfish implementation. The properties are common to,
// and enable management of, all user accounts. The properties include the password requirements and control
// features, such as account lockout. Properties and actions in this service specify general behavior that should
// be followed for typical accounts, however implementations may override these behaviors for special accounts or
// situations to avoid denial of service or other deadlock situations.
type AccountService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccountLockoutCounterResetAfter shall contain the period of time, in seconds, from the last failed login attempt
	// when the AccountLockoutThreshold counter, which counts the number of failed login attempts, is reset to '0'.
	// Then, AccountLockoutThreshold failures are required before the account is locked. This value shall be less than
	// or equal to the AccountLockoutDuration value. The threshold counter also resets to '0' after each successful
	// login. If the AccountLockoutCounterResetEnabled value is 'false', this property shall be ignored.
	AccountLockoutCounterResetAfter string
	// AccountLockoutCounterResetEnabled shall indicate whether the threshold counter is reset after the
	// AccountLockoutCounterResetAfter expires. If 'true', it is reset. If 'false', only a successful login resets the
	// threshold counter and if the user reaches the AccountLockoutThreshold limit, the account shall be locked out
	// indefinitely and only an administrator-issued reset clears the threshold counter. If this property is absent,
	// the default is 'true'.
	AccountLockoutCounterResetEnabled string
	// AccountLockoutDuration shall contain the period of time, in seconds, that an account is locked after the number
	// of failed login attempts reaches the AccountLockoutThreshold value, within the AccountLockoutCounterResetAfter
	// window of time. The value shall be greater than or equal to the AccountLockoutCounterResetAfter value. If this
	// value is '0', no lockout shall occur. If AccountLockoutCounterResetEnabled value is 'false', this property shall
	// be ignored.
	AccountLockoutDuration int
	// AccountLockoutThreshold shall contain the threshold of failed login attempts before a user account is locked. If
	// '0', the account shall never be locked.
	AccountLockoutThreshold int
	// Accounts shall contain a link to a resource collection of type ManagerAccountCollection.
	Accounts string
	// Actions shall contain the available actions for this resource.
	Actions string
	// ActiveDirectory shall contain the first Active Directory external account provider that this account service
	// supports. If the account service supports one or more Active Directory services as an external account provider,
	// this entity shall be populated by default. This entity shall not be present in the additional external account
	// providers resource collection.
	ActiveDirectory string
	// AdditionalExternalAccountProviders shall contain a link to a resource collection of type
	// ExternalAccountProviderCollection that represents the additional external account providers that this account
	// service uses.
	AdditionalExternalAccountProviders string
	// AuthFailureLoggingThreshold shall contain the threshold for when an authorization failure is logged. Logging
	// shall occur after every 'n' occurrences of an authorization failure on the same account, where 'n' represents
	// the value of this property. If the value is '0', logging of authorization failures shall be disabled.
	AuthFailureLoggingThreshold string
	// Description provides a description of this resource.
	Description string
	// LDAP shall contain the first LDAP external account provider that this account service supports. If the account
	// service supports one or more LDAP services as an external account provider, this entity shall be populated by
	// default. This entity shall not be present in the additional external account providers resource collection.
	LDAP string
	// LocalAccountAuth shall govern how the service uses the manager accounts resource collection within this account
	// service as part of authentication. The enumerated values describe the details for each mode.
	LocalAccountAuth string
	// MaxPasswordLength shall contain the maximum password length that the implementation allows for this account
	// service. This property does not apply to accounts from external account providers.
	MaxPasswordLength string
	// MinPasswordLength shall contain the minimum password length that the implementation allows for this account
	// service. This property does not apply to accounts from external account providers.
	MinPasswordLength string
	// OAuth2 shall contain the first OAuth 2.0 external account provider that this account service supports. If the
	// account service supports one or more OAuth 2.0 services as an external account provider, this entity shall be
	// populated by default. This entity shall not be present in the additional external account providers resource
	// collection.
	OAuth2 ExternalAccountProvider
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PasswordExpirationDays shall contain the number of days before account passwords in this account service will
	// expire. The value shall be applied during account creation and password modification unless the
	// PasswordExpiration property is provided. The value 'null' shall indicate that account passwords never expire.
	// This property does not apply to accounts from external account providers.
	PasswordExpirationDays int
	// PrivilegeMap shall contain a link to a resource of type PrivilegeMapping that contains the privileges that are
	// required for a user context to complete a requested operation on a URI associated with this service.
	PrivilegeMap string
	// RestrictedOemPrivileges shall contain an array of OEM privileges that are restricted by the service.
	RestrictedOemPrivileges []string
	// RestrictedPrivileges shall contain an array of Redfish privileges that are restricted by the service.
	RestrictedPrivileges []PrivilegeType
	// Roles shall contain a link to a resource collection of type RoleCollection.
	Roles string
	// ServiceEnabled shall indicate whether the account service is enabled. If 'true', it is enabled. If 'false', it
	// is disabled and users cannot be created, deleted, or modified, and new sessions cannot be started. However,
	// established sessions may still continue to run. Any service, such as the session service, that attempts to
	// access the disabled account service fails. However, this does not affect HTTP Basic Authentication connections.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedAccountTypes shall contain an array of the account types supported by the service.
	SupportedAccountTypes []AccountTypes
	// SupportedOEMAccountTypes shall contain an array of the OEM account types supported by the service.
	SupportedOEMAccountTypes []string
	// TACACSplus shall contain the first TACACS+ external account provider that this account service supports. If the
	// account service supports one or more TACACS+ services as an external account provider, this entity shall be
	// populated by default. This entity shall not be present in the additional external account providers resource
	// collection.
	TACACSplus ExternalAccountProvider
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AccountService object from the raw JSON.
func (accountservice *AccountService) UnmarshalJSON(b []byte) error {
	type temp AccountService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*accountservice = AccountService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	accountservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (accountservice *AccountService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(AccountService)
	original.UnmarshalJSON(accountservice.rawData)

	readWriteFields := []string{
		"AccountLockoutCounterResetAfter",
		"AccountLockoutCounterResetEnabled",
		"AccountLockoutDuration",
		"AccountLockoutThreshold",
		"AuthFailureLoggingThreshold",
		"LocalAccountAuth",
		"MaxPasswordLength",
		"MinPasswordLength",
		"PasswordExpirationDays",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(accountservice).Elem()

	return accountservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetAccountService will get a AccountService instance from the service.
func GetAccountService(c common.Client, uri string) (*AccountService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var accountservice AccountService
	err = json.NewDecoder(resp.Body).Decode(&accountservice)
	if err != nil {
		return nil, err
	}

	accountservice.SetClient(c)
	return &accountservice, nil
}

// ListReferencedAccountServices gets the collection of AccountService from
// a provided reference.
func ListReferencedAccountServices(c common.Client, link string) ([]*AccountService, error) {
	var result []*AccountService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, accountserviceLink := range links.ItemLinks {
		accountservice, err := GetAccountService(c, accountserviceLink)
		if err != nil {
			collectionError.Failures[accountserviceLink] = err
		} else {
			result = append(result, accountservice)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Authentication shall contain the information required to authenticate to the external service.
type Authentication struct {
	// AuthenticationType shall contain the type of authentication used to connect to the external account provider.
	AuthenticationType AuthenticationTypes
	// EncryptionKey shall contain the value of a symmetric encryption key for account services that support some form
	// of encryption, obfuscation, or authentication such as TACACS+. The value shall be 'null' in responses. The
	// property shall accept a hexadecimal string whose length depends on the external account service, such as
	// TACACS+. A TACACS+ service shall use this property to specify the secret key as defined in RFC8907.
	EncryptionKey string
	// EncryptionKeySet shall contain 'true' if a valid value was provided for the EncryptionKey property. Otherwise,
	// the property shall contain 'false'. For a TACACS+ service, the value 'false' shall indicate data obfuscation, as
	// defined in section 4.5 of RFC8907, is disabled.
	EncryptionKeySet bool
	// KerberosKeytab shall contain a Base64-encoded version of the Kerberos keytab for this service. A PATCH or PUT
	// operation writes the keytab. The value shall be 'null' in responses.
	KerberosKeytab string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Password shall contain the password for this service. A PATCH or PUT operation writes the password. The value
	// shall be 'null' in responses.
	Password string
	// Token shall contain the token for this service. A PATCH or PUT operation writes the token. The value shall be
	// 'null' in responses.
	Token string
	// Username shall contain the user name for this service.
	Username string
}

// UnmarshalJSON unmarshals a Authentication object from the raw JSON.
func (authentication *Authentication) UnmarshalJSON(b []byte) error {
	type temp Authentication
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*authentication = Authentication(t.temp)

	// Extract the links to other entities for later

	return nil
}

// ExternalAccountProvider shall contain properties that represent external account provider services that can
// provide accounts for this manager to use for authentication.
type ExternalAccountProvider struct {
	// Authentication shall contain the authentication information for the external account provider.
	Authentication string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates the external account provider uses.
	Certificates string
	// LDAPService shall contain any additional mapping information needed to parse a generic LDAP service. This
	// property should only be present inside the LDAP property.
	LDAPService string
	// OAuth2Service shall contain additional information needed to parse an OAuth 2.0 service. This property should
	// only be present inside an OAuth2 property.
	OAuth2Service OAuth2Service
	// PasswordSet shall contain 'true' if a valid value was provided for the Password property. Otherwise, the
	// property shall contain 'false'.
	PasswordSet string
	// Priority shall contain the assigned priority for the specified external account provider. The value '0' value
	// shall indicate the highest priority. Increasing values shall represent decreasing priority. If an external
	// provider does not have a priority assignment or two or more external providers have the same priority, the
	// behavior shall be determined by the Redfish service. The priority is used to determine the order of
	// authentication and authorization for each external account provider.
	Priority int
	// RemoteRoleMapping shall contain a set of the mapping rules that are used to convert the external account
	// providers account information to the local Redfish role.
	RemoteRoleMapping []RoleMapping
	// ServiceAddresses shall contain the addresses of the account providers to which this external account provider
	// links. The format of this field depends on the type of external account provider. Each item in the array shall
	// contain a single address. Services can define their own behavior for managing multiple addresses.
	ServiceAddresses []string
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// TACACSplusService shall contain additional information needed to parse a TACACS+ services. This property should
	// only be present inside a TACACSplus property.
	TACACSplusService TACACSplusService
}

// UnmarshalJSON unmarshals a ExternalAccountProvider object from the raw JSON.
func (externalaccountprovider *ExternalAccountProvider) UnmarshalJSON(b []byte) error {
	type temp ExternalAccountProvider
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*externalaccountprovider = ExternalAccountProvider(t.temp)

	// Extract the links to other entities for later

	return nil
}

// LDAPSearchSettings shall contain all required settings to search a generic LDAP service.
type LDAPSearchSettings struct {
	// BaseDistinguishedNames shall contain an array of base distinguished names to use to search an external LDAP
	// service.
	BaseDistinguishedNames []string
	// GroupNameAttribute shall contain the attribute name that contains the LDAP group name.
	GroupNameAttribute string
	// GroupsAttribute shall contain the attribute name that contains the groups for an LDAP user entry.
	GroupsAttribute string
	// SSHKeyAttribute shall contain the attribute name that contains the LDAP user's SSH public key.
	SSHKeyAttribute string
	// UsernameAttribute shall contain the attribute name that contains the LDAP user name.
	UsernameAttribute string
}

// UnmarshalJSON unmarshals a LDAPSearchSettings object from the raw JSON.
func (ldapsearchsettings *LDAPSearchSettings) UnmarshalJSON(b []byte) error {
	type temp LDAPSearchSettings
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ldapsearchsettings = LDAPSearchSettings(t.temp)

	// Extract the links to other entities for later

	return nil
}

// LDAPService shall contain all required settings to parse a generic LDAP service.
type LDAPService struct {
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SearchSettings shall contain the required settings to search an external LDAP service.
	SearchSettings string
}

// UnmarshalJSON unmarshals a LDAPService object from the raw JSON.
func (ldapservice *LDAPService) UnmarshalJSON(b []byte) error {
	type temp LDAPService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ldapservice = LDAPService(t.temp)

	// Extract the links to other entities for later

	return nil
}

// OAuth2Service shall contain settings for parsing an OAuth 2.0 service.
type OAuth2Service struct {
	// Audience shall contain an array of allowable RFC7519-defined audience strings of the Redfish service. The values
	// shall uniquely identify the Redfish service. For example, a MAC address or UUID for the manager can uniquely
	// identify the service.
	Audience []string
	// Issuer shall contain the RFC8414-defined issuer string of the OAuth 2.0 service. If the Mode property contains
	// the value 'Discovery', this property shall contain the value of the 'issuer' string from the OAuth 2.0 service's
	// metadata and this property shall be read-only. Clients should configure this property if Mode contains
	// 'Offline'.
	Issuer string
	// Mode shall contain the mode of operation for token validation.
	Mode string
	// OAuthServiceSigningKeys shall contain a Base64-encoded string of the RFC7517-defined signing keys of the issuer
	// of the OAuth 2.0 service. If the Mode property contains the value 'Discovery', this property shall contain the
	// keys found at the URI specified by the 'jwks_uri' string from the OAuth 2.0 service's metadata and this property
	// shall be read-only. Clients should configure this property if Mode contains 'Offline'.
	OAuthServiceSigningKeys string
}

// UnmarshalJSON unmarshals a OAuth2Service object from the raw JSON.
func (oauth2service *OAuth2Service) UnmarshalJSON(b []byte) error {
	type temp OAuth2Service
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*oauth2service = OAuth2Service(t.temp)

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

// RoleMapping shall contain mapping rules that are used to convert the external account providers account
// information to the local Redfish role.
type RoleMapping struct {
	// LocalRole shall contain the RoleId property value within a role resource on this Redfish service to which to map
	// the remote user or group.
	LocalRole string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RemoteGroup shall contain the name of the remote group, or the remote role in the case of a Redfish service,
	// that maps to the local Redfish role to which this entity links.
	RemoteGroup string
	// RemoteUser shall contain the name of the remote user that maps to the local Redfish role to which this entity
	// links.
	RemoteUser string
}

// UnmarshalJSON unmarshals a RoleMapping object from the raw JSON.
func (rolemapping *RoleMapping) UnmarshalJSON(b []byte) error {
	type temp RoleMapping
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*rolemapping = RoleMapping(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TACACSplusService shall contain settings for parsing a TACACS+ service.
type TACACSplusService struct {
	// PasswordExchangeProtocols shall indicate all the allowed TACACS+ password exchange protocol described under
	// section 5.4.2 of RFC8907.
	PasswordExchangeProtocols []TACACSplusPasswordExchangeProtocol
	// PrivilegeLevelArgument shall specify the name of the argument in a TACACS+ Authorization REPLY packet body, as
	// defined in RFC8907, that contains the user's privilege level.
	PrivilegeLevelArgument string
}

// UnmarshalJSON unmarshals a TACACSplusService object from the raw JSON.
func (tacacsplusservice *TACACSplusService) UnmarshalJSON(b []byte) error {
	type temp TACACSplusService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*tacacsplusservice = TACACSplusService(t.temp)

	// Extract the links to other entities for later

	return nil
}
