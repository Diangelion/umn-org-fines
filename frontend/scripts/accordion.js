function toggleOrganizationList() {
  const orgList = document.getElementById("org-list");

  // Toggle the hidden class
  orgList.classList.toggle("hidden");

  // If the list is being shown, fetch the data only if it's empty
  if (
    !orgList.classList.contains("hidden") &&
    orgList.innerHTML.trim() === ""
  ) {
    orgList.dispatchEvent(new Event("htmx:trigger", { bubbles: true }));
  }
}
