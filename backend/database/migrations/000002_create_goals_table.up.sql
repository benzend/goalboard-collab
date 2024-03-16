BEGIN;

CREATE TABLE IF NOT EXISTS goal (
    goal_id           SERIAL PRIMARY KEY,
    name              VARCHAR(50),
    target_per_day    VARCHAR,
    long_term_target  VARCHAR,
    user_id           INT,  -- Define the user_id column
    CONSTRAINT user_fk
        FOREIGN KEY(user_id)
        REFERENCES user_(id)  -- Reference the id column in the user_ table
);

COMMIT;
