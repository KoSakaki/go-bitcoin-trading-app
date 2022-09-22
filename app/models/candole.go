package models

import (
	"fmt"
	"gotrading/bitflyer"
	"log"
	"time"
)

type Candole struct {
	ProductCode string        `json:"product_code"`
	Duration    time.Duration `json:"duration"`
	Time        time.Time     `json:"time"`
	Open        float64       `json:"open"`
	Close       float64       `json:"close"`
	High        float64       `json:"high"`
	Low         float64       `json:"low"`
	Volume      float64       `json:"volume"`
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

func GetAllCandole(productCode string, duration time.Duration, limit int) (dfCandole *DataFrameCandole, err error) {
	tableName := GetCandoleTableName(productCode, duration)
	cmd := fmt.Sprintf(`SELECT * FROM (
		SELECT time, open, close, high, low, volume FROM %s ORDER BY time DESC LIMIT ?
	) ORDER BY time ASC;`, tableName)
	rows, err := DbConnection.Query(cmd, limit)
	if err != nil {
		return
	}
	defer rows.Close()
	dfCandole = &DataFrameCandole{}
	dfCandole.Productcode = productCode
	dfCandole.Duration = duration
	for rows.Next() {
		var candole Candole
		candole.Duration = duration
		rows.Scan(&candole.Time, &candole.Open, &candole.Close, &candole.High, &candole.Low, &candole.Volume)
		dfCandole.Candoles = append(dfCandole.Candoles, candole)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return dfCandole, nil
}
