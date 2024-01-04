//go:build !bench
// +build !bench

package hw10programoptimization

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDomainStat(t *testing.T) {
	data := `{"Id":1,"Name":"Howard Mendoza","Username":"0Oliver","Email":"aliquid_qui_ea@Browsedrive.gov","Phone":"6-866-899-36-79","Password":"InAQJvsq","Address":"Blackbird Place 25"}
{"Id":2,"Name":"Jesse Vasquez","Username":"qRichardson","Email":"mLynch@broWsecat.com","Phone":"9-373-949-64-00","Password":"SiZLeNSGn","Address":"Fulton Hill 80"}
{"Id":3,"Name":"Clarence Olson","Username":"RachelAdams","Email":"RoseSmith@Browsecat.com","Phone":"988-48-97","Password":"71kuz3gA5w","Address":"Monterey Park 39"}
{"Id":4,"Name":"Gregory Reid","Username":"tButler","Email":"5Moore@Teklist.net","Phone":"520-04-16","Password":"r639qLNu","Address":"Sunfield Park 20"}
{"Id":5,"Name":"Janice Rose","Username":"KeithHart","Email":"nulla@Linktype.com","Phone":"146-91-01","Password":"acSBF5","Address":"Russell Trail 61"}`

	t.Run("find 'com'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "com")
		require.NoError(t, err)
		require.Equal(t, DomainStat{
			"browsecat.com": 2,
			"linktype.com":  1,
		}, result)
	})

	t.Run("find 'gov'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "gov")
		require.NoError(t, err)
		require.Equal(t, DomainStat{"browsedrive.gov": 1}, result)
	})

	t.Run("find 'unknown'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "unknown")
		require.NoError(t, err)
		require.Equal(t, DomainStat{}, result)
	})
}

func Test_countDomains(t *testing.T) {
	tests := []struct {
		name           string
		source         users
		domain         string
		expectedResult DomainStat
	}{
		{
			name: "success",
			source: users{
				User{
					Email: "qwerty@google.com",
				},
				User{
					Email: "ololo@google.com",
				},
				User{
					Email: "ololo@amazon.com",
				},
				User{
					Email: "ololo@aMazOn.com",
				},
				User{
					Email: "5Moore.com@Teklist.net",
				},
			},
			domain: "com",
			expectedResult: DomainStat{
				"google.com": 2,
				"amazon.com": 2,
			},
		},
		{
			name: "no matches",
			source: users{
				User{
					Email: "qwerty@google.com",
				},
				User{
					Email: "5Moore.com@Teklist.net",
				},
			},
			domain:         "ru",
			expectedResult: DomainStat{},
		},
		{
			name: "no matches: domain in email",
			source: users{
				User{
					Email: "ru@google.com",
				},
			},
			domain:         "ru",
			expectedResult: DomainStat{},
		},
		{
			name: "no matches: invalid email: domain is missed",
			source: users{
				User{
					Email: "ru@com",
				},
			},
			domain:         "com",
			expectedResult: DomainStat{},
		},
		{
			name: "no matches: invalid email: at(@) is missed",
			source: users{
				User{
					Email: "mail.com",
				},
			},
			domain:         "com",
			expectedResult: DomainStat{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, _ := countDomains(tc.source, tc.domain)
			require.Equal(t, tc.expectedResult, result)
		})
	}
}
