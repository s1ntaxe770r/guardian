package models

type SecretResponse struct {
	Msg     string            `json:"msg"`
	Secrets map[string]string `json:"secrets"`
}

// type EncryptRequest struct {
// 	Msg string `json:"msg"`
// }

// type EncryptResponse struct {
// 	Msg string `json:"msg"`
// }
