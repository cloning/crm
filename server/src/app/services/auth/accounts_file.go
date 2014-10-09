package auth

import (
	"io/ioutil"
	"strings"
)

func parseAccountsFile(filePath string) (map[string]string, error) {
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	userRows := strings.Split(string(fileContent), "\n")
	users := make(map[string]string, len(userRows))
	for _, row := range userRows {
		row = strings.TrimSpace(row)

		if strings.Index(row, "#") == 0 {
			continue
		}

		user := strings.Split(row, ":")

		if len(user) != 2 {
			continue
		}

		users[user[0]] = user[1]
	}

	return users, nil
}
