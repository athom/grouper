package storage

func NewMemoryStorage() (r *MemoryStorage) {
	r = &MemoryStorage{
		relationshipMap: make(map[string][]string),
	}
	return
}

type MemoryStorage struct {
	relationshipMap map[string][]string
}

func (this *MemoryStorage) CreateConnection(id1 string, id2 string) (err error) {
	if v, ok := this.relationshipMap[id1]; ok {
		v = append(v, id2)
		return
	}
	this.relationshipMap[id1] = []string{id2}

	if v, ok := this.relationshipMap[id2]; !ok {
		v = append(v, id1)
		return
	}
	this.relationshipMap[id2] = []string{id1}
	return nil
}

func (this *MemoryStorage) CreateSubscription(id1 string, id2 string) (err error) {
	return
}

func (this *MemoryStorage) BlockConnection(id1 string, id2 string) (err error) {
	return
}

func (this *MemoryStorage) ShowConnections(id string) (r []string, err error) {
	if v, ok := this.relationshipMap[id]; ok {
		r = v
		return
	}
	return
}

func (this *MemoryStorage) CommonConnections(id1 string, id2 string) (r []string, err error) {
	return
}

func (this *MemoryStorage) GetReachableConnections(id string) (r []string, err error) {
	return
}
