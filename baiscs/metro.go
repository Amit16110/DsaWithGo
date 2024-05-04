package baiscs

import (
	"fmt"
	"sort"
)

type MetroCard struct {
	Number  string
	Balance int
}

type Passenger struct {
	MetroCard     *MetroCard
	PassengerType string
}

type Station struct {
	Name            string
	CollectedAmount int
	PassengerCounts map[string]int
}

type Journey struct {
	Passenger *Passenger
	Station   *Station
}

var TravelCharges = map[string]int{
	"ADULT":          200,
	"SENIOR_CITIZEN": 100,
	"KID":            50,
}

func NewMetroCard(number string, balance int) *MetroCard {
	return &MetroCard{
		Number:  number,
		Balance: balance,
	}
}

func (p *Passenger) CheckIn(station *Station) {
	charge := TravelCharges[p.PassengerType]

	if station.PassengerCounts[p.PassengerType] > 0 {
		charge /= 2
	}

	p.MetroCard.Balance -= charge
	station.CollectedAmount += charge
	station.PassengerCounts[p.PassengerType]++
}

func PrintSummary(stations []*Station) {
	for _, station := range stations {
		fmt.Printf("TOTAL_COLLECTION %s %d %d\n", station.Name, station.CollectedAmount, CalculateTotalDiscount(station))
		PrintPassengerSummary(station.PassengerCounts)
	}
}

func CalculateTotalDiscount(station *Station) int {
	totalDiscount := 0

	for passengerType, count := range station.PassengerCounts {
		discount := TravelCharges[passengerType] / 2
		totalDiscount += discount * count
	}

	return totalDiscount
}

func PrintPassengerSummary(passengerCounts map[string]int) {
	type passengerCount struct {
		PassengerType string
		Count         int
	}

	passengerCountList := make([]passengerCount, 0, len(passengerCounts))

	for passengerType, count := range passengerCounts {
		passengerCountList = append(passengerCountList, passengerCount{PassengerType: passengerType, Count: count})
	}

	sort.Slice(passengerCountList, func(i, j int) bool {
		if passengerCountList[i].Count == passengerCountList[j].Count {
			return passengerCountList[i].PassengerType < passengerCountList[j].PassengerType
		}
		return passengerCountList[i].Count > passengerCountList[j].Count
	})

	for _, pc := range passengerCountList {
		fmt.Printf("%s %d\n", pc.PassengerType, pc.Count)
	}
}
