package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/whatvn/denny"
	middleware "github.com/whatvn/denny/middleware/http"
	"log"
	"net/http"
)

func main() {
	server := denny.NewServer(true)

	server.Use(middleware.Logger())

	server.Controller("wrapgin", http.MethodGet, denny.WrapGin(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "wrap gin.HandlerFunc",
		})
	}))

	server.Controller("wrapf", http.MethodGet, denny.WrapF(func(w http.ResponseWriter, r *http.Request) {
		response(w, "wrap http.HandlerFunc")
	}))

	server.Controller("wraph", http.MethodGet, denny.WrapH(&handler{}))

	server.GinHandle("ginhandle", http.MethodGet, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "wrap gin.HandlerFunc",
		})
	})

	server.HTTPHandle("httphandle", http.MethodGet, &handler{})

	if err := server.GraceFulStart(":8080"); err != nil {
		log.Fatal(err)
	}
}

type handler struct {

}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response(w, "wrap http.Handler")
}

func response(w http.ResponseWriter, v string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	b, _ := json.Marshal(map[string]string{
		"hello": v,
	})
	_, _ = w.Write(b)
}