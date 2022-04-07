CREATE TABLE day(
                     id bigserial PRIMARY KEY,
                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                     dayNum integer NOT NULL,
                     week_id bigserial NOT NULL,
                     constraint week_id
                         foreign key (week_id)
                             references week(id)
);