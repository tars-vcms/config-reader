package logic

type ConfigReader interface {
	GetConfig()
	PutConfig()
}