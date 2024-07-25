-- name: CreateAccount :one
INSERT INTO
  accounts (owner, balance, currency)
VALUES
  ($1, $2, $3)
RETURNING
  *;

-- name: GetAccount :one
SELECT
  *
FROM
  accounts
WHERE
  id = $1
LIMIT
  1;

-- name: GetAccountForUpdate :one
SELECT
  *
FROM
  accounts
WHERE
  id = $1
LIMIT
  1 FOR NO KEY
UPDATE;

-- name: ListAccounts :many
SELECT
  *
FROM
  accounts
WHERE
  Owner = $1
ORDER BY
  id
Limit
  $2
Offset
  $3;

-- name: UpdateAccount :one
UPDATE accounts
set
  balance = $2
WHERE
  id = $1
RETURNING
  *;

-- name: AddAccountBalance :one
UPDATE accounts
set
  balance = balance + sqlc.arg (amount)
WHERE
  id = sqlc.arg (id)
RETURNING
  *;

-- name: DeleteAcount :exec
DELETE FROM accounts
WHERE
  id = $1;
