CREATE TABLE users_one (
    id UUID PRIMARY KEY UNIQUE DEFAULT gen_random_uuid(),
    created_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    full_name TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true
);