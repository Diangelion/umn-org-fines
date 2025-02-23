window.AuthModules = {
  getAccessToken: async () =>
    await window.StorageModules.getWithExpiry("access_token"),
  getRefreshToken: async () =>
    await window.StorageModules.getWithExpiry("refresh_token"),
};

function handleLogout() {
  localforage.removeItem("access_token");
  localforage.removeItem("refresh_token");
  window.location.href = "/login";
}
