package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
)

//Структура для сервера
type Server struct {
	store  sessions.CookieStore
	conn   *pgxpool.Pool
	router *mux.Router
}

//Структура для юзера
type UserInfo struct {
	account_type string
	username     string
	password     string
	e_mail       string
	adress       string
}

//Структура для инфы о юзере
type ShowInfo struct {
	username string
	e_mail   string
	adress   string
}

//Структура для курьера
type CourierInfo struct {
	courier_id   string
	first_name   string
	last_name    string
	phone_number string
	e_mail       string
}

//Структура для отображения заказа
type ShowOrder struct {
	Order_list string
	Price      int
}

//Структура для добавления нового ресторана
type AddRest struct {
	Restaurant   string
	Foodtype     string
	time_to_wait string
}

//Структура для добавления нового блюда
type AddDish struct {
	TableName string
	Menu      string
	Position  string
	Price     string
}

//Структура для создания заказа
type CreateOrder struct {
	menu     string
	position string
	price    string
}
