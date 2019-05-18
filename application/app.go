package application

import (
	"fmt"
)

type (
	App interface {
		Start() error
		Stop() error
	}
	app struct {
	}
)

func New() App {
	return &app{
	}
}

func (a *app) Start() error {
	fmt.Println("implement me")
	return nil
}

func (a *app) Stop() error {
	fmt.Println("implement me")
	return nil
}