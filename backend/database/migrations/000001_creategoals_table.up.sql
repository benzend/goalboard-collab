BEGIN;

CREATE TABLE IF NOT EXISTS goals_ (
    goalId          SERIAL PRIMARY KEY,
    Name            VARCHAR(50),
    TargetPerDay    VARCHAR,
    LongTermTarget  VARCHAR,
    user_id         INT,  -- Define the user_id column
    CONSTRAINT user_fk
        FOREIGN KEY(user_id) 
        REFERENCES user_(id)  -- Reference the id column in the user_ table
);

COMMIT;
