function triggerEditProfileForm() {
  if (document.getElementById("edit-profile-form").style.display === "block")
    return;
  document.getElementById("edit-profile-form").style.display = "block";
}
