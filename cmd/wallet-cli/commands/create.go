package commands

import (
	"bufio"
	"context"
	"fmt"
	"github.com/coschain/cobra"
	"github.com/coschain/contentos-go/cmd/wallet-cli/commands/utils"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/contentos-go/rpc/pb"
	"os"
	"strings"
)

var createAccountFee string

var CreateCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "create a new account",
		Example: "create [creator] [name]",
		Args:    cobra.ExactArgs(2),
		Run:     create,
	}

	cmd.Flags().StringVarP(&createAccountFee, "fee", "f", "", `create alice bob --fee 1.000000`)
	utils.ProcessEstimate(cmd)
	return cmd
}

var CreateFromMnemonic = func() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create_from_mnemonic",
		Short: "create from mnemonic",
		Example: "create_from_mnemonic [creator] [name]",
		Args: cobra.ExactArgs(2),
		Run: createFromMnemonic,
	}
	cmd.Flags().StringVarP(&createAccountFee, "fee", "f", "", `create_from_mnemonic alice  bob --fee 1.000000`)
	utils.ProcessEstimate(cmd)
	return cmd
}

func create(cmd *cobra.Command, args []string) {
	defer func() {
		utils.EstimateStamina = false
		createAccountFee = ""
	}()
	c := cmd.Context["rpcclient"]
	client := c.(grpcpb.ApiServiceClient)
	w := cmd.Context["wallet"]
	mywallet := w.(wallet.Wallet)
	r := cmd.Context["preader"]
	preader := r.(utils.PasswordReader)
	creator := args[0]
	creatorAccount, ok := mywallet.GetUnlockedAccount(creator)
	if !ok {
		fmt.Println(fmt.Sprintf("creator: %s should be loaded or created first", creator))
		return
	}
	pubKeyStr, privKeyStr, err := mywallet.GenerateNewKey()
	pubkey, _ := prototype.PublicKeyFromWIF(pubKeyStr)
	name := args[1]
	passphrase, err := utils.GetPassphrase(preader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fee, err := utils.ParseAccountCreateFee(client, createAccountFee)
	if err != nil {
		fmt.Println(err)
		return
	}

	acop := &prototype.AccountCreateOperation{
		Fee:            fee,
		Creator:        &prototype.AccountName{Value: creator},
		NewAccountName: &prototype.AccountName{Value: name},
		PubKey:          pubkey,
	}
	signTx, err := utils.GenerateSignedTxAndValidate(cmd, []interface{}{acop}, creatorAccount)
	if err != nil {
		fmt.Println(err)
		return
	}

	if utils.EstimateStamina {
		req := &grpcpb.EsimateRequest{Transaction:signTx}
		res,err := client.EstimateStamina(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res.Invoice)
		}
	} else {
		req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
		resp, err := client.BroadcastTrx(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		} else {
			if resp.Invoice.Status == prototype.StatusSuccess {
				err = mywallet.Create(name, passphrase, pubKeyStr, privKeyStr)
				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println(fmt.Sprintf("Result: %v", resp))
		}
	}
}

func createFromMnemonic(cmd *cobra.Command, args []string) {
	defer func() {
		utils.EstimateStamina = false
		createAccountFee = ""
	}()
	c := cmd.Context["rpcclient"]
	client := c.(grpcpb.ApiServiceClient)
	w := cmd.Context["wallet"]
	mywallet := w.(wallet.HDWallet)
	r := cmd.Context["preader"]
	preader := r.(utils.PasswordReader)
	creator := args[0]
	creatorAccount, ok := mywallet.GetUnlockedAccount(creator)
	if !ok {
		fmt.Println(fmt.Sprintf("creator: %s should be loaded or created first", creator))
		return
	}
	name := args[1]

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter mnemonic > ")
	mnemonic, _ := reader.ReadString('\n')
	// remove \n
	mnemonic = strings.Replace(mnemonic, "\n", "", -1)

	passphrase, err := utils.GetPassphrase(preader)
	if err != nil {
		fmt.Println(err)
		return
	}

	pubKeyStr, privKeyStr, err := mywallet.GenerateFromMnemonic(mnemonic)
	if err != nil {
		fmt.Println(err)
		return
	}

	pubkey, _ := prototype.PublicKeyFromWIF(pubKeyStr)

	fee, err := utils.ParseAccountCreateFee(client, createAccountFee)
	if err != nil {
		fmt.Println(err)
		return
	}

	acop := &prototype.AccountCreateOperation{
		Fee:            fee,
		Creator:        &prototype.AccountName{Value: creator},
		NewAccountName: &prototype.AccountName{Value: name},
		PubKey:          pubkey,
	}
	signTx, err := utils.GenerateSignedTxAndValidate(cmd, []interface{}{acop}, creatorAccount)
	if err != nil {
		fmt.Println(err)
		return
	}

	if utils.EstimateStamina {
		req := &grpcpb.EsimateRequest{Transaction:signTx}
		res,err := client.EstimateStamina(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res.Invoice)
		}
	} else {
		req := &grpcpb.BroadcastTrxRequest{Transaction: signTx}
		resp, err := client.BroadcastTrx(context.Background(), req)
		if err != nil {
			fmt.Println(err)
		} else {
			if resp.Invoice.Status == prototype.StatusSuccess {
				err = mywallet.Create(name, passphrase, pubKeyStr, privKeyStr)
				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println(fmt.Sprintf("Result: %v", resp))
		}
	}
}
