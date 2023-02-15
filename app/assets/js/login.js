$('#login').on('submit', fazerLogin)

function fazerLogin(event) {
  event.preventDefault()

  $.ajax({
    url: '/login',
    method: 'POST',
    data: {
      email: $('#email').val(),
      senha: $('#senha').val(),
    },
  })
    .done(function () {
      console.log('DEU CERTO')
      window.location = '/home'
    })
    .fail(function (e) {
      console.log(e)
      alert('Usuário ou senha inválidos!')
    })
}
