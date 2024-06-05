package delivery

import (
	h "dataon-test/handler"

	depens "dataon-test/delivery/dependencies"
)

type handler struct {
	orderHandler h.HandlerImpl
}

func SetupHandler(dep depens.Dependency) handler {
	return handler{
		orderHandler:    h.NewOrderHandler(dep.OrderService),
	}
}
