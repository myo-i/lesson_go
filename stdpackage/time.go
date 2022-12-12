/* http://golang.org/pkg/time/

 */

package _stdpackage

import (
	"fmt"
	"time"
)

func Time() {
	t := time.Now()
	fmt.Println(t)

	// RFC3339は頻繁に使用されるので覚えておく
	fmt.Println(t.Format(time.RFC3339))
	
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
