package auth

import (
	"io/ioutil"
	"strings"
)

func parseAuthFile(filePath string) (map[string]string, error) {
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	userRows := strings.Split(string(fileContent), "\n")
	users := make(map[string]string, len(userRows))
	for _, row := range userRows {
		user := strings.Split(row, ":")
		users[user[0]] = user[1]
	}

	return users, nil
}
