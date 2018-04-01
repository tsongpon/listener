package handler

import (
	"encoding/json"
	"fmt"
	"github.com/tsongpon/listener/config"
	"github.com/tsongpon/listener/data"
	"github.com/tsongpon/listener/transport"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func FacebookHookGet(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("hub.verify_token")
	if token == config.GetVerifyToken() {
		io.WriteString(w, r.FormValue("hub.challenge"))
	} else {
		log.Println("token verify fail")
		w.WriteHeader(400)
	}

}

func FacebookHookPost(w http.ResponseWriter, r *http.Request) {
	var changeTransport transport.Transport
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &changeTransport); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	entry := changeTransport.Entry[0]
	changes := entry.Change
	for _, each := range changes {
		var model = data.UserChange{}
		model.UserId = entry.Uid
		model.Time = time.Unix(entry.Time, 0)
		model.Field = each.Field
		model.Value = each.Value
		log.Println("Create user activity log for userId: ", entry.Uid)
		data.Dao.Save(model)
	}
	w.WriteHeader(http.StatusOK)
}

func QueryUserActivities(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userid")
	field := r.FormValue("field")
	sizeStr := r.FormValue("size")
	startStr := r.FormValue("start")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 5
	}
	start, err := strconv.Atoi(startStr)
	if err != nil {
		start = 0
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	activities := data.Dao.QueryActivities(data.Query{UserId: userId, Value: field, Size: size, Start: start})
	if activities == nil {
		activities = []data.UserChange{}
	}
	total := data.Dao.CountActivities(data.Query{UserId: userId, Value: field})
	w.WriteHeader(http.StatusOK)
	response := transport.UserActivities{Total: total, Size: len(activities), Data: activities}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!, but nothing here.")
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Index request", r.RemoteAddr, r.URL)
	t, _ := template.ParseFiles("templates/login.html")
	t.Execute(w, nil)
}
