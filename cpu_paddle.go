package main

type CPUPaddle struct {
	Paddle
}

func (c *CPUPaddle) Update(ballY float32) {
	center := c.Y + c.Height/2
	if center > ballY {
		c.Y -= float32(c.Speed)
	}
	if center <= ballY {
		c.Y += float32(c.Speed)
	}
	c.limitMovement()
}
