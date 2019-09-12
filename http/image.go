package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var (
	STORAGE = os.Getenv("IMAGE_DIRECTORY")
)

const VIEW_ENDPOINT = "/view/*id"
const HANDLE_ENDPOINT = "/upload"

func main() {
}

func ginify(fn func(http.ResponseWriter, *http.Request)) func(*gin.Context) {
	return func(c *gin.Context) {
		fn(c.Writer, c.Request)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Println(r.PostFormValue("filedata"))
	r.ParseForm()
	fmt.Println("Form: ", r.Form)
	for k, v := range r.Form {
		fmt.Println("k:", k, "    v:", v)
	}

	if len(splits) != 2 {
		w.WriteHeader(400)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	path := generatePath()
	err = setImage(path, r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("400 - Bad Request"))
		return
	}

	w.Write([]byte(path))
	fmt.Println("Uploaded Image: " + path)
}

func viewHandler(c *gin.Context) {
	image_path := STORAGE + c.Param("id")
	fmt.Println("Retriving image: ", image_path)

	splits := strings.Split(string(image_path), ".")
	if len(splits) != 2 {
		c.AbortWithStatus(400)
		return
	}

	c.Header("Content-Type", "image/"+string(splits[1]))
	err := getImage(image_path, c.Writer)
	c.Header("Content-Type", "image/"+string(splits[1]))

	if err != nil {
		c.AbortWithStatus(404)
		return
	}
}

func preprocessImage(img io.ReadCloser) error {
	// TODO: downscale image maybe? Verify it indeed is an image file
	return nil
}

func generatePath() string {
	return STORAGE + "/" + getRandomString(12)
}

func setImage(path string, image io.ReadCloser) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	return copyData(image, file)
}

func getImage(path string, resp io.Writer) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	return copyData(file, resp)
}

func copyData(r io.Reader, w io.Writer) error {
	var err error
	buf := make([]byte, 20480)
	size := 0
	bytes := 0
	for err == nil {
		bytes, err = r.Read(buf)
		w.Write(buf[:bytes])
		size += bytes
	}
	if err != io.EOF {
		return err
	}
	return nil
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getRandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
