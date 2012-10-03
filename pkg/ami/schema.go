package ami

type Table struct {
	Name string
	Fields  map[string]string
	Primary string
	Foreign map[string]*Table
}

var ProdstepTable = Table{
	Name: "ProdstepTable",
	Fields: map[string]string{
		"name":         "productionStepName",
		"tag":          "productionStepTag",
		"write_status": "writeStatus",
		"read_status":  "readStatus",
	},
	Primary: "productionStepName",
	Foreign: nil,
}

var FileTable = Table{
	Name: "FileTable",
	Fields: map[string]string{
		"lfn":    "LFN",
		"guid":   "fileFUID",
		"events": "events",
		"size":   "fileSize",
	},
	Primary: "lfn",
	Foreign: nil,
}

var DatasetTable = Table{
	Name: "DatasetTable",
	Fields: map[string]string{
		"ami_status":     "amiStatus",
		"nfiles":         "nFiles",
		"events":         "totalEvents",
		"type":           "dataType",
		"prodsys_status": "prodsysStatus",
		"geometry":       "geometryVersion",
		"version":        "version",
		"transformation": "TransformationPackage",
		"trigger_config": "triggerConfig",
		"atlas_release":  "AtlasRelease",
		"job_config":     "jobConfig",
		"project":        "projectName",
		"dataset_number": "datasetNumber",
		"modified":       "lastModified",
		//"modified-after": "lastModified>",
		//"modified-before": "lastModified<",
		"physics_short":           "physicsShort",
		"history":                 "productionHistory",
		"prod_step":               "prodStep",
		"requested_by":            "requestedBy",
		"name":                    "logicalDatasetName",
		"responsible":             "physicistResponsible",
		"physics_comment":         "physicsComment",
		"modified_by":             "modifiedBy",
		"trash_annotation":        "trashAnnotation",
		"physics_category":        "physicsCategory",
		"trash_date":              "trashDate",
		"trash_trigger":           "trashTrigger",
		"physics_process":         "physicsProcess",
		"principal_physics_group": "principalPhysicsGroup",
		"created":                 "created",
		"created_by":              "createdBy",
		"creation_comment":        "creationComment",
		"stream":                  "streamName",
		"in_container":            "inContainer",
		"run":                     "runNumber",
		"period":                  "period",
		"beam":                    "beamType",
		"conditions_tag":          "conditionsTag",
	},
	Primary: "logicalDatasetName",
	Foreign: map[string]*Table{
		"Files": &FileTable,
	},
}

var NomenclatureTable = Table{
	Name: "NomenclatureTable",
	Fields: map[string]string{
		"template":     "nomenclatureTemplate",
		"name":         "nomenclatureName",
		"tag":          "nomenclatureTag",
		"write_status": "writeStatus",
		"read_status":  "readStatus",
	},
	Primary: "nomenclatureName",
	Foreign: nil,
}

var ProjectTable = Table{
	Name: "ProjectTable",
	Fields: map[string]string{
		"tag":          "projectTag",
		"description":  "description",
		"is_base_type": "isBaseType",
		"read_status":  "readStatus",
		"write_status": "writeStatus",
		"manager":      "projectManager",
	},
	Primary: "projectTag",
	Foreign: map[string]*Table{
		"nomenclature": &NomenclatureTable,
	},
}

var SubprojectTable = Table{
	Name: "SubprojectTable",
	Fields: map[string]string{
		"tag":          "subProjectTag",
		"description":  "description",
		"is_base_type": "isBaseType",
		"read_status":  "readStatus",
		"write_status": "writeStatus",
		"manager":      "projectManager",
	},
	Primary: "subProjectTag",
	Foreign: map[string]*Table{
		"nomenclature": &NomenclatureTable,
	},
}

var TypeTable = Table{
	Name: "TypeTable",
	Fields: map[string]string{
		"type":         "dataType",
		"description":  "description",
		"write_status": "writeStatus",
		"read_status":  "readStatus",
	},
	Primary: "dataType",
	Foreign: nil,
}

var SubtypeTable = Table{
	Name: "SubtypeTable",
	Fields: map[string]string{
		"type":         "subDataType",
		"description":  "description",
		"write_status": "writeStatus",
		"read_status":  "readStatus",
	},
	Primary: "subDataType",
	Foreign: nil,
}

var Tables = map[string]*Table{
	"data_type":      &TypeTable,
	"subData_type":   &SubtypeTable,
	"projects":       &ProjectTable,
	"subProjects":    &SubprojectTable,
	"dataset":        &DatasetTable,
	"nomenclature":   &NomenclatureTable,
	"productionStep": &ProdstepTable,
}

// EOF
