package streamtest

import (
	"github.com/Nik-U/pbc"
	"log"
)

var params = pbc.GenerateA(160, 512)
var pairing = params.NewPairing()
var g = pairing.NewG1().Rand()
var h = pairing.NewG1().Rand()

type Ra struct {
	r1 *pbc.Element
	r2 *pbc.Element
	r3 *pbc.Element
	r4 *pbc.Element
	rS *pbc.Element
}

func ZKPSum() bool {
	ra := Ra{}
	ra.r1 = pairing.NewZr().NewFieldElement()
	ra.r2 = pairing.NewZr().NewFieldElement()
	ra.r3 = pairing.NewZr().NewFieldElement()
	ra.r4 = pairing.NewZr().NewFieldElement()
	ra.rS = pairing.NewZr().Rand()
	y, cm, cmr1 := P1(ra)
	c := V1()
	zr := P2(ra, c)
	res := V2(cm, cmr1, y, c, zr)
	return res

}

func P1(ra Ra) (int32, []*pbc.Element, *pbc.Element) {
	ix := make([]int32, 4)
	ix[0] = 1
	ix[1] = 2
	ix[2] = 1
	ix[3] = 2
	y := ix[0] + ix[1] + ix[2] + ix[3]
	q := pairing.NewG1().NewFieldElement()
	p := pairing.NewG1().NewFieldElement()
	x := make([]*pbc.Element, 4)
	r := make([]*pbc.Element, 4)
	cm := make([]*pbc.Element, 4)

	for i := 0; i < 4; i++ {
		res := ZKPNN(ix[i])
		if !res {
			log.Printf("x为负\n ")
		}
		x[i] = pairing.NewZr().NewFieldElement()
		x[i].SetInt32(ix[i])
		r[i] = pairing.NewZr().Rand()
		cm[i] = pairing.NewG1().NewFieldElement()
		p.PowZn(g, x[i]) //g^x[i]
		q.PowZn(h, r[i]) //h^r[i]
		cm[i].Mul(p, q)  //cm[i]=g^x[i]*h^r[i]
	}
	ra.r1.Set(r[0])
	ra.r2.Set(r[1])
	ra.r3.Set(r[2])
	ra.r4.Set(r[3])

	x0 := pairing.NewZr().NewFieldElement()
	cmr1 := pairing.NewG1().NewFieldElement()
	x0.SetInt32(0)
	cmr1.PowZn(g, x0) //g^0
	q.PowZn(h, ra.rS) //h^r'
	cmr1.Mul(cmr1, q) //cmr1=g^0*h^r'

	return y, cm, cmr1
}

func V1() *pbc.Element {
	c := pairing.NewZr().Rand()
	return c
}

func P2(ra Ra, c *pbc.Element) *pbc.Element {
	zr := pairing.NewZr().NewFieldElement()
	sum := pairing.NewZr().NewFieldElement()
	sum.Add(ra.r1, ra.r2)
	sum.Add(sum, ra.r3)
	sum.Add(sum, ra.r4) //r1+r2+r3+r4
	sum.Mul(c, sum)     //c*(r1+r2+r3+r4)
	zr.Add(ra.rS, sum)  //zr=r'+c*(r1+r2+r3+r4)
	return zr
}

func V2(cm []*pbc.Element, cmr1 *pbc.Element, y int32, c *pbc.Element, zr *pbc.Element) bool {
	res := false
	ey := pairing.NewZr().NewFieldElement()
	ey.SetInt32(y)

	v := pairing.NewZr().NewFieldElement()
	q1 := pairing.NewG1().NewFieldElement()
	q2 := pairing.NewG1().NewFieldElement()
	left := pairing.NewG1().NewFieldElement()
	q1.PowZn(g, v.Mul(c, ey)) //g^(c*ey)
	q2.PowZn(h, zr)           //h^zr
	left.Mul(q1, q2)          //left=g^(c*ey)*h^zr

	sum := pairing.NewG1().NewFieldElement()
	right := pairing.NewG1().NewFieldElement()
	sum.Mul(cm[0], cm[1])
	sum.Mul(sum, cm[2])
	sum.Mul(sum, cm[3]) //sum=cm1*cm2*cm3*cm4
	sum.PowZn(sum, c)
	right.Mul(cmr1, sum) //right=cmr1*sum

	if left.Equals(right) {
		res = true
	}

	return res
}
