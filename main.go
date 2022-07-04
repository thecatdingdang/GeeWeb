package main

import (
	"GeeWeb/Gee"

	"net/http"
)
func main(){
	engine := Gee.New()
	engine.GET("/", func(c *Gee.Context) {
		c.JSON(http.StatusOK,Gee.H{
			"a":"b",
		})
	})
	engine.Run(":8000")
}
