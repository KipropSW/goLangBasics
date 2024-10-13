package domain

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"server/errs"
	"server/logger"
)

type ListingRepositoryDb struct {
	client *sqlx.DB
}

func (d ListingRepositoryDb) ById(id string) (*Listing, *errs.AppError) {
	var l Listing
	getCustomerById := `SELECT listing_id, name, zipcode, location, ward from houses where listing_id = ?`
	//row := d.client.QueryRow(getCustomerById, id)
	err := d.client.Get(&l, getCustomerById, id)
	//err := row.Scan(&l.Id, &l.Name, &l.Zipcode, &l.Location, &l.Ward, &l.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("Listing not found")
		}
		logger.Error("Error while scanning listing" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &l, nil

}

func (d ListingRepositoryDb) FindAll(status string) ([]Listing, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	listings := make([]Listing, 0)

	if status == "" {
		findAllSql := `SELECT listing_id, name, zipcode, location, ward, status from houses`
		//rows, err = d.client.Query(findAllSql)
		err = d.client.Select(&listings, findAllSql)
	} else {
		findAllSql := `SELECT listing_id, name, zipcode, location, ward, status from houses where status = ?`
		//rows, err = d.client.Query(findAllSql, status)
		err = d.client.Select(&listings, findAllSql, status)

	}

	if err != nil {
		logger.Error("Error while querying database" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	//err = sqlx.StructScan(rows, &listings)
	//if err != nil {
	//	logger.Error("Error while scanning listings" + err.Error())
	//	return nil, errs.NewUnexpectedError("Unexpected database error")
	//}

	//for rows.Next() {
	//	var l Listing
	//	err := rows.Scan(&l.Id, &l.Name, &l.Zipcode, &l.Location, &l.Ward, &l.Status)
	//	if err != nil {
	//		logger.Error("Error while scanning listings" + err.Error())
	//		return nil, errs.NewUnexpectedError("Unexpected database error")
	//	}
	//	listings = append(listings, l)
	//}
	return listings, nil
}

func NewListingRepositoryDb(db *sqlx.DB) ListingRepositoryDb {
	return ListingRepositoryDb{db}
}
