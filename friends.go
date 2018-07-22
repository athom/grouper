package grouper

import "github.com/athom/goset"

type FriendId string

func (this *FriendId) Validate() (err error) {
	return
}

type FriendShip interface {
	MakeFriend(id1 FriendId, id2 FriendId) error
	ListFriends(FriendId) ([]FriendId, error)
	CommonFriends(id1 FriendId, id2 FriendId) error
	Subscribe(fromId FriendId, toId FriendId) error
	Block(fromId FriendId, toId FriendId) error
	ReachableIds(id FriendId) ([]FriendId, error)
}

type Storage interface {
	CreateConnection(id1 string, id2 string) error
	ShowConnections(id string) ([]string, error)
}

func NewGrouper(storage Storage) (r *Grouper, err error) {
	r = &Grouper{storage: storage}
	return
}

type Grouper struct {
	storage Storage
}

func (this *Grouper) MakeFriend(id1 FriendId, id2 FriendId) (err error) {
	if err = id1.Validate(); err != nil {
		return err
	}
	if err = id2.Validate(); err != nil {
		return err
	}

	err = this.storage.CreateConnection(string(id1), string(id2))
	return err
}

func (this *Grouper) ListFriends(id1 FriendId) (r []FriendId, err error) {
	if err = id1.Validate(); err != nil {
		return
	}

	var ids []string
	ids, err = this.storage.ShowConnections(string(id1))

	r = goset.Map(ids, func(id string) FriendId {
		return FriendId(id)
	}, []FriendId{}).([]FriendId)

	return
}
