CREATE TABLE IF NOT EXISTS movies (
    id bigint PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title text NOT NULL,
    description text NOT NULL,
    genre text NOT NULL,
    runtime integer NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone
);

ALTER TABLE movies ADD CONSTRAINT movies_runtime_check CHECK (runtime > 0);
