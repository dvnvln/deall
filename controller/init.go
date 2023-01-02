package controller

type Controller struct {
	repo userRepository
}

func New(repo userRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}
