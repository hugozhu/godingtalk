package godingtalk

import (
    "fmt"
    "net/url"
)

type Department struct {
    OAPIResponse
    Id int
    Name string
    ParentId int    
    Order int
    DeptPerimits string
    UserPerimits string
    OuterDept bool
    OuterPermitDepts string
    OuterPermitUsers string
    OrgDeptOwner string
    DeptManagerUseridList string
}

type DepartmentList struct {
    OAPIResponse
    Departments []Department `json:"department"`
}

// DepartmentList is 获取部门列表
func (c *DingTalkClient) DepartmentList() (DepartmentList, error) {
    var data DepartmentList
    err := c.httpRPC("department/list", nil, nil, &data)   
    return data, err
}

//DepartmentDetail is 获取部门详情
func (c *DingTalkClient) DepartmentDetail(id int) (Department, error) {
    var data Department
    params := url.Values{}
    params.Add("id", fmt.Sprintf("%d", id))
    err :=c.httpRPC("department/get", params, nil, &data)
    return data, err
}