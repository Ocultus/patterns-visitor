package figures

type Circle struct {
	Radius  float64
	Dot 	Dot
}

func (c *Circle) Accept(v Visitor) {
    v.VisitForCircle(c)
}

func (c *Circle) GetType() string {
    return "Circle"
}