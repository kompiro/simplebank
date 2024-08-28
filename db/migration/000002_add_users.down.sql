ALTER TABLE if exists "accounts"
DROP CONSTRAINT if EXISTS "ownwer_currency_key";

ALTER TABLE if exists "accounts"
DROP CONSTRAINT if EXISTS "accounts_owner_fkey";

DROP TABLE if exists "users";
