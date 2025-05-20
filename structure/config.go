package structure

const ServerAddress = "127.0.0.1"
const ServerPort = 3000
const MaxBodyBytes = 10 << 20 // 10 MB
const MainAdminUsername = "Administrator"
const MainAdminDefaultPassword = "password123"

type JsonConfiguration struct {
	Version           string `json:"version"`
	Host              string `json:"host"`
	Port              int    `json:"port"`
	MaxBodyBytes      int    `json:"maxBodyBytes"`
	AdminUsername     string `json:"adminUsername"`
	AdminPasswordHash string `json:"adminPasswordHash"`
}
