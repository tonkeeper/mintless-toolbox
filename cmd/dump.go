package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/tonkeeper/tongo"
	"github.com/tonkeeper/tongo/boc"
	"github.com/tonkeeper/tongo/tlb"
)

func newDumpCmd(out io.Writer) *cobra.Command {
	options := dumpOptions{}
	cmd := &cobra.Command{
		Use:  "dump",
		Args: cobra.ExactArgs(1),
		Long: "dump extract data from airdrop file and dump it to a terminal in csv format.",
		RunE: func(cmd *cobra.Command, args []string) error {
			options.airdropFile = args[0]
			return dump(cmd.Context(), options, out)
		},
	}
	return cmd
}

type dumpOptions struct {
	airdropFile string
}

func deserialize(content []byte) ([]*boc.Cell, error) {
	if airdropCells, err := boc.DeserializeBoc(content); err == nil {
		return airdropCells, nil
	}
	return boc.DeserializeBocHex(string(content))
}

func dump(ctx context.Context, options dumpOptions, out io.Writer) error {
	content, err := os.ReadFile(options.airdropFile)
	if err != nil {
		return fmt.Errorf("failed to read airdrop file: %w", err)
	}
	airdropCells, err := deserialize(content)
	if err != nil {
		return fmt.Errorf("failed to deserialize airdrop file: %w", err)

	}
	var m tlb.Hashmap[Address, AirdropData]
	if err := tlb.Unmarshal(airdropCells[0], &m); err != nil {
		return fmt.Errorf("failed to unmarshal airdrop data: %w", err)
	}
	for _, item := range m.Items() {
		_, err := fmt.Fprintf(out, "%s,%v,%v,%v\n", item.Key.ToRaw(), uint64(item.Value.Amount), uint64(item.Value.StartFrom), uint64(item.Value.ExpireAt))
		if err != nil {
			return fmt.Errorf("failed to write to csv file: %w", err)
		}
	}
	return nil
}

type AirdropData struct {
	Amount    tlb.Coins
	StartFrom tlb.Uint48
	ExpireAt  tlb.Uint48
}

type Address struct {
	tlb.MsgAddress
}

func (addr Address) Equal(other any) bool {
	otherAddr, ok := other.(Address)
	if !ok {
		return false
	}
	return addr.MsgAddress == otherAddr.MsgAddress
}

func (addr Address) FixedSize() int {
	return 267
}

func (addr *Address) UnmarshalTLB(c *boc.Cell, decoder *tlb.Decoder) error {
	var msgAddr tlb.MsgAddress
	if err := decoder.Unmarshal(c, &msgAddr); err != nil {
		return err
	}

	*addr = Address{MsgAddress: msgAddr}
	return nil
}
func (addr *Address) ToRaw() string {
	account, err := tongo.AccountIDFromTlb(addr.MsgAddress)
	if err != nil {
		panic(err)
	}
	return account.String()
}
