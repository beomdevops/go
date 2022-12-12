package jwt

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	UserId int
	Name   string
	jwt.RegisteredClaims
}

func NewMyClaims() *MyClaims {
	return &MyClaims{}
}

func (c *MyClaims) BuildUesrId(id int) *MyClaims {
	c.UserId = id
	return c
}
func (c *MyClaims) BuildName(name string) *MyClaims {
	c.Name = name
	return c
}

func GenJwt() string {
	var private_key, err = GetRsaPrivateKey()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var c = NewMyClaims().BuildName("park").BuildUesrId(1)
	c.Issuer = "ISSUER"
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	signedAuthToken, err := token.SignedString(private_key)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(signedAuthToken)
	fmt.Println("---------------------------------------------")
	//ParseJwt(signedAuthToken)

	return signedAuthToken
}

func ParseJwt(token_str string) {
	var claims = NewMyClaims()
	token, _ := jwt.ParseWithClaims(token_str, claims, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})
	fmt.Println(token)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println(claims.Name)
	fmt.Println(claims.UserId)
	fmt.Println(claims.Issuer)
}
