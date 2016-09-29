package chat

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

type User struct {
	conn *websocket.Conn
	room *Room
	in   chan string
	out  chan string
}

func NewUser(conn *websocket.Conn, room *Room) *User {
	user := &User{conn: conn, room: room, in: make(chan string), out: make(chan string)}
	go user.Listen()
	return user
}

func (u *User) Listen() {
	go u.Read()
	go u.Write()
}

func (u *User) Read() {
	defer func() {
		u.room.del <- u
	}()

	u.conn.SetReadDeadline(time.Now().Add(pongWait))
	u.conn.SetPongHandler(func(string) error { u.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := u.conn.ReadMessage()
		if err != nil {
			break
		}
		u.in <- string(msg)
	}
}

func (u *User) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case msg, ok := <-u.out:
			if !ok {
				return
			}
			w, err := u.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write([]byte(msg))

			n := len(u.out)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write([]byte(<-u.out))
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := u.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (u *User) Close() {
	u.conn.Close()
	close(u.in)
	close(u.out)
}

type UserList struct {
	sync.RWMutex
	Users map[*User]struct{}
}

func NewUserList() *UserList {
	return &UserList{Users: make(map[*User]struct{})}
}

func (l *UserList) Add(user *User) {
	l.Lock()
	defer l.Unlock()
	l.Users[user] = struct{}{}
}

func (l *UserList) Remove(user *User) {
	l.Lock()
	defer l.Unlock()
	delete(l.Users, user)
}
