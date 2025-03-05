function toggleOrganizationList(element) {
  const imgElement = element.querySelector("img");
  if (imgElement.src.includes("left")) {
    element.classList.add("shadow-sm");
    document.getElementById("organization-list").classList.remove("hidden");
    imgElement.src = imgElement.src.replace("left", "down");
  } else {
    document.getElementById("organization-list").classList.add("hidden");
    element.classList.remove("shadow-sm");
    imgElement.src = imgElement.src.replace("down", "left");
  }
}
