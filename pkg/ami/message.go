package ami

import (
	"encoding/xml"
	//"time"
)

type Message struct {
	XMLName xml.Name `xml:"AMIMessage"`
	Command string `xml:"command"`
	//Time time.Time `xml:"time"`
	CommandArgs []*CmdArgs `xml:"commandArgs"`
	Result Result `xml:"Result"`
	Status string `xml:"commandStatus"`
	ExecTime float64 `xml:"executionTime"`
}

type CmdArgs struct {
	
}

type Result struct {
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
