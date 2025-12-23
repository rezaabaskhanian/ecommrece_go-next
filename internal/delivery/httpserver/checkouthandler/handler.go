package checkouthandler

import checkoutservcie "github.com/rezaabaskhanian/ecommrece_go-next.git/internal/usecase/checkoutservice"

type Handler struct {
	checkoutSvc checkoutservcie.Service
}

func New(checkoutSvc checkoutservcie.Service) Handler {
	return Handler{checkoutSvc: checkoutSvc}
}
