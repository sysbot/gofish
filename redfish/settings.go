//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ApplyTime is
type ApplyTime string

const (
	// ImmediateApplyTime shall indicate the values within the settings resource are applied immediately. This value
	// may result in an immediate host reset, manager reset, or other side effects.
	ImmediateApplyTime ApplyTime = "Immediate"
	// OnResetApplyTime shall indicate the values within settings resource are applied when the system or service is
	// reset.
	OnResetApplyTime ApplyTime = "OnReset"
	// AtMaintenanceWindowStartApplyTime shall indicate the values within the settings resource are applied during the
	// maintenance window specified by the MaintenanceWindowStartTime and MaintenanceWindowDurationInSeconds
	// properties. A service can perform resets during this maintenance window.
	AtMaintenanceWindowStartApplyTime ApplyTime = "AtMaintenanceWindowStart"
	// InMaintenanceWindowOnResetApplyTime shall indicate the values within the settings resource are applied during
	// the maintenance window specified by the MaintenanceWindowStartTime and MaintenanceWindowDurationInSeconds
	// properties, and if a reset occurs within the maintenance window.
	InMaintenanceWindowOnResetApplyTime ApplyTime = "InMaintenanceWindowOnReset"
)

// MaintenanceWindow shall indicate that a resource has a maintenance window assignment for applying settings or
// operations. Other resources can link to this object to convey a common control surface for the configuration of
// the maintenance window.
type MaintenanceWindow struct {
	// MaintenanceWindowDurationInSeconds shall indicate the end of the maintenance window as the number of seconds
	// after the time specified by the MaintenanceWindowStartTime property.
	MaintenanceWindowDurationInSeconds string
	// MaintenanceWindowStartTime shall indicate the date and time when the service can start to apply the requested
	// settings or operation as part of a maintenance window. Services shall provide a default value if not configured
	// by a user.
	MaintenanceWindowStartTime string
}

// UnmarshalJSON unmarshals a MaintenanceWindow object from the raw JSON.
func (maintenancewindow *MaintenanceWindow) UnmarshalJSON(b []byte) error {
	type temp MaintenanceWindow
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*maintenancewindow = MaintenanceWindow(t.temp)

	// Extract the links to other entities for later

	return nil
}

// OperationApplyTimeSupport shall indicate that a client can request a specific apply time of a create, delete, or
// action operation of a resource.
type OperationApplyTimeSupport struct {
	// MaintenanceWindowDurationInSeconds shall contain the same as the MaintenanceWindowDurationInSeconds property
	// found in the MaintenanceWindow structure on the MaintenanceWindowResource. This property shall be required if
	// the SupportedValues property contains 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowDurationInSeconds string
	// MaintenanceWindowResource shall contain a link to a resource that contains the @Redfish.MaintenanceWindow
	// property that governs this resource. This property shall be required if the SupportedValues property contains
	// 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowResource string
	// MaintenanceWindowStartTime shall contain the same as the MaintenanceWindowStartTime property found in the
	// MaintenanceWindow structure on the MaintenanceWindowResource. Services shall provide a default value if not
	// configured by a user. This property shall be required if the SupportedValues property contains
	// 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowStartTime string
	// SupportedValues shall indicate the types of apply times the client can request when performing a create, delete,
	// or action operation.
	SupportedValues []OperationApplyTime
}

// UnmarshalJSON unmarshals a OperationApplyTimeSupport object from the raw JSON.
func (operationapplytimesupport *OperationApplyTimeSupport) UnmarshalJSON(b []byte) error {
	type temp OperationApplyTimeSupport
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*operationapplytimesupport = OperationApplyTimeSupport(t.temp)

	// Extract the links to other entities for later

	return nil
}

// PreferredApplyTime shall be specified by client to indicate the preferred time to apply the configuration
// settings.
type PreferredApplyTime struct {
	// ApplyTime shall indicate when to apply the values in this settings resource.
	ApplyTime string
	// MaintenanceWindowDurationInSeconds shall indicate the end of the maintenance window as the number of seconds
	// after the time specified by the MaintenanceWindowStartTime property. This property shall be required if the
	// ApplyTime property is 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowDurationInSeconds string
	// MaintenanceWindowStartTime shall indicate the date and time when the service can start to apply the future
	// configuration as part of a maintenance window. Services shall provide a default value if not configured by a
	// user. This property shall be required if the ApplyTime property is 'AtMaintenanceWindowStart' or
	// 'InMaintenanceWindowOnReset'.
	MaintenanceWindowStartTime string
}

// UnmarshalJSON unmarshals a PreferredApplyTime object from the raw JSON.
func (preferredapplytime *PreferredApplyTime) UnmarshalJSON(b []byte) error {
	type temp PreferredApplyTime
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*preferredapplytime = PreferredApplyTime(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Settings shall describe any settings of a resource.
type Settings struct {
	// ETag shall contain the entity tag (ETag) of the resource to which the settings were applied, after the
	// application. The client can check this value against the ETag of this resource to determine whether the resource
	// had other changes.
	ETag string
	// MaintenanceWindowResource shall contain a link to a resource that contains the @Redfish.MaintenanceWindow
	// property that governs this resource. This property should be supported if the SupportedApplyTimes property
	// contains 'AtMaintenanceWindowStart' or 'InMaintenanceWindowOnReset'.
	MaintenanceWindowResource string
	// Messages shall contain an array of messages associated with the settings.
	Messages []Message
	// SettingsObject shall contain the URI of the resource that the client can PUT or PATCH to modify the resource.
	SettingsObject string
	// SupportedApplyTimes shall contain the supported apply time values a client is allowed to request when
	// configuring the settings apply time. Services that do not support clients configuring the apply time can support
	// this property with a single array member in order to inform the client when the settings will be applied.
	SupportedApplyTimes []ApplyTime
	// Time shall indicate the time when the settings were applied to the resource.
	Time string
}

// UnmarshalJSON unmarshals a Settings object from the raw JSON.
func (settings *Settings) UnmarshalJSON(b []byte) error {
	type temp Settings
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*settings = Settings(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetSettings will get a Settings instance from the service.
func GetSettings(c common.Client, uri string) (*Settings, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var settings Settings
	err = json.NewDecoder(resp.Body).Decode(&settings)
	if err != nil {
		return nil, err
	}

	settings.SetClient(c)
	return &settings, nil
}

// ListReferencedSettingss gets the collection of Settings from
// a provided reference.
func ListReferencedSettingss(c common.Client, link string) ([]*Settings, error) {
	var result []*Settings
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, settingsLink := range links.ItemLinks {
		settings, err := GetSettings(c, settingsLink)
		if err != nil {
			collectionError.Failures[settingsLink] = err
		} else {
			result = append(result, settings)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
