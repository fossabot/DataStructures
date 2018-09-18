/*
* @Author: Ximidar
* @Date:   2018-09-02 01:36:21
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-09-18 00:27:29
*/

package mango_structures

import(
	_"encoding/json"
)

type Init_Comm struct {
	Port string `json:"port"`
	Baud int `json:"baud"`
}

type Comm_Status struct {
	Port string `json:"port"`
	Baud string `json:"baud"`
	Connected bool `json:"connected"`
	
}
