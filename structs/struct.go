package structs

func Perimeter(rectanlge Rectangle) float64 {
	return (rectanlge.Width + rectanlge.Height) * 2
}

func Area(rectanlge Rectangle) float64 {
	return rectanlge.Width * rectanlge.Height
}
