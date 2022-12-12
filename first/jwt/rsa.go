package jwt

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	PrivKeyPath = "keys/pri8.pem"       // openssl genrsa -out app.rsa keysize
	PubKeyPath  = "keys/public_key.pem" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

type ParseRSAKeyFactory[T any] interface {
	ParseRsaKey() (T, error)
}

func GetParserFactory(key_type string, byte_key []byte) (interface{}, error) {
	if key_type == "pub" {
		return &ParserRsaPub{key: byte_key}, nil
	}
	return &ParserRsaPrivate{key: byte_key}, nil
}

type ParserRsaPub struct {
	key []byte
}

func (p *ParserRsaPub) ParseRsaKey() (*rsa.PublicKey, error) {
	var data, err = jwt.ParseRSAPublicKeyFromPEM(p.key)
	return data, err
}

type ParserRsaPrivate struct {
	key []byte
}

func (p *ParserRsaPrivate) ParseRsaKey() (*rsa.PrivateKey, error) {
	var data, err = jwt.ParseRSAPrivateKeyFromPEM(p.key)
	return data, err
}

func getByteRsaFile(key_path string) ([]byte, error) {
	var data, err = ioutil.ReadFile(key_path)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func getRsaParsePrivateKey(bite_key []byte) (*rsa.PrivateKey, error) {
	var private_key *rsa.PrivateKey
	factory, err := GetParserFactory("private", bite_key)

	if err != nil {
		return nil, err
	}
	private_key, err = factory.(ParseRSAKeyFactory[*rsa.PrivateKey]).ParseRsaKey()
	if err != nil {
		return nil, err
	}

	return private_key, nil

}

func getRsaParsePubKey(bite_key []byte) (*rsa.PublicKey, error) {
	var pub_key *rsa.PublicKey
	factory, err := GetParserFactory("pub", bite_key)
	if err != nil {
		return nil, err
	}
	pub_key, err = factory.(ParseRSAKeyFactory[*rsa.PublicKey]).ParseRsaKey()
	if err != nil {
		return nil, err
	}

	return pub_key, nil

}

func GetRsaPrivateKey() (*rsa.PrivateKey, error) {
	var private_key *rsa.PrivateKey
	byte_private_key, err := getByteRsaFile(PrivKeyPath)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	private_key, err = getRsaParsePrivateKey(byte_private_key)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return private_key, nil
}

func GetRsaPubKey() (*rsa.PublicKey, error) {

	var verifyKey *rsa.PublicKey

	verifyBytes, err := getByteRsaFile(PubKeyPath)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	verifyKey, err = getRsaParsePubKey(verifyBytes)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return verifyKey, nil
}
