package opengl_exercise

import (
	"math"
)

type Vector struct {
	X, Y, Z float32
}

func (v Vector) Normalize() Vector {
	length := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
	if length < 0.001 {
		return Vector{}
	}

	return Vector{
		X: v.X / length,
		Y: v.Y / length,
		Z: v.Z / length,
	}
}

func (v Vector) AddVector(vector Vector) Vector {
	return Vector{
		X: v.X + vector.X,
		Y: v.Y + vector.Y,
		Z: v.Z + vector.Z,
	}
}

func (v Vector) Negative() Vector {
	return Vector{X: -v.X, Y: -v.Y, Z: -v.Z}
}

func (v Vector) MultiplyScalar(scalar float32) Vector {
	return Vector{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
	}
}

func (v Vector) MultiplyVector(vector Vector) Vector {
	return Vector{
		X: v.X * vector.X,
		Y: v.Y * vector.Y,
		Z: v.Z * vector.Z,
	}
}

func (v Vector) MultiplyMatrix(m *Matrix) Vector {
	return Vector{
		X: v.X*(*m)[0] + v.Y*(*m)[4] + v.Z*(*m)[8] + (*m)[12],
		Y: v.X*(*m)[1] + v.Y*(*m)[5] + v.Z*(*m)[9] + (*m)[13],
		Z: v.X*(*m)[2] + v.Y*(*m)[6] + v.Z*(*m)[10] + (*m)[14],
	}
}

func (lhs Vector) Cross(rhs Vector) Vector {
	return Vector{
		X: lhs.Y*rhs.Z - lhs.Z*rhs.Y,
		Y: lhs.Z*rhs.X - lhs.X*rhs.Z,
		Z: lhs.X*rhs.Y - lhs.Y*rhs.X,
	}
}

func (lhs Vector) Dot(rhs Vector) float32 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}

type Camera struct {
	Position   Vector
	Rotation   Vector
	viewMatrix Matrix
}

func NewCamera() *Camera {
	return &Camera{}
}

func (c *Camera) ViewMatrix() Matrix {
	return c.viewMatrix
}

func (c *Camera) Render() {
	// upward
	up := Vector{Y: 1}

	// where camera is looking by default
	lookAt := Vector{Z: 1}

	// set rotation in radians
	rotation := c.Rotation.MultiplyScalar(math.Pi / 180)

	// create rotation matrix from yaw, pitch and roll values
	rotationMatrix := c.matrixRotationYawPitchRoll(rotation.Y, rotation.X, rotation.Z)

	// transform the lookat and up vector by the rotation matrix so the view is correctly rotated at the origin
	up = up.MultiplyMatrix(&rotationMatrix)
	lookAt = lookAt.MultiplyMatrix(&rotationMatrix)

	// translate rotated camera position to the location of the viewer
	lookAt = lookAt.AddVector(c.Position)

	// finally create view matrix from the three updated vectors
	c.viewMatrix = c.buildViewMatrix(c.Position, lookAt, up)
}

func (c *Camera) matrixRotationYawPitchRoll(yaw, pitch, roll float32) Matrix {
	cYaw := float32(math.Cos(float64(yaw)))
	cPitch := float32(math.Cos(float64(pitch)))
	cRoll := float32(math.Cos(float64(roll)))

	sYaw := float32(math.Sin(float64(yaw)))
	sPitch := float32(math.Sin(float64(pitch)))
	sRoll := float32(math.Sin(float64(roll)))

	// calculate yaw, pitch, roll rotation matrix
	return Matrix{
		cRoll*cYaw + sRoll*sPitch*sYaw,
		sRoll * cPitch,
		-sYaw*cRoll + sRoll*sPitch*cYaw,
		0,

		-sRoll*cYaw + cRoll*sPitch*sYaw,
		cRoll * cPitch,
		sRoll*sYaw + cRoll*sPitch*cYaw,
		0,

		cPitch * sYaw,
		-sPitch,
		cPitch * cYaw,
		0,

		0,
		0,
		0,
		1,
	}
}

func (c *Camera) buildViewMatrix(position, lookAt, up Vector) Matrix {
	zAxis := lookAt.AddVector(position.Negative()).Normalize()
	xAxis := up.Cross(zAxis).Normalize()
	yAxis := zAxis.Cross(xAxis)

	return Matrix{
		xAxis.X, yAxis.X, zAxis.X, 0,
		xAxis.Y, yAxis.Y, zAxis.Y, 0,
		xAxis.Z, xAxis.Z, zAxis.Z, 0,
		-xAxis.Dot(position), yAxis.Dot(position), zAxis.Dot(position), 1,
	}
}
