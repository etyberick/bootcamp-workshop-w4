package usecase

import (
	"errors"

	"github.com/varopxndx/bootcamp-workshop-w4/service"
)

// UseCase struct for communities flows
type UseCase struct {
	service service.ShapeClient
}

// New community UseCase
func New(service service.ShapeClient) *UseCase {
	return &UseCase{service}
}

// GetShapeByName - Returns the information according to the shape name
func (u *UseCase) GetShapeByName(shapeName string) (string, error) {
	switch shapeName {
	case "square":
		return u.service.Square.Info(), nil
	case "circle":
		return u.service.Circle.Info(), nil
	case "hexagon":
		return u.service.Hexagon.Info(), nil
	default:
		return "", errors.New("Shape not found")
	}
}

// SetShapeName - Returns the employee by id
// func (u *UseCase) SetShapeName(employeeID string) (string, error) {
// 	return "", nil
// }
