package main

	
import (
    "fmt"
    "math/rand"
	"time"
)

var samples = [][]int{
	//{x1, x2, b, label}
	{-1, 0, 1, 1}, {0, 1, 1, 1}, 
	{0, -1, 1, -1}, {1, 0, 1, -1}, 
}

var model = []int{0, 0, 0}

func decisionBoundary(model []int, sample []int) bool {
	x1, x2, b, label := sample[0], sample[1], sample[2], sample[3]
	return label*(x1*model[0]+x2*model[1]+b*model[2]) > 0
}

func randIntSymmetrically(n int)int{
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	return r1.Intn(n*2)-n
}

func main() {
	// model[0], model[1], model[2] = randIntSymmetrically(2), randIntSymmetrically(2), randIntSymmetrically(2)
	modelUpdated := false

	for i := 1;; i++ {
		modelUpdated = false
		fmt.Printf("Epoch=%d\n", i)
		
		for _, sample := range samples {
			fmt.Printf("model: x1=%d, x2=%d, b=%d\n", model[0], model[1], model[2])
			if !decisionBoundary(model, sample) {
				label := sample[3]
				model[0] = model[0] + sample[0]*label
				model[1] = model[1] + sample[1]*label
				model[2] = model[2] + sample[2]*label
				modelUpdated = true
				fmt.Printf("model updated: x1=%d, x2=%d, b=%d\n", model[0], model[1], model[2])
			}
		}

		if !modelUpdated{
			break
		}

		fmt.Println()
	}
}
