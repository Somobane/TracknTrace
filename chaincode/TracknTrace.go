/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)

// TnT is a high level smart contract that collaborate together business artifact based smart contracts
type TnT struct {
}
// Assembly Line Structure
type AssemblyLine struct{	
	AssemblyId string `json:"assemblyId"`
	DeviceSerialNo string `json:"deviceSerialNo"`
	DeviceType string `json:"deviceType"`
	FilamentBatchId string `json:"filamentBatchId"`
	LedBatchId string `json:"ledBatchId"`
	CircuitBoardBatchId string `json:"circuitBoardBatchId"`
	WireBatchId string `json:"wireBatchId"`
	CasingBatchId string `json:"casingBatchId"`
	AdaptorBatchId string `json:"adaptorBatchId"`
	StickPodBatchId string `json:"stickPodBatchId"`
	ManufacturingPlant string `json:"manufacturingPlant"`
	AssemblyStatus string `json:"assemblyStatus"`
	AssemblyCreationDate string `json:"assemblyCreationDate"`
	AssemblyLastUpdatedOn string `json:"assemblyLastUpdateOn"`
	AssemblyCreatedBy string `json:"assemblyCreatedBy"`
	AssemblyLastUpdatedBy string `json:"assemblyLastUpdatedBy"`
	}

// Package Line Structure
type PackageLine struct{	
	CaseId string `json:"caseId"`
	HolderAssemblyId string `json:"holderAssemblyId"`
	ChargerAssemblyId string `json:"chargerAssemblyId"`
	PackageStatus string `json:"packageStatus"`
	PackagingDate string `json:"packagingDate"`
	PackageCreationDate string `json:"packagingCreationDate"`
	PackageLastUpdatedOn string `json:"packageLastUpdateOn"`
	ShippingToAddress string `json:"shippingToAddress"`
	PackageCreatedBy string `json:"packageCreatedBy"`
	PackageLastUpdatedBy string `json:"packageLastUpdatedBy"`
	}

// Init initializes the smart contracts
func (t *TnT) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists - AssemblyLine
	_, err := stub.GetTable("AssemblyLine")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table for Assembly Line
	err = stub.CreateTable("AssemblyLine", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "assemblyId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "deviceSerialNo", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "filamentBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ledBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "circuitBoardBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "wireBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "casingBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "adaptorBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "stickPodBatchId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "manufacturingPlant", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assemblyStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assemblyCreationDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assemblyLastUpdateOn", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assemblyCreatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "assemblyLastUpdatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating Assembly Line.")
	}
	


	// Check if table already exists: Packaging Line
	_, err = stub.GetTable("PackagingLine")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("PackagingLine", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "caseId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "holderAssemblyId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "chargerAssemblyId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "packageStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "packagingDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "packageLastUpdateOn", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "shippingToAddress", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "packageCreatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "packageLastUpdatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
		
	})
	if err != nil {
		return nil, errors.New("Failed creating Packaging Line.")
	}
		
	
	return nil, nil
}
//API to create an assembly
func (t *TnT) createAssembly(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
if len(args) != 11 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 11. Got: %d.", len(args))
		}
		//var columns []shim.Column
		//_assemblyId:= rand.New(rand.NewSource(99)).Int31

		//Generate the AssemblyId
		rand.Seed(time.Now().Unix())
		
		_assemblyId := strconv.Itoa(rand.Int())
		_deviceSerialNo:= args[0]
		_deviceType:=args[1]
		_FilamentBatchId:=args[2]
		_LedBatchId:=args[3]
		_CircuitBoardBatchId:=args[4]
		_WireBatchId:=args[5]
		_CasingBatchId:=args[6]
		_AdaptorBatchId:=args[7]
		_StickPodBatchId:=args[8]
		_ManufacturingPlant:=args[9]
		_AssemblyStatus:= args[10]

		_time:= time.Now().Local()

		_AssemblyCreationDate := _time.Format(time.RFC3339)
		_AssemblyLastUpdateOn := _time.Format(time.RFC3339)
		_AssemblyCreatedBy := ""
		_AssemblyLastUpdatedBy := ""

		// Insert a row
		ok, err := stub.InsertRow("AssemblyLine", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: _assemblyId}},
				&shim.Column{Value: &shim.Column_String_{String_: _deviceSerialNo}},
				&shim.Column{Value: &shim.Column_String_{String_: _deviceType}},
				&shim.Column{Value: &shim.Column_String_{String_: _FilamentBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _LedBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _CircuitBoardBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _WireBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _CasingBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _AdaptorBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _StickPodBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _ManufacturingPlant}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyCreationDate}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyLastUpdateOn}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyCreatedBy}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyLastUpdatedBy}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

//Update AssemblyLine status
func (t *TnT) updateAssemblyByID(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 12 {
		return nil, errors.New("Incorrect number of arguments. Expecting 12.")
	} 
	
	_assemblyId := args[0]
	

	// Get the row pertaining to this Assembly Id
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: _assemblyId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssemblyLine", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving AssemblyLine with Assemblyid %s. Error %s",_assemblyId, err.Error())
	}
	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}

	// Delete the row pertaining to this assemblyLineId
	err = stub.DeleteRow(
		"Assemblyline",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}
	
		_deviceSerialNo:= args[1]
		_deviceType:=args[2]
		_FilamentBatchId:=args[3]
		_LedBatchId:=args[4]
		_CircuitBoardBatchId:=args[5]
		_WireBatchId:=args[6]
		_CasingBatchId:=args[7]
		_AdaptorBatchId:=args[8]
		_StickPodBatchId:=args[9]
		_ManufacturingPlant:=args[10]
		_AssemblyStatus:= args[11]

		_time:= time.Now().Local()

		_AssemblyCreationDate := row.Columns[12].GetString_()
		_AssemblyLastUpdateOn := _time.Format(time.RFC3339)
		_AssemblyCreatedBy :=  row.Columns[14].GetString_()
		_AssemblyLastUpdatedBy := ""

		// Insert a row
		ok, err := stub.InsertRow("AssemblyLine", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: _assemblyId}},
				&shim.Column{Value: &shim.Column_String_{String_: _deviceSerialNo}},
				&shim.Column{Value: &shim.Column_String_{String_: _deviceType}},
				&shim.Column{Value: &shim.Column_String_{String_: _FilamentBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _LedBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _CircuitBoardBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _WireBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _CasingBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _AdaptorBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _StickPodBatchId}},
				&shim.Column{Value: &shim.Column_String_{String_: _ManufacturingPlant}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyCreationDate}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyLastUpdateOn}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyCreatedBy}},
				&shim.Column{Value: &shim.Column_String_{String_: _AssemblyLastUpdatedBy}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists in Assemblyline.")
		}
		
	return nil, nil

}

//get all AssemblyLines
func (t *TnT) getAllAssembly(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {	
var columns []shim.Column

	rows, err := stub.GetRows("AssemblyLine", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
 
   
		
	res2E:= []*AssemblyLine{}	
	
	for row := range rows {		
		newApp:= new(AssemblyLine)
		newApp.AssemblyId = row.Columns[0].GetString_()
		newApp.DeviceSerialNo = row.Columns[1].GetString_()
		newApp.DeviceType = row.Columns[2].GetString_()
		newApp.FilamentBatchId = row.Columns[3].GetString_()
		newApp.LedBatchId = row.Columns[4].GetString_()
		newApp.CircuitBoardBatchId = row.Columns[5].GetString_()
		newApp.WireBatchId = row.Columns[6].GetString_()
		newApp.CasingBatchId = row.Columns[7].GetString_()
		newApp.AdaptorBatchId = row.Columns[8].GetString_()
		newApp.StickPodBatchId  = row.Columns[9].GetString_()
		newApp.ManufacturingPlant  = row.Columns[10].GetString_()
		newApp.AssemblyStatus  = row.Columns[11].GetString_()
		newApp.AssemblyCreationDate  = row.Columns[12].GetString_()
		newApp.AssemblyLastUpdatedOn  = row.Columns[13].GetString_()
		newApp.AssemblyCreatedBy  = row.Columns[14].GetString_()
		newApp.AssemblyLastUpdatedBy  = row.Columns[15].GetString_()
		if len(newApp.AssemblyId) > 0{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

//get the Assembly against ID
func (t *TnT) getAssemblyByID(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting AssemblyLineID to query")
	}

	_assemblyId := args[0]
	

	// Get the row pertaining to this assemblyID
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: _assemblyId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("AssemblyLine", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the assembly " + _assemblyId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the assembly " + _assemblyId + "\"}"
		return nil, errors.New(jsonResp)
	}

	//return []byte (row), nil
	 mapB, _ := json.Marshal(row)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

//get all Assembly by status
func (t *TnT) getAllAssemblyByStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {	

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting Assmebly Status to query")
	}

	_AssemblyStatus := args[0]
	
	// Get the row pertaining to this status
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: _AssemblyStatus}}
	columns = append(columns, col1)
	
	rows, err := stub.GetRows("AssemblyLine", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
		
	res2E:= []*AssemblyLine{}	
	
	for row := range rows {		
		newApp:= new(AssemblyLine)
		newApp.AssemblyId = row.Columns[0].GetString_()
		newApp.DeviceSerialNo = row.Columns[1].GetString_()
		newApp.DeviceType = row.Columns[2].GetString_()
		newApp.FilamentBatchId = row.Columns[3].GetString_()
		newApp.LedBatchId = row.Columns[4].GetString_()
		newApp.CircuitBoardBatchId = row.Columns[5].GetString_()
		newApp.WireBatchId = row.Columns[6].GetString_()
		newApp.CasingBatchId = row.Columns[7].GetString_()
		newApp.AdaptorBatchId = row.Columns[8].GetString_()
		newApp.StickPodBatchId  = row.Columns[9].GetString_()
		newApp.ManufacturingPlant  = row.Columns[10].GetString_()
		newApp.AssemblyStatus  = row.Columns[11].GetString_()
		newApp.AssemblyCreationDate  = row.Columns[12].GetString_()
		newApp.AssemblyLastUpdatedOn  = row.Columns[13].GetString_()
		newApp.AssemblyCreatedBy  = row.Columns[14].GetString_()
		newApp.AssemblyLastUpdatedBy  = row.Columns[15].GetString_()
		if len(newApp.AssemblyId) > 0{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


// Invoke callback representing the invocation of a chaincode
func (t *TnT) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "createAssembly" {
		fmt.Printf("Function is createAssembly")
		return t.createAssembly(stub, args)
	} else if function == " updateAssemblyByID" {
		fmt.Printf("Function is updateAssemblyLineStatus")
		return t.updateAssemblyByID(stub, args)
	} 

	return nil, errors.New("Received unknown function invocation")
}


// query queries the chaincode
func (t *TnT) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")

	if function == "getAllAssemblyByStatus" { 
		t := TnT{}
		return t.getAllAssemblyByStatus(stub, args)
	} else 
	if function == "getAssemblyByID" { 
		t := TnT{}
		return t.getAssemblyByID(stub, args)
	}else if function == "getAllAssembly" { 
		t := TnT{}
		return t.getAllAssembly(stub, args)
	}
	
	return nil, errors.New("Received unknown function query")
}

	func main() {
	err := shim.Start(new(TnT))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
