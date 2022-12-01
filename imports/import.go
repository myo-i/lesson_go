package imports

import (
	"fmt"
	"os/user"
	"time"
)

func Imports() {
	fmt.Println("Hello Import!!", time.Now())
	fmt.Println(user.Current())
}
