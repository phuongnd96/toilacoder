package helper

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/crypto/bcrypt"
)

func DecodeDataFromJsonFile(f *os.File, data interface{}) error {
	jsonParser := jsoniter.NewDecoder(f)
	err := jsonParser.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

func EqualIntSlice(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func CheckStringElementInSlice(list []string, str string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash kiểm tra mật khẩu dùng bcrypt
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type MenuItems struct {
	Text string   `json:"text"`
	Link string   `json:"link"`
	Sub  []string `json:"sub"`
}

type FooterItems struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Type  string `json:"type"`
}

type FooterCol struct {
	ColumnName string        `json:"column_name"`
	Items      []FooterItems `json:"items"`
}

type widget struct {
	WidgetType string        `json:"widget_type"`
	Code       template.HTML `json:"code"`
}

type Layout struct {
	LayoutType      int32    `json:"layout_type"`
	Widgets         []widget `json:"widgets"`
	Incontainer     bool     `json:"in_container"`
	BackgroundColor string   `json:"background_color"`
	BackgroundImage string   `json:"background_image"`
}

func ListLayout() []Layout {
	jsonFile, err := os.Open("layout.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var arr []Layout
	err = json.Unmarshal([]byte([]byte(byteValue)), &arr)
	if err != nil {
		fmt.Println(err)
	}
	return arr
}

func ListMenu() []MenuItems {
	jsonFile, err := os.Open("menu.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var arr []MenuItems
	json.Unmarshal([]byte([]byte(byteValue)), &arr)

	return arr
}

type BannerInfo struct {
	Id       int32
	Img      string
	Link     string
	IsActive bool
}


type GetPost struct {
	Id          string
	Title       string
	Thumbnail   string
	Description string
	CreatedAt   string
	CreatedBy   string
}

func ListFooter() []FooterCol {
	jsonFile, err := os.Open("footer.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var arr []FooterCol
	err = json.Unmarshal([]byte([]byte(byteValue)), &arr)
	if err != nil {
		fmt.Println(err)
	}
	return arr
}
