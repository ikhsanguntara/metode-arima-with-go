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

func calculateMAEAndMSE(actualData, predictedData []float64) (float64, float64) {
	if len(actualData) != len(predictedData) {
		panic("Panjang data aktual dan data prediksi harus sama")
	}

	n := float64(len(actualData))
	totalAbsoluteError := 0.0
	totalSquaredError := 0.0

	for i := 0; i < len(actualData); i++ {
		absoluteError := math.Abs(actualData[i] - predictedData[i])
		totalAbsoluteError += absoluteError
		squaredError := math.Pow(absoluteError, 2)
		totalSquaredError += squaredError
	}

	mae := totalAbsoluteError / n
	mse := totalSquaredError / n
	return mae, mse
}

func calculateMAPE(actualData, predictedData []float64) float64 {
	if len(actualData) != len(predictedData) {
		panic("Panjang data aktual dan data prediksi harus sama")
	}

	n := float64(len(actualData))
	totalAbsolutePercentageError := 0.0

	for i := 0; i < len(actualData); i++ {
		percentageError := math.Abs((actualData[i] - predictedData[i]) / actualData[i])
		totalAbsolutePercentageError += percentageError
	}

	mape := (totalAbsolutePercentageError / n) * 100.0
	return mape
}

func main() {
	// Data contoh (gantilah data ini dengan data Anda) (data product perminggu )

	// 	SELECT DATE_FORMAT(DATE_SUB(a.created_at, INTERVAL WEEKDAY(a.created_at) DAY), '%Y-%m-%d') AS start_of_week,
	// 	DATE_FORMAT(DATE_ADD(a.created_at, INTERVAL 6 - WEEKDAY(a.created_at) DAY), '%Y-%m-%d') AS end_of_week,
	// 	SUM(b.qty) AS total_qty
	// FROM transactions a
	// LEFT JOIN transaction_lines b ON a.id = b.transaction_id
	// WHERE b.name = 'Pupuk Kompos (5KG)'
	// AND a.created_at >= '2022-07-01 01:43:26.419'
	// AND a.created_at < NOW()  -- Batas tanggal akhir diganti dengan NOW() untuk data hingga hari ini
	// GROUP BY start_of_week, end_of_week
	// ORDER BY start_of_week ASC;

	data := []float64{
		122,
		113,
		116,
		135,
		90,
		129,
		79,
		118,
		98,
		111,
		153,
		123,
		140,
		144,
		140,
		145,
		73,
		84,
		131,
		137,
		136,
		101,
		135,
		106,
		102,
		128,
		138,
		104,
		116,
		94,
		118,
		126,
		144,
		138,
		109,
		136,
		134,
		121,
		129,
		147,
		113,
		102,
		104,
		70,
		138,
		78,
		86,
		96,
		136,
		89,
		124,
		131,
		93,
		108,
		151,
		116,
		111}

	alpha := 0.3

	// Hitung peramalan dengan metode Exponential Smoothing
	smoothedData := exponentialSmoothing(data, alpha)

	// Prediksi 12 minggu ke depan // tolong dibuatkan param biar bisa di ganti dari web
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

	// Tampilkan hasil peramalan dan prediksi 12 minggu ke depan
	fmt.Println("Data Asli:", data)
	fmt.Println("Hasil Peramalan (Exponential Smoothing):", smoothedData)

	fmt.Println("\nPrediksi 12 Minggu Ke Depan:")
	fmt.Println(predictionData)

	// Menghitung MAE, MSE, dan MAPE untuk hasil peramalan
	mae, mse := calculateMAEAndMSE(data, smoothedData)
	mape := calculateMAPE(data, smoothedData)

	fmt.Println("\nMean Absolute Error (MAE) untuk hasil peramalan:", mae)
	fmt.Println("Mean Squared Error (MSE) untuk hasil peramalan:", mse)
	fmt.Println("Mean Absolute Percentage Error (MAPE) untuk hasil peramalan:", mape)
}
