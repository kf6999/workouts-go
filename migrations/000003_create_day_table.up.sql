CREATE TABLE if not exists day(
                     id int generated always as identity PRIMARY KEY,
                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                     dayNum serial NOT NULL,
                     week_id serial NOT NULL references week(id)
);