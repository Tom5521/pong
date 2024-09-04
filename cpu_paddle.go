package main

type CPUPaddle struct {
	Paddle
}

func (c *CPUPaddle) Update(ballY float) {
	center := c.Y + c.Height/2
	if center > ballY {
		c.Y -= f(c.Speed)
	}
	if center < ballY {
		c.Y += f(c.Speed)
	}
	c.limitMovement()
}
