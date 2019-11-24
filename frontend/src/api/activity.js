import request from '@/utils/request'

export function submit(data) {
  return request({
    url: '/activity/new',
    method: 'post',
    data
  })
}
