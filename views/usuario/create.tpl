{{template "base.html" .}}
{{define "content"}}
<form method="post" onsubmit="return save()">
    Nome: <input type="text" class="form-control" name="nome" id="nome" value=""><br>
    Idade: <input type="number" class="form-control" name="idade" id="idade" value=""><br>
    <button type="submit" class="btn btn-default">Salvar</button>
    <button type="button" class="btn btn-primary" onclick="window.location='/usuario/index'">Voltar</button>
</form>

<script>

    function save() {
        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function () {
            if (this.readyState == 4 && this.status == 201) {
                window.location = '/usuario/index';
            } else if (this.readyState == 4 && this.status != 201) {
                alert('Ocorreu um erro ao salvar o usuário');
                return false;
            }
        }

        var usuario = {
            nome: document.getElementById('nome').value,
            idade: document.getElementById('idade').value
        }

        xhttp.open('POST', '/usuario/save', true);
        xhttp.setRequestHeader('Content-Type', 'application/json');
        xhttp.send(JSON.stringify(usuario));
        return false;
    }

</script>
{{end}}