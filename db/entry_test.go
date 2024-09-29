package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Edojonsnow/bank/utils"
	"github.com/stretchr/testify/require"
)

type CreateEntryParams struct {
	AccountID    int64 `json:"account_id"`
	Amount  int64  `json:"amount"`
}
type UpdateEntryParams struct {
	AccountID    int64 `json:"account_id"`
	Amount  int64  `json:"amount"`
}

type ListEntryParams struct {	
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}


func FetchRandomAccount(t *testing.T) Account{
	account , err := testQueries.GetAccount(context.Background(),utils.RandomInt(1, 12) )
	require.NoError(t, err)
	require.NotEmpty(t, account)
	
	return account
}

// func fetchRandomEntry(t *testing.T) Entry{
// 	entry , err := testQueries.GetEntry(context.Background(), utils.RandomInt(1, 3))
//     require.NoError(t, err)
//     require.NotEmpty(t, entry)
    
//     return entry

// }

func createRandomEntry(t *testing.T) Entry{
	account := FetchRandomAccount(t)
    accountID := account.ID

    arg := CreateEntryParams{
        AccountID:    accountID,
        Amount:  utils.RandomInt(1, 1000),
    }

    entry , err := testQueries.CreateEntry(context.Background(), arg.AccountID, arg.Amount)
    require.NoError(t , err)
    require.NotEmpty(t, entry)

    return entry
}

func TestCreateEntry(t *testing.T){
	
	account := FetchRandomAccount(t)
	accountID := account.ID

	arg := CreateEntryParams{
		AccountID:    accountID,
        Amount:  utils.RandomInt(100, 1000),
	}

	entry , err := testQueries.CreateEntry(context.Background(), arg.AccountID, arg.Amount)
	require.NoError(t , err)
	require.NotEmpty(t, entry)

}



func TestGetEntry(t *testing.T){
	entry := createRandomEntry(t)

    entry2 , err := testQueries.GetEntry(context.Background(), entry.ID)
    require.NoError(t , err)
    require.NotEmpty(t, entry2)
	require.Equal(t, entry.AccountID , entry2.AccountID)
	require.Equal(t, entry.Amount , entry2.Amount)


}

func TestDelete(t *testing.T){

    // entry := fetchRandomEntry(t)
	entry := createRandomEntry(t)


    err := testQueries.DeleteEntry(context.Background(), entry.ID)
    require.NoError(t, err)

    entry2 , err := testQueries.GetEntry(context.Background(), entry.ID)
    require.Error(t , err)
    require.EqualError(t, err, sql.ErrNoRows.Error())
    require.Empty(t, entry2)
}

func TestUpdateEntry(t *testing.T){
	entry :=  createRandomEntry(t)

	arg := UpdateEntryParams{
		AccountID:  entry.ID,
        Amount:  utils.RandomInt(100, 1000),
	}
	entry2 , err := testQueries.UpdateEntry(context.Background(), arg.AccountID, arg.Amount)
	require.NoError(t, err)
	require.NotEmpty(t, entry2 )
	require.Equal(t,arg.Amount, entry2.Amount)

}

func TestListEntries(t *testing.T){
	for i := 0 ; i > 10 ; i ++{
		createRandomEntry(t)
	}
	arg := ListEntryParams{
        Limit:  5,
        Offset: 5,
    }

	entries, err := testQueries.ListEntries(context.Background(), arg.Limit, arg.Offset)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _ , entry := range entries {
        require.NotEmpty(t, entry)     
	}


}