package identity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Identity struct {
	gorm.Model
	Signature string // fully qualified identifier, e.g. code@axilent.com or AxilentEngine@@Twitter
	Context   string // context of id, e.g. email or Twitter
}
