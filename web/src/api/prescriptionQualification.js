import service from "@/utils/request";

export const getListPQ = (params) => {
  return service({
    url: "/prescription/getList",
    method: "get",
    params,
  });
};
export const createPQ = (data) => {
  return service({
    url: "/prescription/create",
    method: "post",
    data,
  });
};
export const updatePQ = (data) => {
  return service({
    url: "/prescription/update",
    method: "post",
    data,
  });
};
export const deletePQ = (params) => {
  return service({
    url: "/prescription/delete",
    method: "delete",
    params,
  });
};
export const addPersonPQ = (data) => {
  return service({
    url: "/prescription/addPersonPQItem",
    method: "post",
    data,
  });
};
export const getPersonPQList = (params) => {
  return service({
    url: "/prescription/getPersonPQList",
    method: "get",
    params,
  });
};
export const deletePersonPQ = (params) => {
  return service({
    url: "/prescription/deletePersonPQ",
    method: "delete",
    params,
  });
};
