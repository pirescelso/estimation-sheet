-- name: InsertCost :exec
INSERT INTO
    costs (
        cost_id,
        baseline_id,
        cost_type,
        description,
        comment,
        amount,
        currency,
        tax,
        apply_inflation,
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

-- name: BulkInsertCost :exec
INSERT INTO
    costs (
        cost_id,
        baseline_id,
        cost_type,
        description,
        comment,
        amount,
        currency,
        tax,
        apply_inflation,
        created_at
    )
VALUES (
        unnest($1::text []),
        unnest($2::text []),
        unnest($3::text []),
        unnest($4::text []),
        unnest($5::text []),
        unnest($6::float8[]),
        unnest($7::text []),
        unnest($8::float8[]),
        unnest($9::boolean[]),
        unnest($10::timestamp[])
    );

-- name: FindCostById :one
SELECT * FROM costs WHERE cost_id = $1;

-- name: UpdateCost :exec
UPDATE costs
SET
    baseline_id = $2,
    cost_type = $3,
    description = $4,
    comment = $5,
    amount = $6,
    currency = $7,
    tax = $8,
    apply_inflation = $9,
    updated_at = $10
WHERE
    cost_id = $1;

-- name: DeleteCost :one
DELETE FROM costs WHERE cost_id = $1 RETURNING *;

-- name: InsertCostAllocation :exec
INSERT INTO
    cost_allocations (
        cost_allocation_id,
        cost_id,
        allocation_date,
        amount,
        created_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: BulkInsertCostAllocation :exec
INSERT INTO
    cost_allocations (
        cost_allocation_id,
        cost_id,
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
-- name: DeleteCostAllocations :execrows
DELETE FROM cost_allocations WHERE cost_id = $1;

-- name: FindCostsByBaselineId :many
SELECT *
FROM costs
WHERE
    baseline_id = $1
ORDER BY cost_type, description ASC;

-- name: FindCostAllocationsByCostId :many
SELECT *
FROM cost_allocations
WHERE
    cost_id = $1
ORDER BY allocation_date ASC;