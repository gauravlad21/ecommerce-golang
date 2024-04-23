package common

type UserType string

const (
	Admin  UserType = "gms"
	Normal UserType = "kgs"
)

func (u UserType) String() string {
	switch u {
	case Admin:
		return "admin"
	case Normal:
		return "normal"
	default:
		return "Unknown"
	}
}

func StringToUserType(str string) UserType {
	switch str {
	case "admin":
		return Admin
	case "normal":
		return Normal
	default:
		return ""
	}
}

func GetUserType() []UserType {
	return []UserType{
		Admin,
		Normal,
	}
}
