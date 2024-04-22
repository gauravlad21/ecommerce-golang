// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package dbsqlc

import (
	"context"
)

const deleteProduct = `-- name: DeleteProduct :exec
UPDATE product
SET is_deleted = true, version=version+1
WHERE id=$1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :many
SELECT id, name, weight, unit, quantity, version, is_deleted, created, updated FROM product
WHERE id=$1 and is_deleted=false
`

func (q *Queries) GetProduct(ctx context.Context, id int32) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProduct, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Weight,
			&i.Unit,
			&i.Quantity,
			&i.Version,
			&i.IsDeleted,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertProduct = `-- name: InsertProduct :one
INSERT INTO product(
    name, weight, unit, quantity, version
)VALUES(
    $1, $2, $3, $4, 1
)
RETURNING id
`

type InsertProductParams struct {
	Name     string
	Weight   int32
	Unit     string
	Quantity int32
}

func (q *Queries) InsertProduct(ctx context.Context, arg InsertProductParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertProduct,
		arg.Name,
		arg.Weight,
		arg.Unit,
		arg.Quantity,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateProductQuantity = `-- name: UpdateProductQuantity :exec
UPDATE product
SET quantity=quantity - $2::int, version=version+1
WHERE id=$1
`

type UpdateProductQuantityParams struct {
	ID             int32
	Descreasecount int32
}

func (q *Queries) UpdateProductQuantity(ctx context.Context, arg UpdateProductQuantityParams) error {
	_, err := q.db.ExecContext(ctx, updateProductQuantity, arg.ID, arg.Descreasecount)
	return err
}
