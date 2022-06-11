package repository

import (
	"fmt"
	"forum/models"
)

func Count(table string) (models.Count, error) {
	//SELECT COUNT(*) FROM user
	rows, err := forumDatabase.QuerryData(fmt.Sprintf("SELECT COUNT(*) FROM %s", table))
	if err != nil {
		return models.Count{}, err
	}
	count := models.Count{}
	for rows.Next() {
		rows.Scan(&count.Nb)
		if err != nil {
			return models.Count{}, err
		}
	}
	return count, nil
}
