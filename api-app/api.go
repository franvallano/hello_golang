package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var collection = getSession().DB("go-app").C("customer")


func CustomerInfo(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	customer_id := params["id"]	

	if !bson.IsObjectIdHex(customer_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(customer_id)

	result := Customer{}
	err := collection.FindId(oid).One(&result)

	if err != nil {
		w.WriteHeader(404)
		log.Fatal(err)
		return
	}else{
		ResponseOkUnique(201, w, result)
	}
}

func CustomerList(w http.ResponseWriter, r *http.Request) {
	var results []Customer
	err := collection.Find(nil).Sort("-_id").All(&results)
	
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
		return
	}else{
		ResponseOkList(201, w, results)
	}

}

func CustomerAdd(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body) 
	
	var new_customer Customer
	err := decoder.Decode(&new_customer)
	defer r.Body.Close()
	
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
		return
	}else{
		collection.Insert(new_customer)
		ResponseOkUnique(201, w, new_customer)
	}
}


func CustomerUpdate(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	customer_id := params["id"]	

	if !bson.IsObjectIdHex(customer_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(customer_id)

	decoder := json.NewDecoder(r.Body) 
	var customer_data Customer
	err := decoder.Decode(&customer_data)
	defer r.Body.Close()
	
	if err != nil {
		w.WriteHeader(500)
		log.Fatal(err)
		return
	}

	document := bson.M{"_id": oid}
	change:= bson.M{"$set": customer_data}

	err = collection.Update(document, change)
	if err != nil {
		w.WriteHeader(404)
		log.Fatal(err)
		return
	}

	ResponseOkUnique(201, w, customer_data)
}

func CustomerDelete(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	customer_id := params["id"]	

	if !bson.IsObjectIdHex(customer_id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(customer_id)

	err := collection.RemoveId(oid)

	if err != nil {
		w.WriteHeader(404)
		log.Fatal(err)
		return
	}else{
		response := new(Response)
		response.SetStatus("success")
		response.SetMessage("The customer has been deleted")

		ResponseOkMessage(201, w, *response)
	}
}

func ResponseOkUnique(status int, w http.ResponseWriter, result Customer ){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

func ResponseOkMessage(status int, w http.ResponseWriter, msg Response){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(msg)
}

func ResponseOkList(status int, w http.ResponseWriter, results []Customer ){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}