package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

//easyjson:json
type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users []User

func getUsers(r io.Reader) (users, error) {
	s := bufio.NewScanner(r)
	result := make(users, 0)

	for s.Scan() {
		var user User
		if err := user.UnmarshalJSON(s.Bytes()); err != nil {
			return nil, fmt.Errorf("user.UnmarshalJSON: %w", err)
		}
		result = append(result, user)
	}

	return result, nil
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	reg, err := regexp.Compile("\\." + domain)
	if err != nil {
		return nil, fmt.Errorf("regexp.Compile: %w", err)
	}

	for _, user := range u {
		matched := reg.Match([]byte(user.Email))

		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
