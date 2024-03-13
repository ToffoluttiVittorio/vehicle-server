package vehicle

import (
	"net/http"
	"strconv"

	"github.com/ToffoluttiVittorio/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	idParse, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		d.logger.Error(
			"Could not find vehicle from store",
			zap.Error(err),
		)
		return
	}

	vehicle, err := d.store.Vehicle().Delete(r.Context(), idParse)

	if err != nil {
		d.logger.Error(
			"Could not delete vehicle from store",
			zap.Error(err),
		)
		return
	}

	if vehicle {
		rw.WriteHeader(http.StatusNoContent)
		return
	}

	rw.WriteHeader(http.StatusNotFound)

	// http.Error(rw, "Not Implemented", http.StatusInternalServerError)
}
