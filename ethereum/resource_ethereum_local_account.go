package ethereum

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceEthereumLocalAccount() *schema.Resource {
	return &schema.Resource{
		Create: createEthereumLocalAccount,
		Read:   readEthereumLocalAccount,
		Update: updateEthereumLocalAccount,
		Delete: deleteEthereumLocalAccount,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private_key": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ValidateFunc: validatePrivateKey,
			},
			"public_key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createEthereumLocalAccount(d *schema.ResourceData, meta interface{}) (err error) {
	var pk *ecdsa.PrivateKey

	k, ok := d.GetOk("private_key")
	if !ok {
		pk, err = crypto.GenerateKey()
	} else {
		pk, err = crypto.HexToECDSA(k.(string))
	}

	if err != nil {
		return err
	}

	hex := common.Bytes2Hex(crypto.FromECDSA(pk))
	d.SetId(hex)

	return readEthereumLocalAccount(d, meta)
}

func readEthereumLocalAccount(d *schema.ResourceData, meta interface{}) error {
	hex := d.Id()

	pk, err := crypto.ToECDSA(common.Hex2Bytes(hex))
	if err != nil {
		return err
	}

	_ = d.Set("public_key", common.Bytes2Hex(crypto.FromECDSAPub(&pk.PublicKey)))
	_ = d.Set("address", crypto.PubkeyToAddress(pk.PublicKey).Hex())
	_ = d.Set("private_key", common.Bytes2Hex(crypto.FromECDSA(pk)))

	return nil
}

func updateEthereumLocalAccount(d *schema.ResourceData, meta interface{}) error {
	return readEthereumLocalAccount(d, meta)
}

func deleteEthereumLocalAccount(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func validatePrivateKey(keyI interface{}, k string) ([]string, []error) {
	key, ok := keyI.(string)
	if !ok {
		return nil, []error{errors.New("private key should be a string")}
	}

	if _, err := crypto.HexToECDSA(key); err != nil {
		return nil, []error{err}
	}

	return nil, nil
}
