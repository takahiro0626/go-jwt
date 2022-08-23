package main

import (
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	email := "test@email.com"

	now := time.Now()
	lifetime := 30 * time.Minute

	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": email,
        "iat":    now.Unix(),
        "exp":    now.Add(lifetime).Unix(),
    })

	fmt.Println(token)
	// &{ 0xc0000a8198 map[alg:HS256 typ:JWT] map[exp:1661260797 iat:1661258997 user_id:test@email.com]  false}
	fmt.Println(token.SignedString([]byte(secret)))
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjEyNjA3OTcsImlhdCI6MTY2MTI1ODk5NywidXNlcl9pZCI6InRlc3RAZW1haWwuY29tIn0.UAlarHRf0EwgawJbzLrEQGJIF-VaTasvtdfDmZmT1B4 <nil>
}