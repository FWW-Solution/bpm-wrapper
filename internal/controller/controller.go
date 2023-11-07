package controller

import "bpm-wrapper/internal/usecase"

type Controller struct {
	Usecase usecase.Usecase
}

func New() *Controller {
	return &Controller{
		// ...
	}
}
