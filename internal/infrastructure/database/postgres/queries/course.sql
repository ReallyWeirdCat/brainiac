-- name: GetCourseByGUID :one
SELECT * FROM course
WHERE guid = $1 AND deleted_at IS NULL
LIMIT 1;
