import service from "@/utils/request";

export const getStaffList = (data) => {
  return service({
    url: "/staff/getStaffList",
    method: "GET",
    params: data,
  });
};

// 获取单个
export const getStaffInfo = (id) => {
  return service({
    url: "/staff/getStaff",
    method: "get",
    params: { id },
  });
};

// 删除此条信息
export const deleteStaff = (id) => {
  return service({
    url: "/staff/deleteStaffInfo",
    method: "delete",
    params: { id },
  });
};

// 更新此条信息
export const updateStaff = (data) => {
  return service({
    url: "/staff/updateStaff",
    method: "post",
    data,
  });
};
// 更新系统配置
export const updateStaffSystemConfig = (data) => {
  return service({
    url: "/staff/updateSystemConfig",
    method: "post",
    data,
  });
};
//  获取擅长疾病列表
export const getAdeptIllnessList = (params) => {
  return service({
    url: "/adeptIllness/getAdeptIllnessList",
    method: "GET",
    params,
  });
};
// 新增一条擅长疾病
export const addAdeptIllness = (data) => {
  return service({
    url: "/adeptIllness/addAdeptIllness",
    method: "post",
    data,
  });
};
// 删除一条擅长疾病
export const deleteAdeptIllness = (params) => {
  return service({
    url: "/adeptIllness/deleteAdeptIllness",
    method: "delete",
    params,
  });
};
