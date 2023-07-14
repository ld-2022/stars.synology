package command

import (
	"fmt"
	"testing"
)

func TestSudoExec(t *testing.T) {
	err, message := SudoExec(`echo 1sss23 > /var/root/test.txt && ls /var/root/`, "123456789..")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(message)
}
