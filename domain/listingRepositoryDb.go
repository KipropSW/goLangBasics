package domain

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"server/errs"
	"time"
)

type ListingRepositoryDb struct {
	client *sql.DB
}

func (d ListingRepositoryDb) ById(id string) (*Listing, *errs.AppError) {
	getCustomerById := `SELECT listing_id, name, zipcode, location, ward from houses where listing_id = ?`
	row := d.client.QueryRow(getCustomerById, id)
	var l Listing
	err := row.Scan(&l.Id, &l.Name, &l.Zipcode, &l.Location, &l.Ward, &l.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("Listing not found")
		}
		log.Println("Error while scanning listing" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &l, nil

}

func (d ListingRepositoryDb) FindAll(status string) ([]Listing, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := `SELECT listing_id, name, zipcode, location, ward, status from houses`
		rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := `SELECT listing_id, name, zipcode, location, ward, status from houses where status = ?`
		rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		log.Println("Error while querying database" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	listings := make([]Listing, 0)
	for rows.Next() {
		var l Listing
		err := rows.Scan(&l.Id, &l.Name, &l.Zipcode, &l.Location, &l.Ward, &l.Status)
		if err != nil {
			log.Println("Error while scanning listings" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		listings = append(listings, l)
	}
	return listings, nil
}

func NewListingRepositoryDb() ListingRepositoryDb {
	client, err := sql.Open("mysql", "root:Silot777@@tcp(localhost:3306)/listing?charset=utf8")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return ListingRepositoryDb{client}
}
