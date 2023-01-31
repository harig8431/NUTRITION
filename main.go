package main

func main()  {
	
	ns:=GetNutritionalScore(NutritionalData{
		Energy:EnergyFromKcal(),
		Sugars:SugarGram(),
		SatturatedFattyAcids:SatturatedFattyAcidsGram(),
		Sodium:Sodiummilligram(),
		Fruits:FruitsPercent(),
		Fibre:Fibregram(),
		Protein:ProteinGram(),
	},Food)

	fmt.Println("Nutritional Score: %d",ns.Value)
}