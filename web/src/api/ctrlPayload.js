import service from "@/utils/request";

/**
 * 创建payload
 * @param {string} job_id
 * @param {string} unique
 * @param {string} business
 * @param {object} {order:Array,payload:string}
 * @returns
 */
export const createPayload = (obj) => {
  return service({
    url: "/payload/createPayload",
    method: "POST",
    data: obj,
  });
};

/**
 * 获取payload列表(*两个参数必传其一)
 * @param {string} ?job_id
 * @param {string} ?unique
 * @returns
 */
export const getPayloadList = (obj) => {
  return service({
    url: "/payload/getPayloadList",
    method: "POST",
    data: obj,
  });
};

/**
 * 获取payload列表(*两个参数必传其一)
 * @param {string} ?payload_id
 * @param {string} ?job_id
 * @param {string} ?payload_content
 * @returns
 */
export const updatePayload = (obj) => {
  return service({
    url: "/payload/updatePayload",
    method: "POST",
    data: obj,
  });
};
/**
 * 获取payload列表
 * @param {string} ?payload_id
 * @returns
 */
export const getPayloadById = (obj) => {
  return service({
    url: "/payload/updatePayload",
    method: "POST",
    data: obj,
  });
};
