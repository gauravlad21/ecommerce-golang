
-- name: InsertOrder :one
INSERT INTO user_order(
    user_id, unique_order_id, total_price, version
)VALUES(
    $1, $2, $3, 1
)
RETURNING id;

-- name: InsertSubOrder :one
INSERT INTO sub_order(
    unique_order_id, product_id, quantity, status, version
)VALUES(
    $1, $2, $3, $4, 1
)
RETURNING id;

-- name: GetOrder :many
SELECT user_order.unique_order_id, sub_order.product_id, sub_order.quantity, sub_order.status as suborderstatus, user_order.user_id, user_order.status as orderstatus, user_order.total_price
FROM user_order JOIN sub_order ON user_order.unique_order_id=sub_order.sub_order
where user_order.unique_order_id=$1 and user_id=$2;

-- name: UpdateUserOrderStatus :exec
UPDATE user_order
SET status=$2, version=version+1
WHERE id=$1;

-- name: UpdateSubOrderStatus :exec
UPDATE sub_order
SET status=$2, version=version+1
WHERE id=$1;