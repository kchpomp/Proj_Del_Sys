package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type user struct {
	account_type string
	username     string
	password     string
	e_mail       string
	adress       string
}

//type rests struct {
//	restaurant   string
//	foodtype     string
//	time_to_wait string
//}

//type menus struct {
//	restaurant  string
//	menu        string
//	delivery    string
//	phonenumber string
//	email       string
//}

// This funtion allows to choose a restaurant in the restaurants table
func findUser(db *sql.DB, login string, password string) (bool, bool) {
	rows, err := db.Query("select * from users where username = $1", login)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		p := user{}
		err := rows.Scan(&p.account_type, &p.username, &p.password, &p.e_mail, &p.adress)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	if len(users) == 0 {
		return false, false
	} else {
		if users[0].password == password {
			return true, true
		} else {
			return true, false
		}
	}
}

//func UserInfo(db *sql.DB, login string) []user {
//	rows, err := db.Query("select * from users where username = $1", login)
//	if err != nil {
//		panic(err)
//	}
//
//	defer rows.Close()
//
//	users := []user{}
//
//	for rows.Next() {
//		p := user{}
//		err := rows.Scan(&p.account_type, &p.username, &p.password, &p.e_mail, &p.adress)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		users = append(users, p)
//	}
//	return users
//}

// This funtion allows to choose a restaurant in the restaurants table
//func findRest(db *sql.DB, restaurant string) (bool, bool) {
//	rows, err := db.Query("select * from restaurants where restaurant = $1", restaurant)
//	if err != nil {
//		panic(err)
//	}
//
//	defer rows.Close()
//
//	restaurants := []rests{}
//
//	for rows.Next() {
//		r := rests{}
//		err := rows.Scan(&r.restaurant, &r.foodtype, &r.time_to_wait)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		restaurants = append(restaurants, r)
//	}
//
//	if len(restaurants) == 0 {
//		return false, false
//	} else {
//		if restaurants[0].restaurant == restaurant {
//			return true, true
//		} else {
//			return true, false
//		}
//	}
//}

//func IndexHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
//	rows, err := db.Query("select * from productdb.Products")
//	if err != nil {
//		log.Println(err)
//	}
//	defer rows.Close()
//	products := []user{}
//
//	for rows.Next() {
//		p := user{}
//		err := rows.Scan(&p.account_type, &p.username, &p.password, &p.e_mail, &p.adress)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		products = append(products, p)
//	}
//}

//// This funtion allows to choose what you would like to choose in Ekiwoki: main course, bar or appetizers pizza if exists
//func findEkiwoki(db *sql.DB, menu string) (bool, bool) {
//	rows, err := db.Query("select * from ekiwoki where menu = $1", menu)
//	if err != nil {
//		panic(err)
//	}
//
//	defer rows.Close()
//
//	menu1 := []menus{}
//
//	for rows.Next() {
//		m := menus{}
//		err := rows.Scan(&m.restaurant, &m.menu, &m.delivery, &m.phonenumber, &m.email)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		menu1 = append(menu1, m)
//	}
//
//	if len(menu1) == 0 {
//		return false, false
//	} else {
//		if menu1[0].menu == menu {
//			return true, true
//		} else {
//			return true, false
//		}
//	}
//}
//
//// This funtion allows to choose what you would like to choose in Papa Milano: main course, bar or appetizers pizza if exists
//func findPapaMilano(db *sql.DB, menu string) (bool, bool) {
//	rows, err := db.Query("select * from papamilano where menu = $1", menu)
//	if err != nil {
//		panic(err)
//	}
//
//	defer rows.Close()
//
//	menu1 := []menus{}
//
//	for rows.Next() {
//		m := menus{}
//		err := rows.Scan(&m.restaurant, &m.menu, &m.delivery, &m.phonenumber, &m.email)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		menu1 = append(menu1, m)
//	}
//
//	if len(menu1) == 0 {
//		return false, false
//	} else {
//		if menu1[0].menu == menu {
//			return true, true
//		} else {
//			return true, false
//		}
//	}
//}
//
//// This funtion allows to choose what you would like to choose in Papa Milano: main course, bar or appetizers pizza if exists
//func findShavaWorld(db *sql.DB, menu string) (bool, bool) {
//	rows, err := db.Query("select * from shavaworld where menu = $1", menu)
//	if err != nil {
//		panic(err)
//	}
//
//	defer rows.Close()
//
//	menu1 := []menus{}
//
//	for rows.Next() {
//		m := menus{}
//		err := rows.Scan(&m.restaurant, &m.menu, &m.delivery, &m.phonenumber, &m.email)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		menu1 = append(menu1, m)
//	}
//
//	if len(menu1) == 0 {
//		return false, false
//	} else {
//		if menu1[0].menu == menu {
//			return true, true
//		} else {
//			return true, false
//		}
//	}
//}
//
//// This funtion allows to choose what you would like to choose in My Burger: main course, bar or appetizers pizza if exists
//func findMyBurger(db *sql.DB, menu string) (bool, bool) {
//	rows, err := db.Query("select * from my_burger where menu = $1", menu)
//	if err != nil {
//		panic(err)
//	}
//
//	defer rows.Close()
//
//	menu1 := []menus{}
//
//	for rows.Next() {
//		m := menus{}
//		err := rows.Scan(&m.restaurant, &m.menu, &m.delivery, &m.phonenumber, &m.email)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		menu1 = append(menu1, m)
//	}
//
//	if len(menu1) == 0 {
//		return false, false
//	} else {
//		if menu1[0].menu == menu {
//			return true, true
//		} else {
//			return true, false
//		}
//	}
//}

//const (
//	COOKIE_NAME = "sessionId"
//)

func main() {
	// Connection to the postgres Database Delivery_System
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "d19995678"
		dbname   = "Delivery_System"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Function to choose a restaurant in the html page
	choose_restaurant := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/Choose_Restaurant.html")
		fmt.Println(w)
		fmt.Println(r)
		//http.Redirect(w, r, "http://localhost:8181/Choose_Restaurant", 303)
	}
	http.HandleFunc("/choose_rest", choose_restaurant)

	// Function for the login page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/Login_Page.html")
	})

	// Function for the Home page
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/go_choose_rest.html")
	})

	// Function for logging in the html page
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

		login := r.FormValue("login")
		password := r.FormValue("password")

		logMat, passMat := findUser(db, login, password)

		if logMat {
			if passMat {
				//http.ServeFile(w, r, "htmls/Choose_Restaurant.html")
				http.Redirect(w, r, "/choose_rest", http.StatusFound)
				fmt.Println(w)
				fmt.Println(r)
				// Choose Ekiwoki Restaurant
			} else {
				fmt.Fprintf(w, "Wrong password for %s", login)
			}
		} else {
			fmt.Fprintf(w, "There is no user with login %s", login)
		}
	})

	// Choose Ekiwoki Restaurant
	http.HandleFunc("/choose_ekiwoki", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./Ekiwoki Menu Page.html")
		fmt.Println(w)
		fmt.Println(r)
		//http.Redirect(w, r, "http://localhost:8181/Ekiwoki%20Menu%20Page", 303)
	})

	// Choose Papa Milano Restaurant
	http.HandleFunc("/choose_papa_milano", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./Papa%20Milano%20Menu%20Page.html")
		http.Redirect(w, r, "http://localhost:8181/Papa%20Milano%20Menu%20Page", 303)
	})

	// Choose My Burger Restaurant
	http.HandleFunc("/choose_my_burger", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./My%20Burger%20Menu.html")
		http.Redirect(w, r, "http://localhost:8181/My%20Burger%20Menu", 303)
	})

	// Choose ShavaWorld Restaurant
	http.HandleFunc("/choose_shavaworld", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./ShavaWorld%20Menu%20Page.html")
		http.Redirect(w, r, "http://localhost:8181/ShavaWorld%20Menu%20Page", 303)
	})

	//Handler Functions Block For Ekiwoki Restaurant
	//Handler for Ekiwoki Menu
	http.HandleFunc("/choose_ekiwoki_menu", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/Ekiwoki_menu_menu_page.html")
	})
	//Handler for Ekiwoki Appetizers
	http.HandleFunc("/choose_ekiwoki_appetizers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/Ekiwoki_menu_appetizers_page.html")
	})
	//Handler for Ekiwoki Bar
	http.HandleFunc("/choose_ekiwoki_bar", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/Ekiwoki_menu_bar_page.html")
	})

	//Handler Functions Block For My Burger Restaurant
	//Handler for My Burger Menu
	http.HandleFunc("/choose_my_burger_menu", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/My_Burger_menu_menu.html")
		http.ServeFile(w, r, "htmls/My_Burger_menu_bar.html")
	})
	//Handler for My Burger Appetizers
	http.HandleFunc("/choose_my_burger_appetizers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/My_Burger_menu_appetizers.html")
	})
	//Handler for My Burger Bar
	http.HandleFunc("/choose_my_burger_bar", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/My_Burger_menu_bar.html")
	})

	//Handler Functions Block For Papa Milano Restaurant
	//Handler for Papa Milano Menu
	http.HandleFunc("/choose_papa_milano_menu", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/Papa_Milano_menu_menu.html")
	})
	//Handler for Papa Milano Appetizers
	http.HandleFunc("/choose_papa_milano_appetizers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/Papa_Milano_menu_appetizers.html")
	})
	//Handler for Papa Milano Pizza
	http.HandleFunc("/choose_papa_milano_pizza", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/Papa_Milano_menu_pizza.html")
	})
	//Handler for Papa Milano Bar
	http.HandleFunc("/choose_papa_milano_bar", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/Papa_Milano_menu_bar.html")
	})

	//Handler Functions Block For ShavaWorld Restaurant
	//Handler for ShavaWorld Menu
	http.HandleFunc("/choose_shavaworld_menu", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/ShavaWorld_menu_menu.html")
	})
	//Handler for ShavaWorld Appetizers
	http.HandleFunc("/choose_shavaworld_appetizers", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/ShavaWorld_menu_appetizers.html")
	})
	//Handler for ShavaWorld Bar
	http.HandleFunc("/choose_shavaworld_bar", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmls/ShavaWorld_menu_bar.html")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
