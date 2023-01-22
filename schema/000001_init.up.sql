CREATE TABLE client (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    user_email VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE account (
    id BIGSERIAL PRIMARY KEY,
    owner_email VARCHAR NOT NULL,
    balance BIGINT,
    currency VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE transaction (
    id BIGSERIAL PRIMARY KEY,
    type VARCHAR NOT NULL,
    from_account_id BIGINT,
    to_account_id BIGINT NOT NULL ,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX ON "client" ("user_email");
CREATE INDEX ON "account" ("owner_email");
CREATE INDEX ON "transaction" ("from_account_id");
CREATE INDEX ON "transaction" ("to_account_id");

ALTER TABLE "account" ADD CONSTRAINT fk_account_client FOREIGN KEY ("owner_email") REFERENCES "client" ("user_email");
ALTER TABLE "transaction" ADD CONSTRAINT fk_transaction_from_client FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");
ALTER TABLE "transaction" ADD CONSTRAINT fk_transaction_to_client FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");


CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON account
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();