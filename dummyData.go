package Ajit

var db map[string]bool

func init() {
	db = make(map[string]bool)
	db["ajithkumar.sinha@srsconsultinginc.com"] = true

}

func UserAvilable(email string) bool {
	if _, ok := db[email]; !ok {
		return false
	}
	return true
}
