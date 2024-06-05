package cmd

import (
	serve "dataon-test/delivery"
	depens "dataon-test/delivery/dependencies"
	"fmt"
)
func Execute(dependency depens.Dependency) {
	handler := serve.SetupHandler(dependency)

	server := serve.Http(handler)
	server.Listen(fmt.Sprintf(":%d", 8000))
}