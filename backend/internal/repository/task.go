package repository

//go:generate mockgen -source=${GOFILE} -destination=task_mock_test.go -package=repository

type Task interface {
	GetAll()
	Create()
	Get()
	Update()
}
