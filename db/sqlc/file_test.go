package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/SabariGanesh-K/21BPS1209_Backend.git/util"
)

// func createRandomUser(t *testing.T) Users {
// 	hashedPassword, err := util.HashPassword(util.RandomString(6))
// 	require.NoError(t, err)

// 	arg := CreateUserParams{
// 		Username:       util.RandomString(15),
// 		HashedPassword: hashedPassword,
// 		FullName:       util.RandomString(15),
// 		Email:          util.RandomEmail(),
// 	}

// 	user, err := testQueries.CreateUser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user)

// 	require.Equal(t, arg.Username, user.Username)
// 	require.Equal(t, arg.HashedPassword, user.HashedPassword)
// 	require.Equal(t, arg.FullName, user.FullName)
// 	require.Equal(t, arg.Email, user.Email)
// 	require.True(t, user.PasswordChangedAt.IsZero())
// 	require.NotZero(t, user.CreatedAt)

// 	return user
// }

 func createRandomFile(t *testing.T) Files {
	user:= createRandomUser(t)
	
	arg := CreateFileParams{
		ID      :   util.RandomString(15),
	Owner     :  user.Username,
	S3Url    :   util.RandomString(15),
	FileName :   util.RandomString(5),
	CreatedAt :  time.Now(),
	Filesize  : "2mb",
	Filetype  : "jpg" ,
		
	}

	file, err := testQueries.CreateFile(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, file)

	require.Equal(t, arg.ID, file.ID)
	require.Equal(t, arg.Owner, file.Owner)
	require.Equal(t, arg.S3Url, file.S3Url)
	require.Equal(t, arg.FileName, file.FileName)
	require.Equal(t, arg.Filesize, file.Filesize)
	require.Equal(t, arg.Filetype, file.Filetype)

	require.NotZero(t, file.CreatedAt)


	return file
}

func TestCreateFile(t *testing.T) {
	createRandomFile(t)
}

// func TestSearchFile(t *testing.T) {
// 	file1 := createRandomFile(t)
// 	arg := SearchFileParams{
// 		ID:file1.ID,
	
		
		
// 	}
// 	file2, err := testQueries.SearchFile(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, file2)

// 	require.Equal(t, file1.ID, file2.ID)
// 	require.Equal(t, file1.Owner, file2.Owner)
// 	require.Equal(t, file1.S3Url, file2.S3Url)
// 	require.Equal(t, file1.FileName, file2.FileName)
// 	require.Equal(t, file1.Filesize, file2.Filesize)
// 	require.Equal(t, file1.Filetype, file2.Filetype)

// }

// func TestUpdateUserOnlyFullName(t *testing.T) {
// 	oldUser := createRandomUser(t)

// 	newFullName := util.RandomString(15)
// 	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
// 		Username: oldUser.Username,
// 		FullName: sql.NullString{
// 			String: newFullName,
// 			Valid:  true,
// 		},
// 	})

// 	require.NoError(t, err)
// 	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
// 	require.Equal(t, newFullName, updatedUser.FullName)
// 	require.Equal(t, oldUser.Email, updatedUser.Email)
// 	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
// }

// func TestUpdateUserOnlyEmail(t *testing.T) {
// 	oldUser := createRandomUser(t)

// 	newEmail := util.RandomEmail()
// 	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
// 		Username: oldUser.Username,
// 		Email: sql.NullString{
// 			String: newEmail,
// 			Valid:  true,
// 		},
// 	})

// 	require.NoError(t, err)
// 	require.NotEqual(t, oldUser.Email, updatedUser.Email)
// 	require.Equal(t, newEmail, updatedUser.Email)
// 	require.Equal(t, oldUser.FullName, updatedUser.FullName)
// 	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
// }

// func TestUpdateUserOnlyPassword(t *testing.T) {
// 	oldUser := createRandomUser(t)

// 	newPassword := util.RandomString(6)
// 	newHashedPassword, err := util.HashPassword(newPassword)
// 	require.NoError(t, err)

// 	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
// 		Username: oldUser.Username,
// 		HashedPassword: sql.NullString{
// 			String: newHashedPassword,
// 			Valid:  true,
// 		},
// 	})

// 	require.NoError(t, err)
// 	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
// 	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
// 	require.Equal(t, oldUser.FullName, updatedUser.FullName)
// 	require.Equal(t, oldUser.Email, updatedUser.Email)
// }

// func TestUpdateUserAllFields(t *testing.T) {
// 	oldUser := createRandomUser(t)

// 	newFullName := util.RandomString(15)
// 	newEmail := util.RandomEmail()
// 	newPassword := util.RandomString(6)
// 	newHashedPassword, err := util.HashPassword(newPassword)
// 	require.NoError(t, err)

// 	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
// 		Username: oldUser.Username,
// 		FullName: sql.NullString{
// 			String: newFullName,
// 			Valid:  true,
// 		},
// 		Email: sql.NullString{
// 			String: newEmail,
// 			Valid:  true,
// 		},
// 		HashedPassword: sql.NullString{
// 			String: newHashedPassword,
// 			Valid:  true,
// 		},
// 	})

// 	require.NoError(t, err)
// 	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
// 	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
// 	require.NotEqual(t, oldUser.Email, updatedUser.Email)
// 	require.Equal(t, newEmail, updatedUser.Email)
// 	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
// 	require.Equal(t, newFullName, updatedUser.FullName)
// }
