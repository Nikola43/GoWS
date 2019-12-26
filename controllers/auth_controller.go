package controllers

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/nikola43/testsocket/models"
	u "github.com/nikola43/testsocket/utils"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path

		//List of endpoints that doesn't require auth
		notAuth := []string{"/api/user/login"}
		//current request path
		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		//Token is missing, returns with error code 403 Unauthorized
		if tokenHeader == "" {
			u.RespondHttpError(w, http.StatusForbidden, "Auth token required")
			return
		}

		fullToken := strings.Split(tokenHeader, " ")
		if len(fullToken) != 2 {
			u.RespondHttpError(w, http.StatusForbidden, "Invalid/Malformed auth token")
			return
		}

		tokenPart := fullToken[1] //Grab the token part, what we are truly interested in
		tk := &models.ApiToken{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		//Malformed token, returns with http code 403 as usual
		if err != nil {
			u.RespondHttpError(w, http.StatusForbidden, "Malformed auth token")
			return
		}

		//Token is invalid, maybe not signed on this server
		if !token.Valid {
			u.RespondHttpError(w, http.StatusForbidden, "Invalid auth token")
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Println("User " + tk.Username) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.Username)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
