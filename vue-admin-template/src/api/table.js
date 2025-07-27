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
export function countServer() {
  return request({
    url: '/api/v1/server/count',
    method: 'get'
  })
}
// WebSocket 连接封装
export function initWebSocket() {
  return new WebSocket(process.env.VUE_APP_WS_API + '/api/v1/ssh/connect')
}

export function wsSSHConnect(host, username, password) {
  return function wsSSHConnectNew() {
    this.socket = initWebSocket()
    // this.socket = new WebSocket('ws://localhost:8080/api/v1/ssh/connect')
    this.socket.onopen = () => {
      // 发送SSH连接参数
      this.socket.send(JSON.stringify({
        host: host,
        username: username,
        password: password // 实际项目应使用加密传输
      }))
    }

    this.socket.onmessage = (event) => {
      this.term.write(event.data)
    }

    this.socket.onclose = () => {
      sessionStorage.removeItem('sshParams')
      this.term.writeln('\r\n连接已关闭')
    }
  }
}
