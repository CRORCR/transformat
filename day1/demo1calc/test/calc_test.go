package test

import (
	"fmt"
	"testing"
	"transformat/day1/demo1calc"
)

func TestAdd(t *testing.T) {
	sum := demo1calc.Add(3, 2)
	if sum != 5 {
		t.Fatal("add is not right, sum:%v expected:5", sum)
	}
	t.Logf("add is ok")
}

func TestSetGenesis(t *testing.T) {
	tests := []struct {
		nodeid  string
		account string
		role    string
	}{

		{"dbf8dcc4c82eb2ea2e1350b0ea94c7e29f5be609736b91f0faf334851d18f8de1a518def870c774649db443fbce5f72246e1c6bc4a901ef33429fdc3244a93b3",
			"0x6a3217d128a76e4777403e092bde8362d4117773", "vali"},
		{"b624a3fb585a48b4c96e4e6327752b1ba82a90a948f258be380ba17ead7c01f6d4ad43d665bb11c50475c058d3aad1ba9a35c0e0c4aa118503bf3ce79609bef6",
			"0x0ead6cdb8d214389909a535d4ccc21a393dddba9", "vali"},
		{"80606b6c1eecb8ce91ca8a49a5a183aa42f335eb0d8628824e715571c1f9d1d757911b80ebc3afab06647da228f36ecf1c39cb561ef7684467c882212ce55cdb",
			"0x8c3d1a9504a36d49003f1652fadb9f06c32a4408", "vali"},
	}

	for _, test := range tests {
		switch test.role {
		case "miner":
			fmt.Print(test.nodeid, test.account, test.role)
		case "vali":
			fmt.Print(test.nodeid, test.account, test.role)
		}
	}
}
