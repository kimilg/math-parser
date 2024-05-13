-- name: GetEquation :one
SELECT e.*,
       json_agg(DISTINCT jsonb_build_object('id', v.id, 'name', v.name, 'vcategory', v.vcategory))
       FILTER (WHERE v.id IS NOT NULL) AS variable
FROM equation e
    LEFT OUTER JOIN equation_variable ev 
        ON e.id = ev.equation_id 
    LEFT OUTER JOIN variable v 
        ON ev.variable_id = v.id
WHERE e.id = $1 
GROUP BY e.id
LIMIT 1;

-- name: GetEquationFromValue :one
SELECT e.*,
       json_agg(DISTINCT jsonb_build_object('id', v.id, 'name', v.name, 'vcategory', v.vcategory))
       FILTER (WHERE v.id IS NOT NULL) AS variable
FROM equation e
    LEFT OUTER JOIN equation_variable ev 
        ON e.id = ev.equation_id 
    LEFT OUTER JOIN variable v 
        ON ev.variable_id = v.id
WHERE e.value = $1
GROUP BY e.id
LIMIT 1;

-- name: ListEquations :many
SELECT e.*, 
       json_agg(DISTINCT jsonb_build_object('id', v.id, 'name', v.name, 'vcategory', v.vcategory)) 
       FILTER (WHERE v.id IS NOT NULL) AS variable
FROM equation e 
    LEFT OUTER JOIN equation_variable ev 
        ON e.id = ev.equation_id 
    LEFT OUTER JOIN variable v 
        ON ev.variable_id = v.id
GROUP BY e.id
ORDER BY e.id;

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