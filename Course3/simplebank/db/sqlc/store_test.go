package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	// checking balance first
	fmt.Println("before transaction: ", account1.Balance, account2.Balance)

	// run a concurrent transfer transaction
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResults)

	for i := 0; i < n; i++ {
		// creating name for each transaction
		// txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			// ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// check account's balance

		fmt.Println("tx: ", fromAccount.Balance, toAccount.Balance)
		balance1 := account1.Balance - fromAccount.Balance // diff is amt going out of acc1
		balance2 := toAccount.Balance - account2.Balance   // diff is amt going in acc2
		require.Equal(t, balance1, balance2)               // if transaction crct then both diff must be same
		require.True(t, balance1 > 0)                      // also diff shouldn't be 0 as amt is going out
		require.True(t, balance1%amount == 0)              // must be divisible by the amt going out

		k := int(balance1 / amount) // this will give the no. of transactions, hence must be divisble
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true

	}

	// check updated account balance
	updatedAcc1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAcc2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println("after transaction: ", updatedAcc1.Balance, updatedAcc2.Balance)
	require.Equal(t, account1.Balance-int64(n)*amount, updatedAcc1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAcc2.Balance)
}

// func to check deadlock situation
func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	// checking balance first
	fmt.Println("before transaction: ", account1.Balance, account2.Balance)

	// run a concurrent transfer transaction
	n := 10
	amount := int64(10)

	// only checking for deadlock errors, not results
	errs := make(chan error)

	for i := 0; i < n; i++ {
		// creating name for each transaction

		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	// check errors
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check updated account balance
	updatedAcc1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAcc2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println("after transaction: ", updatedAcc1.Balance, updatedAcc2.Balance)
	require.Equal(t, account1.Balance, updatedAcc1.Balance)
	require.Equal(t, account2.Balance, updatedAcc2.Balance)
}
