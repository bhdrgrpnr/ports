package data

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"saasteamtest/backend/domain"
	"saasteamtest/backend/models"
	"strings"
)

type PortsHandle struct {
	db DB
}

type DB struct {
	PortTable map[string]models.Port
}

func NewPortHandler(db DB) *PortsHandle {

	file, err := openFile()
	if err != nil {
		log.Fatalf("could not open the file: %v", err)
	}
	file.Seek(0, 0)
	readWithReadLine(file, db)
	defer file.Close()

	return &PortsHandle{db}
}

func (h *PortsHandle) Create(request models.PortRequest) (*models.Port, error) {
	portCreated := request.Port
	h.db.PortTable[request.Index] = portCreated
	return &portCreated, nil
}

func (h *PortsHandle) Update(request models.PortRequest) (*models.Port, error) {

	if _, ok := h.db.PortTable[request.Index]; !ok {
		return nil, domain.ErrIndexNotFound
	} else {
		h.db.PortTable[request.Index] = request.Port
		return &request.Port, nil
	}
}

func readWithReadLine(file *os.File, db DB) {
	reader := bufio.NewReader(file)
	jsonBlob := ""
	read(reader)
	for {
		lineByte, err := read(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("a real error happened here: %v\n", err)
		}
		line := string(lineByte)
		if strings.Contains(line, "},") {
			line = "}}"
			jsonBlob += line
			key, port, errParse := parseJsonBlob("{" + jsonBlob)
			if errParse != nil || port == nil || key == nil {
				continue
			}
			db.PortTable[*key] = *port
			jsonBlob = ""
			continue
		}
		jsonBlob += line
	}
}

func parseJsonBlob(jsonStr string) (*string, *models.Port, error) {
	portMap := map[string]models.Port{}
	err := json.Unmarshal([]byte(jsonStr), &portMap)
	if err != nil {
		fmt.Println(err)
	}

	for key, val := range portMap {
		return &key, &val, nil
	}

	return nil, nil, errors.New("parse error")
}

func read(r *bufio.Reader) ([]byte, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return ln, err
}

func openFile() (*os.File, error) {
	file, err := os.Open("../../../ports.json")
	if err != nil {
		file, err := os.Open("./ports.json")
		if err != nil {
			return nil, errors.New("could not open file")
		}
		return file, nil
	}
	return file, nil
}
