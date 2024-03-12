package storage

import (
	"github.com/ToffoluttiVittorio/vehicle-server/storage/vehiclestore"
)

type Store interface {
	Vehicle() vehiclestore.Store
}
