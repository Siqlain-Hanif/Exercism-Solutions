package thefarm
import ("fmt"
       "errors")
func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
    // Get the total fodder amount for all cows
    fodderAmount, err := fc.FodderAmount(numCows)
    if err != nil {
        return 0, err
    }
    
    // Get the fattening factor
    fatteningFactor, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    
    // Calculate food per cow: (totalFodder * factor) / numCows
    foodPerCow := (fodderAmount * fatteningFactor) / float64(numCows)
    
    return foodPerCow, nil
}
func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
    if numCows <= 0 {
        return 0, errors.New("invalid number of cows")
    }
    
    return DivideFood(fc, numCows)
}
// InvalidCowsError is a custom error type for invalid cow counts
type InvalidCowsError struct {
    numCows int
    message string
}

// Error implements the error interface
func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

// ValidateNumberOfCows validates the number of cows
func ValidateNumberOfCows(numCows int) error {
    if numCows < 0 {
        return &InvalidCowsError{
            numCows: numCows,
            message: "there are no negative cows",
        }
    }
    
    if numCows == 0 {
        return &InvalidCowsError{
            numCows: numCows,
            message: "no cows don't need food",
        }
    }
    
    return nil
}
