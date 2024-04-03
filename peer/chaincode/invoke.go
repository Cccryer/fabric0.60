/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package chaincode

import (
	"fmt"

	"github.com/spf13/cobra"
)

var chaincodeInvokeNonce string //tx nonce
/*
tx nonce will be transformed to uint32 , it will hash with domain to get target
*/

func invokeCmd() *cobra.Command {
	chaincodeInvokeCmd.PersistentFlags().StringVarP(&chaincodeInvokeNonce, "nonce", "n", "",
		fmt.Sprintf("Nonce of transaction for safety"))
	return chaincodeInvokeCmd
}

var chaincodeInvokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: fmt.Sprintf("Invoke the specified %s.", chainFuncName),
	Long:  fmt.Sprintf(`Invoke the specified %s.`, chainFuncName),
	//ValidArgs: []string{"1"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return chaincodeInvoke(cmd, args)
	},
}

func chaincodeInvoke(cmd *cobra.Command, args []string) error {
	return chaincodeInvokeOrQuery(cmd, args, true)
}
