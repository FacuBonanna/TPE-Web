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
                    
                    fila.innerHTML = `<td > ${c.id}</td>
                    <td data-campo="nombre">${c.nombre}</td>
                    <td data-campo="apellido">${c.apellido}</td>
                    <td data-campo="deuda">${c.deuda}$</td>
                    <td data-campo="nro_telefono">${c.nro_telefono}</td>
                    <td><button class = "delete-cliente-tabla"> Eliminar </button>
                        <button class = "update-cliente-tabla"> Modificar </button></td>
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
                    <td data-campo="nombre">${t.nombre}</td>
                    <td data-campo="descripcion_corta">${t.descripcion_corta}</td>
                    <td data-campo="costo">${t.costo}$</td>
                    <td><button class = "delete-tratamiento-tabla"> Eliminar </button>
                        <button class = "update-tratamiento-tabla"> Modificar </button></td>
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

            
            const tbody = document.getElementById("cuerpoTablaTurno")
            tbody.innerHTML = ""

            lista_turnos.forEach(t => {
                const fila =
                    document.createElement("tr")
                    fila.dataset.fecha = t.fecha
                    fila.dataset.hora = t.hora
                    fila.innerHTML = `<td data-campo="fecha"> ${t.fecha}</td>
                    <td data-campo="hora">${t.hora}</td>
                    <td data-campo="nombre_tratamiento">${t.nombre_tratamiento}</td>
                    <td data-campo="nombre_cliente">${t.nombre_cliente}</td>
                    <td><button class = "delete-turno-tabla"> Eliminar </button>
                        <button class = "update-turno-tabla"> Modificar </button></td>
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
                    <td data-campo="regalador">${v.regalador}</td>
                    <td data-campo="cliente">${v.cliente}</td>
                    <td data-campo="tratamiento">${v.tratamiento}</td>
                    <td><button class = "delete-voucher-tabla"> Eliminar </button>
                        <button class = "update-voucher-tabla"> Modificar </button></td>
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
                    <td data-campo="nombre">${datos.nombre}</td>
                    <td data-campo="apellido">${datos.apellido}</td>
                    <td data-campo="deuda"> ${datos.deuda}$</td>
                    <td data-campo="nro_telefono">${datos.nro_telefono}</td>
                    <td><button class = "delete-cliente-tabla"> Eliminar </button>
                        <button class = "update-cliente-tabla"> Modificar </button></td>
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

const crearTratamientoForm = document.getElementById("crear-tratamiento-form");
crearTratamientoForm.addEventListener("submit", async (event) => {
    event.preventDefault(); // Evita que la pagina se recargue al enviar el formulario
    
    const formData = new FormData(crearTratamientoForm);
    const tratamientoData = {
        nombre: formData.get("nombre"),
        descripcion_corta: formData.get("descripcion_corta"),
        costo: parseFloat(formData.get("costo")),
    };
    
    try {
        const response = await fetch("/tratamiento", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(tratamientoData)
        });
        
        if (!response.ok) {
            throw new Error(response.text());
        }
        
        alert("Tratamiento creado exitosamente");
    

        const tbody = document.getElementById("cuerpoTablaTratamiento")
        const tr = await response.json()
                const fila =
                document.createElement("tr")
                    fila.dataset.id = tr.id
                    
                    fila.innerHTML = `<td> ${tr.id}</td>
                    <td data-campo="nombre">${tr.nombre}</td>
                    <td data-campo="descripcion_corta">${tr.descripcion_corta}</td>
                    <td data-campo="costo">${tr.costo}$</td>
                    <td><button class = "delete-tratamiento-tabla"> Eliminar </button>
                        <button class = "update-tratamiento-tabla"> Modificar </button></td>`
                tbody.appendChild(fila); }
    catch (error) {
        console.error("Error al crear el tratamiento:", error);
        alert("Hubo un error al crear el tratamiento. Por favor, inténtalo de nuevo.");
    }
});

const crearVoucherForm = document.getElementById("crear-voucher-form");
crearVoucherForm.addEventListener("submit", async (event) => {
    event.preventDefault(); // Evita que la pagina se recargue al enviar el formulario
    
    const formData = new FormData(crearVoucherForm);
    const voucherData = {
        regalador_id: Number(formData.get("regalador_id")),
        tratamiento_id: Number(formData.get("tratamiento_id")),
        cliente_id: Number(formData.get("cliente_id")),
    };
    
    try {
        const response = await fetch("/voucher", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(voucherData)
        });
        if (!response.ok) {
            throw new Error(response.text());
        }
        
        alert("Voucher creado exitosamente");
    

        const tbody = document.getElementById("cuerpoTablaVoucher")
        const v = await response.json()
                const fila =
                document.createElement("tr")
                    fila.dataset.id = v.id
                    
                    fila.innerHTML = `<td> ${v.id_voucher}</td>
                    <td data-campo="regalador">${v.regalador_id}</td>
                    <td data-campo="tratamiento">${v.tratamiento_id}</td>
                    <td data-campo="cliente">${v.cliente_id}$</td>
                    <td><button class = "delete-voucher-tabla"> Eliminar </button>
                        <button class = "update-voucher-tabla"> Modificar </button></td>`
                tbody.appendChild(fila); }
    catch (error) {
        console.error("Error al crear el voucher:", error);
        alert("Hubo un error al crear el voucher. Por favor, inténtalo de nuevo.");
    }
});

const crearTurnoForm = document.getElementById("crear-turno-form");
crearTurnoForm.addEventListener("submit", async (event) => {
    event.preventDefault(); // Evita que la pagina se recargue al enviar el formulario
    
    const formData = new FormData(crearTurnoForm);
    const turnoData = {
        fecha: formData.get("fecha"),
        hora: Number(formData.get("hora")),
        cliente_id: Number(formData.get("cliente_id")),
        tratamiento_id: Number(formData.get("tratamiento_id")),
    };
    
    try {
        const response = await fetch("/turno", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(turnoData)
        });
        
        if (!response.ok) {
            throw new Error(response.text());
        }
        
        alert("Turno creado exitosamente");
    

        const tbody = document.getElementById("cuerpoTablaTurno")
        const tu = await response.json()
                const fila =
                document.createElement("tr")
                    fila.dataset.fecha = tu.fecha
                    fila.dataset.hora = tu.hora
                    
                    fila.innerHTML = `<td data-campo="fecha"> ${tu.fecha}</td>
                    <td data-campo="hora">${tu.hora}</td>
                    <td data-campo="nombre_tratamiento">${tu.nombre_tratamiento}</td>
                    <td data-campo="nombre_cliente">${tu.nombre_cliente}</td>
                    <td><button class = "delete-turno-tabla"> Eliminar </button>
                        <button class = "update-turno-tabla"> Modificar </button></td>`
                tbody.appendChild(fila); }
    catch (error) {
        console.error("Error al crear el turno:", error);
        alert("Hubo un error al crear el turno. Por favor, inténtalo de nuevo.");
    }
});

// CLIENTES
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

//  Actualizar cliente
clientes.addEventListener("click", (e) => {
  const boton = e.target.closest(".update-cliente-tabla");
  if (!boton) return;

  const fila = boton.closest("tr");

  // Convertir cada celda editable en un input con su contenido actual
  fila.querySelectorAll("td[data-campo]").forEach((td) => {
    let valorActual = td.textContent;
  td.dataset.original = valorActual; // para Cancelar

  const campo = td.dataset.campo;
  const input = document.createElement("input");
  input.name = campo;
  input.placeholder = campo.charAt(0).toUpperCase() + campo.slice(1);
  input.style.marginBottom = "5px"; 
  if (campo == "deuda")
        valorActual = valorActual.replace("$","")    
  input.value = valorActual;
  input.style.width = "100%";
  input.style.maxWidth = "200px";
  input.style.boxSizing = "border-box";
  

  td.textContent = "";
  td.appendChild(input);
  });


  const btnConfirmar = document.createElement("button");
  btnConfirmar.textContent = "Confirmar";
  btnConfirmar.className = "confirmar-cliente-tabla";

  const btnCancelar = document.createElement("button");
  btnCancelar.textContent = "Cancelar";
  btnCancelar.className = "cancelar-cliente-tabla";

  boton.replaceWith(btnConfirmar, btnCancelar);
});

//  Cancelar modificación clientes
clientes.addEventListener("click", (e) => {
  const boton = e.target.closest(".cancelar-cliente-tabla");
  if (!boton) return;

  const fila = boton.closest("tr");

  
  fila.querySelectorAll("td[data-campo]").forEach((td) => {
    td.textContent = td.dataset.original;
    delete td.dataset.original;
  });

  
  const btnModificar = document.createElement("button");
  btnModificar.textContent = "Modificar";
  btnModificar.className = "update-cliente-tabla";

  const btnConfirmar = fila.querySelector(".confirmar-cliente-tabla");
  btnConfirmar.replaceWith(btnModificar);
  boton.remove();
});

//  Confirmar modificación cliente
clientes.addEventListener("click", async (e) => {
  const boton = e.target.closest(".confirmar-cliente-tabla");
  if (!boton) return;

  const fila = boton.closest("tr");

  const datos = { id: fila.dataset.id };

  fila.querySelectorAll("td[data-campo]").forEach((td) => {
    const input = td.querySelector("input");
    datos[td.dataset.campo] = input.value;
  });

  
  id = datos.id
  datos.deuda = parseFloat(datos.deuda)
  datos.nro_telefono = Number(datos.nro_telefono)
  clienteData = {nombre: datos.nombre, apellido:datos.apellido, deuda: datos.deuda, nro_telefono:datos.nro_telefono}
     try {
        const response = await fetch(`/cliente/${datos.id}` ,{
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(clienteData)
        });
        
        if (!response.ok) {
            throw new Error(response.text());
        }
        
        fila.querySelectorAll("td[data-campo]").forEach((td) => {
            if(td.dataset.campo == "deuda"){
                datos[td.dataset.campo]+="$"
            }
            td.textContent = datos[td.dataset.campo];
            delete td.dataset.original;
        });
        const btnModificar = document.createElement("button");
        btnModificar.textContent = "Modificar";
        btnModificar.className = "update-cliente-tabla";

        const btnCancelar = fila.querySelector(".cancelar-cliente-tabla");
        boton.replaceWith(btnModificar);
        btnCancelar.remove();
        alert("Cliente actualizado exitosamente");
     } catch (error){
        console.error("Error al actualizar el cliente: ", error)
     }
});

//  TRATAMIENTOS
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

//  Actualizar tratamiento
tratamientos.addEventListener("click", (e) => {
  const boton = e.target.closest(".update-tratamiento-tabla");
  if (!boton) return;

  const fila = boton.closest("tr");

  // Convertir cada celda editable en un input con su contenido actual
  fila.querySelectorAll("td[data-campo]").forEach((td) => {
    let valorActual = td.textContent;
  td.dataset.original = valorActual; // para Cancelar

  const campo = td.dataset.campo;
  const input = document.createElement("input");
  input.name = campo;
  input.placeholder = campo.charAt(0).toUpperCase() + campo.slice(1);
  input.style.marginBottom = "5px";
  if (campo == "costo")
        valorActual = valorActual.replace("$","")    
  input.value = valorActual;
  input.style.width = "100%";
  input.style.maxWidth = "200px";
  input.style.boxSizing = "border-box";

  td.textContent = "";
  td.appendChild(input);
  });


  const btnConfirmar = document.createElement("button");
  btnConfirmar.textContent = "Confirmar";
  btnConfirmar.className = "confirmar-tratamiento-tabla";

  const btnCancelar = document.createElement("button");
  btnCancelar.textContent = "Cancelar";
  btnCancelar.className = "cancelar-tratamiento-tabla";

  boton.replaceWith(btnConfirmar, btnCancelar);
});

//  Cancelar modificación
tratamientos.addEventListener("click", (e) => {
  const boton = e.target.closest(".cancelar-tratamiento-tabla");
  if (!boton) return;

  const fila = boton.closest("tr");

  
  fila.querySelectorAll("td[data-campo]").forEach((td) => {
    td.textContent = td.dataset.original;
    delete td.dataset.original;
  });

  
  const btnModificar = document.createElement("button");
  btnModificar.textContent = "Modificar";
  btnModificar.className = "update-tratamiento-tabla";

  const btnConfirmar = fila.querySelector(".confirmar-tratamiento-tabla");
  btnConfirmar.replaceWith(btnModificar);
  boton.remove();
});

//  Confirmar modificación
tratamientos.addEventListener("click", async (e) => {
  const boton = e.target.closest(".confirmar-tratamiento-tabla");
  if (!boton) return;

  const fila = boton.closest("tr");

  const datos = { id: fila.dataset.id };

  fila.querySelectorAll("td[data-campo]").forEach((td) => {
    const input = td.querySelector("input");
    datos[td.dataset.campo] = input.value;
  });

  
  id = datos.id
  datos.costo = Number(datos.costo)
  tratamientoData = {nombre: datos.nombre, descripcion_corta: datos.descripcion_corta, costo:datos.costo}
     try {
        const response = await fetch(`/tratamiento/${datos.id}` ,{
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(tratamientoData)
        });
        
        if (!response.ok) {
            throw new Error(response.text());
        }
        
        fila.querySelectorAll("td[data-campo]").forEach((td) => {
            if(td.dataset.campo == "costo"){
                datos[td.dataset.campo] += "$"
            }
            td.textContent = datos[td.dataset.campo];
            delete td.dataset.original;
        });
        const btnModificar = document.createElement("button");
        btnModificar.textContent = "Modificar";
        btnModificar.className = "update-tratamiento-tabla";

        const btnCancelar = fila.querySelector(".cancelar-tratamiento-tabla");
        boton.replaceWith(btnModificar);
        btnCancelar.remove();
        alert("Tratamiento actualizado exitosamente");
     } catch (error){
        console.error("Error al actualizar el tratamiento: ", error)
     }
});

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

const turnos = document.getElementById("cuerpoTablaTurno")
turnos.addEventListener("click",async (e) =>{
    e.preventDefault();
    const boton = e.target.closest(".delete-turno-tabla");
    if(!boton) return;

    const fila = boton.closest("tr");
    const fecha= fila.dataset.fecha;
    const hora= fila.dataset.hora;
    try {
        const res = await fetch(`/turno/${fecha}&${hora}`,{
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            }
        });
        if (!res.ok) {
            throw new Error(res.text())
        }
        turnos.removeChild(fila)
    } catch (error){
        console.error("Error al borrar turno: ", error)
    }
})
