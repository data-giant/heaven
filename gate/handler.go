package gate

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"heaven/storage/mysql"
	"time"
	"strconv"
)

func slient(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	
}

var st = mysql.NewStorage()

func Miner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	record := map[string]string{
		"app_id": ps.ByName("app_id"),
		"client_ip": ps.ByName("client_ip"),
		"client_time": ps.ByName("client_time"),
		"server_time": strconv.FormatInt(time.Now().Unix(), 10),
		"sdk_version": ps.ByName("sdk_version"),
		"resolution": ps.ByName("resolution"),
		"os": ps.ByName("os"),
		"event_id": ps.ByName("id"),
		"page": ps.ByName("page"),
		"arg1": ps.ByName("arg1"),
		"arg2": ps.ByName("arg2"),
		"arg3": ps.ByName("arg3"),
		"args": ps.ByName("args"),
	}
	st.Write(record)
}
