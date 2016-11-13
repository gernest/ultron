$(document).on("click", "#daily-menu-item", function(event) {
  $(this).toggleClass("selected");
  $("#monthly-menu-item").toggleClass("selected");
  $("#future-menu-item").toggleClass("selected");
  $("#daily-container").toggleClass("hide");
  $("#monthly-container").toggleClass("hide");
  $("#future-container").toggleClass("hide");
});

$(document).on("click", "#monthly-menu-item", function(event) {
  $(this).toggleClass("selected");
  $("#daily-menu-item").toggleClass("selected");
  $("#future-menu-item").toggleClass("selected");
  $("#daily-container").toggleClass("hide");
  $("#monthly-container").toggleClass("hide");
  $("#future-container").toggleClass("hide");
});

$(document).on("click", "#future-menu-item", function(event) {
  $(this).toggleClass("selected");
  $("#daily-menu-item").toggleClass("selected");
  $("#monthly-menu-item").toggleClass("selected");
  $("#daily-container").toggleClass("hide");
  $("#monthly-container").toggleClass("hide");
  $("#future-container").toggleClass("hide");
});


