package services

import (
	"context"
	"time"

	"github.com/Chirag711/go-rest-api/config"
	"github.com/Chirag711/go-rest-api/models"
	"github.com/xuri/excelize/v2"

	"go.mongodb.org/mongo-driver/bson"
)

func GenerateUsersExcel() (*excelize.File, error) {

	f := excelize.NewFile()
	sheet := "Users"

	f.SetSheetName("Sheet1", sheet)

	// Header Row
	f.SetCellValue(sheet, "A1", "Name")
	f.SetCellValue(sheet, "B1", "Email")
	f.SetCellValue(sheet, "C1", "Age")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := config.GetCollection("users")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	row := 2

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)

		f.SetCellValue(sheet, "A"+string(rune(row+48)), user.Name)
		f.SetCellValue(sheet, "B"+string(rune(row+48)), user.Email)
		f.SetCellValue(sheet, "C"+string(rune(row+48)), user.Age)

		row++
	}

	return f, nil
}
