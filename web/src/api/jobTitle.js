import service from "@/utils/request";

export const getList = (params) => {
  return service({
    url: "/jobTitle/getList",
    method: "get",
    params,
  });
};
export const createJobTitle = (data) => {
  return service({
    url: "/jobTitle/create",
    method: "post",
    data,
  });
};
export const updateJobTitle = (data) => {
  return service({
    url: "/jobTitle/update",
    method: "post",
    data,
  });
};
export const deleteJobTitle = (params) => {
  return service({
    url: "/jobTitle/delete",
    method: "delete",
    params,
  });
};
