$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(event) {
  event.preventDefault()
  if ($('#senha').val() != $('#confirmar-senha').val()) {
    alert('As senhas não coincidem!')
    return
  }

  $.ajax({
    url: '/usuarios',
    method: 'POST',
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      senha: $('#senha').val(),
    },
  })
    .done(function () {
      alert('Usuário criado com sucesso!')
    })
    .fail(function (err) {
      console.log(err)
      alert('Erro ao cadastrar usuário.')
    })
}
