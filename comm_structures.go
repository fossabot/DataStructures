/*
* @Author: Ximidar
* @Date:   2018-09-02 01:36:21
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-09-02 01:41:47
*/

package mango_structures

import(
	"encoding/json"
)

type Init_Comm struct {
	Port string `json:"port"`
	Baud string `json:"baud"`
}

type Write_Comm struct {
	Message string `json:"message"`
}

