import request from '@/utils/request'

export function submit(data) {
  return request({
    url: '/activity/new',
    method: 'post',
    data
  })
}

export function getActType(token) {
  return request({
    url: '/activity/type',
    method: 'get',
    params: { token }
  })
}
