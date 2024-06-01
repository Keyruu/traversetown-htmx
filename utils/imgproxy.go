package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/keyruu/traversetown-htmx/config"
)

func ResizeURL(url string, width int, height int) string {
	env := config.NewEnv()
	resize := fmt.Sprintf("resize:fill:%d:%d", width, height)
	base := base64.RawURLEncoding.EncodeToString([]byte(url))
	path := fmt.Sprintf("/%s/%s.webp", resize, base)
	return fmt.Sprintf("%s%s", env.ImgproxyUrl, SignURL(path))
}

func SignURL(path string) string {
	env := config.NewEnv()

	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(env.ImgproxyKey); err != nil {
		log.Fatal(err)
	}

	if saltBin, err = hex.DecodeString(env.ImgproxySalt); err != nil {
		log.Fatal(err)
	}

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("/%s%s", signature, path)
}
