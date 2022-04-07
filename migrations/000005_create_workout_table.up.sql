CREATE TABLE workout(
                     id bigserial PRIMARY KEY,
                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                     exerciseName text NOT NULL,
                     setCount integer NOT NULL,
                     weight integer NOT NULL,
                     repGoal integer NOT NULL,
                     repResults integer NOT NULL,
                     sorenessRating integer NOT NULL,
                     pumpRating integer NOT NULL,
                     day_id bigserial NOT NULL,
                     constraint day_id
                        FOREIGN KEY (day_id)
                        REFERENCES day (id),
                     constraint exerciseName
                         foreign key (exerciseName)
                             references exercise(exerciseName)
);