package c3dt

import "github.com/EliCDavis/vector/vector3"

// https://github.com/CesiumGS/3d-tiles/tree/main/specification#box
//
// The boundingVolume.box property is an array of 12 numbers that define an
// oriented bounding box in a right-handed 3-axis (x, y, z) Cartesian
// coordinate system where the z-axis is up. The first three elements define
// the x, y, and z values for the center of the box. The next three elements
// (with indices 3, 4, and 5) define the x-axis direction and half-length. The
// next three elements (indices 6, 7, and 8) define the y-axis direction and
// half-length. The last three elements (indices 9, 10, and 11) define the
// z-axis direction and half-length.
type BoxBoundingVolume []float64

func (bbv BoxBoundingVolume) Center() vector3.Float64 {
	return vector3.New(bbv[0], bbv[1], bbv[2])
}

func NewBoxBoundingVolume(center, xAxis, yAxis, zAxis vector3.Float64) BoxBoundingVolume {
	return BoxBoundingVolume{
		center.X(), center.Y(), center.Z(),
		xAxis.X(), xAxis.Y(), xAxis.Z(),
		yAxis.X(), yAxis.Y(), yAxis.Z(),
		zAxis.X(), zAxis.Y(), zAxis.Z(),
	}
}

// An array of six numbers that define a bounding geographic region in
// EPSG:4979 coordinates with the order [west, south, east, north, minimum
// height, maximum height]. Longitudes and latitudes are in radians. The range
// for latitudes is [-PI/2,PI/2]. The range for longitudes is [-PI,PI]. The
// value that is given as the 'south' of the region shall not be larger than
// the value for the 'north' of the region. The heights are in meters above
// (or below) the WGS84 ellipsoid. The 'minimum height' shall not be larger
// than the 'maximum height'.
type RegionBoundingVolume []float64

// An array of four numbers that define a bounding sphere. The first three
// elements define the x, y, and z values for the center of the sphere. The
// last element (with index 3) defines the radius in meters. The radius shall
// not be negative.
type SphereBoxBoundingVolume []float64

func (bbv SphereBoxBoundingVolume) Center() vector3.Float64 {
	return vector3.New(bbv[0], bbv[1], bbv[2])
}

func (bbv SphereBoxBoundingVolume) Radius() float64 {
	return bbv[3]
}

func NewSphereBoxBoundingVolume(center vector3.Float64, radius float64) SphereBoxBoundingVolume {
	return SphereBoxBoundingVolume{center.X(), center.Y(), center.Z(), radius}
}

// A bounding volume that encloses a tile or its content. At least one bounding
// volume property is required. Bounding volumes include `box`, `region`, or
// `sphere`.
type BoundingVolume struct {
	Box    BoxBoundingVolume       `json:"box,omitempty"`
	Region RegionBoundingVolume    `json:"region,omitempty"`
	Sphere SphereBoxBoundingVolume `json:"sphere,omitempty"`
}
