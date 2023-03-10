package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(connect string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), connect)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

func (s *Storage) Tasks(taskID, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
			($1 = 0 OR id = $1) AND
			($2 = 0 OR author_id = $2)
		ORDER BY id;
	`,
		taskID,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content)
		VALUES ($1, $2) RETURNING id;
		`,
		t.Title,
		t.Content,
	).Scan(&id)
	return id, err
}

func (s *Storage) GetListTasksByAuthor(authorId int) ([]Task, error) {
	query := "SELECT * FROM tasks WHERE author_id = $1"
	rows, err := s.db.Query(context.Background(), query, authorId)
	var tasks []Task
	defer rows.Close()
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) GetListTasksByAssignedId(assigned int) ([]Task, error) {
	query := "SELECT * FROM tasks WHERE assigned_id = $1"
	rows, err := s.db.Query(context.Background(), query, assigned)
	var tasks []Task
	defer rows.Close()
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) UpdateTask(title, content string) error {
	fmt.Scan(&title, &content)
	query := "SET title = $2, content = $3 WHERE id = $1"
	if _, err := s.db.Exec(context.Background(), query, title, content); err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	if _, err := s.db.Exec(context.Background(), query, id); err != nil {
		return err
	}
	return nil
}
