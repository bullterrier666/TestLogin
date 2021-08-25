package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bullterrier666/TestLogin/internal/app/middleware"
	"github.com/form3tech-oss/jwt-go"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	isError    bool   `json:"is_error"`
}

var (
	Login string = "Test"
	Pass  string = "pass"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post to auth")
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json recieved from client:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			isError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if Login != user.Login || Pass != user.Password {
		api.logger.Info("Incorrect authorization data")
		msg := Message{
			StatusCode: 404,
			Message:    "Incorrect authorization data",
			isError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		api.logger.Info("Can not claim jwt-token:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles. Try again",
			isError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		isError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

func (api *API) TestAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get test auth")
	msg := Message{
		StatusCode: 200,
		Message:    "Token works",
		isError:    false,
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(msg)

}
