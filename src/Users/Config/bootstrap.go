package config_users

import (
	_configCore "ddd_golang/src/Core/Infrastructure/Config"
	_userController "ddd_golang/src/Users/Infrastructure/Controller"
)

func Bootstrap(config *_configCore.Config) {
	_userController.NewUserController(config)
}