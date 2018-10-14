/*
* @Author: Ximidar
* @Date:   2018-10-14 10:56:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-14 11:10:29
*/

package file_structure_test

import(
	"fmt"
	"testing"
	FS "github.com/ximidar/Flotilla/data_structures/file_structures"
)

func TestSetup(t *testing.T){
	fmt.Println("Testing Setup!")
	test_path := "somepath"
	fa, err := FS.New_File_Action(FS.SELECT_FILE, test_path)

	if err != nil{
		t.Fatal(err)
	}

	if fa.Action != FS.SELECT_FILE{
		t.Fatal("Action was not set correctly")
	}

	if fa.Path != test_path{
		t.Fatal("Path was not set correctly")
	}
}
