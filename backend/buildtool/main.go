package main

/*Keeping this here for future use
// // Run 'go generate'
// cmd := exec.Command("go", "generate")

// err2 := cmd.Run()
// if err2 != nil {
// 	fmt.Println("Error running 'go generate':", err2)
// 	os.Exit(1)
// }

// cmd.Stdout = os.Stdout
// cmd.Stderr = os.Stderr
*/
import (
	buildtool "github.com/benzend/goalboard/buildtool/runbuild"
)

func main() {
	buildtool.CreateFile()
}
