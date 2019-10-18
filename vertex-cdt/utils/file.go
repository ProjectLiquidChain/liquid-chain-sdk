package utils

import (
	"fmt"
	"os/exec"
)

func DeleteFile(file string) {
	cmd := exec.Command("rm", file)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
func DeleteFolder(folder string) {
	cmd := exec.Command("rm", "-rf", folder)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(out))
}
