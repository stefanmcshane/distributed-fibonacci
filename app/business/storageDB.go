package business

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math/big"
	"strings"
)

const (
	dbTableFibonacci  = "fibonacci"
	dbColumnOrdinal   = "ordinal"
	dbColumnFibonacci = "fibonacci"
)

type BigInt struct {
	big.Int
}

func (b *BigInt) Value() (driver.Value, error) {
	if b != nil {
		return b.String(), nil
	}
	return nil, nil

}
func (b *BigInt) Scan(value interface{}) error {
	var i sql.NullString
	if err := i.Scan(value); err != nil {
		return err
	}
	if _, ok := b.SetString(i.String, 10); ok {
		return nil
	}
	return fmt.Errorf("could not scan type %T into BigInt", value)
}

func (s DBStore) AddFib(ctx context.Context, tup []ordinalTuple) error {

	q := fmt.Sprintf("INSERT INTO %s (%s,%s) VALUES ??? ON CONFLICT DO NOTHING", dbTableFibonacci, dbColumnOrdinal, dbColumnFibonacci)
	var valTups strings.Builder
	for idx, tu := range tup {
		if idx == len(tup)-1 {
			valTups.WriteString(fmt.Sprintf("( %s, %s )", tu.Ordinal.String(), tu.Fib.String()))
			continue
		}
		valTups.WriteString(fmt.Sprintf("( %s, %s ),", tu.Ordinal.String(), tu.Fib.String()))
	}

	query := strings.Replace(q, "???", valTups.String(), 1)

	_, err := s.DB.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (s DBStore) CheckFib(ctx context.Context, ordinal big.Int) (big.Int, error) {
	q := fmt.Sprintf("SELECT %s, %s FROM %s WHERE %s = $1", dbColumnOrdinal, dbColumnFibonacci, dbTableFibonacci, dbColumnOrdinal)
	rows, err := s.DB.Query(q, ordinal.String())
	if err != nil {
		return big.Int{}, err
	}
	defer rows.Close()

	var fibs []ordinalTuple
	for rows.Next() {
		var ord string
		var fib string
		err := rows.Scan(&ord, &fib)
		if err != nil {
			return big.Int{}, err
		}
		f, _ := big.NewInt(0).SetString(fib, 10)
		o, _ := big.NewInt(0).SetString(ord, 10)
		fibs = append(fibs, ordinalTuple{Ordinal: *o, Fib: *f})
	}
	if rows.Err() != nil || len(fibs) == 0 {
		return big.Int{}, NewError(ordinal, ErrorStorageOrdinalNotFound)
	}
	return fibs[0].Fib, nil
}

func (s DBStore) Len() int {
	return 0
}

func (s DBStore) ResultsUnder(ctx context.Context, fibValue big.Int) int {
	q := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE CAST(%s as BIGINT) < %s", dbTableFibonacci, dbColumnFibonacci, fibValue.String())
	row := s.DB.QueryRow(q)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return -1
	}
	return count
}

func (s DBStore) Clear(ctx context.Context) error {
	q := fmt.Sprintf("TRUNCATE TABLE %s", dbTableFibonacci)
	_, err := s.DB.Exec(q)
	if err != nil {
		return err
	}
	fmt.Println("deleted rows from database")
	return nil
}

func (s DBStore) HighestOrdinalStored(ctx context.Context) (big.Int, error) {
	q := fmt.Sprintf("SELECT %s, %s FROM %s order by %s desc limit 1", dbColumnOrdinal, dbColumnFibonacci, dbTableFibonacci, dbColumnOrdinal)
	rows, err := s.DB.Query(q)
	if err != nil {
		return big.Int{}, err
	}
	defer rows.Close()

	var fibs []ordinalTuple
	for rows.Next() {
		var ord string
		var fib string
		err := rows.Scan(&ord, &fib)
		if err != nil {
			return big.Int{}, err
		}
		f, _ := big.NewInt(0).SetString(fib, 10)
		o, _ := big.NewInt(0).SetString(ord, 10)
		fibs = append(fibs, ordinalTuple{Ordinal: *o, Fib: *f})
	}
	if rows.Err() != nil || len(fibs) == 0 {
		return big.Int{}, NewError(*big.NewInt(0), ErrorNoHighestOrdinal)
	}
	return fibs[0].Ordinal, nil
}
