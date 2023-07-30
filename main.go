package main

import (
	"fmt"
	"math"
)

func exponentialSmoothing(data []float64, alpha float64) []float64 {
	// Inisialisasi
	smoothedData := make([]float64, len(data))
	smoothedData[0] = data[0]

	// Peramalan dengan metode Exponential Smoothing
	for i := 1; i < len(data); i++ {
		smoothedData[i] = alpha*data[i] + (1-alpha)*smoothedData[i-1]
	}

	return smoothedData
}

func calculateMAE(actualData, predictedData []float64) float64 {
	if len(actualData) != len(predictedData) {
		panic("Panjang data aktual dan data prediksi harus sama")
	}

	totalAbsoluteError := 0.0
	n := float64(len(actualData))

	for i := 0; i < len(actualData); i++ {
		totalAbsoluteError += math.Abs(actualData[i] - predictedData[i])
	}

	mae := totalAbsoluteError / n
	return mae
}

func main() {
	data := []float64{66, 113, 116, 135, 90, 129, 79, 118, 98, 111, 153, 123, 140, 144, 140, 145, 73, 84, 131, 137, 136, 101, 135, 106, 102, 128, 138, 104, 116, 94, 118, 126, 144, 138, 109, 136, 134, 121, 129, 147, 113, 102, 104, 70, 138, 78, 86, 96, 136, 89, 124, 131, 93, 108, 151, 116, 111}

	alpha := 0.3

	// Hitung peramalan dengan metode Exponential Smoothing
	smoothedData := exponentialSmoothing(data, alpha)

	// Prediksi 12 minggu ke depan
	predictionData := make([]float64, 12)
	predictionData[0] = alpha*data[len(data)-1] + (1-alpha)*smoothedData[len(smoothedData)-1]

	for i := 1; i < 12; i++ {
		predictionData[i] = alpha*data[len(data)-1] + (1-alpha)*predictionData[i-1]
	}

	// Round smoothedData and predictionData to remove decimals
	for i := 0; i < len(smoothedData); i++ {
		smoothedData[i] = math.Round(smoothedData[i])
	}
	for i := 0; i < len(predictionData); i++ {
		predictionData[i] = math.Round(predictionData[i])
	}

	// Tampilkan hasil peramalan dan prediksi 7 minggu ke depan
	fmt.Println("Data Asli:", data)
	fmt.Println("Hasil Peramalan (Exponential Smoothing):", smoothedData)

	fmt.Println("\nPrediksi 7 Minggu Ke Depan:")
	fmt.Println(predictionData)

	// Menghitung MAE untuk hasil peramalan dan prediksi 7 minggu ke depan
	mae := calculateMAE(data, smoothedData)
	fmt.Println("\nMean Absolute Error (MAE) untuk hasil peramalan:", mae)

	// There's no need to calculate MAE for predictionData since it's not the actual data.
}
