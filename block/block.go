package block

// function to add an array of numbers.
func sum(s []int, d chan struct{}) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	<-d
}

func done(t int, d chan struct{}) {
	for i := 0; i < t; i++ {
		d <- struct{}{}
	}
}
