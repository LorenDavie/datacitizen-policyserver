package entities

import (
	"datacitizen.me/identity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DataCitizen struct {
	gorm.Model
	DefaultIdentity identity.Identity
	Identities      []identity.Identity
}

type DataConsumer struct {
	gorm.Model
	Identity identity.Identity
}
type DataProducer struct {
	gorm.Model
	Identity identity.Identity
}

type DataExchange struct {
	gorm.Model
	Identity identity.Identity
}
