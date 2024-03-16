BEGIN;

CREATE TABLE IF NOT EXISTS activity (
    id        SERIAL PRIMARY KEY,
    progress  FLOAT,
    goal_id   INT,
    CONSTRAINT goal_fk
        FOREIGN KEY (goal_id)
        REFERENCES goal (goal_id)  -- Reference the goalId column in the goals_ table
);


COMMIT;
