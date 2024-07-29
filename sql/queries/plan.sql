-- name: InsertPlan :exec
INSERT INTO
    plans (
        plan_id,
        code,
        name,
        assumptions,
        created_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: FindPlanById :one
SELECT * FROM plans WHERE plan_id = $1;

-- name: FindPlanByCode :one
SELECT * FROM plans WHERE code = $1;

-- name: FindAllPlans :many
SELECT * FROM plans;

-- name: UpdatePlan :one
UPDATE plans
SET
    code = $2,
    name = $3,
    assumptions = $4,
    updated_at = $5
WHERE
    plan_id = $1
RETURNING
    *;

-- name: DeletePlan :one
DELETE FROM plans WHERE plan_id = $1 RETURNING *;