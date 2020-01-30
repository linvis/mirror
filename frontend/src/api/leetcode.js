
import request from '@/utils/request'

export function queryLeetcodeRec(token) {
  return request({
    url: '/record/query/leetcode/',
    method: 'get',
    params: { token }
  })
}
