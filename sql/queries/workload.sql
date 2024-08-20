-- name: InsertWorkload :exec
INSERT INTO
    workloads (
        workload_id,
        portfolio_id,
        effort_id,
        hours,
        created_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: BulkInsertWorkload :exec
INSERT INTO
    workloads (
        workload_id,
        portfolio_id,
        effort_id,
        hours,
        created_at
    )
VALUES (
        unnest($1::text []),
        unnest($2::text []),
        unnest($3::text []),
        unnest($4::int[]),
        unnest($5::timestamp[])
    );

-- name: FindWorkloadById :one
SELECT * FROM workloads WHERE workload_id = $1;

-- name: UpdateWorkload :exec
UPDATE workloads
SET
    portfolio_id = $2,
    effort_id = $3,
    hours = $4,
    updated_at = $5
WHERE
    workload_id = $1;

-- name: DeleteWorkload :execrows
DELETE FROM workloads WHERE workload_id = $1 RETURNING *;

-- name: DeleteWorkloadsByPortfolioId :execrows
DELETE FROM workloads WHERE portfolio_id = $1;

-- name: InsertWorkloadAllocation :exec
INSERT INTO
    workload_allocations (
        workload_allocation_id,
        workload_id,
        hours,
        allocation_date,
        created_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: BulkInsertWorkloadAllocation :exec
INSERT INTO
    workload_allocations (
        workload_allocation_id,
        workload_id,
        allocation_date,
        hours,
        created_at
    )
VALUES (
        unnest($1::text []),
        unnest($2::text []),
        unnest($3::date[]),
        unnest($4::int[]),
        unnest($5::timestamp[])
    );

-- name: FindWorkloadAllocations :many
SELECT *
FROM workload_allocations
WHERE
    workload_id = $1
ORDER BY allocation_date ASC;

-- name: FindWorkloadAllocationsGroupedByYear :many
SELECT EXTRACT(
        YEAR
        FROM workload_allocations.allocation_date
    )::int AS year, SUM(workload_allocations.hours)::int AS hours
FROM workload_allocations
WHERE
    workload_id = $1
GROUP BY
    year
ORDER BY year ASC;

-- name: DeleteWorkloadAllocations :execrows
DELETE FROM workload_allocations WHERE workload_id = $1;

-- name: FindWorkloadsByPortfolioId :many
SELECT * FROM workloads WHERE portfolio_id = $1;

-- name: FindWorkloadsByPortfolioIdWithRelations :many
SELECT
    w.workload_id AS workload_id,
    w.portfolio_id AS portfolio_id,
    c.code AS competence_code,
    c.name AS competence_name,
    e.comment AS comment,
    w.hours AS hours,
    w.created_at AS created_at,
    w.updated_at AS updated_at
FROM
    workloads AS w
    INNER JOIN efforts AS e ON w.effort_id = e.effort_id
    INNER JOIN competences AS c ON e.competence_id = c.competence_id
WHERE
    w.portfolio_id = $1
ORDER BY c.code;