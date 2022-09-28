package controllers

import (
	"context"
	"net/http"

	auth "github.com/byeol-i/battery-level-checker/pkg/authentication/firebase"
)

type AuthControllers struct {
	app *auth.FirebaseApp
}

func NewAuthController() *AuthControllers {
	return &AuthControllers{}
}

func (hdl *AuthControllers) CreateNewToken(resp http.ResponseWriter, req *http.Request) {
	
}

func (hdl *AuthControllers) LoginTest(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "text/html")

	script := `
	<html>
	<head>
		<script type="module">

		</script>
	</head>
	<body></body>
  	</html>`
	content := []byte{}
	content = []byte(script)
	resp.Write(
		content,
	)
}

func (hdl *AuthControllers) CreateCustom(resp http.ResponseWriter, req *http.Request) {
	ctx := context.Background();

	token, err := hdl.app.CreateCustomToken(ctx, req.Header.Get("token"))
	if err != nil {
		respondError(resp, 404, "token is not valid")
	}

	respondJSON(resp, 200, "done", token)
}
