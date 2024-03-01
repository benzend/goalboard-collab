BEGIN;

 
CREATE TABLE IF NOT EXISTS goals_ (
	goalId          SERIAL PRIMARY KEY,
    Name            VARCHAR(50)
    TargetPerDay    VARCHAR     
    LongTermTarget  VARCHAR     
);

COMMIT;