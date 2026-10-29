async function listar_clientes() {
    try {
            const res = await fetch("/cliente")
            if (!res.ok) throw new Error("HTTP " + res.status);
            const lista_clientes = await res.json()

            const tbody = document.getElementById("cuerpoTablaCliente")
            tbody.innerHTML = ""

            lista_clientes.forEach(c => {
                const fila =
                    document.createElement("tr")
                    fila.dataset.id = c.id
                    
                    fila.innerHTML = `<td> ${c.id}</td>
                    <td>${c.nombre}</td>
                    <td>${c.apellido}</td>
                    <td>${c.deuda}$</td>
                    <td>${c.nro_telefono}</td>
                    <td><button class = "delete-cliente-tabla"> Eliminar </button></td>
                `; 
                tbody.appendChild(fila);
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

            
            const tbody = document.getElementById("cuerpoTablaTratamiento")
            tbody.innerHTML = ""

            lista_tratamientos.forEach(t => {
                const fila =
                    document.createElement("tr")
                    fila.dataset.id = t.id
                    
                    fila.innerHTML = `<td> ${t.id}</td>
                    <td>${t.nombre}</td>
                    <td>${t.descripcion_corta}</td>
                    <td>${t.costo}$</td>
                    <td><button class = "delete-tratamiento-tabla"> Eliminar </button></td>
                `; 
                tbody.appendChild(fila);
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

            
            const tbody = document.getElementById("cuerpoTablaTurnos")
            tbody.innerHTML = ""

            lista_turnos.forEach(t => {
                const fila =
                    document.createElement("tr")
                    fila.dataset.id = t.id// AGREGAR FECHA Y HORA QUE SE PUEDA MANEJAR CORRECTAMENTE
                    
                    fila.innerHTML = `<td> ${t.fecha}</td>
                    <td>${t.hora}</td>
                    <td>${t.nombre_cliente}</td>
                    <td>${t.nombre_tratamiento}</td>
                    <td><button class = "delete-turno-tabla"> Eliminar </button></td>
                `; 
                tbody.appendChild(fila);
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

            const tbody = document.getElementById("cuerpoTablaVoucher")
            tbody.innerHTML = ""

            lista_voucher.forEach(v => {
                const fila =
                    document.createElement("tr")
                    fila.dataset.id = v.id_voucher
                    
                    fila.innerHTML = `<td> ${v.id_voucher}</td>
                    <td>${v.regalador}</td>
                    <td>${v.cliente}</td>
                    <td>${v.tratamiento}</td>
                    <td><button class = "delete-voucher-tabla"> Eliminar </button></td>
                `; 
                tbody.appendChild(fila);
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
        console.error("Error al borrar cliente: ", error)
    }
})

const tratamientos = document.getElementById("cuerpoTablaTratamiento")
tratamientos.addEventListener("click",async (e) =>{
    e.preventDefault();
    const boton = e.target.closest(".delete-tratamiento-tabla");
    if(!boton) return;

    const fila = boton.closest("tr");
    const idDelete= fila.dataset.id;
    try {
        const res = await fetch(`/tratamiento/${idDelete}`,{
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            }
        });
        if (!res.ok) {
            throw new Error(res.text())
        }
        tratamientos.removeChild(fila)
    } catch (error){
        console.error("Error al borrar tratamiento: ", error)
    }
})

const vouchers = document.getElementById("cuerpoTablaVoucher")
vouchers.addEventListener("click",async (e) =>{
    e.preventDefault();
    const boton = e.target.closest(".delete-voucher-tabla");
    if(!boton) return;

    const fila = boton.closest("tr");
    const idDelete= fila.dataset.id;
    try {
        const res = await fetch(`/voucher/${idDelete}`,{
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            }
        });
        if (!res.ok) {
            throw new Error(res.text())
        }
        vouchers.removeChild(fila)
    } catch (error){
        console.error("Error al borrar voucher: ", error)
    }
})

