package utils

func Exists(userID string, stringArr []*string) bool {
	for _, v := range stringArr {
		if userID == *v {
			return true
		}
	}
	return false
}
