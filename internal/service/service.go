package service

type Service struct {
	UserService UserService
	PostService PostService
}

func NewService(
	userService UserService,
	postService PostService,
) *Service {
	return &Service{
		UserService: userService,
		PostService: postService,
	}
}
