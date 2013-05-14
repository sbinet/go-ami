package ami_test

import (
	"encoding/xml"
	//"fmt"
	"testing"

	"github.com/sbinet/go-ami/ami"
)

const testdata = `<?xml version="1.0" encoding="UTF-8" ?>
<AMIMessage>
	<command entity="router_project">AMIListProject</command>
	<time>2012-09-28  at 11:41:59 AM CEST</time>
	<commandArgs><args argName="amiAdvanced" argValue="ON"/><args argName="amiLang" argValue="english"/><args argName="entity" argValue="router_project"/><args argName="processingStep" argValue="self"/><args argName="project" argValue="self"/><args argName="site" argValue="self"/><args argName="type" argValue="motherDatabase"/></commandArgs>
	<connection  amiLogin="binet" type="router" >successful</connection>
	<connection type="database" project="self" processingStep="self" >successful</connection>
	
<Result entity="router_project">
  <rowset type="router_project">
    <row num="1">
    <field name="identifier" table="router_project" type="INTEGER(3)">1</field>
    <field name="name" table="router_project" type="VARCHAR(50)">self</field>
    <field name="description" table="router_project" type="VARCHAR(65535)">AMI router database</field>
    <field name="projectType" table="router_project" type="VARCHAR(50)"></field>
    </row>
    <row num="56">
    <field name="identifier" table="router_project" type="INTEGER(3)">173</field>
    <field name="name" table="router_project" type="VARCHAR(50)">mc12_001</field>
    <field name="description" table="router_project" type="VARCHAR(65535)">ATLAS_AMI_MC12_01 data</field>
    <field name="projectType" table="router_project" type="VARCHAR(50)">Atlas_Production</field>
    </row>
  </rowset>

</Result>
	<commandStatus>successful</commandStatus><executionTime>0.016</executionTime></AMIMessage>
`

func TestAmiMessage(t *testing.T) {
	msg := ami.Message{}
	err := xml.Unmarshal([]byte(testdata), &msg)
	if err != nil {
		t.Errorf("xml.Unmarshal failed: %v\n", err)
	}
	row := msg.Result.Rows[0]

	if len(msg.Result.Rows) != 2 {
		t.Errorf("expected 2 rows (got: %v)\n", len(msg.Result.Rows))
	}

	row = msg.Result.Rows[0]
	if row.Id() != 1 {
		t.Errorf("expected Num()==1 (got: %v)\n", row.Id())
	}

	if len(row.Fields) != 4 {
		t.Errorf("expected len(row.Fields)==4 (got: %v)\n", len(row.Fields))
	}
	for i, v := range []struct {
		name  string
		data  string
		value interface{}
	}{
		{"identifier", "1", 1},
		{"name", "self", "self"},
		{"description", "AMI router database", "AMI router database"},
		{"projectType", "", ""},
	} {
		if v.data != row.Fields[i].Data {
			t.Errorf("expected [%v] for field [%v] (got: %v)\n",
				v.data, v.name, row.Fields[i].Data)
		}
		if v.value != row.Fields[i].Value() {
			t.Errorf("expected [%v](type=%T) for field [%v] (got: %v (%T,%s))\n",
				v.value, v.value, v.name,
				row.Fields[i].Value(),
				row.Fields[i].Value(),
				row.Fields[i].TypeName)
		}

		if v.value != row.Value()[v.name] {
			t.Errorf("expected [%v](type=%T) for field [%v] (got: %v (%T,%s))\n",
				v.value, v.value, v.name,
				row.Value()[v.name],
				row.Value()[v.name],
				row.Fields[i].TypeName)
		}
	}

	row = msg.Result.Rows[1]
	if row.Id() != 56 {
		t.Errorf("expected Num()==56 (got: %v)\n", row.Id())
	}

	for i, v := range []struct {
		name  string
		data  string
		value interface{}
	}{
		{"identifier", "173", 173},
		{"name", "mc12_001", "mc12_001"},
		{"description", "ATLAS_AMI_MC12_01 data", "ATLAS_AMI_MC12_01 data"},
		{"projectType", "Atlas_Production", "Atlas_Production"},
	} {
		if v.data != row.Fields[i].Data {
			t.Errorf("expected [%v] for field [%v] (got: %v)\n",
				v.data, v.name, row.Fields[i].Data)
		}
		if v.value != row.Fields[i].Value() {
			t.Errorf("expected [%v](type=%T) for field [%v] (got: %v (%T,%s))\n",
				v.value, v.value, v.name,
				row.Fields[i].Value(),
				row.Fields[i].Value(),
				row.Fields[i].TypeName)
		}
		if v.value != row.Value()[v.name] {
			t.Errorf("expected [%v](type=%T) for field [%v] (got: %v (%T,%s))\n",
				v.value, v.value, v.name,
				row.Value()[v.name],
				row.Value()[v.name],
				row.Fields[i].TypeName)
		}
	}

}
