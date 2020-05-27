package setup

import "github.com/ChatFalcon/ChatFalcon/router"

func init() {
	router.Router.Static("/setup", "setup/dist")
}
