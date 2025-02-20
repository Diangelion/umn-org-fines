htmx.config.selfRequestsOnly = false;

const mainElement = document.querySelector("main");
const baseURL = "http://localhost:3334";
const path = window.location.pathname;

if (mainElement) {
  mainElement.setAttribute("hx-get", `${baseURL}${path}`);
  mainElement.setAttribute("hx-trigger", "load");
  mainElement.setAttribute("hx-swap", "innerHTML");
  mainElement.setAttribute("hx-target", "main");
  htmx.process(mainElement);
} else {
  console.error("No main element found!");
}
