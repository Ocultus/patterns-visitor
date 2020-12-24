package figures
type Rectangle struct {
    SideX  float64
	SideY float64
	Dot Dot
}

func (t *Rectangle) Accept(v Visitor) {
    v.VisitForRectangle(t)
}

func (t *Rectangle) GetType() string {
    return "rectangle"
}