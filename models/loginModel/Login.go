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

type ADCodeMessage struct {
    Code string
    Message string
    IsSuccess bool
}

type UnitRole struct {
    GroupRoleName string
    KodeUnit int
}

type RoleMenuView struct {
    ROLE_ID int
    ROLE_NAME string
    MENU_ID int
    MENU_NAME string
    DISPLAY_NAME string
    AKSES string
}


func (RoleMenuView) TableName() string {
	return "ROLE_MENU_VIEW"
}
