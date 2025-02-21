htmx.config.selfRequestsOnly = false;
htmx.config.withCredentials = true;

const mainElement = document.querySelector("main");
const baseURL = "http://127.0.0.1:3334";
const path = window.location.pathname;

if (mainElement) {
  mainElement.setAttribute("hx-get", `${baseURL}${path}`);
  mainElement.setAttribute("hx-trigger", "load from:body");
  htmx.process(mainElement);
} else {
  console.error("No main element found!");
}
