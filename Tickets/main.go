package main

import (
	"database/sql"
	"fmt"
	"homework/ticket"
	"homework/users"
	"homework/report"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "gobaza"
	password = "0208"
	dbname   = "goproject"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var bulim int
	fmt.Print("Bo'limni tanlang: ")
	fmt.Scan(&bulim)

	switch bulim{
	case 1:

		var request int
		fmt.Print("Qanday amal bajarishni xohlaysiz ? ")
		fmt.Scan(&request)
	
		switch request {
		case 1:

			userID, err := users.UserAdd(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("New user ID:", userID)

case 2:
	users.GetUserByID(db)

case 3:
	users.ListUsers(db)

case 4:
	users.UpdateUserByID(db)

case 5:
	users.DeleteUserByID(db)
}

	case 2:

	var request int

	fmt.Print("Qanday amal bajarishni xohlaysiz ? ")
	fmt.Scan(&request)

	switch request {
	case 1:
		fmt.Println("Foydalanuvchilarni ro'yxatga o'tkazish")
		usersList, err := users.ListUsers(db)
		if err != nil {
			fmt.Println("Error listing users:", err)
			return
		}

		fmt.Println("Foydalanuvchilar:")
		for _, u := range usersList {
			fmt.Printf("%d\t%s\t%s\t%s\t%s\t%d\n", u.ID, u.FirstName, u.LastName, u.Email, u.Phone, u.TicketID)
		}
	case 2:
		var from, to string
		fmt.Print("Qayerdan jo'namoqchisiz ? ")
		fmt.Scan(&from )

		fmt.Print("Qayerga bormoqchisiz ? ")
		fmt.Scan(&to)

		ticketID, err := ticket.AddTicket(db, from, to)
		if err != nil {
			fmt.Println("Error adding ticket:", err)
			return
		}
		fmt.Println("Yangi tiket ID:", ticketID)

	case 3:
		var ticketID int
		fmt.Print("Tiket ID ni kiriting: ")
		fmt.Scan(&ticketID)
	
		ticketInfo, err := ticket.GetTicketByID(db, ticketID)
		if err != nil {
			fmt.Println("Error getting ticket information:", err)
			return
		}
	
		fmt.Printf("Ticket ID: %d\tFrom: %s\tTo: %s\tDate: %s\n", ticketInfo.ID, ticketInfo.From, ticketInfo.To, ticketInfo.Date.Format("2006-01-02"))

	case 4:
        tickets, err := ticket.ListTickets(db)
        if err != nil {
            fmt.Println("Error listing tickets:", err)
            return
        }

        fmt.Println("Tickets:")
        for _, t := range tickets {
            fmt.Printf("Ticket ID: %d\tFrom: %s\tTo: %s\tDate: %s\n", t.ID, t.From, t.To, t.Date.Format("2006-01-02"))
        }

	case 5:
		ticket.UpdateTicketByID(db)

	case 6:
		ticket.DeleteTicket(db)
	
	default:
		fmt.Println("Noma'lum buyruq.")
	}

	case 3:

		var from, to string
	fmt.Print("Manzilni kiriting 'From': ")
	fmt.Scan(&from)
	fmt.Print("Manzilni kiriting 'To': ")
	fmt.Scan(&to)

	report.GenerateReport(db, from, to)
	}	
}
