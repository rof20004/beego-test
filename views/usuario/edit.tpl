{{template "base.html" .}}
{{define "content"}}
<form method="post" onsubmit="return update()">
    <input type="hidden" id="usuarioId" value="{{.usuario.ID}}">
    Id: {{.usuario.ID}}<br>
    Nome: <input type="text" class="form-control" name="nome" id="nome" value="{{.usuario.Nome}}"><br>
    Idade: <input type="number" class="form-control" name="idade" id="idade" value="{{.usuario.Idade}}"><br>
    <button type="submit" class="btn btn-default">Atualizar</button>
    <button type="button" class="btn btn-primary" onclick="window.location='/usuario/index'">Voltar</button>
</form>

<script type="text/javascript">

    function update() {
        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function () {
            if (this.readyState == 4 && this.status == 200) {
                window.location = '/usuario/index';
            } else if (this.readyState == 4 && this.status != 200) {
                alert('Ocorreu um erro ao atualizar o usuário');
                return false;
            }
        }

        var usuario = {
            id: document.getElementById('usuarioId').value,
            nome: document.getElementById('nome').value,
            idade: document.getElementById('idade').value
        }

        xhttp.open('PUT', '/usuario/edit/' + usuario.id, true);
        xhttp.setRequestHeader('Content-Type', 'application/json');
        xhttp.send(JSON.stringify(usuario));
        return false;
    }

</script>
{{end}}