//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// PrivilegeType is
type PrivilegeType string

const (
	// LoginPrivilegeType Can log in to the service and read Resources.
	LoginPrivilegeType PrivilegeType = "Login"
	// ConfigureManagerPrivilegeType Can configure managers.
	ConfigureManagerPrivilegeType PrivilegeType = "ConfigureManager"
	// ConfigureUsersPrivilegeType Can configure users and their accounts.
	ConfigureUsersPrivilegeType PrivilegeType = "ConfigureUsers"
	// ConfigureSelfPrivilegeType Can change the password for the current user account and log out of their own
	// sessions.
	ConfigureSelfPrivilegeType PrivilegeType = "ConfigureSelf"
	// ConfigureComponentsPrivilegeType Can configure components that this service manages.
	ConfigureComponentsPrivilegeType PrivilegeType = "ConfigureComponents"
	// NoAuthPrivilegeType shall be used to indicate an operation does not require authentication. This privilege shall
	// not be used in Redfish Roles.
	NoAuthPrivilegeType PrivilegeType = "NoAuth"
	// ConfigureCompositionInfrastructurePrivilegeType shall be used to indicate the user can view and configure
	// composition service resources without matching the Client property in the ResourceBlock or
	// CompositionReservation resources.
	ConfigureCompositionInfrastructurePrivilegeType PrivilegeType = "ConfigureCompositionInfrastructure"
	// AdministrateSystemsPrivilegeType Adminsitrator for systems found in the systems collection. Able to manage boot
	// configuration, keys, and certificates for systems.
	AdministrateSystemsPrivilegeType PrivilegeType = "AdministrateSystems"
	// OperateSystemsPrivilegeType Operator for systems found in the systems colletion. Able to perform resets and
	// configure interfaces.
	OperateSystemsPrivilegeType PrivilegeType = "OperateSystems"
	// AdministrateStoragePrivilegeType Administrator for storage subsystems and storage systems found in the storage
	// collection and storage system collection respectively.
	AdministrateStoragePrivilegeType PrivilegeType = "AdministrateStorage"
	// OperateStorageBackupPrivilegeType Operator for storage backup functionality for storage subsystems and storage
	// systems found in the storage collection and storage system collection respectively.
	OperateStorageBackupPrivilegeType PrivilegeType = "OperateStorageBackup"
)
