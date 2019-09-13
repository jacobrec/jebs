package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jacobrec/jebs/utils"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strings"
)

var (
	STORAGE = os.Getenv("IMAGE_DIRECTORY")
)

const VIEW_ENDPOINT = "/view/*id"
const DELETE_ENDPOINT = "/delete/*id"
const HANDLE_ENDPOINT = "/upload"

func listFiles() []string {
	// get all files
	files, _ := ioutil.ReadDir(STORAGE)

	// sort by last modified date, newest first
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().After(files[j].ModTime())
	})

	// get names
	s := make([]string, len(files))
	for i, f := range files {
		s[i] = f.Name()
	}

	return s
}

func uploadHandler(c *gin.Context) {

	fileHead, err := c.FormFile("filedata")
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	file, err := fileHead.Open()
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	path := generatePath()

	// Get file type
	mime := utils.GuessImageMimeTypes(file)
	ms := strings.Split(mime, "/")
	if len(ms) != 2 {
		c.AbortWithStatus(400)
		return
	}
	path += "." + ms[1]

	// Save file to disk
	file.Seek(0, 0)
	err = setImage(path, file)

	// return file name
	msg := "Upload Sucessful \n" + path + "\nPlease hit back to return to previous page, then refresh the page"
	c.Writer.Write([]byte(msg))
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

func deleteHandler(c *gin.Context) {
	image_path := STORAGE + c.Param("id")
	os.Remove(image_path)
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
