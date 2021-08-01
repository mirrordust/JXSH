package web

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Gate() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Auth TODO
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, present := c.Request.Header["Authorization"]
		if !present {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("missing token"))
			return
		}
		a := auth[0]
		tt := strings.Split(a, " ")
		if len(tt) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("missing token"))
			return
		}
		//type_ := tt[0]
		token := tt[1]
		//log.Printf("type: %v, toekn: %v", type_, token)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("missing token"))
			return
		}

		var req = OAuthRequest{Token: token}
		resp, err := checkToken(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("cannot validate token"))
			return
		}
		if resp.UserId == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("invalid token"))
			return
		}
		c.Set("userId", resp.UserId)

		//host := c.Request.Host
		//url := c.Request.URL
		//method := c.Request.Method
		//log.Printf("%s \t %s \t %s \t %s ", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
	}
}

// TODO
func checkToken(r *OAuthRequest) (OAuthResponse, error) {
	var s string
	if len(r.Token) > 3 {
		s = r.Token[:3]
	} else {
		s = r.Token
	}
	return OAuthResponse{UserId: s}, nil
}
