package repo

import (
	"currency-rate/internal/usecase"
	"database/sql"
	"fmt"
)

var (
	tableNameCurrency = "test.name_currency"
	tableDateCurrency = "test.date_currency"
)

type CurrencyDB struct {
	db *sql.DB
}

func NewCurrencyDB(db *sql.DB) *CurrencyDB {
	return &CurrencyDB{db: db}
}

func (c CurrencyDB) GetCurrencyDate(date string, charName string) (*usecase.CurrencyDTO, error) {
	item := usecase.CurrencyDTO{}
	query := fmt.Sprintf(`SELECT date_start,date_end,
       value, char_name, name FROM %s as t RIGHT JOIN %s as l on t.currency_id = l.currency_id 
                                                                    WHERE char_name = $1 
                                                                    and to_char(date_start, 'yyyy/mm/dd') = $2`,
		tableDateCurrency, tableNameCurrency)
	row := c.db.QueryRow(query, charName, date)
	if err := row.Scan(&item.StartDate, &item.EndDate, &item.Value, &item.CharCode, &item.Name); err != nil {
		return &item, err
	}
	return &item, nil
}

func (c CurrencyDB) InsertCurrencyDate(value []usecase.CurrencyDTO) error {
	var itemId int
	tx, err := c.db.Begin()
	//tx.Prepare() TODO
	if err != nil {
		return err
	}

	for _, v := range value {
		createNameCurrency := fmt.Sprintf(`INSERT INTO %s (char_name, name) VALUES ($1, $2) RETURNING currency_id`,
			tableNameCurrency)

		row := tx.QueryRow(createNameCurrency, v.CharCode, v.Name)
		if err = row.Scan(&itemId); err != nil {
			return tx.Rollback()
		}

		createDateCurrency := fmt.Sprintf(`INSERT INTO %s (date_start, date_end, value, currency_id) VALUES ($1, $2, $3, $4)`, tableDateCurrency)
		_, err = tx.Exec(createDateCurrency, v.StartDate, v.EndDate, v.Value, itemId)
		if err != nil {
			return tx.Rollback()
		}
	}

	return tx.Commit()
}
