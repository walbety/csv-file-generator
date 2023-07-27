package main

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/csv-file-generator/canonical"
	"github.com/walbety/csv-file-generator/config"
	"math/rand"
	"os"
	"time"
)

type Field struct {
	name string
	Type canonical.FieldType
	Size int
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numberRunes = []rune("1234567890")
)

func main() {
	err := config.InitConfigs()
	if err != nil {
		log.WithError(err).Error("Error at initConfig")
	}
	fields := make(map[int]Field, 10)
	i := 0

	for _, field := range config.Cfg.Fields {
		fieldType, ok := canonical.MapFileTypeStringToConst[field.Type]
		if ok {
			i = addFieldToMap(fields, field.Name, fieldType, field.Size, i)
		} else {
			log.Errorf("This type is unsuported: %s", field.Type)
		}
	}

	filename := "saida.csv"
	if config.Cfg.Filename != "" {
		filename = config.Cfg.Filename
	}
	f, err := os.Create(fmt.Sprintf("%s.csv", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(generateHeader(fields))
	if err2 != nil {
		log.Fatal(err2)
	}

	for i := 0; i < config.Cfg.TotalLines; i++ {
		_, err := f.WriteString(generateLine(fields))
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Info("file created!!")
	time.Sleep(time.Second * 3)

}

func addFieldToMap(fields map[int]Field, name string, Type canonical.FieldType, size int, index int) int {
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

func generateField(Type canonical.FieldType, size int) string {

	switch Type {
	case canonical.UUID:
		value, _ := uuid.NewRandom()
		return value.String()

	case canonical.Number:
		result := make([]rune, size)

		for i := range result {
			result[i] = numberRunes[rand.Intn(len(numberRunes))]
		}

		return string(result)

	case canonical.Text:
		result := make([]rune, size)

		for i := range result {
			result[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(result)

	case canonical.Date:
		day := rand.Intn(30) + 1
		month := rand.Intn(9) + 1

		return fmt.Sprintf("%d/0%d/2023", day, month)

	case canonical.Timestamp:
		return "2006-01-02T15:04:05.000-0700"

	case canonical.Float:
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
