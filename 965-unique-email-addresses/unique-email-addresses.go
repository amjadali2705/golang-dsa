func numUniqueEmails(emails []string) int {
    uniqueEmails := make(map[string]bool)

    for _, email := range emails {
        parts := strings.Split(email, "@")
        localName := parts[0]
        domainName := parts[1]

        cleanLocalName := ""
        for _, char := range localName {
            if char == '+' {
                break
            }
            
            if char != '.' {
                cleanLocalName += string(char)
            }
        }
        
        canonicalEmail := cleanLocalName + "@" + domainName
        uniqueEmails[canonicalEmail] = true
    }

    return len(uniqueEmails)
}