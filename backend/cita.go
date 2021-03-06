package backend

import (
	"context"
	"errors"
	"math/big"

	"github.com/cryptape/go-web3/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	getBlockNumberMethod        = "cita_blockNumber"
	getBlockMetadataMethod      = "cita_getMetaData"
	getTransactionReceiptMethod = "eth_getTransactionReceipt"
)

func (b *backend) GetBlockNumber() (*big.Int, error) {
	resp, err := b.provider.SendRequest(getBlockNumberMethod)
	if err != nil {
		return nil, err
	}

	hexNumber, err := resp.GetString()
	if err != nil {
		return nil, err
	}

	return hexutil.DecodeBig(hexNumber)
}

func (b *backend) GetBlockMetadata(blockNumber *big.Int) (*types.BlockMetadata, error) {
	hexBlockNumber := hexutil.EncodeBig(blockNumber)
	resp, err := b.provider.SendRequest(getBlockMetadataMethod, hexBlockNumber)
	if err != nil {
		return nil, err
	}

	var meta types.BlockMetadata
	if err := resp.GetObject(&meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

func (b *backend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	resp, err := b.provider.SendRequest(getTransactionReceiptMethod, txHash.Hex())
	if err != nil {
		return nil, err
	}

	var r types.Receipt
	if err := resp.GetObject(&r); err != nil {
		return nil, err
	}

	if r.ErrorMessage != "" {
		return nil, errors.New(r.ErrorMessage)
	}

	return &r, nil
}
