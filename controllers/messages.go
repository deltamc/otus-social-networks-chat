package controllers

import (
	"github.com/deltamc/otus-social-networks-chat/models/messages"
	"github.com/deltamc/otus-social-networks-chat/models/users"
	"github.com/deltamc/otus-social-networks-chat/responses"
	"net/http"
	"strconv"
)

func HandleMessagesGet(w http.ResponseWriter, r *http.Request, user users.User) {

	m, err := messages.GetMessages(user)

	if err != nil {
		responses.Response500(w, err)
		return
	}

	responses.ResponseJson(w, m)
}


func HandleMessagesPost(w http.ResponseWriter, r *http.Request, user users.User) {

	_= r.ParseForm()

	userTo,_ := strconv.ParseInt(r.FormValue("user_to"),10, 64)

	mes := r.FormValue("message")

	message := messages.Message{
		UserIdTo:   userTo,
		Message:    mes,
	}

	_, err := message.New(user)

	if err != nil {
		responses.Response500(w, err)
		return
	}

	responses.ResponseJson(w, responses.Response200)
}





