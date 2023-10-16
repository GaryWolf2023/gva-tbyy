import service from "@/utils/request";

/**
 * 上传文件到minio
 * @param {formData中需要文件参数:bucketName,objectName,file} data
 * @returns
 */
export const updateFileToMinio = (data) => {
  return service({
    url: "/minio/updateFile",
    method: "post",
    data,
  });
};

/**
 * 上传文件到minio
 * @param {formData中需要文件参数:bucketName,objectName,file} data
 * @returns
 */
export const getFileFromMinio = (data) => {
  return service({
    url: "/minio/getFile",
    method: "get",
    data,
  });
};

/**
 * 删除文件
 * @param {bucketName,objectName} data
 * @returns
 */
export const deleteFileformMinio = (data) => {
  return service({
    url: "/minio/deleteFile",
    method: "delete",
    data,
  });
};
