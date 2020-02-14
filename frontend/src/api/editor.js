import request from '@/utils/request'

export function queryEditorCatalog(token) {
  return request({
    url: '/editor/catalog',
    method: 'get',
    params: { token }
  })
}
