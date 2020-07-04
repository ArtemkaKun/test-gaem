$(document).ready(function() {
    $(".webgl-content").hide();

    let viewportWidth = $( window ).width();
    let viewportHeight = $( window ).height();
    $("#unityContainer").css({"width": viewportWidth, "height": viewportHeight })

    $(".webgl-content").show();
});