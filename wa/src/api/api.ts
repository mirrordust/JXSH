import axios, { AxiosResponse } from "axios";

import { Post, Tag, Collection, User, Credential } from "./model";


/* development mode settings */
const devMode = (!process.env.NODE_ENV || process.env.NODE_ENV === 'development');
const prefix = devMode ? 'http://localhost:4000/api' : process.env.REACT_APP_PREFIX;
if (devMode) {
  // axios拦截器
  axios.interceptors.request.use(function (config) {
    console.log('请求参数：', config);
    return config;
  }, function (error) {
    console.error('请求错误：', error);
    return Promise.reject(error);
  });
  axios.interceptors.response.use(function (response) {
    // 2xx 范围内的状态码都会触发该函数
    console.log('返回结果：', response);
    return response;
  }, function (error) {
    // 超出 2xx 范围的状态码都会触发该函数
    console.error('返回错误：', error);
    return Promise.reject(error);
  });
}

/* api definitions */
type ModelName = 'posts' | 'tags' | 'collections';
type ModelType = Post | Tag | Collection;

// login & logout
function login(user: User) {
  return axios({
    url: `${prefix}/auth/sessions`,
    method: 'post',
    data: { user },
    validateStatus: status => status === 200 || status === 404
  });
}

function logout(credential: Credential) {
  return axios({
    url: `${prefix}/auth/sessions`,
    method: 'delete',
    data: { token: credential.access_token },
    validateStatus: status => status === 204 || status === 404
  });
}

// DB model crud
function getAll(credential: Credential, modelName: ModelName) {
  return axios({
    url: `${prefix}/cms/${modelName}`,
    method: 'get',
    validateStatus: status => status === 200 || status === 401,
    headers: { 'Authorization': `BASIC ${credential.access_token}` },
  });
}

function getById(credential: Credential, modelName: ModelName, id: number) {
  return axios({
    url: `${prefix}/cms/${modelName}/${id}`,
    method: 'get',
    validateStatus: status => status === 200 || status === 401 || status === 404,
    headers: { 'Authorization': `BASIC ${credential.access_token}` },
  });
}

function create(credential: Credential, modelName: ModelName, model: ModelType) {
  return axios({
    url: `${prefix}/cms/${modelName}`,
    method: 'post',
    data: { [modelName.slice(0, -1)]: model },
    validateStatus: status => status === 201 || status === 422,
    headers: { 'Authorization': `BASIC ${credential.access_token}` },
  });
}

function updateById(credential: Credential, modelName: ModelName, id: number, model: ModelType) {
  return axios({
    url: `${prefix}/cms/${modelName}/${id}`,
    method: 'patch',
    data: { [modelName.slice(0, -1)]: model },
    validateStatus: status => status === 200 || status === 422,
    headers: { 'Authorization': `BASIC ${credential.access_token}` },
  });
}

function deleteById(credential: Credential, modelName: ModelName, id: number) {
  return axios({
    url: `${prefix}/cms/${modelName}/${id}`,
    method: 'delete',
    validateStatus: status => status === 204,
    headers: { 'Authorization': `BASIC ${credential.access_token}` },
  });
}

// images api


export const Api = {
  login,
  logout,
  getAll,
  getById,
  create,
  updateById,
  deleteById,
};

export const initialApiStatus: {
  status: 'idle' | 'loading' | 'succeeded' | 'failed';
  error: string | undefined;
} = {
  status: 'idle',
  error: undefined
};

export type ApiStatus = typeof initialApiStatus;

/* API util functions */
export function handleValidateResponse(expectedStatus: number, response: AxiosResponse) {
  if (response.status === expectedStatus) {
    return response.data;
  }
  throw Error(JSON.stringify(response.data));
}
