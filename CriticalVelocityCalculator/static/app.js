function insertChat(who, text, time = 0){
  text.split('\n').forEach(function(x) {
    if (x == "")
      return;
    var elem =
        '<li class="mdl-list__item mdl-list__item--three-line">' +
          '<span class="mdl-list__item-primary-content">' +
            '<i class="material-icons mdl-list__item-avatar">' +
             ((who == "me") ? 'face' : 'person') +
            '</i>' +
            '<span>' + ((who == "me") ? '自己' : '機器人')+ '</span>' +
              '<span class="mdl-list__item-text-body">' +
                 x +
              '</span>' +
            '</span>' +
        '</li>';

    setTimeout(
      function(){
        $("#chat").append(elem);
      }, time);
  });
}


function scrollToBottom() {
  var chat_area = $('.chat-area');
  chat_area.animate({scrollTop: chat_area.prop('scrollHeight')});
}

function chat(text) {
  if (text == "")
    return;
  insertChat("me", text);
  scrollToBottom();
  $.post('/ask', {'in': text})
   .done(function(response) {
     insertChat("you", response, 500);
     scrollToBottom();
   });
}

$('#go').on('click', function(e) {
  var text = $('#input').val();
  $('#input').val('');
  chat(text);
})

$("#input").on("keyup", function(e){
  if (e.which !== 13)
    return;
  var text = $(this).val();
  $(this).val('');
  chat(text);
});

$("#chat").empty();

insertChat("you", "你好，請告訴我你的問題");

