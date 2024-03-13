-- name: AddUserReturnId :one
INSERT INTO users (name) 
VALUES (?) 
RETURNING id;

-- name: AddTime :exec
INSERT OR REPLACE INTO times (id, user_id, time, type)
VALUES (
    (SELECT id FROM times t WHERE t.user_id = ? AND t.type = ?),
    ?, ?, ?
);

-- name: GetTotalTime :one
SELECT SUM(time) total_time
FROM times
WHERE user_id = ?
GROUP BY user_id;

-- name: GetAllTotalTimes :many
SELECT t.user_id, u.name, SUM(t.time) total_time, COUNT(*) quests_done
FROM times t
INNER JOIN users u
    ON t.user_id = u.id
GROUP BY user_id
ORDER BY 
    quests_done DESC,
    total_time ASC;

