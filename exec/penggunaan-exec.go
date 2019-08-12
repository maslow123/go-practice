package main

import "fmt"
import "os/exec"

func main() {
	var output1, _ = exec.Command("ls").Output()
	fmt.Printf("-> ls\n%s\n", string(output1))

	var output2, _ = exec.Command("cmd","/C","mkdir tes").Output()
	fmt.Printf("-> mkdir tes\n%s\n", string(output2))
	
	var output3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf("-> git config user.name\n%s\n", string(output3))
	
	// if runtime.GOOS == "windows" {
	// 	var output, err = exec.Command("cmd","/C", "mkdir tes").Output()
	// } else {
	// 	var output, err = exec.Command("bash", "-c", "git config user.name").Output()
	// }

}