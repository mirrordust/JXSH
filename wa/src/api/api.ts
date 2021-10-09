import axios, { AxiosResponse } from "axios";

import { Collection, Credential, Post, Tag, User } from "./model";


// development mode settings
const devMode = (!process.env.NODE_ENV || process.env.NODE_ENV === 'development');
const prefix = devMode ? 'http://localhost:4000/api' : '/api';
if (prefix) {
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

/*
 * api definitions
 */

// login & logout
function login(user: User) {
  return axios({
    url: `${prefix}/auth/sessions`,
    method: 'post',
    data: {
      user,
    },
    validateStatus: function (status) {
      return status === 200 || status === 404;
    },
  });
}

function logout(credential: Credential) {
  return axios({
    url: `${prefix}/auth/sessions`,
    method: 'delete',
    data: {
      token: credential.accessToken,
    },
    validateStatus: function (status) {
      return status === 204 || status === 404;
    },
  });
}

// DB model crud
function getAll(
  credential: Credential,
  modelName: 'posts' | 'tags' | 'collections'
) {
  return axios({
    url: `${prefix}/cms/${modelName}`,
    method: 'get',
    validateStatus: function (status) {
      return status === 200 || status === 401;
    },
    headers: { 'Authorization': `BASIC ${credential.accessToken}` },
  });
}

function getById(
  credential: Credential,
  modelName: 'posts' | 'tags' | 'collections',
  id: number
) {
  return axios({
    url: `${prefix}/cms/${modelName}/${id}`,
    method: 'get',
    validateStatus: function (status) {
      return status === 200 || status === 401 || status === 404;
    },
    headers: { 'Authorization': `BASIC ${credential.accessToken}` },
  });
}

function create(
  credential: Credential,
  modelName: 'posts' | 'tags' | 'collections',
  model: Post | Tag | Collection
) {

}

function updateById(
  credential: Credential,
  modelName: 'posts' | 'tags' | 'collections',
  model: Post | Tag | Collection
) {

}

function deleteById(
  credential: Credential,
  modelName: 'posts' | 'tags' | 'collections',
  id: number
) {

}

// images api

// export
const Api = {
  login,
  logout,
  getAll,
  getById,
  create,
  updateById,
  deleteById,
};

export default Api;

// utils
function handleResponseError(apiName: string, response: AxiosResponse, expectedStatus: number): string | undefined {
  if (response.status !== expectedStatus) {
    return `[${apiName}] expected status ${expectedStatus}, but got ${response.status}`;
  }
  if (response.data.errors) {
    return `[${apiName}] server returns errors: ${response.data.errors}`;
  }
}

function handleApiExecutionError(apiName: string, error: unknown): string {
  return `[API ${apiName} error] ${error}`;
}
