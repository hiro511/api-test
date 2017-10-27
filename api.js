
function getTimetable(){
  $.getJSON("http://localhost:8080/timetable", function( data ) {
    console.log(data);
  });
}
