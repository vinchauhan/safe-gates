package service

type Service struct {
	codec string
}

func New(codec string) *Service {
	return &Service{codec: codec}
}
