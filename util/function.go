package util

import (
	"math"
	"bytes"
	"crypto/des"
	"errors"
	"encoding/hex"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"unsafe"
	"reflect"
	"encoding/base64"
)

// ApiJsonResult 接口json返回值
type ApiJsonResult struct {
	Code string 		`json:"code"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}

type ApiPagingJsonResult struct {
	TotalCount int 		`json:"totalCount"`
	TotalPageCount int 	`json:"totalPageCount"`
	Page int 			`json:"page"`
	PageSize int 		`json:"pageSize"`
	Lists interface{} 	`json:"lists"`
}
// NewApiJsonResult 创建 ApiJsonResult
func NewApiJsonResult(code string, message string) *ApiJsonResult {
	return &ApiJsonResult{
		Code: code,
		Message: message,
	}
}

// Simple 普通数据返回
func (o *ApiJsonResult) Simple (data interface{}) *ApiJsonResult {
	o.Data = data
	return o
}

// Paging 分页
func (o *ApiJsonResult) Paging(lists interface{}, totalCount, page int, pageSize int) *ApiJsonResult {
	totalPageCount := 0
	if pageSize > 0 {
		totalPageCount = int(math.Ceil(float64(totalCount/pageSize)))
	}
	o.Data = ApiPagingJsonResult{
		TotalCount: totalCount,
		TotalPageCount: totalPageCount,
		Page: page,
		PageSize: pageSize,
		Lists: lists,
	}
	return o
}

// GetApiJsonResult 获取接口json返回值
func GetApiJsonResult(code string, message string, data interface{}) *ApiJsonResult {
	return &ApiJsonResult{
		Code: code,
		Data: data,
		Message: message,
	}
}

// GetApiJsonPagingResult 获取接口分页json返回值
func GetApiJsonPagingResult(code string, message string, lists interface{}, totalCount, page int, pageSize int) *ApiJsonResult {
	totalPageCount := 0
	if pageSize > 0 {
		totalPageCount = int(math.Ceil(float64(totalCount/pageSize)))
	}
	return &ApiJsonResult{
		Code: code,
		Data: ApiPagingJsonResult {
			TotalCount: totalCount,
			TotalPageCount: totalPageCount,
			Page: page,
			PageSize: pageSize,
			Lists: lists,
		},
		Message: message,
	}
}

//BASE64加密
var BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
func B64_Encode(data string) string {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncoding(BASE64Table)
	return coder.EncodeToString(content)
}

func B64_Decode(data string) string {
	coder := base64.NewEncoding(BASE64Table)
	result, _ := coder.DecodeString(data)
	return *(*string)(unsafe.Pointer(&result))
}

//md5加密
func MD5_Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}


func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}


//DES加密
//key := []byte("2fa6c1e9")
//str :="I love this beautiful world!"
//strEncrypted, err := Encrypt(str, key)
//if err != nil {
//log.Fatal(err)
//}
//fmt.Println("Encrypted:", strEncrypted)
//strDecrypted, err := Decrypt(strEncrypted, key)
//if err != nil {
//log.Fatal(err)
//}
//fmt.Println("Decrypted:", strDecrypted)

func DES_Encrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func DES_Decrypt(decrypted string , key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}



//AES加密
//str := "I love this beautiful world!"
//key := []byte{0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
//0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
//0xBA, 0x37, 0x2F, 0x02, 0xC3, 0x92, 0x1F, 0x7D,
//0x7A, 0x3D, 0x5F, 0x06, 0x41, 0x9B, 0x3F, 0x2D,
//}
//strEncrypted,err := Encrypt(str, key)
//if err != nil {
//log.Error("Encrypted err:",err)
//}
//fmt.Println("Encrypted:",strEncrypted)
//strDecrypted,err := Decrypt(strEncrypted, key)
//if err != nil {
//log.Error("Decrypted err:",err)
//}
//fmt.Println("Decrypted:",strDecrypted)
func AES_Encrypt(text string, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}

func AES_Decrypt(encrypted string, key []byte) (string, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
}