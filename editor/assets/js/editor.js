$(function() {
    // init editor
    var editor = ace.edit("editor");
    var autosave;
    editor.setTheme("ace/theme/tomorrow");
    editor.getSession().setMode("ace/mode/yaml");

    // post call to save content of the playground
    function saveEditor() {
      $.post("/save/{{.SketchName}}", {content: editor.getValue()}, function(data) {
	  clearTimeout(autosave);
	  $("#action-save").attr("class", "btn btn-success").find("i").attr("class", "glyphicon glyphicon-saved");
	});
    }
        
    // bind function for autosave on editor keyup
    $("#editor").keyup(function() {
	$("#action-save").attr("class", "btn btn-danger").find("i").attr("class", "glyphicon glyphicon-save");
	if (autosave !== undefined) {
	  clearTimeout(autosave);
	}
	autosave = setTimeout(saveEditor, 5000);
      });

    // bind function on save button
    $("#action-save").click(function() {
	saveEditor();
      });

    // bind function on share button
    $("#action-share").click(function() {
	window.location = "mailto:?subject=J'aimerais partager ce fichier LIA&body=Voici le lien : " + window.location;
      });

    // bind function on + and - buttons
    $("#font-plus, #font-minus").click(function() {
	$("#editor").css("font-size", (parseFloat($("#editor").css("font-size"))+parseFloat($(this).attr("data-value")))+"px");
      });

    // add tooltips to buttons
    $(".actions button, .actions a").tooltip();
  });
