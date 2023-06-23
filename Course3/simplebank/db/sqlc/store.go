package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all funcs to execute db transactions and queries
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// var txKey = struct{}{} // creating object of type struct to pass to the context.WithValue() method

// execTx executes a func within a db transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbError)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains ip parameters of a transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResults contains the result of a transfer transaction
type TransferTxResults struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from 1 acc to another
// It creates a transfer record, add acc entries and update acc balance within 1 transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error) {
	var result TransferTxResults

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// txName := ctx.Value(txKey) // getting back the transaction name

		// fmt.Println(txName, "create transfer")
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// fmt.Println(txName, "create entry 1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// fmt.Println(txName, "create entry 2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {

			// // update acc balance
			// // getaccount -> update its balance   ; but this is often done incorrectly w/o proper locking mech
			// // fmt.Println(txName, "get account 1")
			// // account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
			// // if err != nil {
			// // 	return err
			// // }

			// // fmt.Println(txName, "update account 1")
			// result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			// 	ID:     arg.FromAccountID,
			// 	Amount: -arg.Amount,
			// })
			// if err != nil {
			// 	return err
			// }

			// // fmt.Println(txName, "get account 2")
			// // account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
			// // if err != nil {
			// // 	return err
			// // }

			// // fmt.Println(txName, "update account 2")
			// result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			// 	ID:     arg.ToAccountID,
			// 	Amount: arg.Amount,
			// })
			// if err != nil {
			// 	return err
			// }

			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
			if err != nil {
				return err
			}
		} else { // update toAccount before fromAccount
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}

func addMoney(ctx context.Context, q *Queries, accountID1 int64, amt1 int64, accountID2 int64, amt2 int64) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amt1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amt2,
	})
	return
}
