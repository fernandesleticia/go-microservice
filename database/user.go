type Userstruct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type ApplicationConfig struct {
	Service  string   `json:"service"`
	DataBase User     `json:"database"`
}