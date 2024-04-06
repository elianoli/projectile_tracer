package tracer

import (
	"fmt"
	"math"
	"time"
)

type Projectile struct {
	X, Y, Velocity, Angle float64
	Duration              int
}

// starting values
var x, y, velocity, angle float64

// velocity components
var vx, vy float64

func Shot(newProjectile Projectile) {
	QuickSet(
		newProjectile.X,
		newProjectile.Y,
		newProjectile.Velocity,
		newProjectile.Angle)

	Tracer(newProjectile.Duration)
}

func Tracer(duration int) {

	fmt.Print("\nStarting tracing...\n")

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			// receives completed channel
			case <-done:
				return

			// receives ticker channel
			case <-ticker.C:
				// breaking velocity into vx and vy components and the limiting them to two decimal points
				vx = math.Cos(angle*math.Pi/180) * velocity
				vx = math.Round(vx*100) / 100
				vy = math.Sin(angle*math.Pi/180) * velocity
				vy = math.Round(vy*100) / 100

				// deslocation of x and y
				x += vx
				y += vy

				fmt.Printf("Velocity: %v <=> {%.2v, %.2v}\t", velocity, vx, vy)
				fmt.Printf("Tracing: {%v, %v}\n", x, y)
			}
		}
	}()

	// wait duration before closing the channels
	time.Sleep(time.Duration(duration) * time.Second)
	ticker.Stop()
	done <- true

	fmt.Print("Finishing tracing!\n")

}

func Info() {
	fmt.Printf("Position: (%v, %v)\nVelocity(Units / Second): %v\nAngle(In Degrees): %vº\n", x, y, velocity, angle)
}

func QuickSet(newX, newY, newVelocity, newAngle float64) {
	SetPosition(newX, newY)
	SetVelocity(newVelocity)
	SetAngle(newAngle)

	Info()
}

func SetPosition(newX, newY float64) {
	x = newX
	y = newY
}

func SetVelocity(newVelocity float64) {
	velocity = newVelocity
}

func SetAngle(newAngle float64) {
	angle = newAngle
}
