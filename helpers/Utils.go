package helpers

import (
	"fmt"
	"os"
	"regexp"
)

// PathExist check the path directory if exist
func PathExist(p string) bool {
	if stat, err := os.Stat(p); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// AlphaNum character validation
func AlphaNum(v interface{}) error {

	if v.(string) == "" {
		return nil
	}

	if true == regexp.MustCompile("^[0-9a-zA-Z]+$").MatchString(v.(string)) {
		return nil
	}

	return fmt.Errorf("may only contain alpha-numeric characters")
}
