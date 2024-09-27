package services

//go:generate mockgen -source=${GOFILE} -destination=task_mock_test.go -package=services

type TaskInterface interface {
	GetAll()
	Create()
	Get()
	Update()
}

type service struct{}

func NewTask() TaskInterface {
	return &service{}
}

func (s *service) GetAll() {}

func (s *service) Create() {}

func (s *service) Get() {}

func (s *service) Update() {}
