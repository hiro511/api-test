
function getTimetable(){
  $.ajax({
    url: "http://localhost:8080/timetable",
    dataType: "json",
    type: "GET",
    ifModified: true
  }, function(data) {
    console.log(data);
  });
}
