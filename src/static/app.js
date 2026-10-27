async function listar_clientes() {
    try {
            const res = await fetch("/cliente")
            if (!res.ok) throw new Error("HTTP " + res.status);
            const lista_clientes = await res.json()

            const tbody = document.getElementById("cuerpoTablaCliente")
            tbody.innerHTML = ""

            const clientes_html = document.getElementById("clientes")
            lista_clientes.forEach(c => {
                const fila =
                    document.createElement("tr")
                    fila.dataset.id = c.id
                    
                    fila.innerHTML = `<td> ${c.id}</td>
                    <td>${c.nombre}</td>
                    <td>${c.apellido}</td>
                    <td>$${c.deuda}</td>
                    <td>${c.nro_telefono}</td>
                    <td><button class = "delete-cliente-tabla"> Eliminar </button></td>
                `; 
                tbody.appendChild(fila);
            });
            console.table(lista_clientes);
            
            /*lista_clientes.forEach(c => {
                const cli = document.createElement("li");
                cli.textContent = `${c.id} ${c.nombre} ${c.apellido} - Deuda: ${c.deuda} - Nro. teléfono: ${c.nro_telefono}`;
                clientes_html.appendChild(cli)
            });  */


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
    try {
            const res = await fetch("/voucher")
            if (!res.ok) throw new Error("HTTP " + res.status);
            const lista_voucher = await res.json()
            const vouchers_html = document.getElementById("vouchers")
            lista_voucher.forEach(async v => {
                const voucher = document.createElement("li");
                const regalador = await fetch(`/cliente/${v.cliente_id}`)
                const reg = await regalador.json()
                const receptor = await fetch(`/cliente/${v.cliente_id}`)
                const rec = await receptor.json()
                const tratamiento = await fetch(`/tratamiento/${v.tratamiento_id}`)
                const tra = await tratamiento.json()
                voucher.textContent = `${v.id_voucher} - De: ${reg.nombre} - Para: ${rec.nombre} - Tratamiento: ${tra.nombre}`;
                vouchers_html.appendChild(voucher)
            });
        } catch (err) {
            console.error("Error en fetch:", err);
        }


}

listar_clientes()
listar_tratamientos()
listar_turnos()
listar_vouchers()

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
    

        const tbody = document.getElementById("cuerpoTablaCliente")
        const datos = await response.json()
                const fila =
                    document.createElement("tr")
                    fila.dataset.id = datos.id
                    
                    fila.innerHTML = `<td> ${datos.id}</td>
                    <td>${datos.nombre}</td>
                    <td>${datos.apellido}</td>
                    <td>$${datos.deuda}</td>
                    <td>${datos.nro_telefono}</td>
                    <td><button class = "delete-cliente-tabla"> Eliminar </button></td>
                `; 
                tbody.appendChild(fila); }
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

const clientes = document.getElementById("cuerpoTablaCliente")
clientes.addEventListener("click",async (e) =>{
    e.preventDefault();
    const boton = e.target.closest(".delete-cliente-tabla");
    if(!boton) return;

    const fila = boton.closest("tr");
    const idDelete= fila.dataset.id;
    console.log("ID: ",idDelete)
    try {
        const res = await fetch(`/cliente/${idDelete}`,{
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            }
        });
        if (!res.ok) {
            throw new Error(res.text())
        }
        clientes.removeChild(fila)
    } catch (error){
        console.error("Error al borrar cliente", error)
    }
})

const deleteTratamiento = document.getElementById("deleteTratamiento-button")
deleteTratamiento.addEventListener("click",async (event) =>{
    event.preventDefault();
    const idDelete = 0 ;
    try {
        const res = await fetch(`/tratamiento/${idDelete}`)
        if (!res.ok) {
                throw new Error(res.text)
            }
        
    } catch (error){
        console.error("Error al borrar tratamiento", error)
    }
})

const deleteVoucher = document.getElementById("deleteVoucher-button")
deleteVoucher.addEventListener("click",async (event) =>{
    event.preventDefault();
    const boton = e.target.closest(".delete-cliente-tabla");
    if(!boton) return;

    const fila = boton.closest("tr");
    const idDelete= fila.dataset.id;

    try {
        const res = await fetch(`/voucher/${idDelete}`,{
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            }
        })
        if (!res.ok) {
                throw new Error(res.text)
            }
    } catch (error){
        console.error("Error al borrar voucher", error)
    }
})

const deleteTurno = document.getElementById("deleteTurno-button")
deleteTurno.addEventListener("click",async (event) =>{
    event.preventDefault();
    const fecha = 0 ;
    const hora = 0
    try {
        const res = await fetch(`/turno/${fecha}&${hora}`)
        if (!res.ok) {
                throw new Error(res.text)
            }
    } catch (error){
        console.error("Error al borrar turno", error)
    }
})