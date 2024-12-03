package node

import (
	"pow/engine"

	"encoding/json"
	"strconv"
	"testing"
	"time"
)

func TestRunNode2(t *testing.T) {
	numNodes := 3
	numBlocks := 10
	difficulty := 4

	ch_slice := make([]chan Status, numNodes+1)
	// run node
	for i := 0; i <= numNodes; i++ {
		ch_slice[i] = make(chan Status, 10)
		if i == 0 {
			continue
		}

		node := Node2{ID: i, Channel_Node: ch_slice[i], Channel_Main: ch_slice[0],
			config: config{numBlocks: numBlocks, difficulty: difficulty}}
		go node.RunNode2()
	}

	// run task
	blocks := make([]engine.Block, numBlocks)
	findNonce := make([]bool, numBlocks)
	timestamp := time.Now().Unix()
	block := engine.Block{
		PrevHash:  "0000000000000000",
		Timestamp: timestamp,
		Data:      "Genesis Block",
		Nonce:     0,
		Hash:      "",
	}
	blocks[0] = block
	for i := 1; i <= numNodes; i++ {
		ch_slice[i] <- Status{Block: block, currentBlockNum: 0}
	}

	timeout := 100 * time.Second
	for {
		select {
		case job := <-ch_slice[0]:
			if findNonce[job.currentBlockNum] == true {
				t.Logf("Already find nonce, block%d\n", job.currentBlockNum)
				continue
			}

			findNonce[job.currentBlockNum] = true
			blocks[job.currentBlockNum] = job.Block
			if job.currentBlockNum+1 == numBlocks {
				prettyJSON, err := json.MarshalIndent(blocks, "", "  ")
				if err != nil {
					t.Log("Error marshalling to JSON: ", err)
					return
				}

				t.Logf("Finished find nonce!\n%v\n", string(prettyJSON))
				return
			}
			t.Logf("Main: let's find next block%d\n", job.currentBlockNum+1)

			timestamp := time.Now().Unix()
			for i, ch := range ch_slice {
				if i == 0 {
					continue
				}

				block := engine.Block{PrevHash: job.Block.Hash, Timestamp: timestamp,
					Data:  "Block #" + strconv.Itoa(job.currentBlockNum+1),
					Nonce: 0, Hash: ""}
				ch <- Status{Block: block, currentBlockNum: job.currentBlockNum + 1}
			}
		case <-time.After(timeout):
			t.Logf("All nodes ran successfully for %v seconds", timeout.Seconds())
			return
		}
	}
}
