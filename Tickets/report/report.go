package report

import (
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"
)

func GenerateReport(db *sql.DB, from, to string) {
	query := `
		SELECT t.id AS trip_id, t.from_location, t.to_location, t.date,
			   u.id AS user_id, u.first_name, u.last_name, u.email, u.phone
		FROM tickets t
		JOIN users u ON t.id = u.ticket_id
		WHERE t.from_location = $1 AND t.to_location = $2
		ORDER BY t.date
	`

	rows, err := db.Query(query, from, to)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	defer w.Flush()

	fmt.Fprintf(w, "Ticket ID\tFrom\tTo\tDate\tUser ID\tFirst Name\tLast Name\tEmail\tPhone\n")

	for rows.Next() {
		var ticketID, userID int
		var fromLocation, toLocation, firstName, lastName, email, phone string
		var date string

		err := rows.Scan(&ticketID, &fromLocation, &toLocation, &date, &userID, &firstName, &lastName, &email, &phone)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%d\t%s\t%s\t%s\t%s\n", ticketID, fromLocation, toLocation, date, userID, firstName, lastName, email, phone)
	}
}

