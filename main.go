package main

import (
	"go/projectile_tracer/tracer"
)

func main() {

	// projectile = tracer.Projectile{x, y, velocity, angle, duration}
	projectile := tracer.Projectile{0, 0, 1, 0, 5}

	tracer.Shot(projectile)
}
