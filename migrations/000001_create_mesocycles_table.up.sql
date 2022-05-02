CREATE TABLE if not exists mesocycle(
    id int generated always as identity PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    mesocycleNum serial not null
);