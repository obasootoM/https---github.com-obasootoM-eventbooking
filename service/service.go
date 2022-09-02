package service

import (
	"eventsbook/persistence"
	"net/http"

	"github.com/gorilla/mux"
)


type Service struct {
	database *persistence.Database
}
func NewService(database *persistence.Database) *Service {
	return &Service{
		database: database,
	}
}
func (s *Service) findEvent(w http.ResponseWriter,r *http.Request) {
     
}
func (s *Service) addEvent(w http.ResponseWriter, r *http.Request) {

}
func (s *Service) findAllEvent(w http.ResponseWriter, r *http.Request) {

}

func Server(endpoint string, database persistence.Database) {
	r := mux.NewRouter()
	handler := NewService(&database)
	serviceRouter := r.PathPrefix("").Subrouter()
	serviceRouter.Methods("POST").Path("").HandlerFunc(handler.addEvent)
	serviceRouter.Methods("GET").Path("").HandlerFunc(handler.findAllEvent)
	serviceRouter.Methods("GET").Path("").HandlerFunc(handler.findEvent)
}