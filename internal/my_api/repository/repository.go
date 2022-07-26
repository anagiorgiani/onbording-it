package repository

import (
	"database/sql"
	"log"

	db "github.com/anagiorgiani/onbording-it/internal/my_api/domain"
)

type apiRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) db.ApiRepository {
	return &apiRepository{db: db}
}

func (repo *apiRepository) Get() (*[]db.Api, error) {

	data := []db.Api{}

	rows, err := repo.db.Query(SQLGetAll)

	if err != nil {
		return &data, err
	}
	defer rows.Close()

	for rows.Next() {
		var result db.Api
		if err := rows.Scan(
			&result.Id,
			&result.Data,
		); err != nil {
			return &data, err
		}
		data = append(data, result)
	}
	return &data, nil
}

func (r *apiRepository) Create(data string) (*db.Api, error) {

	stmt, err := r.db.Prepare(SQLCreate)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(data)
	if err != nil {
		return &db.Api{}, err
	}

	insertedId, _ := result.LastInsertId()
	myApi := db.Api{
		Id:   uint64(insertedId),
		Data: data,
	}
	return &myApi, nil
}
