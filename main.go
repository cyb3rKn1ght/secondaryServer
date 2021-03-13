package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	registration "servers/handlers/login"
	msg "servers/handlers/message"
	"servers/redis"
)

type secServState struct {
	State string
}

func main() {

	http.HandleFunc("/message", msg.HandleMessage)
	http.HandleFunc("/login", registration.Login)
	http.HandleFunc("/logout", registration.Logout)

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
	}

	key, err := redis.Redis().Get("SecondaryServer:" + fmt.Sprint(addrs[1])).Result()

	if err != nil {
		if fmt.Sprint(err) == "redis: nil" {

			key = `{"State":"online"}`

		} else {
			fmt.Println(err)
		}
	}

	var scSrvSt secServState

	err = json.Unmarshal([]byte(key), &scSrvSt)
	if err != nil {
		fmt.Println(err)
	}

	scSrvSt.State = "online"

	stringKey, err := json.Marshal(scSrvSt)

	if err != nil {
		fmt.Println(err)
	}

	err = redis.Redis().Set("SecondaryServer:"+fmt.Sprint(addrs[1]), stringKey, 0).Err()

	if err != nil {
		fmt.Println(err)
	}

	http.ListenAndServe(":8080", nil)

}
