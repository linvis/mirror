import request from "@/utils/request";

export function queryLeetcodeRec(token) {
  return request({
    url: "/record/query/leetcode",
    method: "get",
    params: { token }
  });
}

export function queryGithubRec(token) {
  return request({
    url: "/record/query/github",
    method: "get",
    params: { token }
  });
}
