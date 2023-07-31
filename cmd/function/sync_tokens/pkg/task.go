package pkg

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"tokens-api/internal/tokens"
	"tokens-api/pkg/behaviour"
	"tokens-api/pkg/handler"
)

type task struct {
	tokenReader behaviour.ReaderDomain
	tokenWriter behaviour.WriterDomain
}

func NewTask(tokenReader behaviour.ReaderDomain, tokenWriter behaviour.WriterDomain) *task {
	return &task{
		tokenReader,
		tokenWriter,
	}
}

func (t *task) Execute(ctx context.Context) error {
	_, err := t.syncTokensData(ctx)
	if err != nil {
		return err
	}
	return nil
}

//syncTokensData reads cids files and get data using external service
func (t *task) syncTokensData(ctx context.Context) ([]tokens.Tokens, error) {
	tokensDataFile := os.Getenv("TOKENS_DATA_FILE")
	if tokensDataFile == "" {
		tokensDataFile = "cmd/function/sync_tokens/pkg/ipfs_cids.csv"
	}
	file, err := os.Open(tokensDataFile)
	if err != nil {
		return nil, err
	}
	parser := csv.NewReader(file)
	//TODO: we can read and process by batch using concurrency
	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		cid := record[0]
		data, err := t.getDataFromPinataIFPSGateway(cid)
		if err != nil {
			return nil, err
		}
		fmt.Println(data)
		token := tokens.Tokens{
			ID:          cid,
			Name:        data.Name,
			Description: data.Description,
			Image:       data.Image,
		}

		//check if already exists. If already exists, we will ignore
		tok, err := t.tokenReader.ByID(ctx, cid)
		if err != nil && !errors.Is(err, handler.ErrNotFound) {
			return nil, err
		}
		if tok != nil {
			continue
		}

		_, err = t.tokenWriter.Create(ctx, token)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (t *task) getDataFromPinataIFPSGateway(cid string) (*IPFSGatewayResponse, error) {
	var result IPFSGatewayResponse
	url := fmt.Sprintf(`https://blockpartyplatform.mypinata.cloud/ipfs/%s`, cid)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != http.StatusOK {
		msgErr := fmt.Sprintf("error executing get %s", url)
		log.Printf("[ERROR] %s", msgErr)
		return nil, errors.New(msgErr)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		msgErr := fmt.Sprintf("error unmarshalling get %s", url)
		log.Printf("[ERROR] %s", msgErr)
		return nil, errors.New(msgErr)
	}
	return &result, nil

}
