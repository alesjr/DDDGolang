package main

import (
	_configCore "ddd_golang/src/Core/Infrastructure/Config"
	_coreController "ddd_golang/src/Core/Infrastructure/Controller"
	_configUsers "ddd_golang/src/Users/Config"
)

func main()  {
	c := _configCore.Bootstrap()
	_coreController.NewDefaultController(c)

	_configUsers.Bootstrap(c)
	c.AddRoute()
}
