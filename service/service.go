package service

import (
	"encoding/hex"
	"encoding/json"
	"eventsbook/persistence"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)


type Service struct {
	data persistence.Database
}
func NewService(database persistence.Database) *Service {
	return &Service{
		data: database,
	}
}
func (s *Service) findEvent(w http.ResponseWriter,r *http.Request) {
     vars := mux.Vars(r)
	 criteria, ok := vars["searchCriteria"]
	 if !ok {
		w.WriteHeader(400) 
		fmt.Fprintf(w,`{error: cannot describe event you either search by name or id}`)
		return
	 }
	 searchKey, ok := vars["search"]
	 if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w,`{error: cannot describe event you either search by name or id }`)
	 }
	 var persistent persistence.Event
	 var err error
	 switch strings.ToLower(criteria) {
	 case "name":
		persistent, err = s.data.FindEventByName(criteria)
	 case "id":
		id, err := hex.DecodeString(searchKey)	
		if err == nil {
			persistent, err = s.data.FindEvent(id)
			if err != nil {
				fmt.Printf("not present %s", err)
			}
		}
	 }
	 if err != nil {
		fmt.Printf("error %s", err)
	 }
	 w.Header().Set("content-type","application/json;charset=UTF8")
	 encode :=json.NewEncoder(w)
	 err = encode.Encode(persistent)
	 if err != nil {
		fmt.Printf("cannot encode to json %s", err)
	 }

}
func (s *Service) addEvent(w http.ResponseWriter, r *http.Request) {
    e := persistence.Event{}
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&e)
	if err != nil {
		fmt.Fprintf(w,`{error: cannot decode from json} %s`,err)
		w.WriteHeader(500)
		return
	}
	id ,err := s.data.AddEvent(e)
	if err != nil {
		fmt.Fprintf(w,`{error: cannot add event because it does not exist} %s, %d`,err,id)
		w.WriteHeader(500)
		return
	}



}
func (s *Service) findAllEvent(w http.ResponseWriter, r *http.Request) {
   e, err := s.data.FindAllEventAvailable()
   if err != nil {
      fmt.Fprintf(w,`event not present or not correct the used word %s`,err )
	  w.WriteHeader(500)
	  return
   }
   w.Header().Set("content-type","application/json; charset=UTF8")
   encode := json.NewEncoder(w)
   err = encode.Encode(e)
   if err != nil {
	fmt.Fprintf(w, `error %s`, err)
   }

}

func Server(endpoint,tls string, data persistence.Database) (chan error, chan error) {
	r := mux.NewRouter()
	handler := NewService(data)
	serviceRouter := r.PathPrefix("/event").Subrouter()
	serviceRouter.Methods("POST").Path("[searchcriteria]/search").HandlerFunc(handler.addEvent)
	serviceRouter.Methods("GET").Path("/findall").HandlerFunc(handler.findAllEvent)
	serviceRouter.Methods("GET").Path("/findevent").HandlerFunc(handler.findEvent)
	httpApi := make(chan error)
	https := make(chan error)
     go func() {
       httpApi <- http.ListenAndServe(endpoint, r)
	 }()
	 go func ()  {
		https <- http.ListenAndServeTLS(tls, "cert.pem","key.pem",r)
	 }()
	return httpApi, https
}