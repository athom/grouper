package storage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	TableName = "relationship"
)

const (
	RelationShipStatuDoubleConnected = "double_connected"
	RelationShipStatuSubscribed      = "subscribed"
	RelationShipStatuBlocked         = "blocked"
)

type RelationShip struct {
	gorm.Model

	FollowerId string `sql:"not null" gorm:"type:varchar(100);column:follower_id"`
	FolloweeId string `sql:"not null" gorm:"type:varchar(100);column:followee_id"`
	Status     string `sql:"not null" gorm:"type:varchar(20);index;column:status"`
}

func (this RelationShip) TableName() string {
	return TableName
}

func NewMysqlStorage(host string, port int, userName string, password string, database string) (r *MysqlStorage) {
	str := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		userName,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open("mysql", str)
	if err != nil {
		panic(err)
	}

	if !db.HasTable(TableName) {
		db = db.CreateTable(&RelationShip{}).AddUniqueIndex(
			"relastion_ship",
			"follower_id", "followee_id",
		)
		errs := db.GetErrors()
		if len(errs) > 0 {
			panic(errs[0])
		}
	}

	r = &MysqlStorage{db}
	return
}

type MysqlStorage struct {
	db *gorm.DB
}

func (this *MysqlStorage) CreateConnection(id1 string, id2 string) (err error) {
	obj1 := &RelationShip{
		FollowerId: id1,
		FolloweeId: id2,
		Status:     RelationShipStatuDoubleConnected,
	}
	obj2 := &RelationShip{
		FollowerId: id2,
		FolloweeId: id1,
		Status:     RelationShipStatuDoubleConnected,
	}
	this.db.Create(obj1).Create(obj2)
	errs := this.db.GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func (this *MysqlStorage) ShowConnections(id string) (r []string, err error) {
	this.db.Where(&RelationShip{
		FollowerId: id,
	})

	var friends []*RelationShip
	this.db.Select("followee_id").Where(&RelationShip{
		FollowerId: id,
		Status:     RelationShipStatuDoubleConnected,
	}).Find(&friends)

	errs := this.db.GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	for _, u := range friends {
		r = append(r, u.FolloweeId)
	}

	return
}
