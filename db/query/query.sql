-- name: GetEquation :one
SELECT * FROM equations
WHERE id = $1 LIMIT 1;

-- name: ListEquations :many
SELECT * FROM equations
ORDER BY id;

-- name: CreateEquation :one
INSERT INTO equations (
    value
) VALUES (
    $1
 )
RETURNING *;

-- name: UpdateEquation :one
UPDATE equations
set value = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEquation :exec
DELETE FROM equations
WHERE id = $1;