package service

type UserService struct {
}

func (u *UserService) Test(i int) int {
	return i + 1
}
