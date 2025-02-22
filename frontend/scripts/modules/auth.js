window.AuthModules = {
  getAccessToken: async () =>
    await window.StorageModules.getFromLocalForage("access_token"),
  getRefreshToken: async () =>
    await window.StorageModules.getFromLocalForage("refresh_token"),
};

function isTokenExpired(token) {
  try {
    const payload = JSON.parse(atob(token.split(".")[1]));
    return payload.exp * 1000 < Date.now();
  } catch (e) {
    return true;
  }
}

function handleLogout() {
  localforage.removeItem("access_token");
  localforage.removeItem("refresh_token");
  window.location.href = "/login";
}
