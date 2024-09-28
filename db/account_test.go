package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Edojonsnow/bank/utils"
	"github.com/stretchr/testify/require"
)

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}
type UpdateAccountParams struct {
	ID    int64 
	Balance  int64  
}

type ListAccountsParams struct {	
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}



func createRandomAccount(t *testing.T) Account{

	arg := CreateAccountParams {
		Owner: utils.RandomOwner(),
        Balance: utils.RandomBalance(),
        Currency: utils.RandomCurrency(),
	}

	account , err := testQueries.CreateAccount(context.Background(),arg.Owner, arg.Balance, arg.Currency)
	require.NoError(t , err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner , account.Owner)
	require.Equal(t, arg.Balance , account.Balance)
	require.Equal(t, arg.Currency , account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account

}

func TestCreateAccount(t *testing.T){

	createRandomAccount(t)
}

func TestGetAccount(t *testing.T){
	account1 := createRandomAccount(t)
	account2 , err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t , err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID , account2.ID)
	require.Equal(t, account1.Owner , account2.Owner)
	require.Equal(t, account1.Balance , account2.Balance)
	require.Equal(t, account1.Currency , account2.Currency)
	require.WithinDuration(t, account1.CreatedAt , account2.CreatedAt , time.Second)
}

func TestUpdateAccount(t *testing.T){
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:    account1.ID,
        Balance: utils.RandomBalance(),
	}

	account2 , err := testQueries.UpdateAccount(context.Background(), arg.ID, arg.Balance)
	require.NoError(t , err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID , account2.ID)
	require.Equal(t, account1.Owner , account2.Owner)
	require.Equal(t, arg.Balance , account2.Balance)
	require.Equal(t, account1.Currency , account2.Currency)


}

func TestDeleteAccount(t *testing.T){
	account1 := createRandomAccount(t)

    err := testQueries.DeleteAccount(context.Background(), account1.ID)
    require.NoError(t , err)

    account2 , err := testQueries.GetAccount(context.Background(), account1.ID)
    require.Error(t , err)
    require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T){
	for i := 0 ; i > 10 ; i ++{
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
        Limit:  5,
        Offset: 5,
	}

	accounts , err := testQueries.ListAccounts(context.Background(), arg.Limit , arg.Offset)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _ , account := range accounts {
        require.NotEmpty(t, account)     
	}
}