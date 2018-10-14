/*
* @Author: Ximidar
* @Date:   2018-10-10 06:36:00
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-14 10:56:40
*/
package file_structures

import(
	"fmt"
	"errors"
)

const(
	NAME = "FILE_MANAGER."

	// File Structure controls
	SELECT_FILE = NAME + "SELECT_FILE"
	GET_FILE_STRUCTURE = NAME + "GET_FILE_STRUCTURE"
	ADD_FILE = NAME + "ADD_FILE"
	MOVE_FILE = NAME + "MOVE_FILE"
	DELETE_FILE = NAME + "DELETE_FILE"

	// Print Controls
	IS_PRINTING = NAME + "IS_PRINTING"
	IS_PAUSED = NAME + "IS_PAUSED"
	TOGGLE_PAUSE = NAME + "TOGGLE_PAUSE"
	START_PRINT = NAME + "START_PRINT"
	CANCEL_PRINT = NAME + "CANCEL_PRINT"


	// Publishers
	UPDATE_FS = NAME + "UPDATE_FS"
	FILE_PROGRESS = NAME + "FILE_PROGRESS"
	
)

type File_Action struct{
	Action string `json:"action"`
	Path string `json:"path"`
}

func New_File_Action(action string, path string) (*File_Action, error){
	fa := new(File_Action)
	if err := fa.Set_Action(action); err != nil{
		return nil, err
	}
	fa.Path = path
	return fa, nil
}

func (fa *File_Action) Set_Action(action string) (error){
	ok := false
	switch(action){
	case SELECT_FILE, GET_FILE_STRUCTURE, ADD_FILE, MOVE_FILE, DELETE_FILE:
		ok = true
	default:
		ok = false
	}

	if !ok{
		return errors.New(fmt.Sprintf("File Action not available: %v", action))
	}
	fa.Action = action
	return nil
}