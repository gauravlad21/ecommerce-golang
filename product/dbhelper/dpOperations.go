package dbhelper

import (
	"context"
	"database/sql"
	"fmt"

	productCommon "github.com/gauravlad21/ecommerce-golang/product/common"
	"github.com/gauravlad21/ecommerce-golang/product/dbhelper/sqlc/dbsqlc"
	_ "github.com/lib/pq"
)

type DbOperationsIF interface {
	Exec(ctx context.Context, query string) error
	CloseDb(ctx context.Context) error

	InsertProduct(ctx context.Context, req *productCommon.Product, tx ...*sql.Tx) (id int32, err error)
	GetProduct(ctx context.Context, id int32, tx ...*sql.Tx) (product *productCommon.Product, err error)
	DeleteProduct(ctx context.Context, id int32, tx ...*sql.Tx) (err error)
	UpdateProductQuantity(ctx context.Context, req *productCommon.UpdateProductQuantity, tx ...*sql.Tx) (err error)
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
func (dbOps *DbOps) InsertProduct(ctx context.Context, req *productCommon.Product, tx ...*sql.Tx) (id int32, err error) {
	params := dbsqlc.InsertProductParams{Name: req.Name, Weight: req.Weight, Quantity: req.Quantity, Unit: req.Unit, PricePerProduct: req.PricePerItem}
	id, err = GetSqlcQuery(dbOps.DbSqlc, tx...).InsertProduct(ctx, params)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (dbOps *DbOps) GetProduct(ctx context.Context, id int32, tx ...*sql.Tx) (product *productCommon.Product, err error) {
	products, err := GetSqlcQuery(dbOps.DbSqlc, tx...).GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("no product found with id: %v", id)
	}
	firstProduct := products[0]
	product = &productCommon.Product{
		Id:       firstProduct.ID,
		Name:     firstProduct.Name,
		Weight:   firstProduct.Weight,
		Unit:     firstProduct.Unit,
		Quantity: firstProduct.Quantity,
	}
	return product, nil
}

func (dbOps *DbOps) DeleteProduct(ctx context.Context, id int32, tx ...*sql.Tx) (err error) {
	err = GetSqlcQuery(dbOps.DbSqlc, tx...).DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (dbOps *DbOps) UpdateProductQuantity(ctx context.Context, req *productCommon.UpdateProductQuantity, tx ...*sql.Tx) (err error) {
	args := dbsqlc.UpdateProductQuantityParams{ID: req.Id, Descreasecount: req.DescreaseQuantityCount}
	err = GetSqlcQuery(dbOps.DbSqlc, tx...).UpdateProductQuantity(ctx, args)
	if err != nil {
		return err
	}
	return nil
}
