package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if(kind=="car"||kind=="truck"){return true}
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	compare := option1 > option2
    if(compare==true){
        return option2 + " is clearly the better choice."
    }
    return option1 + " is clearly the better choice."

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var devaluePercentage float64;
	if(age<3){
        devaluePercentage = 0.80;
    }else if (age>=3&&age<10){
        devaluePercentage = 0.70;
    }else{
        devaluePercentage = 0.50;
    }
    return originalPrice * devaluePercentage
}
