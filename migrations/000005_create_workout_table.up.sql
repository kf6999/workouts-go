CREATE TABLE if not exists workout(
                     id int generated always as identity PRIMARY KEY,
                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                     exerciseName text NOT NULL references exercise(exerciseName),
                     setCount integer NOT NULL,
                     weight integer NOT NULL,
                     repGoal integer NOT NULL,
                     repResults integer[] NOT NULL,
                     sorenessRating integer NOT NULL,
                     pumpRating integer NOT NULL,
                     day_id integer NOT NULL references day(id)

);