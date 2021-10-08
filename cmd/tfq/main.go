package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	dataSource := os.Args[1]
	argumentName := os.Args[2]
	argumentValue := os.Args[3]

	file := hclwrite.NewEmptyFile()
	rootBody := file.Body()

	dataBlock := rootBody.AppendNewBlock("data", nil)
	dataBlock.SetLabels([]string{dataSource, "this"})
	dataBlock.Body().SetAttributeValue(argumentName, cty.StringVal(argumentValue))

	outputBlock := rootBody.AppendNewBlock("output", nil)
	outputBlock.SetLabels([]string{"this"})
	outputBlock.Body().SetAttributeRaw("value", hclwrite.Tokens{
		&hclwrite.Token{
			Type:  hclsyntax.TokenStringLit,
			Bytes: []byte(fmt.Sprintf("data.%s.this", dataSource)),
		},
	})

	outBytes := hclwrite.Format(file.Bytes())

	tmpDir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	tmpfile, err := ioutil.TempFile(tmpDir, "main-*.tf")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(tmpfile.Name(), outBytes, 0644)

	tf, err := tfexec.NewTerraform(tmpDir, "/usr/local/bin/terraform")
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}
	err = tf.Init(context.Background())
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}
	err = tf.Apply(context.Background())
	if err != nil {
		log.Fatalf("error running Apply: %s", err)
	}
	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	bytes, err := json.Marshal(state.Values.Outputs["this"].Value)
	fmt.Print(string(bytes))
}
