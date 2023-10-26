import service from "@/utils/request";

export const getCountry = () => {
  return service({
    url: "/common/country",
    method: "get",
  });
};
/**
 * 获取省份
 * @param {string} 国家id
 * @returns
 */
export const getProvince = (id) => {
  return service({
    url: "/common/province",
    method: "get",
    params: { countryId: id },
  });
};
/**
 * 获取城市
 * @param {string} 省份id
 * @returns
 */
export const getCity = (id) => {
  return service({
    url: "/common/city",
    method: "get",
    params: { provinceId: id },
  });
};
/**
 * 获取区县
 * @param {string} 城市id
 * @returns
 */
export const getArea = (id) => {
  return service({
    url: "/common/area",
    method: "get",
    params: { cityId: id },
  });
};

/**
 * 获取numberCode
 * @param {string} code
 * @returns
 */
export const getNumberCode = (code) => {
  return service({
    url: "/common/numberCode",
    method: "get",
    params: { code },
  });
};
