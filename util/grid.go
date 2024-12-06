package util

import "image"

type Point = image.Point
type Grid map[Point]string

func Left(p Point) Point {
	return p.Add(image.Pt(-1, 0))
}

func Right(p Point) Point {
	return p.Add(image.Pt(1, 0))
}

func Up(p Point) Point {
	return p.Add(image.Pt(0, -1))
}

func Down(p Point) Point {
	return p.Add(image.Pt(0, 1))
}

func UpLeft(p Point) Point {
	return Up(Left(p))
}

func UpRight(p Point) Point {
	return Up(Right(p))
}

func DownLeft(p Point) Point {
	return Down(Left(p))
}

func DownRight(p Point) Point {
	return Down(Right(p))
}
