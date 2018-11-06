package http

import (
	l "log"
	"net/http"

	"github.com/chatrealtime/websoc"

	"github.com/chatrealtime/endpoints"
	messageDecode "github.com/chatrealtime/http/decode/message"
	roomDecode "github.com/chatrealtime/http/decode/room"
	userDecode "github.com/chatrealtime/http/decode/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints endpoints.Endpoints, websoc *websoc.WebSoc, logger log.Logger) http.Handler {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Use(cors.Handler)
	r.Get("/", Index)
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", httptransport.NewServer(
			endpoints.RegisterUser,
			userDecode.RegisterRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/login", httptransport.NewServer(
			endpoints.LoginUser,
			userDecode.LoginRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})
	r.Route("/rooms", func(r chi.Router) {
		r.Post("/create", httptransport.NewServer(
			endpoints.CreateRoom,
			roomDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/join", httptransport.NewServer(
			endpoints.FindRoom,
			roomDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/get-all", httptransport.NewServer(
			endpoints.GetAllRoom,
			roomDecode.GetAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})
	r.Route("/messages", func(r chi.Router) {
		r.Post("/all", httptransport.NewServer(
			endpoints.GetAllMessageRoom,
			messageDecode.GetAllMessageRoomRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})
	r.Get("/messages", Message)
	r.Route("/chats", func(r chi.Router) {
		r.Post("/", httptransport.NewServer(
			endpoints.CreateMessage,
			messageDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	go ListenMessage(r, websoc)

	return r
}

func Message(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "UI/chat/chat.html")
}

func Index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "UI/login/index.html")
}

func ListenMessage(r *chi.Mux, websoc *websoc.WebSoc) {
	r.Get("/chats", websoc.HandlerChat)
	for {
		select {
		case msg := <-websoc.NewInComingMessage:
			websoc.SendMessage(msg)
		case user := <-websoc.DisconnectUser:
			l.Println("disconnect user ", user.User.Name)
			delete(websoc.ConnectedUser, user.User.Name)
		case err := <-websoc.ErrChannel:
			l.Println("error channel websocket by error ", err.Error())
		}
	}

}
