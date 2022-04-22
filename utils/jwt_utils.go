package utils

func RemoveBearer(token string) (string, error) {
	const Bearer_schema = "Bearer "
	if token == "" || len(token) < 7 {
		return "", nil
	} else {
		return token[len(Bearer_schema):], nil
	}
}
