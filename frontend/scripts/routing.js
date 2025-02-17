const baseURL = "http://localhost:3334";
const mainElement = document.getElementsByTagName("main")?.[0];

if (mainElement) {
  const path = window.location.pathname;
  htmx.ajax("GET", `${baseURL}${path}`, { target: "main", swap: "innerHTML" });
}
