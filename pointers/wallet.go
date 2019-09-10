// https://github.com/quii/learn-go-with-tests/blob/master/pointers-and-errors.md

package main

import "fmt"
import "errors"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// 構造体の値を変更するのでポインタ関数として振る舞いを定義
func (w *Wallet) Deposit(amount Bitcoin) error {
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
