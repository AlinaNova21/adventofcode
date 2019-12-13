package point

type Point2D struct {
	X, Y int
}

func New2D(x, y int) Point2D {
	return Point2D{x, y}
}

func (p Point2D) ID() int {
	return 1e6*p.X + p.Y
}

func (p1 Point2D) Add(p2 Point2D) Point2D {
	return Point2D{p1.X + p2.X, p1.Y + p2.Y}
}

type Point3D struct {
	X, Y, Z float64
}

func New3D(x, y, z float64) Point3D {
	return Point3D{x, y, z}
}

func (p Point3D) ID() float64 {
	return 1e12*p.X + 1e6*p.Y + p.Z
}
