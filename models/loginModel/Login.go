package loginmodel

type RequestLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLDAPData struct {
    ID       string
    Email    string
    Name     string
    FullName string
	TelephoneNumber string
}