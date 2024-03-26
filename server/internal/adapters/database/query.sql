-- name: GetCountry :one
SELECT * FROM Attraction
WHERE id = ? LIMIT 1;


-- name: PostReview :execresult
INSERT INTO Reviews (
    Title,
	Score,
	Description
) VALUES (
    ?,?,?
);


-- name: GetUser :one
SELECT ID, USERNAME, FULLNAME, EMAIL, PASSWORD 
FROM USER
WHERE EMAIL = ?; 