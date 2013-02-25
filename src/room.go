package webrtcing

import (
  "appengine"
  "appengine/datastore"
)

type Room struct {
  User1 string
  User2 string
  // AppEngine Channels API connected properly
  AEConnected1 bool
  AEConnected2 bool
  // JavaScript Channels API connected properly
  JSConnected1 bool
  JSConnected2 bool
}

func (r *Room) OtherUser(user string) string {
  if user == r.User2 {
    return r.User1
  }
  if user == r.User1 {
    return r.User2
  }
  return ""
}

func (r *Room) HasUser(user string) bool {
  if user == r.User2 {
    return true
  }
  if user == r.User1 {
    return true
  }
  return false
}

func (r *Room) AddUser(user string) {
  if r.User1 == "" && r.User2 != user {
    r.User1 = user
  } else if r.User2 == "" && r.User1 != user {
    r.User2 = user
  }
}

func (r *Room) RemoveUser(user string) bool {
  if user == r.User2 {
    r.User2 = ""
    r.AEConnected2 = false
    r.JSConnected2 = false
  }
  if user == r.User1 {
    r.User1 = ""
    r.AEConnected1 = false
    r.JSConnected1 = false
  }
  // returns true if it should be deleted
  return r.Occupants() == 0
}

func (r *Room) AEConnectUser(user string) {
  if user == r.User1 {
    r.AEConnected1 = true
  }
  if user == r.User2 {
    r.AEConnected2 = true
  }
}

func (r *Room) JSConnectUser(user string) {
  if user == r.User1 {
    r.JSConnected1 = true
  }
  if user == r.User2 {
    r.JSConnected2 = true
  }
}

func (r *Room) Connected(user string) bool {
  if user == r.User1 && r.AEConnected1 && r.JSConnected1 {
    return true
  }
  if user == r.User2 && r.AEConnected2 && r.JSConnected2 {
    return true
  }
  return false
}

func (r *Room) Occupants() int {
  occupancy := 0
  if r.User1 != "" { occupancy += 1 }
  if r.User2 != "" { occupancy += 1 }
  return occupancy
}

func GetRoom(c appengine.Context, name string) (*Room, error) {
  k := datastore.NewKey(c, "Room", name, 0, nil)
  r := new(Room)
  err := datastore.Get(c, k, r)
  return r, err;
}

func PutRoom(c appengine.Context, name string, room *Room) error {
  k := datastore.NewKey(c, "Room", name, 0, nil)
  _, err := datastore.Put(c, k, room)
  return err;
}

func DelRoom(c appengine.Context, name string) error {
  k := datastore.NewKey(c, "Room", name, 0, nil)
  err := datastore.Delete(c, k)
  return err;
}
