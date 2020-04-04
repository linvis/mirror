export function validUsername(str) {
  const valid_map = ["admin", "jarvis"];
  return valid_map.indexOf(str.trim()) >= 0;
}
