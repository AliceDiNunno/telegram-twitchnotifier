package usecases

type Usecases interface {
	RegisterForEvents()
	StartWatcher()
}
