import service from "@/utils/request";

export const getPartTimePositionList = (params) => {
  return service({
    url: "/partTimePosition/getList",
    method: "get",
    params,
  });
};
export const AddPartTimePosition = (data) => {
  return service({
    url: "/partTimePosition/add",
    method: "post",
    data,
  });
};
export const deletePartTimePosition = (params) => {
  return service({
    url: "/partTimePosition/delete",
    method: "delete",
    params,
  });
};
export const getJobList = (params) => {
  return service({
    url: "/partTimePosition/getJobList",
    method: "get",
    params,
  });
};
export const createJob = (data) => {
  return service({
    url: "/partTimePosition/createJob",
    method: "post",
    data,
  });
};
export const updateJob = (data) => {
  return service({
    url: "/partTimePosition/updateJob",
    method: "post",
    data,
  });
};
export const deleteJob = (params) => {
  return service({
    url: "/partTimePosition/deleteJob",
    method: "delete",
    params,
  });
};
