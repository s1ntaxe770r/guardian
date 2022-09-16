package main

import (
	"net/http"

	"github.com/s1ntaxe770r/guardian/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/secrets", handlers.GetSecret)
	logrus.Info("servering stuff on 0.0.0.0:8080")
	logrus.Fatal(http.ListenAndServe(":8080", r))
}
