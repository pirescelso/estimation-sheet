-- name: InsertPortfolio :exec
INSERT INTO
    portfolios (
        portfolio_id,
        baseline_id,
        plan_id,
        start_date,
        created_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: UpdatePortfolio :exec
UPDATE portfolios
SET
    baseline_id = $2,
    plan_id = $3,
    start_date = $4,
    updated_at = $5
WHERE
    portfolio_id = $1;

-- name: DeletePortfolio :one
DELETE FROM portfolios WHERE portfolio_id = $1 RETURNING *;

-- name: CountPortfoliosByPlanId :one
SELECT count(*) FROM portfolios WHERE plan_id = $1;

-- name: CountPortfoliosByBaselineId :one
SELECT count(*) FROM portfolios WHERE baseline_id = $1;

-- name: FindPortfolioById :one
SELECT * FROM portfolios WHERE portfolio_id = $1;

-- name: FindPortfolioByIdWithRelations :one
SELECT
    pf.portfolio_id AS portfolio_id,
    pl.code AS plan_code,
    bl.code AS code,
    bl.review AS review,
    bl.title AS title,
    bl.description AS description,
    pf.start_date AS start_date,
    bl.duration AS duration,
    ma.name AS manager,
    es.name AS estimator,
    pf.created_at AS created_at,
    pf.updated_at AS updated_at
FROM
    baselines AS bl
    INNER JOIN portfolios AS pf ON bl.baseline_id = pf.baseline_id
    INNER JOIN users AS ma ON ma.user_id = bl.manager_id
    INNER JOIN users AS es ON es.user_id = bl.estimator_id
    INNER JOIN plans AS pl ON pl.plan_id = pf.plan_id
WHERE
    pf.portfolio_id = $1;

-- name: FindAllPortfoliosByPlanIdWithRelations :many
SELECT
    pf.portfolio_id AS portfolio_id,
    pl.code AS plan_code,
    bl.code AS code,
    bl.review AS review,
    bl.title AS title,
    bl.description AS description,
    pf.start_date AS start_date,
    bl.duration AS duration,
    ma.name AS manager,
    es.name AS estimator,
    pf.created_at AS created_at,
    pf.updated_at AS updated_at
FROM
    baselines AS bl
    INNER JOIN portfolios AS pf ON bl.baseline_id = pf.baseline_id
    INNER JOIN users AS ma ON ma.user_id = bl.manager_id
    INNER JOIN users AS es ON es.user_id = bl.estimator_id
    INNER JOIN plans AS pl ON pl.plan_id = pf.plan_id
WHERE
    pf.plan_id = $1
ORDER BY bl.code, pl.code ASC;

-- name: FindAllPortfoliosWithRelations :many
SELECT
    pf.portfolio_id AS portfolio_id,
    pl.code AS plan_code,
    bl.code AS code,
    bl.review AS review,
    bl.title AS title,
    bl.description AS description,
    pf.start_date AS start_date,
    bl.duration AS duration,
    ma.name AS manager,
    es.name AS estimator,
    pf.created_at AS created_at,
    pf.updated_at AS updated_at
FROM
    baselines AS bl
    INNER JOIN portfolios AS pf ON bl.baseline_id = pf.baseline_id
    INNER JOIN users AS ma ON ma.user_id = bl.manager_id
    INNER JOIN users AS es ON es.user_id = bl.estimator_id
    INNER JOIN plans AS pl ON pl.plan_id = pf.plan_id
ORDER BY bl.code ASC, bl.review DESC, pl.code ASC;