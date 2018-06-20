// Package model_test contains the tests for the model package
package model_test

import (
	"testing"

	"github.com/joincivil/civil-events-crawler/pkg/model"
)

func TestNameToContractTypes(t *testing.T) {
	_type, ok := model.NameToContractTypes.Get("civiltcr")
	if !ok {
		t.Error("Should have found civiltcr as a contract name")
	}
	if _type != model.CivilTcrContractType {
		t.Error("Type should have matched the type for civiltcr")
	}
	_, ok = model.NameToContractTypes.Get("nomatchhere")
	if ok {
		t.Error("Should have not found an invalid name")
	}
	names := model.NameToContractTypes.Names()
	if len(names) == 0 {
		t.Error("Should have returned a valid names list")
	}
}

func TestContractTypeToSpecs(t *testing.T) {
	specs, ok := model.ContractTypeToSpecs.Get(model.CivilTcrContractType)
	if !ok {
		t.Error("Should have found civiltcr specs")
	}
	if specs.SimpleName() != "civiltcr" {
		t.Error("Should have found civiltcr as simple name in spec")
	}
	if specs.Name() != "CivilTCRContract" {
		t.Error("Should have found CivilTCRContract as the name in spec")
	}
	if specs.AbiStr() == "" {
		t.Error("Should have found the abi string in spec")
	}
	if specs.ImportPath() == "" {
		t.Error("Should have found the import path in spec")
	}
	if specs.TypePackage() == "" {
		t.Error("Should have found the type package in spec")
	}

	_, ok = model.ContractTypeToSpecs.Get(model.InvalidContractType)
	if ok {
		t.Error("Should have not found invalid type specs")
	}
	_types := model.ContractTypeToSpecs.Types()
	if len(_types) == 0 {
		t.Error("Should have returned a valid list of specs")
	}
}
