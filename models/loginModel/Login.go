package loginmodel

type RequestLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
    LocationDepartment string `json:"locationDepartment" validate:"required"`
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

type ResponseDataUserLogin struct {
    Username string `json:"username"`
    RoleName []string `json:"roleName"`
}