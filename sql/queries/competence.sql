-- name: InsertCompetence :exec
INSERT INTO
    competences (
        competence_id,
        code,
        name,
        created_at
    )
VALUES ($1, $2, $3, $4);

-- name: UpdateCompetence :exec
UPDATE competences
SET
    code = $2,
    name = $3,
    updated_at = $4
WHERE
    competence_id = $1
RETURNING
    *;

-- name: DeleteCompetence :execrows
DELETE FROM competences WHERE competence_id = $1;

-- name: FindCompetenceById :one
SELECT * FROM competences WHERE competence_id = $1;

-- name: FindAllCompetences :many
SELECT * FROM competences ORDER BY code ASC;