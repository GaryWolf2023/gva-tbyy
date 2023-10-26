import axios from "axios";

const getPrestToken = () => {
  return "123";
};

const prestd = axios.create({
  timeout: 3000,
  baseURL: process.env.VUE_APP_PREST_API,
});

let count = 0;
// gettoken
function getToken_() {
  const hasToken = getPrestToken();
  if (hasToken) {
    count = 0;
    return hasToken;
  } else {
    if (count <= 5) {
      const timer = setTimeout(() => {
        clearTimeout(timer);
        count++;
        getToken_();
      }, 1000);
    } else {
      return null;
    }
  }
}

// 请求拦截器
prestd.interceptors.request.use(
  (config) => {
    const hasToken = getToken_();
    if (hasToken) {
      config.headers.Authorization = hasToken;
      return config;
    } else {
      return;
    }
  },
  (error) => {
    return Promise.reject(error);
  }
);
prestd.interceptors.response.use((response) => {
  return response.data;
});

export default prestd;
