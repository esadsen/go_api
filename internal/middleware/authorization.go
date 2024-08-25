package middleware

import (
	"errors"
	"net/http"

	"github.com/esadsen/go_api/internal/tools"
	"github.com/esadsen/go_api/api"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or taken.")

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		println(err)
		
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database * tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil{
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if (loginDetails == nil || (token != (*loginDetails).AuthToken)){
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}