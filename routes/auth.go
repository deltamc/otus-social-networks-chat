package routes

import (
	c "github.com/deltamc/otus-social-networks-chat/controllers"
	m "github.com/deltamc/otus-social-networks-chat/middlewares"
	"net/http"
)

func Auth() {
	http.HandleFunc("/messages/get", m.Cors(m.Get(m.Jwt(c.HandleMessagesGet))))
	http.HandleFunc("/messages/post", m.Cors(m.Post(m.Jwt(c.HandleMessagesPost))))
}
