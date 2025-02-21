package middleware

import (
	"CoinTransfer/internal/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// const tokenTTL = 12 * time.Hour

// type tokenClaims struct {
// 	jwt.StandardClaims
// 	UserId int `json:"user_id"`
// }

func JWTAuthMiddleware(authService services.Authorization) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		token := tokenParts[1]
		userId, err := authService.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}

// func CreateToken(userId int) (string, error) {

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
// 		jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 		userId,
// 	})

// 	return token.SignedString([]byte(os.Getenv("singingKey")))
// }

// func ParseToken(accessToken string) (int, error) {
// 	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("singingKey")), nil
// 	})

// 	if err != nil {
// 		return 0, err
// 	}

// 	claims, ok := token.Claims.(*tokenClaims)
// 	if !ok || !token.Valid {
// 		return 0, fmt.Errorf("invalid token")
// 	}
// 	return claims.UserId, nil
// }
