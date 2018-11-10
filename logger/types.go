// Copyright 2015 The go-ethereum Authors
// This file is part of Ngin.
//
// Ngin is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Ngin is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Ngin. If not, see <http://www.gnu.org/licenses/>.

package logger

import (
	"math/big"
	"time"
)

type utctime8601 struct{}

func (utctime8601) MarshalJSON() ([]byte, error) {
	timestr := time.Now().UTC().Format(time.RFC3339Nano)
	// Bounds check
	if len(timestr) > 26 {
		timestr = timestr[:26]
	}
	return []byte(`"` + timestr + `Z"`), nil
}

type JsonLog interface {
	EventName() string
}

type LogEvent struct {
	// Guid string      `json:"guid"`
	Ts utctime8601 `json:"ts"`
	// Level string      `json:"level"`
}

type LogStarting struct {
	ClientString    string `json:"client_impl"`
	ProtocolVersion int    `json:"protocol_ver"`
	LogEvent
}

func (l *LogStarting) EventName() string {
	return "starting"
}

type P2PConnected struct {
	RemoteId            string `json:"remote_id"`
	RemoteAddress       string `json:"remote_addr"`
	RemoteVersionString string `json:"remote_version_string"`
	NumConnections      int    `json:"num_connections"`
	LogEvent
}

func (l *P2PConnected) EventName() string {
	return "p2p.connected"
}

type P2PDisconnected struct {
	NumConnections int    `json:"num_connections"`
	RemoteId       string `json:"remote_id"`
	LogEvent
}

func (l *P2PDisconnected) EventName() string {
	return "p2p.disconnected"
}

type MinerNewBlock struct {
	BlockHash     string   `json:"block_hash"`
	BlockNumber   *big.Int `json:"block_number"`
	ChainHeadHash string   `json:"chain_head_hash"`
	BlockPrevHash string   `json:"block_prev_hash"`
	LogEvent
}

func (l *MinerNewBlock) EventName() string {
	return "ngin.miner.new_block"
}

type ChainReceivedNewBlock struct {
	BlockHash     string   `json:"block_hash"`
	BlockNumber   *big.Int `json:"block_number"`
	ChainHeadHash string   `json:"chain_head_hash"`
	BlockPrevHash string   `json:"block_prev_hash"`
	RemoteId      string   `json:"remote_id"`
	LogEvent
}

func (l *ChainReceivedNewBlock) EventName() string {
	return "ngin.chain.received.new_block"
}

type ChainNewHead struct {
	BlockHash     string   `json:"block_hash"`
	BlockNumber   *big.Int `json:"block_number"`
	ChainHeadHash string   `json:"chain_head_hash"`
	BlockPrevHash string   `json:"block_prev_hash"`
	LogEvent
}

func (l *ChainNewHead) EventName() string {
	return "ngin.chain.new_head"
}

type TxReceived struct {
	TxHash   string `json:"tx_hash"`
	RemoteId string `json:"remote_id"`
	LogEvent
}

func (l *TxReceived) EventName() string {
	return "ngin.tx.received"
}
