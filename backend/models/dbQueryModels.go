package models

//Keeping this up top since its a core part of the api as to how to retreive everything else

// Query to fetch the user ID based on the username obtained from the token
var getUserQuery = "SELECT id FROM user_ WHERE username = $1;"

// Note only doing C/R/U leaving off D as this will be handled as a cascading db issue.
//We dont want to just delete goals progress, assuming we would want to delete the whole goal not just a
//of the infromation.

var resetGoalProgress = `
	UPDATE activity
	SET progress = $1
	WHERE user_id = $2;
`
var InsertActivityProgress = "INSERT INTO activity (progress, user_id) VALUES ($1)"

var ShowGoalProgress = `
	SELECT g.name AS goal_name, a.progress AS progress
	FROM goal g
	JOIN activity a ON g.id = a.goal_id
	WHERE g.user_id = $1;
`
