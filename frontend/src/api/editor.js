import request from '@/utils/request'

export function queryEditorCatalog(token) {
  return request({
    url: '/editor/catalog',
    method: 'get',
    params: { token }
  })
}

export function submitNewCatalog(data) {
  return request({
    url: '/editor/catalog/new',
    method: 'post',
    data
  })
}
