package detailmodel

type Type struct {
	Id         string
	Code       string
	Name       string
	Is_deleted string
}

func (Type) TableName() string {
	return "JENIS_PENGADAAN"
}
