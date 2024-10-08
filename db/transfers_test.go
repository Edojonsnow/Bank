package db

import (
	"context"
	"testing"

	"github.com/Edojonsnow/bank/utils"
	"github.com/stretchr/testify/require"
)

type CreateTransferParams struct {

	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`

}


type ListTransferParams struct{
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}


func createRandomTransfer(t  *testing.T) Transfer{
	account1 := FetchRandomAccount(t)
	account2 := FetchRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID ,
		ToAccountID: account2.ID,
		Amount: utils.RandomInt(100, 1000),
	}

	transfer , err := testQueries.CreateTransfer(context.Background(), arg.FromAccountID, arg.ToAccountID, arg.Amount)
	require.NoError(t , err)
	require.NotEmpty(t, transfer)

	return transfer
}



func TestCreateTransfers(t *testing.T){

	createRandomTransfer(t)

}

func TestGetTransfer(t *testing.T){
	transfer :=  createRandomTransfer(t)

	transfer2 , err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t , err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	

}

func TestListTransfers(t *testing.T){
	

	arg := ListTransferParams{
		Limit:  5,
        Offset: 5,
	}
	transfers,err := testQueries.ListTransfers(context.Background(), arg.Limit, arg.Offset)
	require.NoError(t, err)
	require.Len(t, transfers, 5)


}