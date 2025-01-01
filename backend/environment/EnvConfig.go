package environment

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createEnvFile(connectionString, dbUser string) {
	// Define the content of the .env file
	content := fmt.Sprintf(`
	DATABASE_URL=%s
	DB_USER=%s
	`, connectionString, dbUser)

	// Write content to the .env file
	err := os.WriteFile("environment/.env", []byte(content), 0666)
	if err != nil {
		fmt.Printf("Unable to write file: %v\n", err)
		return
	}

	fmt.Println(".env file created successfully!")
}

func LoadEnvFile(envFilePath string) error {

	connectionString := "postgres://postgres:Enable@localhost:5432/test"
	dbUser := "postgres"
	createEnvFile(connectionString, dbUser)

	envFile, err := os.Open(envFilePath)
	if err != nil {
		return err
	}
	defer envFile.Close()

	// Reads data from envFile line by line (like readlines() in Python)
	scanner := bufio.NewScanner(envFile)

	for scanner.Scan() {
		//fmt.Println(scanner.Text()) //Prints Each line of the file

		line := strings.TrimSpace(scanner.Text()) // Current line, trimmed of whitespace

		// Skip empty lines or comments
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line from '=' into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Only select lines with key/val format (one '=' sign)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value) //Set env variable.

	}

	return scanner.Err()
}

//err := os.Setenv(key, value)
