package entities

type MaterialType string

const (
    Bottle  MaterialType = "bottle"  // Measured by units
    Can     MaterialType = "can"     // Measured by units
    Plastic MaterialType = "plastic" // Measured by weight (kg)
    Paper   MaterialType = "paper"   // Measured by weight (kg)
)

type UnitType string

const (
    Unit  UnitType = "unit" 
    Weight UnitType = "kg"   
)

type Material struct {
    ID         int
    Type       MaterialType 
    Unit       UnitType    
    Quantity   float64      
    Reward     float64      
}
