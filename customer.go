package main

import (
	"encoding/json"
	"math"
	"strconv"
)

type customer struct {
	Lat      string `json:"latitude"`
	Long     string `json:"longitude"`
	Name     string `json:"name"`
	UserID   int64  `json:"user_id"`
	location location
}

func (c *customer) unmarshal(bytes []byte) error {
	if err := json.Unmarshal(bytes, c); err != nil {
		return err
	}
	if err := c.setLocation(); err != nil {
		return err
	}
	return nil
}

func (c *customer) setLocation() error {
	lat, err := strconv.ParseFloat(c.Lat, 64)
	if err != nil {
		return err
	}
	long, err := strconv.ParseFloat(c.Long, 64)
	if err != nil {
		return err
	}
	c.location = location{lat, long}
	return nil
}

// Location represents GPS coordinates for any location
type location struct {
	lat, long float64
}

func (l *location) toRadian() (lat, long float64) {
	return radian(l.lat), radian(l.long)
}

func radian(degree float64) float64 {
	return degree * math.Pi / 180
}
