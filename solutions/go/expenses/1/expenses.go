package expenses
import "fmt"
// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var result []Record
    for _, record := range in {
        if predicate(record) {
            result = append(result, record)
        }
    }
    return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	 return func(r Record) bool {
        return r.Day >= p.From && r.Day <= p.To
    }
}

func ByCategory(category string) func(Record) bool {
    return func(r Record) bool {
        return r.Category == category
    }
}

// 4. TotalByPeriod - calculates sum of expenses in a period
func TotalByPeriod(records []Record, period DaysPeriod) float64 {
    filtered := Filter(records, ByDaysPeriod(period))
    total := 0.0
    for _, record := range filtered {
        total += record.Amount
    }
    return total
}

// 5. CategoryExpenses - calculates total for a category in a period with error handling
func CategoryExpenses(records []Record, period DaysPeriod, category string) (float64, error) {
    // First, check if the category exists at all in the records (regardless of period)
    categoryExists := false
    for _, record := range records {
        if record.Category == category {
            categoryExists = true
            break
        }
    }
    
    if !categoryExists {
        return 0, fmt.Errorf("unknown category %s", category)
    }
    
    // Filter by both period and category
    filtered := Filter(records, func(r Record) bool {
        return ByDaysPeriod(period)(r) && ByCategory(category)(r)
    })
    
    // Calculate total
    total := 0.0
    for _, record := range filtered {
        total += record.Amount
    }
    
    return total, nil
}