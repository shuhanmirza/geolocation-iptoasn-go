package infrastructure

type KeyValueStore struct {
	keyValueMap map[string]string
}

func NewKeyValueStore() KeyValueStore {
	return KeyValueStore{keyValueMap: make(map[string]string)}
}

func (kv KeyValueStore) Get(key string) (value string) {
	return kv.keyValueMap[key]
}

func (kv KeyValueStore) Set(key string, value string) {
	kv.keyValueMap[key] = value
}
