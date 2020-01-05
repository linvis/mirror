import request from '@/utils/request'

export function submit(data) {
  return request({
    url: '/evt/new',
    method: 'post',
    data
  })
}

export function getEvtType(token) {
  return request({
    url: '/evt/type',
    method: 'get',
    params: { token }
  })
}

export function getSleepData(token) {
  return request({
    url: '/evt/sleep',
    method: 'get',
    params: { token }
  })
}
