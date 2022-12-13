package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"unsafe"
	//jwt "github.com/golang-jwt/jwt/v4"
)

type Jwk struct {
	Keys []Key `json:"keys"`
}

func NewJwk() *Jwk {
	return &Jwk{}
}

type Key struct {
	E   string `json:"e"`
	Kty string `json:"kty"`
	N   string `json:"n"`
}

func GenJwk() (string, error) {
	jwk := NewJwk()
	fmt.Println("--------jwk------------")
	pub_key, err := GetRsaPubKey()
	if err != nil {
		return "nil", err
	}
	n := base64.URLEncoding.EncodeToString(pub_key.N.Bytes())
	e := base64.URLEncoding.EncodeToString(IntToByteArray(pub_key.E))

	jwk.Keys = append(jwk.Keys, Key{E: e, Kty: "RSA", N: n})

	ret, _ := json.Marshal(jwk)

	fmt.Println(string(ret))

	fmt.Println(n)
	fmt.Println(e)
	fmt.Println("----------------------------")

	return string(ret), nil
}

func IntToByteArray(num int) []byte {
	data := *(*[unsafe.Sizeof(num)]byte)(unsafe.Pointer(&num))
	return data[:]
}
