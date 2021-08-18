package configs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	. "gopkg.in/square/go-jose.v2"
)

type Config struct {
	User     string
	Password string
	Server   string
	Database string
}

// JSON Web Encryption (RFC 7516)
func GetJWE() string {
	fmt.Println("================= START TEST RSA")
	// // Generate a public/private key pair to use for this example.

	// Create the keys
	// priv, pub := GenerateRsaKeyPair()

	// pub_pem, _ := ExportRsaPublicKeyAsPemStr(pub)
	// fmt.Println("pub_pem", pub_pem)

	// priv_pem := ExportRsaPrivateKeyAsPemStr(priv)
	// fmt.Println("priv_pem", priv_pem)

	// pub_parsed, _ := ParseRsaPublicKeyFromPemStr(pub_pem)
	// fmt.Println("pub_parsed", pub_parsed)

	privateKey, err := ParseRsaPrivateKeyFromPemStr("-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAsvjB5Z7lKz1scyp1J333wHYTQSMVmbZWSS4T5KxWQMt6HnNQ\niN+RQLoy5BsJu+i+OJAZY83zdDBxyOE1/HhR6LyZUYiJBfv66kza7gD8gFXHcdJz\nhlBYL6woCd58YMgsClVTCs80Jv8P9piCXIk97MlQEBeFv5gzrjBMQnD6yuX0WCRf\nm2o00DGB6N3ultMTGPFQki0Gmdg0eWz/00ElwHjc9w9V+uBrI/Owg1AXgRtPzdOT\nfiQUR5XnzW8ca908xniC6YfAF6w5JHMGlQdwuuR7fP0IbU+418b6UMp/WldiULoC\nJNgQIeyGmKrIBYE+enum4L1FcM324L8Cr5xkSQIDAQABAoIBAQCMk3IRv4Y4OY32\nw9GzzHEO/m/PSRSNYThTUNsVUfyJ9omcY21NpXW9RsP0gfxaWc0YNq+KwllI2Uh2\nxdcW0RddXM58yWb/iVYkNWv9aE7tKvN2xxww0ukBwRLN9bDygCh/kVeh9PD89HmA\nc3hp4O8AD3xQt3k57f4iAxPTWFTViwvl8Zm2+q8GeXD3TgjagVXQox9rk+38ZOiZ\nODDVmd6GdWd9T408ywRoEKhiLWVcqvIa/oDnYPdvGWJXSPNz+kYeCdJQWGDDh1bD\nB4QQqO7e+9H9FziqHjJDDCHlqhIA+9boXNQV6tijnz6HqXUi8g86fOG8rARBpLUG\n+QvI/5QBAoGBANe72rimxi5HRxEFcKdbmU5wS2fp+9SmE29jwvcjkOGmD5QfQQtA\n7aAmsojsFR9hmgMptoKovdrm3S+43kjnc44bWrqYUo0OkASnVdw8RuniWTtjDO3l\nyxwloayOUNfjsFErmaykbCajYP+8+MZRdc2Ch2CbG1SK118vCmkenIyBAoGBANRg\nWDH4e6JPrt79LrnS3URc7tqYNj1xhNIksaXAqICdOVG9xDk/ywMcOL4v0209N8Df\nXrI53IevEEnA/c2PULEojEw+Hh0cE7l2ZNZSBgIYMO8lDoPsJ30vZWg/VHO7q4S3\n/36SGy+gpq3bj6YG+uKOaiVK65Kwjq+0suuQBJPJAoGAVtgKbjbEpLCQzStRdC2N\ng2P1Fvm8JJhOTpLsUyyAEDYXI1uJhYw3nDa3vCK9Rgq2QvuxuttOwiYZeDKFWPfD\nLruYRHE5Ggt4skbia2vgaBp0LVHsTIAqyUPk11/a72169tIEZNkzqEIwtAJQaxUz\nxOz8rkkCMYBKvshi+noWLwECgYATlFZNNRVg32vFzx2xsAkKNm8m5AWMIJ0YKp3e\nJCbkxJn2OUKP7Juwy7dHgW1CkvRC6dT3JXJ5Z+BejFZZzy6nRo+4r7ljAaWn4Yxm\nXeWD5+sLASWXb/wsBGSa+cu1Z8XmVHhPmPS8rVjwil9x9Q7IOEz85y1UtsZQ2J3u\nLbbeUQKBgGMhCdLtpxVWLSfpiLH6NbW3GFE8EZC4/qd+3VFSOlDmN2+pFBWWGpO5\nPHbBDUgdf/60NxXEU7QRHZl2k6Wn7Fqg7/WoNauM52PZSmFLNMMl0a+gGCRB1Cyf\nPDyPlrl53y5zxese6eloHXxk2A6WpbJFOjmtWaKszvzqQL+VdU0p\n-----END RSA PRIVATE KEY-----")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("privateKey", privateKey)

	// _, err = rsa.GenerateKey(rand.Reader, 2048)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Instantiate an encrypter using RSA-OAEP with AES128-GCM. An error would
	// indicate that the selected algorithm(s) are not currently supported.
	publicKey := &privateKey.PublicKey
	fmt.Println("publicKey", publicKey)
	pub_pem, _ := ExportRsaPublicKeyAsPemStr(publicKey)
	fmt.Println("pub_pem", pub_pem)
	encrypter, err := NewEncrypter(A128GCM, Recipient{Algorithm: RSA_OAEP, Key: publicKey}, nil)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("encrypter", encrypter)

	// Encrypt a sample plaintext. Calling the encrypter returns an encrypted
	// JWE object, which can then be serialized for output afterwards. An error
	// would indicate a problem in an underlying cryptographic primitive.
	var plaintext = []byte("JEKO WW DUET ENCRYPT")
	fmt.Println("plaintext", plaintext)
	object, err := encrypter.Encrypt(plaintext)
	if err != nil {
		fmt.Println(err)
	}

	// Serialize the encrypted object using the full serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.

	serialized := object.FullSerialize()
	cserialized, _ := object.CompactSerialize()
	fmt.Println("===============================")
	fmt.Println("serialized", serialized)
	fmt.Println("serialized", cserialized)
	// serialized {"protected":"eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkExMjhHQ00ifQ","encrypted_key":"AyIhM2XVeCbLV4mcASvtf101Cdj99GXg5pRiF8XEXMjSifG5MOU4uHXw-d1wTZK5hrwr7bqjgbN_REo2p8DwblO6tE2othle4_HnElv6SlrT_eo6WAK5uBVtAjzI-ijZSte5-pQjJ86wCKnDx0y5CLIQ-phbmhK0Ou-RumS-tsVZVkJyEeGfzl7iroM5y8IWdr8XOJL-1AGkTdw9JasssZud4mFacFBbeplxzNzOmWD4jt6P-Mzt_qIJGQI0pJ3zrIxCYyHUagp-Ttv8HY0VM0Kwog8D1LlRHvXPyO2FjLnvv8iV2iSEblgNxcmYWec0hl78uFI56bTPZBij5KxQiA","iv":"GzhDWQxVgVjl7L8o","ciphertext":"Umih_GqlxAExp2z0Czu5DWlBRtE","tag":"OH81QU4Re4C2QSF15bVVDA"}
	// serialized eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkExMjhHQ00ifQ.AyIhM2XVeCbLV4mcASvtf101Cdj99GXg5pRiF8XEXMjSifG5MOU4uHXw-d1wTZK5hrwr7bqjgbN_REo2p8DwblO6tE2othle4_HnElv6SlrT_eo6WAK5uBVtAjzI-ijZSte5-pQjJ86wCKnDx0y5CLIQ-phbmhK0Ou-RumS-tsVZVkJyEeGfzl7iroM5y8IWdr8XOJL-1AGkTdw9JasssZud4mFacFBbeplxzNzOmWD4jt6P-Mzt_qIJGQI0pJ3zrIxCYyHUagp-Ttv8HY0VM0Kwog8D1LlRHvXPyO2FjLnvv8iV2iSEblgNxcmYWec0hl78uFI56bTPZBij5KxQiA.GzhDWQxVgVjl7L8o.Umih_GqlxAExp2z0Czu5DWlBRtE.OH81QU4Re4C2QSF15bVVDA
	// Parse the serialized, encrypted JWE object. An error would indicate that
	// the given input did not represent a valid message.
	fmt.Println("===============================")
	test, err := ParseEncrypted(`{"protected":"eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkExMjhHQ00ifQ","encrypted_key":"AyIhM2XVeCbLV4mcASvtf101Cdj99GXg5pRiF8XEXMjSifG5MOU4uHXw-d1wTZK5hrwr7bqjgbN_REo2p8DwblO6tE2othle4_HnElv6SlrT_eo6WAK5uBVtAjzI-ijZSte5-pQjJ86wCKnDx0y5CLIQ-phbmhK0Ou-RumS-tsVZVkJyEeGfzl7iroM5y8IWdr8XOJL-1AGkTdw9JasssZud4mFacFBbeplxzNzOmWD4jt6P-Mzt_qIJGQI0pJ3zrIxCYyHUagp-Ttv8HY0VM0Kwog8D1LlRHvXPyO2FjLnvv8iV2iSEblgNxcmYWec0hl78uFI56bTPZBij5KxQiA","iv":"GzhDWQxVgVjl7L8o","ciphertext":"Umih_GqlxAExp2z0Czu5DWlBRtE","tag":"OH81QU4Re4C2QSF15bVVDA"}`)
	if err != nil {
		fmt.Println("err", err)
	}

	decryptedTest, err := test.Decrypt(privateKey)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("IKI === ", string(decryptedTest))
	fmt.Println("===============================")

	// object, err = ParseEncrypted("eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkExMjhHQ00ifQ.N7rI01fEqdghD5F22SsTstnmrQrfw5XHXBlgNxfBTNcnwOOyx0hngeSqIfR6VHroO1UbRaROo_funp-Jos1xkUByYxURsGqpEBgsTFGsOwO-HDZCURrHBFFpYgiHxI4MdLTySHHKnce1cYLG_CxOF8MtroNtdaYQsXpL_Ay6AQzcmrpTj_S6z0p0fXixMZ2BB_oqGUkGWViCs2NF7kPoSaS3y3413Bmkyl5QA-Or8jsvN8uT_UewgbFLAsywor8xGVFVXFv4AgQv6PDoa_hDsAIVjGvACHRZzoBfkVXtpal5i1QDH-fAWgwrkbmOC8CdyDM66CYnro_oo-wmpaxeQg.9y9uvEacSqqn8emY.ref5r0vJPSOImGxkn5wPt_GsK7U.xrCnpbvZNpZKMyGbOfLWjg")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("object", object)

	// Now we can decrypt and get back our original plaintext. An error here
	// would indicate that the message failed to decrypt, e.g. because the auth
	// tag was broken or the message was tampered with.
	// decrypted, err := object.Decrypt(privateKey)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf(string(decrypted))
	return string(decryptedTest)
}

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return privkey, &privkey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}

func GetConfig() Config {
	// JWE := GetJWE()
	// fmt.Println("JWE", JWE)
	// file, _ := os.Open("//pmiidsubdev33/invdashboard$/rfid/config.json")
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func (c Config) GetMyConnectionInfo() string {
	return fmt.Sprintf(
		"server=%s;user id=%s;password=%s;database=%s",
		c.Server, c.User, c.Password, c.Database)
}
