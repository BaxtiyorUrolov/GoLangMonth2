// ticket.go fayli

package ticket

import (
	"database/sql"
	"fmt"
	"time"
)

type Ticket struct {
	ID   int
	From string
	To   string
	Date time.Time
}

func AddTicket(db *sql.DB, from, to string) (int, error) {
	newTicket := Ticket{
		From: from,
		To:   to,
		Date: time.Now(),
	}

	var ticketID int
	err := db.QueryRow("INSERT INTO tickets (from_location, to_location, date) VALUES ($1, $2, $3) RETURNING id",
		newTicket.From, newTicket.To, newTicket.Date).Scan(&ticketID)
	if err != nil {
		return 0, err
	}

	return ticketID, nil
}

func GetTicketByID(db *sql.DB, ticketID int) (*Ticket, error) {
	ticket := &Ticket{}
	err := db.QueryRow("SELECT id, from_location, to_location, date FROM tickets WHERE id = $1", ticketID).
		Scan(&ticket.ID, &ticket.From, &ticket.To, &ticket.Date)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func ListTickets(db *sql.DB) ([]Ticket, error) {
    rows, err := db.Query("SELECT id, from_location, to_location, date FROM tickets")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tickets []Ticket
    for rows.Next() {
        var ticket Ticket
        err := rows.Scan(&ticket.ID, &ticket.From, &ticket.To, &ticket.Date)
        if err != nil {
            return nil, err
        }
        tickets = append(tickets, ticket)
    }

    return tickets, nil
}

func UpdateTicketByID(db *sql.DB) {
		tickets, err := ListTickets(db)
		if err != nil {
			panic(err)
		}
	
		fmt.Println("Tickets:")
		for _, t := range tickets {
			fmt.Printf("Ticket ID: %d\tFrom: %s\tTo: %s\tDate: %s\n", t.ID, t.From, t.To, t.Date.Format("2006-01-02"))
		}
	
		if len(tickets) > 0 {
			var ticketID int
			fmt.Print("O'zgartirmoqchi bo'lgan ticketingiz ID raqamini kiriting: ")
			fmt.Scan(&ticketID)
	
			found := false
			for _, t := range tickets {
				if t.ID == ticketID {
					found = true
					break
				}
			}
			if !found {
				fmt.Println("ID topilmadi.")
				return
			}
	
			var newFrom, newTo string
			fmt.Print("Yangi manzilni kiriting 'From' : ")
			fmt.Scan(&newFrom)
			fmt.Print("Yangi manzilni kiriting 'To' : ")
			fmt.Scan(&newTo)
	
			updatedTicket := Ticket{
				From: newFrom,
				To:   newTo,
				Date: time.Now(),
			}
	
			_, err := db.Exec("UPDATE tickets SET from_location = $1, to_location = $2, date = $3 WHERE id = $4",
				updatedTicket.From, updatedTicket.To, updatedTicket.Date, ticketID)
			if err != nil {
				panic(err)
			}
	
			fmt.Println("Ticket muvaffaqiyatli yangilandi!")
		}
}

func DeleteTicket(db *sql.DB) {
	tickets, err := ListTickets(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tickets:")
	for _, t := range tickets {
		fmt.Printf("Ticket ID: %d\tFrom: %s\tTo: %s\tDate: %s\n", t.ID, t.From, t.To, t.Date.Format("2006-01-02"))
	}

	if len(tickets) > 0 {
		var ticketID int
		fmt.Print("O'chirmoqchi bo'lgan ticketingiz ID raqamini kiriting : ")
		fmt.Scan(&ticketID)

		found := false
		for _, t := range tickets {
			if t.ID == ticketID {
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Bunday ticket topilmadi.")
			return
		}

		_, err := db.Exec("DELETE FROM tickets WHERE id = $1", ticketID)
		if err != nil {
			panic(err)
		}

		fmt.Println("Ticket muvaffaqiyatli o'chirildi!")
	}
}
