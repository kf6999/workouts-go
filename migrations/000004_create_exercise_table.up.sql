CREATE TABLE exercise(
    exerciseName text PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    exerciseFocus text NOT NULL,
    videoUrl text NOT NULL,
    tenRM int NOT NULL,
    exerciseType text NOT NULL
);