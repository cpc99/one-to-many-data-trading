package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
)

type RunResults struct {
	n           float64
	BU1v1       float64
	OU1v1       float64
	SU1v1       float64
	Total1v1    float64
	BUnogame    float64
	OUnogame    float64
	SUnogame    float64
	Totalnogame float64
	BUgame      float64
	OUgame      float64
	SUgame      float64
	Totalgame   float64
	x           float64
	pb          float64
	ps          float64
}

func main() {
	Bi := 20.0
	Cb := 1.0
	Cs := 2.0
	Co := 2.0
	Con := 3.0

	x := make([]float64, 101)
	Pb := make([]float64, 101)
	Ps := make([]float64, 101)
	x[0] = 1.0
	Pb[0] = 10.0
	Ps[0] = 3.0
	n := 1.0

	results := make([]*RunResults, 100)
	for i := 0; i < 100; i++ {
		runResults := new(RunResults)
		runResults.n = n

		BU1 := n*Bi*math.Log(1+x[0]) - Pb[0]*x[0]*n - Cb*x[0]*n
		OU1 := (Pb[0]*x[0] - Ps[0]*x[0] - Co*x[0]) * n
		SU1 := Ps[0]*x[0]*n - Cs*x[0]
		/*
			fmt.Printf("BU1v1 equals %v:\n", BU1)
			fmt.Printf("OU1v1 equals %v:\n", OU1)
			fmt.Printf("SU1v1 equals %v:\n", SU1)
		*/
		runResults.BU1v1 = BU1
		runResults.OU1v1 = OU1
		runResults.SU1v1 = SU1
		runResults.Total1v1 = BU1 + OU1 + SU1

		BU2 := n*Bi*math.Log(1+x[0]) - Pb[0]*x[0]*n - Cb*x[0]*n
		OU2 := Pb[0]*x[0]*n - Ps[0]*x[0]*n - Con*x[0]
		SU2 := Ps[0]*x[0]*n - Cs*x[0]
		/*
			fmt.Printf("BU(no game)%v equals %v:\n", i, BU2)
			fmt.Printf("OU(no game)%v equals %v:\n", i, OU2)
			fmt.Printf("SU(no game)%v equals %v:\n", i, SU2)
		*/
		runResults.BUnogame = BU2
		runResults.OUnogame = OU2
		runResults.SUnogame = SU2
		runResults.Totalnogame = BU2 + OU2 + SU2

		m := n*Cb + Co
		s := m + (1.0/2.0)*Cs
		a := n * n * n
		b := 3*n*n*m - (1/4)*n*n*n*Bi
		c := 3*n*m*m - s*n*n*Bi
		d := m*m*m - s*s*n*Bi
		A := b*b - 3*a*c
		B := b*c - 9*a*d
		C := c*c - 3*b*d
		Z := B*B - 4*A*C
		if A == B && B == 0 {
			//x1 := -c / b
			x2 := -c / b
			//x3 := -c / b
			Ps[i+1] = x2
			//fmt.Printf("1x1 equals %v:   x2 equals %v:  x3 equals %v:\n", x1, x2, x3)
		} else if Z > 0 {
			Y1 := A*b + 1.5*a*(-B+math.Sqrt(Z))
			Y2 := A*b + 1.5*a*(-B-math.Sqrt(Z))
			x1 := (-b - (math.Cbrt(Y1) + math.Cbrt(Y2))) / (3 * a)
			/*
				p2 := (-b + 0.5*(math.Cbrt(Y1)+math.Cbrt(Y2))) / (3 * a)
				q2 := (0.5 * math.Cbrt(3) * (math.Cbrt(Y1) - math.Cbrt(Y2))) / (3 * a)
				var x2 complex128 = complex(p2, q2)
				p3 := (-b + 0.5*(math.Cbrt(Y1)+math.Cbrt(Y2))) / (3 * a)
				q3 := -(0.5 * math.Cbrt(3) * (math.Cbrt(Y1) - math.Cbrt(Y2))) / (3 * a)
				var x3 complex128 = complex(p3, q3)
				fmt.Printf("2x1 equals %v:   x2 equals %v:  x3 equals %v:\n", x1, x2, x3)
			*/
			Ps[i+1] = x1
		} else if Z == 0 {
			//x1 := B/A - b/a
			x2 := -B / (2 * A)
			//x3 := -B / (2 * A)
			//fmt.Printf("3x1 equals %v:   x2 equals %v:  x3 equals %v:\n", x1, x2, x3)
			Ps[i+1] = x2
		} else if Z < 0 {
			Y1 := (2*A*b - 3*a*B) / (2 * A * math.Sqrt(A))
			Y2 := math.Acos(Y1)
			//x1 := (-b - 2*math.Sqrt(A)*math.Cos(Y2/3)) / (3 * a)
			x2 := (-b + math.Sqrt(A)*(math.Cos(Y2/3)+math.Sqrt(3)*math.Sin(Y2/3))) / (3 * a)
			//x3 := (-b + math.Sqrt(A)*(math.Cos(Y2/3)-math.Sqrt(3)*math.Sin(Y2/3))) / (3 * a)
			//fmt.Printf("4x1 equals %v:   x2 equals %v:  x3 equals %v:\n", x1, x2, x3)
			Ps[i+1] = x2
		}

		Pb[i+1] = math.Sqrt(n*Bi*(n*Cb+n*Ps[i+1]+Co))/n - Cb

		x[i+1] = n*Bi/((Cb+Pb[i+1])*n) - 1
		/*
			fmt.Printf("x%v equals %v:   Pb%v equals %v:  Ps%v equals %v:\n", i+1, x[i+1], i+1, Pb[i+1], i+1, Ps[i+1])
		*/
		runResults.x = x[i+1]
		runResults.pb = Pb[i+1]
		runResults.ps = Ps[i+1]

		BU3 := n*Bi*math.Log(1+x[i+1]) - Pb[i+1]*x[i+1]*n - Cb*x[i+1]*n
		OU3 := Pb[i+1]*x[i+1]*n - Ps[i+1]*x[i+1]*n - Con*x[i+1]
		SU3 := Ps[i+1]*x[i+1]*n - Cs*x[i+1]
		/*
			fmt.Printf("BU(game)%v equals %v:\n", i, BU3)
			fmt.Printf("OU(game)%v equals %v:\n", i, OU3)
			fmt.Printf("SU(game)%v equals %v:\n", i, SU3)
		*/
		runResults.BUgame = BU3
		runResults.OUgame = OU3
		runResults.SUgame = SU3
		runResults.Totalgame = BU3 + OU3 + SU3

		results[i] = runResults
		n = n + 1.0
	}
	printResults(results)
}

func printResults(results []*RunResults) {
	filename := fmt.Sprintf("%v.csv", "game")
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("can not create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		"n", "BU(1v1)", "OU(1v1)", "SU(1v1)", "BU(no game)", "OU(no game)", "SU(no game)",
		"BU(game)", "OU(game)", "SU(game)", "Total(1v1)", "Total(no game)", "Total(game)", "x", "pb", "ps"})
	for _, res := range results {
		writer.Write([]string{
			fmt.Sprintf("%.4f", res.n),
			fmt.Sprintf("%.4f", res.BU1v1),
			fmt.Sprintf("%.4f", res.OU1v1),
			fmt.Sprintf("%.4f", res.SU1v1),
			fmt.Sprintf("%.4f", res.BUnogame),
			fmt.Sprintf("%.4f", res.OUnogame),
			fmt.Sprintf("%.4f", res.SUnogame),
			fmt.Sprintf("%.4f", res.BUgame),
			fmt.Sprintf("%.4f", res.OUgame),
			fmt.Sprintf("%.4f", res.SUgame),
			fmt.Sprintf("%.4f", res.Total1v1),
			fmt.Sprintf("%.4f", res.Totalnogame),
			fmt.Sprintf("%.4f", res.Totalgame),
			fmt.Sprintf("%.4f", res.x),
			fmt.Sprintf("%.4f", res.pb),
			fmt.Sprintf("%.4f", res.ps)})
	}
	return
}
