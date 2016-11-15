$(document).on("click", "#daily-menu-item", function(event) {
  $(this).addClass("selected");
  $("#monthly-menu-item").removeClass("selected");
  $("#future-menu-item").removeClass("selected");
  $("#daily-container").removeClass("hide");
  $("#monthly-container").addClass("hide");
  $("#future-container").addClass("hide");
});

$(document).on("click", "#monthly-menu-item", function(event) {
  $(this).addClass("selected");
  $("#daily-menu-item").removeClass("selected");
  $("#future-menu-item").removeClass("selected");
  $("#daily-container").addClass("hide");
  $("#monthly-container").removeClass("hide");
  $("#future-container").addClass("hide");
});

$(document).on("click", "#future-menu-item", function(event) {
  $(this).addClass("selected");
  $("#daily-menu-item").removeClass("selected");
  $("#monthly-menu-item").removeClass("selected");
  $("#daily-container").addClass("hide");
  $("#monthly-container").addClass("hide");
  $("#future-container").removeClass("hide");
});


