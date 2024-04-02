package repository

type Interface interface {
	Init() error
	Add(interface{}) (interface{}, error)
	Get(interface{}) (interface{}, error)
	Update(interface{}) (interface{}, error)
	Delete(interface{}) error
}
