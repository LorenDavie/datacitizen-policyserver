package policy

import (
	"datacitizen.me/entities"
	"datacitizen.me/identity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type GradedPermission struct {
	gorm.Model
	Granted bool
	Grade   rune
}

type ScopedGradedPermission struct {
	gorm.Model
	GradedPermission
	Scope string
}

type PermissionException struct {
	gorm.Model
	Granted  bool
	Identity identity.Identity
}

type DataPolicy struct {
	gorm.Model
	Owner             entities.DataCitizen
	GeneralPermission GradedPermission
	ScopedPermissions []ScopedGradedPermission
	Exceptions        []PermissionException
}
