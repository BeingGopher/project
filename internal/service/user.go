package service

type IUser interface {
	Register(username, password, email string) error
	Login(username, password string) (string, error)
}

var (
	iUser IUser
)

func User() IUser {
	return iUser
}

func InitUser(i IUser) {
	iUser = i
}
