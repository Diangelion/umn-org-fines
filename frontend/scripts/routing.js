htmx.config.selfRequestsOnly = false;
// htmx.config.withCredentials = true;

// Set headers dynamically (HTMX events)
document.body.addEventListener("htmx:configRequest", async (event) => {
  const accessToken = await window.AuthModules.getAccessToken();
  const refreshToken = await window.AuthModules.getRefreshToken();
  console.log("refreshToken:", refreshToken);
  console.log("accessToken:", accessToken);
  event.detail.headers["Authorization"] = accessToken;
  event.detail.headers["X-Refresh-Token"] = refreshToken;
});

// Routing start

const mainElement = document.querySelector("main");
const baseURL = "http://127.0.0.1:3334";
const path = window.location.pathname;

if (mainElement) {
  mainElement.setAttribute("hx-get", `${baseURL}${path}`);
  mainElement.setAttribute("hx-trigger", "load once from:body");
  htmx.process(mainElement);
} else {
  console.error("No main element found!");
}
