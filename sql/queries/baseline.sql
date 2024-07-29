-- name: InsertBaseline :exec
INSERT INTO
    baselines (
        baseline_id,
        code,
        review,
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
        $9,
        $10
    );

-- name: FindBaselineById :one
SELECT * FROM baselines WHERE baseline_id = $1;

-- name: FindBaselineByIdWithRelations :one
SELECT baselines.*, managers.name AS manager, estimators.name AS estimator
FROM
    baselines
    INNER JOIN users AS managers ON managers.user_id = baselines.manager_id
    INNER JOIN users AS estimators ON estimators.user_id = baselines.estimator_id
WHERE
    baseline_id = $1;

-- name: UpdateBaseline :exec
UPDATE baselines
SET
    code = $2,
    review = $3,
    title = $4,
    description = $5,
    start_date = $6,
    duration = $7,
    manager_id = $8,
    estimator_id = $9,
    updated_at = $10
WHERE
    baseline_id = $1;

-- name: DeleteBaseline :one
DELETE FROM baselines WHERE baseline_id = $1 RETURNING *;

-- name: FindAllBaselines :many
SELECT baselines.*, managers.name AS manager, estimators.name AS estimator
FROM
    baselines
    INNER JOIN users AS managers ON managers.user_id = baselines.manager_id
    INNER JOIN users AS estimators ON estimators.user_id = baselines.estimator_id
ORDER BY code ASC, review DESC;