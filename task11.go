package main

import "fmt"

/* 11. Реализовать пересечение двух неупорядоченных множеств */
func main() {
	// Сделаем на примере двух пользователей и их общих друзей
	user1 := []string{"kate", "jack", "john", "sam"}
	user2 := []string{"kolya", "kate", "olaf", "taker"}
	m := sameFriends(user1, user2)
	for k, v := range m {
		if v {
			fmt.Println("Общие друзья у данных пользователей:")
			fmt.Printf("%s\n", k)
		}
	}
}

func sameFriends(user1, user2 []string) map[string]bool {
	m := make(map[string]bool)
	result := make(map[string]bool)

	// Заполняем первое множество
	for _, v := range user1 {
		m[v] = true
	}

	// Проверяем второе множество
	for _, v := range user2 {
		if m[v] {
			result[v] = true
		}
	}
	return result
}
