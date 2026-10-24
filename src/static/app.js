async function listar_clientes() {
    try {
            const res = await fetch("/cliente")
            if (!res.ok) throw new Error("HTTP " + res.status);
            const lista_clientes = await res.json()

            
            const clientes_html = document.getElementById("clientes")
            lista_clientes.forEach(c => {
                const cli = document.createElement("li");
                cli.textContent = `${c.nombre} ${c.apellido} - Deuda: ${c.deuda} - Nro. teléfono: ${c.nro_telefono}`;
                clientes_html.appendChild(cli)
            });          
        } catch (err) {
            console.error("Error en fetch:", err);
        }
}

async function listar_tratamientos() {
    try {
            const res = await fetch("/tratamiento")
            if (!res.ok) throw new Error("HTTP " + res.status);
            const lista_tratamientos = await res.json()

            
            const tratamientos_html = document.getElementById("tratamientos")
            lista_tratamientos.forEach(c => {
                const tratamiento = document.createElement("li");
                tratamiento.textContent = `${c.nombre}: ${c.descripcion_corta} - ${c.costo}`;
                tratamientos_html.appendChild(tratamiento)
            });          
        } catch (err) {
            console.error("Error en fetch:", err);
        }
}

async function listar_turnos() {
    try {
            const res = await fetch("/turno")
            if (!res.ok) throw new Error("HTTP " + res.status);
            const lista_turnos = await res.json()

            
            const turnos_html = document.getElementById("turnos")
            lista_turnos.forEach(async t => {
                const turno = document.createElement("li");
                // Habría que evitar
                const cliente = await fetch(`/cliente/${t.cliente_id}`)
                const cli = await cliente.json()
                const tratamiento = await fetch(`/tratamiento/${t.tratamiento_id}`)
                const tra = await tratamiento.json()
                const contenido = `${t.fecha} ${t.hora}hs - ${cli.nombre} ${cli.apellido} - ${tra.nombre}`
                turno.textContent = contenido;
                turnos_html.appendChild(turno)
            });          
        } catch (err) {
            console.error("Error en fetch:", err);
        }
}

async function listar_vouchers(){
    
}
listar_clientes()
listar_tratamientos()
listar_turnos()