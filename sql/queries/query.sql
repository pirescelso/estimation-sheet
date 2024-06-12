-- name: CreateProject :exec
INSERT INTO
    projects (
        project_id,
        description,
        start_date,
        created_at
    )
VALUES ($1, $2, $3, $4);

-- name: CreateCost :exec
INSERT INTO
    costs (
        cost_id,
        project_id,
        cost_type,
        description,
        comment,
        amount,
        currency,
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
        $8
    );

-- name: GetCost :one
SELECT * FROM costs WHERE cost_id = $1;

-- name: UpdateCost :exec
UPDATE costs
SET
    project_id = $2,
    cost_type = $3,
    description = $4,
    comment = $5,
    amount = $6,
    currency = $7,
    updated_at = $8
WHERE
    cost_id = $1;

-- name: CreateInstallment :exec
INSERT INTO
    installments (
        installment_id,
        cost_id,
        payment_date,
        amount,
        created_at
    )
VALUES ($1, $2, $3, $4, $5);

-- name: CreateInstallments :exec
INSERT INTO
    installments (
        installment_id,
        cost_id,
        payment_date,
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

-- name: GetInstallments :many
SELECT
    installment_id,
    cost_id,
    payment_date,
    amount,
    created_at
FROM installments
WHERE
    cost_id = $1;