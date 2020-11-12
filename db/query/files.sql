
-- name: SaveFile :one
INSERT INTO files (
    filename, 
    filesize,
    fileextension,
    uploaded_at,
    uploaded_by
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;


-- name: GetFile :one
SELECT * FROM files
WHERE id = $1 LIMIT 1;

-- name: ListFiles :many
SELECT * FROM files
ORDER BY name;
