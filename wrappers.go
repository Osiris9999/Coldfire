package coldfire

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// F is a wrapper for the Sprintf function.
func F(str string, arg ...interface{}) string {
	return fmt.Sprintf(str, arg...)
}

func f(s string, arg ...interface{}) string {
	return fmt.Sprintf(s, arg...)
}
