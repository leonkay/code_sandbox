package main

import (
	"github.com/leonkay/code_sandbox/golang/websocket/config"
	templates "github.com/leonkay/code_sandbox/golang/websocket/templ"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:		1024,
	WriteBufferSize:	1024,

	// We'll need to check the origin of our connection
  // this will allow us to make requests from our React
  // development server to here.
  // For now, we'll do no checking and just allow any connection
  CheckOrigin: func(r *http.Request) bool { return true },
}

type SocketData struct {
	Message string
}

var clients []websocket.Conn

func reader(conn *websocket.Conn) {
	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("%s error reading: %s\n", conn.RemoteAddr(), err)
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			fmt.Printf("%s error writing : %s\n", conn.RemoteAddr(), err)
			return
		}
	}
}

func htmxBody(message string)string {
		template := `
		<div hx-swap-oob="beforeend:#output">%s<br /></div>
		`
		s := fmt.Sprintf(template, message, message)
		return s
}
func chatReader(conn *websocket.Conn) {
	for {

		msg := SocketData{};
		err := conn.ReadJSON(&msg)

		if err != nil {
			fmt.Printf("%s error reading: %s\n", conn.RemoteAddr(), err)
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), msg)

		// Write message back to browser
		if err = conn.WriteMessage(websocket.TextMessage, []byte(htmxBody(msg.Message))); err != nil {
			fmt.Printf("%s error writing : %s\n", conn.RemoteAddr(), err)
			return
		}
	}
}

func echoWs(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s connected\n", r.Host)

	// upgrade this connection to a WebSocket
  // connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
			log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func chatWs(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s connected\n", r.Host)

	// upgrade this connection to a WebSocket
  // connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
			log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	chatReader(ws)
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()
	fmt.Println(conf.DebugMode)

	http.HandleFunc("/echo", echoWs)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		tar := templates.Hello("test")
		tar.Render(r.Context(), w)
	})

	http.HandleFunc("/wsChat", chatWs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "sample.html")
	})

	http.HandleFunc("/chat",  func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":8080", nil)
}

