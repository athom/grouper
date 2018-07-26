package storage

import (
	"errors"
	"fmt"

	"github.com/athom/goset"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Data model spec
// (follower, followee, is_connected, is_subscribed, is_blocked)
//
// - A B are disconnected:
// 	 case1: really no relationship, empty set
//	 case2: B blocks A (A, B, 0, 0, 1)
//	 case3: A blocks B (B, A, 0, 0, 1)
//   case4: A subscribe B (A, B, 0, 1, 0)
//   case5: B subscribe A (B, A, 0, 1, 0)
//	 case*: {case2, case3, case4, case5}*
//
// - A B are connected:
// (A, B, 1, *, *) and (B, A, 1, *, *)
//
// - A subscribe B
// (A, B, *, 1, *)
//
//  - A block B
//	(B, A, *, *, 1)
//
// - connectable:
//	 case1: empty set
//	 case2: (A, B, 0, *, 0)
//	 case3: (B, A, 0, *, 0)
//	 case4: case2 and case3
//
// - A can receive B's update
//	 case1: (A, B, 1, *, 0)
//	 case2: (A, B, 0, 1, 0)

const (
	TableName = "relationship"
)

const (
	RelationShipStatusConnected  = "connected"
	RelationShipStatusSubscribed = "subscribed"
	RelationShipStatusBlocked    = "blocked"
)

type RelationShip struct {
	gorm.Model

	FollowerId string `sql:"not null" gorm:"type:varchar(100);column:follower_id"`
	FolloweeId string `sql:"not null" gorm:"type:varchar(100);column:followee_id"`
	Connected  string `sql:"not null" gorm:"type:varchar(20);default:'';column:connected"`
	Subscribed string `sql:"not null" gorm:"type:varchar(20);default:'';column:subscribed"`
	Blocked    string `sql:"not null" gorm:"type:varchar(20);default:'';column:blocked"`
}

func (this *RelationShip) SetConected() {
	this.Blocked = RelationShipStatusConnected
}

func (this *RelationShip) SetBlocked() {
	this.Blocked = RelationShipStatusBlocked
}

func (this *RelationShip) SetSubscribed() {
	this.Blocked = RelationShipStatusSubscribed
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

// NOTE very dangerous, only for testing purpose
func (this *MysqlStorage) ResetData() (err error) {
	this.db.Unscoped().Delete(&RelationShip{})
	errs := this.db.GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return
}

func (this *MysqlStorage) CreateConnection(id1 string, id2 string) (err error) {
	var rss []*RelationShip
	this.db.Where("( followee_id = ? or followee_id = ?) AND  blocked = ?",
		id1,
		id2,
		RelationShipStatusBlocked,
	).Find(&rss)
	if len(rss) > 0 {
		err = errors.New("relationship already blocked, can not connect")
		return
	}

	// TODO make it in a transation for the sake of consistency
	obj1 := &RelationShip{
		FollowerId: id1,
		FolloweeId: id2,
		Connected:  RelationShipStatusConnected,
	}
	obj2 := &RelationShip{
		FollowerId: id2,
		FolloweeId: id1,
		Connected:  RelationShipStatusConnected,
	}
	this.db.Create(obj1).Create(obj2)
	errs := this.db.GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	return
}

func (this *MysqlStorage) CreateSubscription(id1 string, id2 string) (err error) {
	obj1 := &RelationShip{
		FollowerId: id1,
		FolloweeId: id2,
		Connected:  RelationShipStatusSubscribed,
	}

	this.db.Create(obj1)
	errs := this.db.GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	return
}

func (this *MysqlStorage) BlockConnection(id1 string, id2 string) (err error) {
	var rs RelationShip
	// NOTE pay attention to the id order
	this.db.FirstOrCreate(&rs, RelationShip{
		FollowerId: id2,
		FolloweeId: id1,
	})

	rs.SetBlocked()
	this.db.Save(rs)
	errs := this.db.GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}
	return
}

func (this *MysqlStorage) ShowConnections(id string) (r []string, err error) {
	var friends []*RelationShip
	this.db.Select("followee_id").Where(&RelationShip{
		FollowerId: id,
		Connected:  RelationShipStatusConnected,
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

func (this *MysqlStorage) CommonConnections(id1 string, id2 string) (r []string, err error) {
	// TODO optimized interset via sql
	var friends, friends1, friends2 []*RelationShip
	this.db.Select("followee_id").Where(&RelationShip{
		FollowerId: id1,
		Connected:  RelationShipStatusConnected,
	}).Find(&friends1)
	errs := this.db.GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	this.db.Select("followee_id").Where(&RelationShip{
		FollowerId: id2,
		Connected:  RelationShipStatusConnected,
	}).Find(&friends2)
	errs = this.db.GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	friends = goset.Intersect(friends1, friends2).([]*RelationShip)
	for _, u := range friends {
		r = append(r, u.FolloweeId)
	}
	return
}

func (this *MysqlStorage) GetReachableConnections(id string) (r []string, err error) {
	var rss []*RelationShip
	this.db.Where("followee_id = ? AND  blocked != ?", id, RelationShipStatusBlocked).Find(&rss)
	errs := this.db.GetErrors()
	if len(errs) > 0 {
		err = errs[0]
		return
	}
	for _, u := range rss {
		r = append(r, u.FollowerId)
	}
	return
}
