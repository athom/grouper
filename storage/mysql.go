package storage

func NewMysqlStorage() (r *MysqlStorage, err error) {
	r = &MysqlStorage{}
	return
}

type MysqlStorage struct {
}

func (this *MysqlStorage) CreateConnection(id1 string, id2 string) (err error) {
	return nil
}
