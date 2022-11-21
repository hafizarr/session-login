package users

import (
	"log"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{repository: repository}
}

var SESSION_ID = "my-session-id"

func newCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}

func newPostgresStore() *pgstore.PGStore {
	url := "postgres://root:root@127.0.0.1:5432/session-login?sslmode=disable"
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err.Error())
		os.Exit(0)
	}

	return store
}

var store = newPostgresStore()

// var store2 = newCookieStore()

func (service *Service) Login(user *UserLogin, c echo.Context) (err error) {
	getUser, err := service.repository.Login(*user)
	if err != nil {
		return
	}

	session, err := store.Get(c.Request(), SESSION_ID)
	if err != nil {
		return
	}

	session.Values["username"] = getUser.Username
	session.Save(c.Request(), c.Response())

	return
}

func (service *Service) Logout(c echo.Context) (err error) {
	session, _ := store.Get(c.Request(), SESSION_ID)
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
	return nil
}

func (service *Service) Register(userRegister *UserRegister, c echo.Context) (err error) {
	getUser, err := service.repository.Register(*userRegister)
	if err != nil {
		return
	}

	session, err := store.Get(c.Request(), SESSION_ID)
	if err != nil {
		return
	}

	session.Values["username"] = getUser.Username
	session.Save(c.Request(), c.Response())

	return
}
