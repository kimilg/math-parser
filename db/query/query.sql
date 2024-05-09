-- name: GetEquation :one
SELECT * FROM equations
WHERE id = $1 LIMIT 1;

-- name: GetEquationFromValue :one
SELECT * FROM equations
WHERE value = $1 LIMIT 1;

-- name: ListEquations :many
SELECT * FROM equations
ORDER BY id;

-- name: CreateEquation :one
INSERT INTO equations (
    value, created_at, updated_at
) VALUES (
    $1, current_timestamp, current_timestamp
 )
RETURNING *;

-- name: UpdateEquation :one
UPDATE equations
set value=$2, updated_at=current_timestamp
WHERE id=$1
RETURNING *;

-- name: DeleteEquation :exec
DELETE FROM equations
WHERE id=$1;