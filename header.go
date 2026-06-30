package go_epay

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"charset":      "utf-8",
	}
}
