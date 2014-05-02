//Authentication REST API
package auth

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"github.com/guregodevo/pastis"
	"time"
)

//Authentication Services
type AuthenticationDomain struct {
	TokenRepository   *TokenRepository
	AccountRepository *AccountRepository
}

type AuthenticationResource struct {
	Service *AuthenticationDomain
}

//Authentication Error type
type AuthenticationError struct {
	When time.Time
	What string
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

//Signup creates a new user account. It will validate that a user with same credentials is not already registered.
func (service *AuthenticationDomain) Signup(email string, password string, companyName string) (string, int64, error) {
	existingAcc, err := service.AccountRepository.FindEmail(email)
	log.Printf("Email '%v' %v \n", existingAcc, err)
	if existingAcc != nil {
		log.Printf("Email '%v' already registered. %v \n", email, existingAcc.Email)
		return "", http.StatusConflict, &AuthenticationError{time.Now(), "The email already exists."}
	}
	if err != nil {
		log.Fatal("Could not check user '%v' \n", email, err)
		return "", http.StatusInternalServerError, &AuthenticationError{time.Now(), "Technical error while getting user record"}
	}
	acc, erro := service.AccountRepository.Create(email, companyName, password)
	if erro != nil {
		log.Fatal("Could not create User '%v' \n", email, erro)
		return "", http.StatusInternalServerError, &AuthenticationError{time.Now(), "Technical error while creating user record"}
	}
	token, erroToken := service.TokenRepository.Put(acc)
	if erroToken != nil {
		log.Fatal("Could not create User '%v' token \n", email, erroToken)
		return "", http.StatusInternalServerError, &AuthenticationError{time.Now(), "Technical Error while getting user token"}
	}
	return token, http.StatusOK, nil
}

//GetUser provides user account data
func (service *AuthenticationDomain) GetUser(email string, password string) (string, error) {
	acc, err := service.AccountRepository.FindEmailAndPassword(email, password)
	if err != nil {
		log.Printf("User '%v':'%v' has input wrong Credentials \n", email, password[1:])
		return "", &AuthenticationError{time.Now(), "Wrong user or password"}
	}
	token := TokenOf(acc)
	return token.Encrypt(), nil
}

//RESTful GET of user account data
func (api AuthenticationResource) GET(values url.Values) (int, interface{}) {
	username := values.Get("username")
	password := values.Get("password")

	token, err := api.Service.GetUser(username, password)

	if err == nil {
		data := map[string]string{"token": token}
		fmt.Println("Return token [%v]", token)
		return http.StatusOK, data
	}
	return http.StatusUnauthorized, pastis.ErrorResponse(err)
}

type AuthenticationRequest struct {
	Username, Password, Company string
}

//RESTful POST of a new user account
func (api AuthenticationResource) POST(values url.Values, t AuthenticationRequest) (int64, interface{}) {
	token, status, err := api.Service.Signup(t.Username, t.Password, t.Company)
	if err == nil {
		data := map[string]string{"token": token}
		fmt.Println("Return token [%v]", token)
		return http.StatusOK, data
	}
	return status, pastis.ErrorResponse(err)
}
