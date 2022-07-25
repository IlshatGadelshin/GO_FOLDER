/*
Реализуйте тип Stopwatch, который будет сохранять время каждой промежуточной фиксации секундомера и выдавать результаты относительно общего времени старта.
Тип должен обладать следующими методами:
Start() — запустить/сбросить секундомер;
SaveSplit() — сохранить промежуточное время;
GetResults() []time.Duration — вернуть текущие результаты.
*/

package main
import (
	"fmt"
	"time"
)
//Пример работы:
func main(){
	sw := Stopwatch{}
    sw.Start()

    time.Sleep(1 * time.Second)
    sw.SaveSplit()

    time.Sleep(500 * time.Millisecond)
    sw.SaveSplit()

    time.Sleep(300 * time.Millisecond)
    sw.SaveSplit()

    fmt.Println(sw.GetResults())
}

type Stopwatch struct{

		start 
		result []float32

}

func (start *Stopwatch) Start() {
	//time.Now()
}