package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/montanaflynn/stats"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	"github.com/sjwhitworth/golearn/trees"
)

// Define variables/containers
type Response struct {
	Coefficients []float64
}

var LinRegresponse Response

func main() {

	//Start Timer
	startTime := time.Now()
	time.Sleep(startTime.Truncate(time.Millisecond).Add(time.Millisecond).Sub(startTime))
	startTime = time.Now()

	/////////////////////
	//DATA PREP
	//Data for LinReg
	//////////////////////
	dis := []float64{4.09, 4.9671, 4.9671, 6.0622, 6.0622, 6.0622, 5.5605, 5.9505, 6.0821, 6.5921, 6.3467, 6.2267, 5.4509, 4.7075, 4.4619, 4.4986, 4.4986, 4.2579, 3.7965, 3.7965, 3.7979, 4.0123, 3.9769, 4.0952, 4.3996, 4.4546, 4.682, 4.4534, 4.4547, 4.239, 4.233, 4.175, 3.99, 3.7872, 3.7598, 3.3603, 3.3779, 3.9342, 3.8473, 5.4011, 5.4011, 5.7209, 5.7209, 5.7209, 5.7209, 5.1004, 5.1004, 5.6894, 5.87, 6.0877, 6.8147, 6.8147, 6.8147, 6.8147, 7.3197, 8.6966, 9.1876, 8.3248, 7.8148, 6.932, 7.2254, 6.8185, 7.2255, 7.9809, 9.2229, 6.6115, 6.6115, 6.498, 6.498, 6.498, 5.2873, 5.2873, 5.2873, 5.2873, 4.2515, 4.5026, 4.0522, 4.0905, 5.0141, 4.5026, 5.4007, 5.4007, 5.4007, 5.4007, 4.7794, 4.4377, 4.4272, 3.7476, 3.4217, 3.4145, 3.0923, 3.0921, 3.6659, 3.6659, 3.615, 3.4952, 3.4952, 3.4952, 3.4952, 3.4952, 2.7778, 2.8561, 2.7147, 2.7147, 2.421, 2.1069, 2.211, 2.1224, 2.4329, 2.5451, 2.7778, 2.6775, 2.3534, 2.548, 2.2565, 2.4631, 2.7301, 2.7474, 2.4775, 2.7592, 2.2577, 2.1974, 2.0869, 1.9444, 2.0063, 1.9929, 1.7572, 1.7883, 1.8125, 1.9799, 2.1185, 2.271, 2.3274, 2.4699, 2.346, 2.1107, 1.9669, 1.8498, 1.6686, 1.6687, 1.6119, 1.4394, 1.3216, 1.4118, 1.3459, 1.4191, 1.5166, 1.4608, 1.5296, 1.5257, 1.618, 1.5916, 1.6102, 1.6232, 1.7494, 1.7455, 1.7364, 1.8773, 1.7573, 1.7659, 1.7984, 1.9709, 2.0407, 2.162, 2.422, 2.2834, 2.0459, 2.4259, 2.1, 2.2625, 2.4259, 2.3887, 2.5961, 2.6463, 2.7019, 3.1323, 3.5549, 3.3175, 2.9153, 2.829, 2.741, 2.5979, 2.7006, 2.847, 2.9879, 3.2797, 3.1992, 3.7886, 4.5667, 4.5667, 6.4798, 6.4798, 6.4798, 6.2196, 6.2196, 5.6484, 7.309, 7.309, 7.309, 7.6534, 7.6534, 6.27, 6.27, 5.118, 5.118, 3.9454, 4.3549, 4.3549, 4.2392, 3.875, 3.8771, 3.665, 3.6526, 3.9454, 3.5875, 3.9454, 3.1121, 3.4211, 2.8893, 3.3633, 2.8617, 3.048, 3.2721, 3.2721, 2.8944, 2.8944, 3.2157, 3.2157, 3.3751, 3.3751, 3.6715, 3.6715, 3.8384, 3.6519, 3.6519, 3.6519, 4.148, 4.148, 6.1899, 6.1899, 6.3361, 6.3361, 7.0355, 7.0355, 7.9549, 7.9549, 8.0555, 8.0555, 7.8265, 7.8265, 7.3967, 7.3967, 8.9067, 8.9067, 9.2203, 9.2203, 6.3361, 1.801, 1.8946, 2.0107, 2.1121, 2.1398, 2.2885, 2.0788, 1.9301, 1.9865, 2.1329, 2.4216, 2.872, 3.9175, 4.429, 4.429, 3.9175, 4.3665, 4.0776, 4.2673, 4.7872, 4.8628, 4.1403, 4.1007, 4.6947, 5.2447, 5.2119, 5.885, 7.3073, 7.3073, 9.0892, 7.3172, 7.3172, 7.3172, 5.1167, 5.1167, 5.1167, 5.5027, 5.5027, 5.9604, 5.9604, 6.32, 7.8278, 7.8278, 7.8278, 5.4917, 5.4917, 5.4917, 4.022, 3.37, 3.0992, 3.1827, 3.3175, 3.1025, 2.5194, 2.6403, 2.834, 3.2628, 3.6023, 3.945, 3.9986, 4.0317, 3.5325, 4.0019, 4.5404, 4.5404, 4.7211, 4.7211, 4.7211, 5.4159, 5.4159, 5.4159, 5.2146, 5.2146, 5.8736, 6.6407, 6.6407, 6.4584, 6.4584, 5.9853, 5.2311, 5.615, 4.8122, 4.8122, 4.8122, 7.0379, 6.2669, 5.7321, 6.4654, 8.0136, 8.0136, 8.5353, 8.344, 8.7921, 8.7921, 10.7103, 10.7103, 12.1265, 10.5857, 10.5857, 2.1222, 2.5052, 2.7227, 2.5091, 2.5182, 2.2955, 2.1036, 1.9047, 1.9047, 1.6132, 1.7523, 1.5106, 1.3325, 1.3567, 1.2024, 1.1691, 1.1296, 1.1742, 1.137, 1.3163, 1.3449, 1.358, 1.3861, 1.3861, 1.4165, 1.5192, 1.5804, 1.5331, 1.4395, 1.4261, 1.4672, 1.5184, 1.5895, 1.7281, 1.9265, 2.1678, 1.77, 1.7912, 1.7821, 1.7257, 1.6768, 1.6334, 1.4896, 1.5004, 1.5888, 1.5741, 1.639, 1.7028, 1.6074, 1.4254, 1.1781, 1.2852, 1.4547, 1.4655, 1.413, 1.5275, 1.5539, 1.5894, 1.6582, 1.8347, 1.8195, 1.6475, 1.8026, 1.794, 1.8589, 1.8746, 1.9512, 2.0218, 2.0635, 1.9096, 1.9976, 1.8629, 1.9356, 1.9682, 2.0527, 2.0882, 2.2004, 2.3158, 2.2222, 2.1247, 2.0026, 1.9142, 1.8206, 1.8172, 1.8662, 2.0651, 2.0048, 1.9784, 1.8956, 1.9879, 2.072, 2.198, 2.2616, 2.185, 2.3236, 2.3552, 2.3682, 2.4527, 2.4961, 2.4358, 2.5806, 2.7792, 2.7831, 2.7175, 2.5975, 2.5671, 2.7344, 2.8016, 2.9634, 3.0665, 2.8715, 2.5403, 2.9084, 2.8237, 3.0334, 3.0993, 2.8965, 2.5329, 2.4298, 2.206, 2.3053, 2.1007, 2.1705, 1.9512, 3.4242, 3.3317, 3.4106, 4.0983, 3.724, 3.9917, 3.5459, 3.1523, 1.8209, 1.7554, 1.8226, 1.8681, 2.1099, 2.3817, 2.3817, 2.7986, 2.7986, 2.8927, 2.4091, 2.3999, 2.4982, 2.4786, 2.2875, 2.1675, 2.3889, 2.505}
	mv := []float64{24, 21.6, 34.7, 33.4, 36.2, 28.7, 22.9, 22.1, 16.5, 18.9, 15, 18.9, 21.7, 20.4, 18.2, 19.9, 23.1, 17.5, 20.2, 18.2, 13.6, 19.6, 15.2, 14.5, 15.6, 13.9, 16.6, 14.8, 18.4, 21, 12.7, 14.5, 13.2, 13.1, 13.5, 18.9, 20, 21, 24.2, 30.8, 34.9, 26.6, 25.3, 24.7, 21.2, 19.3, 20, 16.6, 14.4, 19.4, 19.7, 20.5, 25, 23.4, 18.9, 35.4, 24.7, 31.6, 23.3, 19.6, 18.7, 16, 22.2, 25, 33, 23.5, 19.4, 22, 17.4, 20.9, 24.2, 21.7, 22.8, 23.4, 24.1, 21.4, 20, 20.8, 21.2, 20.3, 28, 23.9, 24.8, 22.9, 23.9, 26.6, 22.5, 22.2, 23.6, 28.7, 22.6, 22, 22.9, 25, 20.6, 28.4, 21.4, 38.7, 43.8, 33.2, 27.5, 26.5, 18.6, 19.3, 20.1, 19.5, 19.5, 20.4, 19.8, 19.4, 21.7, 22.8, 18.8, 18.7, 18.5, 18.3, 21.2, 19.2, 20.4, 19.3, 22, 20.3, 20.5, 17.3, 18.8, 21.4, 15.7, 16.2, 18, 14.3, 19.2, 19.6, 23, 18.4, 15.6, 18.1, 17.4, 17.1, 13.3, 17.8, 14, 14.4, 13.4, 15.6, 11.8, 13.8, 15.6, 14.6, 17.8, 15.4, 21.5, 19.6, 15.3, 19.4, 17, 15.6, 13.1, 41.3, 24.3, 23.3, 27, 50, 50, 50, 22.7, 25, 50, 23.8, 23.8, 22.3, 17.4, 19.1, 23.1, 23.6, 22.6, 29.4, 23.2, 24.6, 29.9, 37.2, 39.8, 36.2, 37.9, 32.5, 26.4, 29.6, 50, 32, 29.8, 34.9, 33, 30.5, 36.4, 31.1, 29.1, 50, 33.3, 30.3, 34.6, 34.9, 32.9, 24.1, 42.3, 48.5, 50, 22.6, 24.4, 22.5, 24.4, 20, 21.7, 19.3, 22.4, 28.1, 23.7, 25, 23.3, 28.7, 21.5, 23, 26.7, 21.7, 27.5, 30.1, 44.8, 50, 37.6, 31.6, 46.7, 31.5, 24.3, 31.7, 41.7, 48.3, 29, 24, 25.1, 31.5, 23.7, 23.3, 27, 20.1, 22.2, 23.7, 17.6, 18.5, 24.3, 20.5, 24.5, 26.2, 24.4, 24.8, 29.6, 42.8, 21.9, 20.9, 44, 50, 36, 30.1, 33.8, 43.1, 48.8, 31, 36.5, 22.8, 30.7, 50, 43.5, 20.7, 21.1, 25.2, 24.4, 35.2, 32.4, 32, 33.2, 33.1, 29.1, 35.1, 45.4, 35.4, 46, 50, 32.2, 22, 20.1, 23.2, 22.3, 24.8, 28.5, 37.3, 27.9, 23.9, 21.7, 28.6, 27.1, 20.3, 22.5, 29, 24.8, 22, 26.4, 33.1, 36.1, 28.4, 33.4, 28.2, 22.8, 20.3, 16.1, 22.1, 19.4, 21.6, 23.8, 16.2, 17.8, 19.8, 23.1, 21, 23.8, 23.1, 20.4, 18.5, 25, 24.6, 23, 22.2, 19.3, 22.6, 19.8, 17.1, 19.4, 22.2, 20.7, 21.1, 19.5, 18.5, 20.6, 19, 18.7, 32.7, 16.5, 23.9, 31.2, 17.5, 17.2, 23.1, 24.5, 26.6, 22.9, 24.1, 18.6, 30.1, 18.2, 20.6, 17.8, 21.7, 22.7, 22.6, 25, 19.9, 20.8, 16.8, 21.9, 27.5, 21.9, 23.1, 50, 50, 50, 50, 50, 13.8, 13.8, 15, 13.9, 13.3, 13.1, 10.2, 10.4, 10.9, 11.3, 12.3, 8.8, 7.2, 10.5, 7.4, 10.2, 11.5, 15.1, 23.2, 9.7, 13.8, 12.7, 13.1, 12.5, 8.5, 5, 6.3, 5.6, 7.2, 12.1, 8.3, 8.5, 5, 11.9, 27.9, 17.2, 27.5, 15, 17.2, 17.9, 16.3, 7, 7.2, 7.5, 10.4, 8.8, 8.4, 16.7, 14.2, 20.8, 13.4, 11.7, 8.3, 10.2, 10.9, 11, 9.5, 14.5, 14.1, 16.1, 14.3, 11.7, 13.4, 9.6, 8.2, 8.4, 12.8, 10.5, 17.1, 14.8, 15.4, 10.8, 11.8, 14.9, 12.6, 14.1, 13, 13.4, 15.2, 16.1, 17.8, 14.4, 14.1, 12.7, 13.5, 14.9, 20, 16.4, 17.7, 19.5, 20.2, 21.4, 19.9, 19, 19.1, 19.1, 20.1, 19.9, 19.6, 23.2, 29.8, 13.8, 13.3, 16.7, 12, 14.6, 21.4, 23, 23.7, 25, 21.8, 20.6, 21.2, 19.1, 20.6, 15.2, 7, 8.1, 13.6, 20.1, 21.8, 24.5, 23.1, 19.7, 18.3, 21.2, 17.5, 16.8, 22.4, 20.6, 23.9, 22, 19}

	//////////////////////
	//Data for RandTree
	/////////////////////

	rand.Seed(44111342)

	// Load csv
	houses, err := base.ParseCSVToInstances("boston.csv", true)
	if err != nil {
		panic(err)
	}

	// Make dataset discrete with Chi-Merge
	filt := filters.NewChiMergeFilter(houses, 0.999)
	for _, a := range base.NonClassFloatAttributes(houses) {
		filt.AddAttribute(a)
	}
	filt.Train()
	housesf := base.NewLazilyFilteredInstances(houses, filt)

	// 60-40 training-test split
	trainData, testData := base.InstancesTrainTestSplit(housesf, 0.60)

	///////////////////////////
	//MODEL 1
	//PERFORMING RANDOM TREE
	///////////////////////////

	nRuns := 100

	//Channel for Model1
	accuracyChan := make(chan float64, nRuns)

	for i := 0; i < nRuns; i++ {

		go func(i int) {
			var localErr error
			// Using two randomly-chosen attributes
			tree := trees.NewRandomTree(2)

			// Train
			localErr = tree.Fit(trainData)
			if localErr != nil {
				panic(localErr)
			}

			// Declare variables
			var predictions base.FixedDataGrid
			var cf evaluation.ConfusionMatrix

			// Predict
			predictions, localErr = tree.Predict(testData)
			if localErr != nil {
				fmt.Println(localErr)
			}

			// Test Predictions
			cf, localErr = evaluation.GetConfusionMatrix(testData, predictions)
			if localErr != nil {
				panic(fmt.Sprintf("Unable to get confusion matrix: %s", localErr.Error()))
			}
			accuracy := evaluation.GetAccuracy(cf)
			accuracyChan <- accuracy
		}(i)
	}

	//Print results for each run
	for i := 0; i < nRuns; i++ {
		accuracy := <-accuracyChan
		fmt.Printf("Run %d (Random Tree): Accuracy = %.2f%%\n", i+1, accuracy*100)
	}

	close(accuracyChan)

	////////////////////////
	//MODEL 2
	// Linear regression
	////////////////////////

	//Channel for Model2
	responseChan := make(chan Response, nRuns)

	for i := 0; i < nRuns; i++ {
		go func(i int) {
			response := RunLinReg(dis, mv, nRuns)
			responseChan <- response
		}(i)
	}

	// Collect and print results
	var coefficientsList []Response

	for i := 0; i < nRuns; i++ {
		select {
		case response := <-responseChan:
			coefficientsList = append(coefficientsList, response)
		}
	}

	// Print for every single run
	for i, response := range coefficientsList {
		fmt.Printf("Run %d (Linear Regression): Coefficients: %v\n", i+1, response.Coefficients)
	}

	//End Timer
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Seconds()
	if elapsedTime == 0 {
		elapsedTime = 0.00000000001
	}
	fmt.Println("total runtime:", elapsedTime, "seconds")
}

// Helper Functions Below //

// Linear Regression Function
func RunLinReg(dis []float64, mv []float64, nRuns int) Response {

	for i := 0; i < nRuns; i++ {

		points, _ := stats.LinearRegression(
			MakeCoordinates(
				dis,
				mv),
		)
		LinRegresponse.Coefficients = EquationLine(points)
	}
	return LinRegresponse
}

// Rounding numbers
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Ceil(val*ratio) / ratio
}

// Equation Line for Linear Regression
func EquationLine(points []stats.Coordinate) []float64 {

	coords := MinMaxCoordinates(points)

	x1 := coords[0].X
	y1 := coords[0].Y
	x2 := coords[len(coords)-1].X
	y2 := coords[len(coords)-1].Y

	m := (y2 - y1) / (x2 - x1)
	b := y1 - m*x1

	container := []float64{b, m}

	return container
}

// MakeCoordinates for Linear Regression
func MakeCoordinates(x, y []float64) []stats.Coordinate {
	container := make([]stats.Coordinate, len(x))

	for i := 0; i < len(x); i++ {
		container[i] = stats.Coordinate{
			x[i], y[i],
		}
	}
	return container
}

// MinMaxCoordinates for Equation Line(Linear Regression)
func MinMaxCoordinates(x []stats.Coordinate) []stats.Coordinate {
	minX := x[0].X
	maxX := x[0].X
	minXY := x[0]
	maxXY := x[0]

	for _, point := range x {
		if point.X < minX {
			minX = point.X
			minXY = point
		}
		if point.X > maxX {
			maxX = point.X
			maxXY = point
		}
	}
	container := []stats.Coordinate{minXY, maxXY}

	return container
}
