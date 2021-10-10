package repository

import (
	"database/sql"

	"github.com/SofaSubs/server/model"
)

type SubsRepository interface {
	FindSubs(page int, size int) ([]model.SubDB, error)
	AddSub(sub model.Sub) (bool, error)
}

type subsRepository struct {
	db *sql.DB
}

func NewSubsRepository(db *sql.DB) SubsRepository {
	return &subsRepository{
		db: db,
	}
}

func (r subsRepository) FindSubs(page int, size int) ([]model.SubDB, error) {
	var subs []model.SubDB
	rows, err := r.db.Query("select * from subs")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tempSubDB model.SubDB
		err = rows.Scan(&tempSubDB.Id, &tempSubDB.Original, &tempSubDB.Translate)

		if err != nil {
			return nil, err
		}

		subs = append(subs, tempSubDB)
	}

	return subs, nil
}

func (r subsRepository) AddSub(sub model.Sub) (bool, error) {
	if _, err := r.db.Exec("insert into subs(original, translate) values('" + sub.Original + "', '" + sub.Translate + "')"); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
