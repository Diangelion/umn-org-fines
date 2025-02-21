import { getFromLocalForage } from "./scripts";

function isTokenExpired(token) {
  try {
    const payload = JSON.parse(atob(token.split(".")[1]));
    return payload.exp * 1000 < Date.now();
  } catch (e) {
    return true;
  }
}

export function getAccessToken() {
  const accessToken = JSON.parse(getFromLocalForage("access_token"));
  return accessToken;
}

export function getRefreshToken() {
  const refreshToken = JSON.parse(getFromLocalForage("refresh_token"));

  // If no refresh token, force logout
  if (!refreshToken || isTokenExpired(refreshToken)) {
    handleLogout();
    return null;
  }

  return refreshToken;
}

function handleLogout() {
  localforage.removeItem("access_token");
  localforage.removeItem("refresh_token");
  window.location.href = "/login";
}
