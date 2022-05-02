CREATE TABLE if not exists workout(
                     id int generated always as identity PRIMARY KEY,
                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                     exerciseName text NOT NULL references exercise(exerciseName) default 'choose',
                     setCount integer NOT NULL default 3,
                     weight integer NOT NULL default 0,
                     repGoal integer NOT NULL default 0,
                     repResults integer[] NOT NULL default '{}',
                     sorenessRating integer NOT NULL default 0,
                     pumpRating integer NOT NULL default 0,
                     day_id serial NOT NULL references day(id)
);