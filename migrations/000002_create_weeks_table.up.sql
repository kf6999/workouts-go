CREATE TABLE if not exists week
(
    id           int generated always as identity PRIMARY KEY,
    created_at   timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    weekNum      integer                     NOT NULL,
    mesocycle_id integer                     NOT NULL references mesocycle (id)
);