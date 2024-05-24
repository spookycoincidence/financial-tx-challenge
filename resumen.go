package resumen

type Resumen struct {
	SaldoTotal             float64
	NumTransaccionesPorMes map[string]int
	PromedioDebitoPorMes   map[string]float64
	PromedioCreditoPorMes  map[string]float64
}
