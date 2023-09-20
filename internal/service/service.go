package service

type Service struct {
	userService UserService
	postService PostService
}

func NewService(
	userService UserService,
	postService PostService,
) *Service {
	return &Service{
		userService: userService,
		postService: postService,
	}
}
