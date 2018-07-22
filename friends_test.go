package grouper

import (
	"testing"

	"github.com/athom/goset"
	"github.com/athom/grouper/storage"
)

var testStorage = storage.NewMemoryStorage()

func TestMakeFriend(t *testing.T) {
	g, err := NewGrouper(testStorage)
	if err != nil {
		t.Fail()
	}
	email1 := "alex@test.com"
	email2 := "bob@test.com"
	id1 := FriendId(email1)
	id2 := FriendId(email2)
	err = g.MakeFriend(id1, id2)
	if err != nil {
		t.Fail()
	}

	ids, err := testStorage.ShowConnections(email1)
	if err != nil {
		t.Fail()
	}
	if !goset.IsIncluded(ids, email2) {
		t.Errorf("%v should connected with %v but not got %v", id1, id2, ids)
	}

}
