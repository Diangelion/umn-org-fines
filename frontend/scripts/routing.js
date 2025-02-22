htmx.config.selfRequestsOnly = false;
// htmx.config.withCredentials = true;

(async () => {
  const mainElement = document.querySelector("main");
  const baseURL = "http://127.0.0.1:3334";
  const path = window.location.pathname;

  if (mainElement) {
    const [accessToken, refreshToken] = await Promise.all([
      window.AuthModules.getAccessToken(),
      window.AuthModules.getRefreshToken(),
    ]);
    const headerSent = {
      Authorization: accessToken,
      "X-Refresh-Token": refreshToken,
    };
    mainElement.setAttribute("hx-headers", JSON.stringify(headerSent));
    mainElement.setAttribute("hx-get", `${baseURL}${path}`);
    mainElement.setAttribute("hx-trigger", "load once from:body");
    htmx.process(mainElement);
  } else {
    console.error("No main element found!");
  }
})();
