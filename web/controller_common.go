package web

import (
	"github.com/athom/grouper"
	"github.com/athom/grouper/storage"
	"github.com/gin-gonic/gin"
)

var ok = gin.H{"success": true}

type Controller struct {
	core grouper.FriendShip
}

func (this *Controller) Setup(conf *config) {
	var s grouper.Storage
	storageType := conf.StorageType
	switch storageType {
	case storage.StorageTypeMysql:
		mysqlConf := conf.MySqlSettings
		s = storage.NewMysqlStorage(
			mysqlConf.Host,
			mysqlConf.Port,
			mysqlConf.UserName,
			mysqlConf.Password,
			mysqlConf.Database,
		)
	case storage.StorageTypeMemory:
		s = storage.NewMemoryStorage()
	default:
		s = storage.NewMemoryStorage()
	}

	this.core = grouper.NewGrouper(s)
}
