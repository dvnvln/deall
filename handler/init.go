package handler

type Handler struct {
	controller userController
}

func New(controller userController) *Handler {
	return &Handler{
		controller: controller,
	}
}
