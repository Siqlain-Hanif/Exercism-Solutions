package lasagna

func PreparationTime(layers []string, avgPreparationTime int) int{
    if (avgPreparationTime==0){
        avgPreparationTime=2;
    }
    return len(layers) *avgPreparationTime 
}
func Quantities(layers []string) (noodlesNeeded int, sauceNeeded float64){
    // sauceNeeded:=0.0
    // noodlesNeeded:=0
    for i:=0; i<len(layers); i++ {
        if(layers[i]=="noodles"){
            noodlesNeeded+=50
        }
        if(layers[i]=="sauce"){
            sauceNeeded+=0.2
        }
    }
    return
}

func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
func ScaleRecipe(quantities  []float64, scale int) []float64{
    scaledQuantities :=[]float64{}
    otherPortionScale := float64(scale)/2.0
    for i:=0;i<len(quantities);i++{
        scaledQuantities=append(scaledQuantities,quantities[i] * otherPortionScale)
    }
    return scaledQuantities
}