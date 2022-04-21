package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/stretchr/objx"
)

type ChatUser interface {
	UniqueID() string
	AvatarURL() string
}

type chatUser struct {
	goth.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}

func (u chatUser) AvatarURL() string {
	return u.User.AvatarURL
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.next.ServeHTTP(w, r)
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		queryParams := r.URL.Query()
		queryParams.Add("provider", provider)
		r.URL.RawQuery = queryParams.Encode()

		if _, err := gothic.CompleteUserAuth(w, r); err == nil {
			w.Header().Set("Location", "/chat")
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}

		fmt.Println("New login")
		gothic.BeginAuthHandler(w, r)
	case "callback":
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			// http.Error(w, fmt.Sprintf("Error when trying to get user from %s: %s", provider, err), http.StatusInternalServerError)
			// return
			log.Fatalln("Error when trying to get user from", provider, "-", err)
		}

		chatUser := &chatUser{User: user}

		m := md5.New()
		io.WriteString(m, strings.ToLower(user.Email))
		// userId := fmt.Sprintf("%x", m.Sum(nil))
		chatUser.uniqueID = fmt.Sprintf("%x", m.Sum(nil))
		avatarURL, err := avatars.GetAvatarURL(chatUser)
		if err != nil {
			log.Fatalln("Error when trying to GetAvatarURL", "-", err)
		}

		authCookieValue := objx.New(map[string]any{
			"userid":     chatUser.uniqueID,
			"name":       user.Name,
			"avatar_url": avatarURL,
			// "email":      user.Email,
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})

		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
