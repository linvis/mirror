import request from '@/utils/request'

export function submitSleepRec(data) {
  return request({
    url: '/record/submit/sleep',
    method: 'post',
    data
  })
}

// export function querySleepRec(token) {
//     return request({
//         url: '/record/query/sleep',
//         method: 'get',
//         params: { token }
//     })
// }

export function querySleepRec(time, token) {
  return request({
    url: '/record/query/sleep/' + time,
    method: 'get',
    params: { token }
  })
}
