package user_service

type UserService interface {
	CreateUser(request UserRequest) (UserResponse, error)
	GetUser(request UserRequest) (UserResponse, error)
	UpdateUser(request UserRequest) (UserResponse, error)
	DeleteUser(request UserRequest) (UserResponse, error)

	// APIs for Admin Only!
	GetUsers(request UserRequest) (UserResponse, error)
	GetUsersByRole(request UserRequest) (UserResponse, error)
	DeleteUsers(request UserRequest) (UserResponse, error)
}
