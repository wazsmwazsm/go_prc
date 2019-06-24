package main
import (
	"fmt"
	"os"
)

func main() {
	// 1) os.StartProcess //
	/* Linux: */
	env := os.Environ()
	procAttr := &os.ProcAttr{
				Env: env,
				Files: []*os.File{
					os.Stdin,
					os.Stdout,
					os.Stderr,
				},
			}
	// 1st example: list files
	// ls -l 的结果会输出到标准输出
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)  
	if err != nil {
			fmt.Printf("Error %v starting process!", err)  //
			os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)
}
