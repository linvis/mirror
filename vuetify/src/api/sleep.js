import request from "@/utils/request";

export function submitSleepRec(data) {
  return request({
    url: "/record/submit/sleep",
    method: "post",
    data
  });
}

// export function querySleepRec(token) {
//     return request({
//         url: '/record/query/sleep',
//         method: 'get',
//         params: { token }
//     })
// }

export function querySleepRec(token) {
  return request({
    url: "/record/query/sleep",
    method: "get",
    params: { token }
  });
}

export function querySleepRecAnalysis(token) {
  return request({
    url: "/record/analysis/sleep",
    method: "get",
    params: { token }
  });
}
