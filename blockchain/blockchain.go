package blockchain

import "../block"

type Blockchain struct {
	genesisBlock    block.Block
	mostRecentBlock block.Block
}
