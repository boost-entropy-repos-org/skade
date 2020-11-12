
-- name: CreateReport :one
INSERT INTO reports (
    file,
    malicious
) VALUES (
    $1, $2
) RETURNING *;


-- name: GetReport :one
SELECT * FROM reports
WHERE id = $1 LIMIT 1;

-- name: ListReports :many
SELECT * FROM reports
ORDER BY name;
