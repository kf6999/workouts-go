CREATE TABLE mesocycle(
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    mesocycleNum integer NOT NULL
);