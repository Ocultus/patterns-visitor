package figures

type Square struct {
	Dot Dot
    Side  float64
}

func (s *Square) Accept(v Visitor) {
    v.VisitForSquare(s)
}

func (s *Square) GetType() string {
    return "Square"
}