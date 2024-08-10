-- name: InsertEffort :exec
INSERT INTO
    efforts (
        effort_id,
        baseline_id,
        competence_id,
        comment,
        hours,
        created_at
    )
VALUES ($1, $2, $3, $4, $5, $6);

-- name: BulkInsertEffort :exec
INSERT INTO
    efforts (
        effort_id,
        baseline_id,
        competence_id,
        comment,
        hours,
        created_at
    )
VALUES (
        unnest($1::text []),
        unnest($2::text []),
        unnest($3::text []),
        unnest($4::text []),
        unnest($5::int[]),
        unnest($6::timestamp[])
    );

-- name: UpdateEffort :exec
UPDATE efforts
SET
    baseline_id = $2,
    competence_id = $3,
    comment = $4,
    hours = $5,
    updated_at = $6
WHERE
    effort_id = $1;

-- name: DeleteEffort :execrows
DELETE FROM efforts WHERE effort_id = $1;

-- name: FindEffortById :one
SELECT * FROM efforts WHERE effort_id = $1;

-- name: FindEffortsByBaselineIdWithRelations :many
SELECT
    e.effort_id AS effort_id,
    e.baseline_id AS baseline_id,
    e.competence_id AS competence_id,
    c.code AS competence_code,
    c.name AS competence_name,
    e.comment AS comment,
    e.hours AS hours,
    e.created_at AS created_at,
    e.updated_at AS updated_at
FROM efforts AS e
    INNER JOIN competences AS c ON e.competence_id = c.competence_id
WHERE
    e.baseline_id = $1
ORDER BY c.code ASC;

-- name: InsertEffortAllocation :exec
INSERT INTO
    effort_allocations (
        effort_allocation_id,
        effort_id,
        allocation_date,
        hours,
        created_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: BulkInsertEffortAllocation :exec
INSERT INTO
    effort_allocations (
        effort_allocation_id,
        effort_id,
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

-- name: DeleteEffortAllocations :execrows
DELETE FROM effort_allocations WHERE effort_id = $1;

-- name: FindEffortAllocationsByEffortId :many
SELECT *
FROM effort_allocations
WHERE
    effort_id = $1
ORDER BY allocation_date ASC;