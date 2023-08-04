package services

type IUserService interface {
	UserCreate() error
}

type userService struct {
}
