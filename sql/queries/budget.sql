-- name: InsertBudget :exec
INSERT INTO
    budgets (
        budget_id,
        portfolio_id,
        cost_id,
        amount,
        created_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: BulkInsertBudget :exec
INSERT INTO
    budgets (
        budget_id,
        portfolio_id,
        cost_id,
        amount,
        created_at
    )
VALUES (
        unnest($1::text []),
        unnest($2::text []),
        unnest($3::text []),
        unnest($4::float8[]),
        unnest($5::timestamp[])
    );

-- name: FindBudgetById :one
SELECT * FROM budgets WHERE budget_id = $1;

-- name: UpdateBudget :exec
UPDATE budgets
SET
    portfolio_id = $2,
    cost_id = $3,
    amount = $4,
    updated_at = $5
WHERE
    budget_id = $1;

-- name: DeleteBudget :execrows
DELETE FROM budgets WHERE budget_id = $1 RETURNING *;

-- name: DeleteBudgetsByPortfolioId :execrows
DELETE FROM budgets WHERE portfolio_id = $1;

-- name: InsertBudgetAllocation :exec
INSERT INTO
    budget_allocations (
        budget_allocation_id,
        budget_id,
        amount,
        allocation_date,
        created_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: BulkInsertBudgetAllocation :exec
INSERT INTO
    budget_allocations (
        budget_allocation_id,
        budget_id,
        allocation_date,
        amount,
        created_at
    )
VALUES (
        unnest($1::text []),
        unnest($2::text []),
        unnest($3::date[]),
        unnest($4::float8[]),
        unnest($5::timestamp[])
    );

-- name: FindBudgetAllocations :many
SELECT *
FROM budget_allocations
WHERE
    budget_id = $1
ORDER BY allocation_date ASC;

;

-- name: DeleteBudgetAllocations :execrows
DELETE FROM budget_allocations WHERE budget_id = $1;

-- name: FindBudgetsByPortfolioId :many
SELECT * FROM budgets WHERE portfolio_id = $1;

-- name: FindBudgetsByPortfolioIdWithRelations :many
SELECT
    bu.budget_id AS budget_id,
    bu.portfolio_id AS portfolio_id,
    co.cost_type AS cost_type,
    co.description AS description,
    co.comment AS comment,
    co.amount AS cost_amount,
    co.currency AS cost_currency,
    co.tax AS cost_tax,
    co.apply_inflation AS cost_apply_inflation,
    bu.amount AS amount,
    bu.created_at AS created_at,
    bu.updated_at AS updated_at
FROM budgets AS bu
    INNER JOIN costs AS co ON bu.cost_id = co.cost_id
WHERE
    bu.portfolio_id = $1
ORDER BY co.cost_type, co.description;