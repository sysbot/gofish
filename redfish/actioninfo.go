//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ParameterTypes is
type ParameterTypes string

const (
	// BooleanParameterTypes A boolean.
	BooleanParameterTypes ParameterTypes = "Boolean"
	// NumberParameterTypes A number.
	NumberParameterTypes ParameterTypes = "Number"
	// NumberArrayParameterTypes An array of numbers.
	NumberArrayParameterTypes ParameterTypes = "NumberArray"
	// StringParameterTypes A string.
	StringParameterTypes ParameterTypes = "String"
	// StringArrayParameterTypes An array of strings.
	StringArrayParameterTypes ParameterTypes = "StringArray"
	// ObjectParameterTypes An embedded JSON object.
	ObjectParameterTypes ParameterTypes = "Object"
	// ObjectArrayParameterTypes An array of JSON objects.
	ObjectArrayParameterTypes ParameterTypes = "ObjectArray"
)

// ActionInfo shall represent the supported parameters and other information for a Redfish action on a target
// within a Redfish implementation. Supported parameters can differ among vendors and even among resource
// instances. This data can ensure that action requests from applications contain supported parameters.
type ActionInfo struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Parameters shall list the parameters included in the specified Redfish action for this resource.
	Parameters []Parameters
}

// UnmarshalJSON unmarshals a ActionInfo object from the raw JSON.
func (actioninfo *ActionInfo) UnmarshalJSON(b []byte) error {
	type temp ActionInfo
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*actioninfo = ActionInfo(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetActionInfo will get a ActionInfo instance from the service.
func GetActionInfo(c common.Client, uri string) (*ActionInfo, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var actioninfo ActionInfo
	err = json.NewDecoder(resp.Body).Decode(&actioninfo)
	if err != nil {
		return nil, err
	}

	actioninfo.SetClient(c)
	return &actioninfo, nil
}

// ListReferencedActionInfos gets the collection of ActionInfo from
// a provided reference.
func ListReferencedActionInfos(c common.Client, link string) ([]*ActionInfo, error) {
	var result []*ActionInfo
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, actioninfoLink := range links.ItemLinks {
		actioninfo, err := GetActionInfo(c, actioninfoLink)
		if err != nil {
			collectionError.Failures[actioninfoLink] = err
		} else {
			result = append(result, actioninfo)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// Parameters shall contain information about a parameter included in a Redfish action for this resource.
type Parameters struct {
	common.Entity
	// AllowableNumbers shall indicate the allowable numeric values, inclusive ranges of values, and incremental step
	// values values for this parameter as applied to this action target, as defined in the 'Allowable values for
	// numbers and durations' clause of the Redfish Specification. For arrays, this property shall represent the
	// allowable values for each array member. This property shall only be present for numeric parameters or string
	// parameters that specify a duration.
	AllowableNumbers []string
	// AllowablePattern shall contain a regular expression that describes the allowable values for this parameter as
	// applied to this action target. For arrays, this property shall represent the allowable values for each array
	// member. This property shall only be present for string parameters.
	AllowablePattern string
	// AllowableValues shall indicate the allowable values for this parameter as applied to this action target. For
	// arrays, this property shall represent the allowable values for each array member.
	AllowableValues []string
	// ArraySizeMaximum shall contain the maximum number of array elements that this service supports for this
	// parameter. This property shall not be present for non-array parameters.
	ArraySizeMaximum int
	// ArraySizeMinimum shall contain the minimum number of array elements required by this service for this parameter.
	// This property shall not be present for non-array parameters.
	ArraySizeMinimum int
	// DataType shall contain the JSON property type for this parameter.
	DataType ParameterTypes
	// MaximumValue shall contain the maximum value that this service supports. For arrays, this property shall
	// represent the maximum value for each array member. This property shall not be present for non-integer or number
	// parameters.
	MaximumValue float64
	// MinimumValue shall contain the minimum value that this service supports. For arrays, this property shall
	// represent the minimum value for each array member. This property shall not be present for non-integer or number
	// parameters.
	MinimumValue float64
	// ObjectDataType shall describe the entity type definition in '@odata.type' format for the parameter. This
	// property shall be required for parameters with a data type of 'Object' or 'ObjectArray', and shall not be
	// present for parameters with other data types.
	ObjectDataType string
	// Required shall indicate whether the parameter is required to complete this action.
	Required string
}

// UnmarshalJSON unmarshals a Parameters object from the raw JSON.
func (parameters *Parameters) UnmarshalJSON(b []byte) error {
	type temp Parameters
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*parameters = Parameters(t.temp)

	// Extract the links to other entities for later

	return nil
}
