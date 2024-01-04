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

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	s := bufio.NewScanner(r)
	cnt := 0
	for s.Scan() {
		var user User
		if err = user.UnmarshalJSON(s.Bytes()); err != nil {
			return
		}
		result[cnt] = user
		cnt++
	}

	return
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
