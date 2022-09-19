package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/deltamc/otus-social-networks-chat/models/users"
	"github.com/deltamc/otus-social-networks-chat/responses"
	"io/ioutil"
	"net/http"
	"os"
)

func Jwt(h handlerAuth) handler {

	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		usersUrl := fmt.Sprintf("%s/getUserByToken", os.Getenv("USER_SERVICE"))

		req, err := http.NewRequest("GET", usersUrl, bytes.NewBuffer([]byte("")))

		if err != nil {
			responses.Response500(w, err)
			return
		}

		req.Header.Set("Authorization", authHeader)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			responses.Response500(w, err)
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(resp.StatusCode)
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)
		var user users.User
		json.Unmarshal(body, &user)

		h(w, r, user)
	}
}
