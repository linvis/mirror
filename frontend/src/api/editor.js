import request from '@/utils/request'

export function queryEditorCatalog(token) {
  return request({
    url: '/editor/catalog',
    method: 'get',
    params: { token }
  })
}

export function updateEditorCatalog(data) {
  return request({
    url: '/editor/catalog/update',
    method: 'post',
    data
  })
}
