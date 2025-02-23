package internal

const (
	MinWorkers             = 2
	MaxWorkers             = 7
	MinQueueCapacity       = 10
	MaxQueueCapacity       = 25
	MinApplicationInterval = 1
	MaxApplicationInterval = 10
	MinServingDuration     = 2
	MaxServingDuration     = 30
	MinProfitRange         = 3
	MaxProfitRange         = 50
	MinModelingStep        = 10
	MaxModelingStep        = 60
	MinLunchDuration       = 0
	MaxLunchDuration       = 60

	NormalDistribution  = "normal distribution"
	UniformDistribution = "uniform distribution"

	WorkerSalary = 2
)

func ValidateWorkers(val int) bool {
	return val >= MinWorkers && val <= MaxWorkers
}

func ValidateQueueCapacity(val int) bool {
	return val >= MinQueueCapacity && val <= MaxQueueCapacity
}

func ValidateApplicationInterval(val int) bool {
	return val >= MinApplicationInterval && val <= MaxApplicationInterval
}

func ValidateServingDuration(left, right int) bool {
	return left >= MinServingDuration && left <= MaxServingDuration &&
		right >= MinServingDuration && right <= MaxServingDuration && left <= right
}

func ValidateProfitRange(left, right int) bool {
	return left >= MinProfitRange && left <= MaxProfitRange &&
		right >= MinProfitRange && right <= MaxProfitRange && left <= right
}

func ValidateModelingStep(val int) bool {
	return val >= MinModelingStep && val <= MaxModelingStep
}

func ValidateLunchDuration(val int) bool {
	return val >= MinLunchDuration && val <= MaxLunchDuration
}
