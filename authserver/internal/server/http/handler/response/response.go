package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	AppCode        int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequestFormat(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "invalid request format.",
		ErrorText:      err.Error(),
	}
}

func ErrInternalServer(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error",
		ErrorText:      err.Error(),
	}
}

var UserRegistered = &Response{HTTPStatusCode: 200, StatusText: "User successfully registered and authenticated."}
var UserAuthenticated = &Response{HTTPStatusCode: 200, StatusText: "User successfully authenticated."}
var UserSignInOk = &Response{HTTPStatusCode: 200, StatusText: "Sign in ok."}
var RefreshAccessToken = &Response{HTTPStatusCode: 200, StatusText: "Refresh access token ok."}
var Logout = &Response{HTTPStatusCode: 200, StatusText: "Logout ok."}

var UserLogout = &Response{HTTPStatusCode: 200, StatusText: "User logout."}
var StatusCreated = &Response{HTTPStatusCode: 201, StatusText: "User Created."}

var ErrStatusBadRequest = &Response{HTTPStatusCode: 400, StatusText: "Status Bad Request."}

var ErrNotAuthorized = &Response{HTTPStatusCode: 401, StatusText: "User is not authorized."}
var ErrNotAuthenticated = &Response{HTTPStatusCode: 401, StatusText: "User not authenticated."}

var ErrStatusForbidden = &Response{HTTPStatusCode: 403, StatusText: "Forbidden."}

var ErrLoginIsTaken = &Response{HTTPStatusCode: 409, StatusText: "Login is taken."}

var ErrStatusUnprocessableEntity = &Response{HTTPStatusCode: 409, StatusText: "Status unprocessable entity."}

var ErrStatusBadGateway = &Response{HTTPStatusCode: 502, StatusText: "Status Bad Gateway."}
