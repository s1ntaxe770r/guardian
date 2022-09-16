package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/s1ntaxe770r/guardian/models"
	"github.com/sirupsen/logrus"
)

func GetSecret(rw http.ResponseWriter, r *http.Request) {
	encryptionKey := os.Getenv("GUARDIAN_ENCRYPTION_KEY")
	message := os.Getenv("GUARDIAN_MESSAGE")

	if encryptionKey == "" && message == "" {
		logrus.Error("How unfortunate. I hold no secrets ")
		resp := models.SecretResponse{
			Msg:     "I hold no secrets",
			Secrets: map[string]string{},
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
		json.NewEncoder(rw).Encode(resp)
		return
	}
	logrus.Info("Whats this?")
	maps := map[string]string{}
	maps["encryption_key"] = encryptionKey
	maps["lets hear"] = message
	resp := models.SecretResponse{
		Msg:     "Egads!! I have some secrets, I must inform you at once",
		Secrets: maps,
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(resp)
}
