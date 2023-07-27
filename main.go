package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/google/uuid"
)

type Field struct {
	name string
	Type FieldType
	Size int
}

type (
	FieldType string
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numberRunes = []rune("1234567890")
)

const (
	UUID      FieldType = "UUID"
	Number    FieldType = "NUMBER"
	Date      FieldType = "DATE"
	Timestamp FieldType = "TIMESTAMP"
	Text      FieldType = "TEXT"
	Float     FieldType = "FLOAT"

	MAX_LINES = 170
	// MAX_LINES = 17000
)

func main() {
	fields := make(map[int]Field, 10)
	i := 0

	i = addFieldToMap(fields, "paymentId", Text, 34, i)
	i = addFieldToMap(fields, "TID", UUID, 34, i)
	i = addFieldToMap(fields, "NSU", Number, 10, i)
	i = addFieldToMap(fields, "Date", Date, 10, i)
	i = addFieldToMap(fields, "PaymentDate", Date, 10, i)
	i = addFieldToMap(fields, "AuthCode", Number, 16, i)
	i = addFieldToMap(fields, "CNPJ", Number, 14, i)
	i = addFieldToMap(fields, "ValorTransacao", Float, 5, i)
	i = addFieldToMap(fields, "ValorParcela", Float, 5, i)
	i = addFieldToMap(fields, "Bandeira", Text, 14, i)
	i = addFieldToMap(fields, "CNPJ", Text, 14, i)

	f, err := os.Create("saida.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(generateHeader(fields))
	if err2 != nil {
		log.Fatal(err2)
	}

	for i := 0; i < MAX_LINES; i++ {
		_, err := f.WriteString(generateLine(fields))
		if err != nil {
			log.Fatal(err)
		}
	}

}

func addFieldToMap(fields map[int]Field, name string, Type FieldType, size int, index int) int {
	fields[index] = Field{
		name: name,
		Type: Type,
		Size: size,
	}
	return index + 1
}

func generateHeader(fields map[int]Field) string {
	var result string
	for i := 0; i < len(fields); i++ {
		result += fmt.Sprint(fields[i].name, ";")
	}
	result += "\n"
	return result
}

func generateLine(fields map[int]Field) string {
	var result string
	for i := 0; i < len(fields); i++ {
		result += fmt.Sprint(generateField(fields[i].Type, fields[i].Size), ";")
	}
	result += "\n"
	return result
}

func generateField(Type FieldType, size int) string {

	switch Type {
	case UUID:
		value, _ := uuid.NewRandom()
		return value.String()

	case Number:
		result := make([]rune, size)

		for i := range result {
			result[i] = numberRunes[rand.Intn(len(numberRunes))]
		}

		return string(result)

	case Text:
		result := make([]rune, size)

		for i := range result {
			result[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(result)

	case Date:
		day := rand.Intn(30) + 1
		month := rand.Intn(9) + 1

		return fmt.Sprintf("%d/0%d/2023", day, month)

	case Timestamp:
		return "2006-01-02T15:04:05.000-0700"

	case Float:
		intValue := make([]rune, size)
		decimalValue := make([]rune, 2)

		for i := range intValue {
			intValue[i] = numberRunes[rand.Intn(len(numberRunes))]
		}
		for i := range decimalValue {
			decimalValue[i] = numberRunes[rand.Intn(len(numberRunes))]
		}

		return fmt.Sprintf("%s.%s", string(intValue), string(decimalValue))

	}

	return "ERROR"
}
