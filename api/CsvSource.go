package api

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func InsertCsv(w http.ResponseWriter, r *http.Request) {
	Table_Name := r.URL.Query().Get("Table_Name")
	content, err := io.ReadAll(r.Body)
	if err != nil {
		ReturnDBError(w, err)
		return
	}
	body := string(content)
	lookup, err := getlookup(body, Table_Name)
	if err != nil {
		ReturnDBError(w, err)
	}
	datainsert(w, Table_Name, []byte(body), lookup)
}

func ReadCsvFile(path string) ([]byte, string, error) {
	csvFile, err := os.Open(path)

	if err != nil {
		return nil, "", errors.New("the file is not found || wrong root")
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	content, _ := reader.ReadAll()

	if len(content) < 1 {
		return nil, "", fmt.Errorf("something wrong, the file maybe empty or length of the lines are not the same")
	}

	headersArr := make([]string, 0)
	for _, headE := range content[0] {
		headersArr = append(headersArr, headE)
	}

	//Remove the header row
	content = content[1:]

	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headersArr[j] + `":`)
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				buffer.WriteString(y)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(y))
			} else {
				buffer.WriteString((`"` + y + `"`))
			}
			//end of property
			if j < len(d)-1 {
				buffer.WriteString(",")
			}

		}
		//end of object of the array
		buffer.WriteString("}")
		if i < len(content)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(`]`)
	rawMessage := json.RawMessage(buffer.String())
	x, _ := json.MarshalIndent(rawMessage, "", "  ")
	newFileName := filepath.Base(path)
	newFileName = newFileName[0:len(newFileName)-len(filepath.Ext(newFileName))] + ".json"
	r := filepath.Dir(path)
	return x, filepath.Join(r, newFileName), nil
}
