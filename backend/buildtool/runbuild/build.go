package buildtool

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type envVars struct {
	DB_PASSWORD string
	DB_USERNAME string
}

func CreateFile() {
	E := &envVars{}

	envFile, err := os.Create(".env")
	if err != nil {
		fmt.Println("Error creating .env file:", err)
		os.Exit(1)
	}
	defer envFile.Close()

	//to do get cmd input pass to the set env as string

	fmt.Println("input text:")
	var w1, w2, w3 string
	n, err := fmt.Scan(&w1, &w2, &w3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read text: %s %s %s-\n", w1, w2, w3)

	fmt.Printf("number of items read: %d\n", n)
	fmt.Printf("read line: %s %s %s-\n", w1, w2, w3)

	os.Setenv("DB_USERNAME", getdbb)
	os.Setenv("DB_PASSWORDY", E.DB_PASSWORD)

	// Run 'go generate'
	cmd := exec.Command("go", "generate")

	err2 := cmd.Run()
	if err2 != nil {
		fmt.Println("Error running 'go generate':", err2)
		os.Exit(1)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Code generation completed successfully.")
}
