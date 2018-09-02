/*
* @Author: Ximidar
* @Date:   2018-09-02 01:37:29
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-09-02 01:42:03
*/
package mango_structures

import(
	"encoding/json"
)

type Reply_String struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}

type Reply_JSON struct{
	Success bool `json:"success"`
	JSON_Message []byte `json:"message"`
}
