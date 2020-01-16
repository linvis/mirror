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

export function submitRec(evt, data) {
  return request({
    url: '/record/submit/' + evt,
    method: 'post',
    data
  })
}

export function queryRec(evt, token) {
  return request({
    url: '/record/query/' + evt,
    method: 'get',
    params: { token }
  })
}
