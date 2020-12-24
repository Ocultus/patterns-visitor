package figures

import (
	"math"
)

type Visitor interface {
	VisitForSquare(*Square) float64
	VisitForCircle(*Circle) float64
	VisitForRectangle(*Rectangle) float64
}

type AreaCalc struct {
	Area float64
}

func (p *AreaCalc) VisitForCircle(c *Circle) float64 {
	return c.Radius * c.Radius * math.Pi
}

func (p* AreaCalc) VisitForSquare(s *Square) float64{
	return s.Side * s.Side
}

func (p* AreaCalc) VisitForRectangle(r *Rectangle) float64{
	return r.SideX * r.SideY 
}

func New() *AreaCalc{
	return &AreaCalc{
		Area: 0,
	}
}