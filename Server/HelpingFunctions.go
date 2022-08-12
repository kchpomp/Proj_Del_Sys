package main

import (
	"context"
	"fmt"
	"strconv"
)

// Функция для поиска юзеров в таблице юзеров, возвращает два bool значения для проверки логина
func (s Server) findUser(username string, password string) (bool, bool) {
	rows, err := s.conn.Query(context.Background(), "select * from users where username = $1", username)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	users := []UserInfo{}

	for rows.Next() {
		p := UserInfo{}
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

//Функция, позволяющая вытащить из базы данных информацию о пользователе
func (s Server) ShowInfo(username string) (string, string, string) {
	rows, err := s.conn.Query(context.Background(), "select * from users where username = $1", username)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	user_data := []ShowInfo{}
	for rows.Next() {
		p := ShowInfo{}
		err := rows.Scan(&p.username, &p.e_mail, &p.adress)
		if err != nil {
			fmt.Println(err)
			continue
		}
		user_data = append(user_data, p)
	}

	usname := user_data[0].username
	e_mail := user_data[0].e_mail
	adress := user_data[0].adress

	return usname, e_mail, adress
}

// Функция для выбора блюда из базы данных и формирования цены заказа
// Позволяет выбрать из базы данных позицию, добавить в список и сформировать общую цену заказа
func (s Server) ChooseDish(table_name string, menu_name string, position string) ([]string, int) {

	//Делаем запрос
	rows, err := s.conn.Query(context.Background(), "select position, price from "+table_name+"_"+menu_name+" where position = $1", position)
	if err != nil {
		panic(err)
	}

	//Отложенный запрос
	defer rows.Close()

	//Создаем срез структур
	order := []CreateOrder{}
	for rows.Next() {
		p := CreateOrder{}
		err := rows.Scan(&p.position, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		order = append(order, p)
		fmt.Println(order)
	}
	order_list = append(order_list, order[0].position)
	a, err := strconv.Atoi(order[0].price)
	if err != nil {
		panic(err)
	} else {
		result += a
	}

	////Задаем общую цену заказа и список заказа
	//order_list = []string{}
	//result = 0
	//for _, i := range order {
	//	j, err := strconv.Atoi(i.price)
	//	if err != nil {
	//		panic(err)
	//	} else {
	//		result += j
	//	}
	//	order_list = append(order_list, i.position)
	//}
	fmt.Println(order_list)
	fmt.Println(result)
	return order_list, result
}
