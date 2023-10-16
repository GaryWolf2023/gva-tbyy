import service from "@/utils/request";

// 床架模板
export const createPayloadTemp = (data) => {
  return service({
    url: "/payload/createTemp",
    method: "POST",
    data,
  });
};

// 更新模板
export const updatePayloadTemp = (data) => {
  return service({
    url: "/payload/updateTemp",
    method: "put",
    data,
  });
};

// 更新模板文件
export const updatePayloadOfFile = (data) => {
  return service({
    url: "/payload/updatePayloadOfFile",
    method: "POST",
    data,
  });
};

// 删除模板
export const deletePayloadTemp = (params) => {
  return service({
    url: "/payload/deleteTemp",
    method: "delete",
    params,
  });
};

// 获取模板列表
export const getPayloadTempList = (data) => {
  return service({
    url: "/payload/getTempList",
    method: "post",
    data,
  });
};

// 获取单个模板信息
export const getPayloadTemp = (id) => {
  return service({
    url: "/payload/getTemp",
    method: "GET",
    params: { id },
  });
};

// 文件管理中获取模板文件
export const tempManageGetFile = (params) => {
  return service({
    url: "/payload/getFile",
    method: "GET",
    params,
  });
};

// 页面获取模板文件
/**
 * @param {} tempName
 * @param {} tempType
 * @param {} ID
 */
export const getFileOfTemp = (params) => {
  return service({
    url: "/payload/getFileOfTemp",
    method: "GET",
    params,
  });
};
