package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func AddCustomer(customer Customer) error {
	var err error

	if _, err = SQL.Exec("INSERT INTO customer (customer_id, store_id, first_name, last_name, email, address_id, activebool, create_date, last_update, active) VALUES (DEFAULT, 2, $1, $2, $3, 595, true, DEFAULT, DEFAULT, 1)",
		customer.Firstname, customer.Lastname, customer.Email); err != nil {
		return err
	}
	return err
}

func ListCustumers(query string) ([]Customer, error) {
	var err error
	var result []Customer
	var customer Customer
	var rows *sql.Rows

	if rows, err = SQL.Query("SELECT customer_id, first_name, last_name, email FROM customer WHERE "+
		"(first_name like $1) OR (last_name like $1) OR (email like $1)ORDER BY customer_id", "%"+query+"%"); err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&customer.ID, &customer.Firstname, &customer.Lastname, &customer.Email); err != nil {
			return nil, err
		}
		result = append(result, customer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, err
}

func ListFilms(query string) ([]Film, error) {
	var err error
	var result []Film
	var film Film
	var rows *sql.Rows

	if rows, err = SQL.Query("SELECT film_id, title, description, release_year, rental_rate, length FROM film WHERE title like '%" +
		query + "%' ORDER BY film_id"); err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&film.ID, &film.Title, &film.Description, &film.Year, &film.Rate, &film.Length); err != nil {
			return nil, err
		}
		result = append(result, film)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, err
}

func CheckUser(email string, password string) (User, error) {
	result := User{}

	query := fmt.Sprintf("SELECT staff_id, first_name, last_name, email,job_title, password, picture "+
		"FROM staff WHERE email = '%s' and password = '%s' LIMIT 1", email, password)
	if err := SQL.QueryRow(query).Scan(
		&result.ID, &result.Firstname, &result.Lastname, &result.Email, &result.JobTitle, &result.Password, &result.Image); err != nil {
		return User{}, err
	}

	return result, nil
}
