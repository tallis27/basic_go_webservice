package main

import (
	"net/http"

	"github.com/tallis27/basic_go_webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":5000", nil)
}
