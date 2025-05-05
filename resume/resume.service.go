package resume

import (
	"fmt"
	"log"
)

func GetByUsername(username string) (MyResume, error) {
	resume, err := FindByUsername(username)
	if err != nil {
		log.Println("Error getting resume:", err)
		return MyResume{}, fmt.Errorf("failed to get resume: %w", err)
	}

	return resume, nil
}
