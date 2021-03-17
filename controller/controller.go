package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// UseCase interface
type UseCase interface {
	GetShapeByName(name string) (string, error)
	//SetShapeName(name string) (string, error)
}

// Controller struct
type Controller struct {
	useCase UseCase
	logger  *logrus.Logger
	render  *render.Render
}

// New returns a controller
func New(
	u UseCase,
	logger *logrus.Logger,
	r *render.Render,
) *Controller {
	return &Controller{u, logger, r}
}

// GetShapeByName logic
func (c *Controller) GetShapeByName(w http.ResponseWriter, r *http.Request) {
	shapeName := mux.Vars(r)["shape_name"]

	logger := c.logger.WithFields(logrus.Fields{
		"func":       "Get Shape by Name",
		"shape_name": shapeName,
	})
	logger.Debug("in")

	body, err := c.useCase.GetShapeByName(shapeName)

	// Handle the 404 error

	if err != nil {
		logger.WithError(err).Error("getting shape")

		c.render.Text(w, 200, "Shape not found")
		return
	}

	c.render.Text(w, 200, body)
}
