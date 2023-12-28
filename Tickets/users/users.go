package users

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	TicketID  int
}

func UserAdd(db *sql.DB) (int, error) {
	var firstName, lastName, email, phone string
	var ticketID int

	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter email: ")
	fmt.Scan(&email)
	fmt.Print("Enter phone: ")
	fmt.Scan(&phone)
	fmt.Print("Enter ticket ID: ")
	fmt.Scan(&ticketID)

	newUser := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		TicketID:  ticketID,
	}

	var userID int
	err := db.QueryRow("INSERT INTO users (first_name, last_name, email, phone, ticket_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		newUser.FirstName, newUser.LastName, newUser.Email, newUser.Phone, newUser.TicketID).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("error adding user: %v", err)
	}

	return userID, nil
}

func GetUserByID(db *sql.DB) {
	var userID int
	fmt.Print("User ID raqamini kiriting: ")
	fmt.Scan(&userID)

	user := &User{}
	err := db.QueryRow("SELECT id, first_name, last_name, email, phone, ticket_id FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.TicketID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User ID: %d\tFirst Name: %s\tLast Name: %s\tEmail: %s\tPhone: %s\tTicket ID: %d\n",
		user.ID, user.FirstName, user.LastName, user.Email, user.Phone, user.TicketID)
}

func ListUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, first_name, last_name, email, phone, ticket_id FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.TicketID)
		if err != nil {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
		users = append(users, user)
	}

	fmt.Println("Users:")
	for _, u := range users {
		fmt.Printf("User ID: %d\tFirst Name: %s\tLast Name: %s\tEmail: %s\tPhone: %s\tTicket ID: %d\n",
			u.ID, u.FirstName, u.LastName, u.Email, u.Phone, u.TicketID)
	}

	return users, nil
}


func UpdateUserByID(db *sql.DB) {

	usersList, err := ListUsers(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Users:")
	for _, u := range usersList {
		fmt.Printf("User ID: %d\tFirst Name: %s\tLast Name: %s\tEmail: %s\tPhone: %s\tTicket ID: %d\n",
			u.ID, u.FirstName, u.LastName, u.Email, u.Phone, u.TicketID)
	}

	if len(usersList) > 0 {
		var userID int
		fmt.Print("O'zgartirmochi bo'lgan useringiz ID raqamini kiriting : ")
		fmt.Scan(&userID)

		found := false
		for _, u := range usersList {
			if u.ID == userID {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Bunday Id topilmadi.")
			return
		}

		var newFirstName, newLastName, newEmail, newPhone string
		var newTicketID int

		fmt.Print("Enter the new First Name: ")
		fmt.Scan(&newFirstName)
		fmt.Print("Enter the new Last Name: ")
		fmt.Scan(&newLastName)
		fmt.Print("Enter the new Email: ")
		fmt.Scan(&newEmail)
		fmt.Print("Enter the new Phone: ")
		fmt.Scan(&newPhone)
		fmt.Print("Enter the new Ticket ID: ")
		fmt.Scan(&newTicketID)

		updatedUser := User{
			FirstName: newFirstName,
			LastName:  newLastName,
			Email:     newEmail,
			Phone:     newPhone,
			TicketID:  newTicketID,
		}

		_, err := db.Exec("UPDATE users SET first_name = $1, last_name = $2, email = $3, phone = $4, ticket_id = $5 WHERE id = $6",
			updatedUser.FirstName, updatedUser.LastName, updatedUser.Email, updatedUser.Phone, updatedUser.TicketID, userID)
		if err != nil {
			panic(err)
		}

		fmt.Println("User muvaffaqiyatli yangilandi!")
	}
}

func DeleteUserByID(db *sql.DB) {
	usersList, err := ListUsers(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Users:")
	for _, u := range usersList {
		fmt.Printf("User ID: %d\tFirst Name: %s\tLast Name: %s\tEmail: %s\tPhone: %s\tTicket ID: %d\n",
			u.ID, u.FirstName, u.LastName, u.Email, u.Phone, u.TicketID)
	}

	if len(usersList) > 0 {
		var userID int
		fmt.Print("Enter the User ID you want to delete: ")
		fmt.Scan(&userID)

		found := false
		for _, u := range usersList {
			if u.ID == userID {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("User topilmadi.")
			return
		}

		_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
		if err != nil {
			panic(err)
		}

		fmt.Println("User muvaffaqiyatli o'chirildi!")
	}
}
