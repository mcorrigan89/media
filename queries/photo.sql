-- name: GetPhotoByID :one
SELECT sqlc.embed(photo) FROM photo
WHERE photo.id = $1;

-- name: CreatePhoto :one
INSERT INTO photo (bucket, asset_id, width, height, size, owner_id) 
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;