package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//easyjson:json
type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	s := bufio.NewScanner(r)
	stat := DomainStat{}

	for s.Scan() {
		user := &User{}
		if err := user.UnmarshalJSON(s.Bytes()); err != nil {
			return nil, fmt.Errorf("user.UnmarshalJSON: %w", err)
		}

		dom := getDomainFromEmail(user.Email, domain)
		if dom == "" {
			continue
		}

		num := stat[dom]
		num++
		stat[dom] = num
	}

	return stat, nil
}

func getDomainFromEmail(email, domain string) string {
	if !strings.HasSuffix(email, "."+domain) {
		return ""
	}

	sp := strings.SplitN(email, "@", 2)
	if len(sp) != 2 {
		return ""
	}

	return strings.ToLower(sp[1])
}
