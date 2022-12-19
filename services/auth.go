package services

import (
	"assignment2/dto"
	"assignment2/repository"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	b64 "encoding/base64"
	"fmt"
	"io"
)

var secret = "hacktiv8hacktiv8hacktiv8hacktiv8"

type AuthService struct {
	AuthRepo repository.AuthRepository
	DB       *sql.DB
}

func NewAuthService(db *sql.DB) AuthServiceProvider {
	return AuthService{
		DB:       db,
		AuthRepo: repository.AuthRepository{db},
	}
}

type AuthServiceProvider interface {
	Login(req dto.LoginDto) (res dto.LoginDto, err error)
	GetUser(username string) (res dto.RegisterDto, err error)
	UserRegister(req dto.RegisterDto) (res dto.RegisterDto, err error)
}

func (s AuthService) Login(req dto.LoginDto) (res dto.LoginDto, err error) {
	user, err := s.AuthRepo.GetUser(req.Username)
	if err != nil {
		return res, err
	}
	user.Password = decrypt(user.Password, secret)
	if user.Password != req.Password {
		return req, fmt.Errorf("Password yang Anda masukkan salah")
	}
	return req, nil
}

func (s AuthService) GetUser(username string) (res dto.RegisterDto, err error) {
	user, err := s.AuthRepo.GetUser(username)
	if err != nil {
		return res, err
	}
	return user, nil
}

func (s AuthService) UserRegister(req dto.RegisterDto) (res dto.RegisterDto, err error) {
	req.Password = encrypt(req.Password, secret)
	fmt.Printf("pass : %v", req.Password)
	user, err := s.AuthRepo.GetUser(req.Username)
	if err != nil {
		fmt.Printf("error get : %v\n", err.Error())
		if err == sql.ErrNoRows {
			err = s.AuthRepo.InsertUser(req)
			if err != nil {
				fmt.Printf("error insert : %v\n", err.Error())
				return user, err
			}
		} else {
			return user, err
		}
	}
	return user, nil
}

func encrypt(text string, passphrase string) string {
	salt := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err.Error())
	}

	key, iv := __DeriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	pad := __PKCS5Padding([]byte(text), block.BlockSize())
	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(pad))
	ecb.CryptBlocks(encrypted, pad)

	return b64.StdEncoding.EncodeToString([]byte("Salted__" + string(salt) + string(encrypted)))
}

// Decrypts encrypted text with the passphrase
func decrypt(encrypted string, passphrase string) string {
	ct, _ := b64.StdEncoding.DecodeString(encrypted)
	if len(ct) < 16 || string(ct[:8]) != "Salted__" {
		return ""
	}

	salt := ct[8:16]
	ct = ct[16:]
	key, iv := __DeriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	dst := make([]byte, len(ct))
	cbc.CryptBlocks(dst, ct)

	return string(__PKCS5Trimming(dst))
}

func __PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func __PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func __DeriveKeyAndIv(passphrase string, salt string) (string, string) {
	salted := ""
	dI := ""

	for len(salted) < 48 {
		md := md5.New()
		md.Write([]byte(dI + passphrase + salt))
		dM := md.Sum(nil)
		dI = string(dM[:16])
		salted = salted + dI
	}

	key := salted[0:32]
	iv := salted[32:48]

	return key, iv
}
