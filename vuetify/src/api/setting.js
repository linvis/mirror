import request from "@/utils/request";

export function submitSettingProgram(data) {
  return request({
    url: "/setting/program",
    method: "post",
    data
  });
}

export function querySettingProgram(token) {
  return request({
    url: "/setting/program",
    method: "get",
    params: { token }
  });
}
