package main

import (
	cmd "dataon-test/cmd"
	depens "dataon-test/delivery/dependencies"
)

func main()  {
	dependency := depens.SetupDependencies()
	cmd.Execute(dependency)
}
