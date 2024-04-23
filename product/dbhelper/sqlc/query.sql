-- name: InsertProduct :one
INSERT INTO product(
    name, weight, unit, quantity, price_per_product, version
)VALUES(
    $1, $2, $3, $4, $5, 1
)
RETURNING id;

-- name: GetProduct :many
SELECT * FROM product
WHERE id=$1 and is_deleted=false;

-- name: DeleteProduct :exec
UPDATE product
SET is_deleted = true, version=version+1
WHERE id=$1;

-- name: UpdateProductQuantity :exec
UPDATE product
SET quantity=quantity - @descreaseCount::int, version=version+1
WHERE id=$1;