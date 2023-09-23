package repository

import (
	"context"
	"database/sql"
	"go_project_template/configs/db"
	"go_project_template/internal/user/model"
)

type IUserRepository interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id int64) (model.User, error)
	AddNewUser(ctx context.Context, newUser model.User) (int64, error)
	UpdateUser(ctx context.Context, user model.User) (int64, error)
	DeleteUser(ctx context.Context, id int64) error
}

type UserRepository struct {
	db db.DBInterface
}

func NewUserRepository(db db.DBInterface) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Query Implementation

func (q *UserRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User

	sqlStatement := `
	SELECT * FROM public.users
	`

	// Querying
	rows, err := q.db.QueryContext(ctx, sqlStatement)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User

		// read data
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (q *UserRepository) GetUserById(ctx context.Context, id int64) (model.User, error) {
	var user model.User

	queryStatement := `
	SELECT
		*
	FROM public.users
	WHERE id=$1
	`
	err := q.db.QueryRowContext(ctx, queryStatement, id).Scan(
		&user.ID,
		&user.Name,
	)

	if err != nil {
		return user, sql.ErrNoRows
	}

	return user, nil
}

func (q *UserRepository) AddNewUser(ctx context.Context, newUser model.User) (int64, error) {

	var newId int64

	sqlStatement := `
	INSERT INTO public.users (name) VALUES ($1) RETURNING id
	`

	err := q.db.QueryRowContext(ctx, sqlStatement, newUser.Name).Scan(&newId)

	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (q *UserRepository) UpdateUser(ctx context.Context, user model.User) (int64, error) {

	sqlStatement := `
	UPDATE public.users SET name=$2 WHERE id=$1
	`

	res, err := q.db.ExecContext(ctx, sqlStatement, user.ID, user.Name)

	if err != nil {
		return 0, err
	}

	rowsAffeced, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}
	return rowsAffeced, nil
}

func (q *UserRepository) DeleteUser(ctx context.Context, id int64) error {

	// create sql statement to delete user from database
	sqlStatement := `DELETE FROM public.users WHERE id=$1`

	// execute sql statement
	res, err := q.db.ExecContext(ctx, sqlStatement, id)

	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected < 1 {
		return sql.ErrNoRows
	}
	return nil
}
