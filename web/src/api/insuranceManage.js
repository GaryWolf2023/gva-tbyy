import service from "@/utils/request";

export const getInsuranceList = (params) => {
  return service({
    url: "/insurance/getInsuranceList",
    method: "get",
    params,
  });
};
export const getInsuranceInfo = (params) => {
  return service({
    url: "/insurance/getInsuranceInfo",
    method: "get",
    params,
  });
};
export const createInsurance = (data) => {
  return service({
    url: "/insurance/createInsurance",
    method: "post",
    data,
  });
};
export const updateInsurance = (data) => {
  return service({
    url: "/insurance/updateInsurance",
    method: "post",
    data,
  });
};
export const deleteInsurance = (params) => {
  return service({
    url: "/insurance/deleteInsurance",
    method: "delete",
    params,
  });
};
export const getPersonInsuranceList = (params) => {
  return service({
    url: "/insurance/getPersonInsuranceList",
    method: "get",
    params,
  });
};
export const addPersonInsurance = (data) => {
  return service({
    url: "/insurance/addPersonInsurance",
    method: "post",
    data,
  });
};
export const deletePersonInsurance = (data) => {
  return service({
    url: "/insurance/deletePersonInsurance",
    method: "post",
    data,
  });
};
