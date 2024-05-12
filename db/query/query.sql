-- name: GetEquation :one
SELECT * FROM equation
WHERE id = $1 LIMIT 1;

-- name: GetEquationFromValue :one
SELECT * FROM equation
WHERE value = $1 LIMIT 1;

-- name: ListEquations :many
SELECT * FROM equation
ORDER BY id;

-- name: CreateEquation :one
INSERT INTO equation (
    value, category, created_at, updated_at
) VALUES (
    $1, $2, current_timestamp, current_timestamp
 )
RETURNING *;

-- name: UpdateEquation :one
UPDATE equation
set value=$2, updated_at=current_timestamp
WHERE id=$1
RETURNING *;

-- name: DeleteEquation :exec
DELETE FROM equation
WHERE id=$1;