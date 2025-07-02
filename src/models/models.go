package models

type RDPConnectionParams struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type PromptInputParams struct {
	Label        string
	DefaultValue string
	IsPassword   bool
}
