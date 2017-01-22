{{template "base.html" .}}
{{define "content"}}
<form method="delete" onsubmit="return remover()">

    <input type="hidden" id="usuarioId" value="{{.usuario.ID}}" >
    Nome: <label class="control-label">{{.usuario.Nome}}</label><br>
    Idade: <label class="control-label">{{.usuario.Idade}}</label><br>
    
    <button type="submit" class="btn btn-danger">Remover</button>
    <button type="button" class="btn btn-primary" onclick="window.location='/usuario/index'">Voltar</button>

</form>

<script>
    
    function remover() {
        if (!confirm('Você deseja realmente remover o usuário?')) {
            return false;
        }

        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 200) {
                window.location = '/usuario/index';
            } else if (this.readyState == 4) {
                alert('Ocorreu um erro ao remover usuário');
                return false;
            }
        }

        var id = document.getElementById('usuarioId').value;
        xhttp.open('DELETE', '/usuario/delete/' + id, true);
        xhttp.send();
        return false;
    }

</script>
{{end}}