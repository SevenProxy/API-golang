package client

type UserUseCaseImpl struct {
	UserRepository PropsUser
}

func UserUseCase(userRepository PropsUser) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		UserRepository: userRepository,
	}
}