package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"ttms01/dao"
	"ttms01/model"
	"ttms01/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	user.Username = r.FormValue("username")
	user.Password = r.FormValue("password")
	user.Email = r.FormValue("email")
	checkuser, _ := dao.CheckUserName(user.Username)
	fmt.Println("checkuser1")
	if checkuser.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/register.html"))
		fmt.Println("checkuser2")
		t.Execute(w, nil)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		dao.AddUser(user)
		page, _ := dao.GetPageMovie("1")
		page.IsLogin = true
		fmt.Println("checkuser3")
		page.Username = user.Username
		t.Execute(w, page)
	}
}

//func Login(w http.ResponseWriter, r *http.Request) {
//	flag, session := dao.IsLogin(r)
//	fmt.Println(flag, "dwad")
//	if flag {
//		t := template.Must(template.ParseFiles("views/index.html"))
//		t.Execute(w, session)
//	} else {
//		username := r.PostFormValue("username")
//		password := r.PostFormValue("password")
//		fmt.Println(username, password)
//		user, _ := dao.CheckUserName(username)
//		if user.ID > 0 {
//			uuid := utils.CreatUUID()
//			sess := &model.Session{
//				SessionID: uuid,
//				UserName:  username,
//				UserID:    user.ID,
//			}
//
//			dao.AddSession(sess)
//			cookie := http.Cookie{
//				Name:     "user",
//				Value:    uuid,
//				HttpOnly: true,
//			}
//			http.SetCookie(w, &cookie)
//			fmt.Println("登录成功")
//			session.IsLogin = true
//			t := template.Must(template.ParseFiles("views/index.html"))
//			_ = t.Execute(w, session)
//		} else {
//			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
//			t.Execute(w, "")
//		}
//	}
//}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login1")
	flag, _ := dao.IsLogin(r)
	page, _ := dao.GetPageMovie("1")
	page.Username = r.FormValue("username")
	if flag {
		fmt.Println("fakelogin2")
		page.IsLogin = true
		if dao.IsAdmin(page.Username) {
			page.IsAdmin = true
		} else {
			page.IsAdmin = false
		}
		t := template.Must(template.ParseFiles("views/pages/login.html"))
		t.Execute(w, page)
	} else {
		fmt.Println("login2")
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		admin, _ := dao.CheckAdmin(username, password)
		if admin.ID > 0 {
			fmt.Println("login3")
			uuid := utils.CreatUUID()
			sess := &model.Session{
				SessionID: uuid,
				UserName:  username,
				UserID:    admin.ID,
			}

			//dao.AddSession(sess)
			dao.AddAdminSession(sess)
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			fmt.Println("login4")
			http.SetCookie(w, &cookie)
			//fmt.Println("登录成功")
			page.IsLogin = true
			page.IsAdmin = true
			//测试
			//GetPageMovie(w, r)

			//page.IsLogin = true
			//if dao.IsAdmin(page.Username) {
			//	page.IsAdmin = true
			//} else {
			//	page.IsAdmin = false
			//}

			now := time.Now()

			// 定义格式化模板
			format := "2006-01-02-15:04" // 24小时制，小时范围是0-23

			// 格式化时间
			formattedTime := now.Format(format)

			page.ComingMovies = dao.GetComingMovies("", "6", formattedTime)
			page.Movies = dao.GetHotMovies("", "6", formattedTime)
			page.ClassicMovies = dao.GetClassicMovies("", "6", formattedTime)

			//

			t := template.Must(template.ParseFiles("views/index.html"))
			_ = t.Execute(w, page)
		} else {
			user, _ := dao.CheckUserNameAndPassword(username, password)
			fmt.Println("loginn1")
			if user.ID > 0 {
				uuid := utils.CreatUUID()
				sess := &model.Session{
					SessionID: uuid,
					UserName:  username,
					UserID:    user.ID,
				}
				fmt.Println("loginn1,1")
				dao.AddSession(sess)
				cookie := http.Cookie{
					Name:     "user",
					Value:    uuid,
					HttpOnly: true,
				}
				fmt.Println("loginn1.2")
				http.SetCookie(w, &cookie)
				//fmt.Println("登录成功")
				page.IsLogin = true
				page.IsAdmin = false

				//
				fmt.Println("loginn2")

				now := time.Now()

				// 定义格式化模板
				format := "2006-01-02-15:04" // 24小时制，小时范围是0-23

				// 格式化时间
				formattedTime := now.Format(format)

				page.ComingMovies = dao.GetComingMovies("", "6", formattedTime)
				page.Movies = dao.GetHotMovies("", "6", formattedTime)
				page.ClassicMovies = dao.GetClassicMovies("", "6", formattedTime)
				//

				t := template.Must(template.ParseFiles("views/index.html"))
				_ = t.Execute(w, page)
			} else {
				fmt.Println("loginn4.1")
				page.IsLogin = false
				t := template.Must(template.ParseFiles("views/pages/user/login.html"))
				t.Execute(w, nil)
			}
		}
	}
}

//func Login(w http.ResponseWriter, r *http.Request) {
//	flag, _ := dao.IsLogin(r)
//	page, _ := dao.GetPageMovie("1")
//	page.Username = r.FormValue("username")
//	if flag {
//		page.IsLogin = true
//		t := template.Must(template.ParseFiles("views/pages/login.html"))
//		t.Execute(w, page)
//	} else {
//		username := r.PostFormValue("username")
//		password := r.PostFormValue("password")
//		admin,_:=dao.CheckAdmin(username,password)
//		if admin.ID>0{
//
//		}
//		user, _ := dao.CheckUserNameAndPassword(username, password)
//		if user.ID > 0 {
//
//			uuid := utils.CreatUUID()
//			sess := &model.Session{
//				SessionID: uuid,
//				UserName:  username,
//				UserID:    user.ID,
//			}
//
//			dao.AddSession(sess)
//			cookie := http.Cookie{
//				Name:     "user",
//				Value:    uuid,
//				HttpOnly: true,
//			}
//			http.SetCookie(w, &cookie)
//			//fmt.Println("登录成功")
//			page.IsLogin = true
//
//			t := template.Must(template.ParseFiles("views/index.html"))
//			_ = t.Execute(w, page)
//		} else {
//			page.IsLogin = false
//			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
//			t.Execute(w, nil)
//		}
//	}
//}

func Exit(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	GetPageMovie(w, r)
}

func JumpToRegister(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/pages/user/Register.html")
	t.Execute(w, nil)
}

func JumpToLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/pages/user/login.html")
	t.Execute(w, nil)
}

func SubmitComment(w http.ResponseWriter, r *http.Request) {
	moviename := r.FormValue("moviename")
	speaker := r.FormValue("username")
	comment := r.FormValue("comment")
	At := r.FormValue("At")
	time := time.Now()

	Comment := &model.Comment{
		Speaker: speaker,
		Word:    comment,
		Time:    string(time.Format("2006-01-02-15:04:05")),
		Movie:   moviename,
		At:      At,
	}
	dao.AddComment(Comment)
	GetPageMovie(w, r)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	Comment := &model.Comment{
		Word:    r.FormValue("comment"),
		Time:    r.FormValue("time"),
		Movie:   r.FormValue("movie"),
		Speaker: r.FormValue("speaker"),
	}
	fmt.Println(Comment)
	dao.DeleteComment(Comment)
	GetPageMovie(w, r)
}
