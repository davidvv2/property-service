package headers

type Header interface {
	Add(key string, value string)
	Remove(key string)
	List() map[string][]string
	Get(key string) string
}
