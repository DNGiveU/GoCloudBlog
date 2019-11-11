package login

type Login struct {
	Password string 	`json:"password"`
	UserName string 	`json:"userName"`
	Type string			`json:"type"`
}
