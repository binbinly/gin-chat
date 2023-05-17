package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Payload is the data of the JSON web token.
// 主要是配合jwt来生成用户登录token
type Payload struct {
	UserID int
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (any, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the payloads if the token was valid.
func Parse(tokenString string, secret string) (*Payload, error) {
	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return nil, err
	}
	// Read the token if it's valid.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payloads := &Payload{}
		payloads.UserID = int(claims["user_id"].(float64))
		return payloads, nil
	}

	// Other errors.
	return nil, err
}

// Sign signs the payload with the specified secret.
func Sign(ctx context.Context, payload map[string]any, secret string, timeout int64) (tokenString string, err error) {
	// The token content.
	// iss: （Issuer）签发者
	// iat: （Issued At）签发时间，用Unix时间戳表示
	// exp: （Expiration Time）过期时间，用Unix时间戳表示
	// aud: （Audience）接收该JWT的一方
	// sub: （Subject）该JWT的主题
	// nbf: （Not Before）不要早于这个时间
	// jti: （JWT ID）用于标识JWT的唯一ID
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["nbf"] = now
	claims["iat"] = now
	claims["exp"] = now + timeout

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
