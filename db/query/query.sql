-- name: GetEquation :one
SELECT e.*,
       json_agg(DISTINCT jsonb_build_object('id', v.id, 'name', v.name, 'vcategory', v.vcategory, 'arguments', v.arguments))
       FILTER (WHERE v.id IS NOT NULL) AS variables
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
       json_agg(DISTINCT jsonb_build_object('id', v.id, 'name', v.name, 'vcategory', v.vcategory, 'arguments', v.arguments))
       FILTER (WHERE v.id IS NOT NULL) AS variables
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
       json_agg(DISTINCT jsonb_build_object('id', v.id, 'name', v.name, 'vcategory', v.vcategory, 'arguments', v.arguments)) 
       FILTER (WHERE v.id IS NOT NULL) AS variables
FROM equation e 
    LEFT OUTER JOIN equation_variable ev 
        ON e.id = ev.equation_id 
    LEFT OUTER JOIN variable v 
        ON ev.variable_id = v.id
GROUP BY e.id
ORDER BY e.id;

-- name: InsertEquation :one
INSERT INTO equation (value, category, cause, effect, created_at, updated_at) 
VALUES ($1, $2, $3, $4, current_timestamp, current_timestamp)
RETURNING *;

-- name: InsertVariable :one
WITH inserted_id AS (
    INSERT INTO variable (name, vcategory, arguments, created_at, updated_at)
        VALUES ($1, $2, $3, current_timestamp, current_timestamp)
        ON CONFLICT (name, vcategory) DO NOTHING
        RETURNING id
), 
inserted_id_union AS (
    SELECT id 
    FROM inserted_id
    UNION
    SELECT v.id
    FROM variable v
    WHERE v.name = $1
      AND v.vcategory = $2
),
eq_var AS (
    INSERT INTO equation_variable (equation_id, variable_id)
        VALUES ($4, (SELECT id FROM inserted_id_union))
        ON CONFLICT (equation_id, variable_id) DO NOTHING
        RETURNING *
)
SELECT * FROM eq_var
UNION 
SELECT * FROM equation_variable ev
    WHERE ev.equation_id = $4 AND
          ev.variable_id = (SELECT id FROM inserted_id_union);

-- name: UpdateEquation :one
UPDATE equation
SET value=$2, category=$3, cause=$4, effect=$5, updated_at=current_timestamp
WHERE id=$1
RETURNING *;

-- name: DeleteEquation :exec
DELETE FROM equation
WHERE id=$1;

-- name: DeleteVariable :exec
DELETE FROM variable
WHERE id=$1;