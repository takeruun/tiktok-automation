package main

import (
	"app/utils"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const AUTH_URL = "https://www.tiktok.com/v2/auth/authorize/"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/auth-init", authInit)
	e.GET("/auth-callback", authCallback)

	e.Logger.Fatal(e.Start(":3000"))
}

func authInit(c echo.Context) error {
	// Generate a random CSRF state
	csrfState := make([]byte, 16)
	if _, err := rand.Read(csrfState); err != nil {
		return err
	}
	csrfStateHex := hex.EncodeToString(csrfState)

	// Set the CSRF state as a cookie with a max age of 60 seconds
	cookie := new(http.Cookie)
	cookie.Name = "csrfState"
	cookie.Value = csrfStateHex
	cookie.MaxAge = 60
	c.SetCookie(cookie)

	challengeCode := utils.GenerateChallengeCode()
	url := AUTH_URL +
		"?client_key=" + os.Getenv("CLIENT_KEY") +
		"&code_challenge=" + challengeCode +
		"&redirect_uri=http://localhost:3000/auth-callback" +
		"&scope=user.info.basic" +
		"&response_type=code" +
		"&code_challenge_method=S256" +
		"&state=" + csrfStateHex

	return c.Redirect(http.StatusFound, url)
}

func authCallback(c echo.Context) error {
	return c.String(http.StatusOK, "Auth")
}
