package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	btc := Bitcoin(10)
	wallet.Deposite(btc)
	AssertBalance(t, wallet, Bitcoin(10))
}

func TestWithdraw(t *testing.T) {
	t.Run("Withdraw normal", func(t *testing.T) {
		wallet := Wallet{Bitcoin(100)}
		e := wallet.Withdraw(Bitcoin(10))
		AssertBalance(t, wallet, Bitcoin(90))
		AssertNoError(t, e)
	})

	t.Run("Withdraw Insufficent", func(t *testing.T) {
		wallet := Wallet{Bitcoin(5)}
		AssertError(t, wallet.Withdraw(Bitcoin(10)), errorrrInsufficientFunds)
		AssertBalance(t, wallet, Bitcoin(5))
	})
}

func AssertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Either got no error, or the error aint matching")
	}
}

func AssertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("Error Withdrawing")
	}
}
