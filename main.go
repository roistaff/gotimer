package main
import(
	"time"
	"fmt"
//	"atomicgo.dev/keyboard"
)
func timeInput()int{
	var timestr string
	fmt.Scan(&timestr)
	timertime,_ := time.ParseDuration(timestr)
	fmt.Print("")
	return int(timertime.Seconds())
}
func main(){
	timertime := timeInput()
	time.Sleep(time.Duration(timertime) * time.Second)
	fmt.Println("It's time.")
}

