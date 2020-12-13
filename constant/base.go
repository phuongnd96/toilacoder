package constant

import (
	"time"

	"github.com/kataras/iris/sessions"
)

var (
	CookieNameForSessionID = "Yhe8GUPoOJs1lHj6HLei"
	Sess                   = sessions.New(sessions.Config{
		Cookie:       CookieNameForSessionID,
		AllowReclaim: true,
		Expires:      time.Duration(time.Hour * time.Duration(48)), /*Có giá trị trong 2 ngày*/
	})
)

const DOMAIN = "https://toilacoder.com"

const REGEXPCharacters = "[^A-Za-z0-9ÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚĂĐĨŨƠàáâãèéêìíòóôõùúăđĩũơƯĂẠẢẤẦẨẪẬẮẰẲẴẶẸẺẼỀỀỂẾưăạảấầẩẫậắằẳẵặẹẻẽếềềểỄỆỈỊỌỎỐỒỔỖỘỚỜỞỠỢỤỦỨỪễệỉịọỏốồổỗộớờởỡợụủứừỬỮỰỲỴÝỶỸửữựỳỵýỷỹ]+"

var Account map[string]string = map[string]string{
	"lap-trinh-vien": "123456",
}

var MapCategory map[string]string = map[string]string{
	"flutter": "Flutter",
	"go":      "Go",
	"javascript": "Javascript",
}
