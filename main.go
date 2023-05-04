package main

import (
	"github.com/NSDN/nya-server/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":10127")
}
