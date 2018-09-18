/*
* @Author: Ximidar
* @Date:   2018-09-02 01:37:29
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-09-15 17:28:09
*/
package mango_structures

import(
	_"encoding/json"
)

type Reply_String struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}

type Reply_JSON struct{
	Success bool `json:"success"`
	Message []byte `json:"message"`
}
