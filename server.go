package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func startServer() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		authURL := googleConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		html := `<h1>Codeforces Calendar Sync</h1>
                 <p>Sync upcoming contests to your Google Calendar automatically.</p>
                 <a href="` + authURL + `">Login with Google</a>`
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
	})

	r.GET("/auth/callback", func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.String(http.StatusBadRequest, "Code not found")
			return
		}

		// 1. Exchange Code for Token
		token, err := googleConfig.Exchange(context.Background(), code)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
			return
		}

		// 2. Get User Email
		client := googleConfig.Client(context.Background(), token)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get user info: "+err.Error())
			return
		}
		defer resp.Body.Close()

		var userInfo struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
			c.String(http.StatusInternalServerError, "Failed to parse user info")
			return
		}

		// 3. Save to DB
		if err := saveUser(userInfo.Email, token); err != nil {
			c.String(http.StatusInternalServerError, "Failed to save user: "+err.Error())
			return
		}

		c.String(http.StatusOK, "Success! You are now synced. Email: "+userInfo.Email)
	})

	r.Run(":8080")
}
