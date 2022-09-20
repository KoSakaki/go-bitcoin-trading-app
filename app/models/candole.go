package models

import (
	"fmt"
	"gotrading/bitflyer"
	"log"
	"time"
)

type Candole struct {
	ProductCode string
	Duration    time.Duration
	Time        time.Time
	Open        float64
	Close       float64
	High        float64
	Low         float64
	Volume      float64
}

func NewCandole(productCode string, duration time.Duration, timeDate time.Time, open, close, high, low, volume float64) *Candole {
	return &Candole{
		productCode,
		duration,
		timeDate,
		open,
		close,
		high,
		low,
		volume,
	}
}

func (c *Candole) TableName() string {
	return GetCandoleTableName(c.ProductCode, c.Duration)
}

func (c *Candole) Create() error {
	cmd := fmt.Sprintf("INSERT INTO %s(time, open, close, high, low, volume) VALUES (?, ?, ?, ?, ?, ?)", c.TableName())
	_, err := DbConnection.Exec(cmd, c.Time.Format(time.RFC3339), c.Open, c.Close, c.High, c.Low, c.Volume)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
func (c *Candole) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET open = ?, close = ?, high = ?, low = ?, volume = ? WHERE time = ?", c.TableName())
	_, err := DbConnection.Exec(cmd, c.Open, c.Close, c.High, c.Low, c.Volume, c.Time.Format(time.RFC3339))
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func GetCandole(productCode string, duration time.Duration, dateTime time.Time) *Candole {
	tableName := GetCandoleTableName(productCode, duration)
	cmd := fmt.Sprintf("SELECT time, open, close, high, low, volume FROM %s WHERE time = ?", tableName)
	row := DbConnection.QueryRow(cmd, dateTime.Format(time.RFC3339))
	var candole Candole
	err := row.Scan(&candole.Time, &candole.Open, &candole.Close, &candole.High, &candole.Low, &candole.Volume)
	if err != nil {
		// log.Println(err)
		return nil
	}
	return NewCandole(productCode, duration, candole.Time, candole.Open, candole.Close, candole.High, candole.Low, candole.Volume)
}

func CreateCandleWithDuration(ticker bitflyer.Ticker, productCode string, duration time.Duration) bool {
	currentCandole := GetCandole(productCode, duration, ticker.TruncateDateTime(duration))
	price := ticker.GetMidPrice()
	if currentCandole == nil {
		candole := NewCandole(productCode, duration, ticker.TruncateDateTime(duration), price, price, price, price, ticker.Volume)
		candole.Create()
		return true
	}

	if currentCandole.High <= price {
		currentCandole.High = price
	} else if currentCandole.Low >= price {
		currentCandole.Low = price
	}
	currentCandole.Volume += ticker.Volume
	currentCandole.Close = price
	currentCandole.Save()
	return false
}
