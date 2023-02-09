package algorithmfoo

import "math"

// 区域内点相交算法

// Coordinate right-handed coordinate system
type Coordinate struct {
	X int
	Y int
}

type Size struct {
	Width  uint
	Height uint
}

type Unit interface {
	Width() uint
	Height() uint
}

type RenderContext interface {
	Unit
	Coordinate() Coordinate
	// CoincidenceCheck
	// @param  RenderContext
	// @return RenderContext coincidence context
	//         bool          is coincidence
	CoincidenceCheck(RenderContext) (RenderContext, bool)
}

type BasicContext struct {
	c Coordinate
	s Size
}

func NewBasicContext(s Size) BasicContext {
	return BasicContext{s: s}
}

func (ctx *BasicContext) SetCoordinate(c Coordinate) {
	ctx.c = c
}

func (c BasicContext) Size() Size {
	return c.s
}

func (c BasicContext) Width() uint {
	return c.s.Width
}

func (c BasicContext) Height() uint {
	return c.s.Height
}

func (c BasicContext) Coordinate() Coordinate {
	return c.c
}

// ┌──────────┐
// │ Terminal │
// │     ┌───────────┐
// │     │ Container │
// └─────│           │
//
//	│ * Point   │
//	└───────────┘
//
// Point is in Container but not in terminal.
// Thats why Container also need coincidence check.
// @param ctx sth to render's ctx
func (c BasicContext) CoincidenceCheck(ctx RenderContext) (RenderContext, bool) {
	newCtx := BasicContext{}
	renderStartX := ctx.Coordinate().X
	renderEndX := renderStartX + int(ctx.Width())
	if renderEndX <= c.c.X || c.c.X+int(c.Width()) <= renderStartX {
		// outer left || outer right
		return newCtx, false
	} else {
		newCtx.c.X = int(math.Max(float64(c.c.X), float64(renderStartX)))
	}

	renderStartY := ctx.Coordinate().Y
	renderEndY := renderStartY + int(ctx.Height())
	if renderEndY <= c.c.Y || c.c.Y+int(c.Height()) <= renderStartY {
		// outer left || outer right
		return newCtx, false
	} else {
		newCtx.c.Y = int(math.Max(float64(c.c.Y), float64(renderStartY)))
	}

	newCtx.s.Width = uint(math.Min(float64(c.c.X+int(c.Width())), float64(renderEndX)) - float64(newCtx.c.X))
	newCtx.s.Height = uint(math.Min(float64(c.c.Y+int(c.Height())), float64(renderEndY)) - float64(newCtx.c.Y))
	return newCtx, true
}