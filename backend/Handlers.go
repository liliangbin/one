package main

import (
	"net/http"
	"fmt"
	"./handlers"
	"encoding/json"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!  this is my first go model ")
}

func allowCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, If-Modified-Since") //header的类型
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

/*func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]d
	fmt.Fprintln(w, "Todo show:", todoId)
}
*/
/*func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	//解析出错
	if err :=json.Unmarshal(body,&todo);err!=nil{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err:=json.NewEncoder(w).Encode(err);err!=nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err:=json.NewEncoder(w).Encode(todo);err!=nil{
		panic(err)
	}
}
*/

func TheMax(w http.ResponseWriter, r *http.Request) {
	allowCORS(w)
	w.WriteHeader(http.StatusOK)
	one := handlers.Max()
	if err := json.NewEncoder(w).Encode(one); err != nil {
		panic(err)
	}
}
func FindById(w http.ResponseWriter, r *http.Request) {
	allowCORS(w)
	w.WriteHeader(http.StatusOK)
	r.ParseForm()
	fullId := r.Form["full_id"][0]
	fmt.Println(fullId)
	id, err := strconv.Atoi(fullId)
	if err != nil {
		panic(nil)
	}
	one := handlers.FindId(id)

	if err := json.NewEncoder(w).Encode(one); err != nil {
		panic(err)
	}

}
