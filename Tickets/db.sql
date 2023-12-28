CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    from_location VARCHAR(30),
    to_location VARCHAR(50),
    date TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    email VARCHAR(50),
    phone VARCHAR(20),
    ticket_id INT REFERENCES tickets(id)
);
