package main

import "github.com/geo-albin/bookStore/pkg/routes"

func main() {
	routes.AddRoutes()
	routes.StartAndServe()
}
