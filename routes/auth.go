package routes

import (
	c "github.com/deltamc/otus-social-networks-users/controllers"
	m "github.com/deltamc/otus-social-networks-users/middlewares"
	"net/http"
)

func Auth() {
	http.HandleFunc("/refresh", m.Cors(m.Post(m.Jwt(c.HandleRefresh))))
	http.HandleFunc("/getUserByToken", m.Cors(m.Get(m.Jwt(c.HandleMy))))
	http.HandleFunc("/me", m.Cors(m.Get(m.Jwt(c.HandleMy))))
	http.HandleFunc("/profile", m.Cors(m.Post(m.Jwt(c.HandleProfile))))
	http.HandleFunc("/logout", m.Cors(m.Post(m.Jwt(c.HandleLogout))))
}
