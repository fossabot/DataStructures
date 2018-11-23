/*
* @Author: Ximidar
* @Date:   2018-10-10 06:36:00
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-11-22 20:31:50
 */

package FileStructures

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/go-nats"
	"github.com/ximidar/Flotilla/Flotilla_File_Manager/Files"
)

const (
	// Name is the name of the program to call
	Name = "FILE_MANAGER."

	// File Structure controls

	// SelectFile will Select a file
	SelectFile = Name + "SELECT_FILE"
	// GetFileStructure will get the JSON representation of the current file structure
	GetFileStructure = Name + "GET_FILE_STRUCTURE"
	// AddFile will add a file to the file structure
	AddFile = Name + "ADD_FILE"
	// MoveFile will move a file from one place to another
	MoveFile = Name + "MOVE_FILE"
	// DeleteFile will delete a file
	DeleteFile = Name + "DELETE_FILE"

	// Publishers

	// UpdateFS will be called any time there is an update to the file system
	UpdateFS = Name + "UPDATE_FS"
	// FileProgress will be called every time the file progresses (0 - 100%)
	FileProgress = Name + "FILE_PROGRESS"
)

// FileAction is an object to convey what action you want
// the File system to take
type FileAction struct {
	Action string `json:"action"`
	Path   string `json:"path"`
}

// NewFileAction will construct a FileAction object
func NewFileAction(action string, path string) (*FileAction, error) {
	fa := new(FileAction)
	if err := fa.SetAction(action); err != nil {
		return nil, err
	}
	fa.Path = path
	return fa, nil
}

// NewFileActionFromMSG will construct a File action from a nats msg
func NewFileActionFromMSG(msg *nats.Msg) (*FileAction, error) {
	fa := new(FileAction)

	err := json.Unmarshal(msg.Data, &fa)

	if err != nil {
		return nil, err
	}

	return fa, nil
}

// SetAction will make sure a real action is taking place.
func (fa *FileAction) SetAction(action string) error {
	var ok bool
	switch action {
	case SelectFile, GetFileStructure, AddFile, MoveFile, DeleteFile:
		ok = true
	default:
		ok = false
	}

	if !ok {
		return fmt.Errorf("File Action not available: %v", action)
	}
	fa.Action = action
	return nil
}

// SendAction will use the supplied nats conn to send the File Action
func (fa *FileAction) SendAction(nc *nats.Conn, timeout time.Duration) (reply *nats.Msg, err error) {

	data, err := json.Marshal(fa)

	if err != nil {
		return nil, err
	}

	reply, err = nc.Request(fa.Action, data, timeout)

	if err != nil {
		return nil, err
	}

	return reply, nil

}

//FSProgress is a struct for getting and setting the File Progress
type FSProgress struct {
	File        *Files.File
	CurrentLine int64
	ReadBytes   int64
	Progress    float64
}

// NewFSProgressFromMSG will construct a FSProgress object from a received message
func NewFSProgressFromMSG(msg *nats.Msg) (*FSProgress, error) {

	// make FSP
	fsp := new(FSProgress)

	// unmarshal the data from the msg
	err := json.Unmarshal(msg.Data, fsp)
	if err != nil {
		return nil, err
	}

	return fsp, nil

}
