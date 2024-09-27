package controller

type controller struct{}

func New() ServerInterface {
	return &controller{}
}
