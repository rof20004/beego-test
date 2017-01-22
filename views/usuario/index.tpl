{{template "base.html" .}}
{{define "content"}}
<button type="button" class="btn btn-default" onclick="javascript:window.location='/usuario/create'">Ir para cadastro</button>

<br><br>

<table class="table table-hover">
    <thead>
        <tr>
            <th>Nome</th>
            <th>Idade</th>
            <th>Editar</th>
            <th>Visualizar</th>
        </tr>
    </thead>
    <tbody id="list-usuario">
        
    </tbody>
</table>

<script>
    
    window.onload = getUsuarios();

    function getUsuarios() {
        var xhttp = new XMLHttpRequest();
        xhttp.onreadystatechange = function () {
            if (this.readyState == 4 && this.status == 200) {
                var usuarios = JSON.parse(this.responseText);
                var listUsuario = document.getElementById('list-usuario');
                for (var i = 0; i < usuarios.length; i++) {
                    var tr = document.createElement('tr');
                    
                    var tdNome = document.createElement('td');
                    tdNome.textContent = usuarios[i].nome;

                    var tdIdade = document.createElement('td');
                    tdIdade.textContent = usuarios[i].idade;
                    
                    var tdEditar = document.createElement('td');
                    var aEdit = document.createElement('a');
                    aEdit.textContent = 'Edit';
                    aEdit.setAttribute('href', '/usuario/edit/' + usuarios[i].id);
                    tdEditar.appendChild(aEdit);

                    var tdVisualizar = document.createElement('td');
                    var aView = document.createElement('a');
                    aView.textContent = 'View';
                    aView.setAttribute('href', '/usuario/view/' + usuarios[i].id);
                    tdVisualizar.appendChild(aView);

                    tr.appendChild(tdNome);
                    tr.appendChild(tdIdade);
                    tr.appendChild(tdEditar);
                    tr.appendChild(tdVisualizar);

                    listUsuario.appendChild(tr);
                }
            }
        }

        xhttp.open('GET', '/usuario/list', true);
        xhttp.send();
    }

</script>
{{end}}