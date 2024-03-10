BEGIN 

CREATE TABLE IF NOT EXISTS activity_ (
    id        SERIAL PRIMARY KEY,
    Progress  FLOAT,
    goal_id   INT,
    CONSTRAINT goal_fk
        FOREIGN KEY (goal_id) 
        REFERENCES goals_ (goalId)  -- Reference the goalId column in the goals_ table
);
 

COMMIT;