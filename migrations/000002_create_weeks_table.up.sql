CREATE TABLE week(
                          id bigserial PRIMARY KEY,
                          created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                          weekNum integer NOT NULL,
                          mesocycle_id bigserial NOT NULL,
                 constraint mesocycle_id
                 foreign key (mesocycle_id)
                 references mesocycle(id)
);