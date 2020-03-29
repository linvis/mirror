import request from "@/utils/request";

export function queryEditorCatalog(token) {
  return request({
    url: "/editor/catalog",
    method: "get",
    params: { token }
  });
}

export function updateEditorCatalog(data) {
  return request({
    url: "/editor/catalog/update",
    method: "post",
    data
  });
}

export function saveDocument(data) {
  return request({
    url: "/editor/document/new",
    method: "post",
    data
  });
}

export function queryAllDocument(token) {
  return request({
    url: "/editor/document/queryall",
    method: "get"
  });
}

export function queryDocumentByID(id, token) {
  return request({
    url: "/editor/document/query/" + id,
    method: "get"
  });
}
