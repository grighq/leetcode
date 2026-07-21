package uniqueemailaddresses

import "strings"

func numUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]struct{})
	for _, email := range emails {
		emailParts := strings.Split(email, "@")
		localName, domainName := emailParts[0], emailParts[1]
		localName = strings.Split(localName, "+")[0]
		localName = strings.ReplaceAll(localName, ".", "")
		uniqueEmails[localName+"@"+domainName] = struct{}{}
	}

	return len(uniqueEmails)
}
