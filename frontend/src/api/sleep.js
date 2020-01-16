
import request from '@/utils/request'

export function getSleepRec(token) {
  return request({
    url: '/record/sleep/query',
    method: 'get',
    params: { token }
  })
}

export function submitSleepRec(data) {
  return request({
    url: '/record/sleep/new',
    method: 'post',
    data
  })
}
