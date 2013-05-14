package ami

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	XMLName xml.Name `xml:"AMIMessage"`
	Cmd     string   `xml:"command"`
	//Time      time.Time    `xml:"time"`
	CmdArgs   []xml_cmdarg `xml:"commandArgs>args"`
	Result    Result       `xml:"Result"`
	CmdStatus string       `xml:"commandStatus"`
	//ExecTime  float64      `xml:"executionTime"`
}

func (msg *Message) Status() bool {
	return msg.CmdStatus == "successful"
}

type xml_cmdarg struct {
	Name  string `xml:"argName,attr"`
	Value string `xml:"argValue,attr"`
}

type Result struct {
	Rows []*Row `xml:"rowset>row"`
}

type xml_rowfield struct {
	Name     string `xml:"name,attr"`
	Table    string `xml:"table,attr"`
	TypeName string `xml:"type,attr"`
	Data     string `xml:",chardata"`
}

type Row struct {
	Num    string         `xml:"num,attr"`
	Fields []xml_rowfield `xml:"field"`
}

func (r *Row) Id() int {
	i, err := strconv.Atoi(r.Num)
	if err != nil {
		panic(fmt.Sprintf("ami.Row.Num: %v\n", err))
	}
	return i
}

func (r *Row) Get(key string) interface{} {
	for _, f := range r.Fields {
		if key == f.Name {
			return f.Value()
		}
	}
	panic(fmt.Sprintf("ami.Row.Get: no such key [%s]", key))
}

func (r *Row) Value() map[string]interface{} {
	o := make(map[string]interface{}, len(r.Fields))
	for _, f := range r.Fields {
		o[f.Name] = f.Value()
	}
	return o
}

func (r *xml_rowfield) Value() interface{} {
	tn := strings.ToLower(r.TypeName)
	if strings.HasPrefix(tn, "integer") || strings.HasPrefix(tn, "number") {
		val, err := strconv.Atoi(r.Data)
		if err != nil {
			panic(fmt.Sprintf("ami.Row.Value: %v\n", err))
		}
		return val
	}

	if strings.HasPrefix(tn, "varchar") || tn == "text" ||
		strings.HasPrefix(tn, "char") {
		return r.Data
	}

	if strings.HasPrefix(tn, "datetime") {
		layout := "2006-01-02 15:04:05"
		val, err := time.Parse(layout, r.Data)
		if err != nil {
			panic(fmt.Sprintf("ami.Row.Value: %v\n", err))
		}
		return val
	}

	panic("ami.Row.Value: unhandled typename [" + r.TypeName + "] (value=" + r.Data + ")")
}

/*
<?xml version="1.0" encoding="UTF-8" ?>
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
[...]
    <row num="56">
    <field name="identifier" table="router_project" type="INTEGER(3)">173</field>
    <field name="name" table="router_project" type="VARCHAR(50)">mc12_001</field>
    <field name="description" table="router_project" type="VARCHAR(65535)">ATLAS_AMI_MC12_01 data</field>
    <field name="projectType" table="router_project" type="VARCHAR(50)">Atlas_Production</field>
    </row>
  </rowset>

</Result>
	<commandStatus>successful</commandStatus><executionTime>0.016</executionTime></AMIMessage>

*/

// EOF
