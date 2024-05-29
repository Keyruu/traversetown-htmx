package image

import (
	"fmt"

	"github.com/keyruu/traversetown-htmx/config"
)

func OptimizedURL(url string) string {
	env := config.NewEnv()
	resize := "resize:fill:1024:1024"
	blur := "blur:20"
	quality := "q:10"
	path := fmt.Sprintf("/%s/%s/%s/plain/%s@webp", quality, blur, resize, url)
	return fmt.Sprintf("%s%s", env.ImgproxyUrl, SignURL(path))
}
