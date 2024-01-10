CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    shorturl VARCHAR(255) UNIQUE NOT NULL,
    longurl TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

go get .
go run .

