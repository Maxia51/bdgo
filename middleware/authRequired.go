package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthRequired middleware managing acces to a route regarding the role of the user
func AuthRequired(role ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)

		claims, err := checkTokenValidity(token)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fmt.Println(claims["role"])

		var hasRole bool;

		for _, element := range role {
			// index is the index where we are
			// element is the element from someSlice for where we are
			fmt.Println(element)

			if element == claims["role"] {
				hasRole = true
				break
			}
		}

		if hasRole {
			c.Next()
			return
		}

		c.JSON(http.StatusOK, gin.H{"error": "you don't have acces to this resource"})
		c.AbortWithStatus(http.StatusForbidden)
		return
		
	}
}

func checkTokenValidity(tokenString string) (jwt.MapClaims, error) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err == nil && token.Valid {
		return token.Claims.(jwt.MapClaims), nil
	}

	return nil, err
}
