package main

import (
	"encoding/json"
	"fmt"
	"github.com/ipandtcp/godingtalk"
	"os"
	"time"
)

func main() {

	c := godingtalk.NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	err := c.RefreshAccessToken()
	if err != nil {
		panic(err)
	}

	var compntValues []godingtalk.ProcInstCompntValues
	compntValues = append(compntValues, godingtalk.ProcInstCompntValues{Name: "单行输入框", Value: "单行输入框输入的内容"})

	// 明细控件是数组套数组!!
	detailCompntValues := [][]godingtalk.ProcInstCompntValues{
		[]godingtalk.ProcInstCompntValues{
			godingtalk.ProcInstCompntValues{Name: "明细内单行输入框", Value: "明细内单行输入框的内容"},
		},
	}
	detailValues, _ := json.Marshal(detailCompntValues)
	compntValues = append(compntValues, godingtalk.ProcInstCompntValues{Name: "明细1", Value: string(detailValues)})

	procInstData := godingtalk.TopAPICreateProcInst{
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
		panic(err)
	} else {
		fmt.Printf("%+v\n", procInstID)
	}

	procInst, err := c.TopAPIGetProcInst(procInstID)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v\n", procInst)
	}

	listResp, err := c.TopAPIListProcInst("PROC-FF6YHQ9WQ2-RWDT8XCUTV0U5IAT7JBM1-8MD0TNEJ-6", time.Now().AddDate(0, 0, -10), time.Now(), 10, 0, nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v\n", listResp)
	}
}
