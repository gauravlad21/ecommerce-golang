package dbhelper

import (
	"context"
	"database/sql"
	"fmt"

	userAuthCommon "github.com/gauravlad21/ecommerce-golang/user-auth/common"
	"github.com/gauravlad21/ecommerce-golang/user-auth/dbhelper/sqlc/dbsqlc"
	_ "github.com/lib/pq"
)

type DbOperationsIF interface {
	Exec(ctx context.Context, query string) error
	CloseDb(ctx context.Context) error

	InsertUser(ctx context.Context, req *userAuthCommon.UserAuthBody, tx ...*sql.Tx) (id int32, err error)
	GetUser(ctx context.Context, email string, tx ...*sql.Tx) (usr *userAuthCommon.User, err error)
	GetUserById(ctx context.Context, id int32, tx ...*sql.Tx) (usr *userAuthCommon.User, err error)
}

type DbOps struct {
	DB     *sql.DB
	DbSqlc *dbsqlc.Queries
}

func New(db *sql.DB) DbOperationsIF {
	return &DbOps{DbSqlc: dbsqlc.New(db), DB: db}
}

func (dbOps *DbOps) Exec(ctx context.Context, query string) error {
	_, err := dbOps.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (dbOps *DbOps) CloseDb(ctx context.Context) error {
	return dbOps.DB.Close()
}

// start from below
func (dbOps *DbOps) InsertUser(ctx context.Context, req *userAuthCommon.UserAuthBody, tx ...*sql.Tx) (id int32, err error) {
	args := dbsqlc.InsertUserParams{Email: req.Email, Password: req.Password}
	id, err = GetSqlcQuery(dbOps.DbSqlc, tx...).InsertUser(ctx, args)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (dbOps *DbOps) GetUser(ctx context.Context, email string, tx ...*sql.Tx) (usr *userAuthCommon.User, err error) {
	users, err := GetSqlcQuery(dbOps.DbSqlc, tx...).GetUser(ctx, email)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no user found with email: %v", email)
	}
	usr = &userAuthCommon.User{ID: users[0].ID, Email: users[0].Email, Password: users[0].Password}
	return usr, nil
}

func (dbOps *DbOps) GetUserById(ctx context.Context, id int32, tx ...*sql.Tx) (usr *userAuthCommon.User, err error) {
	users, err := GetSqlcQuery(dbOps.DbSqlc, tx...).GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("no user found with id: %v", id)
	}
	usr = &userAuthCommon.User{ID: users[0].ID, Email: users[0].Email, Password: users[0].Password}
	return usr, nil
}
