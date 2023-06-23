package cmd

import (
	"os"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
	"github.com/n-creativesystem/short-url/pkg/utils/credentials/crypto/filekms"
	"github.com/spf13/cobra"
)

func makeCryptKeyCommand() *cobra.Command {
	var (
		file string
	)
	cmd := cobra.Command{
		Use:           "crypto-key",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			executeCryptoKey(file)
		},
	}
	pflags := cmd.Flags()
	pflags.StringVarP(&file, "file", "f", "", "keyset binary file")
	return &cmd
}

func executeCryptoKey(file string) {
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		panic(err)
	}
	f := os.Stdout
	if file != "" {
		f, err = os.Create(file)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}
	w := keyset.NewBinaryWriter(f)
	if err := kh.Write(w, filekms.NewMasterKey()); err != nil {
		panic(err)
	}
}
