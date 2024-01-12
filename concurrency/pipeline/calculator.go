package pipeline

func LaunchPipeline(amount int) int {
	firstChan := generator(amount)
	secondChan := power(firstChan, amount)
	thirdChan := sum(secondChan, amount)

	return <-thirdChan
}

func generator(amount int) <-chan int {
	out := make(chan int, amount)
	go func() {
		for i := 0; i <= amount; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func power(in <-chan int, amount int) <-chan int {
	out := make(chan int, amount)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int, amount int) <-chan int {
	out := make(chan int, amount)
	go func() {
		sum := 0
		for i := range in {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}
