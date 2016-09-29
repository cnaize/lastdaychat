package chat

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	mainTemplate  = template.Must(template.ParseFiles("templates/main.html.tpl"))
	roomsTemplate = template.Must(template.ParseFiles("templates/room.html.tpl"))
)

type Chat struct {
	router *mux.Router
	rooms  *RoomList
}

func NewChat() *Chat {
	return &Chat{rooms: NewRoomList()}
}

func (c *Chat) Remove(room *Room) {
	room.Close()
	c.rooms.Remove(room)
}

func (c *Chat) Run(defaultPort string) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	fmt.Printf("Chat run on port: %s...\n", port)
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/", c.mainShow)
	r.HandleFunc("/rooms", c.roomsCreate).Methods(http.MethodPost)
	r.HandleFunc("/rooms", c.roomsShow).Methods(http.MethodGet)
	r.HandleFunc("/users", c.userCount).Methods(http.MethodGet)
	r.HandleFunc("/ws", c.wsHandler)
	cors := handlers.CORS(handlers.AllowedOrigins([]string{"http://lastday.herokuapp.com"}), handlers.AllowedMethods([]string{http.MethodGet}))
	return http.ListenAndServe(":"+port, cors(r))
}

func (c *Chat) mainShow(w http.ResponseWriter, r *http.Request) {
	mainTemplate.Execute(w, nil)
}

func (c *Chat) roomsCreate(w http.ResponseWriter, r *http.Request) {
	room := NewRoom(c)
	c.rooms.Add(room)

	u := url.URL{Path: "/rooms", RawQuery: fmt.Sprintf("roomId=%s", room.id)}
	http.Redirect(w, r, u.String(), http.StatusFound)
}

func (c *Chat) roomsShow(w http.ResponseWriter, r *http.Request) {
	v := make(map[string]string)
	v["roomId"] = r.URL.Query().Get("roomId")
	v["host"] = r.Host
	roomsTemplate.Execute(w, v)
}

func (c *Chat) userCount(w http.ResponseWriter, r *http.Request) {
	room, err := c.rooms.Get(r.URL.Query().Get("roomId"))
	if err != nil {
		fmt.Fprint(w, "not found")
		return
	}
	room.users.RLock()
	defer room.users.RUnlock()
	fmt.Fprint(w, len(room.users.Users))
}

func (c *Chat) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	room, err := c.rooms.Get(r.URL.Query().Get("roomId"))
	if err != nil {
		conn.Close()
		return
	}
	room.add <- conn
}
