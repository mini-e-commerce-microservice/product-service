package product

type service struct{}

type ServiceOption struct {
}

func New(opt ServiceOption) *service {
	return &service{}
}
