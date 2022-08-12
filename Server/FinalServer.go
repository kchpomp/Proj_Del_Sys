package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	_ "github.com/martini-contrib/render"
	"html/template"
	"net/http"
	"strings"
)

var (
	key = []byte("1FOXCKBJU59WKCVV")
	//Глобальная переменная названия ресторана
	rest_name_for_URL string
	order_list        []string
	result            int
	table_name        string
	menu_name         string
	position          string
)

//Функция для создания сервера
func newServer(db_url string) *Server {
	s := &Server{
		router: mux.NewRouter(),
		store:  *sessions.NewCookieStore(key),
		conn:   connect_db(db_url),
	}

	s.createRouter()

	return s
}

//Функция для соединения с базой данных
func connect_db(url string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		fmt.Println("config Database Fail")
		fmt.Print(err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)

	return conn
}

func (s Server) MainPage(w http.ResponseWriter, r *http.Request) {
	session, _ := s.store.Get(r, "cookie-name")

	// Проверка аутентификации
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		//Если нет, то переход на страницу входа
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
		//Если да, то переход на домашнюю страницу
	} else {
		username := session.Values["username"].(string)
		//Ищем запись о пользователе в таблице users
		rows, err := s.conn.Query(context.Background(), "select * from users where username = $1", username)
		if err != nil {
			panic(err)
		}

		usrs := []UserInfo{}
		var usr UserInfo

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&usr.account_type, &usr.username, &usr.password, &usr.e_mail, &usr.adress)
			if err != nil {
				fmt.Println(err)
				continue
			}
			usrs = append(usrs, usr)
		}

		//Если нашли запись о пользователе, то, в таком случае
		//Выполняем переадресацию на нужную домашнюю страницу в зависимости от типа аккаунта пользователя
		if usrs[0].account_type == "client" {
			http.Redirect(w, r, "/client_homepage", http.StatusSeeOther)
		} else if usrs[0].account_type == "administrator" {
			http.Redirect(w, r, "/admin_homepage", http.StatusSeeOther)
		} else if usrs[0].account_type == "courier" {
			http.Redirect(w, r, "/courier_homepage", http.StatusSeeOther)
		}
	}
}

// Функция "обслуживания" страницы аутентификации
func (s Server) LoginPage(w http.ResponseWriter, r *http.Request) {
	html_path := "C:/Users/user/go/Delivery_System/htmls/Login_Page.html"
	t, err := template.ParseFiles(html_path)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

// Функция для страницы аутентификации, с помощью findUser выполяет поиск в базе данных и если такой пользователь есть,
//то происходит редирект на страницу выбора ресторана
func (s Server) LoginProcedure(w http.ResponseWriter, r *http.Request) {
	session, _ := s.store.Get(r, "cookie-name")

	// Формируются переменные из полей login и password
	login := r.FormValue("login")
	password := r.FormValue("password")

	//Используется функция поиска юзера в базе данных
	logMat, passMat := s.findUser(login, password)

	// Если логин и пароль совпадают, то сохраняются значения сессии
	// и происходит переход на MainPage
	// в результате, так как сессия сохранена, то в "/" должен произойти переход на домашнюю страницу
	if logMat {
		if passMat {
			session.Values["username"] = login
			session.Values["authenticated"] = true
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Fprintf(w, "Wrong password for %s", login)
		}
	} else {
		fmt.Fprintf(w, "There is no user with login %s", login)
	}
}

// Функция для обслуживания процедуры логаута
func (s Server) LogoutProcedure(w http.ResponseWriter, r *http.Request) {
	session, _ := s.store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Функция "обслуживания" домашней страницы пользователя
func (s Server) ClientHomePage(w http.ResponseWriter, r *http.Request) {
	html_path := "C:/Users/user/go/Delivery_System/htmls/Client_Homepage.html"
	t, err := template.ParseFiles(html_path)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

// Функция "обслуживания" страницы выбора ресторана
func (s Server) ChooseRestaurantPage(w http.ResponseWriter, r *http.Request) {
	html_path := "C:/Users/user/go/Delivery_System/htmls/Choose_Restaurant.html"
	t, err := template.ParseFiles(html_path)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

// Функция "обслуживания" страницы информации о заказе для клиента
func (s Server) OrderInfoPageClient(w http.ResponseWriter, r *http.Request) {
	Order_lst := strings.Join(order_list, ", ")
	Price := result
	fmt.Println("Order final list is:", Order_lst)
	fmt.Println("Order final price is:", Price)
	if len(order_list) != 0 {
		Show_order_info := ShowOrder{
			Order_list: Order_lst,
			Price:      Price,
		}
		fmt.Println("Structure is: ", Show_order_info)
		html_path := "C:/Users/user/go/Delivery_System/htmls/Order_Info_User.html"
		t, err := template.ParseFiles(html_path)
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, Show_order_info)
	} else {
		empty := "Empty"
		Show_order_info := ShowOrder{
			Order_list: empty,
			Price:      0,
		}
		fmt.Println("Structure is: ", Show_order_info)
		html_path := "C:/Users/user/go/Delivery_System/htmls/Order_Info_User.html"
		t, err := template.ParseFiles(html_path)
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, Show_order_info)
	}
}

// Функция переадресации на страницу ресторана
func (s Server) RedirectToRestaurant(w http.ResponseWriter, r *http.Request) {

	rest_name_for_URL = r.FormValue("restaurant")

	html_path := "C:/Users/user/go/Delivery_System/htmls/" + rest_name_for_URL + " Menu Page.html"

	t, err := template.ParseFiles(html_path)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

// Функция переадресации на страницу меню ресторана
func (s Server) RedirectToMenu(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Params are: ", r.URL.Query())

	menu_name = r.FormValue("menu")

	html_path := "C:/Users/user/go/Delivery_System/htmls/" + rest_name_for_URL + " menu " + menu_name + ".html"

	t, err := template.ParseFiles(html_path)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

// Функция "обслуживания" домашней страницы пользователя
// Данная функция показывает на домашней странице пользователя информацию о пользователе
func (s Server) ClientInfo(w http.ResponseWriter, r *http.Request) {
	session, _ := s.store.Get(r, "cookie-name")
	html_path := "C:/Users/user/go/Delivery_System/htmls/Client_Homepage.html"
	//http.ServeFile(w, r, html_path)
	t, err := template.ParseFiles(html_path)
	if err != nil {
		fmt.Println(err)
	}

	username := session.Values["username"].(string)
	name, e_mail, adress := s.ShowInfo(username)

	show_client_info := ShowInfo{
		username: name,
		e_mail:   e_mail,
		adress:   adress,
	}

	t.Execute(w, show_client_info)
}

//Функция для выбора блюда из меню ресторана
//Данная функция формирует значения, затем с помощью ChooseDish обращается к базе данных и выбирает заданные блюда
//После чего с помощью ShowOrderInfo изменяется значение структуры ShowOrder
//В результате на странице информации о заказе либо будет показан список заказа и его цена, либо будет пусто
func (s Server) SelectDishFromMenu(w http.ResponseWriter, r *http.Request) {
	table_name = strings.ToLower(strings.TrimSpace(rest_name_for_URL))
	position = r.URL.Query().Get("position")

	html_path := "C:/Users/user/go/Delivery_System/htmls/" + rest_name_for_URL + " menu " + menu_name + ".html"
	s.ChooseDish(table_name, menu_name, position)
	t, err := template.ParseFiles(html_path)
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

//Функция для показывания информации о курьере
func (s Server) CourierInfo(w http.ResponseWriter, r *http.Request) {
	session, _ := s.store.Get(r, "cookie-name")

	courier_name := session.Values["username"]

	html_path := "C:/Users/user/go/Delivery_System/htmls/Courier_Homepage.html"
	http.ServeFile(w, r, html_path)

	t, err := template.ParseFiles(html_path)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := s.conn.Query(context.Background(), "select courier_id, first_name, last_name, phone_number, e_mail from couriers where courier_id = $1", courier_name)
	if err != nil {
		panic(err)
	}

	courier := []CourierInfo{}

	defer rows.Close()

	for rows.Next() {
		p := CourierInfo{}
		err := rows.Scan(&p.courier_id, &p.first_name, &p.last_name, &p.phone_number, &p.e_mail)
		if err != nil {
			fmt.Println(err)
			continue
		}
		courier = append(courier, p)
	}
	t.Execute(w, courier)
}

func (s Server) createRouter() {
	s.router.HandleFunc("/", s.MainPage).Methods("GET")

	//Страница аутентификации
	s.router.HandleFunc("/login", s.LoginPage).Methods("GET")
	s.router.HandleFunc("/login", s.LoginProcedure).Methods("POST")
	s.router.HandleFunc("/logout", s.LogoutProcedure).Methods("GET")

	//Домашняя страница пользователя
	s.router.HandleFunc("/client_homepage", s.ClientHomePage).Methods("GET")
	s.router.HandleFunc("/client_homepage", s.ClientHomePage).Methods("POST")

	//Страница выбора ресторана
	s.router.HandleFunc("/choose_restaurant", s.ChooseRestaurantPage).Methods("GET")
	s.router.HandleFunc("/rest_page", s.RedirectToRestaurant).Methods("GET")

	//Страница информации о заказе для пользователя
	s.router.HandleFunc("/show_order_for_client", s.OrderInfoPageClient).Methods("GET")

	//Страница выбора одного из меню
	s.router.HandleFunc("/choose_menu", s.RedirectToMenu).Methods("GET")
	//Добавление блюда в заказ
	s.router.HandleFunc("/choose_dish", s.SelectDishFromMenu).Methods("GET")

	//Страница информации о курьере
	s.router.HandleFunc("/courier_info", s.CourierInfo).Methods(("POST"))

	//
	//s.router.HandleFunc("/register", s.RegisterPage).Methods(("GET"))
	//s.router.HandleFunc("/register", s.RegisterProcedure).Methods(("POST"))
	//
	//s.router.HandleFunc("/test", s.Test).Methods(("GET"))
	//s.router.HandleFunc("/test", s.TestCheck).Methods(("POST"))
	//
	//s.router.HandleFunc("/task", s.Task).Methods(("GET"))
	//s.router.HandleFunc("/task", s.TaskSubmit).Methods(("POST"))
	//
	//s.router.HandleFunc("/edittask", s.EditTask).Methods(("GET"))
	//s.router.HandleFunc("/edittask", s.SaveTask).Methods(("POST"))
	//s.router.HandleFunc("/deletetask", s.DeleteTask).Methods(("GET"))
	//
	//s.router.HandleFunc("/userinfo", s.UserInfo).Methods(("GET"))
	//
	//s.router.HandleFunc("/delete_account", s.DeleteAccount).Methods(("GET"))
	//
	//s.router.HandleFunc("/translate", s.Translate).Methods(("POST"))
	//
	//s.router.HandleFunc("/profile", s.ProfilePage).Methods(("GET"))
}
