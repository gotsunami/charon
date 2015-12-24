$(function() {
    // init editor
    var editor = ace.edit("editor");
    var autosave = undefined;
    editor.setTheme("ace/theme/tomorrow");
    editor.getSession().setMode("ace/mode/yaml");

    // post call to save content of the playground
    function saveEditor() {
      $.post("/save/"+sketchName, {content: editor.getValue()}, function(data) {
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
	return false;
      });

    // bind function on save button
    $("#action-save").click(function() {
	saveEditor();
	return false;
      });

    // bind function on share button
    $("#action-share").click(function() {
	window.location = "mailto:?subject=J'aimerais partager ce fichier LIA&body=Voici le lien : " + window.location;
	return false;
      });

    // bind function on + and - buttons
    $("#font-plus, #font-minus").click(function() {
	$("#editor").css("font-size", (parseFloat($("#editor").css("font-size"))+parseFloat($(this).attr("data-value")))+"px");
	return false;
      });

    // bind function for fullscreen mode
    $('#editor-full').click(function() {
	$('#editor').fullscreen();
	return false;
      });

    // bind function for fullscreen mode
    $('#editor-clear').click(function() {
	editor.setValue("models:");
	return false;
      });

    // add tooltips to buttons
    $(".actions button, .actions a").tooltip();

    $("#example iframe").load(function() {
	    $(this).parent().fadeIn();
	});
    $("#example .close").click(function() {
	    $(this).parent().fadeOut();
	});
    
    $("a[data-example]").click(function() {
	    $("#example iframe").attr("src", "example/"+$(this).attr("data-example"));
	});
  });

