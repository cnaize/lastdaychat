package chat

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	id    string
	chat  *Chat
	users *UserList
	msgs  chan string
	add   chan *websocket.Conn
	del   chan *User
}

func NewRoom(chat *Chat) *Room {
	r := &Room{
		id:    GenerateRoomId(),
		chat:  chat,
		users: NewUserList(),
		msgs:  make(chan string),
		add:   make(chan *websocket.Conn),
		del:   make(chan *User),
	}
	go r.Listen()
	return r
}

func (r *Room) Listen() {
	for {
		select {
		case msg, ok := <-r.msgs:
			if !ok {
				return
			}
			r.Broadcast(msg)
		case conn, ok := <-r.add:
			if !ok {
				return
			}
			r.Add(conn)
		case user, ok := <-r.del:
			if !ok {
				return
			}
			r.Remove(user)
		}
	}
}

func (r *Room) Broadcast(msg string) {
	r.users.RLock()
	defer r.users.RUnlock()

	for u, _ := range r.users.Users {
		u.out <- msg
	}
}

func (r *Room) Add(conn *websocket.Conn) {
	user := NewUser(conn, r)
	r.users.Add(user)
	go func() {
		for {
			msg, ok := <-user.in
			if !ok {
				break
			}
			r.msgs <- msg
		}
	}()
}

func (r *Room) Remove(user *User) {
	user.Close()
	r.users.Remove(user)

	if len(r.users.Users) == 0 {
		r.chat.Remove(r)
	}
}

func (r *Room) Close() {
	r.users.RLock()
	defer r.users.RUnlock()

	for u, _ := range r.users.Users {
		r.del <- u
	}
	close(r.msgs)
	close(r.add)
	close(r.del)
}

type RoomList struct {
	sync.RWMutex
	Rooms map[string]*Room
}

func NewRoomList() *RoomList {
	return &RoomList{Rooms: make(map[string]*Room)}
}

func (l *RoomList) Get(roomId string) (*Room, error) {
	l.RLock()
	defer l.RUnlock()
	room, ok := l.Rooms[roomId]
	if !ok {
		return nil, fmt.Errorf("Room %q not found", roomId)
	}
	return room, nil
}

func (l *RoomList) Add(room *Room) {
	l.Lock()
	defer l.Unlock()
	l.Rooms[room.id] = room
}

func (l *RoomList) Remove(room *Room) {
	l.Lock()
	defer l.Unlock()
	delete(l.Rooms, room.id)
}
