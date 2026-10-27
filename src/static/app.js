async function listar_clientes() {
    try {
            const res = await fetch("/cliente")
            if (!res.ok) throw new Error("HTTP " + res.status);
            const lista_clientes = await res.json()

            
            const clientes_html = document.getElementById("clientes")
            lista_clientes.forEach(c => {
                const cli = document.createElement("li");
                cli.textContent = `${c.id} ${c.nombre} ${c.apellido} - Deuda: ${c.deuda} - Nro. teléfono: ${c.nro_telefono}`;
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

const crearClienteForm = document.getElementById("crear-cliente-form");
crearClienteForm.addEventListener("submit", async (event) => {
    event.preventDefault(); // Evita que la pagina se recargue al enviar el formulario
    
    const formData = new FormData(crearClienteForm);
    const clienteData = {
        nombre: formData.get("name"),
        apellido: formData.get("surname"),
        deuda: parseFloat(formData.get("debt")),
        nro_telefono: Number(formData.get("phone"))
    };
    
    try {
        const response = await fetch("/cliente", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(clienteData)
        });
        
        if (!response.ok) {
            throw new Error("Error al crear el cliente");
        }
        
        alert("Cliente creado exitosamente");
       const clientes_html = document.getElementById("clientes")
            
        const cli = document.createElement("li");
        const datos = await response.json()
        cli.textContent = `${datos.id} ${datos.nombre} ${datos.apellido} - Deuda: ${datos.deuda} - Nro. teléfono: ${datos.nro_telefono}`;
        clientes_html.appendChild(cli)      
    }
    catch (error) {
        console.error("Error al crear el cliente:", error);
        alert("Hubo un error al crear el cliente. Por favor, inténtalo de nuevo.");
    }
});

const buscarClienteForm = document.getElementById("buscar-cliente-form");
buscarClienteForm.addEventListener("submit", async (event) => {
    event.preventDefault(); //  evita que la página se recargue

    const formData = new FormData(buscarClienteForm);
    const id = Number(formData.get("id_cliente"));


    try{
        const response = await fetch("/cliente/" + id);
            if (!response.ok) {
            throw new Error("error al buscar cliente")
            }
        const c = await response.json();
        alert(`${c.nombre} ${c.apellido} - Deuda: ${c.deuda} - Nro. teléfono: ${c.nro_telefono}`);
    } catch (error) {
        console.error("error al buscar cliente", error);
        alert("Hubo un error al buscar el cliente: " + error.message);
    }

});

const actualizarClienteForm = document.getElementById("")