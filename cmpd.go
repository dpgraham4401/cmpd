package main

import (
	"flag"
	"fmt"
	"math"
)

type InterestData struct {
    initialAmount float64
    rate float64
    frequency float64
    years int
    compoundInt float64
    totalAmount float64
}

type Compound interface {
    ComputeCompound() float64
}

func (d *InterestData) ComputeCompound() {
    d.totalAmount = d.initialAmount * math.Pow((1 + (d.rate/100)/d.frequency), d.frequency * float64(d.years)) 
    d.compoundInt = d.totalAmount - d.initialAmount
}


func main() {
    var data = InterestData{}
    flag.Float64Var(&data.initialAmount, "p", 0, "Principle (initial) amount")
    flag.Float64Var(&data.rate, "r", 0, "Rate of (annual) interest")
    flag.Float64Var(&data.frequency, "f", 1, "Compound frequency per annum")
    flag.IntVar(&data.years, "y", 5, "Number of years")
    flag.Parse()
    data.ComputeCompound()

    fmt.Println("=================================================")
    fmt.Printf("Total Amount: %.2f    Compound Int.: %.2f\n", 
    data.totalAmount, data.compoundInt)
    fmt.Println("=================================================")
    fmt.Printf("Initial Amount: %.2f\n",data.initialAmount) 
    fmt.Printf("Interest Rate: %.2f%%\n",data.rate) 
    fmt.Printf("Numbers of Years: %d\n",data.years) 
}
