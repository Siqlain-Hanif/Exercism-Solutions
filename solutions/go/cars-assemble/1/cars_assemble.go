package cars

func FailedProductionRate (productionRate int, successRate float64) float64{
    failureRate := (100.0-successRate)/100.0;
    return failureRate* float64(productionRate)
}
// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    failedProductionRate := FailedProductionRate(productionRate,successRate)
	return float64(productionRate)- failedProductionRate;    
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    carsPerHour:=CalculateWorkingCarsPerHour(productionRate, successRate);
	 
    return int(carsPerHour/60.0)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const groupCost,individualCost = 95000,10000
    individualCars := int(carsCount%10);
    groupCars := int(carsCount/10);
    return uint(groupCars*groupCost + individualCars*individualCost)
}
