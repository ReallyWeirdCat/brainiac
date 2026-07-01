-- name: GetCourseByGUID :one
SELECT * FROM course
WHERE guid = $1
LIMIT 1;
