package dto

// TableUsers struct which contains ...
// an array of users
type TableUsers struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

// ResponseMessage ..
type ResponseMessage struct {
	Data    interface{} `json:"data"`
	Error   *string     `json:"error"`
	Success bool        `json:"success"`
}
