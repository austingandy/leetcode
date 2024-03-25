package main

type Employee struct {
	ID               string
	Name             string
	EmploymentType   string
	DOB              string
	IdentifiedGender string
	Title            string
	Compensation     Compensation
	Department       string
	WorkLocation     string
	StartDate        string
}

type Compensation struct {
	CompensationType string
	Currency         string
	Amount           float64
}

func aggregateMetric[T comparable, U any](
	employees []Employee,
	getKey func(Employee) T,
	getValue func(Employee) U,
	aggregate func(U, U) U,
) map[T]U {
	m := make(map[T]U)
	for _, e := range employees {
		m[getKey(e)] = aggregate(m[getKey(e)], getValue(e))
	}
	return m
}

func SumComp(employees []Employee) map[string]float64 {
	return aggregateMetric(employees, getDept, getComp, sum)
}

func AvgComp(employees []Employee) map[string]float64 {
	return aggregateMetric(employees, getDept, getComp, makeAverage())
}

func getComp(e Employee) float64 {
	switch e.Compensation.CompensationType {
	case "HOURLY":
		return e.Compensation.Amount * 40 * 52
	case "WEEKLY":
		return e.Compensation.Amount * 52
	case "ANNUAL":
		return e.Compensation.Amount
	case "MONTHLY":
		return e.Compensation.Amount * 12
	}
	return 0
}

func getDept(e Employee) string {
	return e.Department
}

func sum(a, b float64) float64 {
	return a + b
}

func makeAverage() func(a, b float64) float64 {
	count := 0
	return func(a, b float64) float64 {
		count += 1
		return ((a * float64(count-1)) + b) / float64(count)
	}
}
