-- name: InsertBaseline :exec
INSERT INTO
    baselines (
        baseline_id,
        code,
        title,
        description,
        start_date,
        duration,
        manager_id,
        estimator_id,
        created_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9
    );

-- name: FindBaselineById :one
SELECT * FROM baselines WHERE baseline_id = $1;

-- name: UpdateBaseline :exec
UPDATE baselines
SET
    code = $2,
    title = $3,
    description = $4,
    start_date = $5,
    duration = $6,
    manager_id = $7,
    estimator_id = $8,
    updated_at = $9
WHERE
    baseline_id = $1;

-- name: DeleteBaseline :one
DELETE FROM baselines WHERE baseline_id = $1 RETURNING *;