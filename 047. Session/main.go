package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// User holds a users account information
type User struct {
	Username      string
	Authenticated bool
}

// store will hold all session data
var store *sessions.CookieStore

// tpl holds all parsed templates
var tpl *template.Template

//帳號密碼
var userPwdPair 

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(User{})

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func Login(c *gin.Context) {

	if c.Request.Method == "GET" {
		// Get a session. Get() always returns a session, even if empty.
		session, err := store.Get(c.Request, "session-name")
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		session.ID

		// 顯示登入介面
		t, _ := template.ParseFiles("047. Session/login.gtpl")
		t.Execute(c.Writer, nil)

		msgs, ok := c.Request.URL.Query()["msg"]
		if !ok || len(msgs[0]) < 1 {
			log.Println("Url Param 'key' is missing")
		}

		if "nilUsernameOrPwd" == msgs[0] {
			fmt.Fprintf(c.Writer, "帳號或密碼不可以空白")
		} else {
			fmt.Fprintf(c.Writer, "%s", strings.Join(msgs, ","))
		}
	}

	if c.Request.Method == "POST" {
		// 若請求登入資料，處理登入的邏輯判斷
		err := c.Request.ParseForm()
		if nil == err {
			username := c.Request.Form["username"][0]
			password := c.Request.Form["password"][0]
			fmt.Printf("u:%s, p:%s\n", username, password)

			if "" == username || "" == password {
				// 如果帳號或密碼空白，重新導向回登入頁面，並且帶提示資訊
				c.Redirect(302, "login?msg=nilUsernameOrPwd")
			} else {
				// 如果沒有空白，開始驗證資訊
				if userPwdPair[username] == password {
					fmt.Fprintf(c.Writer, "登入成功\n")
					//產生並儲存 Session
					sess, err := store.New(c.Request, username)
					err = sess.Save(c.Request, c.Writer)
					if nil != err {
						fmt.Fprintf(c.Writer, "session contant: %v", sess.Values)
					} else {
						fmt.Println(err.Error())
					}
				} else {
					fmt.Fprintf(c.Writer, "登入失敗，帳號或密碼不符")
				}
			}
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	router := gin.Default()
	//設定存取的路由
	router.GET("/login", Login)
	router.POST("/login", Login)
	//router.POST("/log", MyHandler)
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
