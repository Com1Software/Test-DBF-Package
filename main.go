package main

import (
	// "strings"
	"fmt"
	//---------------------- Windows
	"comone/dbf"
	//---------------------- Mac/Android/Linux
	//"../dbf"
)
//------- (c) 2022 Com1 Software Development
//----------------------------------------------------------------
func main() {
	fmt.Println("testdbf (c) 2022 Com1 Software Development")
	testCtl := 1
	dbffile := "csc030.dbf"
	if dbf.CheckForFile(dbffile) {
		dbfdata := dbf.LoadFile(dbffile)
		fmt.Printf("Number of Records in Table = %d\n", dbf.GetRecordCount(dbfdata))
		fmt.Printf("Record Length = %d\n", dbf.GetRecordLength(dbfdata))
		fmt.Printf("First Record Position = %d\n", dbf.GetFirstRecordPosition(dbfdata))
		fmt.Printf("Field Count = %d\n", dbf.GetFieldCount(dbfdata))
		switch {
		//--------------------------- General Display
		case testCtl == 0:
			fmt.Println("-------- General Display")
			fctl := true
			if fctl {
				for i := 0; i < dbf.GetFieldCount(dbfdata); i++ {
					fmt.Printf("-->%s\n", dbf.GetFieldName(dbfdata, i))
					fmt.Printf("-->%d\n", dbf.GetFieldLength(dbfdata, i))
					fmt.Printf("--> %s\n", dbf.GetRecordField(dbfdata, dbf.GetRecord(dbfdata, i), i))
				}
			}
			dctl := false
			if dctl {
				for i := 0; i < dbf.GetRecordCount(dbfdata); i++ {
					fmt.Printf("-->%s\n", dbf.GetRecord(dbfdata, i))
				}
			}
			//----------------- Test for GetRecordField
		case testCtl == 1:
			fmt.Println("-------- Test for GetRecordField")
			fmt.Printf("Display Record 0 Field 0  --> %s\n", dbf.GetRecordField(dbfdata, dbf.GetRecord(dbfdata, 0), 0))
			fmt.Printf("Field Length >%d\n", dbf.GetFieldLength(dbfdata, 0))
			fmt.Printf("Actual Len --> %d\n", len(dbf.GetRecordField(dbfdata, dbf.GetRecord(dbfdata, 0), 0)))
			fmt.Printf("Display Record 0 Field 1  --> %s\n", dbf.GetRecordField(dbfdata, dbf.GetRecord(dbfdata, 0), 1))
			fmt.Printf("Field Length >%d\n", dbf.GetFieldLength(dbfdata, 1))
			fmt.Printf("Actual Len --> %d\n", len(dbf.GetRecordField(dbfdata, dbf.GetRecord(dbfdata, 0), 1)))

		}
	}
}
