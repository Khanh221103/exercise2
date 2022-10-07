package main

import (
	"context"
	"fmt"
	"time"
)

}
func ex1() {
	for i := 0; i < 3; i++ {
		timeWait := time.After(3 * time.Second)
		a := time.Now().UnixMilli()
		fmt.Printf("%d\n", a)
		if i < 2 {
			<-timeWait
		}
	}
	fmt.Println("Kết thúc")
}

func ex2() {
	now := time.Now().Unix()
	days := now / 60 / 60 /24
	fmt.Println("Day: ", days)
}

func ex3(ctx context.Context) {
	canceledChannel := make(chan bool)

	go func() {
		<-ctx.Done()
		fmt.Println(ctx.Err())
		canceledChannel <- true
	}()

	isCanceledChannel := <-canceledChannel
	if isCanceledChannel {
		close(canceledChannel)
		return
	}

	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("Sleep!")
	}
}

func ex4() {
	timeUnix := time.Unix(0,1592190294764144364)
	fmt.Println(timeUnix)
	fmt.Printf("Số phút: %d\n", timeUnix.Minute())
}

func ex5() {
	timeUnix := time.Unix(1592190385, 0)
	fmt.Println(timeUnix)
	fmt.Printf("Day: %q, Number of day: %d\n", timeUnix.Weekday(), int(timeUnix.Weekday()))

}

func ex7() {
	_, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	context := context.WithValue(context.Background(), "timeNow", time.Now().UnixNano())
	time.Sleep(3 * time.Second)
	value1 := time.Now().UnixNano()
	value2 := context.Value("timeNow").(int64)
	fmt.Print(value1 - value2)
}

func ex8() {
	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)
	go func ()  {
		for {
			<-ticker.C
                fmt.Printf("%d done\n", time.Now().Unix())
		}
	}()
    done <-true 
}

func study() {
	fmt.Println("i'm study")
}

func ex9() {
	time.AfterFunc(100 * time.Millisecond, study)
	time.Sleep(time.Second)
}

func main() {
	fmt.Println("==========EX1==========")
	ex1()

	fmt.Println("==========EX2==========")
	ex2()

	fmt.Println("==========EX3==========")
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	ex3(ctx)
	defer cancel()

	fmt.Println("==========EX4==========")
	ex4()

	fmt.Println("==========EX5==========")
	ex5()

	fmt.Println("==========EX7==========")
	ex7()

	fmt.Println("==========EX8==========")
	ex8()

	fmt.Println("==========EX9==========")
	ex9()
}


 
