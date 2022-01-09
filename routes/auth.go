package routes

import (
	c "github.com/deltamc/otus-social-networks-chat/controllers"
	m "github.com/deltamc/otus-social-networks-chat/middlewares"
	"net/http"
)

func Auth() {
	http.HandleFunc("/refresh", m.Cors(m.Post(m.Jwt(c.HandleRefresh))))
	http.HandleFunc("/me", m.Cors(m.Get(m.Jwt(c.HandleMy))))
	//http.HandleFunc("/friends", m.Cors(m.Get(m.Jwt(c.HandleFriends))))
	//http.HandleFunc("/make_friend", m.Cors(m.Post(m.Jwt(c.HandleMakeFriend))))
	http.HandleFunc("/profile", m.Cors(m.Post(m.Jwt(c.HandleProfile))))
	http.HandleFunc("/logout", m.Cors(m.Post(m.Jwt(c.HandleLogout))))


	http.HandleFunc("/messages/get", m.Cors(m.Get(m.Jwt(c.HandleMessagesGet))))
	http.HandleFunc("/messages/post", m.Cors(m.Post(m.Jwt(c.HandleMessagesPost))))

}
