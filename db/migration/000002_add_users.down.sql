ALTER TABLE "accounts" IF EXISTS DROP CONSTRAINT IF EXISTS "owner_currency_key";

ALTER TABLE "accounts" IF EXISTS DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

DROP TABLE IF EXISTS "users";