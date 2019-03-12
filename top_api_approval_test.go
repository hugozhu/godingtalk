package godingtalk

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTopAPIProcInst(t *testing.T) {
	var compntValues []ProcInstCompntValues
	compntValues = append(compntValues, ProcInstCompntValues{Name: "单行输入框", Value: "单行输入框输入的内容"})

	detailCompntValues := [][]ProcInstCompntValues{[]ProcInstCompntValues{ProcInstCompntValues{Name: "明细内单行输入框", Value: "明细内单行输入框的内容"}}}
	detailValues, _ := json.Marshal(detailCompntValues)
	compntValues = append(compntValues, ProcInstCompntValues{Name: "明细1", Value: string(detailValues)})

	procInstData := TopAPICreateProcInst{
		Approvers:        []string{"085354234826136236"},
		CCList:           []string{"085354234826136236"},
		CCPosition:       "START",
		DeptID:           4207088,
		OriginatorUID:    "085354234826136236",
		ProcessCode:      "PROC-FF6YHQ9WQ2-RWDT8XCUTV0U5IAT7JBM1-8MD0TNEJ-6",
		FormCompntValues: compntValues,
	}
	procInstID, err := c.TopAPICreateProcInst(procInstData)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v\n", procInstID)
	}

	procInst, err := c.TopAPIGetProcInst(procInstID)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v\n", procInst)
	}

	listResp, err := c.TopAPIListProcInst("PROC-FF6YHQ9WQ2-RWDT8XCUTV0U5IAT7JBM1-8MD0TNEJ-6", time.Now().AddDate(0, 0, -10), time.Now(), 10, 0, nil)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v\n", listResp)
	}
}
