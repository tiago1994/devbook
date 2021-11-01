package repository

import (
	"database/sql"
	"devbook/src/models"
)

type Publication struct {
	db *sql.DB
}

func NewPublicationRepository(db *sql.DB) *Publication {
	return &Publication{db}
}

func (repository Publication) Create(publication models.Publication) (uint64, error) {
	statment, statmentError := repository.db.Prepare("INSERT INTO publications (title, content, created_id) VALUES (?, ?, ?)")
	if statmentError != nil {
		return 0, statmentError
	}

	defer statment.Close()
	result, resultError := statment.Exec(publication.Title, publication.Content, publication.CreatedID)
	if resultError != nil {
		return 0, resultError
	}

	lastId, errorLastId := result.LastInsertId()
	if errorLastId != nil {
		return 0, errorLastId
	}

	return uint64(lastId), nil
}

func (repository Publication) GetBy(publicationID uint64) (models.Publication, error) {
	rows, requestError := repository.db.Query("SELECT p.*, u.nick FROM publications p INNER JOIN users u ON u.id = p.created_id WHERE p.id = ?", publicationID)
	if requestError != nil {
		return models.Publication{}, requestError
	}

	defer rows.Close()
	var publication models.Publication
	if rows.Next() {
		if requestError = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.CreatedID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.CreatedNick,
		); requestError != nil {
			return models.Publication{}, requestError
		}
	}

	return publication, nil
}

func (repository Publication) GetAll(userID uint64) ([]models.Publication, error) {
	rows, requestError := repository.db.Query(`
		SELECT DISTINCT p.*, u.nick FROM publications p
		INNER JOIN users u ON u.id = p.created_id
		INNER JOIN followers s on p.created_id = s.user_id
		WHERE u.id = ? OR s.follower_id = ? ORDER BY 1 DESC`, userID, userID)
	if requestError != nil {
		return nil, requestError
	}

	defer rows.Close()
	var publications []models.Publication
	for rows.Next() {
		var publication models.Publication
		if requestError = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.CreatedID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.CreatedNick,
		); requestError != nil {
			return nil, requestError
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publication) UpdatePublication(publicationID uint64, publication models.Publication) error {
	statment, statmentError := repository.db.Prepare("UPDATE publications SET title = ?, content = ? WHERE id = ?")
	if statmentError != nil {
		return statmentError
	}

	defer statment.Close()
	_, resultError := statment.Exec(publication.Title, publication.Content, publicationID)
	if resultError != nil {
		return resultError
	}

	return nil
}

func (repository Publication) DeletePublication(publicationID uint64) error {
	statment, statmentError := repository.db.Prepare("DELETE FROM publications WHERE id = ?")
	if statmentError != nil {
		return statmentError
	}

	defer statment.Close()
	_, resultError := statment.Exec(publicationID)
	if resultError != nil {
		return resultError
	}

	return nil
}

func (repository Publication) GetByUser(userID uint64) ([]models.Publication, error) {
	rows, requestError := repository.db.Query(`
		SELECT p.*, u.nick FROM publications p
		JOIN users u ON u.id = p.created_id
		WHERE p.created_id = ?`, userID)
	if requestError != nil {
		return nil, requestError
	}

	defer rows.Close()
	var publications []models.Publication
	for rows.Next() {
		var publication models.Publication
		if requestError = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.CreatedID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.CreatedNick,
		); requestError != nil {
			return nil, requestError
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publication) Like(publicationID uint64) error {
	statment, statmentError := repository.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = ?")
	if statmentError != nil {
		return statmentError
	}

	defer statment.Close()
	_, resultError := statment.Exec(publicationID)
	if resultError != nil {
		return resultError
	}

	return nil
}

func (repository Publication) Dislike(publicationID uint64) error {
	statment, statmentError := repository.db.Prepare("UPDATE publications SET likes = CASE WHEN likes > 0 THEN likes - 1 ELSE 0 END WHERE id = ?")
	if statmentError != nil {
		return statmentError
	}

	defer statment.Close()
	_, resultError := statment.Exec(publicationID)
	if resultError != nil {
		return resultError
	}

	return nil
}
