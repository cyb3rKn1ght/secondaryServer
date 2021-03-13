package message

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"servers/redis"
	"time"
)

type server struct {
	State   string
	History []string
}

// HandleMessage ...
func HandleMessage(w http.ResponseWriter, r *http.Request) {
	userName, nameOk := r.URL.Query()["name"]
	msg, msgOk := r.URL.Query()["message"]

	if nameOk && msgOk {
		fmt.Println(userName[0] + ": " + msg[0])

		addrs, err := net.InterfaceAddrs()

		if err != nil {
			fmt.Println(err)
		}

		key, err := redis.Redis().Get("SecondaryServer:" + fmt.Sprint(addrs[1])).Result()

		if err != nil {
			fmt.Println(err)
		}

		var server server

		err = json.Unmarshal([]byte(key), &server)
		if err != nil {
			fmt.Println(err)
		}

		date := time.Now().UTC().Format(time.RFC3339)

		server.History = append(server.History, date+" "+userName[0]+": "+msg[0])

		fmt.Println(server)

		serverToString, err := json.Marshal(server)

		if err != nil {
			fmt.Println(err)
		}

		err = redis.Redis().Set("SecondaryServer:"+fmt.Sprint(addrs[1]), serverToString, 0).Err()

	}

}
