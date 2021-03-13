package login

import (
	"fmt"
	"net/http"
	"servers/redis"
)

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	userName, nameOk := r.URL.Query()["name"]

	if nameOk {
		err := redis.Redis().Set(userName[0]+r.RemoteAddr, "online", 0).Err()

		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Logout ...
func Logout(w http.ResponseWriter, r *http.Request) {
	userName, nameOk := r.URL.Query()["name"]

	if nameOk {
		err := redis.Redis().Set(userName[0]+r.RemoteAddr, "offline", 0).Err()

		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
