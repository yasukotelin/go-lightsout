package main

type Cell struct {
	IsLightOn bool
}

const (
	LightsOn  = " !"
	LightsOut = " ."
)

func (c *Cell) String() string {
	if c.IsLightOn {
		return LightsOn
	}
	return LightsOut
}
