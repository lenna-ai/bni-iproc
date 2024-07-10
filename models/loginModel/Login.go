package loginmodel

type RequestLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
    Ldap bool `json:"isLdap" validate:"required"`
}

type UserLDAPData struct {
    ID       string
    Email    string
    Name     string
    FullName string
	TelephoneNumber string
}

type DataUserLogin struct {
    Username string `json:"username"`
	Password string `json:"password"`
    RoleName []string `json:"roleName"`
}