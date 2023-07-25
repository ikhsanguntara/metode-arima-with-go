package main

import (
	"fmt"
)

// Fungsi untuk menghitung rata-rata dari sebuah slice float64
func mean(data []float64) float64 {
	sum := 0.0
	for _, val := range data {
		sum += val
	}
	return sum / float64(len(data))
}

// Fungsi untuk menghitung selisih data dengan nilai rata-rata
func subtractMean(data []float64) []float64 {
	meanValue := mean(data)
	result := make([]float64, len(data))
	for i, val := range data {
		result[i] = val - meanValue
	}
	return result
}

// Fungsi untuk menghitung varians dari sebuah slice float64
func variance(data []float64) float64 {
	meanValue := mean(data)
	sumSquaredDiff := 0.0
	for _, val := range data {
		diff := val - meanValue
		sumSquaredDiff += diff * diff
	}
	return sumSquaredDiff / float64(len(data)-1)
}

// Fungsi untuk menghitung autokorelasi
func autocorrelation(data []float64, lag int) float64 {
	n := len(data)
	meanData := mean(data)
	varianceData := variance(data)

	var numerator, denominator float64
	for i := 0; i < n-lag; i++ {
		numerator += (data[i] - meanData) * (data[i+lag] - meanData)
	}
	denominator = float64(n-lag) * varianceData

	return numerator / denominator
}

// Fungsi untuk menghitung koefisien autoregressive (AR)
func calculateARCoefficients(data []float64, p int) []float64 {
	coefficients := make([]float64, p)

	for i := 1; i <= p; i++ {
		coefficients[i-1] = autocorrelation(data, i)
	}

	return coefficients
}

// Fungsi untuk menghitung residual dari model ARIMA
func calculateResidual(data []float64, arCoefficients []float64) []float64 {
	n := len(data)
	p := len(arCoefficients)
	residual := make([]float64, n-p)

	copy(residual, data[:p])

	for i := p; i < n; i++ {
		prediction := float64(0)
		for j := 0; j < p; j++ {
			prediction += arCoefficients[j] * data[i-j-1]
		}
		residual = append(residual, data[i]-prediction)
	}

	return residual
}

// Fungsi untuk menghitung moving average (MA)
func calculateMACoefficients(residual []float64, q int) []float64 {
	coefficients := make([]float64, q)

	for i := 1; i <= q; i++ {
		coefficients[i-1] = autocorrelation(residual, i)
	}

	return coefficients
}

// Fungsi untuk melakukan prediksi menggunakan model ARIMA
func predictARIMA(data []float64, p, d, q, predictionLength int) []float64 {
	// Lakukan differencing jika d > 0
	for i := 0; i < d; i++ {
		diffData := make([]float64, len(data)-1)
		for j := 1; j < len(data); j++ {
			diffData[j-1] = data[j] - data[j-1]
		}
		data = diffData
	}

	arCoefficients := calculateARCoefficients(data, p)

	// Lakukan residual calculation
	residual := calculateResidual(data, arCoefficients)

	// Lakukan moving average jika q > 0
	if q > 0 {
		maCoefficients := calculateMACoefficients(residual, q)

		for i := 0; i < predictionLength; i++ {
			prediction := float64(0)
			for j := 1; j <= q; j++ {
				if len(residual) >= j {
					prediction += maCoefficients[j-1] * residual[len(residual)-j]
				}
			}
			residual = append(residual, prediction)
		}
	}

	// Lakukan inverse differencing jika d > 0
	if d > 0 {
		for i := 0; i < d; i++ {
			inverseDiffData := make([]float64, len(residual))
			inverseDiffData[0] = data[len(data)-1]
			for j := 1; j < len(residual); j++ {
				inverseDiffData[j] = inverseDiffData[j-1] + residual[j-1]
			}
			residual = inverseDiffData
		}
	}

	// Lakukan inverse AR jika p > 0
	if p > 0 {
		residual = inverseARIMA(residual, p, d, arCoefficients)
	}

	return residual[len(residual)-predictionLength:]
}

// Fungsi untuk melakukan inverse ARIMA
func inverseARIMA(residual []float64, p, d int, arCoefficients []float64) []float64 {
	n := len(residual)
	data := make([]float64, n)

	copy(data, residual)

	// Lakukan inverse differencing jika d > 0
	if d > 0 {
		for i := n - 1; i >= d; i-- {
			data[i] = data[i] + data[i-d]
		}
	}

	// Lakukan inverse AR jika p > 0
	if p > 0 {
		for i := n - 1; i >= p; i-- {
			prediction := float64(0)
			for j := 0; j < p; j++ {
				prediction += arCoefficients[j] * data[i-j-1]
			}
			data[i-p] = data[i-p] + prediction
		}
	}

	return data
}

func main() {
	// Data contoh (gantilah data ini dengan data Anda)
	data := []float64{
		14, 45, 71, 76, 89, 28, 48, 57, 46, 60, 45,
		65, 110, 72, 69, 57, 48, 103, 40, 51, 33, 56,
		56, 67, 76, 22, 39, 86, 54, 32, 91, 68, 72,
		70, 52, 75, 91, 39, 49, 60, 81, 59, 43, 75,
		69, 71, 46, 50, 55, 51, 73, 80, 73,
	}

	// Parameter ARIMA (gantilah parameter ini sesuai kebutuhan)
	p := 1 // Order autoregressive (AR)
	d := 0 // Derajat differencing (I)
	q := 1 // Order moving average (MA)

	// Panjang prediksi (dalam minggu)
	predictionLength := 12 // 12 minggu

	// Lakukan prediksi
	predictions := predictARIMA(data, p, d, q, predictionLength)

	// Tampilkan hasil data asli
	fmt.Println("Data Asli:")
	for i, val := range data {
		fmt.Printf("Minggu %d: %.2f\n", i+1, val)
	}

	// Tampilkan hasil prediksi
	fmt.Println("Hasil Prediksi (Bentuk Asli):")
	for i, pred := range predictions {
		fmt.Printf("Minggu %d: %.2f\n", len(data)+i+1, pred+data[len(data)-1])
	}
}
