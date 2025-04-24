import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/api/v1/server/list',
    method: 'get',
    params
  })
}
export function updateServer(data) {
  return request({
    url: '/api/v1/server/upd',
    method: 'post',
    data
  })
}
export function delServer(data) {
  return request({
    url: '/api/v1/server/del',
    method: 'post',
    data
  })
}
export function addServer(data) {
  return request({
    url: '/api/v1/server/add',
    method: 'post',
    data
  })
}
export function getOneByID(ip) {
  return request({
    url: `/api/v1/server/get/${ip}`,
    method: 'get'
  })
}
