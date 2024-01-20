package password

func Authenticate(pwd string) bool {
	if pwd == "hello" {
		return true
	} else {
		return false
	}
}
