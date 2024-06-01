package EventBus

type EventListener interface {
	Listen(request []byte) (map[string]interface{}, error)
}
