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
