//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// AttributeType is
type AttributeType string

const (
	// EnumerationAttributeType A list of the known possible enumerated values.
	EnumerationAttributeType AttributeType = "Enumeration"
	// StringAttributeType Free-form text in their values.
	StringAttributeType AttributeType = "String"
	// IntegerAttributeType An integer value.
	IntegerAttributeType AttributeType = "Integer"
	// BooleanAttributeType A flag with a 'true' or 'false' value.
	BooleanAttributeType AttributeType = "Boolean"
	// PasswordAttributeType shall be null in responses.
	PasswordAttributeType AttributeType = "Password"
)

// DependencyType is
type DependencyType string

const (
	// MapDependencyType A simple mapping dependency. If the condition evaluates to 'true', the attribute or state
	// changes to the mapped value.
	MapDependencyType DependencyType = "Map"
)

// MapFromCondition is
type MapFromCondition string

const (
	// EQUMapFromCondition The logical operation for 'Equal'.
	EQUMapFromCondition MapFromCondition = "EQU"
	// NEQMapFromCondition The logical operation for 'Not Equal'.
	NEQMapFromCondition MapFromCondition = "NEQ"
	// GTRMapFromCondition The logical operation for 'Greater than'.
	GTRMapFromCondition MapFromCondition = "GTR"
	// GEQMapFromCondition The logical operation for 'Greater than or Equal'.
	GEQMapFromCondition MapFromCondition = "GEQ"
	// LSSMapFromCondition The logical operation for 'Less than'.
	LSSMapFromCondition MapFromCondition = "LSS"
	// LEQMapFromCondition The logical operation for 'Less than or Equal'.
	LEQMapFromCondition MapFromCondition = "LEQ"
)

// MapFromProperty is
type MapFromProperty string

const (
	// CurrentValueMapFromProperty The dependency on an attribute's CurrentValue.
	CurrentValueMapFromProperty MapFromProperty = "CurrentValue"
	// DefaultValueMapFromProperty The dependency on an attribute's DefaultValue.
	DefaultValueMapFromProperty MapFromProperty = "DefaultValue"
	// ReadOnlyMapFromProperty The dependency on an attribute's ReadOnly state.
	ReadOnlyMapFromProperty MapFromProperty = "ReadOnly"
	// WriteOnlyMapFromProperty The dependency on an attribute's WriteOnly state.
	WriteOnlyMapFromProperty MapFromProperty = "WriteOnly"
	// GrayOutMapFromProperty The dependency on an attribute's GrayOut state.
	GrayOutMapFromProperty MapFromProperty = "GrayOut"
	// HiddenMapFromProperty The dependency on an attribute's Hidden state.
	HiddenMapFromProperty MapFromProperty = "Hidden"
	// LowerBoundMapFromProperty The dependency on an attribute's LowerBound.
	LowerBoundMapFromProperty MapFromProperty = "LowerBound"
	// UpperBoundMapFromProperty The dependency on an attribute's UpperBound.
	UpperBoundMapFromProperty MapFromProperty = "UpperBound"
	// MinLengthMapFromProperty The dependency on an attribute's MinLength.
	MinLengthMapFromProperty MapFromProperty = "MinLength"
	// MaxLengthMapFromProperty The dependency on an attribute's MaxLength.
	MaxLengthMapFromProperty MapFromProperty = "MaxLength"
	// ScalarIncrementMapFromProperty The dependency on an attribute's ScalarIncrement.
	ScalarIncrementMapFromProperty MapFromProperty = "ScalarIncrement"
)

// MapTerms is
type MapTerms string

const (
	// ANDMapTerms The operation used for logical 'AND' of dependency terms.
	ANDMapTerms MapTerms = "AND"
	// ORMapTerms The operation used for logical 'OR' of dependency terms.
	ORMapTerms MapTerms = "OR"
)

// MapToProperty is
type MapToProperty string

const (
	// CurrentValueMapToProperty The dependency that affects an attribute's CurrentValue.
	CurrentValueMapToProperty MapToProperty = "CurrentValue"
	// DefaultValueMapToProperty The dependency that affects an attribute's DefaultValue.
	DefaultValueMapToProperty MapToProperty = "DefaultValue"
	// ReadOnlyMapToProperty The dependency that affects an attribute's ReadOnly state.
	ReadOnlyMapToProperty MapToProperty = "ReadOnly"
	// WriteOnlyMapToProperty The dependency that affects an attribute's WriteOnly state.
	WriteOnlyMapToProperty MapToProperty = "WriteOnly"
	// GrayOutMapToProperty The dependency that affects an attribute's GrayOut state.
	GrayOutMapToProperty MapToProperty = "GrayOut"
	// HiddenMapToProperty The dependency that affects an attribute's Hidden state.
	HiddenMapToProperty MapToProperty = "Hidden"
	// ImmutableMapToProperty The dependency that affects an attribute's Immutable state.
	ImmutableMapToProperty MapToProperty = "Immutable"
	// HelpTextMapToProperty The dependency that affects an attribute's HelpText.
	HelpTextMapToProperty MapToProperty = "HelpText"
	// WarningTextMapToProperty The dependency that affects an attribute's WarningText.
	WarningTextMapToProperty MapToProperty = "WarningText"
	// DisplayNameMapToProperty The dependency that affects an attribute's DisplayName.
	DisplayNameMapToProperty MapToProperty = "DisplayName"
	// DisplayOrderMapToProperty The dependency that affects an attribute's DisplayName.
	DisplayOrderMapToProperty MapToProperty = "DisplayOrder"
	// LowerBoundMapToProperty The dependency that affects an attribute's LowerBound.
	LowerBoundMapToProperty MapToProperty = "LowerBound"
	// UpperBoundMapToProperty The dependency that affects an attribute's UpperBound.
	UpperBoundMapToProperty MapToProperty = "UpperBound"
	// MinLengthMapToProperty The dependency that affects an attribute's MinLength.
	MinLengthMapToProperty MapToProperty = "MinLength"
	// MaxLengthMapToProperty The dependency that affects an attribute's MaxLength.
	MaxLengthMapToProperty MapToProperty = "MaxLength"
	// ScalarIncrementMapToProperty The dependency that affects an attribute's ScalarIncrement.
	ScalarIncrementMapToProperty MapToProperty = "ScalarIncrement"
	// ValueExpressionMapToProperty The dependency that affects an attribute's ValueExpression.
	ValueExpressionMapToProperty MapToProperty = "ValueExpression"
)

// AttributeRegistry shall represent an attribute registry for a Redfish implementation.
type AttributeRegistry struct {
	common.Entity
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// Language shall contain an RFC5646-conformant language code.
	Language string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OwningEntity shall represent the publisher of this attribute registry.
	OwningEntity string
	// RegistryEntries shall list attributes for this component, along with their possible values, dependencies, and
	// other metadata.
	RegistryEntries string
	// RegistryVersion shall contain the version of this attribute registry.
	RegistryVersion string
	// SupportedSystems shall contain an array containing a list of systems that this attribute registry supports.
	SupportedSystems []SupportedSystems
}

// UnmarshalJSON unmarshals a AttributeRegistry object from the raw JSON.
func (attributeregistry *AttributeRegistry) UnmarshalJSON(b []byte) error {
	type temp AttributeRegistry
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*attributeregistry = AttributeRegistry(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetAttributeRegistry will get a AttributeRegistry instance from the service.
func GetAttributeRegistry(c common.Client, uri string) (*AttributeRegistry, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var attributeregistry AttributeRegistry
	err = json.NewDecoder(resp.Body).Decode(&attributeregistry)
	if err != nil {
		return nil, err
	}

	attributeregistry.SetClient(c)
	return &attributeregistry, nil
}

// ListReferencedAttributeRegistrys gets the collection of AttributeRegistry from
// a provided reference.
func ListReferencedAttributeRegistrys(c common.Client, link string) ([]*AttributeRegistry, error) {
	var result []*AttributeRegistry
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, attributeregistryLink := range links.ItemLinks {
		attributeregistry, err := GetAttributeRegistry(c, attributeregistryLink)
		if err != nil {
			collectionError.Failures[attributeregistryLink] = err
		} else {
			result = append(result, attributeregistry)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// AttributeValue shall describe a possible enumeration attribute value.
type AttributeValue struct {
	// ValueDisplayName shall contain a string representing the user-readable display string of the value for the
	// attribute in the defined language.
	ValueDisplayName string
	// ValueName shall contain a string representing the value name for the attribute. ValueName is a unique string
	// within the list of possible values in the Value array for an attribute.
	ValueName string
}

// UnmarshalJSON unmarshals a AttributeValue object from the raw JSON.
func (attributevalue *AttributeValue) UnmarshalJSON(b []byte) error {
	type temp AttributeValue
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*attributevalue = AttributeValue(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Attributes shall describe an attribute and its possible values and other metadata.
type Attributes struct {
	// AttributeName shall contain the name of this attribute that is unique in this attribute registry.
	AttributeName string
	// CurrentValue shall contain the placeholder of the current value for the attribute, to aid in evaluating
	// dependencies. The evaluation results of the Dependencies array may affect the current attribute value.
	CurrentValue float64
	// DefaultValue shall contain the default value for the attribute.
	DefaultValue float64
	// DisplayName shall contain the user-readable display string for the attribute in the defined language.
	DisplayName string
	// DisplayOrder shall contain the ascending order, as a number, in which this attribute appears relative to other
	// attributes.
	DisplayOrder int
	// GrayOut shall indicate whether this attribute is grayed out. A grayed-out attribute is not active and is grayed
	// out in user interfaces but the attribute value can be modified. The evaluation results of the Dependencies array
	// may affect the grayed-out state of an attribute.
	GrayOut bool
	// HelpText shall contain the help text for the attribute.
	HelpText string
	// Hidden shall indicate whether this attribute is hidden in user interfaces. The evaluation results of the
	// Dependencies array may affect the hidden state of an attribute.
	Hidden bool
	// Immutable shall indicate whether this attribute is immutable. Immutable attributes shall not be modified and
	// typically reflect a hardware state.
	Immutable bool
	// IsSystemUniqueProperty shall indicate whether this attribute is unique.
	IsSystemUniqueProperty bool
	// LowerBound shall contain a number indicating the lower limit for an integer attribute.
	LowerBound int
	// MaxLength shall contain the maximum character length of an attribute of the String type.
	MaxLength int
	// MenuPath shall contain the menu hierarchy of this attribute, in the form of a path to the menu names. It shall
	// start with './' to indicate the root menu, followed by the menu names with '/' characters to delineate the menu
	// traversal.
	MenuPath string
	// MinLength shall contain a number indicating the minimum character length of an attribute of the String type.
	MinLength int
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ReadOnly shall indicate whether this attribute is read-only. A read-only attribute cannot be modified, and
	// should be grayed out in user interfaces. The evaluation results of the Dependencies array may affect the read-
	// only state of an attribute.
	ReadOnly bool
	// ResetRequired shall indicate whether a system or device reset is required for this attribute value change to
	// take effect.
	ResetRequired bool
	// ScalarIncrement shall contain a number indicating the amount to increment or decrement an integer attribute each
	// time a user requests a value change. The '0' value indicates a free-form numeric user-input attribute.
	ScalarIncrement int
	// Type shall contain an enumeration that describes the attribute type.
	Type string
	// UefiDevicePath shall contain the UEFI Specification-defined UEFI device path that qualifies and locates this
	// device for this attribute.
	UefiDevicePath string
	// UefiKeywordName shall contain the UEFI Specification-defined keyword for this attribute.
	UefiKeywordName string
	// UefiNamespaceId shall contain the UEFI Specification-defined namespace ID for this attribute.
	UefiNamespaceId string
	// UpperBound shall contain a number indicating the upper limit for an integer attribute.
	UpperBound int
	// Value shall contain an array containing the possible values of an attribute of the Enumeration type.
	Value []AttributeValue
	// ValueExpression shall contain a valid regular expression, according to the Perl regular expression dialect, that
	// validates the attribute value. Applies to only string and integer attributes.
	ValueExpression string
	// WarningText shall contain the warning text for the attribute.
	WarningText string
	// WriteOnly shall indicate whether this attribute is write-only. A write-only attribute reverts to its initial
	// value after settings are applied.
	WriteOnly bool
}

// UnmarshalJSON unmarshals a Attributes object from the raw JSON.
func (attributes *Attributes) UnmarshalJSON(b []byte) error {
	type temp Attributes
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*attributes = Attributes(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Dependencies shall describe a dependency of attributes on this component.
type Dependencies struct {
	// Dependency shall contain the dependency expression for one or more attributes in this attribute registry.
	Dependency string
	// DependencyFor shall contain the AttributeName of the attribute whose change triggers the evaluation of this
	// dependency expression.
	DependencyFor string
	// Type shall contain an enumeration that describes the type for the attribute dependency.
	Type string
}

// UnmarshalJSON unmarshals a Dependencies object from the raw JSON.
func (dependencies *Dependencies) UnmarshalJSON(b []byte) error {
	type temp Dependencies
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*dependencies = Dependencies(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Dependency shall describe the dependency expression for one or more attributes in this attribute registry.
type Dependency struct {
	// MapFrom shall contain an array containing the map-from conditions for a dependency of the Map type.
	MapFrom []MapFrom
	// MapToAttribute shall contain the AttributeName of the attribute that is affected by this dependency expression.
	MapToAttribute string
	// MapToProperty shall contain the metadata property for the attribute that the MapFromAttribute property specifies
	// that evaluates this dependency expression. For example, this value could be the MapFromAttribute CurrentValue or
	// ReadOnly state.
	MapToProperty string
	// MapToValue The value that the property in MapToProperty in the attribute specified in MapToAttribute changes to
	// if the dependency expression evaluates to 'true'.
	MapToValue float64
}

// UnmarshalJSON unmarshals a Dependency object from the raw JSON.
func (dependency *Dependency) UnmarshalJSON(b []byte) error {
	type temp Dependency
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*dependency = Dependency(t.temp)

	// Extract the links to other entities for later

	return nil
}

// MapFrom shall describe a map-from condition for a dependency of the Map type.
type MapFrom struct {
	// MapFromAttribute shall contain the AttributeName for the attribute to use to evaluate this dependency expression
	// term.
	MapFromAttribute string
	// MapFromCondition shall contain the condition to use to evaluate this dependency expression. For example, 'EQU'
	// or 'NEQ'.
	MapFromCondition string
	// MapFromProperty shall contain the metadata property for the attribute that the MapFromAttribute property
	// specifies to use to evaluate this dependency expression. For example, this value could be the MapFromAttribute
	// CurrentValue, or ReadOnly state.
	MapFromProperty string
	// MapFromValue The value that the property in MapFromProperty in the attribute in MapFromAttribute to use to
	// evaluate this dependency expression.
	MapFromValue float64
	// MapTerms shall contain the logical term that combines two or more MapFrom conditions in this dependency
	// expression. For example, 'AND' for logical AND, or 'OR' for logical OR. If multiple logical terms are present in
	// a dependency expression, they should be evaluated in array order, meaning they are evaluated left-to-right when
	// displayed as a logic expression.
	MapTerms string
}

// UnmarshalJSON unmarshals a MapFrom object from the raw JSON.
func (mapfrom *MapFrom) UnmarshalJSON(b []byte) error {
	type temp MapFrom
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*mapfrom = MapFrom(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Menus shall describe an attribute's menu and its hierarchy.
type Menus struct {
	// DisplayName shall contain the user-readable display string of the menu in the defined language.
	DisplayName string
	// DisplayOrder shall contain the ascending order, as a number, in which this menu appears relative to other menus.
	DisplayOrder int
	// GrayOut shall indicate whether this menu is grayed out. A grayed-only menu is not accessible in user interfaces.
	GrayOut bool
	// Hidden shall indicate whether this menu is hidden in user interfaces. The evaluation results of the Dependencies
	// array may affect the hidden state of a menu.
	Hidden bool
	// MenuName shall contain the name of this menu that is unique in this attribute registry.
	MenuName string
	// MenuPath shall contain the menu hierarchy of this menu, in the form of a path to the menu names. It shall start
	// with './' to indicate the root menu, followed by the menu names with '/' characters to delineate the menu
	// traversal.
	MenuPath string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ReadOnly shall indicate whether this menu is read-only. A read-only menu is not accessible in user interfaces,
	// and all properties contained in that menu and its sub-menus are read-only.
	ReadOnly bool
}

// UnmarshalJSON unmarshals a Menus object from the raw JSON.
func (menus *Menus) UnmarshalJSON(b []byte) error {
	type temp Menus
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*menus = Menus(t.temp)

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

// RegistryEntries shall describe a list of all attributes for this component, along with their possible values,
// dependencies, and other metadata.
type RegistryEntries struct {
	// Attributes shall contain an array containing the attributes and their possible values and other metadata in the
	// attribute registry.
	Attributes []Attributes
	// Dependencies shall contain an array containing a list of dependencies of attributes on this component.
	Dependencies []Dependencies
	// Menus shall contain an array containing the attributes menus and their hierarchy in the attribute registry.
	Menus []Menus
}

// UnmarshalJSON unmarshals a RegistryEntries object from the raw JSON.
func (registryentries *RegistryEntries) UnmarshalJSON(b []byte) error {
	type temp RegistryEntries
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*registryentries = RegistryEntries(t.temp)

	// Extract the links to other entities for later

	return nil
}

// SupportedSystems shall describe a system that this attribute registry supports.
type SupportedSystems struct {
	// FirmwareVersion The version of the component firmware image to which this attribute registry applies.
	FirmwareVersion string
	// ProductName shall contain the product name of the computer system to which this attribute registry applies.
	ProductName string
	// SystemId shall contain the system ID that identifies the systems to which this attribute registry applies. This
	// can be identified by one or more properties in the computer system resource, such as Model, SubModel, or SKU.
	SystemId string
}

// UnmarshalJSON unmarshals a SupportedSystems object from the raw JSON.
func (supportedsystems *SupportedSystems) UnmarshalJSON(b []byte) error {
	type temp SupportedSystems
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*supportedsystems = SupportedSystems(t.temp)

	// Extract the links to other entities for later

	return nil
}
