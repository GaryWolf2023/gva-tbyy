import service from "@/utils/request";

export const createPayloadTemp = (data) => {
  return service({
    url: "/payload/createTemp",
    method: "POST",
    data,
  });
};
export const updatePayloadTemp = (data) => {
  return service({
    url: "/payload/updateTemp",
    method: "put",
    data,
  });
};
export const deletePayloadTemp = (params) => {
  return service({
    url: "/payload/deleteTemp",
    method: "delete",
    params,
  });
};
export const getPayloadTempList = (params) => {
  return service({
    url: "/payload/getTempList",
    method: "get",
    params,
  });
};
export const getPayloadTemp = (params) => {
  return service({
    url: "/payload/getTemp",
    method: "get",
    params,
  });
};
