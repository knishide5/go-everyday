package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
	  if got != want {
		  t.Errorf("got %s want %s", got, want)
	  }
	}

	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
	  wallet := Wallet{}
	  wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
    wallet := Wallet{Bitcoin(30)}
	  wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(20))
	})

  t.Run("Withdraw insufficient funds", func(t *testing.T) {
    wallet := Wallet{Bitcoin(10)}
	  err := wallet.Withdraw(Bitcoin(30))

		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err)
	})
}
