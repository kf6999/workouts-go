CREATE TABLE if not exists day(
                     id int generated always as identity PRIMARY KEY,
                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                     dayNum integer NOT NULL,
                     week_id integer NOT NULL references week(id)
);